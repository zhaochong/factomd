// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package adminBlock

import (
	"encoding/json"
	"fmt"

	"github.com/FactomProject/factomd/common/constants"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
)

// Administrative Block
// This is a special block which accompanies this Directory Block.
// It contains the signatures and organizational data needed to validate previous and future Directory Blocks.
// This block is included in the DB body. It appears there with a pair of the Admin AdminChainID:SHA256 of the block.
// For more details, please go to:
// https://github.com/FactomProject/FactomDocs/blob/master/factomDataStructureDetails.md#administrative-block
type AdminBlock struct {
	//Marshalized
	Header    interfaces.IABlockHeader
	ABEntries []interfaces.IABEntry //Interface
}

var _ interfaces.IAdminBlock = (*AdminBlock)(nil)
var _ interfaces.Printable = (*AdminBlock)(nil)
var _ interfaces.BinaryMarshallableAndCopyable = (*AdminBlock)(nil)
var _ interfaces.DatabaseBatchable = (*AdminBlock)(nil)

func (c *AdminBlock) Init() {
	if c.Header == nil {
		h := new(ABlockHeader)
		h.Init()
		c.Header = h
	}
}

func (c *AdminBlock) String() string {
	callTime := time.Now().UnixNano()
	defer adminBlockString.Observe(float64(time.Now().UnixNano() - callTime))
	c.Init()
	var out primitives.Buffer

	fh, _ := c.BackReferenceHash()
	if fh == nil {
		fh = primitives.NewZeroHash()
	}
	out.WriteString(fmt.Sprintf("%20s %x\n", "Primary Hash:", c.DatabasePrimaryIndex().Bytes()))
	out.WriteString(fmt.Sprintf("%20s %x\n", "512 Sha3:", fh.Bytes()))

	out.WriteString(c.GetHeader().String())
	out.WriteString("Entries: \n")
	for _, entry := range c.ABEntries {
		out.WriteString(entry.String() + "\n")
	}

	return (string)(out.DeepCopyBytes())
}

func (c *AdminBlock) UpdateState(state interfaces.IState) error {
	callTime := time.Now().UnixNano()
	defer adminBlockUpdateState.Observe(float64(time.Now().UnixNano() - callTime))
	c.Init()
	if state == nil {
		return fmt.Errorf("No State provided")
	}

	dbSigs := []*DBSignatureEntry{}
	for _, entry := range c.ABEntries {
		if entry.Type() == constants.TYPE_DB_SIGNATURE {
			dbSigs = append(dbSigs, entry.(*DBSignatureEntry))
		} else {
			err := entry.UpdateState(state)
			if err != nil {
				return err
			}
		}
	}

	for _, dbSig := range dbSigs {
		//list.State.ProcessLists.Get(currentDBHeight).DBSignatures
		state.AddDBSig(c.GetDBHeight()-1, dbSig.IdentityAdminChainID, &dbSig.PrevDBSig)
	}

	// Clear any keys that are now too old to be valid
	state.UpdateAuthSigningKeys(c.GetHeader().GetDBHeight())
	return nil
}

func (c *AdminBlock) AddDBSig(serverIdentity interfaces.IHash, sig interfaces.IFullSignature) error {
	callTime := time.Now().UnixNano()
	defer adminBlockAddDBSig.Observe(float64(time.Now().UnixNano() - callTime))
	if serverIdentity == nil {
		return fmt.Errorf("No serverIdentity provided")
	}
	if sig == nil {
		return fmt.Errorf("No sig provided")
	}

	entry, err := NewDBSignatureEntry(serverIdentity, sig)
	if err != nil {
		return err
	}
	return c.AddEntry(entry)
}

func (c *AdminBlock) AddFedServer(identityChainID interfaces.IHash) error {
	callTime := time.Now().UnixNano()
	defer adminBlockAddFedServer.Observe(float64(time.Now().UnixNano() - callTime))	
	c.Init()
	if identityChainID == nil {
		return fmt.Errorf("No identityChainID provided")
	}

	entry := NewAddFederatedServer(identityChainID, c.GetHeader().GetDBHeight()+1) // Goes in the NEXT block
	return c.AddEntry(entry)
}

