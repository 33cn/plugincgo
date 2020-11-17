package state

import (
	"fmt"
	"testing"

	"github.com/33cn/chain33/common/address"
	"github.com/33cn/plugincgo/plugin/dapp/jvm/types"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

func Test_NewContractAccount(t *testing.T) {
	contractAccount := NewContractAccount("", nil)
	if nil != contractAccount {
		assert.Equal(t, nil, contractAccount)
	}
}

func Test_GetState_SetState(t *testing.T) {
	env := setupTestEnv()
	exectorName := "user.jvm.Dice"
	contractAddr := address.ExecAddress(exectorName)
	memoryStateDB := NewMemoryStateDB(exectorName, env.stateDB, env.localDB, env.base.GetCoinsAccount(), 10)
	contractAccount := NewContractAccount(contractAddr, memoryStateDB)
	memoryStateDB.accounts[contractAddr] = contractAccount
	key := "Test_GetState"
	value := []byte{1, 2}
	err := contractAccount.SetState(key, value)
	assert.Equal(t, nil, err)
	result := contractAccount.GetState(key)
	assert.Equal(t, value, result)

	var storageChg storageChange
	storageChg.account = contractAddr
	storageChg.key = []byte(key)
	kv := storageChg.getData(memoryStateDB)
	assert.Equal(t, 1, len(kv))

	delete(contractAccount.stateCache, key)
	result = contractAccount.GetState(key)
	assert.Equal(t, value, result)

	result = contractAccount.GetState("key")
	assert.Equal(t, 0, len(result))
}

func Test_SetValue2Local(t *testing.T) {
	env := setupTestEnv()
	exectorName := "user.jvm.Dice"
	contractAddr := address.ExecAddress(exectorName)
	memoryStateDB := NewMemoryStateDB(exectorName, env.stateDB, env.localDB, env.base.GetCoinsAccount(), 10)
	contractAccount := NewContractAccount(contractAddr, memoryStateDB)
	key := "Test_GetState"
	value := []byte{1, 2}
	err := contractAccount.SetValue2Local(key, value, "")
	assert.Equal(t, types.ErrSetLocalNotAllowed, err)
}

func Test_restoreData(t *testing.T) {
	env := setupTestEnv()
	exectorName := "user.jvm.Dice"
	contractAddr := address.ExecAddress(exectorName)
	memoryStateDB := NewMemoryStateDB(exectorName, env.stateDB, env.localDB, env.base.GetCoinsAccount(), 10)
	contractAccount := NewContractAccount(contractAddr, memoryStateDB)

	var content types.JVMContractData
	content.Name = exectorName
	content.Addr = contractAddr
	content.Code = []byte{1, 2}
	data, _ := proto.Marshal(&content)
	contractAccount.restoreData(data)
	assert.Equal(t, 1, int(contractAccount.Data.Code[0]))
	assert.Equal(t, 2, int(contractAccount.Data.Code[1]))
	assert.Equal(t, exectorName, contractAccount.GetExecName())
	log := contractAccount.BuildDataLog()
	assert.Equal(t, types.TyLogContractDataJvm, int(log.Ty))
	var logJVMContractData types.LogJVMContractData
	_ = proto.Unmarshal(log.Log, &logJVMContractData)
	assert.Equal(t, exectorName, logJVMContractData.Name)

	var contentNew types.JVMContractData
	contractAccount.Data = contentNew
	contractAccount.restoreData([]byte{1, 2})
	assert.Equal(t, 0, len(contractAccount.Data.Code))
}

func Test_GetDataKV(t *testing.T) {
	env := setupTestEnv()
	exectorName := "user.jvm.Dice"
	contractAddr := address.ExecAddress(exectorName)
	memoryStateDB := NewMemoryStateDB(exectorName, env.stateDB, env.localDB, env.base.GetCoinsAccount(), 10)
	contractAccount := NewContractAccount(contractAddr, memoryStateDB)

	kvs := contractAccount.GetDataKV()
	assert.Equal(t, 1, len(kvs))
}

func Test_GetStateKey(t *testing.T) {
	env := setupTestEnv()
	exectorName := "user.jvm.Dice"
	contractAddr := address.ExecAddress(exectorName)
	memoryStateDB := NewMemoryStateDB(exectorName, env.stateDB, env.localDB, env.base.GetCoinsAccount(), 10)
	contractAccount := NewContractAccount(contractAddr, memoryStateDB)

	key := contractAccount.GetStateKey()
	keyLocal := []byte("mavl-" + exectorName + "-state: " + contractAddr)
	assert.Equal(t, keyLocal, key)

	localKey := "local"
	IsPara = true
	expect := fmt.Sprintf(string([]byte("LODB"))+"-"+Title+exectorName+"-data-%v:%v", contractAddr, localKey)
	localDataKey := contractAccount.GetLocalDataKey(contractAddr, localKey)
	assert.Equal(t, expect, localDataKey)
}
