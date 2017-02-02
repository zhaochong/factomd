// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package factoid

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/FactomProject/factomd/common/constants"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
)

// The default FactoidSignature doesn't care about indexing.  We will extend this
// FactoidSignature for multisig
type FactoidSignature struct {
	Signature [constants.SIGNATURE_LENGTH]byte // The FactoidSignature
}

var _ interfaces.ISignature = (*FactoidSignature)(nil)

func (s *FactoidSignature) Verify([]byte) bool {
	callTime := time.Now().UnixNano()
	defer factoidSignatureVerify.Observe(float64(time.Now().UnixNano() - callTime))
	fmt.Println("Verify is Broken")
	return true
}

func (sig *FactoidSignature) Bytes() []byte {
	callTime := time.Now().UnixNano()
	defer factoidSignatureBytes.Observe(float64(time.Now().UnixNano() - callTime))
	return sig.Signature[:]
}

func (s *FactoidSignature) GetKey() []byte {
	callTime := time.Now().UnixNano()
	defer factoidSignatureGetKey.Observe(float64(time.Now().UnixNano() - callTime))
	return s.Signature[32:]
}

func (s *FactoidSignature) GetHash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer factoidSignatureGetHash.Observe(float64(time.Now().UnixNano() - callTime))
	return nil
}

func (h *FactoidSignature) MarshalText() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer factoidSignatureMarshalText.Observe(float64(time.Now().UnixNano() - callTime))
	return []byte(hex.EncodeToString(h.Signature[:])), nil
}

func (s *FactoidSignature) JSONByte() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer factoidSignatureJSONByte.Observe(float64(time.Now().UnixNano() - callTime))
	return primitives.EncodeJSON(s)
}

func (s *FactoidSignature) JSONString() (string, error) {
	callTime := time.Now().UnixNano()
	defer factoidSignatureJSONString.Observe(float64(time.Now().UnixNano() - callTime))
	return primitives.EncodeJSONString(s)
}

func (s *FactoidSignature) JSONBuffer(b *bytes.Buffer) error {
	callTime := time.Now().UnixNano()
	defer factoidSignatureJSONBuffer.Observe(float64(time.Now().UnixNano() - callTime))
	return primitives.EncodeJSONToBuffer(s, b)
}

func (s FactoidSignature) String() string {
	callTime := time.Now().UnixNano()
	defer factoidSignatureString.Observe(float64(time.Now().UnixNano() - callTime))
	txt, err := s.CustomMarshalText()
	if err != nil {
		return "<error>"
	}
	return string(txt)
}

// Checks that the FactoidSignatures are the same.
func (s1 *FactoidSignature) IsEqual(sig interfaces.IBlock) []interfaces.IBlock {
	callTime := time.Now().UnixNano()
	defer factoidSignatureIsEqual.Observe(float64(time.Now().UnixNano() - callTime))
	s2, ok := sig.(*FactoidSignature)
	if !ok || // Not the right kind of interfaces.IBlock
		s1.Signature != s2.Signature { // Not the right rcd
		r := make([]interfaces.IBlock, 0, 5)
		return append(r, s1)
	}
	return nil
}

// Index is ignored.  We only have one FactoidSignature
func (s *FactoidSignature) SetSignature(sig []byte) error {
	callTime := time.Now().UnixNano()
	defer factoidSignatureSetSignature.Observe(float64(time.Now().UnixNano() - callTime))
	if len(sig) != constants.SIGNATURE_LENGTH {
		return fmt.Errorf("Bad FactoidSignature.  Should not happen")
	}
	copy(s.Signature[:], sig)
	return nil
}

func (s *FactoidSignature) GetSignature() *[constants.SIGNATURE_LENGTH]byte {
	callTime := time.Now().UnixNano()
	defer factoidSignatureGetSignature.Observe(float64(time.Now().UnixNano() - callTime))
	return &s.Signature
}

func (s FactoidSignature) MarshalBinary() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer factoidSignatureMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	var out primitives.Buffer

	out.Write(s.Signature[:])

	return out.DeepCopyBytes(), nil
}

func (s FactoidSignature) CustomMarshalText() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer factoidSignatureCustomMarshalText.Observe(float64(time.Now().UnixNano() - callTime))
	var out primitives.Buffer

	out.WriteString(" FactoidSignature: ")
	out.WriteString(hex.EncodeToString(s.Signature[:]))
	out.WriteString("\n")

	return out.DeepCopyBytes(), nil
}

func (s *FactoidSignature) UnmarshalBinaryData(data []byte) ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer factoidSignatureUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))
	copy(s.Signature[:], data[:constants.SIGNATURE_LENGTH])
	return data[constants.SIGNATURE_LENGTH:], nil
}

func (s *FactoidSignature) UnmarshalBinary(data []byte) error {
	callTime := time.Now().UnixNano()
	defer factoidSignatureUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	_, err := s.UnmarshalBinaryData(data)
	return err
}

func NewED25519Signature(priv, data []byte) *FactoidSignature {
	callTime := time.Now().UnixNano()
	defer factoidSignatureNewED25519Signature.Observe(float64(time.Now().UnixNano() - callTime))
	sig := primitives.Sign(priv, data)
	fs := new(FactoidSignature)
	copy(fs.Signature[:], sig[:constants.SIGNATURE_LENGTH])
	return fs
}
