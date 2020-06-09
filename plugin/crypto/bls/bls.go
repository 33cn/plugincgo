// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bls

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/33cn/chain33/common/crypto"
	"github.com/herumi/bls-eth-go-binary/bls"
)

const (
	BLSPrivateKeyLength = 32
	BLSPublicKeyLength  = 48
	BLSSignatureLength  = 96
	BLSMsgHashSize      = 32
)

// Driver driver
type Driver struct{}

func New() *Driver {
	bls.Init(bls.BLS12_381)
	//AggregateVerify len(msgVec) == 32 * len(pubVec) 需要这个设置
	bls.SetETHmode(bls.EthModeDraft07)
	return &Driver{}
}

// GenKey create private key
func (d Driver) GenKey() (k crypto.PrivKey, e error) {
	defer func() {
		if r := recover(); r != nil {
			k = nil
			e = errors.New("panic")
		}

	}()
	privKeyBytes := new([BLSPrivateKeyLength]byte)
	priv := bls.SecretKey{}
	priv.SetByCSPRNG()

	privBytes := priv.Serialize()
	copy(privKeyBytes[:], privBytes[:])
	return PrivKeyBLS(*privKeyBytes), nil
}

// PrivKeyFromBytes create private key from bytes
func (d Driver) PrivKeyFromBytes(b []byte) (privKey crypto.PrivKey, err error) {
	if len(b) != BLSPrivateKeyLength {
		return nil, errors.New("invalid bls priv key byte")
	}

	priv := &bls.SecretKey{}
	err = priv.Deserialize(b)
	if err != nil {
		return nil, err
	}
	privKeyBytes := new([BLSPrivateKeyLength]byte)
	privBytes := priv.Serialize()
	copy(privKeyBytes[:], privBytes[:])
	return PrivKeyBLS(*privKeyBytes), nil
}

// PubKeyFromBytes create public key from bytes
func (d Driver) PubKeyFromBytes(b []byte) (pubKey crypto.PubKey, err error) {
	if len(b) != BLSPublicKeyLength {
		return nil, errors.New("invalid bls pub key byte")
	}
	pubKeyBytes := new([BLSPublicKeyLength]byte)
	copy(pubKeyBytes[:], b[:])
	return PubKeyBLS(*pubKeyBytes), nil
}

// SignatureFromBytes create signature from bytes
func (d Driver) SignatureFromBytes(b []byte) (sig crypto.Signature, err error) {
	sigBytes := new([BLSSignatureLength]byte)
	copy(sigBytes[:], b[:])
	return SignatureBLS(*sigBytes), nil
}

func (d Driver) Aggregate(sigs []crypto.Signature) (crypto.Signature, error) {
	if len(sigs) == 0 {
		return nil, errors.New("no signatures to aggregate")
	}
	bSigs := make([]bls.Sign, 0, len(sigs))
	for i, sig := range sigs {
		bSig, err := ConvertToSignature(sig)
		if err != nil {
			return nil, fmt.Errorf("%v(index: %d)", err, i)
		}
		bSigs = append(bSigs, *bSig)
	}
	sign := &bls.Sign{}
	sign.Aggregate(bSigs)
	ret := new([BLSSignatureLength]byte)
	copy(ret[:], sign.Serialize())
	return SignatureBLS(*ret), nil
}

func (d Driver) AggregatePublic(pubs []crypto.PubKey) (crypto.PubKey, error) {
	if len(pubs) == 0 {
		return nil, errors.New("no public keys to aggregate")
	}
	//blank public key
	bPubs := &bls.PublicKey{}
	for i, pub := range pubs {
		bPub, err := ConvertToPublicKey(pub)
		if err != nil {
			return nil, fmt.Errorf("%v(index: %d)", err, i)
		}
		bPubs.Add(bPub)
	}
	ret := new([BLSPublicKeyLength]byte)
	copy(ret[:], bPubs.Serialize())
	return PubKeyBLS(*ret), nil
}

func (d Driver) VerifyAggregatedOne(pubs []crypto.PubKey, m []byte, sig crypto.Signature) error {
	bPubs := make([]bls.PublicKey, 0, len(pubs))
	for i, pub := range pubs {
		bPub, err := ConvertToPublicKey(pub)
		if err != nil {
			return fmt.Errorf("%v(index: %d)", err, i)
		}
		bPubs = append(bPubs, *bPub)
	}

	bSig, err := ConvertToSignature(sig)
	if err != nil {
		return err
	}

	if bSig.FastAggregateVerify(bPubs, m) {
		return nil
	}
	return errors.New("bls signature mismatch")
}

//由于cgo库的规则，需要每个ms都不大于32byte,或者后面
func (d Driver) VerifyAggregatedN(pubs []crypto.PubKey, ms [][]byte, sig crypto.Signature) error {
	bPubs := make([]bls.PublicKey, 0, len(pubs))
	for i, pub := range pubs {
		bPub, err := ConvertToPublicKey(pub)
		if err != nil {
			return fmt.Errorf("%v(index: %d)", err, i)
		}
		bPubs = append(bPubs, *bPub)
	}

	bSig, err := ConvertToSignature(sig)
	if err != nil {
		return err
	}

	if len(bPubs) != len(ms) {
		return fmt.Errorf("different length of pubs and messages, %d vs %d", len(bPubs), len(ms))
	}
	aggms := []byte{}
	for i, m := range ms {
		if len(m) != BLSMsgHashSize {
			return fmt.Errorf("the %d msg hash len not 32,len=%d", i, len(m))
		}
		aggms = append(aggms, m...)
	}

	if bSig.AggregateVerify(bPubs, aggms) {
		return nil
	}
	return errors.New("bls signature mismatch")
}

