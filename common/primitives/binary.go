// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package primitives

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/FactomProject/ed25519"
	"github.com/FactomProject/factomd/common/interfaces"
	"time"
)

type Buffer struct {
	bytes.Buffer
}

func (b *Buffer) DeepCopyBytes() []byte {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesBufferDeepCopyBytes.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return b.Next(b.Len())
}

func NewBuffer(buf []byte) *Buffer {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesNewBuffer.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	tmp := new(Buffer)
	tmp.Buffer = *bytes.NewBuffer(buf)
	return tmp
}

func AreBytesEqual(b1, b2 []byte) bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesAreBytesEqual.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if len(b1) != len(b2) {
		return false
	}
	for i := range b1 {
		if b1[i] != b2[i] {
			return false
		}
	}
	return true
}

func AreBinaryMarshallablesEqual(b1, b2 interfaces.BinaryMarshallable) (bool, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesAreBinaryMarshallablesEqual.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if b1 == nil {
		if b2 == nil {
			return true, nil
		}
		return false, nil
	}
	if b2 == nil {
		return false, nil
	}

	bytes1, err := b1.MarshalBinary()
	if err != nil {
		return false, err
	}
	bytes2, err := b2.MarshalBinary()
	if err != nil {
		return false, err
	}
	return AreBytesEqual(bytes1, bytes2), nil
}

func EncodeBinary(bytes []byte) string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesEncodeBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return hex.EncodeToString(bytes)
}

func DecodeBinary(bytes string) ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesDecodeBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return hex.DecodeString(bytes)
}

type ByteSlice32 [32]byte

var _ interfaces.Printable = (*ByteSlice32)(nil)
var _ interfaces.BinaryMarshallable = (*ByteSlice32)(nil)

func StringToByteSlice32(s string) *ByteSlice32 {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesStringToByteSlice32.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	bin, err := DecodeBinary(s)
	if err != nil {
		return nil
	}
	bs := new(ByteSlice32)
	err = bs.UnmarshalBinary(bin)
	if err != nil {
		return nil
	}
	return bs
}

func Byte32ToByteSlice32(b [32]byte) *ByteSlice32 {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByte32ToByteSlice32.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	bs := new(ByteSlice32)
	err := bs.UnmarshalBinary(b[:])
	if err != nil {
		return nil
	}
	return bs
}

func (bs *ByteSlice32) MarshalBinary() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSlice32MarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return bs[:], nil
}

func (bs *ByteSlice32) Fixed() [32]byte {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSlice32Fixed.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return *bs
}

func (bs *ByteSlice32) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSlice32UnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error unmarshalling: %v", r)
		}
	}()
	if data == nil || len(data) < 32 {
		return nil, fmt.Errorf("Not enough data to unmarshal")
	}
	copy(bs[:], data[:32])
	newData = data[32:]
	return
}

func (bs *ByteSlice32) UnmarshalBinary(data []byte) (err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSlice32UnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	_, err = bs.UnmarshalBinaryData(data)
	return
}

func (e *ByteSlice32) JSONByte() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSlice32JSONByte.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return EncodeJSON(e)
}

func (e *ByteSlice32) JSONString() (string, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSlice32JSONString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return EncodeJSONString(e)
}

func (bs *ByteSlice32) String() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSlice32String.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return fmt.Sprintf("%x", bs[:])
}

func (bs *ByteSlice32) MarshalText() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSlice32MarshalText.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return []byte(bs.String()), nil
}

type ByteSlice64 [64]byte

var _ interfaces.Printable = (*ByteSlice64)(nil)
var _ interfaces.BinaryMarshallable = (*ByteSlice64)(nil)

func (bs *ByteSlice64) MarshalBinary() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSlice64MarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return bs[:], nil
}

func (bs *ByteSlice64) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSlice64UnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error unmarshalling: %v", r)
		}
	}()
	if data == nil || len(data) < 64 {
		return nil, fmt.Errorf("Not enough data to unmarshal")
	}
	copy(bs[:], data[:64])
	newData = data[64:]
	return
}

func (bs *ByteSlice64) UnmarshalBinary(data []byte) (err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSlice64UnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	_, err = bs.UnmarshalBinaryData(data)
	return
}

func (e *ByteSlice64) JSONByte() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSlice64JSONByte.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return EncodeJSON(e)
}

func (e *ByteSlice64) JSONString() (string, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSlice64JSONString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return EncodeJSONString(e)
}

func (bs *ByteSlice64) String() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSlice64String.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return fmt.Sprintf("%x", bs[:])
}

func (bs *ByteSlice64) MarshalText() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSlice64MarshalText.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return []byte(bs.String()), nil
}

type ByteSlice6 [6]byte

var _ interfaces.Printable = (*ByteSlice6)(nil)
var _ interfaces.BinaryMarshallable = (*ByteSlice6)(nil)

func (bs *ByteSlice6) MarshalBinary() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSlice6MarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return bs[:], nil
}

func (bs *ByteSlice6) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSlice6UnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error unmarshalling: %v", r)
		}
	}()
	if data == nil || len(data) < 6 {
		return nil, fmt.Errorf("Not enough data to unmarshal")
	}
	copy(bs[:], data[:6])
	newData = data[6:]
	return
}

func (bs *ByteSlice6) UnmarshalBinary(data []byte) (err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSlice6UnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	_, err = bs.UnmarshalBinaryData(data)
	return
}

