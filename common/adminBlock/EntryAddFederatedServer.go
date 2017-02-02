package adminBlock

import (
	"encoding/binary"
	"fmt"

	"github.com/FactomProject/factomd/common/constants"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
)

// DB Signature Entry -------------------------
type AddFederatedServer struct {
	IdentityChainID interfaces.IHash
	DBHeight        uint32
}

var _ interfaces.IABEntry = (*AddFederatedServer)(nil)
var _ interfaces.BinaryMarshallable = (*AddFederatedServer)(nil)

func (e *AddFederatedServer) Init() {
	if e.IdentityChainID == nil {
		e.IdentityChainID = primitives.NewZeroHash()
	}
}

func (e *AddFederatedServer) String() string {
	callTime := time.Now().UnixNano()
	defer entryAddFederatedServerString.Observe(float64(time.Now().UnixNano() - callTime))	
	e.Init()
	var out primitives.Buffer
	out.WriteString(fmt.Sprintf("    E: %35s -- %17s %8x %12s %8d",
		"AddFedServer",
		"IdentityChainID", e.IdentityChainID.Bytes()[3:5],
		"DBHeight",
		e.DBHeight))
	return (string)(out.DeepCopyBytes())
}

func (c *AddFederatedServer) UpdateState(state interfaces.IState) error {
	callTime := time.Now().UnixNano()
	defer entryAddFederatedServerUpdateState.Observe(float64(time.Now().UnixNano() - callTime))	
	c.Init()
	state.AddFedServer(c.DBHeight, c.IdentityChainID)
	authorityDeltaString := fmt.Sprintf("AdminBlock (AddFedMsg DBHt: %d) \n ^ %s", c.DBHeight, c.IdentityChainID.String()[5:10])
	state.AddStatus(authorityDeltaString)
	state.AddAuthorityDelta(authorityDeltaString)
	state.UpdateAuthorityFromABEntry(c)
	return nil
}

// Create a new DB Signature Entry
func NewAddFederatedServer(identityChainID interfaces.IHash, dbheight uint32) (e *AddFederatedServer) {
	callTime := time.Now().UnixNano()
	defer entryAddFederatedServerNewAddFederatedServer.Observe(float64(time.Now().UnixNano() - callTime))	
	if identityChainID == nil {
		return nil
	}
	e = new(AddFederatedServer)
	e.DBHeight = dbheight
	e.IdentityChainID = primitives.NewHash(identityChainID.Bytes())
	return
}

func (e *AddFederatedServer) Type() byte {
	callTime := time.Now().UnixNano()
	defer entryAddFederatedServerType.Observe(float64(time.Now().UnixNano() - callTime))	
	return constants.TYPE_ADD_FED_SERVER
}

func (e *AddFederatedServer) MarshalBinary() (data []byte, err error) {
	callTime := time.Now().UnixNano()
	defer entryAddFederatedServerMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))	
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

func (e *AddFederatedServer) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	callTime := time.Now().UnixNano()
	defer entryAddFederatedServerUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))	
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

func (e *AddFederatedServer) UnmarshalBinary(data []byte) (err error) {
	callTime := time.Now().UnixNano()
	defer entryAddFederatedServerUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))	
	_, err = e.UnmarshalBinaryData(data)
	return
}

func (e *AddFederatedServer) JSONByte() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer entryAddFederatedServerJSONByte.Observe(float64(time.Now().UnixNano() - callTime))	
	return primitives.EncodeJSON(e)
}

func (e *AddFederatedServer) JSONString() (string, error) {
	callTime := time.Now().UnixNano()
	defer entryAddFederatedServerJSONString.Observe(float64(time.Now().UnixNano() - callTime))	
	return primitives.EncodeJSONString(e)
}

func (e *AddFederatedServer) IsInterpretable() bool {
	return false
}

func (e *AddFederatedServer) Interpret() string {
	return ""
}

func (e *AddFederatedServer) Hash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer entryAddFederatedServerHash.Observe(float64(time.Now().UnixNano() - callTime))	
	bin, err := e.MarshalBinary()
	if err != nil {
		panic(err)
	}
	return primitives.Sha(bin)
}
