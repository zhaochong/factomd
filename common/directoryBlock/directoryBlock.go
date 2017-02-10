// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package directoryBlock

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/FactomProject/factomd/common/constants"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
	"time"
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
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddirectoryBlockDirectoryBlockInit.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if c.Header == nil {
		h := new(DBlockHeader)
		h.Init()
		c.Header = h
	}
}

func (c *DirectoryBlock) SetEntryHash(hash, chainID interfaces.IHash, index int) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddirectoryBlockDirectoryBlockSetEntryHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

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
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddirectoryBlockDirectoryBlockSetABlockHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	hash := aBlock.DatabasePrimaryIndex()
	c.SetEntryHash(hash, aBlock.GetChainID(), 0)
	return nil
}

func (c *DirectoryBlock) SetECBlockHash(ecBlock interfaces.IEntryCreditBlock) error {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddirectoryBlockDirectoryBlockSetECBlockHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	hash := ecBlock.DatabasePrimaryIndex()
	c.SetEntryHash(hash, ecBlock.GetChainID(), 1)
	return nil
}

func (c *DirectoryBlock) SetFBlockHash(fBlock interfaces.IFBlock) error {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddirectoryBlockDirectoryBlockSetFBlockHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	hash := fBlock.DatabasePrimaryIndex()
	c.SetEntryHash(hash, fBlock.GetChainID(), 2)
	return nil
}

func (c *DirectoryBlock) GetEntryHashes() []interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddirectoryBlockDirectoryBlockGetEntryHashes.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	entries := c.DBEntries[:]
	answer := make([]interfaces.IHash, len(entries))
	for i, entry := range entries {
		answer[i] = entry.GetKeyMR()
	}
	return answer
}

func (c *DirectoryBlock) GetEntrySigHashes() []interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddirectoryBlockDirectoryBlockGetEntrySigHashes.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return nil
}

func (c *DirectoryBlock) Sort() {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddirectoryBlockDirectoryBlockSort.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

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
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddirectoryBlockDirectoryBlockGetEntryHashesForBranch.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	entries := c.DBEntries[:]
	answer := make([]interfaces.IHash, 2*len(entries))
	for i, entry := range entries {
		answer[2*i] = entry.GetChainID()
		answer[2*i+1] = entry.GetKeyMR()
	}
	return answer
}

func (c *DirectoryBlock) GetDBEntries() []interfaces.IDBEntry {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddirectoryBlockDirectoryBlockGetDBEntries.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return c.DBEntries
}

func (c *DirectoryBlock) GetEBlockDBEntries() []interfaces.IDBEntry {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddirectoryBlockDirectoryBlockGetEBlockDBEntries.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

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
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddirectoryBlockDirectoryBlockGetKeyMR.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	keyMR, err := c.BuildKeyMerkleRoot()
	if err != nil {
		panic("Failed to build the key MR")
	}

	c.KeyMR = keyMR

	return c.KeyMR
}

func (c *DirectoryBlock) GetHeader() interfaces.IDirectoryBlockHeader {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddirectoryBlockDirectoryBlockGetHeader.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	c.Init()
	return c.Header
}

func (c *DirectoryBlock) SetHeader(header interfaces.IDirectoryBlockHeader) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddirectoryBlockDirectoryBlockSetHeader.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	c.Header = header
}

func (c *DirectoryBlock) SetDBEntries(dbEntries []interfaces.IDBEntry) error {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddirectoryBlockDirectoryBlockSetDBEntries.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	c.DBEntries = dbEntries
	c.GetHeader().SetBlockCount(uint32(len(dbEntries)))
	_, err := c.BuildBodyMR()
	if err != nil {
		return err
	}
	return nil
}

func (c *DirectoryBlock) New() interfaces.BinaryMarshallableAndCopyable {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddirectoryBlockDirectoryBlockNew.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	dBlock := new(DirectoryBlock)
	dBlock.Header = NewDBlockHeader()
	dBlock.DBHash = primitives.NewZeroHash()
	dBlock.KeyMR = primitives.NewZeroHash()
	return dBlock
}

func (c *DirectoryBlock) GetDatabaseHeight() uint32 {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddirectoryBlockDirectoryBlockGetDatabaseHeight.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	c.Init()
	return c.GetHeader().GetDBHeight()
}

func (c *DirectoryBlock) GetChainID() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddirectoryBlockDirectoryBlockGetChainID.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.NewHash(constants.D_CHAINID)
}

func (c *DirectoryBlock) DatabasePrimaryIndex() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddirectoryBlockDirectoryBlockDatabasePrimaryIndex.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return c.GetKeyMR()
}

func (c *DirectoryBlock) DatabaseSecondaryIndex() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddirectoryBlockDirectoryBlockDatabaseSecondaryIndex.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return c.GetHash()
}

func (e *DirectoryBlock) JSONByte() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddirectoryBlockDirectoryBlockJSONByte.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSON(e)
}

func (e *DirectoryBlock) JSONString() (string, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddirectoryBlockDirectoryBlockJSONString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSONString(e)
}

func (e *DirectoryBlock) String() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddirectoryBlockDirectoryBlockString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

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
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddirectoryBlockDirectoryBlockMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

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
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddirectoryBlockDirectoryBlockBuildBodyMR.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

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
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddirectoryBlockDirectoryBlockHeaderHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	binaryEBHeader, err := b.GetHeader().MarshalBinary()
	if err != nil {
		return nil, err
	}
	return primitives.Sha(binaryEBHeader), nil
}

func (b *DirectoryBlock) BodyKeyMR() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddirectoryBlockDirectoryBlockBodyKeyMR.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	key, _ := b.BuildBodyMR()
	return key
}

func (b *DirectoryBlock) BuildKeyMerkleRoot() (keyMR interfaces.IHash, err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddirectoryBlockDirectoryBlockBuildKeyMerkleRoot.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

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
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddirectoryBlockUnmarshalDBlock.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

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
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddirectoryBlockDirectoryBlockUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

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
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddirectoryBlockDirectoryBlockGetTimestamp.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return h.GetHeader().GetTimestamp()
}

func (b *DirectoryBlock) UnmarshalBinary(data []byte) (err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddirectoryBlockDirectoryBlockUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	_, err = b.UnmarshalBinaryData(data)
	return
}

func (b *DirectoryBlock) GetHash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddirectoryBlockDirectoryBlockGetHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return b.GetFullHash()
}

func (b *DirectoryBlock) GetFullHash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddirectoryBlockDirectoryBlockGetFullHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	binaryDblock, err := b.MarshalBinary()
	if err != nil {
		return nil
	}
	b.DBHash = primitives.Sha(binaryDblock)
	return b.DBHash
}

func (b *DirectoryBlock) AddEntry(chainID interfaces.IHash, keyMR interfaces.IHash) error {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddirectoryBlockDirectoryBlockAddEntry.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

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
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddirectoryBlockNewDirectoryBlock.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

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
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddirectoryBlockCheckBlockPairIntegrity.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

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
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddirectoryBlockDirectoryBlockMarshalJSON.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	e.GetKeyMR()
	e.GetFullHash()

	return json.Marshal(ExpandedDBlock(e))
}
