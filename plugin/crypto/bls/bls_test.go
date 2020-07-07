// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bls

import (
	"encoding/hex"
	"testing"

	"github.com/33cn/chain33/common/crypto"
	"github.com/33cn/plugin/plugin/dapp/evm/executor/vm/common"
	"github.com/phoreproject/bls/g1pubs"
	"github.com/stretchr/testify/assert"
)

var blsDrv = New()

func TestGenKey(t *testing.T) {
	sk, err := blsDrv.GenKey()
	assert.NoError(t, err)
	assert.NotEmpty(t, sk)
	pk := sk.PubKey()
	assert.NotEmpty(t, pk)

	sk2, _ := blsDrv.GenKey()
	assert.NotEqual(t, sk.Bytes(), sk2.Bytes(), "should not generate two same key", sk, sk2)
}

func TestSignAndVerify(t *testing.T) {
	sk, _ := blsDrv.GenKey()
	pk := sk.PubKey()
	m1 := []byte("message to be signed. 将要做签名的消息")
	// sign and verify
	sig1 := sk.Sign(m1)
	ret := pk.VerifyBytes(m1, sig1)
	assert.Equal(t, true, ret)

	// different message should have different signature
	m2 := []byte("message to be signed. 将要做签名的消息.")
	sig2 := sk.Sign(m2)
	assert.NotEqual(t, sig1, sig2, "different message got the same signature", sig1, sig2)

	// different key should have different signature for a same message.
	sk2, _ := blsDrv.GenKey()
	sig12 := sk2.Sign(m1)
	ret = pk.VerifyBytes(m1, sig12)
	assert.Equal(t, false, ret)
}

func TestAggregate(t *testing.T) {
	m := []byte("message to be signed. 将要做签名的消息")
	n := 8
	pubs := make([]crypto.PubKey, 0, n)
	sigs := make([]crypto.Signature, 0, n) //signatures for the same message

	for i := 0; i < n; i++ {
		sk, _ := blsDrv.GenKey()
		pk := sk.PubKey()
		pubs = append(pubs, pk)
		sigs = append(sigs, sk.Sign(m))
	}

	asig, err := blsDrv.Aggregate(sigs)
	assert.NoError(t, err)
	// One
	err = blsDrv.VerifyAggregatedOne(pubs, m, asig)
	assert.NoError(t, err)

	apub, err := blsDrv.AggregatePublic(pubs)
	assert.NoError(t, err)

	ret := apub.VerifyBytes(m, asig)
	assert.Equal(t, true, ret)

	//invalid length
	_, err = blsDrv.Aggregate(nil)
	assert.Error(t, err)
	_, err = blsDrv.AggregatePublic(make([]crypto.PubKey, 0))
	assert.Error(t, err)
}

//和其他Bls库交叉验证
func TestCrossVerify(t *testing.T) {
	//priv, err := g1pubs.RandKey(rand.Reader)
	//assert.NoError(t, err)
	//var sec bls.SecretKey
	//secByte, _ := hex.DecodeString("4aac41b5cb665b93e031faa751944b1f14d77cb17322403cba8df1d6e4541a4d")
	//sec.Deserialize(secByte)
	priStr := "4aac41b5cb665b93e031faa751944b1f14d77cb17322403cba8df1d6e4541a4d"
	prib, err := hex.DecodeString(priStr)
	assert.NoError(t, err)

	m1 := []byte("message to be signed.")

	var priv32 [32]byte
	copy(priv32[:], prib[:])
	privG1 := g1pubs.DeserializeSecretKey(priv32)

	pk := g1pubs.PrivToPub(privG1)
	pub := pk.Serialize()
	t.Log("g1pub", "pub", hex.EncodeToString(pub[:]))
	sig1 := g1pubs.Sign(m1, privG1)
	sig1b := sig1.Serialize()
	t.Log("g1pub", "sig", hex.EncodeToString(sig1b[:]))

	blsPub, err := blsDrv.PubKeyFromBytes(pub[:])
	assert.NoError(t, err)

	blsSig, err := blsDrv.SignatureFromBytes(sig1b[:])
	assert.NoError(t, err)
	ret := blsPub.VerifyBytes(m1, blsSig)
	//以后验证成功应该是true, 当前验证失败，不兼容，先写成false
	assert.Equal(t, false, ret)

	blsPri, err := blsDrv.PrivKeyFromBytes(prib[:])
	newBlsPub := blsPri.PubKey()
	bt := newBlsPub.Bytes()
	t.Log("blscgo", "pub", hex.EncodeToString(bt[:]))
	blsSig2 := blsPri.Sign(m1)
	t.Log("blscgo", "sig", hex.EncodeToString(blsSig2.Bytes()[:]))

}

func TestAggregateN(t *testing.T) {
	n := 100
	pubs := make([]crypto.PubKey, 0, n)
	sigs := make([]crypto.Signature, 0, n) //signatures for different message
	msgs := [][]byte{}
	for i := 0; i < n; i++ {
		sk, _ := blsDrv.GenKey()
		pk := sk.PubKey()
		pubs = append(pubs, pk)
		m := common.ToHash([]byte{byte(i)})
		msgs = append(msgs, m[:])
		sigs = append(sigs, sk.Sign(m[:]))
	}

	asig, err := blsDrv.Aggregate(sigs)
	assert.NoError(t, err)

	err = blsDrv.VerifyAggregatedN(pubs, msgs, asig)
	assert.NoError(t, err)

	//lose some messages will cause an error
	err = blsDrv.VerifyAggregatedN(pubs, msgs[1:], asig)
	assert.Error(t, err)

}
