// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package messages

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"time"

	"github.com/FactomProject/factomd/common/constants"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
)

// Communicate a Admin Block Change

type ChangeServerKeyMsg struct {
	MessageBase
	Timestamp        interfaces.Timestamp // Message Timestamp
	IdentityChainID  interfaces.IHash     // ChainID of new server
	AdminBlockChange byte
	KeyType          byte
	KeyPriority      byte
	Key              interfaces.IHash

	Signature interfaces.IFullSignature
}

var _ interfaces.IMsg = (*ChangeServerKeyMsg)(nil)
var _ Signable = (*ChangeServerKeyMsg)(nil)

func (m *ChangeServerKeyMsg) GetRepeatHash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer messagesChangeServerKeyMsgGetRepeatHash.Observe(float64(time.Now().UnixNano() - callTime))
	return m.GetMsgHash()
}

func (m *ChangeServerKeyMsg) GetHash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer messagesChangeServerKeyMsgGetHash.Observe(float64(time.Now().UnixNano() - callTime))
	return m.GetMsgHash()
}

func (m *ChangeServerKeyMsg) GetMsgHash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer messagesChangeServerKeyMsgGetMsgHash.Observe(float64(time.Now().UnixNano() - callTime))
	if m.MsgHash == nil {
		data, err := m.MarshalForSignature()
		if err != nil {
			return nil
		}
		m.MsgHash = primitives.Sha(data)
	}
	return m.MsgHash
}

func (m *ChangeServerKeyMsg) Type() byte {
	callTime := time.Now().UnixNano()
	defer messagesChangeServerKeyMsgType.Observe(float64(time.Now().UnixNano() - callTime))
	return constants.CHANGESERVER_KEY_MSG
}

func (m *ChangeServerKeyMsg) Int() int {
	return -1
}

func (m *ChangeServerKeyMsg) Bytes() []byte {
	return nil
}

func (m *ChangeServerKeyMsg) GetTimestamp() interfaces.Timestamp {
	callTime := time.Now().UnixNano()
	defer messagesChangeServerKeyMsgGetTimestamp.Observe(float64(time.Now().UnixNano() - callTime))
	return m.Timestamp
}

func (m *ChangeServerKeyMsg) Validate(state interfaces.IState) int {
	callTime := time.Now().UnixNano()
	defer messagesChangeServerKeyMsgValidate.Observe(float64(time.Now().UnixNano() - callTime))
	// Check to see if identity exists and is audit or fed server
	if !state.VerifyIsAuthority(m.IdentityChainID) {
		fmt.Println("ChangeServerKey Error. Server is not an authority")
		return -1
	}

	// Should only be 20 bytes in the hash if btc key add
	if m.AdminBlockChange == constants.TYPE_ADD_BTC_ANCHOR_KEY {
		for _, b := range m.Key.Bytes()[21:] {
			if b != 0 {
				fmt.Println("ChangeServerKey Error. Newkey is invalid length")
				return -1
			}
		}
	}

	// Check signatures
	bytes, err := m.MarshalForSignature()
	if err != nil || m.Signature == nil {
		return -1
	}
	sig := m.Signature.GetSignature()
	authSigned, err := state.VerifyAuthoritySignature(bytes, sig, state.GetLeaderHeight())
	if err != nil || authSigned != 1 { // authSigned = 1 for fed signed
		return -1
	}

	isVer, err := m.VerifySignature()
	if err != nil || !isVer {
		// if there is an error during signature verification
		// or if the signature is invalid
		// the message is considered invalid
		return -1
	}

	return 1
}

// Returns true if this is a message for this server to execute as
// a leader.
func (m *ChangeServerKeyMsg) ComputeVMIndex(state interfaces.IState) {
	callTime := time.Now().UnixNano()
	defer messagesChangeServerKeyMsgComputeVMIndex.Observe(float64(time.Now().UnixNano() - callTime))
	m.VMIndex = state.ComputeVMIndex(constants.ADMIN_CHAINID)
}

// Execute the leader functions of the given message
func (m *ChangeServerKeyMsg) LeaderExecute(state interfaces.IState) {
	callTime := time.Now().UnixNano()
	defer messagesChangeServerKeyMsgLeaderExecute.Observe(float64(time.Now().UnixNano() - callTime))
	state.LeaderExecute(m)
}

