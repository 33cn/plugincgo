package state

import (
	"testing"

	"github.com/33cn/chain33/common/address"
	"github.com/stretchr/testify/assert"
)

func Test_createAccountChange(t *testing.T) {
	env := setupTestEnv()
	exectorName := "user.jvm.Dice"
	memoryStateDB := NewMemoryStateDB(exectorName, env.stateDB, env.localDB, env.base.GetCoinsAccount(), 10)

	addr := address.ExecAddress(exectorName)
	memoryStateDB.CreateAccount(addr, opener, exectorName)

	var createAccChange createAccountChange
	createAccChange.account = addr
	kv := createAccChange.getData(memoryStateDB)
	assert.Equal(t, 1, len(kv))
}

func Test_codeChange(t *testing.T) {
	env := setupTestEnv()
	exectorName := "user.jvm.Dice"
	memoryStateDB := NewMemoryStateDB(exectorName, env.stateDB, env.localDB, env.base.GetCoinsAccount(), 10)

	addr := address.ExecAddress(exectorName)
	memoryStateDB.CreateAccount(addr, opener, exectorName)

	memoryStateDB.SetCodeAndAbi(addr, []byte{1, 2}, nil)
	var codeChg codeChange
	codeChg.account = addr
	kv := codeChg.getData(memoryStateDB)
	assert.Equal(t, 1, len(kv))
}

func Test_localStorageChange(t *testing.T) {
	env := setupTestEnv()
	exectorName := "user.jvm.Dice"
	memoryStateDB := NewMemoryStateDB(exectorName, env.stateDB, env.localDB, env.base.GetCoinsAccount(), 10)

	addr := address.ExecAddress(exectorName)
	memoryStateDB.CreateAccount(addr, opener, exectorName)

	memoryStateDB.SetCodeAndAbi(addr, []byte{1, 2}, nil)
	var localStorageChg localStorageChange
	localStorageChg.account = addr
	kv := localStorageChg.getData(memoryStateDB)
	assert.Equal(t, 0, len(kv))

	logs := localStorageChg.getLog(memoryStateDB)
	assert.Equal(t, 1, len(logs))
}
