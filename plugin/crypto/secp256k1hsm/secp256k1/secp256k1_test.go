package secp256k1

import (
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"testing"

	goSh256 "crypto/sha256"
	"gotest.tools/assert"

	"github.com/33cn/chain33/common/crypto"

	"github.com/33cn/chain33/system/crypto/secp256k1"
	"github.com/33cn/plugin/plugin/dapp/evm/executor/vm/common"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	btcSecp256k1 "github.com/btcsuite/btcd/btcec"
)

//secp256k1签名DER编码格式
// 0x30 <length> 0x02
//<length r> r
//0x02 <length s>
//s

//30440220
//7C12FF568B6DA03EF2CD5681EE45EDF846172771AC6F9369B50FEBC95B1CF68F
//0220
//8B6916A3CC7423D9044E77ABECE410B69B7BD82C22ECCC76061B4BE79141D1A3
func Test_VerifySecp256k1SigFromTass_forChain33(t *testing.T) {
	require := require.New(t)

	//var pubKey secp256k1.PubKeySecp256k1

	pubBytes := common.FromHex("04C24FBA65F8CD81223D2935EDEA663048A1BEFB5A78BC67C80DCB5A1D601F898C35EA242D2E76CACE9EE5A61DBDA29A5076707325FE20B5A80DB0CA6D02C5D983")

	secpPubKey, err := ethCrypto.UnmarshalPubkey(pubBytes)
	require.Equal(nil, err)
	pub33Bytes := ethCrypto.CompressPubkey(secpPubKey)

	c := &secp256k1.Driver{}
	pubKey, err := c.PubKeyFromBytes(pub33Bytes)
	require.Equal(nil, err)

	//msg := []byte("12345678123456781234567812345678")
	msg := []byte("456789")

	hash := crypto.Sha256(msg)
	fmt.Println("hash = ", common.Bytes2Hex(hash))
	//0xfed9efbd5a8ef6820d639dbcb831daf9d6308312cc73d6188beb54a9a148e29a

	sig, err := c.SignatureFromBytes(common.FromHex("304502207C12FF568B6DA03EF2CD5681EE45EDF846172771AC6F9369B50FEBC95B1CF68F0221008B6916A3CC7423D9044E77ABECE410B69B7BD82C22ECCC76061B4BE79141D1A3"))
	//sig, err := c.SignatureFromBytes(common.FromHex("304402207C12FF568B6DA03EF2CD5681EE45EDF846172771AC6F9369B50FEBC95B1CF68F02208B6916A3CC7423D9044E77ABECE410B69B7BD82C22ECCC76061B4BE79141D1A3"))
	require.Equal(nil, err)

	result := pubKey.VerifyBytes(msg, sig)
	require.Equal(true, result)

	privateKeySlice := common.FromHex("300B155F751964276C0536230BD9B16FE7A86533C3CBAA7575E8D0431DBEDF23")
	privateKey, err := c.PrivKeyFromBytes(privateKeySlice)
	require.Equal(nil, err)
	sig2 := privateKey.Sign(msg)
	fmt.Println("sig2 = ", common.Bytes2Hex(sig2.Bytes()))
	sig3 := privateKey.Sign(msg)
	fmt.Println("sig3 = ", common.Bytes2Hex(sig3.Bytes()))
}

//在以太坊上的验证签名的有效性
//注意：从加密中导出的签名信息中RS信息中的首字节必须大于０，否则签名验证失败
func Test_Verify4Eth(t *testing.T) {
	pub := common.FromHex("04C24FBA65F8CD81223D2935EDEA663048A1BEFB5A78BC67C80DCB5A1D601F898C35EA242D2E76CACE9EE5A61DBDA29A5076707325FE20B5A80DB0CA6D02C5D983")
	sig := common.FromHex("2F2F8EF10E6C9075CAB44DE3C4F904817220537C1E7DCFADD502C03F14F5B3974C405EA9BB189B85F15B91C82CE5D6191D66238ECCCE83FA8F8FF83173F1586F00")

	msg := []byte("456789")
	hash := crypto.Sha256(msg)
	fmt.Println("hash  = ", common.Bytes2Hex(hash))
    hash1 := goSh256.Sum256(msg)
	fmt.Println("hash1 = ", common.Bytes2Hex(hash1[:]))

	pubRecoverd, err := ethCrypto.Ecrecover(hash[:], sig)
	require.Equal(t, nil, err)
	fmt.Println("pubRecoverd = ", common.Bytes2Hex(pubRecoverd))

	VerifyResult := ethCrypto.VerifySignature(pub, hash[:], sig[:64])
	assert.Equal(t, true, VerifyResult)
}

