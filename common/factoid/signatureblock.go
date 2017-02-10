// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package factoid

import (
	"fmt"

	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
	"time"
)

/**************************************
 * ISign
 *
 * Interface for RCB Signatures
 *
 * The signature block holds the signatures that validate one of the RCBs.
 * Each signature has an index, so if the RCD is a multisig, you can know
 * how to apply the signatures to the addresses in the RCD.
 **************************************/

type SignatureBlock struct {
	Signatures []interfaces.ISignature
}

var _ interfaces.ISignatureBlock = (*SignatureBlock)(nil)

func (b *SignatureBlock) IsSameAs(s interfaces.ISignatureBlock) bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidSignatureBlockIsSameAs.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if s == nil {
		return b == nil
	}

	sigs := s.GetSignatures()
	if len(b.Signatures) != len(sigs) {
		return false
	}
	for i := range b.Signatures {
		if b.Signatures[i].IsSameAs(sigs[i]) == false {
			return false
		}
	}

	return true
}

func (b SignatureBlock) UnmarshalBinary(data []byte) error {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidSignatureBlockUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	_, err := b.UnmarshalBinaryData(data)
	return err
}

func (e *SignatureBlock) JSONByte() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidSignatureBlockJSONByte.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSON(e)
}

func (e *SignatureBlock) JSONString() (string, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidSignatureBlockJSONString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSONString(e)
}

func (b SignatureBlock) String() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidSignatureBlockString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	txt, err := b.CustomMarshalText()
	if err != nil {
		return "<error>"
	}
	return string(txt)
}

func (s *SignatureBlock) AddSignature(sig interfaces.ISignature) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidSignatureBlockAddSignature.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if len(s.Signatures) > 0 {
		s.Signatures[0] = sig
	} else {
		s.Signatures = append(s.Signatures, sig)
	}
}

func (s SignatureBlock) GetSignature(index int) interfaces.ISignature {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidSignatureBlockGetSignature.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if len(s.Signatures) <= index {
		return nil
	}
	return s.Signatures[index]
}

func (s SignatureBlock) GetSignatures() []interfaces.ISignature {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidSignatureBlockGetSignatures.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if s.Signatures == nil {
		s.Signatures = make([]interfaces.ISignature, 1, 1)
		s.Signatures[0] = new(FactoidSignature)
	}
	return s.Signatures
}

func (a SignatureBlock) MarshalBinary() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidSignatureBlockMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	var out primitives.Buffer

	for _, sig := range a.GetSignatures() {
		data, err := sig.MarshalBinary()
		if err != nil {
			return nil, fmt.Errorf("Signature failed to Marshal in RCD_1")
		}
		out.Write(data)
	}

	return out.DeepCopyBytes(), nil
}

func (s SignatureBlock) CustomMarshalText() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidSignatureBlockCustomMarshalText.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	var out primitives.Buffer

	out.WriteString("Signature Block: \n")
	for _, sig := range s.Signatures {
		out.WriteString(" signature: ")
		txt, err := sig.CustomMarshalText()
		if err != nil {
			return nil, err
		}
		out.Write(txt)
		out.WriteString("\n ")

	}

	return out.DeepCopyBytes(), nil
}

func (s *SignatureBlock) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidSignatureBlockUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	s.Signatures = make([]interfaces.ISignature, 1)
	s.Signatures[0] = new(FactoidSignature)
	data, err = s.Signatures[0].UnmarshalBinaryData(data)
	if err != nil {
		return nil, fmt.Errorf("Failure to unmarshal Signature")
	}

	return data, nil
}

func NewSingleSignatureBlock(priv, data []byte) *SignatureBlock {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidNewSingleSignatureBlock.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	s := new(SignatureBlock)
	s.AddSignature(NewED25519Signature(priv, data))
	return s
}
