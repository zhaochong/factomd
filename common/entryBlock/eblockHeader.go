package entryBlock

import (
	"encoding/binary"
	"fmt"

	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
)

// EBlockHeader holds relevent metadata about the Entry Block and the data
// nessisary to verify the previous block in the Entry Block Chain.
type EBlockHeader struct {
	ChainID      interfaces.IHash
	BodyMR       interfaces.IHash
	PrevKeyMR    interfaces.IHash
	PrevFullHash interfaces.IHash
	EBSequence   uint32
	DBHeight     uint32
	EntryCount   uint32
}

var _ interfaces.Printable = (*EBlockHeader)(nil)
var _ interfaces.IEntryBlockHeader = (*EBlockHeader)(nil)

func (e *EBlockHeader) Init() {
	if e.ChainID == nil {
		e.ChainID = primitives.NewZeroHash()
	}
	if e.BodyMR == nil {
		e.BodyMR = primitives.NewZeroHash()
	}
	if e.PrevKeyMR == nil {
		e.PrevKeyMR = primitives.NewZeroHash()
	}
	if e.PrevFullHash == nil {
		e.PrevFullHash = primitives.NewZeroHash()
	}
}

// NewEBlockHeader initializes a new empty Entry Block Header.
func NewEBlockHeader() *EBlockHeader {
	callTime := time.Now().UnixNano()
	defer eBlockHeaderNewEBlockHeader.Observe(float64(time.Now().UnixNano() - callTime))	
	e := new(EBlockHeader)
	e.Init()
	return e
}

func (e *EBlockHeader) JSONByte() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer eBlockHeaderJSONByte.Observe(float64(time.Now().UnixNano() - callTime))	
	return primitives.EncodeJSON(e)
}

func (e *EBlockHeader) JSONString() (string, error) {
	callTime := time.Now().UnixNano()
	defer eBlockHeaderJSONString.Observe(float64(time.Now().UnixNano() - callTime))	
	return primitives.EncodeJSONString(e)
}

func (e *EBlockHeader) String() string {
	callTime := time.Now().UnixNano()
	defer eBlockHeaderString.Observe(float64(time.Now().UnixNano() - callTime))	
	e.Init()
	var out primitives.Buffer
	out.WriteString("  Entry Block Header\n")
	out.WriteString(fmt.Sprintf("    %20s: %x\n", "ChainID", e.ChainID.Bytes()[:3]))
	out.WriteString(fmt.Sprintf("    %20s: %x\n", "BodyMR", e.BodyMR.Bytes()[:3]))
	out.WriteString(fmt.Sprintf("    %20s: %x\n", "PrevKeyMR", e.PrevKeyMR.Bytes()[:3]))
	out.WriteString(fmt.Sprintf("    %20s: %x\n", "PrevFullHash", e.PrevFullHash.Bytes()[:3]))
	out.WriteString(fmt.Sprintf("    %20s: %10v\n", "EBSequence", e.EBSequence))
	out.WriteString(fmt.Sprintf("    %20s: %10v\n", "DBHeight", e.DBHeight))
	out.WriteString(fmt.Sprintf("    %20s: %x\n", "EntryCount", e.EntryCount))
	return (string)(out.DeepCopyBytes())
}

func (c *EBlockHeader) GetChainID() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer eBlockHeaderGetChainID.Observe(float64(time.Now().UnixNano() - callTime))	
	return c.ChainID
}

func (c *EBlockHeader) SetChainID(chainID interfaces.IHash) {
	callTime := time.Now().UnixNano()
	defer eBlockHeaderSetChainID.Observe(float64(time.Now().UnixNano() - callTime))	
	c.ChainID = chainID
}

func (c *EBlockHeader) GetBodyMR() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer eBlockHeaderGetBodyMR.Observe(float64(time.Now().UnixNano() - callTime))	
	return c.BodyMR
}

func (c *EBlockHeader) SetBodyMR(bodyMR interfaces.IHash) {
	callTime := time.Now().UnixNano()
	defer eBlockHeaderSetBodyMR.Observe(float64(time.Now().UnixNano() - callTime))	
	c.BodyMR = bodyMR
}

func (c *EBlockHeader) GetPrevKeyMR() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer eBlockHeaderGetPrevKeyMR.Observe(float64(time.Now().UnixNano() - callTime))	
	return c.PrevKeyMR
}

