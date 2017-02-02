package dbInfo

import (
	"bytes"
	"encoding/gob"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
)

type DirBlockInfo struct {
	// Serial hash for the directory block
	DBHash    interfaces.IHash
	DBHeight  uint32 //directory block height
	Timestamp int64  // time of this dir block info being created
	// BTCTxHash is the Tx hash returned from rpcclient.SendRawTransaction
	BTCTxHash interfaces.IHash // use string or *btcwire.ShaHash ???
	// BTCTxOffset is the index of the TX in this BTC block
	BTCTxOffset int32
	// BTCBlockHeight is the height of the block where this TX is stored in BTC
	BTCBlockHeight int32
	//BTCBlockHash is the hash of the block where this TX is stored in BTC
	BTCBlockHash interfaces.IHash // use string or *btcwire.ShaHash ???
	// DBMerkleRoot is the merkle root of the Directory Block
	// and is written into BTC as OP_RETURN data
	DBMerkleRoot interfaces.IHash
	// A flag to to show BTC anchor confirmation
	BTCConfirmed bool
}

var _ interfaces.Printable = (*DirBlockInfo)(nil)
var _ interfaces.BinaryMarshallableAndCopyable = (*DirBlockInfo)(nil)
var _ interfaces.DatabaseBatchable = (*DirBlockInfo)(nil)
var _ interfaces.IDirBlockInfo = (*DirBlockInfo)(nil)

func NewDirBlockInfo() *DirBlockInfo {
	callTime := time.Now().UnixNano()
	defer dbInfoNewDirBlockInfo.Observe(float64(time.Now().UnixNano() - callTime))	

	dbi := new(DirBlockInfo)
	dbi.DBHash = primitives.NewZeroHash()
	dbi.BTCTxHash = primitives.NewZeroHash()
	dbi.BTCBlockHash = primitives.NewZeroHash()
	dbi.DBMerkleRoot = primitives.NewZeroHash()
	return dbi
}

func (e *DirBlockInfo) String() string {
	callTime := time.Now().UnixNano()
	defer dbInfoString.Observe(float64(time.Now().UnixNano() - callTime))	
	str, _ := e.JSONString()
	return str
}

func (e *DirBlockInfo) JSONByte() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer dbInfoJSONByte.Observe(float64(time.Now().UnixNano() - callTime))	
	return primitives.EncodeJSON(e)
}

func (e *DirBlockInfo) JSONString() (string, error) {
	callTime := time.Now().UnixNano()
	defer dbInfoJSONString.Observe(float64(time.Now().UnixNano() - callTime))	
	return primitives.EncodeJSONString(e)
}

func (e *DirBlockInfo) JSONBuffer(b *bytes.Buffer) error {
	callTime := time.Now().UnixNano()
	defer dbInfoJSONBuffer.Observe(float64(time.Now().UnixNano() - callTime))	
	return primitives.EncodeJSONToBuffer(e, b)
}

func (c *DirBlockInfo) New() interfaces.BinaryMarshallableAndCopyable {
	callTime := time.Now().UnixNano()
	defer dbInfoNew.Observe(float64(time.Now().UnixNano() - callTime))	
	return NewDirBlockInfo()
}

func (c *DirBlockInfo) GetDatabaseHeight() uint32 {
	callTime := time.Now().UnixNano()
	defer dbInfoGetDatabaseHeight.Observe(float64(time.Now().UnixNano() - callTime))	
	return c.DBHeight
}

func (c *DirBlockInfo) GetDBHeight() uint32 {
	callTime := time.Now().UnixNano()
	defer dbInfoGetDBHeight.Observe(float64(time.Now().UnixNano() - callTime))	
	return c.DBHeight
}

func (c *DirBlockInfo) GetBTCConfirmed() bool {
	callTime := time.Now().UnixNano()
	defer dbInfoGetBTCConfirmed.Observe(float64(time.Now().UnixNano() - callTime))	
	return c.BTCConfirmed
}

func (c *DirBlockInfo) GetChainID() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer dbInfoGetChainID.Observe(float64(time.Now().UnixNano() - callTime))	
	id := make([]byte, 32)
	copy(id, []byte("DirBlockInfo"))
	return primitives.NewHash(id)
}

func (c *DirBlockInfo) DatabasePrimaryIndex() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer dbInfoDatabasePrimaryIndex.Observe(float64(time.Now().UnixNano() - callTime))	
	return c.DBMerkleRoot
}

func (c *DirBlockInfo) DatabaseSecondaryIndex() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer dbInfoDatabaseSecondaryIndex.Observe(float64(time.Now().UnixNano() - callTime))	
	return c.DBHash
}

func (e *DirBlockInfo) GetDBMerkleRoot() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer dbInfoGetDBMerkleRoot.Observe(float64(time.Now().UnixNano() - callTime))	
	return e.DBMerkleRoot
}

func (e *DirBlockInfo) GetBTCTxHash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer dbInfoGetBTCTxHash.Observe(float64(time.Now().UnixNano() - callTime))	
	return e.BTCTxHash
}

func (e *DirBlockInfo) GetTimestamp() interfaces.Timestamp {
	callTime := time.Now().UnixNano()
	defer dbInfoGetTimestamp.Observe(float64(time.Now().UnixNano() - callTime))	
	return primitives.NewTimestampFromMilliseconds(uint64(e.Timestamp))
}

func (e *DirBlockInfo) GetBTCBlockHeight() int32 {
	callTime := time.Now().UnixNano()
	defer dbInfoGetBTCBlockHeight.Observe(float64(time.Now().UnixNano() - callTime))	
	return e.BTCBlockHeight
}

