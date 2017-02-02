// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package directoryBlock

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/FactomProject/factomd/common/constants"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
)

var _ = fmt.Print

type DirectoryBlock struct {
	//Marshalized
	Header    interfaces.IDirectoryBlockHeader
	DBEntries []interfaces.IDBEntry

	//Not Marshalized
	DBHash interfaces.IHash
	KeyMR  interfaces.IHash
}

var _ interfaces.Printable = (*DirectoryBlock)(nil)
var _ interfaces.BinaryMarshallableAndCopyable = (*DirectoryBlock)(nil)
var _ interfaces.IDirectoryBlock = (*DirectoryBlock)(nil)
var _ interfaces.DatabaseBatchable = (*DirectoryBlock)(nil)
var _ interfaces.DatabaseBlockWithEntries = (*DirectoryBlock)(nil)

func (c *DirectoryBlock) Init() {
	if c.Header == nil {
		h := new(DBlockHeader)
		h.Init()
		c.Header = h
	}
}

func (c *DirectoryBlock) SetEntryHash(hash, chainID interfaces.IHash, index int) {
	callTime := time.Now().UnixNano()
	defer directoryBlockSetEntryHash.Observe(float64(time.Now().UnixNano() - callTime))
	if len(c.DBEntries) < index {
		ent := make([]interfaces.IDBEntry, index)
		copy(ent, c.DBEntries)
		c.DBEntries = ent
	}
	dbe := new(DBEntry)
	dbe.ChainID = chainID
	dbe.KeyMR = hash
	c.DBEntries[index] = dbe
}

func (c *DirectoryBlock) SetABlockHash(aBlock interfaces.IAdminBlock) error {
	callTime := time.Now().UnixNano()
	defer directoryBlockSetABlockHash.Observe(float64(time.Now().UnixNano() - callTime))
	hash := aBlock.DatabasePrimaryIndex()
	c.SetEntryHash(hash, aBlock.GetChainID(), 0)
	return nil
}

func (c *DirectoryBlock) SetECBlockHash(ecBlock interfaces.IEntryCreditBlock) error {
	callTime := time.Now().UnixNano()
	defer directoryBlockSetECBlockHash.Observe(float64(time.Now().UnixNano() - callTime))
	hash := ecBlock.DatabasePrimaryIndex()
	c.SetEntryHash(hash, ecBlock.GetChainID(), 1)
	return nil
}

func (c *DirectoryBlock) SetFBlockHash(fBlock interfaces.IFBlock) error {
	callTime := time.Now().UnixNano()
	defer directorySetFBlockHash.Observe(float64(time.Now().UnixNano() - callTime))
	hash := fBlock.DatabasePrimaryIndex()
	c.SetEntryHash(hash, fBlock.GetChainID(), 2)
	return nil
}

func (c *DirectoryBlock) GetEntryHashes() []interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer directoryBlockGetEntryHashes.Observe(float64(time.Now().UnixNano() - callTime))
	entries := c.DBEntries[:]
	answer := make([]interfaces.IHash, len(entries))
	for i, entry := range entries {
		answer[i] = entry.GetKeyMR()
	}
	return answer
}

func (c *DirectoryBlock) GetEntrySigHashes() []interfaces.IHash {
	return nil
}

func (c *DirectoryBlock) Sort() {
	callTime := time.Now().UnixNano()
	defer directoryBlockSort.Observe(float64(time.Now().UnixNano() - callTime))
	done := false
	for i := 3; !done && i < len(c.DBEntries)-1; i++ {
		done = true
		for j := 3; j < len(c.DBEntries)-1-i+3; j++ {
			comp := bytes.Compare(c.DBEntries[j].GetChainID().Bytes(),
				c.DBEntries[j+1].GetChainID().Bytes())
			if comp > 0 {
				h := c.DBEntries[j]
				c.DBEntries[j] = c.DBEntries[j+1]
				c.DBEntries[j+1] = h
				done = false
			}
		}
	}
}

func (c *DirectoryBlock) GetEntryHashesForBranch() []interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer directoryBlockGetEntryHashesForBranch.Observe(float64(time.Now().UnixNano() - callTime))
	entries := c.DBEntries[:]
	answer := make([]interfaces.IHash, 2*len(entries))
	for i, entry := range entries {
		answer[2*i] = entry.GetChainID()
		answer[2*i+1] = entry.GetKeyMR()
	}
	return answer
}

func (c *DirectoryBlock) GetDBEntries() []interfaces.IDBEntry {
	callTime := time.Now().UnixNano()
	defer directoryBlockGetDBEntries.Observe(float64(time.Now().UnixNano() - callTime))
	return c.DBEntries
}

