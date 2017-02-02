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
type RequestBlock struct {
	MessageBase
	Timestamp interfaces.Timestamp

	//TODO: figure whether this should be signed or not?

	//Not marshalled
	hash interfaces.IHash
}

var _ interfaces.IMsg = (*RequestBlock)(nil)

func (a *RequestBlock) IsSameAs(b *RequestBlock) bool {
	callTime := time.Now().UnixNano()
	defer messagesRequestBlockIsSameAs.Observe(float64(time.Now().UnixNano() - callTime))
	if b == nil {
		return false
	}
	if a.Timestamp.GetTimeMilli() != b.Timestamp.GetTimeMilli() {
		return false
	}

	//TODO: expand

	return true
}

func (m *RequestBlock) Process(uint32, interfaces.IState) bool { return true }

func (m *RequestBlock) GetRepeatHash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer messagesRequestBlockGetRepeatHash.Observe(float64(time.Now().UnixNano() - callTime))
	return m.GetMsgHash()
}

func (m *RequestBlock) GetHash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer messagesRequestBlockGetHash.Observe(float64(time.Now().UnixNano() - callTime))
	if m.hash == nil {
		data, err := m.MarshalForSignature()
		if err != nil {
			panic(fmt.Sprintf("Error in RequestBlock.GetHash(): %s", err.Error()))
		}
		m.hash = primitives.Sha(data)
	}
	return m.hash
}

func (m *RequestBlock) GetMsgHash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer messagesRequestBlockGetMsgHash.Observe(float64(time.Now().UnixNano() - callTime))
	if m.MsgHash == nil {
		data, err := m.MarshalBinary()
		if err != nil {
			return nil
		}
		m.MsgHash = primitives.Sha(data)
	}
	return m.MsgHash
}

func (m *RequestBlock) GetTimestamp() interfaces.Timestamp {
	callTime := time.Now().UnixNano()
	defer messagesRequestBlockGetTimestamp.Observe(float64(time.Now().UnixNano() - callTime))
	return m.Timestamp
}

func (m *RequestBlock) Type() byte {
	callTime := time.Now().UnixNano()
	defer messagesRequestBlockType.Observe(float64(time.Now().UnixNano() - callTime))
	return constants.REQUEST_BLOCK_MSG
}

func (m *RequestBlock) Int() int {
	return -1
}

func (m *RequestBlock) Bytes() []byte {
	return nil
}

func (m *RequestBlock) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	callTime := time.Now().UnixNano()
	defer messagesRequestBlockUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error unmarshalling RequestBlock: %v", r)
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

	return newData, nil
}

func (m *RequestBlock) UnmarshalBinary(data []byte) error {
	callTime := time.Now().UnixNano()
	defer messagesRequestBlockUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	_, err := m.UnmarshalBinaryData(data)
	return err
}

func (m *RequestBlock) MarshalForSignature() (data []byte, err error) {
	callTime := time.Now().UnixNano()
	defer messagesRequestBlockMarshalForSignature.Observe(float64(time.Now().UnixNano() - callTime))
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

func (m *RequestBlock) MarshalBinary() (data []byte, err error) {
	callTime := time.Now().UnixNano()
	defer messagesRequestBlockMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	//TODO: sign or delete
	return m.MarshalForSignature()
}

func (m *RequestBlock) String() string {
	callTime := time.Now().UnixNano()
	defer messagesRequestBlock.Observe(float64(time.Now().UnixNano() - callTime))
	return "Request Block"
}

func (m *RequestBlock) DBHeight() int {
	return 0
}

func (m *RequestBlock) ChainID() []byte {
	return nil
}

func (m *RequestBlock) ListHeight() int {
	return 0
}

func (m *RequestBlock) SerialHash() []byte {
	return nil
}

func (m *RequestBlock) Signature() []byte {
	return nil
}

// Validate the message, given the state.  Three possible results:
//  < 0 -- Message is invalid.  Discard
//  0   -- Cannot tell if message is Valid
//  1   -- Message is valid
func (m *RequestBlock) Validate(state interfaces.IState) int {
	return 0
}

func (m *RequestBlock) ComputeVMIndex(state interfaces.IState) {
}

func (m *RequestBlock) LeaderExecute(state interfaces.IState) {
}

func (m *RequestBlock) FollowerExecute(interfaces.IState) {
}

func (e *RequestBlock) JSONByte() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer messagesRequestBlockJSONByte.Observe(float64(time.Now().UnixNano() - callTime))
	return primitives.EncodeJSON(e)
}

func (e *RequestBlock) JSONString() (string, error) {
	callTime := time.Now().UnixNano()
	defer messagesRequestBlockJSONString.Observe(float64(time.Now().UnixNano() - callTime))
	return primitives.EncodeJSONString(e)
}

func (e *RequestBlock) JSONBuffer(b *bytes.Buffer) error {
	callTime := time.Now().UnixNano()
	defer messagesRequestBlockJSONBuffer.Observe(float64(time.Now().UnixNano() - callTime))
	return primitives.EncodeJSONToBuffer(e, b)
}
