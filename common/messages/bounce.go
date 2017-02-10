// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package messages

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/FactomProject/factomd/common/constants"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
)

type Bounce struct {
	MessageBase
	Name      string
	Number    int32
	Timestamp interfaces.Timestamp
	Stamps    []interfaces.Timestamp
	Data      []byte
	size      int
}

var _ interfaces.IMsg = (*Bounce)(nil)

func (m *Bounce) AddData(dataSize int) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesBounceAddData.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	m.Data = make([]byte, dataSize)
	for i, _ := range m.Data {
		m.Data[i] = byte(rand.Int())
	}
}

func (m *Bounce) GetRepeatHash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesBounceGetRepeatHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return m.GetMsgHash()
}

// We have to return the haswh of the underlying message.
func (m *Bounce) GetHash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesBounceGetHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return m.GetMsgHash()
}

func (m *Bounce) SizeOf() int {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesBounceSizeOf.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	m.GetMsgHash()
	return m.size
}

func (m *Bounce) GetMsgHash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesBounceGetMsgHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	data, err := m.MarshalForSignature()

	m.size = len(data)

	if err != nil {
		return nil
	}
	m.MsgHash = primitives.Sha(data)
	return m.MsgHash
}

func (m *Bounce) Type() byte {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesBounceType.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return constants.BOUNCE_MSG
}

func (m *Bounce) GetTimestamp() interfaces.Timestamp {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesBounceGetTimestamp.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return m.Timestamp
}

func (m *Bounce) VerifySignature() (bool, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesBounceVerifySignature.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return true, nil
}

// Validate the message, given the state.  Three possible results:
//  < 0 -- Message is invalid.  Discard
//  0   -- Cannot tell if message is Valid
//  1   -- Message is valid
func (m *Bounce) Validate(state interfaces.IState) int {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesBounceValidate.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return 1
}

// Returns true if this is a message for this server to execute as
// a leader.
func (m *Bounce) ComputeVMIndex(state interfaces.IState) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesBounceComputeVMIndex.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

}

// Execute the leader functions of the given message
// Leader, follower, do the same thing.
func (m *Bounce) LeaderExecute(state interfaces.IState) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesBounceLeaderExecute.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

}

func (m *Bounce) FollowerExecute(state interfaces.IState) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesBounceFollowerExecute.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

}

// Acknowledgements do not go into the process list.
func (e *Bounce) Process(dbheight uint32, state interfaces.IState) bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesBounceProcess.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return true
}

func (e *Bounce) JSONByte() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesBounceJSONByte.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSON(e)
}

func (e *Bounce) JSONString() (string, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesBounceJSONString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSONString(e)
}

func (m *Bounce) Sign(key interfaces.Signer) error {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesBounceSign.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return nil
}

func (m *Bounce) GetSignature() interfaces.IFullSignature {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesBounceGetSignature.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return nil
}

func (m *Bounce) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesBounceUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error unmarshalling: %v", r)
		}
	}()
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

	lenData, newData := binary.BigEndian.Uint32(newData[0:4]), newData[4:]

	m.Data = make([]byte, lenData)
	copy(m.Data, newData)
	newData = newData[lenData:]

	return
}

func (m *Bounce) UnmarshalBinary(data []byte) error {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesBounceUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	_, err := m.UnmarshalBinaryData(data)
	return err
}

func (m *Bounce) MarshalForSignature() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesBounceMarshalForSignature.Observe(float64(time.Now().UnixNano() - callTime))
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

	binary.Write(&buf, binary.BigEndian, int32(len(m.Data)))
	buf.Write(m.Data)

	return buf.DeepCopyBytes(), nil
}

func (m *Bounce) MarshalBinary() (data []byte, err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesBounceMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return m.MarshalForSignature()
}

func (m *Bounce) String() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesBounceString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	// bbbb Origin: 2016-09-05 12:26:20.426954586 -0500 CDT left Bounce Start:             2016-09-05 12:26:05 Hops:     1 Size:    43 Last Hop Took 14.955 Average Hop: 14.955
	now := time.Now()
	t := fmt.Sprintf("%2d:%02d:%02d.%03d", now.Hour(), now.Minute(), now.Second(), now.Nanosecond()/1000000)
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

	str := fmt.Sprintf("Origin: %12s %30s-%04d Bounce Start: %12s Hops: %5d [Size: %12s] ",
		t,
		strings.TrimSpace(m.Name),
		m.Number,
		t2,
		len(m.Stamps),
		sz)
	var sum int64
	for i := 0; i < len(m.Stamps)-1; i++ {
		sum += m.Stamps[i+1].GetTimeMilli() - m.Stamps[i].GetTimeMilli()
	}
	elapse := primitives.NewTimestampNow().GetTimeMilli() - m.Stamps[len(m.Stamps)-1].GetTimeMilli()
	sum += elapse
	sign := " "
	if sum < 0 {
		sign = "-"
		sum = sum * -1
	}
	avg := sum / (int64(len(m.Stamps)))
	str = str + fmt.Sprintf("Last Hop Took %3d.%03d Average Hop: %s%3d.%03d Hash: %x", elapse/1000, elapse%1000, sign, avg/1000, avg%1000, m.GetHash().Bytes()[:4])
	return str
}

func (a *Bounce) IsSameAs(b *Bounce) bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesBounceIsSameAs.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return true
}
