package pbftlibbyz

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
	"testing"
	"time"

	"github.com/33cn/chain33/rpc/jsonclient"
	rpctypes "github.com/33cn/chain33/rpc/types"
	"github.com/33cn/chain33/types"
	"github.com/stretchr/testify/assert"
)

func rpc(method string, params, res interface{}) (e error) {
	jsonclient, err := jsonclient.NewJSONClient("http://127.0.0.1:5005")
	if err != nil {
		log.Println(err)
		return err
	}
	err = jsonclient.Call(method, params, res)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func TestPbft(t *testing.T) {
	// 构建docker镜像
	bi := exec.Command("/bin/sh", "-c", "./test/build-docker.sh")
	err := bi.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Docker build success!")

	//启动5个docker节点
	cmd := exec.Command("/bin/sh", "-c", "./test/run-docker.sh")
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		log.Fatal(err)
	}

	cmd.Start()
	reader := bufio.NewReader(stdout)
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Print(line)
	}
	cmd.Wait()

	time.Sleep(30 * time.Second)
	var res interface{}

	for i := 0; i < 20; i++ {
		rpc("Chain33.IsSync", new(types.ReqNil), &res)
		if res.(bool) {
			fmt.Println("********* IsSync! *********")
			break
		}
		time.Sleep(10 * time.Second)
	}

	fmt.Println("\n********* Get Peers Info! *********")
	err = rpc("Chain33.GetPeerInfo", new(types.ReqNil), &res)
	assert.Nil(t, err)

	fmt.Println("\n********** Gen Seed! ********")
	err = rpc("Chain33.GenSeed", &types.GenSeedLang{Lang: 0}, &res)
	assert.Nil(t, err)

	seed := res.(map[string]interface{})["seed"]
	fmt.Println("\n********** Save Seed! ********")
	err = rpc("Chain33.SaveSeed", &types.SaveSeedByPw{Seed: seed.(string), Passwd: "pwd"}, &res)
	assert.Nil(t, err)

	fmt.Println("\n********** UnLock Wallt! ********")
	err = rpc("Chain33.UnLock", &types.WalletUnLock{Passwd: "pwd", WalletOrTicket: false}, &res)
	assert.Nil(t, err)

	fmt.Println("\n********** Import PrivKey! ********")
	err = rpc("Chain33.ImportPrivkey", &types.ReqWalletImportPrivkey{Privkey: "CC38546E9E659D15E6B4893F0AB32A06D103931A8230B0BDE71459D2B27D6944", Label: "origin"}, &res)
	assert.Nil(t, err)
	origin := res.(map[string]interface{})["acc"].(map[string]interface{})["addr"].(string)

	// 初始化两个账户alex，bob来测试转账交易
	fmt.Println("\n********** Create New Account Alex! ********")
	err = rpc("Chain33.NewAccount", &types.ReqNewAccount{Label: "alex"}, &res)
	assert.Nil(t, err)
	alex := res.(map[string]interface{})["acc"].(map[string]interface{})["addr"].(string)

	fmt.Println("\n********** Create New Account Bob! ********")
	err = rpc("Chain33.NewAccount", &types.ReqNewAccount{Label: "bob"}, &res)
	assert.Nil(t, err)
	bob := res.(map[string]interface{})["acc"].(map[string]interface{})["addr"].(string)

	fmt.Println("\n********** Create Raw Transaction! ********")
	err = rpc("Chain33.CreateRawTransaction", &types.CreateTx{To: alex, Amount: 10000000000, IsToken: false, IsWithdraw: false}, &res)
	assert.Nil(t, err)

	fmt.Println("\n********** Sign Raw Tx! ********")
	err = rpc("Chain33.SignRawTx", &types.ReqSignRawTx{Addr: origin, TxHex: res.(string), Expire: "1h"}, &res)
	assert.Nil(t, err)

	fmt.Println("\n********** Send Transaction! ********")
	err = rpc("Chain33.SendTransaction", &rpctypes.RawParm{Data: res.(string)}, &res)
	assert.Nil(t, err)

	fmt.Println("\n********** Query Transaction! ********")
	hash := res.(string)

	for i := 0; i < 10; i++ {
		err := rpc("Chain33.QueryTransaction", &rpctypes.QueryParm{Hash: hash}, &res)
		if err == nil {
			break
		}
		time.Sleep(3 * time.Second)
	}

	fmt.Println("\n********** Create Raw Transaction! ********")
	err = rpc("Chain33.CreateRawTransaction", &types.CreateTx{To: bob, Amount: 10000000000, IsToken: false, IsWithdraw: false}, &res)
	assert.Nil(t, err)

	fmt.Println("\n********** Sign Raw Tx! ********")
	err = rpc("Chain33.SignRawTx", &types.ReqSignRawTx{Addr: origin, TxHex: res.(string), Expire: "1h", Index: 0}, &res)
	assert.Nil(t, err)

	fmt.Println("\n********** Send Transaction! ********")
	err = rpc("Chain33.SendTransaction", &rpctypes.RawParm{Data: res.(string)}, &res)
	assert.Nil(t, err)

	fmt.Println("\n********** Query Transaction! ********")
	hash = res.(string)

	for i := 0; i < 10; i++ {
		err := rpc("Chain33.QueryTransaction", &rpctypes.QueryParm{Hash: hash}, &res)
		if err == nil {
			break
		}
		time.Sleep(3 * time.Second)
	}

	fmt.Println("\n********** Get Accounts! ********")
	err = rpc("Chain33.GetAccounts", &types.ReqNil{}, &res)
	assert.Nil(t, err)

	// 开始十笔交易，alex把所有钱都转到bob账户
	fmt.Println("\n********* Test for ten transactions from alex to bob! *********")
	for i := 0; i < 10; i++ {
		rpc("Chain33.CreateRawTransaction", &types.CreateTx{To: bob, Amount: 1000000000, IsToken: false, IsWithdraw: false}, &res)

		rpc("Chain33.SignRawTx", &types.ReqSignRawTx{Addr: alex, TxHex: res.(string), Expire: "1h", Index: 0}, &res)

		rpc("Chain33.SendTransaction", &rpctypes.RawParm{Data: res.(string)}, &res)

		hash := res.(string)
		for i := 0; i < 10; i++ {
			err := rpc("Chain33.QueryTransaction", &rpctypes.QueryParm{Hash: hash}, &res)
			if err == nil {
				break
			}
			time.Sleep(3 * time.Second)
		}
		fmt.Println("\n********* Get tx, finish one transaction! *********")
	}
	// 显示所有账户的余额
	fmt.Println("\n********** Get Accounts! ********")
	err = rpc("Chain33.GetAccounts", &types.ReqNil{}, &res)
	assert.Nil(t, err)

	// 正确结果应该是bob余额20000000000，alex余额为0，会进行自动检查
	for _, value := range res.(map[string]interface{})["wallets"].([]interface{}) {
		if value.(map[string]interface{})["label"] == "alex" {
			assert.Equal(t, value.(map[string]interface{})["acc"].(map[string]interface{})["balance"].(float64), 0)
		}
		if value.(map[string]interface{})["label"] == "bob" {
			assert.Equal(t, value.(map[string]interface{})["acc"].(map[string]interface{})["balance"].(float64), 20000000000)
		}
	}
	fmt.Println("********* Balance Check Passed! *********")

	err = exec.Command("/bin/sh", "-c", "./test/rm-docker.sh").Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Docker rm success!")

}
