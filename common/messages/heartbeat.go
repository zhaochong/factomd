// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package messages

import (
	"encoding/binary"
	"fmt"

	"github.com/FactomProject/factomd/common/constants"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
	"time"
)

//A placeholder structure for messages
type Heartbeat struct {
	MessageBase
	Timestamp       interfaces.Timestamp
	SecretNumber    uint32
	DBHeight        uint32
	DBlockHash      interfaces.IHash //Hash of last Directory Block
	IdentityChainID interfaces.IHash //Identity Chain ID

	Signature interfaces.IFullSignature

	//Not marshalled
	hash     interfaces.IHash
	sigvalid bool
}

var _ interfaces.IMsg = (*Heartbeat)(nil)
var _ Signable = (*Heartbeat)(nil)

func (a *Heartbeat) IsSameAs(b *Heartbeat) bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesHeartbeatIsSameAs.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if b == nil {
		return false
	}
	if a.Timestamp.GetTimeMilli() != b.Timestamp.GetTimeMilli() {
		return false
	}

	if a.DBlockHash == nil && b.DBlockHash != nil {
		return false
	}
	if a.DBlockHash != nil {
		if a.DBlockHash.IsSameAs(b.DBlockHash) == false {
			return false
		}
	}

	if a.IdentityChainID == nil && b.IdentityChainID != nil {
		return false
	}
	if a.IdentityChainID != nil {
		if a.IdentityChainID.IsSameAs(b.IdentityChainID) == false {
			return false
		}
	}

	if a.Signature == nil && b.Signature != nil {
		return false
	}
	if a.Signature != nil {
		if a.Signature.IsSameAs(b.Signature) == false {
			return false
		}
	}

	return true
}

func (m *Heartbeat) Process(uint32, interfaces.IState) bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesHeartbeatProcess.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return true
}

func (m *Heartbeat) GetRepeatHash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesHeartbeatGetRepeatHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return m.GetMsgHash()
}

func (m *Heartbeat) GetHash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesHeartbeatGetHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if m.hash == nil {
		data, err := m.MarshalForSignature()
		if err != nil {
			panic(fmt.Sprintf("Error in CommitChain.GetHash(): %s", err.Error()))
		}
		m.hash = primitives.Sha(data)
	}
	return m.hash
}

func (m *Heartbeat) GetMsgHash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesHeartbeatGetMsgHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if m.MsgHash == nil {
		data, err := m.MarshalBinary()
		if err != nil {
			return nil
		}
		m.MsgHash = primitives.Sha(data)
	}
	return m.MsgHash
}

func (m *Heartbeat) GetTimestamp() interfaces.Timestamp {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesHeartbeatGetTimestamp.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return m.Timestamp
}

func (m *Heartbeat) Type() byte {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesHeartbeatType.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return constants.HEARTBEAT_MSG
}

func (m *Heartbeat) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesHeartbeatUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error unmarshalling HeartBeat: %v", r)
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

	m.SecretNumber, newData = binary.BigEndian.Uint32(newData[0:4]), newData[4:]
	m.DBHeight, newData = binary.BigEndian.Uint32(newData[0:4]), newData[4:]

	hash := new(primitives.Hash)

	newData, err = hash.UnmarshalBinaryData(newData)
	if err != nil {
		return nil, err
	}
	m.DBlockHash = hash

	hash = new(primitives.Hash)
	newData, err = hash.UnmarshalBinaryData(newData)
	if err != nil {
		return nil, err
	}
	m.IdentityChainID = hash

	if len(newData) > 0 {
		sig := new(primitives.Signature)
		newData, err = sig.UnmarshalBinaryData(newData)
		if err != nil {
			return nil, err
		}
		m.Signature = sig
	}

	return nil, nil
}

func (m *Heartbeat) UnmarshalBinary(data []byte) error {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesHeartbeatUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	_, err := m.UnmarshalBinaryData(data)
	return err
}

