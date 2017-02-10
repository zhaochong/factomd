// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package messages

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/FactomProject/factomd/common/constants"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
	"time"
)

// Communicate a Directory Block State

type AddServerMsg struct {
	MessageBase
	Timestamp     interfaces.Timestamp // Message Timestamp
	ServerChainID interfaces.IHash     // ChainID of new server
	ServerType    int                  // 0 = Federated, 1 = Audit

	Signature interfaces.IFullSignature
}

var _ interfaces.IMsg = (*AddServerMsg)(nil)
var _ Signable = (*AddServerMsg)(nil)

func (m *AddServerMsg) GetRepeatHash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesAddServerMsgGetRepeatHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return m.GetMsgHash()
}

func (m *AddServerMsg) GetHash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesAddServerMsgGetHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return m.GetMsgHash()
}

func (m *AddServerMsg) GetMsgHash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesAddServerMsgGetMsgHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if m.MsgHash == nil {
		data, err := m.MarshalForSignature()
		if err != nil {
			return nil
		}
		m.MsgHash = primitives.Sha(data)
	}
	return m.MsgHash
}

func (m *AddServerMsg) Type() byte {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesAddServerMsgType.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return constants.ADDSERVER_MSG
}

func (m *AddServerMsg) GetTimestamp() interfaces.Timestamp {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesAddServerMsgGetTimestamp.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return m.Timestamp
}

func (m *AddServerMsg) Validate(state interfaces.IState) int {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesAddServerMsgValidate.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	//return 1
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
func (m *AddServerMsg) ComputeVMIndex(state interfaces.IState) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesAddServerMsgComputeVMIndex.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	m.VMIndex = state.ComputeVMIndex(constants.ADMIN_CHAINID)
}

// Execute the leader functions of the given message
func (m *AddServerMsg) LeaderExecute(state interfaces.IState) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesAddServerMsgLeaderExecute.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	state.LeaderExecute(m)
}

func (m *AddServerMsg) FollowerExecute(state interfaces.IState) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesAddServerMsgFollowerExecute.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	state.FollowerExecuteMsg(m)
}

// Acknowledgements do not go into the process list.
func (e *AddServerMsg) Process(dbheight uint32, state interfaces.IState) bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesAddServerMsgProcess.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return state.ProcessAddServer(dbheight, e)
}

func (e *AddServerMsg) JSONByte() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesAddServerMsgJSONByte.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSON(e)
}

func (e *AddServerMsg) JSONString() (string, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesAddServerMsgJSONString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSONString(e)
}

func (m *AddServerMsg) Sign(key interfaces.Signer) error {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesAddServerMsgSign.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	signature, err := SignSignable(m, key)
	if err != nil {
		return err
	}
	m.Signature = signature
	return nil
}

func (m *AddServerMsg) GetSignature() interfaces.IFullSignature {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesAddServerMsgGetSignature.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return m.Signature
}

func (m *AddServerMsg) VerifySignature() (bool, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesAddServerMsgVerifySignature.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return VerifyMessage(m)
}

func (m *AddServerMsg) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesAddServerMsgUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	defer func() {
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

func (m *AddServerMsg) UnmarshalBinary(data []byte) error {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesAddServerMsgUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	_, err := m.UnmarshalBinaryData(data)
	return err
}

func (m *AddServerMsg) MarshalForSignature() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesAddServerMsgMarshalForSignature.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

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

func (m *AddServerMsg) MarshalBinary() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesAddServerMsgMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

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

func (m *AddServerMsg) String() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesAddServerMsgString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	var stype string
	if m.ServerType == 0 {
		stype = "Federated"
	} else {
		stype = "Audit"
	}
	return fmt.Sprintf("AddServer (%s): ChainID: %x Time: %x Msg Hash %x ",
		stype,
		m.ServerChainID.Bytes()[:3],
		&m.Timestamp,
		m.GetMsgHash().Bytes()[:3])

}

func (m *AddServerMsg) IsSameAs(b *AddServerMsg) bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesAddServerMsgIsSameAs.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

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

func NewAddServerMsg(state interfaces.IState, serverType int) interfaces.IMsg {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesNewAddServerMsg.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	msg := new(AddServerMsg)
	msg.ServerChainID = state.GetIdentityChainID()
	msg.ServerType = serverType
	msg.Timestamp = state.GetTimestamp()

	return msg

}

func NewAddServerByHashMsg(state interfaces.IState, serverType int, newServerHash interfaces.IHash) interfaces.IMsg {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesNewAddServerByHashMsg.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	msg := new(AddServerMsg)
	msg.ServerChainID = newServerHash
	msg.ServerType = serverType
	msg.Timestamp = state.GetTimestamp()

	return msg
}
