// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package messages

import (
	"bytes"
	"fmt"
	"time"

	"github.com/FactomProject/factomd/common/constants"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
)

//A placeholder structure for messages
type AuditServerFault struct {
	MessageBase
	Timestamp interfaces.Timestamp

	Signature interfaces.IFullSignature

	//Not marshalled
	hash interfaces.IHash
}

var _ interfaces.IMsg = (*AuditServerFault)(nil)
var _ Signable = (*AuditServerFault)(nil)

func (m *AuditServerFault) GetRepeatHash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer messagesAuditServerFaultGetRepeatHash.Observe(float64(time.Now().UnixNano() - callTime))
	return m.GetMsgHash()
}

func (a *AuditServerFault) IsSameAs(b *AuditServerFault) bool {
	callTime := time.Now().UnixNano()
	defer messagesAuditServerFaultIsSameAs.Observe(float64(time.Now().UnixNano() - callTime))
	if b == nil {
		return false
	}
	if a.Timestamp.GetTimeMilli() != b.Timestamp.GetTimeMilli() {
		return false
	}

	//TODO: expand

	if a.Signature == nil && b.Signature != nil {
		return false
	}
	if a.Signature != nil {
		if a.Signature.IsSameAs(b.Signature) == false {
			return false
		}
	}

	return true
}

func (m *AuditServerFault) Sign(key interfaces.Signer) error {
	callTime := time.Now().UnixNano()
	defer messagesAuditServerFaultSign.Observe(float64(time.Now().UnixNano() - callTime))
	signature, err := SignSignable(m, key)
	if err != nil {
		return err
	}
	m.Signature = signature
	return nil
}

func (m *AuditServerFault) GetSignature() interfaces.IFullSignature {
	callTime := time.Now().UnixNano()
	defer messagesAuditServerFaultGetSignature.Observe(float64(time.Now().UnixNano() - callTime))
	return m.Signature
}

func (m *AuditServerFault) VerifySignature() (bool, error) {
	callTime := time.Now().UnixNano()
	defer messagesAuditServerFaultVerifySignature.Observe(float64(time.Now().UnixNano() - callTime))
	return VerifyMessage(m)
}

func (e *AuditServerFault) Process(uint32, interfaces.IState) bool {
	callTime := time.Now().UnixNano()
	defer messagesAuditServerFaultProcess.Observe(float64(time.Now().UnixNano() - callTime))
	panic("AuditServerFault object should never have its Process() method called")
}

func (m *AuditServerFault) GetTimestamp() interfaces.Timestamp {
	callTime := time.Now().UnixNano()
	defer messagesAuditServerFaultGetTimestamp.Observe(float64(time.Now().UnixNano() - callTime))
	return m.Timestamp
}

func (m *AuditServerFault) GetHash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer messagesAuditServerFaultGetHash.Observe(float64(time.Now().UnixNano() - callTime))

	return nil
}

func (m *AuditServerFault) GetMsgHash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer messagesAuditServerFaultGetMsgHash.Observe(float64(time.Now().UnixNano() - callTime))
	if m.MsgHash == nil {
		data, err := m.MarshalBinary()
		if err != nil {
			return nil
		}
		m.MsgHash = primitives.Sha(data)
	}
	return m.MsgHash
}

func (m *AuditServerFault) Type() byte {
	callTime := time.Now().UnixNano()
	defer messagesAuditServerFaultType.Observe(float64(time.Now().UnixNano() - callTime))
	return constants.AUDIT_SERVER_FAULT_MSG
}

func (m *AuditServerFault) Int() int {
	callTime := time.Now().UnixNano()
	defer messagesAuditServerFaultInt.Observe(float64(time.Now().UnixNano() - callTime))
	return -1
}

func (m *AuditServerFault) Bytes() []byte {
	callTime := time.Now().UnixNano()
	defer messagesAuditServerFaultBytes.Observe(float64(time.Now().UnixNano() - callTime))
	return nil
}

func (m *AuditServerFault) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	callTime := time.Now().UnixNano()
	defer messagesAuditServerFaultUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error unmarshalling AuditServerFault: %v", r)
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

	//TODO: expand

	if len(newData) > 0 {
		m.Signature = new(primitives.Signature)
		newData, err = m.Signature.UnmarshalBinaryData(newData)
		if err != nil {
			return nil, err
		}
	}

	return newData, nil
}

func (m *AuditServerFault) UnmarshalBinary(data []byte) error {
	callTime := time.Now().UnixNano()
	defer messagesAuditServerFaultUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	_, err := m.UnmarshalBinaryData(data)
	return err
}

func (m *AuditServerFault) MarshalForSignature() (data []byte, err error) {
	callTime := time.Now().UnixNano()
	defer messagesAuditServerFaultMarshalForSignature.Observe(float64(time.Now().UnixNano() - callTime))
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

func (m *AuditServerFault) MarshalBinary() (data []byte, err error) {
	callTime := time.Now().UnixNano()
	defer messagesAuditServerFaultMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
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

func (m *AuditServerFault) String() string {
	callTime := time.Now().UnixNano()
	defer messagesAuditServerFault.Observe(float64(time.Now().UnixNano() - callTime))
	return "AuditFault"
}

func (m *AuditServerFault) DBHeight() int {
	callTime := time.Now().UnixNano()
	defer messagesAuditServerFaultDBHeight.Observe(float64(time.Now().UnixNano() - callTime))
	return 0
}

func (m *AuditServerFault) ChainID() []byte {
	callTime := time.Now().UnixNano()
	defer messagesAuditServerFaultChainID.Observe(float64(time.Now().UnixNano() - callTime))
	return nil
}

func (m *AuditServerFault) ListHeight() int {
	callTime := time.Now().UnixNano()
	defer messagesAuditServerFaultListHeight.Observe(float64(time.Now().UnixNano() - callTime))
	return 0
}

func (m *AuditServerFault) SerialHash() []byte {
	callTime := time.Now().UnixNano()
	defer messagesAuditServerFaultSerialHash.Observe(float64(time.Now().UnixNano() - callTime))
	return nil
}

// Validate the message, given the state.  Three possible results:
//  < 0 -- Message is invalid.  Discard
//  0   -- Cannot tell if message is Valid
//  1   -- Message is valid
func (m *AuditServerFault) Validate(state interfaces.IState) int {
	callTime := time.Now().UnixNano()
	defer messagesAuditServerFaultValidate.Observe(float64(time.Now().UnixNano() - callTime))
	return 0
}

// Returns true if this is a message for this server to execute as
// a leader.
func (m *AuditServerFault) ComputeVMIndex(state interfaces.IState) {

}

// Execute the leader functions of the given message
func (m *AuditServerFault) LeaderExecute(state interfaces.IState) {

}

func (m *AuditServerFault) FollowerExecute(interfaces.IState) {
}

func (e *AuditServerFault) JSONByte() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer messagesAuditServerFaultJSONByte.Observe(float64(time.Now().UnixNano() - callTime))
	return primitives.EncodeJSON(e)
}

func (e *AuditServerFault) JSONString() (string, error) {
	callTime := time.Now().UnixNano()
	defer messagesAuditServerFaultJSONString.Observe(float64(time.Now().UnixNano() - callTime))
	return primitives.EncodeJSONString(e)
}

func (e *AuditServerFault) JSONBuffer(b *bytes.Buffer) error {
	callTime := time.Now().UnixNano()
	defer messagesAuditServerFaultJSONBuffer.Observe(float64(time.Now().UnixNano() - callTime))
	return primitives.EncodeJSONToBuffer(e, b)
}
