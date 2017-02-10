package adminBlock

import (
	"encoding/binary"
	"fmt"

	"github.com/FactomProject/factomd/common/constants"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
	"time"
)

// DB Signature Entry -------------------------
type AddAuditServer struct {
	IdentityChainID interfaces.IHash
	DBHeight        uint32
}

var _ interfaces.IABEntry = (*AddAuditServer)(nil)
var _ interfaces.BinaryMarshallable = (*AddAuditServer)(nil)

func (e *AddAuditServer) Init() {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdadminBlockAddAuditServerInit.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if e.IdentityChainID == nil {
		e.IdentityChainID = primitives.NewZeroHash()
	}
}

func (e *AddAuditServer) String() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdadminBlockAddAuditServerString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	e.Init()
	var out primitives.Buffer
	out.WriteString(fmt.Sprintf("    E: %20s -- %17s %8x %12s %8d",
		"AddAuditServer",
		"IdentityChainID", e.IdentityChainID.Bytes()[3:5],
		"DBHeight", e.DBHeight))
	return (string)(out.DeepCopyBytes())
}

func (c *AddAuditServer) UpdateState(state interfaces.IState) error {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdadminBlockAddAuditServerUpdateState.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	c.Init()
	state.AddAuditServer(c.DBHeight, c.IdentityChainID)
	authorityDeltaString := fmt.Sprintf("AdminBlock (AddAudMsg DBHt: %d) \n v %s", c.DBHeight, c.IdentityChainID.String()[5:10])
	state.AddStatus(authorityDeltaString)
	state.AddAuthorityDelta(authorityDeltaString)
	state.UpdateAuthorityFromABEntry(c)
	return nil
}

// Create a new DB Signature Entry
func NewAddAuditServer(identityChainID interfaces.IHash, dbheight uint32) (e *AddAuditServer) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdadminBlockNewAddAuditServer.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if identityChainID == nil {
		return nil
	}
	e = new(AddAuditServer)
	e.DBHeight = dbheight
	e.IdentityChainID = primitives.NewHash(identityChainID.Bytes())
	return
}

func (e *AddAuditServer) Type() byte {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdadminBlockAddAuditServerType.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return constants.TYPE_ADD_AUDIT_SERVER
}

func (e *AddAuditServer) MarshalBinary() (data []byte, err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdadminBlockAddAuditServerMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	e.Init()
	var buf primitives.Buffer

	buf.Write([]byte{e.Type()})

	data, err = e.IdentityChainID.MarshalBinary()
	if err != nil {
		return nil, err
	}
	buf.Write(data)

	binary.Write(&buf, binary.BigEndian, e.DBHeight)

	return buf.DeepCopyBytes(), nil
}

func (e *AddAuditServer) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdadminBlockAddAuditServerUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error unmarshalling Add Federated Server Entry: %v", r)
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
		panic(err.Error())
	}

	e.DBHeight, newData = binary.BigEndian.Uint32(newData[0:4]), newData[4:]

	return
}

func (e *AddAuditServer) UnmarshalBinary(data []byte) (err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdadminBlockAddAuditServerUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	_, err = e.UnmarshalBinaryData(data)
	return
}

func (e *AddAuditServer) JSONByte() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdadminBlockAddAuditServerJSONByte.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSON(e)
}

func (e *AddAuditServer) JSONString() (string, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdadminBlockAddAuditServerJSONString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSONString(e)
}

func (e *AddAuditServer) IsInterpretable() bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdadminBlockAddAuditServerIsInterpretable.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return false
}

func (e *AddAuditServer) Interpret() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdadminBlockAddAuditServerInterpret.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return ""
}

func (e *AddAuditServer) Hash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdadminBlockAddAuditServerHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	bin, err := e.MarshalBinary()
	if err != nil {
		panic(err)
	}
	return primitives.Sha(bin)
}
