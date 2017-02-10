// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package primitives

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/FactomProject/ed25519"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives/random"
	"time"
)

// PrivateKey contains Public/Private key pair
type PrivateKey struct {
	Key *[ed25519.PrivateKeySize]byte
	Pub *PublicKey
}

var _ interfaces.Signer = (*PrivateKey)(nil)

func (e *PrivateKey) Init() {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesPrivateKeyInit.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if e.Key == nil {
		e.Key = new([ed25519.PrivateKeySize]byte)
	}
	if e.Pub == nil {
		e.Pub = new(PublicKey)
	}
}

func RandomPrivateKey() interfaces.Signer {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesRandomPrivateKey.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return NewPrivateKeyFromHexBytes(random.RandByteSliceOfLen(ed25519.PrivateKeySize))
}

func (pk *PrivateKey) CustomMarshalText2(string) ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesPrivateKeyCustomMarshalText2.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return ([]byte)(hex.EncodeToString(pk.Key[:]) + pk.Pub.String()), nil
}

func (pk *PrivateKey) Public() []byte {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesPrivateKeyPublic.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return pk.Pub[:]
}

func (pk *PrivateKey) AllocateNew() {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesPrivateKeyAllocateNew.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	pk.Key = new([ed25519.PrivateKeySize]byte)
	pk.Pub = new(PublicKey)
}

// Create a new private key from a hex string
func NewPrivateKeyFromHex(s string) (*PrivateKey, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesNewPrivateKeyFromHex.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	privKeybytes, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}
	if privKeybytes == nil {
		return nil, errors.New("Invalid private key input string!")
	}
	pk := new(PrivateKey)
	if len(privKeybytes) == ed25519.PrivateKeySize-ed25519.PublicKeySize {
		_, privKeybytes, err = GenerateKeyFromPrivateKey(privKeybytes)
		if err != nil {
			return nil, err
		}
	}
	if len(privKeybytes) != ed25519.PrivateKeySize {
		return nil, errors.New("Invalid private key input string!")
	}
	pk.AllocateNew()
	copy(pk.Key[:], privKeybytes)
	err = pk.Pub.UnmarshalBinary(privKeybytes[len(privKeybytes)-ed25519.PublicKeySize:])
	if err != nil {
		return nil, err
	}
	return pk, nil
}

func NewPrivateKeyFromHexBytes(privKeybytes []byte) *PrivateKey {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesNewPrivateKeyFromHexBytes.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	pk := new(PrivateKey)
	pk.AllocateNew()
	copy(pk.Key[:], privKeybytes)
	pk.Pub.UnmarshalBinary(ed25519.GetPublicKey(pk.Key)[:])
	return pk
}

// Sign signs msg with PrivateKey and return Signature
func (pk *PrivateKey) Sign(msg []byte) (sig interfaces.IFullSignature) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesPrivateKeySign.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	sig = new(Signature)
	sig.SetPub(pk.Pub[:])
	s := ed25519.Sign(pk.Key, msg)
	sig.SetSignature(s[:])
	return
}

// Sign signs msg with PrivateKey and return Signature
func (pk *PrivateKey) MarshalSign(msg interfaces.BinaryMarshallable) (sig interfaces.IFullSignature) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesPrivateKeyMarshalSign.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	data, _ := msg.MarshalBinary()
	return pk.Sign(data)
}

//Generate creates new PrivateKey / PublciKey pair or returns error
func (pk *PrivateKey) GenerateKey() error {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesPrivateKeyGenerateKey.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return err
	}
	pk.Key = priv
	pk.Pub = new(PublicKey)
	err = pk.Pub.UnmarshalBinary(pub[:])
	return err
}

// Returns hex-encoded string of first 32 bytes of key (private key portion)
func (pk *PrivateKey) PrivateKeyString() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesPrivateKeyPrivateKeyString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return hex.EncodeToString(pk.Key[:32])
}
func (pk *PrivateKey) PublicKeyString() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesPrivateKeyPublicKeyString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return pk.Pub.String()
}

/******************PublicKey*******************************/

// PublicKey contains only Public part of Public/Private key pair
type PublicKey [ed25519.PublicKeySize]byte

var _ interfaces.Verifier = (*PublicKey)(nil)

func (c *PublicKey) Copy() (*PublicKey, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesPublicKeyCopy.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	h := new(PublicKey)
	bytes, err := c.MarshalBinary()
	if err != nil {
		return nil, err
	}
	err = h.UnmarshalBinary(bytes)
	if err != nil {
		return nil, err
	}
	return h, nil
}

func (a PublicKey) Fixed() [ed25519.PublicKeySize]byte {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesPublicKeyFixed.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return a
}

func (a *PublicKey) IsSameAs(b *PublicKey) bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesPublicKeyIsSameAs.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if b == nil {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func (pk *PublicKey) MarshalText() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesPublicKeyMarshalText.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return []byte(pk.String()), nil
}

func (pk *PublicKey) UnmarshalText(b []byte) error {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesPublicKeyUnmarshalText.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	p, err := hex.DecodeString(string(b))
	if err != nil {
		return err
	}
	copy(pk[:], p)
	return nil
}

func (pk *PublicKey) String() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesPublicKeyString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return hex.EncodeToString(pk[:])
}

func PubKeyFromString(instr string) (pk PublicKey) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesPubKeyFromString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	p, _ := hex.DecodeString(instr)
	copy(pk[:], p)
	return
}

func (k *PublicKey) Verify(msg []byte, sig *[ed25519.SignatureSize]byte) bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesPublicKeyVerify.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return ed25519.VerifyCanonical((*[32]byte)(k), msg, sig)
}

func (k *PublicKey) MarshalBinary() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesPublicKeyMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	var buf Buffer
	buf.Write(k[:])
	return buf.DeepCopyBytes(), nil
}

func (k *PublicKey) UnmarshalBinaryData(p []byte) ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesPublicKeyUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if len(p) < ed25519.PublicKeySize {
		return nil, fmt.Errorf("Invalid data passed")
	}
	copy(k[:], p)
	return p[ed25519.PublicKeySize:], nil
}

func (k *PublicKey) UnmarshalBinary(p []byte) (err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesPublicKeyUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	_, err = k.UnmarshalBinaryData(p)
	return
}

// Verify returns true iff sig is a valid signature of message by publicKey.
func Verify(publicKey *[ed25519.PublicKeySize]byte, message []byte, sig *[ed25519.SignatureSize]byte) bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesVerify.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return ed25519.VerifyCanonical(publicKey, message, sig)
}

// Verify returns true iff sig is a valid signature of message by publicKey.
func VerifySlice(p []byte, message []byte, s []byte) bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesVerifySlice.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	sig := new([ed25519.PrivateKeySize]byte)
	pub := new([ed25519.PublicKeySize]byte)
	copy(sig[:], s)
	copy(pub[:], p)
	return Verify(pub, message, sig)
}
