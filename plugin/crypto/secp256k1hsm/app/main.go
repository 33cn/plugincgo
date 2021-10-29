package main

import (
	"fmt"
	"math/big"
	"time"

	"github.com/33cn/chain33/common"
	"github.com/33cn/chain33/common/crypto"
	"github.com/33cn/chain33/system/crypto/sm2"
	"github.com/33cn/plugincgo/plugin/crypto/secp256k1hsm/adapter"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	gmsm_sm2 "github.com/tjfoc/gmsm/sm2"
)

func verifySM2Signature(rBytes, sBytes, msg []byte) bool {
	xBytes, _ := common.FromHex("0000000000000000000000000000000000000000000000000000000000000000FD4241057FEC6CBEEC501F7E1763751B8F6DFCFB910FB634FBB76A16639EF172")
	yBytes, _ := common.FromHex("00000000000000000000000000000000000000000000000000000000000000001C6DA89F9C1A5EE9B6108E5A2A5FE336962630A34DBA1AF428451E1CE63BB3CF")
	x := new(big.Int).SetBytes(xBytes)
	y := new(big.Int).SetBytes(yBytes)

	publicKey := &gmsm_sm2.PublicKey{
		X: x,
		Y: y,
	}
	var pubSM2 sm2.PubKeySM2
	copy(pubSM2[:], gmsm_sm2.Compress(publicKey))

	r := new(big.Int).SetBytes(rBytes)
	s := new(big.Int).SetBytes(sBytes)
	signature := sm2.SignatureSM2(sm2.Serialize(r, s))

	return pubSM2.VerifyBytes(msg, signature)
}

func main() {
	if err := adapter.OpenHSMSession(); nil != err {
		panic("Failed to OpenHSMSession")
	}
	fmt.Println("Succeed to OpenHSMSession")
	fmt.Println("   ")
	fmt.Println("   ")
	fmt.Println("   ")

	passwd := "a1234567"
	//for keyIndex := 2; keyIndex <= 2; keyIndex++ {
	    keyIndex := 2
		if err := adapter.GetPrivateKeyAccessRight(passwd, keyIndex); nil != err {
			panic("Failed to GetPrivateKeyAccessRight")
		}

		for i := 0; i < 20; i++ {
			time.Sleep(time.Millisecond*1000)
			verifySecp256k1(keyIndex)
		}


		if err := adapter.ReleaeAccessRight(keyIndex); nil != err {
			panic("Failed to GetPrivateKeyAccessRight")
		}
	//}
	adapter.CloseHSMSession()
}

func verifySecp256k1(keyIndex int) {
	msg, _ := common.FromHex("456789")
	r, s, err := adapter.SignSecp256k1(msg, keyIndex)
	if err != nil {
		panic("Failed to SignSecp256k1 due to:" + err.Error())
	}
	fmt.Println("   ")
	fmt.Println("   ")
	fmt.Println("   ")
	fmt.Println(" keyIndex is ", keyIndex)
	fmt.Println("signature R=", common.ToHex(r))
	fmt.Println("signature S=", common.ToHex(s))
	hash := crypto.Sha256(msg)

	sig := adapter.MakeRSVsignature(r, s)
	fmt.Println(" sig ", common.ToHex(sig))

	pubRecoverd, err := ethCrypto.Ecrecover(hash[:], sig)
	fmt.Println(" pubRecoverd is ", common.ToHex(pubRecoverd))
	secpPubKey, err := ethCrypto.UnmarshalPubkey(pubRecoverd)
	if nil != err {
		panic("ethCrypto.UnmarshalPubkey failed")
	}
	recoveredAddr := ethCrypto.PubkeyToAddress(*secpPubKey)
	fmt.Println(" recoveredAddr is ", recoveredAddr.String())
}

func verifySM2() {
	//msg := []byte("112233445566112233445566112233445566112233445566")
	msg, _ := common.FromHex("112233445566112233445566112233445566112233445566")
	r, s, err := adapter.SignSM2Internal(msg, 10)
	if err != nil {
		panic("Failed to SignSM2Internal due to:" + err.Error())
	}
	fmt.Println("signature R=", common.ToHex(r))
	fmt.Println("signature S=", common.ToHex(s))

	///////构建公钥////////
	xBytes, _ := common.FromHex("0000000000000000000000000000000000000000000000000000000000000000FD4241057FEC6CBEEC501F7E1763751B8F6DFCFB910FB634FBB76A16639EF172")
	yBytes, _ := common.FromHex("00000000000000000000000000000000000000000000000000000000000000001C6DA89F9C1A5EE9B6108E5A2A5FE336962630A34DBA1AF428451E1CE63BB3CF")
	x := new(big.Int).SetBytes(xBytes)
	y := new(big.Int).SetBytes(yBytes)

	publicKey := &gmsm_sm2.PublicKey{
		X: x,
		Y: y,
	}
	var pubSM2 sm2.PubKeySM2
	copy(pubSM2[:], gmsm_sm2.Compress(publicKey))

	///////开始循环验证签名////////
	now := time.Now()
	now.Nanosecond()

	msLater := now.Nanosecond()
	fmt.Printf("msLater = %d\n", msLater)
	time.Sleep(time.Millisecond)
	msLater = now.Nanosecond()
	fmt.Printf("msLater = %d\n", time.Now().Nanosecond())
	fmt.Printf("msLater Sec = %d\n", time.Now().Second())
	for i := 0; i < 10*1000; i++ {
		adapter.SignSM2Internal(msg, 10)
		//rBytes, sBytes, _ := adapter.SignSM2Internal(msg, 10)
		//fmt.Println("rBytes = ", common.ToHex(rBytes))
		//fmt.Println("sBytes = ", common.ToHex(sBytes))
		//r := new(big.Int).SetBytes(rBytes)
		//s := new(big.Int).SetBytes(sBytes)
		//signature := sm2.SignatureSM2(sm2.Serialize(r, s))
		//if !pubSM2.VerifyBytes(msg, signature) {
		//	panic("Failed to do VerifyBytes")
		//}
		//fmt.Println("Succeed to do VerifyBytes for times = ", i)
	}

	fmt.Println("      ")
	fmt.Printf("testLater = %d\n", time.Now().Nanosecond())
	fmt.Printf("testLater sec = %d\n", time.Now().Second())
	fmt.Println("      ")
	fmt.Println(" ^-^ Successful ^-^  ")
}
