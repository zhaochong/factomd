// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package messages

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"time"

	"github.com/FactomProject/factomd/common/constants"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
)

//Structure to request missing messages in a node's process list
type MissingMsgResponse struct {
	MessageBase

	Timestamp   interfaces.Timestamp
	AckResponse interfaces.IMsg
	MsgResponse interfaces.IMsg

	//No signature!

	//Not marshalled
	hash interfaces.IHash
}

var _ interfaces.IMsg = (*MissingMsgResponse)(nil)

func (a *MissingMsgResponse) IsSameAs(b *MissingMsgResponse) bool {
	callTime := time.Now().UnixNano()
	defer messagesMissingMsgNewMissingMsgIsSameAs.Observe(float64(time.Now().UnixNano() - callTime))
	if b == nil {
		return false
	}
	if a.Timestamp.GetTimeMilli() != b.Timestamp.GetTimeMilli() {
		return false
	}

	if a.MsgResponse.GetHash() != b.MsgResponse.GetHash() {
		fmt.Println("MissingMsgResponse IsNotSameAs because GetHash mismatch")
		return false
	}

	if a.AckResponse.GetHash() != b.AckResponse.GetHash() {
		fmt.Println("MissingMsgResponse IsNotSameAs because Ack GetHash mismatch")
		return false
	}

	return true
}

func (m *MissingMsgResponse) Process(uint32, interfaces.IState) bool {
	return true
}

func (m *MissingMsgResponse) GetRepeatHash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer messagesMissingMsgNewMissingMsgGetRepeatHash.Observe(float64(time.Now().UnixNano() - callTime))
	return m.GetMsgHash()
}

func (m *MissingMsgResponse) GetHash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer messagesMissingMsgNewMissingMsgGetHash.Observe(float64(time.Now().UnixNano() - callTime))
	if m.hash == nil {
		data, err := m.MarshalBinary()
		if err != nil {
			panic(fmt.Sprintf("Error in MissingMsg.GetHash(): %s", err.Error()))
		}
		m.hash = primitives.Sha(data)
	}
	return m.hash
}

func (m *MissingMsgResponse) GetMsgHash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer messagesMissingMsgNewMissingMsgGetMsgHash.Observe(float64(time.Now().UnixNano() - callTime))
	if m.MsgHash == nil {
		data, err := m.MarshalBinary()
		if err != nil {
			return nil
		}
		m.MsgHash = primitives.Sha(data)
	}
	return m.MsgHash
}

func (m *MissingMsgResponse) GetTimestamp() interfaces.Timestamp {
	callTime := time.Now().UnixNano()
	defer messagesMissingMsgNewMissingMsgGetTimestamp.Observe(float64(time.Now().UnixNano() - callTime))
	return m.Timestamp
}

func (m *MissingMsgResponse) Type() byte {
	callTime := time.Now().UnixNano()
	defer messagesMissingMsgNewMissingMsgType.Observe(float64(time.Now().UnixNano() - callTime))
	return constants.MISSING_MSG_RESPONSE
}

func (m *MissingMsgResponse) Int() int {
	callTime := time.Now().UnixNano()
	defer messagesMissingMsgNewMissingMsgInt.Observe(float64(time.Now().UnixNano() - callTime))
	return -1
}

func (m *MissingMsgResponse) Bytes() []byte {
	callTime := time.Now().UnixNano()
	defer messagesMissingMsgNewMissingMsgBytes.Observe(float64(time.Now().UnixNano() - callTime))
	return nil
}

func (m *MissingMsgResponse) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	callTime := time.Now().UnixNano()
	defer messagesMissingMsgNewMissingMsgUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error unmarshalling: %v", r)
		}
	}()
	newData = data
	if newData[0] != m.Type() {
		return nil, fmt.Errorf("%s", "Invalid Message type")
	}
	newData = newData[1:]

	m.Timestamp = new(primitives.Timestamp)
	newData, err = m.Timestamp.UnmarshalBinaryData(newData)
	if err != nil {
		return nil, err
	}

	b, newData := newData[0], newData[1:]

	if b == 1 {
		m.AckResponse = new(Ack)
		newData, err = m.AckResponse.UnmarshalBinaryData(newData)

		if err != nil {
			return nil, err
		}
	}

	mr, err := UnmarshalMessage(newData)

	if err != nil {
		return nil, err
	}
	m.MsgResponse = mr

	m.Peer2Peer = true // Always a peer2peer request.

	return data, nil
}

func (m *MissingMsgResponse) UnmarshalBinary(data []byte) error {
	callTime := time.Now().UnixNano()
	defer messagesMissingMsgNewMissingMsgUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	_, err := m.UnmarshalBinaryData(data)
	return err
}