func (c *DirectoryBlock) GetEBlockDBEntries() []interfaces.IDBEntry {
	callTime := time.Now().UnixNano()
	defer directoryBlockGetEBlockDBEntries.Observe(float64(time.Now().UnixNano() - callTime))
	answer := []interfaces.IDBEntry{}
	for _, v := range c.DBEntries {
		if v.GetChainID().String() == "000000000000000000000000000000000000000000000000000000000000000a" {
			continue
		}
		if v.GetChainID().String() == "000000000000000000000000000000000000000000000000000000000000000f" {
			continue
		}
		if v.GetChainID().String() == "000000000000000000000000000000000000000000000000000000000000000c" {
			continue
		}
		answer = append(answer, v)
	}
	return answer
}

func (c *DirectoryBlock) GetKeyMR() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer directoryBlockGetKeyMR.Observe(float64(time.Now().UnixNano() - callTime))
	keyMR, err := c.BuildKeyMerkleRoot()
	if err != nil {
		panic("Failed to build the key MR")
	}

	c.KeyMR = keyMR

	return c.KeyMR
}

func (c *DirectoryBlock) GetHeader() interfaces.IDirectoryBlockHeader {
	callTime := time.Now().UnixNano()
	defer directoryBlockGetHeader.Observe(float64(time.Now().UnixNano() - callTime))
	c.Init()
	return c.Header
}

func (c *DirectoryBlock) SetHeader(header interfaces.IDirectoryBlockHeader) {
	callTime := time.Now().UnixNano()
	defer directoryBlockSetHeader.Observe(float64(time.Now().UnixNano() - callTime))
	c.Header = header
}

func (c *DirectoryBlock) SetDBEntries(dbEntries []interfaces.IDBEntry) error {
	callTime := time.Now().UnixNano()
	defer directoryBlockSetDBEntries.Observe(float64(time.Now().UnixNano() - callTime))
	c.DBEntries = dbEntries
	c.GetHeader().SetBlockCount(uint32(len(dbEntries)))
	_, err := c.BuildBodyMR()
	if err != nil {
		return err
	}
	return nil
}

func (c *DirectoryBlock) New() interfaces.BinaryMarshallableAndCopyable {
	callTime := time.Now().UnixNano()
	defer directoryBlockNew.Observe(float64(time.Now().UnixNano() - callTime))
	dBlock := new(DirectoryBlock)
	dBlock.Header = NewDBlockHeader()
	dBlock.DBHash = primitives.NewZeroHash()
	dBlock.KeyMR = primitives.NewZeroHash()
	return dBlock
}

func (c *DirectoryBlock) GetDatabaseHeight() uint32 {
	callTime := time.Now().UnixNano()
	defer directoryBlockGetDatabaseHeight.Observe(float64(time.Now().UnixNano() - callTime))
	c.Init()
	return c.GetHeader().GetDBHeight()
}

func (c *DirectoryBlock) GetChainID() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer directoryBlockGetChainID.Observe(float64(time.Now().UnixNano() - callTime))
	return primitives.NewHash(constants.D_CHAINID)
}

func (c *DirectoryBlock) DatabasePrimaryIndex() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer directoryBlockDatabasePrimaryIndex.Observe(float64(time.Now().UnixNano() - callTime))
	return c.GetKeyMR()
}

func (c *DirectoryBlock) DatabaseSecondaryIndex() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer directoryBlockDatabaseSecondaryIndex.Observe(float64(time.Now().UnixNano() - callTime))
	return c.GetHash()
}

func (e *DirectoryBlock) JSONByte() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer directoryBlockJSONByte.Observe(float64(time.Now().UnixNano() - callTime))
	return primitives.EncodeJSON(e)
}

func (e *DirectoryBlock) JSONString() (string, error) {
	callTime := time.Now().UnixNano()
	defer directoryBlockJSONString.Observe(float64(time.Now().UnixNano() - callTime))
	return primitives.EncodeJSONString(e)
}

func (e *DirectoryBlock) String() string {
	callTime := time.Now().UnixNano()
	defer directoryBlockString.Observe(float64(time.Now().UnixNano() - callTime))
	e.Init()
	var out primitives.Buffer

	kmr := e.GetKeyMR()
	out.WriteString(fmt.Sprintf("%20s %v\n", "KeyMR:", kmr.String()))

	kmr = e.BodyKeyMR()
	out.WriteString(fmt.Sprintf("%20s %v\n", "BodyMR:", kmr.String()))

	fh := e.GetFullHash()
	out.WriteString(fmt.Sprintf("%20s %v\n", "FullHash:", fh.String()))

	out.WriteString(e.GetHeader().String())
	out.WriteString("Entries: \n")
	for i, entry := range e.DBEntries {
		out.WriteString(fmt.Sprintf("%5d %s", i, entry.String()))
	}

	return (string)(out.DeepCopyBytes())

}

