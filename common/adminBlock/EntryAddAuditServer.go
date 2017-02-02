package adminBlock

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/FactomProject/factomd/common/constants"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
)

// DB Signature Entry -------------------------
type AddAuditServer struct {
	IdentityChainID interfaces.IHash
	DBHeight        uint32
}

var _ interfaces.IABEntry = (*AddAuditServer)(nil)
var _ interfaces.BinaryMarshallable = (*AddAuditServer)(nil)

func (e *AddAuditServer) String() string {
	callTime := time.Now().UnixNano()
	defer entryAddAuditServerString.Observe(float64(time.Now().UnixNano() - callTime))	
	var out primitives.Buffer
	out.WriteString(fmt.Sprintf("    E: %20s -- %17s %8x %12s %8d",
		"AddAuditServer",
		"IdentityChainID", e.IdentityChainID.Bytes()[3:5],
		"DBHeight", e.DBHeight))
	return (string)(out.DeepCopyBytes())
}

func (c *AddAuditServer) UpdateState(state interfaces.IState) error {
	callTime := time.Now().UnixNano()
	defer entryAddAuditServerUpdateState.Observe(float64(time.Now().UnixNano() - callTime))	
	state.AddAuditServer(c.DBHeight, c.IdentityChainID)
	authorityDeltaString := fmt.Sprintf("AdminBlock (AddAudMsg DBHt: %d) \n v %s", c.DBHeight, c.IdentityChainID.String()[5:10])
	state.AddStatus(authorityDeltaString)
	state.AddAuthorityDelta(authorityDeltaString)
	state.UpdateAuthorityFromABEntry(c)
	return nil
}

// Create a new DB Signature Entry
func NewAddAuditServer(identityChainID interfaces.IHash, dbheight uint32) (e *AddAuditServer) {
	callTime := time.Now().UnixNano()
	defer entryAddAuditServerNewAddAuditServer.Observe(float64(time.Now().UnixNano() - callTime))	
	e = new(AddAuditServer)
	e.DBHeight = dbheight
	e.IdentityChainID = primitives.NewHash(identityChainID.Bytes())
	return
}

func (e *AddAuditServer) Type() byte {
	callTime := time.Now().UnixNano()
	defer entryAddAuditServerType.Observe(float64(time.Now().UnixNano() - callTime))	
	return constants.TYPE_ADD_AUDIT_SERVER
}

func (e *AddAuditServer) MarshalBinary() (data []byte, err error) {
	callTime := time.Now().UnixNano()
	defer entryAddAuditServerMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))	
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
	callTime := time.Now().UnixNano()
	defer entryAddAuditServerUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))	
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error unmarshalling Add Federated Server Entry: %v", r)
		}
	}()

	newData = data
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
	callTime := time.Now().UnixNano()
	defer entryAddAuditServerUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))	
	_, err = e.UnmarshalBinaryData(data)
	return
}

func (e *AddAuditServer) JSONByte() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer entryAddAuditServerJSONByte.Observe(float64(time.Now().UnixNano() - callTime))	
	return primitives.EncodeJSON(e)
}

func (e *AddAuditServer) JSONString() (string, error) {
	callTime := time.Now().UnixNano()
	defer entryAddAuditServerJSONByte.Observe(float64(time.Now().UnixNano() - callTime))	
	return primitives.EncodeJSONString(e)
}

func (e *AddAuditServer) JSONBuffer(b *bytes.Buffer) error {
	callTime := time.Now().UnixNano()
	defer entryAddAuditServerJSONBuffer.Observe(float64(time.Now().UnixNano() - callTime))	
	return primitives.EncodeJSONToBuffer(e, b)
}

func (e *AddAuditServer) IsInterpretable() bool {
	return false
}

func (e *AddAuditServer) Interpret() string {
	return ""
}

func (e *AddAuditServer) Hash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer entryAddAuditServerHash.Observe(float64(time.Now().UnixNano() - callTime))	
	bin, err := e.MarshalBinary()
	if err != nil {
		panic(err)
	}
	return primitives.Sha(bin)
}