// PrivKeyBLS PrivKey
type PrivKeyBLS [BLSPrivateKeyLength]byte

// Bytes convert to bytes
func (privKey PrivKeyBLS) Bytes() []byte {
	s := make([]byte, BLSPrivateKeyLength)
	copy(s, privKey[:])
	return s
}

// Sign create signature
func (privKey PrivKeyBLS) Sign(msg []byte) crypto.Signature {
	k := bls.SecretKey{}
	k.Deserialize(privKey[:])

	sig := k.SignByte(msg).Serialize()
	s := new([BLSSignatureLength]byte)
	copy(s[:], sig[:])
	return SignatureBLS(*s)
}

// PubKey convert to public key
func (privKey PrivKeyBLS) PubKey() crypto.PubKey {
	k := bls.SecretKey{}
	err := k.Deserialize(privKey[:])
	if err != nil {
		return nil
	}
	pub := k.GetPublicKey().Serialize()
	s := new([BLSPublicKeyLength]byte)
	copy(s[:], pub[:])
	return PubKeyBLS(*s)
}

// Equals check privkey is equal
func (privKey PrivKeyBLS) Equals(other crypto.PrivKey) bool {
	if otherSecp, ok := other.(PrivKeyBLS); ok {
		return bytes.Equal(privKey[:], otherSecp[:])
	}
	return false
}

// String convert to string
func (privKey PrivKeyBLS) String() string {
	return fmt.Sprintf("PrivKeyBLS{*****}")
}

// PubKeyBLS PubKey
type PubKeyBLS [BLSPublicKeyLength]byte

// Bytes convert to bytes
func (pubKey PubKeyBLS) Bytes() []byte {
	s := make([]byte, BLSPublicKeyLength)
	copy(s, pubKey[:])
	return s
}

// VerifyBytes verify signature
func (pubKey PubKeyBLS) VerifyBytes(msg []byte, sig crypto.Signature) bool {
	pk := &bls.PublicKey{}
	err := pk.Deserialize(pubKey[:])
	if err != nil {
		fmt.Println("invalid bls pubkey")
		return false
	}

	g2sig, err := ConvertToSignature(sig)
	if err != nil {
		fmt.Println("ConvertToSignature fail:", err)
		return false
	}

	return g2sig.VerifyByte(pk, msg)
}

// String convert to string
func (pubKey PubKeyBLS) String() string {
	return fmt.Sprintf("PubKeyBLS{%X}", pubKey[:])
}

// KeyString Must return the full bytes in hex.
// Used for map keying, etc.
func (pubKey PubKeyBLS) KeyString() string {
	return fmt.Sprintf("%X", pubKey[:])
}

// Equals check public key is equal
func (pubKey PubKeyBLS) Equals(other crypto.PubKey) bool {
	if otherSecp, ok := other.(PubKeyBLS); ok {
		return bytes.Equal(pubKey[:], otherSecp[:])
	}
	return false
}

// ConvertToPublicKey convert to BLS PublicKey
func ConvertToPublicKey(pub crypto.PubKey) (*bls.PublicKey, error) {
	pubBLS, ok := pub.(PubKeyBLS)
	if !ok {
		return nil, errors.New("invalid bls public key")
	}
	p := &bls.PublicKey{}
	err := p.Deserialize(pubBLS[:])
	if err != nil {
		return nil, err
	}
	return p, nil
}

// SignatureBLS Signature
type SignatureBLS [BLSSignatureLength]byte

// SignatureS signature struct
type SignatureS struct {
	crypto.Signature
}

// Bytes convert signature to bytes
func (sig SignatureBLS) Bytes() []byte {
	s := make([]byte, len(sig))
	copy(s, sig[:])
	return s
}

// IsZero check signature is zero
func (sig SignatureBLS) IsZero() bool { return len(sig) == 0 }

// String convert signature to string
func (sig SignatureBLS) String() string {
	fingerprint := make([]byte, len(sig[:]))
	copy(fingerprint, sig[:])
	return fmt.Sprintf("/%X.../", fingerprint)

}

// Equals check signature equals
func (sig SignatureBLS) Equals(other crypto.Signature) bool {
	if otherEd, ok := other.(SignatureBLS); ok {
		return bytes.Equal(sig[:], otherEd[:])
	}
	return false
}

// ConvertToSignature convert to BLS Signature
func ConvertToSignature(sig crypto.Signature) (*bls.Sign, error) {
	// unwrap if needed
	if wrap, ok := sig.(SignatureS); ok {
		sig = wrap.Signature
	}
	sigBLS, ok := sig.(SignatureBLS)
	if !ok {
		return nil, errors.New("invalid bls signature")
	}
	s := bls.Sign{}
	err := s.Deserialize(sigBLS[:])
	if err != nil {
		return nil, err
	}
	return &s, nil
}

// Name name
const Name = "bls_cgo"

// ID id
const ID = 259

func init() {
	crypto.Register(Name, New())
	crypto.RegisterType(Name, ID)
}