func (m *ChangeServerKeyMsg) FollowerExecute(state interfaces.IState) {
	callTime := time.Now().UnixNano()
	defer messagesChangeServerKeyMsgFollowerExecute.Observe(float64(time.Now().UnixNano() - callTime))
	state.FollowerExecuteMsg(m)
}

// Acknowledgements do not go into the process list.
func (e *ChangeServerKeyMsg) Process(dbheight uint32, state interfaces.IState) bool {
	callTime := time.Now().UnixNano()
	defer messagesChangeServerKeyMsgProcess.Observe(float64(time.Now().UnixNano() - callTime))
	return state.ProcessChangeServerKey(dbheight, e)
}

func (e *ChangeServerKeyMsg) JSONByte() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer messagesChangeServerKeyMsgJSONByte.Observe(float64(time.Now().UnixNano() - callTime))
	return primitives.EncodeJSON(e)
}

func (e *ChangeServerKeyMsg) JSONString() (string, error) {
	callTime := time.Now().UnixNano()
	defer messagesChangeServerKeyMsgJSONString.Observe(float64(time.Now().UnixNano() - callTime))
	return primitives.EncodeJSONString(e)
}

func (e *ChangeServerKeyMsg) JSONBuffer(b *bytes.Buffer) error {
	callTime := time.Now().UnixNano()
	defer messagesChangeServerKeyMsgJSONBuffer.Observe(float64(time.Now().UnixNano() - callTime))
	return primitives.EncodeJSONToBuffer(e, b)
}

func (m *ChangeServerKeyMsg) Sign(key interfaces.Signer) error {
	callTime := time.Now().UnixNano()
	defer messagesChangeServerKeyMsgSign.Observe(float64(time.Now().UnixNano() - callTime))
	signature, err := SignSignable(m, key)
	if err != nil {
		return err
	}
	m.Signature = signature
	return nil
}

func (m *ChangeServerKeyMsg) GetSignature() interfaces.IFullSignature {
	callTime := time.Now().UnixNano()
	defer messagesChangeServerKeyMsg.Observe(float64(time.Now().UnixNano() - callTime))
	return m.Signature
}

func (m *ChangeServerKeyMsg) VerifySignature() (bool, error) {
	callTime := time.Now().UnixNano()
	defer messagesChangeServerKeyMsgVerifySignature.Observe(float64(time.Now().UnixNano() - callTime))
	return VerifyMessage(m)
}

func (m *ChangeServerKeyMsg) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	callTime := time.Now().UnixNano()
	defer messagesChangeServerKeyMsgUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))
	defer func() {
		return
		if r := recover(); r != nil {
			err = fmt.Errorf("Error unmarshalling Add Server Message: %v", r)
		}
	}()
	newData = data
	if newData[0] != m.Type() {
		return nil, fmt.Errorf("Invalid Message type")
	}
	newData = newData[1:]

	m.Timestamp = new(primitives.Timestamp)
	newData, err = m.Timestamp.UnmarshalBinaryData(newData)
	if err != nil {
		return nil, err
	}

	m.IdentityChainID = new(primitives.Hash)
	newData, err = m.IdentityChainID.UnmarshalBinaryData(newData)
	if err != nil {
		return nil, err
	}

	m.AdminBlockChange = newData[0]
	newData = newData[1:]

	m.KeyType = newData[0]
	newData = newData[1:]

	m.KeyPriority = newData[0]
	newData = newData[1:]

	m.Key = new(primitives.Hash)
	newData, err = m.Key.UnmarshalBinaryData(newData)
	if err != nil {
		return nil, err
	}

	if len(newData) > 32 {
		m.Signature = new(primitives.Signature)
		newData, err = m.Signature.UnmarshalBinaryData(newData)
		if err != nil {
			return nil, err
		}
	}
	return
}

func (m *ChangeServerKeyMsg) UnmarshalBinary(data []byte) error {
	callTime := time.Now().UnixNano()
	defer messagesChangeServerKeyMsgUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	_, err := m.UnmarshalBinaryData(data)
	return err
}