func Test_secp256k1(t *testing.T) {
	require := require.New(t)

	c := &secp256k1.Driver{}

	priv, err := c.PrivKeyFromBytes(common.FromHex("CC38546E9E659D15E6B4893F0AB32A06D103931A8230B0BDE71459D2B27D6944"))
	require.Nil(err)
	t.Logf("priv:%X, len:%d", priv.Bytes(), len(priv.Bytes()))

	pub := priv.PubKey()
	require.NotNil(pub)
	t.Logf("pub:%X, len:%d", pub.Bytes(), len(pub.Bytes()))

	//msg := []byte("12345678123456781234567812345678")
	//msg := []byte("hello world")
	msg := []byte("456789")
	signature := priv.Sign(msg)
	t.Logf("sign:%X, len:%d", signature.Bytes(), len(signature.Bytes()))
	t.Logf("signature in hex format:%s", common.Bytes2Hex(signature.Bytes()))
	//0x3045022100f4009ab47dc32880b3e0bfad47885e9cfd1fd2228e804b38fb7f0f5ea6c02405022061422eb681fdd5078aa3971770cf22ce4ef12e9116995e4a3e141e23f5403014
	ok := pub.VerifyBytes(msg, signature)
	require.Equal(true, ok)
}

func Test_btcsecp256k1(t *testing.T) {
	msg := []byte("456789")
	priv, _ := btcSecp256k1.PrivKeyFromBytes(btcSecp256k1.S256(), common.FromHex("CC38546E9E659D15E6B4893F0AB32A06D103931A8230B0BDE71459D2B27D6944"))

	for i := 0; i < 10; i++ {
		sig, err := btcec.SignCompact(btcec.S256(), priv, crypto.Sha256(msg), true)
		assert.Equal(t, nil, err)

		pub, compressed, err := btcec.RecoverCompact(btcec.S256(), sig, crypto.Sha256(msg))
		assert.Equal(t, nil, err)
		fmt.Println("i is", i)
		fmt.Println("The recoverd pubkey is", common.Bytes2Hex(pub.SerializeCompressed()))
		fmt.Println("The compressed is", compressed)
		fmt.Println("   ")
		fmt.Println("   ")
		assert.Equal(t, "0x02504fa1c28caaf1d5a20fefb87c50a49724ff401043420cb3ba271997eb5a4387", common.Bytes2Hex(pub.SerializeCompressed()))
	}

	//rawerrInfo := common.FromHex("0x65766d3a20657865637574696f6e2072657665727465642c64657461696c3a2008c379a00000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000003f425459206465706f736974732072657175697265207468652027746f6b656e27206164647265737320746f20626520746865206e756c6c206164647265737300")
	//rawerrInfo := common.FromHex("0x3f425459206465706f736974732072657175697265207468652027746f6b656e27206164647265737320746f20626520746865206e756c6c206164647265737300")

	rawerrInfo := common.FromHex("0x65766d3a20657865637574696f6e2072657665727465642c64657461696c3a2008c379a0")
	errInfo := filterInvisibleChar(rawerrInfo)
	fmt.Println(string(errInfo))

	fmt.Println(string(rawerrInfo))
}

func filterInvisibleChar(in []byte) []byte{
	var out []byte
	for i:= 0; i < len(in); i++ {
		if in[i] < 32 || in[i] > 126 {
			continue
		}
		out = append(out, in[i])
	}

   return out
}