func (m *Heartbeat) MarshalForSignature() (data []byte, err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesHeartbeatMarshalForSignature.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if m.DBlockHash == nil || m.IdentityChainID == nil {
		return nil, fmt.Errorf("Message is incomplete")
	}

	var buf primitives.Buffer
	buf.Write([]byte{m.Type()})
	if d, err := m.Timestamp.MarshalBinary(); err != nil {
		return nil, err
	} else {
		buf.Write(d)
	}

	binary.Write(&buf, binary.BigEndian, m.SecretNumber)
	binary.Write(&buf, binary.BigEndian, m.DBHeight)

	if d, err := m.DBlockHash.MarshalBinary(); err != nil {
		return nil, err
	} else {
		buf.Write(d)
	}

	if d, err := m.IdentityChainID.MarshalBinary(); err != nil {
		return nil, err
	} else {
		buf.Write(d)
	}

	return buf.DeepCopyBytes(), nil
}

func (m *Heartbeat) MarshalBinary() (data []byte, err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesHeartbeatMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	resp, err := m.MarshalForSignature()
	if err != nil {
		return nil, err
	}
	sig := m.GetSignature()
	if sig != nil {
		sigBytes, err := sig.MarshalBinary()
		if err != nil {
			return nil, err
		}
		return append(resp, sigBytes...), nil
	}
	return resp, nil
}

func (m *Heartbeat) String() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesHeartbeatString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return fmt.Sprintf("HeartBeat ID[%x] dbht %d ts %d", m.IdentityChainID.Bytes()[3:5], m.DBHeight, m.Timestamp.GetTimeSeconds())
}

func (m *Heartbeat) ChainID() []byte {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesHeartbeatChainID.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return nil
}

func (m *Heartbeat) ListHeight() int {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesHeartbeatListHeight.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return 0
}

func (m *Heartbeat) SerialHash() []byte {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesHeartbeatSerialHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return nil
}

// Validate the message, given the state.  Three possible results:
//  < 0 -- Message is invalid.  Discard
//  0   -- Cannot tell if message is Valid
//  1   -- Message is valid
func (m *Heartbeat) Validate(state interfaces.IState) int {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesHeartbeatValidate.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	now := state.GetTimestamp()

	if now.GetTimeSeconds()-m.Timestamp.GetTimeSeconds() > 60 {
		return -1
	}

	if m.GetSignature() == nil {
		// the message has no signature (and so is invalid)
		return -1
	}

	// Ignore old heartbeats
	if m.DBHeight <= state.GetHighestSavedBlk() {
		return -1
	}

	if !m.sigvalid {
		isVer, err := m.VerifySignature()
		if err != nil || !isVer {
			// if there is an error during signature verification
			// or if the signature is invalid
			// the message is considered invalid
			return -1
		}
		m.sigvalid = true
	}

	return 1
}

// Returns true if this is a message for this server to execute as
// a leader.
func (m *Heartbeat) ComputeVMIndex(state interfaces.IState) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesHeartbeatComputeVMIndex.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

}

// Execute the leader functions of the given message
func (m *Heartbeat) LeaderExecute(state interfaces.IState) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesHeartbeatLeaderExecute.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	m.FollowerExecute(state)
}

func (m *Heartbeat) FollowerExecute(state interfaces.IState) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesHeartbeatFollowerExecute.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	for _, auditServer := range state.GetAuditServers(state.GetLeaderHeight()) {
		if auditServer.GetChainID().IsSameAs(m.IdentityChainID) {
			if m.IdentityChainID.IsSameAs(state.GetIdentityChainID()) {
				if m.SecretNumber != state.GetSalt(m.Timestamp) {
					panic("We have seen a heartbeat using our Identity that isn't ours")
				}
			}
			auditServer.SetOnline(true)
		}
	}
}

func (e *Heartbeat) JSONByte() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesHeartbeatJSONByte.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSON(e)
}

func (e *Heartbeat) JSONString() (string, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesHeartbeatJSONString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSONString(e)
}

func (m *Heartbeat) Sign(key interfaces.Signer) error {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesHeartbeatSign.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	signature, err := SignSignable(m, key)
	if err != nil {
		return err
	}
	m.Signature = signature
	return nil
}

func (m *Heartbeat) GetSignature() interfaces.IFullSignature {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesHeartbeatGetSignature.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return m.Signature
}

func (m *Heartbeat) VerifySignature() (bool, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesHeartbeatVerifySignature.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return VerifyMessage(m)
}
