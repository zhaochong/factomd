// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package messages

import (
	"encoding/binary"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/FactomProject/factomd/common/constants"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
)

type BounceReply struct {
	MessageBase
	Name      string
	Number    int32
	Timestamp interfaces.Timestamp
	Stamps    []interfaces.Timestamp
	size      int
}

var _ interfaces.IMsg = (*BounceReply)(nil)

func (m *BounceReply) GetRepeatHash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesBounceReplyGetRepeatHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return m.GetMsgHash()
}

// We have to return the haswh of the underlying message.
func (m *BounceReply) GetHash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesBounceReplyGetHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return m.GetMsgHash()
}

func (m *BounceReply) SizeOf() int {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesBounceReplySizeOf.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	m.GetMsgHash()
	return m.size
}

func (m *BounceReply) GetMsgHash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesBounceReplyGetMsgHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	data, err := m.MarshalForSignature()

	m.size = len(data)

	if err != nil {
		return nil
	}
	m.MsgHash = primitives.Sha(data)
	return m.MsgHash
}

func (m *BounceReply) Type() byte {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesBounceReplyType.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return constants.BOUNCEREPLY_MSG
}

func (m *BounceReply) GetTimestamp() interfaces.Timestamp {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesBounceReplyGetTimestamp.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return m.Timestamp
}

func (m *BounceReply) VerifySignature() (bool, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesBounceReplyVerifySignature.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return true, nil
}

// Validate the message, given the state.  Three possible results:
//  < 0 -- Message is invalid.  Discard
//  0   -- Cannot tell if message is Valid
//  1   -- Message is valid
func (m *BounceReply) Validate(state interfaces.IState) int {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesBounceReplyValidate.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return 1
}

// Returns true if this is a message for this server to execute as
// a leader.
func (m *BounceReply) ComputeVMIndex(state interfaces.IState) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesBounceReplyComputeVMIndex.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

}

// Execute the leader functions of the given message
// Leader, follower, do the same thing.
func (m *BounceReply) LeaderExecute(state interfaces.IState) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesBounceReplyLeaderExecute.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

}

func (m *BounceReply) FollowerExecute(state interfaces.IState) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesBounceReplyFollowerExecute.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

}

// Acknowledgements do not go into the process list.
func (e *BounceReply) Process(dbheight uint32, state interfaces.IState) bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesBounceReplyProcess.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return true
}

func (e *BounceReply) JSONByte() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesBounceReplyJSONByte.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSON(e)
}

func (e *BounceReply) JSONString() (string, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesBounceReplyJSONString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSONString(e)
}

func (m *BounceReply) Sign(key interfaces.Signer) error {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesBounceReplySign.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return nil
}

func (m *BounceReply) GetSignature() interfaces.IFullSignature {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesBounceReplyGetSignature.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return nil
}

func (m *BounceReply) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesBounceReplyUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error unmarshalling: %v", r)
		}
	}()

	m.SetPeer2Peer(true)

	newData = data

	if newData[0] != m.Type() {
		return nil, errors.New("Invalid Message type")
	}
	newData = newData[1:]

	m.Name = string(newData[:32])
	newData = newData[32:]

	m.Number, newData = int32(binary.BigEndian.Uint32(newData[0:4])), newData[4:]

	m.Timestamp = new(primitives.Timestamp)
	newData, err = m.Timestamp.UnmarshalBinaryData(newData)
	if err != nil {
		return nil, err
	}

	numTS, newData := binary.BigEndian.Uint32(newData[0:4]), newData[4:]

	for i := uint32(0); i < numTS; i++ {
		ts := new(primitives.Timestamp)
		newData, err = ts.UnmarshalBinaryData(newData)
		if err != nil {
			return nil, err
		}
		m.Stamps = append(m.Stamps, ts)
	}
	return
}

func (m *BounceReply) UnmarshalBinary(data []byte) error {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesBounceReplyUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	_, err := m.UnmarshalBinaryData(data)
	return err
}

func (m *BounceReply) MarshalForSignature() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesBounceReplyMarshalForSignature.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	var buf primitives.Buffer

	binary.Write(&buf, binary.BigEndian, m.Type())

	var buff [32]byte

	copy(buff[:32], []byte(fmt.Sprintf("%32s", m.Name)))
	buf.Write(buff[:])

	binary.Write(&buf, binary.BigEndian, m.Number)

	t := m.GetTimestamp()
	data, err := t.MarshalBinary()
	if err != nil {
		return nil, err
	}
	buf.Write(data)

	binary.Write(&buf, binary.BigEndian, int32(len(m.Stamps)))

	for _, ts := range m.Stamps {
		data, err := ts.MarshalBinary()
		if err != nil {
			return nil, err
		}
		buf.Write(data)
	}

	return buf.DeepCopyBytes(), nil
}

func (m *BounceReply) MarshalBinary() (data []byte, err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesBounceReplyMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return m.MarshalForSignature()
}

func (m *BounceReply) String() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesBounceReplyString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	// bbbb Origin: 2016-09-05 12:26:20.426954586 -0500 CDT left BounceReply Start:             2016-09-05 12:26:05 Hops:     1 Size:    43 Last Hop Took 14.955 Average Hop: 14.955
	now := time.Now()
	t := fmt.Sprintf("%2d:%2d:%2d.%03d", now.Hour(), now.Minute(), now.Second(), now.Nanosecond()/1000000)
	mill := m.Timestamp.GetTimeMilli()
	mills := mill % 1000
	mill = mill / 1000
	secs := mill % 60
	mill = mill / 60
	mins := mill % 60
	mill = mill / 60
	hrs := mill % 24
	t2 := fmt.Sprintf("%2d:%2d:%2d.%03d", hrs, mins, secs, mills)
	b := m.SizeOf() % 1000
	kb := (m.SizeOf() / 1000) % 1000
	mb := (m.SizeOf() / 1000 / 1000)
	sz := fmt.Sprintf("%d,%03d", kb, b)
	if mb > 0 {
		sz = fmt.Sprintf("%d,%03d,%03d", mb, kb, b)
	}

	str := fmt.Sprintf("Origin: %12s  %10s-%03d-%03d BounceReply Start: %12s Hops: %5d Size: [%12s] ",
		t,
		strings.TrimSpace(m.Name),
		m.Number,
		len(m.Stamps),
		t2,
		len(m.Stamps),
		sz)

	elapse := primitives.NewTimestampNow().GetTimeMilli() - m.Stamps[len(m.Stamps)-1].GetTimeMilli()

	str = str + fmt.Sprintf("Last Hop Took %d.%03d", elapse/1000, elapse%1000)
	return str
}

func (a *BounceReply) IsSameAs(b *BounceReply) bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesBounceReplyIsSameAs.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return true
}
