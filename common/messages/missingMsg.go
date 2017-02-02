// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package messages

import (
	"encoding/binary"
	"fmt"
	"time"

	"github.com/FactomProject/factomd/common/constants"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
)

//Structure to request missing messages in a node's process list
type MissingMsg struct {
	MessageBase

	Timestamp         interfaces.Timestamp
	Asking            interfaces.IHash
	DBHeight          uint32
	SystemHeight      uint32 // Might as well check for a missing Server Fault
	ProcessListHeight []uint32

	//No signature!

	//Not marshalled
	hash interfaces.IHash
}

var _ interfaces.IMsg = (*MissingMsg)(nil)

func (a *MissingMsg) IsSameAs(b *MissingMsg) bool {
	callTime := time.Now().UnixNano()
	defer messagesMissingMsgIsSameAs.Observe(float64(time.Now().UnixNano() - callTime))
	if b == nil {
		return false
	}
	if a.Timestamp.GetTimeMilli() != b.Timestamp.GetTimeMilli() {
		return false
	}

	if a.DBHeight != b.DBHeight {
		return false
	}

	if a.VMIndex != b.VMIndex {
		return false
	}

	if len(a.ProcessListHeight) != len(b.ProcessListHeight) {
		return false
	}
	for i, v := range a.ProcessListHeight {
		if v != b.ProcessListHeight[i] {
			return false
		}
	}

	return true
}

func (m *MissingMsg) Process(uint32, interfaces.IState) bool {
	panic("MissingMsg should not have its Process() method called")
}

func (m *MissingMsg) GetRepeatHash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer messagesMissingMsgGetRepeatHash.Observe(float64(time.Now().UnixNano() - callTime))
	return m.GetMsgHash()
}

func (m *MissingMsg) GetHash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer messagesMissingMsgGetHash.Observe(float64(time.Now().UnixNano() - callTime))
	if m.hash == nil {
		data, err := m.MarshalBinary()
		if err != nil {
			panic(fmt.Sprintf("Error in MissingMsg.GetHash(): %s", err.Error()))
		}
		m.hash = primitives.Sha(data)
	}
	return m.hash
}

func (m *MissingMsg) GetMsgHash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer messagesMissingMsgGetMsgHash.Observe(float64(time.Now().UnixNano() - callTime))
	if m.MsgHash == nil {
		data, err := m.MarshalBinary()
		if err != nil {
			return nil
		}
		m.MsgHash = primitives.Sha(data)
	}
	return m.MsgHash
}

func (m *MissingMsg) GetTimestamp() interfaces.Timestamp {
	callTime := time.Now().UnixNano()
	defer messagesMissingMsgGetTimestamp.Observe(float64(time.Now().UnixNano() - callTime))
	return m.Timestamp
}

func (m *MissingMsg) Type() byte {
	callTime := time.Now().UnixNano()
	defer messagesMissingMsgType.Observe(float64(time.Now().UnixNano() - callTime))
	return constants.MISSING_MSG
}

func (m *MissingMsg) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	callTime := time.Now().UnixNano()
	defer messagesMissingMsg.Observe(float64(time.Now().UnixNano() - callTime))
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

	m.Asking = new(primitives.Hash)
	newData, err = m.Asking.UnmarshalBinaryData(newData)
	if err != nil {
		return nil, err
	}

	m.VMIndex, newData = int(newData[0]), newData[1:]
	m.DBHeight, newData = binary.BigEndian.Uint32(newData[0:4]), newData[4:]
	m.SystemHeight, newData = binary.BigEndian.Uint32(newData[0:4]), newData[4:]

	// Get all the missing messages...
	lenl, newData := binary.BigEndian.Uint32(newData[0:4]), newData[4:]
	for i := 0; i < int(lenl); i++ {
		var height uint32
		height, newData = binary.BigEndian.Uint32(newData[0:4]), newData[4:]
		m.ProcessListHeight = append(m.ProcessListHeight, height)
	}

	m.Peer2Peer = true // Always a peer2peer request.

	return data, nil
}

func (m *MissingMsg) UnmarshalBinary(data []byte) error {
	callTime := time.Now().UnixNano()
	defer messagesMissingMsgUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	_, err := m.UnmarshalBinaryData(data)
	return err
}

