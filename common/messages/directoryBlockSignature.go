// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package messages

import (
	"encoding/binary"
	"fmt"

	"github.com/FactomProject/factomd/common/constants"
	"github.com/FactomProject/factomd/common/directoryBlock"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
	"time"
)

type DirectoryBlockSignature struct {
	MessageBase
	Timestamp interfaces.Timestamp
	DBHeight  uint32
	//DirectoryBlockKeyMR   interfaces.IHash
	DirectoryBlockHeader  interfaces.IDirectoryBlockHeader
	ServerIdentityChainID interfaces.IHash

	Signature interfaces.IFullSignature

	// Signature that goes into the admin block
	// Signature of directory block header
	DBSignature interfaces.IFullSignature
	SysHeight   uint32
	SysHash     interfaces.IHash

	//Not marshalled
	Matches   bool
	Processed bool
	hash      interfaces.IHash
}

var _ interfaces.IMsg = (*DirectoryBlockSignature)(nil)
var _ Signable = (*DirectoryBlockSignature)(nil)

func (a *DirectoryBlockSignature) IsSameAs(b *DirectoryBlockSignature) bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesDirectoryBlockSignatureIsSameAs.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if b == nil {
		return false
	}

	if a.Timestamp.GetTimeMilli() != b.Timestamp.GetTimeMilli() {
		return false
	}
	if a.DBHeight != b.DBHeight {
		return false
	}

	if a.DirectoryBlockHeader == nil && b.DirectoryBlockHeader != nil {
		return false
	}
	if a.DirectoryBlockHeader != nil {
		if a.DirectoryBlockHeader.GetPrevFullHash().IsSameAs(b.DirectoryBlockHeader.GetPrevFullHash()) == false {
			return false
		}
	}

	if a.ServerIdentityChainID == nil && b.ServerIdentityChainID != nil {
		return false
	}
	if a.ServerIdentityChainID != nil {
		if a.ServerIdentityChainID.IsSameAs(b.ServerIdentityChainID) == false {
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

func (e *DirectoryBlockSignature) Process(dbheight uint32, state interfaces.IState) bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesDirectoryBlockSignatureProcess.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return state.ProcessDBSig(dbheight, e)
}

func (m *DirectoryBlockSignature) GetRepeatHash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesDirectoryBlockSignatureGetRepeatHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return m.GetMsgHash()
}

func (m *DirectoryBlockSignature) GetHash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesDirectoryBlockSignatureGetHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return m.GetMsgHash()
}

func (m *DirectoryBlockSignature) GetMsgHash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesDirectoryBlockSignatureGetMsgHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	data, _ := m.MarshalForSignature()
	if data == nil {
		return nil
	}
	m.MsgHash = primitives.Sha(data)

	return m.MsgHash
}

func (m *DirectoryBlockSignature) GetTimestamp() interfaces.Timestamp {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesDirectoryBlockSignatureGetTimestamp.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if m.Timestamp == nil {
		m.Timestamp = new(primitives.Timestamp)
	}
	return m.Timestamp
}

func (m *DirectoryBlockSignature) Type() byte {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesDirectoryBlockSignatureType.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return constants.DIRECTORY_BLOCK_SIGNATURE_MSG
}

// Validate the message, given the state.  Three possible results:
//  < 0 -- Message is invalid.  Discard
//  0   -- Cannot tell if message is Valid
//  1   -- Message is valid
func (m *DirectoryBlockSignature) Validate(state interfaces.IState) int {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesDirectoryBlockSignatureValidate.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if m.DBHeight < state.GetLLeaderHeight() {
		state.AddStatus(fmt.Sprintf("DirectoryBlockSignature: Fail dbht: %v %s", state.GetLLeaderHeight(), m.String()))
		return -1
	}

	if m.DBHeight > state.GetLLeaderHeight() {
		//state.AddStatus(fmt.Sprintf("DirectoryBlockSignature: Wait dbht: %v %s", state.GetLLeaderHeight(), m.String()))
		return 0
	}

	found, _ := state.GetVirtualServers(m.DBHeight, 9, m.ServerIdentityChainID)

	if found == false {
		state.AddStatus(fmt.Sprintf("DirectoryBlockSignature: Fail dbht: %v Server not found %x %s",
			state.GetLLeaderHeight(),
			m.ServerIdentityChainID.Bytes()[3:5],
			m.String()))
		return 0
	}

	if m.IsLocal() {
		return 1
	}

	isVer, err := m.VerifySignature()
	if err != nil || !isVer {
		state.AddStatus(fmt.Sprintf("DirectoryBlockSignature: Fail to Verify Sig dbht: %v %s", state.GetLLeaderHeight(), m.String()))
		// if there is an error during signature verification
		// or if the signature is invalid
		// the message is considered invalid
		return -1
	}

	marshalledMsg, _ := m.MarshalForSignature()
	authorityLevel, err := state.VerifyAuthoritySignature(marshalledMsg, m.Signature.GetSignature(), m.DBHeight)
	if err != nil || authorityLevel < 1 {
		//This authority is not a Fed Server (it's either an Audit or not an Authority at all)
		state.AddStatus(fmt.Sprintf("DirectoryBlockSignature: Fail to Verify Sig (not from a Fed Server) dbht: %v %s", state.GetLLeaderHeight(), m.String()))
		return -1
	}

	return 1
}

// Returns true if this is a message for this server to execute as
// a leader.
func (m *DirectoryBlockSignature) ComputeVMIndex(state interfaces.IState) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesDirectoryBlockSignatureComputeVMIndex.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

}

// Execute the leader functions of the given message
func (m *DirectoryBlockSignature) LeaderExecute(state interfaces.IState) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesDirectoryBlockSignatureLeaderExecute.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	state.LeaderExecuteDBSig(m)
}