func (b *DirectoryBlock) MarshalBinary() (data []byte, err error) {
	callTime := time.Now().UnixNano()
	defer directoryBlockMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	b.Init()
	var buf primitives.Buffer

	b.Sort()

	b.BuildBodyMR()

	count := uint32(len(b.GetDBEntries()))
	b.GetHeader().SetBlockCount(count)

	data, err = b.GetHeader().MarshalBinary()
	if err != nil {
		return
	}
	buf.Write(data)

	for i := uint32(0); i < count; i++ {
		data, err = b.GetDBEntries()[i].MarshalBinary()
		if err != nil {
			return
		}
		buf.Write(data)
	}

	return buf.DeepCopyBytes(), err
}

func (b *DirectoryBlock) BuildBodyMR() (interfaces.IHash, error) {
	callTime := time.Now().UnixNano()
	defer directoryBlockBuildBodyMR.Observe(float64(time.Now().UnixNano() - callTime))
	hashes := make([]interfaces.IHash, len(b.GetDBEntries()))
	for i, entry := range b.GetDBEntries() {
		data, err := entry.MarshalBinary()
		if err != nil {
			return nil, err
		}
		hashes[i] = primitives.Sha(data)
	}

	if len(hashes) == 0 {
		hashes = append(hashes, primitives.Sha(nil))
	}

	merkleTree := primitives.BuildMerkleTreeStore(hashes)
	merkleRoot := merkleTree[len(merkleTree)-1]

	b.GetHeader().SetBodyMR(merkleRoot)

	return merkleRoot, nil
}

func (b *DirectoryBlock) HeaderHash() (interfaces.IHash, error) {
	callTime := time.Now().UnixNano()
	defer directoryBlockHeaderHash.Observe(float64(time.Now().UnixNano() - callTime))
	binaryEBHeader, err := b.GetHeader().MarshalBinary()
	if err != nil {
		return nil, err
	}
	return primitives.Sha(binaryEBHeader), nil
}

func (b *DirectoryBlock) BodyKeyMR() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer directoryBlockBodyKeyMR.Observe(float64(time.Now().UnixNano() - callTime))
	key, _ := b.BuildBodyMR()
	return key
}

func (b *DirectoryBlock) BuildKeyMerkleRoot() (keyMR interfaces.IHash, err error) {
	callTime := time.Now().UnixNano()
	defer directoryBlokcBuildKeyMerkleRoot.Observe(float64(time.Now().UnixNano() - callTime))
	// Create the Entry Block Key Merkle Root from the hash of Header and the Body Merkle Root

	hashes := make([]interfaces.IHash, 0, 2)
	bodyKeyMR := b.BodyKeyMR() //This needs to be called first to build the header properly!!
	headerHash, err := b.HeaderHash()
	if err != nil {
		return nil, err
	}
	hashes = append(hashes, headerHash)
	hashes = append(hashes, bodyKeyMR)
	merkle := primitives.BuildMerkleTreeStore(hashes)
	keyMR = merkle[len(merkle)-1] // MerkleRoot is not marshalized in Dir Block

	b.KeyMR = keyMR

	b.GetFullHash() // Create the Full Hash when we create the keyMR

	return primitives.NewHash(keyMR.Bytes()), nil
}

func UnmarshalDBlock(data []byte) (interfaces.IDirectoryBlock, error) {
	callTime := time.Now().UnixNano()
	defer directoryBlockUnmarshalDBlock.Observe(float64(time.Now().UnixNano() - callTime))
	dBlock := new(DirectoryBlock)
	dBlock.Header = NewDBlockHeader()
	dBlock.DBHash = primitives.NewZeroHash()
	dBlock.KeyMR = primitives.NewZeroHash()
	err := dBlock.UnmarshalBinary(data)
	if err != nil {
		return nil, err
	}
	return dBlock, nil
}

func (b *DirectoryBlock) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	callTime := time.Now().UnixNano()
	defer directoryBlockUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error unmarshalling Directory Block: %v", r)
		}
	}()

	newData = data

	var fbh interfaces.IDirectoryBlockHeader = new(DBlockHeader)

	newData, err = fbh.UnmarshalBinaryData(newData)
	if err != nil {
		return
	}
	b.SetHeader(fbh)

	count := b.GetHeader().GetBlockCount()
	entries := make([]interfaces.IDBEntry, count)
	for i := uint32(0); i < count; i++ {
		entries[i] = new(DBEntry)
		newData, err = entries[i].UnmarshalBinaryData(newData)
		if err != nil {
			return
		}
	}

	err = b.SetDBEntries(entries)
	if err != nil {
		return
	}

	return
}

