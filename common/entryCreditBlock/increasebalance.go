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

type IncreaseBalance struct {
	ECPubKey *primitives.ByteSlice32
	TXID     interfaces.IHash
	Index    uint64
	NumEC    uint64
}

var _ interfaces.Printable = (*IncreaseBalance)(nil)

var _ interfaces.BinaryMarshallable = (*IncreaseBalance)(nil)
var _ interfaces.ShortInterpretable = (*IncreaseBalance)(nil)
var _ interfaces.IECBlockEntry = (*IncreaseBalance)(nil)

func (e *IncreaseBalance) Init() {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockIncreaseBalanceInit.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if e.ECPubKey == nil {
		e.ECPubKey = new(primitives.ByteSlice32)
	}
	if e.TXID == nil {
		e.TXID = primitives.NewZeroHash()
	}
}

func (e *IncreaseBalance) String() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockIncreaseBalanceString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	e.Init()
	var out primitives.Buffer
	out.WriteString(fmt.Sprintf(" %-20s\n", "IncreaseBalance"))
	out.WriteString(fmt.Sprintf("   %-20s %x\n", "ECPubKey", e.ECPubKey[:3]))
	out.WriteString(fmt.Sprintf("   %-20s %x\n", "TXID", e.TXID.Bytes()[:3]))
	out.WriteString(fmt.Sprintf("   %-20s %d\n", "Index", e.Index))
	out.WriteString(fmt.Sprintf("   %-20s %x\n", "NumEC", e.NumEC))

	return (string)(out.DeepCopyBytes())
}

func NewIncreaseBalance() *IncreaseBalance {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockNewIncreaseBalance.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	r := new(IncreaseBalance)
	r.Init()
	return r
}

func (a *IncreaseBalance) GetEntryHash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockIncreaseBalanceGetEntryHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return nil
}

func (e *IncreaseBalance) Hash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockIncreaseBalanceHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	bin, err := e.MarshalBinary()
	if err != nil {
		panic(err)
	}
	return primitives.Sha(bin)
}

func (e *IncreaseBalance) GetHash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockIncreaseBalanceGetHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return e.Hash()
}

func (e *IncreaseBalance) GetSigHash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockIncreaseBalanceGetSigHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return nil
}

func (b *IncreaseBalance) ECID() byte {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockIncreaseBalanceECID.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return ECIDBalanceIncrease
}

func (b *IncreaseBalance) IsInterpretable() bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockIncreaseBalanceIsInterpretable.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return false
}

func (b *IncreaseBalance) Interpret() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockIncreaseBalanceInterpret.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return ""
}

func (b *IncreaseBalance) MarshalBinary() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockIncreaseBalanceMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	b.Init()
	buf := new(primitives.Buffer)

	buf.Write(b.ECPubKey[:])

	buf.Write(b.TXID.Bytes())

	primitives.EncodeVarInt(buf, b.Index)

	primitives.EncodeVarInt(buf, b.NumEC)

	return buf.DeepCopyBytes(), nil
}

func (b *IncreaseBalance) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockIncreaseBalanceUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error unmarshalling IncreaseBalance: %v", r)
		}
	}()

	buf := primitives.NewBuffer(data)
	hash := make([]byte, 32)

	_, err = buf.Read(hash)
	if err != nil {
		return
	}
	b.ECPubKey = new(primitives.ByteSlice32)
	copy(b.ECPubKey[:], hash)

	_, err = buf.Read(hash)
	if err != nil {
		return
	}
	if b.TXID == nil {
		b.TXID = primitives.NewZeroHash()
	}
	b.TXID.SetBytes(hash)

	tmp := make([]byte, 0)
	b.Index, tmp = primitives.DecodeVarInt(buf.DeepCopyBytes())

	b.NumEC, tmp = primitives.DecodeVarInt(tmp)

	newData = tmp
	return
}

func (b *IncreaseBalance) UnmarshalBinary(data []byte) (err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockIncreaseBalanceUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	_, err = b.UnmarshalBinaryData(data)
	return
}

func (e *IncreaseBalance) JSONByte() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockIncreaseBalanceJSONByte.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSON(e)
}

func (e *IncreaseBalance) JSONString() (string, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockIncreaseBalanceJSONString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSONString(e)
}

func (e *IncreaseBalance) GetTimestamp() interfaces.Timestamp {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockIncreaseBalanceGetTimestamp.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return nil
}