func (m *DirectoryBlockSignature) FollowerExecute(state interfaces.IState) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesDirectoryBlockSignatureFollowerExecute.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	state.FollowerExecuteMsg(m)
}

func (m *DirectoryBlockSignature) Sign(key interfaces.Signer) error {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesDirectoryBlockSignatureSign.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	// Signature that goes into admin block
	err := m.SignHeader(key)
	if err != nil {
		return err
	}

	signature, err := SignSignable(m, key)
	if err != nil {
		return err
	}
	m.Signature = signature
	return nil
}

func (m *DirectoryBlockSignature) SignHeader(key interfaces.Signer) error {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesDirectoryBlockSignatureSignHeader.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	header, err := m.DirectoryBlockHeader.MarshalBinary()
	if err != nil {
		return err
	}
	m.DBSignature = key.Sign(header)
	return nil
}

func (m *DirectoryBlockSignature) GetSignature() interfaces.IFullSignature {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesDirectoryBlockSignatureGetSignature.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return m.Signature
}

func (m *DirectoryBlockSignature) VerifySignature() (bool, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesDirectoryBlockSignatureVerifySignature.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return VerifyMessage(m)
}

func (m *DirectoryBlockSignature) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesDirectoryBlockSignatureUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error unmarshalling: %v", r)
		}
	}()

	newData = data
	if newData[0] != m.Type() {
		return nil, fmt.Errorf("Invalid Message type")
	}
	newData = newData[1:]

	// TimeStamp
	m.Timestamp = new(primitives.Timestamp)
	newData, err = m.Timestamp.UnmarshalBinaryData(newData)
	if err != nil {
		return nil, err
	}

	m.SysHeight, newData = binary.BigEndian.Uint32(newData[0:4]), newData[4:]
	hash := new(primitives.Hash)
	newData, err = hash.UnmarshalBinaryData(newData)
	if err != nil {
		return nil, err
	}
	m.SysHash = hash

	m.DBHeight, newData = binary.BigEndian.Uint32(newData[0:4]), newData[4:]
	m.VMIndex, newData = int(newData[0]), newData[1:]

	header := directoryBlock.NewDBlockHeader()
	newData, err = header.UnmarshalBinaryData(newData)
	if err != nil {
		return nil, err
	}
	m.DirectoryBlockHeader = header

	hash = new(primitives.Hash)
	newData, err = hash.UnmarshalBinaryData(newData)
	if err != nil {
		return nil, err
	}
	m.ServerIdentityChainID = hash

	//if len(newData) > 0 {
	sig := new(primitives.Signature)
	newData, err = sig.UnmarshalBinaryData(newData)
	if err != nil {
		return nil, err
	}
	m.DBSignature = sig
	//}

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

func (m *DirectoryBlockSignature) UnmarshalBinary(data []byte) error {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesDirectoryBlockSignatureUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	_, err := m.UnmarshalBinaryData(data)
	return err
}

func (m *DirectoryBlockSignature) MarshalForSignature() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesDirectoryBlockSignatureMarshalForSignature.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if m.DirectoryBlockHeader == nil {
		m.DirectoryBlockHeader = directoryBlock.NewDBlockHeader()
	}

	var buf primitives.Buffer
	buf.Write([]byte{m.Type()})

	t := m.GetTimestamp()
	data, err := t.MarshalBinary()
	if err != nil {
		return nil, err
	}
	buf.Write(data)

	binary.Write(&buf, binary.BigEndian, m.SysHeight)
	if m.SysHash == nil {
		m.SysHash = primitives.NewZeroHash()
	}
	hash, err := m.SysHash.MarshalBinary()
	if err != nil {
		return nil, err
	}
	buf.Write(hash)

	binary.Write(&buf, binary.BigEndian, m.DBHeight)
	binary.Write(&buf, binary.BigEndian, byte(m.VMIndex))

	header, err := m.DirectoryBlockHeader.MarshalBinary()
	if err != nil {
		return nil, err
	}
	buf.Write(header)

	hash, err = m.ServerIdentityChainID.MarshalBinary()
	if err != nil {
		return nil, err
	}
	buf.Write(hash)

	if m.DBSignature != nil {
		dbSig, err := m.DBSignature.MarshalBinary()
		if err != nil {
			return nil, err
		}
		buf.Write(dbSig)
	} else {
		blankSig := make([]byte, constants.SIGNATURE_LENGTH)
		buf.Write(blankSig)
		buf.Write(blankSig[:32])
	}

	return buf.DeepCopyBytes(), nil
}

func (m *DirectoryBlockSignature) MarshalBinary() (data []byte, err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesDirectoryBlockSignatureMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	var sig interfaces.IFullSignature
	resp, err := m.MarshalForSignature()
	if err == nil {
		sig = m.GetSignature()
	}

	if sig != nil {
		sigBytes, err := sig.MarshalBinary()
		if err != nil {
			return resp, nil
		}
		return append(resp, sigBytes...), nil
	}
	return resp, nil
}

func (m *DirectoryBlockSignature) String() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesDirectoryBlockSignatureString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return fmt.Sprintf("%6s-VM%3d:          DBHt:%5d -- Signer[:3]=%x PrevDBKeyMR[:3]=%x hash[:3]=%x",
		"DBSig",
		m.VMIndex,
		m.DBHeight,
		m.ServerIdentityChainID.Bytes()[2:6],
		m.DirectoryBlockHeader.GetPrevKeyMR().Bytes()[:3],
		m.GetHash().Bytes()[:3])

}

func (e *DirectoryBlockSignature) JSONByte() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesDirectoryBlockSignatureJSONByte.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSON(e)
}

func (e *DirectoryBlockSignature) JSONString() (string, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesDirectoryBlockSignatureJSONString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSONString(e)
}