func (c *AdminBlock) AddAuditServer(identityChainID interfaces.IHash) error {
	callTime := time.Now().UnixNano()
	defer adminBlockAddAuditServer.Observe(float64(time.Now().UnixNano() - callTime))	
	c.Init()
	if identityChainID == nil {
		return fmt.Errorf("No identityChainID provided")
	}

	entry := NewAddAuditServer(identityChainID, c.GetHeader().GetDBHeight()+1) // Goes in the NEXT block
	return c.AddEntry(entry)
}

func (c *AdminBlock) RemoveFederatedServer(identityChainID interfaces.IHash) error {
	callTime := time.Now().UnixNano()
	defer adminBlockRemoveFederatedServer.Observe(float64(time.Now().UnixNano() - callTime))	
	c.Init()
	if identityChainID == nil {
		return fmt.Errorf("No identityChainID provided")
	}

	entry := NewRemoveFederatedServer(identityChainID, c.GetHeader().GetDBHeight()+1) // Goes in the NEXT block
	return c.AddEntry(entry)
}

func (c *AdminBlock) AddMatryoshkaHash(identityChainID interfaces.IHash, mHash interfaces.IHash) error {
	callTime := time.Now().UnixNano()
	defer adminBlockAddMatryoshkaHash.Observe(float64(time.Now().UnixNano() - callTime))	
	if identityChainID == nil {
		return fmt.Errorf("No identityChainID provided")
	}
	if mHash == nil {
		return fmt.Errorf("No mHash provided")
	}
	entry := NewAddReplaceMatryoshkaHash(identityChainID, mHash)
	return c.AddEntry(entry)
}

func (c *AdminBlock) AddFederatedServerSigningKey(identityChainID interfaces.IHash, publicKey [32]byte) error {
	callTime := time.Now().UnixNano()
	defer adminBlockAddFederatedServerSigningKey.Observe(float64(time.Now().UnixNano() - callTime))	
	c.Init()
	if identityChainID == nil {
		return fmt.Errorf("No identityChainID provided")
	}

	p := new(primitives.PublicKey)
	err := p.UnmarshalBinary(publicKey[:])
	if err != nil {
		return err
	}
	entry := NewAddFederatedServerSigningKey(identityChainID, byte(0), *p, c.GetHeader().GetDBHeight()+1)
	return c.AddEntry(entry)
}

func (c *AdminBlock) AddFederatedServerBitcoinAnchorKey(identityChainID interfaces.IHash, keyPriority byte, keyType byte, ecdsaPublicKey [20]byte) error {
	callTime := time.Now().UnixNano()
	defer adminBlockAddFederatedServerBitcoinAnchorKey.Observe(float64(time.Now().UnixNano() - callTime))	
	if identityChainID == nil {
		return fmt.Errorf("No identityChainID provided")
	}

	b := new(primitives.ByteSlice20)
	err := b.UnmarshalBinary(ecdsaPublicKey[:])
	if err != nil {
		return err
	} else {
		entry := NewAddFederatedServerBitcoinAnchorKey(identityChainID, keyPriority, keyType, *b)
		return c.AddEntry(entry)
	}
	return nil
}

func (c *AdminBlock) AddEntry(entry interfaces.IABEntry) error {
	callTime := time.Now().UnixNano()
	defer adminBlockAddEntry.Observe(float64(time.Now().UnixNano() - callTime))	
	if entry == nil {
		return fmt.Errorf("No entry provided")
	}

	if entry.Type() == constants.TYPE_SERVER_FAULT {
		//Server Faults needs to be ordered in a specific way
		return c.AddServerFault(entry)
	}

	for i := range c.ABEntries {
		//Server Faults are always the last entry in an AdminBlock
		if c.ABEntries[i].Type() == constants.TYPE_SERVER_FAULT {
			c.ABEntries = append(c.ABEntries[:i], append([]interfaces.IABEntry{entry}, c.ABEntries[i:]...)...)
			return nil
		}
	}
	c.ABEntries = append(c.ABEntries, entry)
	return nil
}