func (e *DirBlockInfo) MarshalBinary() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer dbInfoMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))	
	var data primitives.Buffer

	enc := gob.NewEncoder(&data)

	err := enc.Encode(newDirBlockInfoCopyFromDBI(e))
	if err != nil {
		return nil, err
	}
	return data.DeepCopyBytes(), nil
}

func (e *DirBlockInfo) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	callTime := time.Now().UnixNano()
	defer dbInfoUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))	
	dec := gob.NewDecoder(primitives.NewBuffer(data))
	dbic := newDirBlockInfoCopy()
	err = dec.Decode(dbic)
	if err != nil {
		return nil, err
	}
	e.parseDirBlockInfoCopy(dbic)
	return nil, nil
}

func (e *DirBlockInfo) UnmarshalBinary(data []byte) (err error) {
	callTime := time.Now().UnixNano()
	defer dbInfoUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))	
	_, err = e.UnmarshalBinaryData(data)
	return
}

func (e *DirBlockInfo) SetTimestamp(timestamp interfaces.Timestamp) {
	callTime := time.Now().UnixNano()
	defer dbInfoSetTimestamp.Observe(float64(time.Now().UnixNano() - callTime))	
	e.Timestamp = timestamp.GetTimeMilli()
}

type dirBlockInfoCopy struct {
	// Serial hash for the directory block
	DBHash    interfaces.IHash
	DBHeight  uint32 //directory block height
	Timestamp int64  // time of this dir block info being created
	// BTCTxHash is the Tx hash returned from rpcclient.SendRawTransaction
	BTCTxHash interfaces.IHash // use string or *btcwire.ShaHash ???
	// BTCTxOffset is the index of the TX in this BTC block
	BTCTxOffset int32
	// BTCBlockHeight is the height of the block where this TX is stored in BTC
	BTCBlockHeight int32
	//BTCBlockHash is the hash of the block where this TX is stored in BTC
	BTCBlockHash interfaces.IHash // use string or *btcwire.ShaHash ???
	// DBMerkleRoot is the merkle root of the Directory Block
	// and is written into BTC as OP_RETURN data
	DBMerkleRoot interfaces.IHash
	// A flag to to show BTC anchor confirmation
	BTCConfirmed bool
}

func newDirBlockInfoCopyFromDBI(dbi *DirBlockInfo) *dirBlockInfoCopy {
	callTime := time.Now().UnixNano()
	defer dbInfonewDirBlockInfoCopyFromDBI.Observe(float64(time.Now().UnixNano() - callTime))	
	dbic := new(dirBlockInfoCopy)
	dbic.DBHash = dbi.DBHash
	dbic.DBHeight = dbi.DBHeight
	dbic.Timestamp = dbi.Timestamp
	dbic.BTCTxHash = dbi.BTCTxHash
	dbic.BTCTxOffset = dbi.BTCTxOffset
	dbic.BTCBlockHeight = dbi.BTCBlockHeight
	dbic.BTCBlockHash = dbi.BTCBlockHash
	dbic.DBMerkleRoot = dbi.DBMerkleRoot
	dbic.BTCConfirmed = dbi.BTCConfirmed
	return dbic
}

func newDirBlockInfoCopy() *dirBlockInfoCopy {
	callTime := time.Now().UnixNano()
	defer dbInfonewDirBlockInfoCopy.Observe(float64(time.Now().UnixNano() - callTime))	
	dbi := new(dirBlockInfoCopy)
	dbi.DBHash = primitives.NewZeroHash()
	dbi.BTCTxHash = primitives.NewZeroHash()
	dbi.BTCBlockHash = primitives.NewZeroHash()
	dbi.DBMerkleRoot = primitives.NewZeroHash()
	return dbi
}

func (dbic *DirBlockInfo) parseDirBlockInfoCopy(dbi *dirBlockInfoCopy) {
	callTime := time.Now().UnixNano()
	defer dbInfoparseDirBlockInfoCopy.Observe(float64(time.Now().UnixNano() - callTime))	
	dbic.DBHash = dbi.DBHash
	dbic.DBHeight = dbi.DBHeight
	dbic.Timestamp = dbi.Timestamp
	dbic.BTCTxHash = dbi.BTCTxHash
	dbic.BTCTxOffset = dbi.BTCTxOffset
	dbic.BTCBlockHeight = dbi.BTCBlockHeight
	dbic.BTCBlockHash = dbi.BTCBlockHash
	dbic.DBMerkleRoot = dbi.DBMerkleRoot
	dbic.BTCConfirmed = dbi.BTCConfirmed
}

// NewDirBlockInfoFromDirBlock creates a DirDirBlockInfo from DirectoryBlock
func NewDirBlockInfoFromDirBlock(dirBlock interfaces.IDirectoryBlock) *DirBlockInfo {
	callTime := time.Now().UnixNano()
	defer dbInfoNewDirBlockInfoFromDirBlock.Observe(float64(time.Now().UnixNano() - callTime))	
	dbi := new(DirBlockInfo)
	dbi.DBHash = dirBlock.GetHash()
	dbi.DBHeight = dirBlock.GetDatabaseHeight()
	dbi.DBMerkleRoot = dirBlock.GetKeyMR()
	dbi.SetTimestamp(dirBlock.GetHeader().GetTimestamp())
	dbi.BTCTxHash = primitives.NewZeroHash()
	dbi.BTCBlockHash = primitives.NewZeroHash()
	dbi.BTCConfirmed = false
	return dbi
}
