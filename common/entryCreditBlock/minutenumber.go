// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package entryCreditBlock

import (
	"fmt"

	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
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
	callTime := time.Now().UnixNano()
	defer entryCreditBlockminuteNumberHash.Observe(float64(time.Now().UnixNano() - callTime))	
	bin, err := e.MarshalBinary()
	if err != nil {
		panic(err)
	}
	return primitives.Sha(bin)
}

func (e *MinuteNumber) GetHash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer entryCreditBlockminuteNumberGetHash.Observe(float64(time.Now().UnixNano() - callTime))	
	return e.Hash()
}

func (e *MinuteNumber) GetSigHash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer entryCreditBlockminuteNumberGetSigHash.Observe(float64(time.Now().UnixNano() - callTime))	
	return nil
}

func (a *MinuteNumber) GetEntryHash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer entryCreditBlockminuteNumberGetEntryHash.Observe(float64(time.Now().UnixNano() - callTime))	
	return nil
}

func (b *MinuteNumber) IsInterpretable() bool {
	callTime := time.Now().UnixNano()
	defer entryCreditBlockminuteNumberIsInterpretable.Observe(float64(time.Now().UnixNano() - callTime))	
	return true
}

func (b *MinuteNumber) Interpret() string {
	callTime := time.Now().UnixNano()
	defer entryCreditBlockminuteNumberInterpret.Observe(float64(time.Now().UnixNano() - callTime))	
	return fmt.Sprintf("MinuteNumber %v", b.Number)
}

func NewMinuteNumber(number uint8) *MinuteNumber {
	callTime := time.Now().UnixNano()
	defer entryCreditBlockminuteNumberNewMinuteNumber.Observe(float64(time.Now().UnixNano() - callTime))	
	mn := new(MinuteNumber)
	mn.Number = number
	return mn
}

func (m *MinuteNumber) ECID() byte {
	callTime := time.Now().UnixNano()
	defer entryCreditBlockminuteNumberECID.Observe(float64(time.Now().UnixNano() - callTime))	
	return ECIDMinuteNumber
}

func (m *MinuteNumber) MarshalBinary() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer entryCreditBlockminuteNumberMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))	
	buf := new(primitives.Buffer)
	buf.WriteByte(m.Number)
	return buf.DeepCopyBytes(), nil
}

func (m *MinuteNumber) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	callTime := time.Now().UnixNano()
	defer entryCreditBlockminuteNumberUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))	
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
	callTime := time.Now().UnixNano()
	defer entryCreditBlockminuteNumberUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))	
	_, err = m.UnmarshalBinaryData(data)
	return
}

func (e *MinuteNumber) JSONByte() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer entryCreditBlockminuteNumberJSONByte.Observe(float64(time.Now().UnixNano() - callTime))	
	return primitives.EncodeJSON(e)
}

func (e *MinuteNumber) JSONString() (string, error) {
	callTime := time.Now().UnixNano()
	defer entryCreditBlockminuteNumberJSONString.Observe(float64(time.Now().UnixNano() - callTime))	
	return primitives.EncodeJSONString(e)
}

func (e *MinuteNumber) String() string {
	callTime := time.Now().UnixNano()
	defer entryCreditBlockminuteNumberString.Observe(float64(time.Now().UnixNano() - callTime))	
	var out primitives.Buffer
	out.WriteString(fmt.Sprintf(" %-20s\n", "MinuteNumber"))
	out.WriteString(fmt.Sprintf("   %-20s %d\n", "Number", e.Number))
	return (string)(out.DeepCopyBytes())
}

func (e *MinuteNumber) GetTimestamp() interfaces.Timestamp {
	return nil
}
