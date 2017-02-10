package entryBlock

import (
	"encoding/binary"
	"fmt"

	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
	"time"
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
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryBlockEBlockHeaderInit.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

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
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryBlockNewEBlockHeader.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	e := new(EBlockHeader)
	e.Init()
	return e
}

func (e *EBlockHeader) JSONByte() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryBlockEBlockHeaderJSONByte.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSON(e)
}

func (e *EBlockHeader) JSONString() (string, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryBlockEBlockHeaderJSONString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSONString(e)
}

func (e *EBlockHeader) String() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryBlockEBlockHeaderString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

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
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryBlockEBlockHeaderGetChainID.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return c.ChainID
}

func (c *EBlockHeader) SetChainID(chainID interfaces.IHash) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryBlockEBlockHeaderSetChainID.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	c.ChainID = chainID
}

func (c *EBlockHeader) GetBodyMR() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryBlockEBlockHeaderGetBodyMR.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return c.BodyMR
}

func (c *EBlockHeader) SetBodyMR(bodyMR interfaces.IHash) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryBlockEBlockHeaderSetBodyMR.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	c.BodyMR = bodyMR
}

func (c *EBlockHeader) GetPrevKeyMR() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryBlockEBlockHeaderGetPrevKeyMR.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return c.PrevKeyMR
}

func (c *EBlockHeader) SetPrevKeyMR(prevKeyMR interfaces.IHash) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryBlockEBlockHeaderSetPrevKeyMR.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	c.PrevKeyMR = prevKeyMR
}

func (c *EBlockHeader) GetPrevFullHash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryBlockEBlockHeaderGetPrevFullHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return c.PrevFullHash
}

func (c *EBlockHeader) SetPrevFullHash(prevFullHash interfaces.IHash) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryBlockEBlockHeaderSetPrevFullHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	c.PrevFullHash = prevFullHash
}

func (c *EBlockHeader) GetEBSequence() uint32 {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryBlockEBlockHeaderGetEBSequence.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return c.EBSequence
}

func (c *EBlockHeader) SetEBSequence(sequence uint32) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryBlockEBlockHeaderSetEBSequence.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	c.EBSequence = sequence
}

func (c *EBlockHeader) GetDBHeight() uint32 {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryBlockEBlockHeaderGetDBHeight.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return c.DBHeight
}

func (c *EBlockHeader) SetDBHeight(dbHeight uint32) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryBlockEBlockHeaderSetDBHeight.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	c.DBHeight = dbHeight
}

func (c *EBlockHeader) GetEntryCount() uint32 {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryBlockEBlockHeaderGetEntryCount.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return c.EntryCount
}

func (c *EBlockHeader) SetEntryCount(entryCount uint32) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryBlockEBlockHeaderSetEntryCount.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	c.EntryCount = entryCount
}

// marshalHeaderBinary returns a serialized binary Entry Block Header
func (e *EBlockHeader) MarshalBinary() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryBlockEBlockHeaderMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

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
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryBlockEBlockHeaderUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

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
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryBlockEBlockHeaderUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	_, err = e.UnmarshalBinaryData(data)
	return
}
