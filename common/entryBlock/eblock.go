// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package entryBlock

import (
	"fmt"

	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
)

// EBlock is the Entry Block. It holds the hashes of the Entries and its Merkle
// Root is written into the Directory Blocks. Each Entry Block represents all
// of the entries for a paticular Chain during a 10 minute period.
type EBlock struct {
	Header interfaces.IEntryBlockHeader
	Body   *EBlockBody
}

var _ interfaces.Printable = (*EBlock)(nil)
var _ interfaces.BinaryMarshallableAndCopyable = (*EBlock)(nil)
var _ interfaces.BinaryMarshallable = (*EBlock)(nil)
var _ interfaces.DatabaseBatchable = (*EBlock)(nil)
var _ interfaces.IEntryBlock = (*EBlock)(nil)
var _ interfaces.DatabaseBlockWithEntries = (*EBlock)(nil)

func (c *EBlock) Init() {
	if c.Header == nil {
		h := new(EBlockHeader)
		h.Init()
		c.Header = h
	}
	if c.Body == nil {
		c.Body = new(EBlockBody)
	}
}

func (c *EBlock) GetEntryHashes() []interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer entryBlockGetEntryHashes.Observe(float64(time.Now().UnixNano() - callTime))	
	return c.GetBody().GetEBEntries()
}

func (c *EBlock) GetEntrySigHashes() []interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer entryBlockGetEntrySigHashes.Observe(float64(time.Now().UnixNano() - callTime))	
	return nil
}

func (c *EBlock) New() interfaces.BinaryMarshallableAndCopyable {
	callTime := time.Now().UnixNano()
	defer entryBlockNew.Observe(float64(time.Now().UnixNano() - callTime))	
	return NewEBlock()
}

func (e *EBlock) GetWelds() [][]byte {
	callTime := time.Now().UnixNano()
	defer entryBlockGetWelds.Observe(float64(time.Now().UnixNano() - callTime))	
	e.Init()
	var answer [][]byte
	for _, entry := range e.GetEntryHashes() {
		answer = append(answer, primitives.DoubleSha(append(entry.Bytes(), e.GetChainID().Bytes()...)))
	}
	return answer
}

func (e *EBlock) GetWeldHashes() []interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer entryBlockGetWeldHashes.Observe(float64(time.Now().UnixNano() - callTime))	
	var answer []interfaces.IHash
	for _, h := range e.GetWelds() {
		hash := primitives.NewZeroHash()
		hash.SetBytes(h)
		answer = append(answer, hash)
	}
	return answer
}

func (c *EBlock) GetDatabaseHeight() uint32 {
	callTime := time.Now().UnixNano()
	defer entryBlockGetDatabaseHeight.Observe(float64(time.Now().UnixNano() - callTime))	
	return c.GetHeader().GetDBHeight()
}

func (c *EBlock) GetChainID() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer entryBlockGetChainID.Observe(float64(time.Now().UnixNano() - callTime))	
	return c.GetHeader().GetChainID()
}

func (c *EBlock) GetHashOfChainID() []byte {
	callTime := time.Now().UnixNano()
	defer entryBlockGetHashOfChainID.Observe(float64(time.Now().UnixNano() - callTime))	
	return primitives.DoubleSha(c.GetChainID().Bytes())
}

func (c *EBlock) GetHashOfChainIDHash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer entryBlockGetHashOfChainIDHash.Observe(float64(time.Now().UnixNano() - callTime))	
	hash := primitives.NewZeroHash()
	hash.SetBytes(c.GetHashOfChainID())
	return hash
}

func (c *EBlock) DatabasePrimaryIndex() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer entryBlockDatabasePrimaryIndex.Observe(float64(time.Now().UnixNano() - callTime))	
	key, _ := c.KeyMR()
	return key
}

func (c *EBlock) DatabaseSecondaryIndex() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer entryBlockDatabaseSecondaryIndex.Observe(float64(time.Now().UnixNano() - callTime))	
	h, _ := c.Hash()
	return h
}

func (c *EBlock) GetHeader() interfaces.IEntryBlockHeader {
	callTime := time.Now().UnixNano()
	defer entryBlockGetHeader.Observe(float64(time.Now().UnixNano() - callTime))	
	c.Init()
	return c.Header
}