func (m *MissingMsgResponse) MarshalBinary() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer messagesMissingMsgNewMissingMsgMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	var buf primitives.Buffer

	binary.Write(&buf, binary.BigEndian, m.Type())

	t := m.GetTimestamp()
	data, err := t.MarshalBinary()
	if err != nil {
		return nil, err
	}
	buf.Write(data)

	if m.AckResponse == nil {
		buf.WriteByte(0)
	} else {
		buf.WriteByte(1)

		ackData, err := m.AckResponse.MarshalBinary()
		if err != nil {
			return nil, err
		}
		buf.Write(ackData)
	}

	msgData, err := m.MsgResponse.MarshalBinary()
	if err != nil {
		return nil, err
	}
	buf.Write(msgData)

	var mmm MissingMsgResponse

	bb := buf.DeepCopyBytes()

	//TODO: delete this once we have unit tests
	if unmarshalErr := mmm.UnmarshalBinary(bb); unmarshalErr != nil {
		fmt.Println("MissingMsgResponse failed to marshal/unmarshal: ", unmarshalErr)
		return nil, unmarshalErr
	}

	return bb, nil
}

func (m *MissingMsgResponse) String() string {
	callTime := time.Now().UnixNano()
	defer messagesMissingMsgNewMissingMsgString.Observe(float64(time.Now().UnixNano() - callTime))
	ack, ok := m.AckResponse.(*Ack)
	if !ok {
		return fmt.Sprint("MissingMsgResponse (no Ack) <-- ", m.MsgResponse.String())
	}
	return fmt.Sprintf("MissingMsgResponse <-- DBHeight:%3d vm=%3d PL Height:%3d msgHash[%x]", ack.DBHeight, ack.VMIndex, ack.Height, m.GetMsgHash().Bytes()[:3])
}

func (m *MissingMsgResponse) ChainID() []byte {
	return nil
}

func (m *MissingMsgResponse) ListHeight() int {
	return 0
}

// Validate the message, given the state.  Three possible results:
//  < 0 -- Message is invalid.  Discard
//  0   -- Cannot tell if message is Valid
//  1   -- Message is valid
func (m *MissingMsgResponse) Validate(state interfaces.IState) int {
	callTime := time.Now().UnixNano()
	defer messagesMissingMsgNewMissingMsgValidate.Observe(float64(time.Now().UnixNano() - callTime))
	if m.AckResponse == nil {
		return -1
	}
	if m.MsgResponse == nil {
		return -1
	}
	return 1
}

func (m *MissingMsgResponse) ComputeVMIndex(state interfaces.IState) {

}

func (m *MissingMsgResponse) LeaderExecute(state interfaces.IState) {
	callTime := time.Now().UnixNano()
	defer messagesMissingMsgNewMissingMsgLeaderExecute.Observe(float64(time.Now().UnixNano() - callTime))
	m.FollowerExecute(state)
}

func (m *MissingMsgResponse) FollowerExecute(state interfaces.IState) {
	callTime := time.Now().UnixNano()
	defer messagesMissingMsgNewMissingMsgFollowerExecute.Observe(float64(time.Now().UnixNano() - callTime))

	state.FollowerExecuteMMR(m)

	return
}

func (e *MissingMsgResponse) JSONByte() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer messagesMissingMsgNewMissingMsgJSONByte.Observe(float64(time.Now().UnixNano() - callTime))
	return primitives.EncodeJSON(e)
}

func (e *MissingMsgResponse) JSONString() (string, error) {
	callTime := time.Now().UnixNano()
	defer messagesMissingMsgNewMissingMsgJSONString.Observe(float64(time.Now().UnixNano() - callTime))
	return primitives.EncodeJSONString(e)
}

func (e *MissingMsgResponse) JSONBuffer(b *bytes.Buffer) error {
	callTime := time.Now().UnixNano()
	defer messagesMissingMsgNewMissingMsgJSONBuffer.Observe(float64(time.Now().UnixNano() - callTime))
	return primitives.EncodeJSONToBuffer(e, b)
}

func NewMissingMsgResponse(state interfaces.IState, msgResponse interfaces.IMsg, ackResponse interfaces.IMsg) interfaces.IMsg {
	callTime := time.Now().UnixNano()
	defer messagesMissingMsgNewMissingMsg.Observe(float64(time.Now().UnixNano() - callTime))

	msg := new(MissingMsgResponse)

	msg.Peer2Peer = true // Always a peer2peer request.
	msg.Timestamp = state.GetTimestamp()
	msg.MsgResponse = msgResponse
	msg.AckResponse = ackResponse

	return msg
}
