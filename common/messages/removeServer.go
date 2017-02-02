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

// Communicate a Directory Block State

type RemoveServerMsg struct {
	MessageBase
	Timestamp     interfaces.Timestamp // Message Timestamp
	ServerChainID interfaces.IHash     // ChainID of new server
	ServerType    int                  // 0 = Federated, 1 = Audit

	Signature interfaces.IFullSignature
}

var _ interfaces.IMsg = (*RemoveServerMsg)(nil)
var _ Signable = (*RemoveServerMsg)(nil)

func (m *RemoveServerMsg) GetRepeatHash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer messagesRemoveServerMsgGetRepeatHash.Observe(float64(time.Now().UnixNano() - callTime))
	return m.GetMsgHash()
}

func (m *RemoveServerMsg) GetHash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer messagesRemoveServerMsgGetHash.Observe(float64(time.Now().UnixNano() - callTime))
	return m.GetMsgHash()
}

func (m *RemoveServerMsg) GetMsgHash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer messagesRemoveServerMsgGetMsgHash.Observe(float64(time.Now().UnixNano() - callTime))
	if m.MsgHash == nil {
		data, err := m.MarshalForSignature()
		if err != nil {
			return nil
		}
		m.MsgHash = primitives.Sha(data)
	}
	return m.MsgHash
}

func (m *RemoveServerMsg) Type() byte {
	callTime := time.Now().UnixNano()
	defer messagesRemoveServerMsgType.Observe(float64(time.Now().UnixNano() - callTime))
	return constants.REMOVESERVER_MSG
}

func (m *RemoveServerMsg) Int() int {
	callTime := time.Now().UnixNano()
	defer messagesRemoveServerMsgInt.Observe(float64(time.Now().UnixNano() - callTime))
	return -1
}

func (m *RemoveServerMsg) Bytes() []byte {
	callTime := time.Now().UnixNano()
	defer messagesRemoveServerMsgBytes.Observe(float64(time.Now().UnixNano() - callTime))
	return nil
}

func (m *RemoveServerMsg) GetTimestamp() interfaces.Timestamp {
	callTime := time.Now().UnixNano()
	defer messagesRemoveServerMsgGetTimestamp.Observe(float64(time.Now().UnixNano() - callTime))
	return m.Timestamp
}

func (m *RemoveServerMsg) Validate(state interfaces.IState) int {
	callTime := time.Now().UnixNano()
	defer messagesRemoveServerMsgValidate.Observe(float64(time.Now().UnixNano() - callTime))
	// Check to see if identity exists and is audit or fed server
	if !state.VerifyIsAuthority(m.ServerChainID) {
		//fmt.Printf("RemoveServerMsg Error: [%s] is not a server, cannot be removed\n", m.ServerChainID.String()[:8])
		return -1
	}

	authoritativeKey := state.GetNetworkSkeletonKey().Bytes()
	if m.GetSignature() == nil || bytes.Compare(m.GetSignature().GetKey(), authoritativeKey) != 0 {
		// the message was not signed with the proper authoritative signing key (from conf file)
		// it is therefore considered invalid
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
func (m *RemoveServerMsg) ComputeVMIndex(state interfaces.IState) {
	callTime := time.Now().UnixNano()
	defer messagesRemoveServerMsgComputeVMIndex.Observe(float64(time.Now().UnixNano() - callTime))
	m.VMIndex = state.ComputeVMIndex(constants.ADMIN_CHAINID)
}

// Execute the leader functions of the given message
func (m *RemoveServerMsg) LeaderExecute(state interfaces.IState) {
	callTime := time.Now().UnixNano()
	defer messagesRemoveServerMsg.Observe(float64(time.Now().UnixNano() - callTime))
	state.LeaderExecute(m)
}

func (m *RemoveServerMsg) FollowerExecute(state interfaces.IState) {
	callTime := time.Now().UnixNano()
	defer messagesRemoveServerMsg.Observe(float64(time.Now().UnixNano() - callTime))
	state.FollowerExecuteMsg(m)
}

// Acknowledgements do not go into the process list.
func (e *RemoveServerMsg) Process(dbheight uint32, state interfaces.IState) bool {
	callTime := time.Now().UnixNano()
	defer messagesRemoveServerMsgProcess.Observe(float64(time.Now().UnixNano() - callTime))
	return state.ProcessRemoveServer(dbheight, e)
}

func (e *RemoveServerMsg) JSONByte() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer messagesRemoveServerMsgJSONByte.Observe(float64(time.Now().UnixNano() - callTime))
	return primitives.EncodeJSON(e)
}

func (e *RemoveServerMsg) JSONString() (string, error) {
	callTime := time.Now().UnixNano()
	defer messagesRemoveServerMsgJSONString.Observe(float64(time.Now().UnixNano() - callTime))
	return primitives.EncodeJSONString(e)
}

func (e *RemoveServerMsg) JSONBuffer(b *bytes.Buffer) error {
	callTime := time.Now().UnixNano()
	defer messagesRemoveServerMsgJSONBuffer.Observe(float64(time.Now().UnixNano() - callTime))
	return primitives.EncodeJSONToBuffer(e, b)
}

func (m *RemoveServerMsg) Sign(key interfaces.Signer) error {
	callTime := time.Now().UnixNano()
	defer messagesRemoveServerMsgSign.Observe(float64(time.Now().UnixNano() - callTime))
	signature, err := SignSignable(m, key)
	if err != nil {
		return err
	}
	m.Signature = signature
	return nil
}

func (m *RemoveServerMsg) GetSignature() interfaces.IFullSignature {
	callTime := time.Now().UnixNano()
	defer messagesRemoveServerMsg.Observe(float64(time.Now().UnixNano() - callTime))
	return m.Signature
}

func (m *RemoveServerMsg) VerifySignature() (bool, error) {
	callTime := time.Now().UnixNano()
	defer messagesRemoveServerMsg.Observe(float64(time.Now().UnixNano() - callTime))
	return VerifyMessage(m)
}

func (m *RemoveServerMsg) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	callTime := time.Now().UnixNano()
	defer messagesRemoveServerMsg.Observe(float64(time.Now().UnixNano() - callTime))
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

	m.ServerChainID = new(primitives.Hash)
	newData, err = m.ServerChainID.UnmarshalBinaryData(newData)
	if err != nil {
		return nil, err
	}

	m.ServerType = int(newData[0])
	newData = newData[1:]

	if len(newData) > 32 {
		m.Signature = new(primitives.Signature)
		newData, err = m.Signature.UnmarshalBinaryData(newData)
		if err != nil {
			return nil, err
		}
	}
	return
}

