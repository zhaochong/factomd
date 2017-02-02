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
	"github.com/FactomProject/factomd/log"
)

var _ = log.Printf

type EOM struct {
	MessageBase
	Timestamp interfaces.Timestamp
	Minute    byte

	DBHeight  uint32
	SysHeight uint32
	SysHash   interfaces.IHash
	ChainID   interfaces.IHash
	Signature interfaces.IFullSignature
	FactoidVM bool

	//Not marshalled
	Processed  bool
	hash       interfaces.IHash
	MarkerSent bool // If we have set EOM markers on blocks like Factoid blocks and such.
}

//var _ interfaces.IConfirmation = (*EOM)(nil)
var _ Signable = (*EOM)(nil)
var _ interfaces.IMsg = (*EOM)(nil)

func (a *EOM) IsSameAs(b *EOM) bool {
	callTime := time.Now().UnixNano()
	defer messagesEOMIsSameAs.Observe(float64(time.Now().UnixNano() - callTime))
	if b == nil {
		return false
	}
	if a.Timestamp.GetTimeMilli() != b.Timestamp.GetTimeMilli() {
		return false
	}
	if a.Minute != b.Minute {
		return false
	}
	if a.DBHeight != b.DBHeight {
		return false
	}

	if a.ChainID == nil && b.ChainID != nil {
		return false
	}
	if a.ChainID != nil {
		if a.ChainID.IsSameAs(b.ChainID) == false {
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

func (e *EOM) Process(dbheight uint32, state interfaces.IState) bool {
	callTime := time.Now().UnixNano()
	defer messagesEOMProcess.Observe(float64(time.Now().UnixNano() - callTime))
	return state.ProcessEOM(dbheight, e)
}

func (m *EOM) GetRepeatHash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer messagesEOMGetRepeatHash.Observe(float64(time.Now().UnixNano() - callTime))
	return m.GetMsgHash()
}

func (m *EOM) GetHash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer messagesEOMGetHash.Observe(float64(time.Now().UnixNano() - callTime))
	return m.GetMsgHash()
}

func (m *EOM) GetMsgHash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer messagesEOMGetMsgHash.Observe(float64(time.Now().UnixNano() - callTime))
	if m.MsgHash == nil {
		data, err := m.MarshalForSignature()
		if err != nil {
			return nil
		}
		m.MsgHash = primitives.Sha(data)
	}
	return m.MsgHash
}

func (m *EOM) GetTimestamp() interfaces.Timestamp {
	callTime := time.Now().UnixNano()
	defer messagesEOMGetTimestamp.Observe(float64(time.Now().UnixNano() - callTime))
	if m.Timestamp == nil {
		m.Timestamp = new(primitives.Timestamp)
	}
	return m.Timestamp
}

func (m *EOM) Int() int {
	callTime := time.Now().UnixNano()
	defer messagesEOMInt.Observe(float64(time.Now().UnixNano() - callTime))
	return int(m.Minute)
}

func (m *EOM) Bytes() []byte {
	callTime := time.Now().UnixNano()
	defer messagesEOMBytes.Observe(float64(time.Now().UnixNano() - callTime))
	var ret []byte
	return append(ret, m.Minute)
}

func (m *EOM) Type() byte {
	callTime := time.Now().UnixNano()
	defer messagesEOMType.Observe(float64(time.Now().UnixNano() - callTime))
	return constants.EOM_MSG
}

// Validate the message, given the state.  Three possible results:
//  < 0 -- Message is invalid.  Discard
//  0   -- Cannot tell if message is Valid
//  1   -- Message is valid
func (m *EOM) Validate(state interfaces.IState) int {
	callTime := time.Now().UnixNano()
	defer messagesEOMValidate.Observe(float64(time.Now().UnixNano() - callTime))
	if m.IsLocal() {
		return 1
	}

	// Ignore old EOM
	if m.DBHeight <= state.GetHighestSavedBlk() {
		return -1
	}

	found, _ := state.GetVirtualServers(m.DBHeight, int(m.Minute), m.ChainID)
	if !found { // Only EOM from federated servers are valid.
		return -1
	}

	// Check signature
	eomSigned, err := m.VerifySignature()
	if err != nil {
		return -1
	}
	if !eomSigned {
		return -1
	}
	return 1
}

// Returns true if this is a message for this server to execute as
// a leader.
func (m *EOM) ComputeVMIndex(state interfaces.IState) {
}

// Execute the leader functions of the given message
func (m *EOM) LeaderExecute(state interfaces.IState) {
	callTime := time.Now().UnixNano()
	defer messagesEOMLeaderExecute.Observe(float64(time.Now().UnixNano() - callTime))
	state.LeaderExecuteEOM(m)
}

func (m *EOM) FollowerExecute(state interfaces.IState) {
	callTime := time.Now().UnixNano()
	defer messagesEOMFollowerExecute.Observe(float64(time.Now().UnixNano() - callTime))
	state.FollowerExecuteEOM(m)
}

func (e *EOM) JSONByte() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer messagesEOMJSONByte.Observe(float64(time.Now().UnixNano() - callTime))
	return primitives.EncodeJSON(e)
}

func (e *EOM) JSONString() (string, error) {
	callTime := time.Now().UnixNano()
	defer messagesEOMJSONString.Observe(float64(time.Now().UnixNano() - callTime))
	return primitives.EncodeJSONString(e)
}

