// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package messages

import (
	"bytes"
	"fmt"
	"time"

	"encoding/binary"

	"github.com/FactomProject/factomd/common/constants"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
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
	callTime := time.Now().UnixNano()
	defer messagesHeartbeatIsSameAs.Observe(float64(time.Now().UnixNano() - callTime))
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
	return true
}

func (m *Heartbeat) GetRepeatHash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer messagesHeartbeatGetRepeatHash.Observe(float64(time.Now().UnixNano() - callTime))
	return m.GetMsgHash()
}

func (m *Heartbeat) GetHash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer messagesHeartbeatGetHash.Observe(float64(time.Now().UnixNano() - callTime))
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
	callTime := time.Now().UnixNano()
	defer messagesHeartbeat.Observe(float64(time.Now().UnixNano() - callTime))
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
	callTime := time.Now().UnixNano()
	defer messagesHeartbeatGetTimestamp.Observe(float64(time.Now().UnixNano() - callTime))
	return m.Timestamp
}

func (m *Heartbeat) Type() byte {
	callTime := time.Now().UnixNano()
	defer messagesHeartbeatType.Observe(float64(time.Now().UnixNano() - callTime))
	return constants.HEARTBEAT_MSG
}

func (m *Heartbeat) Int() int {
	callTime := time.Now().UnixNano()
	defer messagesHeartbeatInt.Observe(float64(time.Now().UnixNano() - callTime))
	return -1
}

func (m *Heartbeat) Bytes() []byte {
	callTime := time.Now().UnixNano()
	defer messagesHeartbeatBytes.Observe(float64(time.Now().UnixNano() - callTime))
	return nil
}

func (m *Heartbeat) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	callTime := time.Now().UnixNano()
	defer messagesHeartbeatUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))
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
	callTime := time.Now().UnixNano()
	defer messagesHeartbeatUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	_, err := m.UnmarshalBinaryData(data)
	return err
}

func (m *Heartbeat) MarshalForSignature() (data []byte, err error) {
	callTime := time.Now().UnixNano()
	defer messagesHeartbeatMarshalForSignature.Observe(float64(time.Now().UnixNano() - callTime))
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
	callTime := time.Now().UnixNano()
	defer messagesHeartbeatMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
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
	callTime := time.Now().UnixNano()
	defer messagesHeartbeatString.Observe(float64(time.Now().UnixNano() - callTime))
	return fmt.Sprintf("HeartBeat ID[%x] dbht %d ts %d", m.IdentityChainID.Bytes()[3:5], m.DBHeight, m.Timestamp.GetTimeSeconds())
}

func (m *Heartbeat) ChainID() []byte {
	return nil
}

func (m *Heartbeat) ListHeight() int {
	return 0
}

func (m *Heartbeat) SerialHash() []byte {
	return nil
}

// Validate the message, given the state.  Three possible results:
//  < 0 -- Message is invalid.  Discard
//  0   -- Cannot tell if message is Valid
//  1   -- Message is valid
func (m *Heartbeat) Validate(state interfaces.IState) int {
	callTime := time.Now().UnixNano()
	defer messagesHeartbeatValidate.Observe(float64(time.Now().UnixNano() - callTime))
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

}

// Execute the leader functions of the given message
func (m *Heartbeat) LeaderExecute(state interfaces.IState) {
	callTime := time.Now().UnixNano()
	defer messagesHeartbeatLeaderExecute.Observe(float64(time.Now().UnixNano() - callTime))
	m.FollowerExecute(state)
}

func (m *Heartbeat) FollowerExecute(state interfaces.IState) {
	callTime := time.Now().UnixNano()
	defer messagesHeartbeat.Observe(float64(time.Now().UnixNano() - callTime))
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
	callTime := time.Now().UnixNano()
	defer messagesHeartbeatJSONByte.Observe(float64(time.Now().UnixNano() - callTime))
	return primitives.EncodeJSON(e)
}

func (e *Heartbeat) JSONString() (string, error) {
	callTime := time.Now().UnixNano()
	defer messagesHeartbeatJSONString.Observe(float64(time.Now().UnixNano() - callTime))
	return primitives.EncodeJSONString(e)
}

func (e *Heartbeat) JSONBuffer(b *bytes.Buffer) error {
	callTime := time.Now().UnixNano()
	defer messagesHeartbeatJSONBuffer.Observe(float64(time.Now().UnixNano() - callTime))
	return primitives.EncodeJSONToBuffer(e, b)
}

func (m *Heartbeat) Sign(key interfaces.Signer) error {
	callTime := time.Now().UnixNano()
	defer messagesHeartbeat.Observe(float64(time.Now().UnixNano() - callTime))
	signature, err := SignSignable(m, key)
	if err != nil {
		return err
	}
	m.Signature = signature
	return nil
}

func (m *Heartbeat) GetSignature() interfaces.IFullSignature {
	callTime := time.Now().UnixNano()
	defer messagesHeartbeat.Observe(float64(time.Now().UnixNano() - callTime))
	return m.Signature
}

func (m *Heartbeat) VerifySignature() (bool, error) {
	callTime := time.Now().UnixNano()
	defer messagesHeartbeat.Observe(float64(time.Now().UnixNano() - callTime))
	return VerifyMessage(m)
}
