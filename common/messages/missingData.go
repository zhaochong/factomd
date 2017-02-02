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

//Structure to request missing messages in a node's process list
type MissingData struct {
	MessageBase
	Timestamp interfaces.Timestamp

	RequestHash interfaces.IHash

	//No signature!
}

var _ interfaces.IMsg = (*MissingData)(nil)

func (a *MissingData) IsSameAs(b *MissingData) bool {
	callTime := time.Now().UnixNano()
	defer messagesMissingDataIsSameAs.Observe(float64(time.Now().UnixNano() - callTime))
	if b == nil {
		return false
	}
	if a.Timestamp.GetTimeMilli() != b.Timestamp.GetTimeMilli() {
		return false
	}

	if a.RequestHash == nil && b.RequestHash != nil {
		return false
	}
	if a.RequestHash != nil {
		if a.RequestHash.IsSameAs(b.RequestHash) == false {
			return false
		}
	}

	return true
}

func (m *MissingData) Process(uint32, interfaces.IState) bool {
	callTime := time.Now().UnixNano()
	defer messagesMissingDataProcess.Observe(float64(time.Now().UnixNano() - callTime))
	return true
}

func (m *MissingData) GetRepeatHash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer messagesMissingDataGetRepeatHash.Observe(float64(time.Now().UnixNano() - callTime))
	return m.GetMsgHash()
}

func (m *MissingData) GetHash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer messagesMissingDataGetHash.Observe(float64(time.Now().UnixNano() - callTime))
	return m.GetMsgHash()
}

func (m *MissingData) GetMsgHash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer messagesMissingDataGetMsgHash.Observe(float64(time.Now().UnixNano() - callTime))
	if m.MsgHash == nil {
		data, err := m.MarshalBinary()
		if err != nil {
			return nil
		}
		m.MsgHash = primitives.Sha(data)
	}
	return m.MsgHash
}

func (m *MissingData) GetTimestamp() interfaces.Timestamp {
	callTime := time.Now().UnixNano()
	defer messagesMissingDataGetTimestamp.Observe(float64(time.Now().UnixNano() - callTime))
	return m.Timestamp
}

func (m *MissingData) Type() byte {
	callTime := time.Now().UnixNano()
	defer messagesMissingDataType.Observe(float64(time.Now().UnixNano() - callTime))
	return constants.MISSING_DATA
}

func (m *MissingData) Int() int {
	callTime := time.Now().UnixNano()
	defer messagesMissingDataInt.Observe(float64(time.Now().UnixNano() - callTime))
	return -1
}

func (m *MissingData) Bytes() []byte {
	callTime := time.Now().UnixNano()
	defer messagesMissingDataBytes.Observe(float64(time.Now().UnixNano() - callTime))
	return nil
}

func (m *MissingData) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	callTime := time.Now().UnixNano()
	defer messagesMissingDataUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))
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

	m.RequestHash = primitives.NewHash(constants.ZERO_HASH)
	newData, err = m.RequestHash.UnmarshalBinaryData(newData)
	if err != nil {
		return nil, err
	}

	m.Peer2Peer = true // Always a peer2peer request.

	return data, nil
}

func (m *MissingData) UnmarshalBinary(data []byte) error {
	callTime := time.Now().UnixNano()
	defer messagesMissingData.Observe(float64(time.Now().UnixNano() - callTime))
	_, err := m.UnmarshalBinaryData(data)
	return err
}

func (m *MissingData) MarshalBinary() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer messagesMissingDataMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	var buf primitives.Buffer
	buf.Write([]byte{m.Type()})
	if d, err := m.Timestamp.MarshalBinary(); err != nil {
		return nil, err
	} else {
		buf.Write(d)
	}

	if d, err := m.RequestHash.MarshalBinary(); err != nil {
		return nil, err
	} else {
		buf.Write(d)
	}

	return buf.DeepCopyBytes(), nil
}

func (m *MissingData) String() string {
	callTime := time.Now().UnixNano()
	defer messagesMissingDataString.Observe(float64(time.Now().UnixNano() - callTime))
	return fmt.Sprintf("MissingData: [%x]", m.RequestHash.Bytes()[:5])
}

// Validate the message, given the state.  Three possible results:
//  < 0 -- Message is invalid.  Discard
//  0   -- Cannot tell if message is Valid
//  1   -- Message is valid
func (m *MissingData) Validate(state interfaces.IState) int {
	return 1
}

func (m *MissingData) ComputeVMIndex(state interfaces.IState) {
}

func (m *MissingData) LeaderExecute(state interfaces.IState) {
	callTime := time.Now().UnixNano()
	defer messagesMissingDataLeaderExecute.Observe(float64(time.Now().UnixNano() - callTime))
	m.FollowerExecute(state)
}

func (m *MissingData) FollowerExecute(state interfaces.IState) {
	callTime := time.Now().UnixNano()
	defer messagesMissingDataFollowerExecute.Observe(float64(time.Now().UnixNano() - callTime))
	var dataObject interfaces.BinaryMarshallable
	//var dataHash interfaces.IHash
	rawObject, dataType, err := state.LoadDataByHash(m.RequestHash)

	if rawObject != nil && err == nil { // If I don't have this message, ignore.
		switch dataType {
		case 0: // DataType = entry
			dataObject = rawObject.(interfaces.IEBEntry)
			//dataHash = dataObject.(interfaces.IEBEntry).GetHash()
		case 1: // DataType = eblock
			dataObject = rawObject.(interfaces.IEntryBlock)
			//dataHash, _ = dataObject.(interfaces.IEntryBlock).Hash()
		default:
			return
		}

		msg := NewDataResponse(state, dataObject, dataType, m.RequestHash)

		msg.SetOrigin(m.GetOrigin())
		msg.SetNetworkOrigin(m.GetNetworkOrigin())
		msg.SendOut(state, msg)
	}
	return
}

func (e *MissingData) JSONByte() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer messagesMissingDataJSONByte.Observe(float64(time.Now().UnixNano() - callTime))
	return primitives.EncodeJSON(e)
}

func (e *MissingData) JSONString() (string, error) {
	callTime := time.Now().UnixNano()
	defer messagesMissingDataJSONString.Observe(float64(time.Now().UnixNano() - callTime))
	return primitives.EncodeJSONString(e)
}

func (e *MissingData) JSONBuffer(b *bytes.Buffer) error {
	callTime := time.Now().UnixNano()
	defer messagesMissingDataJSONBuffer.Observe(float64(time.Now().UnixNano() - callTime))
	return primitives.EncodeJSONToBuffer(e, b)
}

func NewMissingData(state interfaces.IState, requestHash interfaces.IHash) interfaces.IMsg {
	callTime := time.Now().UnixNano()
	defer messagesMissingData.Observe(float64(time.Now().UnixNano() - callTime))

	msg := new(MissingData)

	msg.Peer2Peer = true // Always a peer2peer request.
	msg.Timestamp = state.GetTimestamp()
	msg.RequestHash = requestHash

	return msg
}