func (c *EBlock) GetBody() interfaces.IEBlockBody {
	callTime := time.Now().UnixNano()
	defer entryBlockGetBody.Observe(float64(time.Now().UnixNano() - callTime))	
	c.Init()
	return c.Body
}

// AddEBEntry creates a new Entry Block Entry from the provided Factom Entry
// and adds it to the Entry Block Body.
func (e *EBlock) AddEBEntry(entry interfaces.IEBEntry) error {
	callTime := time.Now().UnixNano()
	defer entryBlockAddEBEntry.Observe(float64(time.Now().UnixNano() - callTime))	
	e.Init()
	e.GetBody().AddEBEntry(entry.GetHash())
	if err := e.BuildHeader(); err != nil {
		return err
	}
	return nil
}

// AddEndOfMinuteMarker adds the End of Minute to the Entry Block. The End of
// Minut byte becomes the last byte in a 32 byte slice that is added to the
// Entry Block Body as an Entry Block Entry.
func (e *EBlock) AddEndOfMinuteMarker(m byte) error {
	callTime := time.Now().UnixNano()
	defer entryBlockAddEndOfMinuteMarker.Observe(float64(time.Now().UnixNano() - callTime))	
	e.Init()
	e.GetBody().AddEndOfMinuteMarker(m)
	if err := e.BuildHeader(); err != nil {
		return err
	}
	return nil
}

// BuildHeader updates the Entry Block Header to include information about the
// Entry Block Body. BuildHeader should be run after the Entry Block Body has
// included all of its EntryEntries.
func (e *EBlock) BuildHeader() error {
	callTime := time.Now().UnixNano()
	defer entryBlockBuildHeader.Observe(float64(time.Now().UnixNano() - callTime))	
	e.GetHeader().SetBodyMR(e.GetBody().MR())
	e.GetHeader().SetEntryCount(uint32(len(e.GetEntryHashes())))
	return nil
}

func (e *EBlock) GetHash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer entryBlockGetHash.Observe(float64(time.Now().UnixNano() - callTime))	
	h, _ := e.Hash()
	return h
}

// Hash returns the simple Sha256 hash of the serialized Entry Block. Hash is
// used to provide the PrevFullHash to the next Entry Block in a Chain.
func (e *EBlock) Hash() (interfaces.IHash, error) {
	callTime := time.Now().UnixNano()
	defer entryBlockHash.Observe(float64(time.Now().UnixNano() - callTime))	
	p, err := e.MarshalBinary()
	if err != nil {
		return nil, err
	}
	return primitives.Sha(p), nil
}

func (e *EBlock) HeaderHash() (interfaces.IHash, error) {
	callTime := time.Now().UnixNano()
	defer entryBlockHeaderHash.Observe(float64(time.Now().UnixNano() - callTime))	
	e.BuildHeader()
	header, err := e.GetHeader().MarshalBinary()
	if err != nil {
		return nil, err
	}
	h := primitives.Sha(header)
	return h, nil
}

func (e *EBlock) BodyKeyMR() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer entryBlockBodyKeyMR.Observe(float64(time.Now().UnixNano() - callTime))	
	e.BuildHeader()
	return e.GetHeader().GetBodyMR()
}

// KeyMR returns the hash of the hash of the Entry Block Header concatinated
// with the Merkle Root of the Entry Block Body. The Body Merkle Root is
// calculated by the func (e *EBlockBody) MR() which is called by the func
// (e *EBlock) BuildHeader().
func (e *EBlock) KeyMR() (interfaces.IHash, error) {
	// Sha(Sha(header) + BodyMR)
	callTime := time.Now().UnixNano()
	defer entryBlockKeyMR.Observe(float64(time.Now().UnixNano() - callTime))	
	e.BuildHeader()
	h, err := e.HeaderHash()
	if err != nil {
		return nil, err
	}
	return primitives.Sha(append(h.Bytes(), e.GetHeader().GetBodyMR().Bytes()...)), nil
}

// MarshalBinary returns the serialized binary form of the Entry Block.
func (e *EBlock) MarshalBinary() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer entryBlockMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))	
	e.Init()
	buf := new(primitives.Buffer)

	if err := e.BuildHeader(); err != nil {
		return nil, err
	}
	if p, err := e.GetHeader().MarshalBinary(); err != nil {
		return nil, err
	} else {
		buf.Write(p)
	}

	if p, err := e.marshalBodyBinary(); err != nil {
		return nil, err
	} else {
		buf.Write(p)
	}

	return buf.DeepCopyBytes(), nil
}

