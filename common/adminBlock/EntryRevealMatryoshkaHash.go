package adminBlock

import (
	"fmt"

	"github.com/FactomProject/factomd/common/constants"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
	"time"
)

type RevealMatryoshkaHash struct {
	IdentityChainID interfaces.IHash
	MHash           interfaces.IHash
}

var _ interfaces.Printable = (*RevealMatryoshkaHash)(nil)
var _ interfaces.BinaryMarshallable = (*RevealMatryoshkaHash)(nil)
var _ interfaces.IABEntry = (*RevealMatryoshkaHash)(nil)

func (e *RevealMatryoshkaHash) Init() {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdadminBlockRevealMatryoshkaHashInit.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if e.IdentityChainID == nil {
		e.IdentityChainID = primitives.NewZeroHash()
	}
	if e.MHash == nil {
		e.MHash = primitives.NewZeroHash()
	}
}

func (m *RevealMatryoshkaHash) Type() byte {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdadminBlockRevealMatryoshkaHashType.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return constants.TYPE_REVEAL_MATRYOSHKA
}

func NewRevealMatryoshkaHash(identityChainID interfaces.IHash, mHash interfaces.IHash) *RevealMatryoshkaHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdadminBlockNewRevealMatryoshkaHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	e := new(RevealMatryoshkaHash)
	e.IdentityChainID = identityChainID
	e.MHash = mHash
	return e
}

func (c *RevealMatryoshkaHash) UpdateState(state interfaces.IState) error {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdadminBlockRevealMatryoshkaHashUpdateState.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	c.Init()
	return nil
}

func (e *RevealMatryoshkaHash) MarshalBinary() (data []byte, err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdadminBlockRevealMatryoshkaHashMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	e.Init()
	var buf primitives.Buffer

	buf.Write([]byte{e.Type()})
	buf.Write(e.IdentityChainID.Bytes())
	buf.Write(e.MHash.Bytes())

	return buf.DeepCopyBytes(), nil
}

func (e *RevealMatryoshkaHash) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdadminBlockRevealMatryoshkaHashUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

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
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdadminBlockRevealMatryoshkaHashUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	_, err = e.UnmarshalBinaryData(data)
	return
}

func (e *RevealMatryoshkaHash) JSONByte() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdadminBlockRevealMatryoshkaHashJSONByte.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSON(e)
}

func (e *RevealMatryoshkaHash) JSONString() (string, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdadminBlockRevealMatryoshkaHashJSONString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSONString(e)
}

func (e *RevealMatryoshkaHash) String() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdadminBlockRevealMatryoshkaHashString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	e.Init()
	str := fmt.Sprintf("    E: %35s -- %17s %8x %12s %x",
		"RevealMatryoshkaHash",
		"IdentityChainID", e.IdentityChainID.Bytes()[3:5],
		"Hash", e.MHash.Bytes()[:5])
	return str
}

func (e *RevealMatryoshkaHash) IsInterpretable() bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdadminBlockRevealMatryoshkaHashIsInterpretable.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return false
}

func (e *RevealMatryoshkaHash) Interpret() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdadminBlockRevealMatryoshkaHashInterpret.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return ""
}

func (e *RevealMatryoshkaHash) Hash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdadminBlockRevealMatryoshkaHashHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	bin, err := e.MarshalBinary()
	if err != nil {
		panic(err)
	}
	return primitives.Sha(bin)
}
