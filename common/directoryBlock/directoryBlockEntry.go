// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package directoryBlock

import (
	"fmt"

	"github.com/FactomProject/factomd/common/constants"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
)

var _ = fmt.Print

type DBEntry struct {
	ChainID interfaces.IHash
	KeyMR   interfaces.IHash // Different MR in EBlockHeader
}

var _ interfaces.Printable = (*DBEntry)(nil)
var _ interfaces.BinaryMarshallable = (*DBEntry)(nil)
var _ interfaces.IDBEntry = (*DBEntry)(nil)

func (c *DBEntry) GetChainID() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer directoryBlockEntryGetChainID.Observe(float64(time.Now().UnixNano() - callTime))	
	return c.ChainID
}

func (c *DBEntry) SetChainID(chainID interfaces.IHash) {
	callTime := time.Now().UnixNano()
	defer directoryBlockEntrySetChainID.Observe(float64(time.Now().UnixNano() - callTime))	
	c.ChainID = chainID
}

func (c *DBEntry) GetKeyMR() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer directoryBlockEntryGetKeyMR.Observe(float64(time.Now().UnixNano() - callTime))	
	return c.KeyMR
}

func (c *DBEntry) SetKeyMR(keyMR interfaces.IHash) {
	callTime := time.Now().UnixNano()
	defer directoryBlockEntrySetKeyMR.Observe(float64(time.Now().UnixNano() - callTime))	
	c.KeyMR = keyMR
}

func (e *DBEntry) MarshalBinary() (data []byte, err error) {
	callTime := time.Now().UnixNano()
	defer directoryBlockEntryMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))	
	var buf primitives.Buffer

	data, err = e.ChainID.MarshalBinary()
	if err != nil {
		return
	}
	buf.Write(data)

	if e.KeyMR == nil {
		data, err = primitives.NewHash(constants.ZERO_HASH).MarshalBinary()
	} else {
		data, err = e.KeyMR.MarshalBinary()
	}
	if err != nil {
		return
	}
	buf.Write(data)

	return buf.DeepCopyBytes(), nil
}

func (e *DBEntry) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	callTime := time.Now().UnixNano()
	defer directoryBlockEntryUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))	
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error unmarshalling Directory Block Entry: %v", r)
		}
	}()
	newData = data
	e.ChainID = new(primitives.Hash)
	newData, err = e.ChainID.UnmarshalBinaryData(newData)
	if err != nil {
		return
	}

	e.KeyMR = new(primitives.Hash)
	newData, err = e.KeyMR.UnmarshalBinaryData(newData)
	if err != nil {
		return
	}

	return
}

func (e *DBEntry) UnmarshalBinary(data []byte) (err error) {
	callTime := time.Now().UnixNano()
	defer directoryBlockEntryUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))	
	_, err = e.UnmarshalBinaryData(data)
	return
}

func (e *DBEntry) ShaHash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer directoryBlockEntryShaHash.Observe(float64(time.Now().UnixNano() - callTime))	
	byteArray, _ := e.MarshalBinary()
	return primitives.Sha(byteArray)
}

func (e *DBEntry) JSONByte() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer directoryBlockEntryJSONByte.Observe(float64(time.Now().UnixNano() - callTime))	
	return primitives.EncodeJSON(e)
}

func (e *DBEntry) JSONString() (string, error) {
	callTime := time.Now().UnixNano()
	defer directoryBlockEntryJSONString.Observe(float64(time.Now().UnixNano() - callTime))	
	return primitives.EncodeJSONString(e)
}

func (e *DBEntry) String() string {
	callTime := time.Now().UnixNano()
	defer directoryBlockEntryString.Observe(float64(time.Now().UnixNano() - callTime))	
	var out primitives.Buffer
	out.WriteString("ChainID: " + e.GetChainID().String() + "\n")
	out.WriteString("      KeyMR:   " + e.GetKeyMR().String() + "\n")
	return (string)(out.DeepCopyBytes())
}
