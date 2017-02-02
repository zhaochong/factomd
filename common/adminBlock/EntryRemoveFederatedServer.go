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
type RemoveFederatedServer struct {
	IdentityChainID interfaces.IHash
	DBHeight        uint32
}

var _ interfaces.IABEntry = (*RemoveFederatedServer)(nil)
var _ interfaces.BinaryMarshallable = (*RemoveFederatedServer)(nil)

func (e *RemoveFederatedServer) String() string {
	callTime := time.Now().UnixNano()
	defer entryRemoveFederatedServerString.Observe(float64(time.Now().UnixNano() - callTime))	
	var out primitives.Buffer
	out.WriteString(fmt.Sprintf("    E: %35s -- %17s %8x %12s %8d",
		"Remove Federated Server",
		"IdentityChainID",
		e.IdentityChainID.Bytes()[3:5],
		"DBHeight",
		e.DBHeight))
	return (string)(out.DeepCopyBytes())
}

func (c *RemoveFederatedServer) UpdateState(state interfaces.IState) error {
	callTime := time.Now().UnixNano()
	defer entryRemoveFederatedServerUpdateState.Observe(float64(time.Now().UnixNano() - callTime))	
	if len(state.GetFedServers(c.DBHeight)) != 0 {
		state.RemoveFedServer(c.DBHeight, c.IdentityChainID)
	}
	if state.GetOut() {
		state.Println(fmt.Sprintf("Removed Federated Server: %x", c.IdentityChainID.Bytes()[:4]))
	}
	authorityDeltaString := fmt.Sprintf("AdminBlock (RemoveFedMsg DBHt: %d) \n v %s", c.DBHeight, c.IdentityChainID.String()[5:10])
	state.AddStatus(authorityDeltaString)
	state.AddAuthorityDelta(authorityDeltaString)
	state.UpdateAuthorityFromABEntry(c)
	return nil
}

// Create a new DB Signature Entry
func NewRemoveFederatedServer(identityChainID interfaces.IHash, dbheight uint32) (e *RemoveFederatedServer) {
	callTime := time.Now().UnixNano()
	defer entryRemoveFederatedServerNewRemoveFederatedServer.Observe(float64(time.Now().UnixNano() - callTime))	
	e = new(RemoveFederatedServer)
	e.IdentityChainID = primitives.NewHash(identityChainID.Bytes())
	e.DBHeight = dbheight
	return
}

func (e *RemoveFederatedServer) Type() byte {
	callTime := time.Now().UnixNano()
	defer entryRemoveFederatedServerType.Observe(float64(time.Now().UnixNano() - callTime))	
	return constants.TYPE_REMOVE_FED_SERVER
}

func (e *RemoveFederatedServer) MarshalBinary() (data []byte, err error) {
	callTime := time.Now().UnixNano()
	defer entryRemoveFederatedServerMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))	
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

func (e *RemoveFederatedServer) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	callTime := time.Now().UnixNano()
	defer entryRemoveFederatedServerUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))	
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error unmarshalling Remove Federated Server: %v", r)
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

	e.DBHeight, newData = binary.BigEndian.Uint32(newData[0:4]), newData[4:]

	return
}

func (e *RemoveFederatedServer) UnmarshalBinary(data []byte) (err error) {
	callTime := time.Now().UnixNano()
	defer entryRemoveFederatedServerUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))	
	_, err = e.UnmarshalBinaryData(data)
	return
}

func (e *RemoveFederatedServer) JSONByte() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer entryRemoveFederatedServerJSONByte.Observe(float64(time.Now().UnixNano() - callTime))	
	return primitives.EncodeJSON(e)
}

func (e *RemoveFederatedServer) JSONString() (string, error) {
	callTime := time.Now().UnixNano()
	defer entryRemoveFederatedServerJSONString.Observe(float64(time.Now().UnixNano() - callTime))	
	return primitives.EncodeJSONString(e)
}

func (e *RemoveFederatedServer) JSONBuffer(b *bytes.Buffer) error {
	callTime := time.Now().UnixNano()
	defer entryRemoveFederatedServerJSONBuffer.Observe(float64(time.Now().UnixNano() - callTime))	
	return primitives.EncodeJSONToBuffer(e, b)
}

func (e *RemoveFederatedServer) IsInterpretable() bool {
	return false
}

func (e *RemoveFederatedServer) Interpret() string {
	return ""
}

func (e *RemoveFederatedServer) Hash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer entryRemoveFederatedServerHash.Observe(float64(time.Now().UnixNano() - callTime))	
	bin, err := e.MarshalBinary()
	if err != nil {
		panic(err)
	}
	return primitives.Sha(bin)
}
