package testHelper

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/FactomProject/ed25519"
	"github.com/FactomProject/factomd/common/factoid"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
	"time"
)

func NewPrivKeyString(n uint64) string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdtestHelperNewPrivKeyString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	buf := new(primitives.Buffer)
	if err := binary.Write(buf, binary.BigEndian, n); err != nil {
		panic(err)
	}

	priv := fmt.Sprintf("000000000000000000000000000000000000000000000000%x", buf.DeepCopyBytes())
	return priv
}

//Create 32 bit private key (without the public key part)
func NewPrivKey(n uint64) []byte {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdtestHelperNewPrivKey.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	priv := NewPrivKeyString(n)
	p, err := hex.DecodeString(priv)
	if err != nil {
		panic(err)
	}
	return p
}

//Create a full 64 bit key holding both private and public key
func NewFullPrivKey(n uint64) []byte {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdtestHelperNewFullPrivKey.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	priv := NewPrivKey(n)
	pub := PrivateKeyToEDPub(priv)
	return append(priv, pub...)
}

func NewPrimitivesPrivateKey(n uint64) *primitives.PrivateKey {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdtestHelperNewPrimitivesPrivateKey.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.NewPrivateKeyFromHexBytes(NewFullPrivKey(n))
}

func NewFactoidAddressStrings(n uint64) (string, string, string) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdtestHelperNewFactoidAddressStrings.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	//ec9f1cefa00406b80d46135a53504f1f4182d4c0f3fed6cca9281bc020eff973
	//000000000000000000000000000000000000000000000000XXXXXXXXXXXXXXXX
	priv := NewPrivKeyString(n)
	privKey, pubKey, add, err := factoid.PrivateKeyStringToEverythingString(priv)
	if err != nil {
		panic(err)
	}
	return privKey, pubKey, add
}

func NewFactoidAddress(n uint64) interfaces.IAddress {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdtestHelperNewFactoidAddress.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	_, pub, _ := NewFactoidAddressStrings(n)
	add, err := factoid.PublicKeyStringToFactoidAddress(pub)
	if err != nil {
		panic(err)
	}
	return add
}

func NewFactoidRCDAddressString(n uint64) string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdtestHelperNewFactoidRCDAddressString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	add := NewFactoidAddress(n)
	return add.String()
}

func NewFactoidRCDAddress(n uint64) interfaces.IRCD {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdtestHelperNewFactoidRCDAddress.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	_, pub, _ := NewFactoidAddressStrings(n)
	add, err := factoid.PublicKeyStringToFactoidRCDAddress(pub)
	if err != nil {
		panic(err)
	}
	return add
}

func NewECAddress(n uint64) interfaces.IAddress {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdtestHelperNewECAddress.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	_, pub, _ := NewFactoidAddressStrings(n)
	add, err := factoid.PublicKeyStringToECAddress(pub)
	if err != nil {
		panic(err)
	}
	return add
}

func NewECAddressPublicKeyString(n uint64) string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdtestHelperNewECAddressPublicKeyString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	add := NewECAddress(n)
	return add.String()
}

func NewECAddressString(n uint64) string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdtestHelperNewECAddressString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	pub := NewECAddressPublicKeyString(n)
	ecAddress, err := factoid.PublicKeyStringToECAddressString(pub)
	if err != nil {
		panic(err)
	}
	return ecAddress
}

func PrivateKeyToEDPub(priv []byte) []byte {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdtestHelperPrivateKeyToEDPub.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	priv2 := new([ed25519.PrivateKeySize]byte)
	copy(priv2[:], priv)
	pub := ed25519.GetPublicKey(priv2)
	return pub[:]
}