func (m *MissingMsg) MarshalBinary() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer messagesMissingMsgMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	var buf primitives.Buffer

	binary.Write(&buf, binary.BigEndian, m.Type())

	t := m.GetTimestamp()
	data, err := t.MarshalBinary()
	if err != nil {
		return nil, err
	}
	buf.Write(data)

	if m.Asking == nil {
		m.Asking = primitives.NewHash(constants.ZERO_HASH)
	}
	data, err = m.Asking.MarshalBinary()
	if err != nil {
		return nil, err
	}
	buf.Write(data)

	buf.WriteByte(uint8(m.VMIndex))
	binary.Write(&buf, binary.BigEndian, m.DBHeight)
	binary.Write(&buf, binary.BigEndian, m.SystemHeight)

	binary.Write(&buf, binary.BigEndian, uint32(len(m.ProcessListHeight)))
	for _, h := range m.ProcessListHeight {
		binary.Write(&buf, binary.BigEndian, h)
	}

	bb := buf.DeepCopyBytes()

	return bb, nil
}

func (m *MissingMsg) String() string {
	callTime := time.Now().UnixNano()
	defer messagesMissingMsgString.Observe(float64(time.Now().UnixNano() - callTime))
	str := ""
	for _, n := range m.ProcessListHeight {
		str = fmt.Sprintf("%s%d,", str, n)
	}
	return fmt.Sprintf("MissingMsg --> Asking %x DBHeight:%3d vm=%3d Hts::[%s] Sys: %d msgHash[%x]",
		m.Asking.Bytes()[:8],
		m.DBHeight,
		m.VMIndex,
		str,
		m.SystemHeight,
		m.GetMsgHash().Bytes()[:3])
}

func (m *MissingMsg) ChainID() []byte {
	return nil
}

func (m *MissingMsg) ListHeight() int {
	return 0
}

// Validate the message, given the state.  Three possible results:
//  < 0 -- Message is invalid.  Discard
//  0   -- Cannot tell if message is Valid
//  1   -- Message is valid
func (m *MissingMsg) Validate(state interfaces.IState) int {
	callTime := time.Now().UnixNano()
	defer messagesMissingMsg.Observe(float64(time.Now().UnixNano() - callTime))
	if m.Asking == nil {
		return -1
	}
	if m.Asking.IsZero() {
		return -1
	}
	return 1
}

func (m *MissingMsg) ComputeVMIndex(state interfaces.IState) {
}

func (m *MissingMsg) LeaderExecute(state interfaces.IState) {
	callTime := time.Now().UnixNano()
	defer messagesMissingMsgLeaderExecute.Observe(float64(time.Now().UnixNano() - callTime))
	m.FollowerExecute(state)
}

func (m *MissingMsg) FollowerExecute(state interfaces.IState) {
	callTime := time.Now().UnixNano()
	defer messagesMissingMsgFollowerExecute.Observe(float64(time.Now().UnixNano() - callTime))
	state.FollowerExecuteMissingMsg(m)
}

func (e *MissingMsg) JSONByte() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer messagesMissingMsgJSONByte.Observe(float64(time.Now().UnixNano() - callTime))
	return primitives.EncodeJSON(e)
}

func (e *MissingMsg) JSONString() (string, error) {
	callTime := time.Now().UnixNano()
	defer messagesMissingMsgJSONString.Observe(float64(time.Now().UnixNano() - callTime))
	return primitives.EncodeJSONString(e)
}

// AddHeight: Add a Missing Message Height to the request
func (e *MissingMsg) AddHeight(h uint32) {
	callTime := time.Now().UnixNano()
	defer messagesMissingMsgAddHeight.Observe(float64(time.Now().UnixNano() - callTime))
	e.ProcessListHeight = append(e.ProcessListHeight, h)
}

// NewMissingMsg: Build a missing Message request, and add the first Height
func NewMissingMsg(state interfaces.IState, vm int, dbHeight uint32, processlistHeight uint32) *MissingMsg {
	callTime := time.Now().UnixNano()
	defer messagesMissingMsgNewMissingMsg.Observe(float64(time.Now().UnixNano() - callTime))
	msg := new(MissingMsg)

	msg.Asking = state.GetIdentityChainID()
	msg.Peer2Peer = true // Always a peer2peer request // .
	msg.VMIndex = vm
	msg.Timestamp = state.GetTimestamp()
	msg.DBHeight = dbHeight
	msg.ProcessListHeight = append(msg.ProcessListHeight, processlistHeight)
	msg.SystemHeight = uint32(state.GetSystemHeight(dbHeight))
	return msg
}