func (e *ByteSlice6) JSONByte() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSlice6JSONByte.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return EncodeJSON(e)
}

func (e *ByteSlice6) JSONString() (string, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSlice6JSONString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return EncodeJSONString(e)
}

func (bs *ByteSlice6) String() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSlice6String.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return fmt.Sprintf("%x", bs[:])
}

func (bs *ByteSlice6) MarshalText() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSlice6MarshalText.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return []byte(bs.String()), nil
}

type ByteSliceSig [ed25519.SignatureSize]byte

var _ interfaces.Printable = (*ByteSliceSig)(nil)
var _ interfaces.BinaryMarshallable = (*ByteSliceSig)(nil)

func (bs *ByteSliceSig) MarshalBinary() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSliceSigMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return bs[:], nil
}

func (bs *ByteSliceSig) GetFixed() ([ed25519.SignatureSize]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSliceSigGetFixed.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	answer := [ed25519.SignatureSize]byte{}
	copy(answer[:], bs[:])

	return answer, nil
}

func (bs *ByteSliceSig) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSliceSigUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error unmarshalling: %v", r)
		}
	}()
	if data == nil || len(data) < ed25519.SignatureSize {
		return nil, fmt.Errorf("Not enough data to unmarshal")
	}
	copy(bs[:], data[:ed25519.SignatureSize])
	newData = data[ed25519.SignatureSize:]
	return
}

func (bs *ByteSliceSig) UnmarshalBinary(data []byte) (err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSliceSigUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if len(data) < ed25519.SignatureSize {
		return fmt.Errorf("Byte slice too short to unmarshal")
	}
	copy(bs[:], data[:ed25519.SignatureSize])
	return
}

func (e *ByteSliceSig) JSONByte() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSliceSigJSONByte.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return EncodeJSON(e)
}

func (e *ByteSliceSig) JSONString() (string, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSliceSigJSONString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return EncodeJSONString(e)
}

func (bs *ByteSliceSig) String() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSliceSigString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return fmt.Sprintf("%x", bs[:])
}

func (bs *ByteSliceSig) MarshalText() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSliceSigMarshalText.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return []byte(bs.String()), nil
}

func (bs *ByteSliceSig) UnmarshalText(text []byte) error {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSliceSigUnmarshalText.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	b, err := hex.DecodeString(string(text))
	if err != nil {
		return err
	}
	return bs.UnmarshalBinary(b)
}

type ByteSlice20 [20]byte

var _ interfaces.Printable = (*ByteSlice20)(nil)
var _ interfaces.BinaryMarshallable = (*ByteSlice20)(nil)

func (bs *ByteSlice20) MarshalBinary() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSlice20MarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return bs[:], nil
}

func (bs *ByteSlice20) GetFixed() ([20]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSlice20GetFixed.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	answer := [20]byte{}
	copy(answer[:], bs[:])

	return answer, nil
}

func (bs *ByteSlice20) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSlice20UnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error unmarshalling: %v", r)
		}
	}()
	if data == nil || len(data) < 20 {
		return nil, fmt.Errorf("Not enough data to unmarshal")
	}
	copy(bs[:], data[:20])
	newData = data[20:]
	return
}

func (bs *ByteSlice20) UnmarshalBinary(data []byte) (err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSlice20UnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	_, err = bs.UnmarshalBinaryData(data)
	return
}

func (e *ByteSlice20) JSONByte() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSlice20JSONByte.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return EncodeJSON(e)
}

func (e *ByteSlice20) JSONString() (string, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSlice20JSONString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return EncodeJSONString(e)
}

func (bs *ByteSlice20) String() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSlice20String.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return fmt.Sprintf("%x", bs[:])
}

func (bs *ByteSlice20) MarshalText() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSlice20MarshalText.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return []byte(bs.String()), nil
}

type ByteSlice struct {
	Bytes []byte
}

var _ interfaces.Printable = (*ByteSlice)(nil)
var _ interfaces.BinaryMarshallable = (*ByteSlice)(nil)
var _ interfaces.BinaryMarshallableAndCopyable = (*ByteSlice)(nil)

func StringToByteSlice(s string) *ByteSlice {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesStringToByteSlice.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	bin, err := DecodeBinary(s)
	if err != nil {
		return nil
	}
	bs := new(ByteSlice)
	err = bs.UnmarshalBinary(bin)
	if err != nil {
		return nil
	}
	return bs
}

func (bs *ByteSlice) New() interfaces.BinaryMarshallableAndCopyable {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSliceNew.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return new(ByteSlice)
}

func (bs *ByteSlice) MarshalBinary() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSliceMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return bs.Bytes[:], nil
}

func (bs *ByteSlice) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSliceUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error unmarshalling: %v", r)
		}
	}()
	bs.Bytes = make([]byte, len(data))
	copy(bs.Bytes[:], data)
	return nil, nil
}

func (bs *ByteSlice) UnmarshalBinary(data []byte) (err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSliceUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	_, err = bs.UnmarshalBinaryData(data)
	return
}

func (e *ByteSlice) JSONByte() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSliceJSONByte.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return EncodeJSON(e)
}

func (e *ByteSlice) JSONString() (string, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSliceJSONString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return EncodeJSONString(e)
}

func (bs *ByteSlice) String() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSliceString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return fmt.Sprintf("%x", bs.Bytes[:])
}

func (bs *ByteSlice) MarshalText() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdprimitivesByteSliceMarshalText.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return []byte(bs.String()), nil
}