func (c *AdminBlock) AddServerFault(serverFault interfaces.IABEntry) error {
	callTime := time.Now().UnixNano()
	defer adminBlockAddServerFault.Observe(float64(time.Now().UnixNano() - callTime))	
	if serverFault == nil {
		return fmt.Errorf("No serverFault provided")
	}

	sf, ok := serverFault.(*ServerFault)
	if ok == false {
		return fmt.Errorf("Entry is not serverFault")
	}

	for i := range c.ABEntries {
		if c.ABEntries[i].Type() == constants.TYPE_SERVER_FAULT {
			//Server Faults need to follow a deterministic order
			if c.ABEntries[i].(*ServerFault).Compare(sf) > 0 {
				c.ABEntries = append(c.ABEntries[:i], append([]interfaces.IABEntry{sf}, c.ABEntries[i:]...)...)
				return nil
			}
		}
	}
	c.ABEntries = append(c.ABEntries, sf)
	return nil
}

func (c *AdminBlock) GetHeader() interfaces.IABlockHeader {
	callTime := time.Now().UnixNano()
	defer adminBlockGetHeader.Observe(float64(time.Now().UnixNano() - callTime))	
	c.Init()
	return c.Header
}

func (c *AdminBlock) SetHeader(header interfaces.IABlockHeader) {
	callTime := time.Now().UnixNano()
	defer adminBlockSetHeader.Observe(float64(time.Now().UnixNano() - callTime))	
	c.Header = header
}

func (c *AdminBlock) GetABEntries() []interfaces.IABEntry {
	callTime := time.Now().UnixNano()
	defer adminBlockGetABEntries.Observe(float64(time.Now().UnixNano() - callTime))	
	return c.ABEntries
}

func (c *AdminBlock) GetDBHeight() uint32 {
	callTime := time.Now().UnixNano()
	defer adminBlockGetDBHeight.Observe(float64(time.Now().UnixNano() - callTime))	
	return c.GetHeader().GetDBHeight()
}

func (c *AdminBlock) SetABEntries(abentries []interfaces.IABEntry) {
	callTime := time.Now().UnixNano()
	defer adminBlockSetABEntries.Observe(float64(time.Now().UnixNano() - callTime))	
	c.ABEntries = abentries
}

func (c *AdminBlock) New() interfaces.BinaryMarshallableAndCopyable {
	callTime := time.Now().UnixNano()
	defer adminBlockNew.Observe(float64(time.Now().UnixNano() - callTime))	
	return NewAdminBlock(nil)
}

func (c *AdminBlock) GetDatabaseHeight() uint32 {
	callTime := time.Now().UnixNano()
	defer adminBlockGetDatabaseHeight.Observe(float64(time.Now().UnixNano() - callTime))	
	return c.GetHeader().GetDBHeight()
}

func (c *AdminBlock) GetChainID() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer adminBlockGetChainID.Observe(float64(time.Now().UnixNano() - callTime))	
	return c.GetHeader().GetAdminChainID()
}

func (c *AdminBlock) DatabasePrimaryIndex() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer adminBlockDatabasePrimaryIndex.Observe(float64(time.Now().UnixNano() - callTime))	
	key, _ := c.LookupHash()
	return key
}

func (c *AdminBlock) DatabaseSecondaryIndex() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer adminBlockDatabaseSecondaryIndex.Observe(float64(time.Now().UnixNano() - callTime))	
	key, _ := c.BackReferenceHash()
	return key
}

func (c *AdminBlock) GetHash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer adminBlockGetHash.Observe(float64(time.Now().UnixNano() - callTime))	
	h, _ := c.GetKeyMR()
	return h
}

func (c *AdminBlock) GetKeyMR() (interfaces.IHash, error) {
	callTime := time.Now().UnixNano()
	defer adminBlockGetKeyMR.Observe(float64(time.Now().UnixNano() - callTime))	
	return c.BackReferenceHash()
}

// Returns the SHA512Half hash for the admin block
func (b *AdminBlock) BackReferenceHash() (interfaces.IHash, error) {
	callTime := time.Now().UnixNano()
	defer adminBlockBackReferenceHash.Observe(float64(time.Now().UnixNano() - callTime))	
	var binaryAB []byte
	binaryAB, err := b.MarshalBinary()
	if err != nil {
		return nil, err
	}
	return primitives.Sha512Half(binaryAB), nil
}

