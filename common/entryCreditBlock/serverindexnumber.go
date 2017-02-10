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
	ServerIndexNumberSize = 1
)

type ServerIndexNumber struct {
	ServerIndexNumber uint8
}

var _ interfaces.Printable = (*ServerIndexNumber)(nil)
var _ interfaces.BinaryMarshallable = (*ServerIndexNumber)(nil)
var _ interfaces.ShortInterpretable = (*ServerIndexNumber)(nil)
var _ interfaces.IECBlockEntry = (*ServerIndexNumber)(nil)

func (e *ServerIndexNumber) String() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockServerIndexNumberString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	var out primitives.Buffer
	out.WriteString(fmt.Sprintf(" %-20s\n", "ServerIndexNumber"))
	out.WriteString(fmt.Sprintf("   %-20s %d\n", "Number", e.ServerIndexNumber))
	return (string)(out.DeepCopyBytes())
}

func (e *ServerIndexNumber) Hash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockServerIndexNumberHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	bin, err := e.MarshalBinary()
	if err != nil {
		panic(err)
	}
	return primitives.Sha(bin)
}

func (e *ServerIndexNumber) GetHash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockServerIndexNumberGetHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return e.Hash()
}

func (a *ServerIndexNumber) GetEntryHash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockServerIndexNumberGetEntryHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return nil
}

func (e *ServerIndexNumber) GetSigHash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockServerIndexNumberGetSigHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return nil
}

func (b *ServerIndexNumber) IsInterpretable() bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockServerIndexNumberIsInterpretable.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return true
}

func (b *ServerIndexNumber) Interpret() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockServerIndexNumberInterpret.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return fmt.Sprintf("ServerIndexNumber %v", b.ServerIndexNumber)
}

func NewServerIndexNumber() *ServerIndexNumber {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockNewServerIndexNumber.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return new(ServerIndexNumber)
}

func NewServerIndexNumber2(number uint8) *ServerIndexNumber {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockNewServerIndexNumber2.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	sin := new(ServerIndexNumber)
	sin.ServerIndexNumber = number
	return sin
}

func (s *ServerIndexNumber) ECID() byte {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockServerIndexNumberECID.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return ECIDServerIndexNumber
}

func (s *ServerIndexNumber) MarshalBinary() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockServerIndexNumberMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	buf := new(primitives.Buffer)
	buf.WriteByte(s.ServerIndexNumber)
	return buf.DeepCopyBytes(), nil
}

func (s *ServerIndexNumber) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockServerIndexNumberUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error unmarshalling ServerIndexNumber: %v", r)
		}
	}()

	buf := primitives.NewBuffer(data)
	var c byte
	if c, err = buf.ReadByte(); err != nil {
		return
	} else {
		s.ServerIndexNumber = c
	}
	newData = buf.DeepCopyBytes()
	return
}

func (s *ServerIndexNumber) UnmarshalBinary(data []byte) (err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockServerIndexNumberUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	_, err = s.UnmarshalBinaryData(data)
	return
}

func (e *ServerIndexNumber) JSONByte() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockServerIndexNumberJSONByte.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSON(e)
}

func (e *ServerIndexNumber) JSONString() (string, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockServerIndexNumberJSONString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSONString(e)
}

func (e *ServerIndexNumber) GetTimestamp() interfaces.Timestamp {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockServerIndexNumberGetTimestamp.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return nil
}