func (e *EOM) JSONBuffer(b *bytes.Buffer) error {
	callTime := time.Now().UnixNano()
	defer messagesEOMJSONBuffer.Observe(float64(time.Now().UnixNano() - callTime))
	return primitives.EncodeJSONToBuffer(e, b)
}

func (m *EOM) Sign(key interfaces.Signer) error {
	callTime := time.Now().UnixNano()
	defer messagesEOMSign.Observe(float64(time.Now().UnixNano() - callTime))
	signature, err := SignSignable(m, key)
	if err != nil {
		return err
	}
	m.Signature = signature
	return nil
}

func (m *EOM) GetSignature() interfaces.IFullSignature {
	callTime := time.Now().UnixNano()
	defer messagesEOMGetSignature.Observe(float64(time.Now().UnixNano() - callTime))
	return m.Signature
}

func (m *EOM) VerifySignature() (bool, error) {
	callTime := time.Now().UnixNano()
	defer messagesEOMVerifySignature.Observe(float64(time.Now().UnixNano() - callTime))
	return VerifyMessage(m)
}

func (m *EOM) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	callTime := time.Now().UnixNano()
	defer messagesEOMUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error unmarshalling EOM message: %v", r)
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

	m.ChainID = primitives.NewHash(constants.ZERO_HASH)
	newData, err = m.ChainID.UnmarshalBinaryData(newData)
	if err != nil {
		return nil, err
	}

	m.Minute, newData = newData[0], newData[1:]

	if m.Minute < 0 || m.Minute >= 10 {
		return nil, fmt.Errorf("Minute number is out of range")
	}

	m.VMIndex = int(newData[0])
	newData = newData[1:]
	m.FactoidVM = uint8(newData[0]) == 1
	newData = newData[1:]

	m.DBHeight, newData = binary.BigEndian.Uint32(newData[0:4]), newData[4:]
	m.SysHeight, newData = binary.BigEndian.Uint32(newData[0:4]), newData[4:]

	m.SysHash = primitives.NewHash(constants.ZERO_HASH)
	newData, err = m.SysHash.UnmarshalBinaryData(newData)

	if len(newData) > 0 {
		sig := new(primitives.Signature)
		newData, err = sig.UnmarshalBinaryData(newData)
		if err != nil {
			return nil, err
		}
		m.Signature = sig
	}

	return data, nil
}

func (m *EOM) UnmarshalBinary(data []byte) error {
	callTime := time.Now().UnixNano()
	defer messagesEOMUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	_, err := m.UnmarshalBinaryData(data)
	return err
}

func (m *EOM) MarshalForSignature() (data []byte, err error) {
	callTime := time.Now().UnixNano()
	defer messagesEOMMarshalForSignature.Observe(float64(time.Now().UnixNano() - callTime))
	var buf primitives.Buffer
	buf.Write([]byte{m.Type()})
	if d, err := m.Timestamp.MarshalBinary(); err != nil {
		return nil, err
	} else {
		buf.Write(d)
	}

	if d, err := m.ChainID.MarshalBinary(); err != nil {
		return nil, err
	} else {
		buf.Write(d)
	}

	binary.Write(&buf, binary.BigEndian, m.Minute)
	binary.Write(&buf, binary.BigEndian, uint8(m.VMIndex))
	if m.FactoidVM {
		binary.Write(&buf, binary.BigEndian, uint8(1))
	} else {
		binary.Write(&buf, binary.BigEndian, uint8(0))
	}
	return buf.DeepCopyBytes(), nil
}

func (m *EOM) MarshalBinary() (data []byte, err error) {
	callTime := time.Now().UnixNano()
	defer messagesEOM.Observe(float64(time.Now().UnixNano() - callTime))
	var buf primitives.Buffer
	resp, err := m.MarshalForSignature()
	if err != nil {
		return nil, err
	}
	buf.Write(resp)

	binary.Write(&buf, binary.BigEndian, m.DBHeight)
	binary.Write(&buf, binary.BigEndian, m.SysHeight)

	if m.SysHash == nil {
		m.SysHash = primitives.NewHash(constants.ZERO_HASH)
	}
	if d, err := m.SysHash.MarshalBinary(); err != nil {
		return nil, err
	} else {
		buf.Write(d)
	}

	sig := m.GetSignature()
	if sig != nil {
		sigBytes, err := sig.MarshalBinary()
		if err != nil {
			return nil, err
		}
		buf.Write(sigBytes)
	}
	return buf.DeepCopyBytes(), nil
}

func (m *EOM) String() string {
	callTime := time.Now().UnixNano()
	defer messagesEOMString.Observe(float64(time.Now().UnixNano() - callTime))
	local := ""
	if m.IsLocal() {
		local = "local"
	}
	f := "-"
	if m.FactoidVM {
		f = "F"
	}
	return fmt.Sprintf("%6s-VM%3d: Min:%4d DBHt:%5d FF %2d -%1s-Leader[%x] hash[%x] %s",
		"EOM",
		m.VMIndex,
		m.Minute,
		m.DBHeight,
		m.SysHeight,
		f,
		m.ChainID.Bytes()[:4],
		m.GetMsgHash().Bytes()[:3],
		local)
}