// Returns the SHA256 hash for the admin block
func (b *AdminBlock) LookupHash() (interfaces.IHash, error) {
	callTime := time.Now().UnixNano()
	defer adminBlockLookupHash.Observe(float64(time.Now().UnixNano() - callTime))	
	var binaryAB []byte
	binaryAB, err := b.MarshalBinary()
	if err != nil {
		return nil, err
	}
	return primitives.Sha(binaryAB), nil
}

// Add an Admin Block entry to the block
func (b *AdminBlock) AddABEntry(e interfaces.IABEntry) error {
	callTime := time.Now().UnixNano()
	defer adminBlockAddABEntry.Observe(float64(time.Now().UnixNano() - callTime))	
	return b.AddEntry(e)
}

// Add an Admin Block entry to the start of the block entries
func (b *AdminBlock) AddFirstABEntry(e interfaces.IABEntry) (err error) {
	callTime := time.Now().UnixNano()
	defer adminBlockAddABEntry.Observe(float64(time.Now().UnixNano() - callTime))	
	b.ABEntries = append(b.ABEntries, nil)
	copy(b.ABEntries[1:], b.ABEntries[:len(b.ABEntries)-1])
	b.ABEntries[0] = e
	return
}

// Write out the AdminBlock to binary.
func (b *AdminBlock) MarshalBinary() ([]byte, error) {
	b.Init()
	// Marshal all the entries into their own thing (need the size)
	callTime := time.Now().UnixNano()
	defer adminBlockMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))	
	var buf2 primitives.Buffer
	for _, v := range b.ABEntries {
		data, err := v.MarshalBinary()
		if err != nil {
			return nil, err
		}
		buf2.Write(data)
	}

	b.GetHeader().SetMessageCount(uint32(len(b.ABEntries)))
	b.GetHeader().SetBodySize(uint32(buf2.Len()))

	data, err := b.GetHeader().MarshalBinary()
	if err != nil {
		return nil, err
	}

	var buf primitives.Buffer
	buf.Write(data)

	// Write the Body out
	buf.Write(buf2.DeepCopyBytes())

	return buf.DeepCopyBytes(), err
}

func UnmarshalABlock(data []byte) (interfaces.IAdminBlock, error) {
	callTime := time.Now().UnixNano()
	defer adminBlockUnmarshalABlock.Observe(float64(time.Now().UnixNano() - callTime))	
	block := NewAdminBlock(nil)
	err := block.UnmarshalBinary(data)
	if err != nil {
		return nil, err
	}

	return block, nil
}

func (b *AdminBlock) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	callTime := time.Now().UnixNano()
	defer adminBlockUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))	
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error unmarshalling Admin Block: %v", r)
		}
	}()
	newData = data
	h := new(ABlockHeader)
	newData, err = h.UnmarshalBinaryData(newData)
	if err != nil {
		return
	}
	b.Header = h

	b.ABEntries = make([]interfaces.IABEntry, int(b.GetHeader().GetMessageCount()))
	for i := uint32(0); i < b.GetHeader().GetMessageCount(); i++ {
		switch newData[0] {
		case constants.TYPE_MINUTE_NUM:
			b.ABEntries[i] = new(EndOfMinuteEntry)
		case constants.TYPE_DB_SIGNATURE:
			b.ABEntries[i] = new(DBSignatureEntry)
		case constants.TYPE_REVEAL_MATRYOSHKA:
			b.ABEntries[i] = new(RevealMatryoshkaHash)
		case constants.TYPE_ADD_MATRYOSHKA:
			b.ABEntries[i] = new(AddReplaceMatryoshkaHash)
		case constants.TYPE_ADD_SERVER_COUNT:
			b.ABEntries[i] = new(IncreaseServerCount)
		case constants.TYPE_ADD_FED_SERVER:
			b.ABEntries[i] = new(AddFederatedServer)
		case constants.TYPE_ADD_AUDIT_SERVER:
			b.ABEntries[i] = new(AddAuditServer)
		case constants.TYPE_REMOVE_FED_SERVER:
			b.ABEntries[i] = new(RemoveFederatedServer)
		case constants.TYPE_ADD_FED_SERVER_KEY:
			b.ABEntries[i] = new(AddFederatedServerSigningKey)
		case constants.TYPE_ADD_BTC_ANCHOR_KEY:
			b.ABEntries[i] = new(AddFederatedServerBitcoinAnchorKey)
		case constants.TYPE_SERVER_FAULT:
			b.ABEntries[i] = new(ServerFault)
		default:
			fmt.Printf("AB UNDEFINED ENTRY %x for block %v\n", newData[0], b.GetHeader().GetDBHeight())
			panic("Undefined Admin Block Entry Type")
		}
		newData, err = b.ABEntries[i].UnmarshalBinaryData(newData)
		if err != nil {
			return
		}
	}
	return
}