func (m *ChangeServerKeyMsg) MarshalForSignature() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer messagesChangeServerKeyMsgMarshalForSignature.Observe(float64(time.Now().UnixNano() - callTime))
	var buf primitives.Buffer

	binary.Write(&buf, binary.BigEndian, m.Type())

	t := m.GetTimestamp()
	data, err := t.MarshalBinary()
	if err != nil {
		return nil, err
	}
	buf.Write(data)

	data, err = m.IdentityChainID.MarshalBinary()
	if err != nil {
		return nil, err
	}
	buf.Write(data)

	binary.Write(&buf, binary.BigEndian, uint8(m.AdminBlockChange))
	binary.Write(&buf, binary.BigEndian, uint8(m.KeyType))
	binary.Write(&buf, binary.BigEndian, uint8(m.KeyPriority))

	data, err = m.Key.MarshalBinary()
	if err != nil {
		return nil, err
	}
	buf.Write(data)

	return buf.DeepCopyBytes(), nil
}

func (m *ChangeServerKeyMsg) MarshalBinary() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer messagesChangeServerKeyMsgMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	var buf primitives.Buffer

	data, err := m.MarshalForSignature()
	if err != nil {
		return nil, err
	}
	buf.Write(data)

	if m.Signature != nil {
		data, err = m.Signature.MarshalBinary()
		if err != nil {
			MarshalBinary
			return nil, err
		}
		buf.Write(data)
	}

	return buf.DeepCopyBytes(), nil
}

func (m *ChangeServerKeyMsg) String() string {
	callTime := time.Now().UnixNano()
	defer messagesChangeServerKeyMsgString.Observe(float64(time.Now().UnixNano() - callTime))
	var mtype string
	if m.AdminBlockChange == constants.TYPE_ADD_MATRYOSHKA {
		mtype = "MHash"
	} else if m.AdminBlockChange == constants.TYPE_ADD_FED_SERVER_KEY {
		mtype = "Signing Key"
	} else if m.AdminBlockChange == constants.TYPE_ADD_BTC_ANCHOR_KEY {
		mtype = "BTC Key"
	} else {
		mtype = "other"
	}
	return fmt.Sprintf("ChangeServerKey (%s): ChainID: %x Time: %x  Key: %x Msg Hash %x ",
		mtype,
		m.IdentityChainID.Bytes()[:3],
		&m.Timestamp,
		m.Key.Bytes()[:3],
		m.GetMsgHash().Bytes()[:3])

}

func (m *ChangeServerKeyMsg) IsSameAs(b *ChangeServerKeyMsg) bool {
	callTime := time.Now().UnixNano()
	defer messagesChangeServerKeyMsgIsSameAs.Observe(float64(time.Now().UnixNano() - callTime))
	if b == nil {
		return false
	}
	if m.Timestamp.GetTimeMilli() != b.Timestamp.GetTimeMilli() {
		return false
	}
	if !m.IdentityChainID.IsSameAs(b.IdentityChainID) {
		return false
	}
	if m.AdminBlockChange != b.AdminBlockChange {
		return false
	}
	if m.KeyType != b.KeyType {
		return false
	}
	if m.KeyPriority != b.KeyPriority {
		return false
	}
	if !m.Key.IsSameAs(b.Key) {
		return false
	}
	if m.Signature == nil && b.Signature != nil {
		return false
	}
	if m.Signature != nil {
		if m.Signature.IsSameAs(b.Signature) == false {
			return false
		}
	}
	return true
}

func NewChangeServerKeyMsg(state interfaces.IState, identityChain interfaces.IHash, adminChange byte, keyPriority byte, keyType byte, key interfaces.IHash) interfaces.IMsg {
	callTime := time.Now().UnixNano()
	defer messagesChangeServerKeyMsg.Observe(float64(time.Now().UnixNano() - callTime))
	msg := new(ChangeServerKeyMsg)
	msg.IdentityChainID = identityChain
	msg.AdminBlockChange = adminChange
	msg.KeyType = keyType
	msg.KeyPriority = keyPriority
	msg.Key = key
	msg.Timestamp = state.GetTimestamp()

	return msg

}