func UnmarshalEBlock(data []byte) (interfaces.IEntryBlock, error) {
	callTime := time.Now().UnixNano()
	defer entryBlockUnmarshalEBlock.Observe(float64(time.Now().UnixNano() - callTime))	
	block, _, err := UnmarshalEBlockData(data)
	return block, err
}

func UnmarshalEBlockData(data []byte) (interfaces.IEntryBlock, []byte, error) {
	callTime := time.Now().UnixNano()
	defer entryBlockUnmarshalEBlockData.Observe(float64(time.Now().UnixNano() - callTime))	
	block := NewEBlock()

	data, err := block.UnmarshalBinaryData(data)
	if err != nil {
		return nil, nil, err
	}

	return block, data, nil
}

// UnmarshalBinary populates the Entry Block object from the serialized binary
// data.
func (e *EBlock) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	callTime := time.Now().UnixNano()
	defer entryBlockUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))	
	newData = data

	if e.Header == nil {
		e.Header = new(EBlockHeader)
	}

	newData, err = e.GetHeader().UnmarshalBinaryData(newData)
	if err != nil {
		return
	}

	newData, err = e.unmarshalBodyBinaryData(newData)
	if err != nil {
		return
	}

	return
}

func (e *EBlock) UnmarshalBinary(data []byte) (err error) {
	callTime := time.Now().UnixNano()
	defer entryBlockUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))	
	_, err = e.UnmarshalBinaryData(data)
	return
}

// marshalBodyBinary returns a serialized binary Entry Block Body
func (e *EBlock) marshalBodyBinary() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer entryBlockmarshalBodyBinary.Observe(float64(time.Now().UnixNano() - callTime))	
	e.Init()
	buf := new(primitives.Buffer)

	for _, v := range e.GetEntryHashes() {
		buf.Write(v.Bytes())
	}

	return buf.DeepCopyBytes(), nil
}

// unmarshalBodyBinary builds the Entry Block Body from the serialized binary.
func (e *EBlock) unmarshalBodyBinaryData(data []byte) (newData []byte, err error) {
	callTime := time.Now().UnixNano()
	defer entryBlockunmarshalBodyBinaryData.Observe(float64(time.Now().UnixNano() - callTime))	
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error unmarshalling: %v", r)
		}
	}()
	e.Init()

	buf := primitives.NewBuffer(data)
	hash := make([]byte, 32)

	for i := uint32(0); i < e.GetHeader().GetEntryCount(); i++ {
		if _, err := buf.Read(hash); err != nil {
			return nil, err
		}

		h := primitives.NewZeroHash()
		h.SetBytes(hash)
		e.GetBody().AddEBEntry(h)
	}

	newData = buf.DeepCopyBytes()
	return
}

func (e *EBlock) unmarshalBodyBinary(data []byte) (err error) {
	callTime := time.Now().UnixNano()
	defer entryBlockunmarshalBodyBinary.Observe(float64(time.Now().UnixNano() - callTime))	
	_, err = e.unmarshalBodyBinaryData(data)
	return
}

func (e *EBlock) JSONByte() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer entryBlockJSONByte.Observe(float64(time.Now().UnixNano() - callTime))	
	return primitives.EncodeJSON(e)
}

func (e *EBlock) JSONString() (string, error) {
	callTime := time.Now().UnixNano()
	defer entryBlockJSONString.Observe(float64(time.Now().UnixNano() - callTime))	
	return primitives.EncodeJSONString(e)
}

func (e *EBlock) String() string {
	callTime := time.Now().UnixNano()
	defer entryBlockString.Observe(float64(time.Now().UnixNano() - callTime))	
	e.Init()
	str := e.GetHeader().String()
	str = str + e.GetBody().String()
	return str
}

/*****************************************************
 * Support Routines
 *****************************************************/

// NewEBlock returns a blank initialized Entry Block with all of its fields
// zeroed.
func NewEBlock() *EBlock {
	callTime := time.Now().UnixNano()
	defer entryBlockNewEBlock.Observe(float64(time.Now().UnixNano() - callTime))	
	e := new(EBlock)
	e.Header = NewEBlockHeader()
	e.Body = NewEBlockBody()
	return e
}
