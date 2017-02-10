// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package messages

import (
	"time"

	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
)

type MessageBase struct {
	FullMsgHash interfaces.IHash

	Origin        int    // Set and examined on a server, not marshaled with the message
	NetworkOrigin string // Hash of the network peer/connection where the message is from
	Peer2Peer     bool   // The nature of this message type, not marshaled with the message
	LocalOnly     bool   // This message is only a local message, is not broadcasted and may skip verification

	NoResend  bool // Don't resend this message if true.
	ResendCnt int  // Put a limit on resends

	LeaderChainID interfaces.IHash
	MsgHash       interfaces.IHash // Cache of the hash of a message
	VMIndex       int              // The Index of the VM responsible for this message.
	VMHash        []byte           // Basis for selecting a VMIndex
	Minute        byte
	resend        int64 // Time to resend (milliseconds)
	expire        int64 // Time to expire (milliseconds)

	Stalled     bool // This message is currently stalled
	MarkInvalid bool
	Sigvalid    bool
}

func resend(state interfaces.IState, msg interfaces.IMsg, cnt int, delay int) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesresend.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	for i := 0; i < cnt; i++ {
		state.NetworkOutMsgQueue() <- msg
		time.Sleep(time.Duration(delay) * time.Second)
	}
}

func (m *MessageBase) SendOut(state interfaces.IState, msg interfaces.IMsg) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMessageBaseSendOut.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	// Dont' resend if we are behind
	if m.ResendCnt > 1 && state.GetHighestKnownBlock()-state.GetHighestSavedBlk() > 4 {
		return
	}
	if m.NoResend {
		return
	}

	if m.ResendCnt > 4 {
		return
	}
	m.ResendCnt++

	switch msg.(interface{}).(type) {
	//case ServerFault:
	//	go resend(state, msg, 20, 1)
	case FullServerFault:
		go resend(state, msg, 2, 5)
	case ServerFault:
		go resend(state, msg, 2, 5)
	case MissingMsg:
		go resend(state, msg, 1, 1)
	case DBStateMissing:
		go resend(state, msg, 1, 1)
	default:
		go resend(state, msg, 1, 1)
	}
}

func (m *MessageBase) GetNoResend() bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMessageBaseGetNoResend.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return m.NoResend
}

func (m *MessageBase) SetNoResend(v bool) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMessageBaseSetNoResend.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	m.NoResend = v
}

func (m *MessageBase) IsValid() bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMessageBaseIsValid.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return m.Sigvalid
}

func (m *MessageBase) SetValid() {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMessageBaseSetValid.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	m.Sigvalid = true
}

// To suppress how many messages are sent to the NetworkInvalid Queue, we mark them, and only
// send them once.
func (m *MessageBase) MarkSentInvalid(b bool) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMessageBaseMarkSentInvalid.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	m.MarkInvalid = b
}

func (m *MessageBase) SentInvlaid() bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMessageBaseSentInvlaid.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return m.MarkInvalid
}

// Try and Resend.  Return true if we should keep the message, false if we should give up.
func (m *MessageBase) Resend(s interfaces.IState) (rtn bool) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMessageBaseResend.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	now := s.GetTimestamp().GetTimeMilli()
	if m.resend == 0 {
		m.resend = now
		return false
	}
	if now-m.resend > 20000 && len(s.NetworkOutMsgQueue()) < 1000 {
		m.resend = now
		return true
	}
	return false
}

// Try and Resend.  Return true if we should keep the message, false if we should give up.
func (m *MessageBase) Expire(s interfaces.IState) (rtn bool) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMessageBaseExpire.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	now := s.GetTimestamp().GetTimeMilli()
	if m.expire == 0 {
		m.expire = now
	}
	if now-m.expire > 5*60*1000 { // Keep messages for some length before giving up.
		rtn = true
	}
	return
}

func (m *MessageBase) IsStalled() bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMessageBaseIsStalled.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return m.Stalled
}
func (m *MessageBase) SetStall(b bool) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMessageBaseSetStall.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	m.Stalled = b
}

func (m *MessageBase) GetFullMsgHash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMessageBaseGetFullMsgHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if m.FullMsgHash == nil {
		m.FullMsgHash = primitives.NewZeroHash()
	}
	return m.FullMsgHash
}

func (m *MessageBase) SetFullMsgHash(hash interfaces.IHash) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMessageBaseSetFullMsgHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	m.GetFullMsgHash().SetBytes(hash.Bytes())
}

func (m *MessageBase) GetOrigin() int {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMessageBaseGetOrigin.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return m.Origin
}

func (m *MessageBase) SetOrigin(o int) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMessageBaseSetOrigin.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	m.Origin = o
}

func (m *MessageBase) GetNetworkOrigin() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMessageBaseGetNetworkOrigin.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return m.NetworkOrigin
}

func (m *MessageBase) SetNetworkOrigin(o string) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMessageBaseSetNetworkOrigin.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	m.NetworkOrigin = o
}

// Returns true if this is a response to a peer to peer
// request.
func (m *MessageBase) IsPeer2Peer() bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMessageBaseIsPeer2Peer.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return m.Peer2Peer
}

func (m *MessageBase) SetPeer2Peer(f bool) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMessageBaseSetPeer2Peer.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	m.Peer2Peer = f
}

func (m *MessageBase) IsLocal() bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMessageBaseIsLocal.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return m.LocalOnly
}

func (m *MessageBase) SetLocal(v bool) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMessageBaseSetLocal.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	m.LocalOnly = v
}

func (m *MessageBase) GetLeaderChainID() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMessageBaseGetLeaderChainID.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if m.LeaderChainID == nil {
		m.LeaderChainID = primitives.NewZeroHash()
	}
	return m.LeaderChainID
}

func (m *MessageBase) SetLeaderChainID(hash interfaces.IHash) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMessageBaseSetLeaderChainID.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	m.LeaderChainID = hash
}

func (m *MessageBase) GetVMIndex() (index int) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMessageBaseGetVMIndex.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	index = m.VMIndex
	return
}

func (m *MessageBase) SetVMIndex(index int) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMessageBaseSetVMIndex.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	m.VMIndex = index
}

func (m *MessageBase) GetVMHash() []byte {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMessageBaseGetVMHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return m.VMHash
}

func (m *MessageBase) SetVMHash(vmhash []byte) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMessageBaseSetVMHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	m.VMHash = vmhash
}

func (m *MessageBase) GetMinute() byte {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMessageBaseGetMinute.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return m.Minute
}

func (m *MessageBase) SetMinute(minute byte) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmessagesMessageBaseSetMinute.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	m.Minute = minute
}
