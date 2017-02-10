// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package messages

import (
	"encoding/binary"
	"fmt"

	"github.com/FactomProject/factomd/common/constants"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
	"time"
)

// Communicate a Directory Block State

type DBStateMissing struct {
	MessageBase
	Timestamp interfaces.Timestamp

	DBHeightStart uint32 // First block missing
	DBHeightEnd   uint32 // Last block missing.

	//Not signed!
}

var _ interfaces.IMsg = (*DBStateMissing)(nil)

func (a *DBStateMissing) IsSameAs(b *DBStateMissing) bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesDBStateMissingIsSameAs.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if b == nil {
		return false
	}
	if a.Timestamp.GetTimeMilli() != b.Timestamp.GetTimeMilli() {
		return false
	}
	if a.DBHeightStart != b.DBHeightStart {
		return false
	}
	if a.DBHeightEnd != b.DBHeightEnd {
		return false
	}

	return true
}

func (m *DBStateMissing) GetRepeatHash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesDBStateMissingGetRepeatHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return m.GetMsgHash()
}

func (m *DBStateMissing) GetHash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesDBStateMissingGetHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return m.GetMsgHash()
}

func (m *DBStateMissing) GetMsgHash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesDBStateMissingGetMsgHash.Observe(float64(time.Now().UnixNano() - callTime))
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

func (m *DBStateMissing) Type() byte {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesDBStateMissingType.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return constants.DBSTATE_MISSING_MSG
}

func (m *DBStateMissing) GetTimestamp() interfaces.Timestamp {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesDBStateMissingGetTimestamp.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return m.Timestamp
}

// Validate the message, given the state.  Three possible results:
//  < 0 -- Message is invalid.  Discard
//  0   -- Cannot tell if message is Valid
//  1   -- Message is valid
func (m *DBStateMissing) Validate(state interfaces.IState) int {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesDBStateMissingValidate.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if m.DBHeightStart > m.DBHeightEnd {
		return -1
	}
	return 1
}

func (m *DBStateMissing) ComputeVMIndex(state interfaces.IState) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesDBStateMissingComputeVMIndex.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

}

// Execute the leader functions of the given message
func (m *DBStateMissing) LeaderExecute(state interfaces.IState) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesDBStateMissingLeaderExecute.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	m.FollowerExecute(state)
}

// Only send the same block again after 15 seconds.
func (m *DBStateMissing) send(dbheight uint32, state interfaces.IState) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesDBStateMissingsend.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	send := true

	now := state.GetTimestamp()
	sents := state.GetDBStatesSent()
	var keeps []*interfaces.DBStateSent

	for _, v := range sents {
		if now.GetTimeSeconds()-v.Sent.GetTimeSeconds() < 1 {
			if v.DBHeight == dbheight {
				send = false
			}
			keeps = append(keeps, v)
		}
	}
	if send {
		msg, err := state.LoadDBState(dbheight)
		if msg != nil && err == nil {
			msg.SetOrigin(m.GetOrigin())
			msg.SetNetworkOrigin(m.GetNetworkOrigin())
			msg.SetNoResend(false)
			msg.SendOut(state, msg)
			state.IncDBStateAnswerCnt()
			v := new(interfaces.DBStateSent)
			v.DBHeight = dbheight
			v.Sent = now
			keeps = append(keeps, v)
		}
		state.SetDBStatesSent(keeps)
	}
}

func (m *DBStateMissing) FollowerExecute(state interfaces.IState) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesDBStateMissingFollowerExecute.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if len(state.NetworkOutMsgQueue()) > 100 {
		return
	}

	// TODO: Likely need to consider a limit on how many blocks we reply with.  For now,
	// just give them what they ask for.
	start := m.DBHeightStart
	end := m.DBHeightEnd
	if end-start > 200 {
		end = start + 200
	}
	for dbs := start; dbs <= end; dbs++ {
		m.send(dbs, state)
	}

	return
}

// Acknowledgements do not go into the process list.
func (e *DBStateMissing) Process(dbheight uint32, state interfaces.IState) bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesDBStateMissingProcess.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	panic("Ack object should never have its Process() method called")
}

func (e *DBStateMissing) JSONByte() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesDBStateMissingJSONByte.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSON(e)
}

func (e *DBStateMissing) JSONString() (string, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesDBStateMissingJSONString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSONString(e)
}

func (m *DBStateMissing) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesDBStateMissingUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error unmarshalling Directory Block State Missing Message: %v", r)
		}
	}()
	newData = data
	if newData[0] != m.Type() {
		return nil, fmt.Errorf("Invalid Message type")
	}
	newData = newData[1:]

	m.Peer2Peer = true // This is always a Peer2peer message

	m.Timestamp = new(primitives.Timestamp)
	newData, err = m.Timestamp.UnmarshalBinaryData(newData)
	if err != nil {
		return nil, err
	}

	m.DBHeightStart, newData = binary.BigEndian.Uint32(newData[0:4]), newData[4:]
	m.DBHeightEnd, newData = binary.BigEndian.Uint32(newData[0:4]), newData[4:]

	return
}

func (m *DBStateMissing) UnmarshalBinary(data []byte) error {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesDBStateMissingUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	_, err := m.UnmarshalBinaryData(data)
	return err
}

func (m *DBStateMissing) MarshalForSignature() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesDBStateMissingMarshalForSignature.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	var buf primitives.Buffer

	binary.Write(&buf, binary.BigEndian, m.Type())

	t := m.GetTimestamp()
	data, err := t.MarshalBinary()
	if err != nil {
		return nil, err
	}
	buf.Write(data)

	binary.Write(&buf, binary.BigEndian, m.DBHeightStart)
	binary.Write(&buf, binary.BigEndian, m.DBHeightEnd)

	return buf.DeepCopyBytes(), nil
}

func (m *DBStateMissing) MarshalBinary() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesDBStateMissingMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return m.MarshalForSignature()
}

func (m *DBStateMissing) String() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesDBStateMissingString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return fmt.Sprintf("DBStateMissing: %d-%d", m.DBHeightStart, m.DBHeightEnd)
}

func NewDBStateMissing(state interfaces.IState, dbheightStart uint32, dbheightEnd uint32) interfaces.IMsg {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesNewDBStateMissing.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	msg := new(DBStateMissing)

	msg.Peer2Peer = true // Always a peer2peer request.
	msg.Timestamp = state.GetTimestamp()
	msg.DBHeightStart = dbheightStart
	msg.DBHeightEnd = dbheightEnd

	return msg
}
