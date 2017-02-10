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

//Requests entry blocks from a range of DBlocks

type MissingEntryBlocks struct {
	MessageBase
	Timestamp interfaces.Timestamp

	DBHeightStart uint32 // First block missing
	DBHeightEnd   uint32 // Last block missing.

	//Not signed!
}

var _ interfaces.IMsg = (*MissingEntryBlocks)(nil)

func (a *MissingEntryBlocks) IsSameAs(b *MissingEntryBlocks) bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMissingEntryBlocksIsSameAs.Observe(float64(time.Now().UnixNano() - callTime))
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

func (m *MissingEntryBlocks) GetRepeatHash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMissingEntryBlocksGetRepeatHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return m.GetMsgHash()
}

func (m *MissingEntryBlocks) GetHash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMissingEntryBlocksGetHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return m.GetMsgHash()
}

func (m *MissingEntryBlocks) GetMsgHash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMissingEntryBlocksGetMsgHash.Observe(float64(time.Now().UnixNano() - callTime))
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

func (m *MissingEntryBlocks) Type() byte {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMissingEntryBlocksType.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return constants.MISSING_ENTRY_BLOCKS
}

func (m *MissingEntryBlocks) GetTimestamp() interfaces.Timestamp {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMissingEntryBlocksGetTimestamp.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return m.Timestamp
}

// Validate the message, given the state.  Three possible results:
//  < 0 -- Message is invalid.  Discard
//  0   -- Cannot tell if message is Valid
//  1   -- Message is valid
func (m *MissingEntryBlocks) Validate(state interfaces.IState) int {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMissingEntryBlocksValidate.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if m.DBHeightStart > m.DBHeightEnd {
		return -1
	}
	return 1
}

func (m *MissingEntryBlocks) ComputeVMIndex(state interfaces.IState) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMissingEntryBlocksComputeVMIndex.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

}

// Execute the leader functions of the given message
func (m *MissingEntryBlocks) LeaderExecute(state interfaces.IState) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMissingEntryBlocksLeaderExecute.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	m.FollowerExecute(state)
}

func (m *MissingEntryBlocks) FollowerExecute(state interfaces.IState) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMissingEntryBlocksFollowerExecute.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if len(state.NetworkOutMsgQueue()) > 1000 {
		return
	}
	start := m.DBHeightStart
	end := m.DBHeightEnd
	if end-start > 20 {
		end = start + 20
	}
	db := state.GetAndLockDB()
	defer state.UnlockDB()

	resp := NewEntryBlockResponse(state).(*EntryBlockResponse)

	for i := start; i <= end; i++ {
		dblk, err := db.FetchDBlockByHeight(i)
		if err != nil {
			return
		}
		if dblk == nil {
			return
		}
		for _, v := range dblk.GetDBEntries() {
			if v.GetChainID().IsMinuteMarker() == true {
				continue
			}
			eBlock, err := db.FetchEBlock(v.GetKeyMR())
			if err != nil {
				return
			}
			resp.EBlocks = append(resp.EBlocks, eBlock)
			for _, v := range eBlock.GetBody().GetEBEntries() {
				entry, err := db.FetchEntry(v)
				if err != nil {
					return
				}
				resp.Entries = append(resp.Entries, entry)
			}
		}
	}

	resp.SetOrigin(m.GetOrigin())
	resp.SetNetworkOrigin(m.GetNetworkOrigin())
	resp.SendOut(state, resp)
	state.IncDBStateAnswerCnt()

	return
}

// Acknowledgements do not go into the process list.
func (e *MissingEntryBlocks) Process(dbheight uint32, state interfaces.IState) bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMissingEntryBlocksProcess.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	panic("Ack object should never have its Process() method called")
}

func (e *MissingEntryBlocks) JSONByte() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMissingEntryBlocksJSONByte.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSON(e)
}

func (e *MissingEntryBlocks) JSONString() (string, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMissingEntryBlocksJSONString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSONString(e)
}

func (m *MissingEntryBlocks) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMissingEntryBlocksUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))
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

func (m *MissingEntryBlocks) UnmarshalBinary(data []byte) error {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMissingEntryBlocksUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	_, err := m.UnmarshalBinaryData(data)
	return err
}

func (m *MissingEntryBlocks) MarshalForSignature() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMissingEntryBlocksMarshalForSignature.Observe(float64(time.Now().UnixNano() - callTime))
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

func (m *MissingEntryBlocks) MarshalBinary() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMissingEntryBlocksMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return m.MarshalForSignature()
}

func (m *MissingEntryBlocks) String() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMissingEntryBlocksString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return fmt.Sprintf("MissingEntryBlocks: %d-%d", m.DBHeightStart, m.DBHeightEnd)
}

func NewMissingEntryBlocks(state interfaces.IState, dbheightStart uint32, dbheightEnd uint32) interfaces.IMsg {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesNewMissingEntryBlocks.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	msg := new(MissingEntryBlocks)

	msg.Peer2Peer = true // Always a peer2peer request.
	msg.Timestamp = state.GetTimestamp()
	msg.DBHeightStart = dbheightStart
	msg.DBHeightEnd = dbheightEnd

	return msg
}