// Read in the binary into the Admin block.
func (b *AdminBlock) UnmarshalBinary(data []byte) (err error) {
	callTime := time.Now().UnixNano()
	defer adminBlockUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))	
	_, err = b.UnmarshalBinaryData(data)
	return
}

// Read in the binary into the Admin block.
func (b *AdminBlock) GetDBSignature() interfaces.IABEntry {
	callTime := time.Now().UnixNano()
	defer adminBlockGetDBSignature.Observe(float64(time.Now().UnixNano() - callTime))	
	b.Init()
	for i := uint32(0); i < b.GetHeader().GetMessageCount(); i++ {
		if b.ABEntries[i].Type() == constants.TYPE_DB_SIGNATURE {
			return b.ABEntries[i]
		}
	}

	return nil
}

func (e *AdminBlock) JSONByte() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer adminBlockJSONByte.Observe(float64(time.Now().UnixNano() - callTime))	
	return primitives.EncodeJSON(e)
}

func (e *AdminBlock) JSONString() (string, error) {
	callTime := time.Now().UnixNano()
	defer adminBlockJSONString.Observe(float64(time.Now().UnixNano() - callTime))	
	return primitives.EncodeJSONString(e)
}

type ExpandedABlock AdminBlock

func (e AdminBlock) MarshalJSON() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer adminBlockMarshalJSON.Observe(float64(time.Now().UnixNano() - callTime))	
	backRefHash, err := e.BackReferenceHash()
	if err != nil {
		return nil, err
	}

	lookupHash, err := e.LookupHash()
	if err != nil {
		return nil, err
	}

	return json.Marshal(struct {
		ExpandedABlock
		BackReferenceHash string
		LookupHash        string
	}{
		ExpandedABlock:    ExpandedABlock(e),
		BackReferenceHash: backRefHash.String(),
		LookupHash:        lookupHash.String(),
	})
}

/*********************************************************************
 * Support
 *********************************************************************/

func NewAdminBlock(prev interfaces.IAdminBlock) interfaces.IAdminBlock {
	callTime := time.Now().UnixNano()
	defer adminBlockNewAdminBlock.Observe(float64(time.Now().UnixNano() - callTime))	
	block := new(AdminBlock)
	block.Init()
	if prev != nil {
		block.GetHeader().SetPrevBackRefHash(primitives.NewZeroHash())
		block.GetHeader().SetDBHeight(prev.GetDBHeight() + 1)
	} else {
		block.GetHeader().SetPrevBackRefHash(primitives.NewZeroHash())
	}
	return block
}

func CheckBlockPairIntegrity(block interfaces.IAdminBlock, prev interfaces.IAdminBlock) error {
	callTime := time.Now().UnixNano()
	defer adminBlockCheckBlockPairIntegrity.Observe(float64(time.Now().UnixNano() - callTime))	
	if block == nil {
		return fmt.Errorf("No block specified")
	}

	if prev == nil {
		if block.GetHeader().GetPrevBackRefHash().IsZero() == false {
			return fmt.Errorf("Invalid PrevBackRefHash")
		}
		if block.GetHeader().GetDBHeight() != 0 {
			return fmt.Errorf("Invalid DBHeight")
		}
	} else {
		if block.GetHeader().GetPrevBackRefHash().IsSameAs(prev.GetHash()) == false {
			return fmt.Errorf("Invalid PrevBackRefHash")
		}
		if block.GetHeader().GetDBHeight() != (prev.GetHeader().GetDBHeight() + 1) {
			return fmt.Errorf("Invalid DBHeight")
		}
	}

	return nil
}