//签名值rsv：4E55CB10F11ECDD66807E303FA6A7797F9067D15AA2F4BC33BBA733BBF314B23324465CA8D3F89638756332CE3E556757277E158E90D1E39AA7D1014ABE331220000001C
//私钥：300B155F751964276C0536230BD9B16FE7A86533C3CBAA7575E8D0431DBEDF23
//公钥X：4C4D145791FB81AE5F5CC6B8290E12AB73818B1EAAA42A95C26F488DFCBD6887
//公钥y：976481BDEBB48B2796A72FCB2A48624AC33FE0B294529054B015BD1B537C6CDF
//签名数据（摘要）：1234567890123456123456789012345612345678901234561234567890123456

func Test_Ethsecp256k1(t *testing.T) {

	hash := common.FromHex("1234567890123456123456789012345612345678901234561234567890123456")
	sig := common.FromHex("4E55CB10F11ECDD66807E303FA6A7797F9067D15AA2F4BC33BBA733BBF314B23324465CA8D3F89638756332CE3E556757277E158E90D1E39AA7D1014ABE3312201")

	pubRecoverd, err := ethCrypto.Ecrecover(hash[:], sig)
	fmt.Println(" pubRecoverd is ", common.Bytes2Hex(pubRecoverd))
	secpPubKey, err := ethCrypto.UnmarshalPubkey(pubRecoverd)
	if nil != err {
		panic("ethCrypto.UnmarshalPubkey failed")
	}
	recoveredAddr := ethCrypto.PubkeyToAddress(*secpPubKey)
	fmt.Println(" recoveredAddr is ", recoveredAddr.String())

	privateKeySlice := common.FromHex("300B155F751964276C0536230BD9B16FE7A86533C3CBAA7575E8D0431DBEDF23")

	privateKey, err := ethCrypto.ToECDSA(privateKeySlice)
	assert.Equal(t, nil, err)

	calcAddr := ethCrypto.PubkeyToAddress(privateKey.PublicKey)
	fmt.Println(" calcAddr is ", calcAddr.String())

	pub2 := ethCrypto.FromECDSAPub(&privateKey.PublicKey)
	fmt.Println(" pub2 is ", common.Bytes2Hex(pub2))
}

func Test_chain33secp256k1(t *testing.T) {
	hash := common.FromHex("1234567890123456123456789012345612345678901234561234567890123456")
	sig := common.FromHex("4E55CB10F11ECDD66807E303FA6A7797F9067D15AA2F4BC33BBA733BBF314B23324465CA8D3F89638756332CE3E556757277E158E90D1E39AA7D1014ABE33122")
	pubKey := common.FromHex("044C4D145791FB81AE5F5CC6B8290E12AB73818B1EAAA42A95C26F488DFCBD6887976481BDEBB48B2796A72FCB2A48624AC33FE0B294529054B015BD1B537C6CDF")

	sigSecp256k1 := makeDERsignature(sig[:32], sig[32:])

	pub, err := btcSecp256k1.ParsePubKey(pubKey[:], btcSecp256k1.S256())
	assert.Equal(t, nil, err)

	sig2, err := btcSecp256k1.ParseDERSignature(sigSecp256k1[:], btcSecp256k1.S256())
	assert.Equal(t, nil, err)
	res := sig2.Verify(hash, pub)
	assert.Equal(t, true, res)

}

func makeDERsignature(rb, sb []byte) []byte {
	if rb[0] > 0x7F {
		rb = append([]byte{0}, rb...)
	}

	if sb[0] > 0x7F {
		sb = append([]byte{0}, sb...)
	}
	// total length of returned signature is 1 byte for each magic and
	// length (6 total), plus lengths of r and s
	length := 6 + len(rb) + len(sb)
	b := make([]byte, length)

	b[0] = 0x30
	b[1] = byte(length - 2)
	b[2] = 0x02
	b[3] = byte(len(rb))
	offset := copy(b[4:], rb) + 4
	b[offset] = 0x02
	b[offset+1] = byte(len(sb))
	copy(b[offset+2:], sb)
	return b
}