func (m *RemoveServerMsg) UnmarshalBinary(data []byte) error {
	callTime := time.Now().UnixNano()
	defer messagesRemoveServerMsgUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	_, err := m.UnmarshalBinaryData(data)
	return err
}

func (m *RemoveServerMsg) MarshalForSignature() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer messagesRemoveServerMsgMarshalForSignature.Observe(float64(time.Now().UnixNano() - callTime))
	var buf primitives.Buffer

	binary.Write(&buf, binary.BigEndian, m.Type())

	t := m.GetTimestamp()
	data, err := t.MarshalBinary()
	if err != nil {
		return nil, err
	}
	buf.Write(data)

	data, err = m.ServerChainID.MarshalBinary()
	if err != nil {
		return nil, err
	}
	buf.Write(data)

	binary.Write(&buf, binary.BigEndian, uint8(m.ServerType))

	return buf.DeepCopyBytes(), nil
}

func (m *RemoveServerMsg) MarshalBinary() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer messagesRemoveServerMsgMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	var buf primitives.Buffer

	data, err := m.MarshalForSignature()
	if err != nil {
		return nil, err
	}
	buf.Write(data)

	if m.Signature != nil {
		data, err = m.Signature.MarshalBinary()
		if err != nil {
			return nil, err
		}
		buf.Write(data)
	}

	return buf.DeepCopyBytes(), nil
}

func (m *RemoveServerMsg) String() string {
	callTime := time.Now().UnixNano()
	defer messagesRemoveServerMsgString.Observe(float64(time.Now().UnixNano() - callTime))
	var stype string
	if m.ServerType == 0 {
		stype = "Federated"
	} else {
		stype = "Audit"
	}
	return fmt.Sprintf("RemoveServer (%s): ChainID: %x Time: %x Msg Hash %x ",
		stype,
		m.ServerChainID.Bytes()[:3],
		&m.Timestamp,
		m.GetMsgHash().Bytes()[:3])

}

func (m *RemoveServerMsg) IsSameAs(b *RemoveServerMsg) bool {
	callTime := time.Now().UnixNano()
	defer messagesRemoveServerMsgIsSameAs.Observe(float64(time.Now().UnixNano() - callTime))
	if b == nil {
		return false
	}
	if m.Timestamp.GetTimeMilli() != b.Timestamp.GetTimeMilli() {
		return false
	}
	if !m.ServerChainID.IsSameAs(b.ServerChainID) {
		return false
	}
	if m.ServerType != b.ServerType {
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

func NewRemoveServerMsg(state interfaces.IState, chainId interfaces.IHash, serverType int) interfaces.IMsg {
	callTime := time.Now().UnixNano()
	defer messagesRemoveServerMsgNewRemoveServerMsg.Observe(float64(time.Now().UnixNano() - callTime))
	msg := new(RemoveServerMsg)
	msg.ServerChainID = chainId
	msg.ServerType = serverType
	msg.Timestamp = state.GetTimestamp()

	return msg

}
