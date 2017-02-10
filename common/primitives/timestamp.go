// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package primitives

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"time"

	"github.com/FactomProject/factomd/common/interfaces"
)

func GetTimeMilli() uint64 {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesGetTimeMilli.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return uint64(time.Now().UnixNano()) / 1000000 // 10^-9 >> 10^-3
}

func GetTime() uint64 {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesGetTime.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return uint64(time.Now().Unix())
}

//A structure for handling timestamps for messages
type Timestamp uint64 //in miliseconds
var _ interfaces.BinaryMarshallable = (*Timestamp)(nil)
var _ interfaces.Timestamp = (*Timestamp)(nil)

func NewTimestampNow() *Timestamp {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesNewTimestampNow.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	t := new(Timestamp)
	t.SetTimeNow()
	return t
}

func NewTimestampFromSeconds(s uint32) *Timestamp {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesNewTimestampFromSeconds.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	t := new(Timestamp)
	*t = Timestamp(int64(s) * 1000)
	return t
}

func NewTimestampFromMinutes(s uint32) *Timestamp {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesNewTimestampFromMinutes.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	t := new(Timestamp)
	*t = Timestamp(int64(s) * 60000)
	return t
}

func NewTimestampFromMilliseconds(s uint64) *Timestamp {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesNewTimestampFromMilliseconds.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	t := new(Timestamp)
	*t = Timestamp(s)
	return t
}

func (t *Timestamp) SetTimestamp(b interfaces.Timestamp) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesTimestampSetTimestamp.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if b == nil {
		t.SetTimeMilli(0)
	}
	t.SetTimeMilli(b.GetTimeMilli())
}

func (t *Timestamp) SetTimeNow() {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesTimestampSetTimeNow.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	*t = Timestamp(GetTimeMilli())
}

func (t *Timestamp) SetTimeMilli(miliseconds int64) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesTimestampSetTimeMilli.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	t.SetTime(uint64(miliseconds))
}

func (t *Timestamp) SetTime(miliseconds uint64) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesTimestampSetTime.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	*t = Timestamp(miliseconds)
}

func (t *Timestamp) SetTimeSeconds(seconds int64) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesTimestampSetTimeSeconds.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	t.SetTime(uint64(seconds * 1000))
}

func (t *Timestamp) GetTime() time.Time {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesTimestampGetTime.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return time.Unix(int64(*t/1000), int64(((*t)%1000)*1000))
}

func (t *Timestamp) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesTimestampUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if data == nil || len(data) < 6 {
		return nil, fmt.Errorf("Not enough data to unmarshal")
	}
	hd, data := binary.BigEndian.Uint32(data[:]), data[4:]
	ld, data := binary.BigEndian.Uint16(data[:]), data[2:]
	*t = Timestamp((uint64(hd) << 16) + uint64(ld))
	return data, nil
}

func (t *Timestamp) UnmarshalBinary(data []byte) error {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesTimestampUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	_, err := t.UnmarshalBinaryData(data)
	return err
}

func (t *Timestamp) GetTimeSeconds() int64 {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesTimestampGetTimeSeconds.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return int64(*t / 1000)
}

func (t *Timestamp) GetTimeMinutesUInt32() uint32 {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesTimestampGetTimeMinutesUInt32.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return uint32(*t / 60000)
}

func (t *Timestamp) GetTimeMilli() int64 {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesTimestampGetTimeMilli.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return int64(*t)
}

func (t *Timestamp) GetTimeMilliUInt64() uint64 {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesTimestampGetTimeMilliUInt64.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return uint64(*t)
}

func (t *Timestamp) GetTimeSecondsUInt32() uint32 {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesTimestampGetTimeSecondsUInt32.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return uint32(*t / 1000)
}

func (t *Timestamp) MarshalBinary() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesTimestampMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	var out bytes.Buffer
	hd := uint32(*t >> 16)
	ld := uint16(*t & 0xFFFF)
	binary.Write(&out, binary.BigEndian, uint32(hd))
	binary.Write(&out, binary.BigEndian, uint16(ld))
	return out.Bytes(), nil
}

func (t *Timestamp) String() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesTimestampString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return t.GetTime().Format("2006-01-02 15:04:05")
}
