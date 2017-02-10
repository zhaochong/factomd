// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package entryCreditBlock

import (
	"fmt"

	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
	"time"
)

const (
	MinuteNumberSize = 1
)

type MinuteNumber struct {
	Number uint8
}

var _ interfaces.Printable = (*MinuteNumber)(nil)
var _ interfaces.BinaryMarshallable = (*MinuteNumber)(nil)
var _ interfaces.ShortInterpretable = (*MinuteNumber)(nil)
var _ interfaces.IECBlockEntry = (*MinuteNumber)(nil)

func (e *MinuteNumber) Hash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockMinuteNumberHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	bin, err := e.MarshalBinary()
	if err != nil {
		panic(err)
	}
	return primitives.Sha(bin)
}

func (e *MinuteNumber) GetHash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockMinuteNumberGetHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return e.Hash()
}

func (e *MinuteNumber) GetSigHash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockMinuteNumberGetSigHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return nil
}

func (a *MinuteNumber) GetEntryHash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockMinuteNumberGetEntryHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return nil
}

func (b *MinuteNumber) IsInterpretable() bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockMinuteNumberIsInterpretable.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return true
}

func (b *MinuteNumber) Interpret() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockMinuteNumberInterpret.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return fmt.Sprintf("MinuteNumber %v", b.Number)
}

func NewMinuteNumber(number uint8) *MinuteNumber {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockNewMinuteNumber.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	mn := new(MinuteNumber)
	mn.Number = number
	return mn
}

func (m *MinuteNumber) ECID() byte {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockMinuteNumberECID.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return ECIDMinuteNumber
}

func (m *MinuteNumber) MarshalBinary() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockMinuteNumberMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	buf := new(primitives.Buffer)
	buf.WriteByte(m.Number)
	return buf.DeepCopyBytes(), nil
}

func (m *MinuteNumber) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockMinuteNumberUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error unmarshalling MinuteNumber: %v", r)
		}
	}()

	buf := primitives.NewBuffer(data)
	var c byte
	if c, err = buf.ReadByte(); err != nil {
		return
	} else {
		m.Number = c
	}
	newData = buf.DeepCopyBytes()
	return
}

func (m *MinuteNumber) UnmarshalBinary(data []byte) (err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockMinuteNumberUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	_, err = m.UnmarshalBinaryData(data)
	return
}

func (e *MinuteNumber) JSONByte() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockMinuteNumberJSONByte.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSON(e)
}

func (e *MinuteNumber) JSONString() (string, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockMinuteNumberJSONString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSONString(e)
}

func (e *MinuteNumber) String() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockMinuteNumberString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	var out primitives.Buffer
	out.WriteString(fmt.Sprintf(" %-20s\n", "MinuteNumber"))
	out.WriteString(fmt.Sprintf("   %-20s %d\n", "Number", e.Number))
	return (string)(out.DeepCopyBytes())
}

func (e *MinuteNumber) GetTimestamp() interfaces.Timestamp {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockMinuteNumberGetTimestamp.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return nil
}