func (h *DirectoryBlock) GetTimestamp() interfaces.Timestamp {
	callTime := time.Now().UnixNano()
	defer directoryBlockGetTimestamp.Observe(float64(time.Now().UnixNano() - callTime))
	return h.GetHeader().GetTimestamp()
}

func (b *DirectoryBlock) UnmarshalBinary(data []byte) (err error) {
	callTime := time.Now().UnixNano()
	defer directoryBlockUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	_, err = b.UnmarshalBinaryData(data)
	return
}

func (b *DirectoryBlock) GetHash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer directoryBlockGetHash.Observe(float64(time.Now().UnixNano() - callTime))
	return b.GetFullHash()
}

func (b *DirectoryBlock) GetFullHash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer directoryBlockGetFullHash.Observe(float64(time.Now().UnixNano() - callTime))
	binaryDblock, err := b.MarshalBinary()
	if err != nil {
		return nil
	}
	b.DBHash = primitives.Sha(binaryDblock)
	return b.DBHash
}

func (b *DirectoryBlock) AddEntry(chainID interfaces.IHash, keyMR interfaces.IHash) error {
	callTime := time.Now().UnixNano()
	defer directoryBlockAddEntry.Observe(float64(time.Now().UnixNano() - callTime))
	var dbentry interfaces.IDBEntry
	dbentry = new(DBEntry)
	dbentry.SetChainID(chainID)
	dbentry.SetKeyMR(keyMR)

	if b.DBEntries == nil {
		b.DBEntries = []interfaces.IDBEntry{}
	}

	return b.SetDBEntries(append(b.DBEntries, dbentry))
}

/*********************************************************************
 * Support
 *********************************************************************/

func NewDirectoryBlock(prev interfaces.IDirectoryBlock) interfaces.IDirectoryBlock {
	callTime := time.Now().UnixNano()
	defer directoryBlockNewDirectoryBlock.Observe(float64(time.Now().UnixNano() - callTime))
	newdb := new(DirectoryBlock)

	newdb.Header = new(DBlockHeader)
	newdb.GetHeader().SetVersion(constants.VERSION_0)

	if prev != nil {
		newdb.GetHeader().SetPrevFullHash(prev.GetFullHash())
		newdb.GetHeader().SetPrevKeyMR(prev.GetKeyMR())
		newdb.GetHeader().SetDBHeight(prev.GetHeader().GetDBHeight() + 1)
	} else {
		newdb.GetHeader().SetPrevFullHash(primitives.NewZeroHash())
		newdb.GetHeader().SetPrevKeyMR(primitives.NewZeroHash())
		newdb.GetHeader().SetDBHeight(0)
	}

	newdb.SetDBEntries(make([]interfaces.IDBEntry, 0))

	newdb.AddEntry(primitives.NewHash(constants.ADMIN_CHAINID), primitives.NewZeroHash())
	newdb.AddEntry(primitives.NewHash(constants.EC_CHAINID), primitives.NewZeroHash())
	newdb.AddEntry(primitives.NewHash(constants.FACTOID_CHAINID), primitives.NewZeroHash())

	return newdb
}

func CheckBlockPairIntegrity(block interfaces.IDirectoryBlock, prev interfaces.IDirectoryBlock) error {
	callTime := time.Now().UnixNano()
	defer directoryBlockCheckBlockPairIntegrity.Observe(float64(time.Now().UnixNano() - callTime))
	if block == nil {
		return fmt.Errorf("No block specified")
	}

	if prev == nil {
		if block.GetHeader().GetPrevKeyMR().IsZero() == false {
			return fmt.Errorf("Invalid PrevKeyMR")
		}
		if block.GetHeader().GetPrevFullHash().IsZero() == false {
			return fmt.Errorf("Invalid PrevFullHash")
		}
		if block.GetHeader().GetDBHeight() != 0 {
			return fmt.Errorf("Invalid DBHeight")
		}
	} else {
		if block.GetHeader().GetPrevKeyMR().IsSameAs(prev.GetKeyMR()) == false {
			return fmt.Errorf("Invalid PrevKeyMR")
		}
		if block.GetHeader().GetPrevFullHash().IsSameAs(prev.GetFullHash()) == false {
			return fmt.Errorf("Invalid PrevFullHash")
		}
		if block.GetHeader().GetDBHeight() != (prev.GetHeader().GetDBHeight() + 1) {
			return fmt.Errorf("Invalid DBHeight")
		}
	}

	return nil
}

type ExpandedDBlock DirectoryBlock

func (e DirectoryBlock) MarshalJSON() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer directoryBlockMarshalJSON.Observe(float64(time.Now().UnixNano() - callTime))
	e.GetKeyMR()
	e.GetFullHash()

	return json.Marshal(ExpandedDBlock(e))
}
