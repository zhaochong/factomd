package adminBlock

import (
	"fmt"

	"github.com/FactomProject/factomd/common/constants"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
	"time"
)

type AddReplaceMatryoshkaHash struct {
	IdentityChainID interfaces.IHash
	MHash           interfaces.IHash
}

var _ interfaces.Printable = (*AddReplaceMatryoshkaHash)(nil)
var _ interfaces.BinaryMarshallable = (*AddReplaceMatryoshkaHash)(nil)
var _ interfaces.IABEntry = (*AddReplaceMatryoshkaHash)(nil)

func (e *AddReplaceMatryoshkaHash) Init() {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdadminBlockAddReplaceMatryoshkaHashInit.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if e.IdentityChainID == nil {
		e.IdentityChainID = primitives.NewZeroHash()
	}
	if e.MHash == nil {
		e.MHash = primitives.NewZeroHash()
	}
}

func (e *AddReplaceMatryoshkaHash) String() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdadminBlockAddReplaceMatryoshkaHashString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	e.Init()
	var out primitives.Buffer
	out.WriteString(fmt.Sprintf("    E: %35s -- %17s %8x %12s %8s",
		"AddReplaceMatryoshkaHash",
		"IdentityChainID", e.IdentityChainID.Bytes()[3:5],
		"MHash", e.MHash.String()[:8]))
	return (string)(out.DeepCopyBytes())
}

func (m *AddReplaceMatryoshkaHash) Type() byte {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdadminBlockAddReplaceMatryoshkaHashType.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return constants.TYPE_ADD_MATRYOSHKA
}

func (c *AddReplaceMatryoshkaHash) UpdateState(state interfaces.IState) error {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdadminBlockAddReplaceMatryoshkaHashUpdateState.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	c.Init()
	state.UpdateAuthorityFromABEntry(c)
	return nil
}

func NewAddReplaceMatryoshkaHash(identityChainID interfaces.IHash, mHash interfaces.IHash) *AddReplaceMatryoshkaHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdadminBlockNewAddReplaceMatryoshkaHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	e := new(AddReplaceMatryoshkaHash)
	e.IdentityChainID = identityChainID
	e.MHash = mHash
	return e
}

func (e *AddReplaceMatryoshkaHash) MarshalBinary() (data []byte, err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdadminBlockAddReplaceMatryoshkaHashMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	e.Init()
	var buf primitives.Buffer

	buf.Write([]byte{e.Type()})
	buf.Write(e.IdentityChainID.Bytes())
	buf.Write(e.MHash.Bytes())

	return buf.DeepCopyBytes(), nil
}

func (e *AddReplaceMatryoshkaHash) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdadminBlockAddReplaceMatryoshkaHashUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error unmarshalling Add Replace Matryoshka Hash: %v", r)
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

func (e *AddReplaceMatryoshkaHash) UnmarshalBinary(data []byte) (err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdadminBlockAddReplaceMatryoshkaHashUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	_, err = e.UnmarshalBinaryData(data)
	return
}

func (e *AddReplaceMatryoshkaHash) JSONByte() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdadminBlockAddReplaceMatryoshkaHashJSONByte.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSON(e)
}

func (e *AddReplaceMatryoshkaHash) JSONString() (string, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdadminBlockAddReplaceMatryoshkaHashJSONString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSONString(e)
}

func (e *AddReplaceMatryoshkaHash) IsInterpretable() bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdadminBlockAddReplaceMatryoshkaHashIsInterpretable.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return false
}

func (e *AddReplaceMatryoshkaHash) Interpret() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdadminBlockAddReplaceMatryoshkaHashInterpret.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return ""
}

func (e *AddReplaceMatryoshkaHash) Hash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdadminBlockAddReplaceMatryoshkaHashHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	bin, err := e.MarshalBinary()
	if err != nil {
		panic(err)
	}
	return primitives.Sha(bin)
}
