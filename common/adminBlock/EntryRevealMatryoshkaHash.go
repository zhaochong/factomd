package adminBlock

import (
	"fmt"

	"github.com/FactomProject/factomd/common/constants"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
)

type RevealMatryoshkaHash struct {
	IdentityChainID interfaces.IHash
	MHash           interfaces.IHash
}

var _ interfaces.Printable = (*RevealMatryoshkaHash)(nil)
var _ interfaces.BinaryMarshallable = (*RevealMatryoshkaHash)(nil)
var _ interfaces.IABEntry = (*RevealMatryoshkaHash)(nil)

func (e *RevealMatryoshkaHash) Init() {
	if e.IdentityChainID == nil {
		e.IdentityChainID = primitives.NewZeroHash()
	}
	if e.MHash == nil {
		e.MHash = primitives.NewZeroHash()
	}
}

func (m *RevealMatryoshkaHash) Type() byte {
	callTime := time.Now().UnixNano()
	defer entryRevealMatryoshkaHashType.Observe(float64(time.Now().UnixNano() - callTime))	
	return constants.TYPE_REVEAL_MATRYOSHKA
}

func NewRevealMatryoshkaHash(identityChainID interfaces.IHash, mHash interfaces.IHash) *RevealMatryoshkaHash {
	callTime := time.Now().UnixNano()
	defer entryRevealMatryoshkaHashNewRevealMatryoshkaHash.Observe(float64(time.Now().UnixNano() - callTime))	
	e := new(RevealMatryoshkaHash)
	e.IdentityChainID = identityChainID
	e.MHash = mHash
	return e
}

func (c *RevealMatryoshkaHash) UpdateState(state interfaces.IState) error {
	c.Init()
	return nil
}

func (e *RevealMatryoshkaHash) MarshalBinary() (data []byte, err error) {
	callTime := time.Now().UnixNano()
	defer entryRevealMatryoshkaHashMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))	
	e.Init()
	var buf primitives.Buffer

	buf.Write([]byte{e.Type()})
	buf.Write(e.IdentityChainID.Bytes())
	buf.Write(e.MHash.Bytes())

	return buf.DeepCopyBytes(), nil
}

func (e *RevealMatryoshkaHash) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	callTime := time.Now().UnixNano()
	defer entryRevealMatryoshkaHashUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))	
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error unmarshalling Reveal Matryoshka Hash: %v", r)
		}
	}()
	newData = data
	if newData[0] != e.Type() {
		return nil, fmt.Errorf("Invalid Entry type")
	}
	newData = newData[1:]

	e.IdentityChainID = new(primitives.Hash)
	newData, err = e.IdentityChainID.UnmarshalBinaryData(newData)
	if err != nil {
		return
	}
	e.MHash = new(primitives.Hash)
	newData, err = e.MHash.UnmarshalBinaryData(newData)
	if err != nil {
		return
	}

	return
}

func (e *RevealMatryoshkaHash) UnmarshalBinary(data []byte) (err error) {
	callTime := time.Now().UnixNano()
	defer entryRevealMatryoshkaHashUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))	
	_, err = e.UnmarshalBinaryData(data)
	return
}

func (e *RevealMatryoshkaHash) JSONByte() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer entryRevealMatryoshkaHashJSONByte.Observe(float64(time.Now().UnixNano() - callTime))	
	return primitives.EncodeJSON(e)
}

func (e *RevealMatryoshkaHash) JSONString() (string, error) {
	callTime := time.Now().UnixNano()
	defer entryRevealMatryoshkaHashJSONString.Observe(float64(time.Now().UnixNano() - callTime))	
	return primitives.EncodeJSONString(e)
}

func (e *RevealMatryoshkaHash) String() string {
	callTime := time.Now().UnixNano()
	defer entryRevealMatryoshkaHashString.Observe(float64(time.Now().UnixNano() - callTime))	
	e.Init()
	str := fmt.Sprintf("    E: %35s -- %17s %8x %12s %x",
		"RevealMatryoshkaHash",
		"IdentityChainID", e.IdentityChainID.Bytes()[3:5],
		"Hash", e.MHash.Bytes()[:5])
	return str
}

func (e *RevealMatryoshkaHash) IsInterpretable() bool {
	return false
}

func (e *RevealMatryoshkaHash) Interpret() string {
	return ""
}

func (e *RevealMatryoshkaHash) Hash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer entryRevealMatryoshkaHashHash.Observe(float64(time.Now().UnixNano() - callTime))	
	bin, err := e.MarshalBinary()
	if err != nil {
		panic(err)
	}
	return primitives.Sha(bin)
}
