// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package messages

import (
	"fmt"

	"time"

	"github.com/FactomProject/factomd/common/constants"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
)

//A placeholder structure for messages
type SignatureTimeout struct {
	MessageBase
	Timestamp interfaces.Timestamp

	Signature interfaces.IFullSignature

	//Not marshalled
	hash interfaces.IHash
}

var _ interfaces.IMsg = (*SignatureTimeout)(nil)
var _ Signable = (*SignatureTimeout)(nil)

func (a *SignatureTimeout) IsSameAs(b *SignatureTimeout) bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesSignatureTimeoutIsSameAs.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if b == nil {
		return false
	}
	if a.Timestamp.GetTimeMilli() != b.Timestamp.GetTimeMilli() {
		return false
	}

	if a.Signature == nil && b.Signature != nil {
		return false
	}
	if a.Signature != nil {
		if a.Signature.IsSameAs(b.Signature) == false {
			return false
		}
	}
	//TODO: expand

	return true
}

func (m *SignatureTimeout) Process(uint32, interfaces.IState) bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesSignatureTimeoutProcess.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////
	return true
}

func (m *SignatureTimeout) GetRepeatHash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesSignatureTimeoutGetRepeatHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return m.GetMsgHash()
}

func (m *SignatureTimeout) GetHash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesSignatureTimeoutGetHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if m.hash == nil {
		data, err := m.MarshalForSignature()
		if err != nil {
			panic(fmt.Sprintf("Error in CommitChain.GetHash(): %s", err.Error()))
		}
		m.hash = primitives.Sha(data)
	}
	return m.hash
}

func (m *SignatureTimeout) GetMsgHash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesSignatureTimeoutGetMsgHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if m.MsgHash == nil {
		data, err := m.MarshalBinary()
		if err != nil {
			return nil
		}
		m.MsgHash = primitives.Sha(data)
	}
	return m.MsgHash
}

func (m *SignatureTimeout) GetTimestamp() interfaces.Timestamp {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesSignatureTimeoutGetTimestamp.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return m.Timestamp
}

func (m *SignatureTimeout) Type() byte {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesSignatureTimeoutType.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return constants.SIGNATURE_TIMEOUT_MSG
}

func (m *SignatureTimeout) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesSignatureTimeoutUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error unmarshalling: %v", r)
		}
	}()
	newData = data
	if newData[0] != m.Type() {
		return nil, fmt.Errorf("Invalid Message type")
	}
	newData = newData[1:]

	m.Timestamp = new(primitives.Timestamp)
	newData, err = m.Timestamp.UnmarshalBinaryData(newData)
	if err != nil {
		return nil, err
	}

	if len(newData) > 0 {
		m.Signature = new(primitives.Signature)
		newData, err = m.Signature.UnmarshalBinaryData(newData)
		if err != nil {
			return nil, err
		}
	}

	return newData, nil
}

func (m *SignatureTimeout) UnmarshalBinary(data []byte) error {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesSignatureTimeoutUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	_, err := m.UnmarshalBinaryData(data)
	return err
}

func (m *SignatureTimeout) MarshalForSignature() (data []byte, err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesSignatureTimeoutMarshalForSignature.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	var buf primitives.Buffer
	buf.Write([]byte{m.Type()})
	if d, err := m.Timestamp.MarshalBinary(); err != nil {
		return nil, err
	} else {
		buf.Write(d)
	}

	//TODO: expand

	return buf.DeepCopyBytes(), nil
}
func (m *SignatureTimeout) MarshalBinary() (data []byte, err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesSignatureTimeoutMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	resp, err := m.MarshalForSignature()
	if err != nil {
		return nil, err
	}
	sig := m.GetSignature()

	if sig != nil {
		sigBytes, err := sig.MarshalBinary()
		if err != nil {
			return nil, err
		}
		return append(resp, sigBytes...), nil
	}
	return resp, nil
}

func (m *SignatureTimeout) GetSignature() interfaces.IFullSignature {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesSignatureTimeoutGetSignature.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return m.Signature
}

func (m *SignatureTimeout) VerifySignature() (bool, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesSignatureTimeoutVerifySignature.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return VerifyMessage(m)
}

func (m *SignatureTimeout) Sign(key interfaces.Signer) error {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesSignatureTimeoutSign.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	signature, err := SignSignable(m, key)
	if err != nil {
		return err
	}
	m.Signature = signature
	return nil
}

func (m *SignatureTimeout) String() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesSignatureTimeoutString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return "Signature Timeout"
}

func (m *SignatureTimeout) DBHeight() int {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesSignatureTimeoutDBHeight.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return 0
}

func (m *SignatureTimeout) ChainID() []byte {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesSignatureTimeoutChainID.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return nil
}

func (m *SignatureTimeout) ListHeight() int {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesSignatureTimeoutListHeight.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return 0
}

func (m *SignatureTimeout) SerialHash() []byte {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesSignatureTimeoutSerialHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return nil
}

// Validate the message, given the state.  Three possible results:
//  < 0 -- Message is invalid.  Discard
//  0   -- Cannot tell if message is Valid
//  1   -- Message is valid
func (m *SignatureTimeout) Validate(state interfaces.IState) int {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesSignatureTimeoutValidate.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return 0
}

func (m *SignatureTimeout) ComputeVMIndex(state interfaces.IState) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesSignatureTimeoutComputeVMIndex.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

}

func (m *SignatureTimeout) LeaderExecute(state interfaces.IState) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesSignatureTimeoutLeaderExecute.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

}

func (m *SignatureTimeout) FollowerExecute(interfaces.IState) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesSignatureTimeoutFollowerExecute.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

}

func (e *SignatureTimeout) JSONByte() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesSignatureTimeoutJSONByte.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSON(e)
}

func (e *SignatureTimeout) JSONString() (string, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesSignatureTimeoutJSONString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSONString(e)
}
