// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package factoid

import (
	"encoding/hex"
	"fmt"

	"github.com/FactomProject/factomd/common/constants"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
	"time"
)

// The default FactoidSignature doesn't care about indexing.  We will extend this
// FactoidSignature for multisig
type FactoidSignature struct {
	Signature [constants.SIGNATURE_LENGTH]byte // The FactoidSignature
}

var _ interfaces.ISignature = (*FactoidSignature)(nil)

func (s *FactoidSignature) IsSameAs(sig interfaces.ISignature) bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidFactoidSignatureIsSameAs.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.AreBytesEqual(s.Bytes(), sig.Bytes())
}

func (s *FactoidSignature) Verify([]byte) bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidFactoidSignatureVerify.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	fmt.Println("Verify is Broken")
	return true
}

func (sig *FactoidSignature) Bytes() []byte {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidFactoidSignatureBytes.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return sig.Signature[:]
}

func (s *FactoidSignature) GetKey() []byte {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidFactoidSignatureGetKey.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return s.Signature[32:]
}

func (h *FactoidSignature) MarshalText() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidFactoidSignatureMarshalText.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return []byte(hex.EncodeToString(h.Signature[:])), nil
}

func (s *FactoidSignature) JSONByte() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidFactoidSignatureJSONByte.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSON(s)
}

func (s *FactoidSignature) JSONString() (string, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidFactoidSignatureJSONString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSONString(s)
}

func (s FactoidSignature) String() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidFactoidSignatureString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	txt, err := s.CustomMarshalText()
	if err != nil {
		return "<error>"
	}
	return string(txt)
}

// Index is ignored.  We only have one FactoidSignature
func (s *FactoidSignature) SetSignature(sig []byte) error {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidFactoidSignatureSetSignature.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if len(sig) != constants.SIGNATURE_LENGTH {
		return fmt.Errorf("Bad FactoidSignature.  Should not happen")
	}
	copy(s.Signature[:], sig)
	return nil
}

func (s *FactoidSignature) GetSignature() *[constants.SIGNATURE_LENGTH]byte {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidFactoidSignatureGetSignature.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return &s.Signature
}

func (s FactoidSignature) MarshalBinary() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidFactoidSignatureMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	var out primitives.Buffer

	out.Write(s.Signature[:])

	return out.DeepCopyBytes(), nil
}

func (s FactoidSignature) CustomMarshalText() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidFactoidSignatureCustomMarshalText.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	var out primitives.Buffer

	out.WriteString(" FactoidSignature: ")
	out.WriteString(hex.EncodeToString(s.Signature[:]))
	out.WriteString("\n")

	return out.DeepCopyBytes(), nil
}

func (s *FactoidSignature) UnmarshalBinaryData(data []byte) ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidFactoidSignatureUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if data == nil || len(data) < constants.SIGNATURE_LENGTH {
		return nil, fmt.Errorf("Not enough data to unmarshal")
	}
	copy(s.Signature[:], data[:constants.SIGNATURE_LENGTH])
	return data[constants.SIGNATURE_LENGTH:], nil
}

func (s *FactoidSignature) UnmarshalBinary(data []byte) error {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidFactoidSignatureUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	_, err := s.UnmarshalBinaryData(data)
	return err
}

func NewED25519Signature(priv, data []byte) *FactoidSignature {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidNewED25519Signature.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	sig := primitives.Sign(priv, data)
	fs := new(FactoidSignature)
	copy(fs.Signature[:], sig[:constants.SIGNATURE_LENGTH])
	return fs
}