func (c *EBlockHeader) SetPrevKeyMR(prevKeyMR interfaces.IHash) {
	callTime := time.Now().UnixNano()
	defer eBlockHeaderGetPrevKeyMR.Observe(float64(time.Now().UnixNano() - callTime))	
	c.PrevKeyMR = prevKeyMR
}

func (c *EBlockHeader) GetPrevFullHash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer eBlockHeaderGetPrevFullHash.Observe(float64(time.Now().UnixNano() - callTime))	
	return c.PrevFullHash
}

func (c *EBlockHeader) SetPrevFullHash(prevFullHash interfaces.IHash) {
	callTime := time.Now().UnixNano()
	defer eBlockHeaderSetPrevFullHash.Observe(float64(time.Now().UnixNano() - callTime))	
	c.PrevFullHash = prevFullHash
}

func (c *EBlockHeader) GetEBSequence() uint32 {
	callTime := time.Now().UnixNano()
	defer eBlockHeaderGetEBSequence.Observe(float64(time.Now().UnixNano() - callTime))	
	return c.EBSequence
}

func (c *EBlockHeader) SetEBSequence(sequence uint32) {
	callTime := time.Now().UnixNano()
	defer eBlockHeaderSetEBSequence.Observe(float64(time.Now().UnixNano() - callTime))	
	c.EBSequence = sequence
}

func (c *EBlockHeader) GetDBHeight() uint32 {
	callTime := time.Now().UnixNano()
	defer eBlockHeaderGetDBHeight.Observe(float64(time.Now().UnixNano() - callTime))	
	return c.DBHeight
}

func (c *EBlockHeader) SetDBHeight(dbHeight uint32) {
	callTime := time.Now().UnixNano()
	defer eBlockHeaderSetDBHeight.Observe(float64(time.Now().UnixNano() - callTime))	
	c.DBHeight = dbHeight
}

func (c *EBlockHeader) GetEntryCount() uint32 {
	callTime := time.Now().UnixNano()
	defer eBlockHeaderGetEntryCount.Observe(float64(time.Now().UnixNano() - callTime))	
	return c.EntryCount
}

func (c *EBlockHeader) SetEntryCount(entryCount uint32) {
	callTime := time.Now().UnixNano()
	defer eBlockHeaderSetEntryCount.Observe(float64(time.Now().UnixNano() - callTime))	
	c.EntryCount = entryCount
}

// marshalHeaderBinary returns a serialized binary Entry Block Header
func (e *EBlockHeader) MarshalBinary() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer eBlockHeaderMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))	
	e.Init()
	buf := new(primitives.Buffer)

	buf.Write(e.ChainID.Bytes())
	buf.Write(e.BodyMR.Bytes())
	buf.Write(e.PrevKeyMR.Bytes())
	buf.Write(e.PrevFullHash.Bytes())

	if err := binary.Write(buf, binary.BigEndian, e.EBSequence); err != nil {
		return nil, err
	}

	if err := binary.Write(buf, binary.BigEndian, e.DBHeight); err != nil {
		return nil, err
	}

	if err := binary.Write(buf, binary.BigEndian, e.EntryCount); err != nil {
		return nil, err
	}

	return buf.DeepCopyBytes(), nil
}

// unmarshalHeaderBinary builds the Entry Block Header from the serialized binary.
func (e *EBlockHeader) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	callTime := time.Now().UnixNano()
	defer eBlockHeaderUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))	
	e.Init()
	buf := primitives.NewBuffer(data)
	hash := make([]byte, 32)
	newData = data

	if _, err = buf.Read(hash); err != nil {
		return
	} else {
		e.ChainID.SetBytes(hash)
	}

	if _, err = buf.Read(hash); err != nil {
		return
	} else {
		e.BodyMR.SetBytes(hash)
	}

	if _, err = buf.Read(hash); err != nil {
		return
	} else {
		e.PrevKeyMR.SetBytes(hash)
	}

	if _, err = buf.Read(hash); err != nil {
		return
	} else {
		e.PrevFullHash.SetBytes(hash)
	}

	if err = binary.Read(buf, binary.BigEndian, &e.EBSequence); err != nil {
		return
	}

	if err = binary.Read(buf, binary.BigEndian, &e.DBHeight); err != nil {
		return
	}

	if err = binary.Read(buf, binary.BigEndian, &e.EntryCount); err != nil {
		return
	}

	newData = buf.DeepCopyBytes()

	return
}

func (e *EBlockHeader) UnmarshalBinary(data []byte) (err error) {
	callTime := time.Now().UnixNano()
	defer eBlockHeaderUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))	
	_, err = e.UnmarshalBinaryData(data)
	return
}
