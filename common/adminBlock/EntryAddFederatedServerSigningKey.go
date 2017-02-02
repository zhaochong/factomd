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
type AddFederatedServerSigningKey struct {
	IdentityChainID interfaces.IHash
	KeyPriority     byte
	PublicKey       primitives.PublicKey
	DBHeight        uint32
}

var _ interfaces.IABEntry = (*AddFederatedServerSigningKey)(nil)
var _ interfaces.BinaryMarshallable = (*AddFederatedServerSigningKey)(nil)

func (c *AddFederatedServerSigningKey) UpdateState(state interfaces.IState) error {
	callTime := time.Now().UnixNano()
	defer entryAddFederatedServerSigningKeyUpdateState.Observe(float64(time.Now().UnixNano() - callTime))	
	state.UpdateAuthorityFromABEntry(c)
	return nil
}

func (e *AddFederatedServerSigningKey) String() string {
	callTime := time.Now().UnixNano()
	defer entryAddFederatedServerSigningKeyString.Observe(float64(time.Now().UnixNano() - callTime))	
	var out primitives.Buffer
	out.WriteString(fmt.Sprintf("    E: %35s -- %17s %8x %12s %8x %12s %8s %12s %d",
		"AddFederatedServerSigningKey",
		"IdentityChainID", e.IdentityChainID.Bytes()[3:5],
		"KeyPriority", e.KeyPriority,
		"PublicKey", e.PublicKey.String()[:8],
		"DBHeight", e.DBHeight))
	return (string)(out.DeepCopyBytes())
}

// Create a new DB Signature Entry
func NewAddFederatedServerSigningKey(identityChainID interfaces.IHash, keyPriority byte, publicKey primitives.PublicKey, height uint32) (e *AddFederatedServerSigningKey) {
	callTime := time.Now().UnixNano()
	defer entryAddFederatedServerSigningKeyNewAddFederatedServerSigningKey.Observe(float64(time.Now().UnixNano() - callTime))	
	e = new(AddFederatedServerSigningKey)
	e.IdentityChainID = identityChainID
	e.KeyPriority = keyPriority
	e.PublicKey = publicKey
	e.DBHeight = height
	return
}

func (e *AddFederatedServerSigningKey) Type() byte {
	callTime := time.Now().UnixNano()
	defer entryAddFederatedServerSigningKeyType.Observe(float64(time.Now().UnixNano() - callTime))	
	return constants.TYPE_ADD_FED_SERVER_KEY
}

func (e *AddFederatedServerSigningKey) MarshalBinary() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer entryAddFederatedServerSigningKeyMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))	
	var buf primitives.Buffer

	buf.Write([]byte{e.Type()})

	data, err := e.IdentityChainID.MarshalBinary()
	if err != nil {
		return nil, err
	}
	buf.Write(data)

	buf.Write([]byte{e.KeyPriority})

	data, err = e.PublicKey.MarshalBinary()
	if err != nil {
		return nil, err
	}
	buf.Write(data)

	binary.Write(&buf, binary.BigEndian, e.DBHeight)

	return buf.DeepCopyBytes(), nil
}

func (e *AddFederatedServerSigningKey) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	callTime := time.Now().UnixNano()
	defer entryAddFederatedServerSigningKeyUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))	
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error unmarshalling Add Federated server Signing Key Entry: %v", r)
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

	e.KeyPriority, newData = newData[0], newData[1:]

	newData, err = e.PublicKey.UnmarshalBinaryData(newData)
	if err != nil {
		return
	}

	e.DBHeight, newData = binary.BigEndian.Uint32(newData[0:4]), newData[4:]

	return
}

func (e *AddFederatedServerSigningKey) UnmarshalBinary(data []byte) (err error) {
	callTime := time.Now().UnixNano()
	defer entryAddFederatedServerSigningKeyUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))	
	_, err = e.UnmarshalBinaryData(data)
	return
}

func (e *AddFederatedServerSigningKey) JSONByte() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer entryAddFederatedServerSigningKeyJSONByte.Observe(float64(time.Now().UnixNano() - callTime))	
	return primitives.EncodeJSON(e)
}

func (e *AddFederatedServerSigningKey) JSONString() (string, error) {
	callTime := time.Now().UnixNano()
	defer entryAddFederatedServerSigningKeyJSONString.Observe(float64(time.Now().UnixNano() - callTime))	
	return primitives.EncodeJSONString(e)
}

func (e *AddFederatedServerSigningKey) JSONBuffer(b *bytes.Buffer) error {
	callTime := time.Now().UnixNano()
	defer entryAddFederatedServerSigningKeyJSONBuffer.Observe(float64(time.Now().UnixNano() - callTime))	
	return primitives.EncodeJSONToBuffer(e, b)
}

func (e *AddFederatedServerSigningKey) IsInterpretable() bool {
	return false
}

func (e *AddFederatedServerSigningKey) Interpret() string {
	return ""
}

func (e *AddFederatedServerSigningKey) Hash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer entryAddFederatedServerSigningKeyHash.Observe(float64(time.Now().UnixNano() - callTime))	
	bin, err := e.MarshalBinary()
	if err != nil {
		panic(err)
	}
	return primitives.Sha(bin)
}
