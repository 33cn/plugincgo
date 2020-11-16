package state

import (
	"github.com/33cn/chain33/common/address"
	"github.com/33cn/plugincgo/plugin/dapp/jvm/types"
	"testing"
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
    key := "Test_GetState"
    value := []byte{1,2}
	err := contractAccount.SetState(key, value)
	assert.Equal(t, nil, err)
	result := contractAccount.GetState(key)
	assert.Equal(t, value, result)

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
	value := []byte{1,2}
	err := contractAccount.SetValue2Local(key, value, "")
	assert.Equal(t, types.ErrSetLocalNotAllowed, err)
}





