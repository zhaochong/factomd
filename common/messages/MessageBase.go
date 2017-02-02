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

	NoResend bool // Don't resend this message if true.

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
	callTime := time.Now().UnixNano()
	defer messagesMessageBaseresend.Observe(float64(time.Now().UnixNano() - callTime))
	for i := 0; i < cnt; i++ {
		state.NetworkOutMsgQueue() <- msg
		time.Sleep(time.Duration(delay) * time.Second)
	}
}

func (m *MessageBase) SendOut(state interfaces.IState, msg interfaces.IMsg) {
	callTime := time.Now().UnixNano()
	defer messagesMessageBaseSendOut.Observe(float64(time.Now().UnixNano() - callTime))

	if m.NoResend {
		return
	}

	switch msg.(interface{}).(type) {
	//case ServerFault:
	//	go resend(state, msg, 20, 1)
	case FullServerFault:
		go resend(state, msg, 10, 2)
	case ServerFault:
		go resend(state, msg, 10, 2)
	case MissingMsg:
		go resend(state, msg, 1, 1)
	case DBStateMissing:
		go resend(state, msg, 1, 1)
	default:
		go resend(state, msg, 1, 1)
	}
}

func (m *MessageBase) GetNoResend() bool {
	callTime := time.Now().UnixNano()
	defer messagesMessageBaseGetNoResend.Observe(float64(time.Now().UnixNano() - callTime))
	return m.NoResend
}

func (m *MessageBase) SetNoResend(v bool) {
	callTime := time.Now().UnixNano()
	defer messagesMessageBaseSetNoResend.Observe(float64(time.Now().UnixNano() - callTime))
	m.NoResend = v
}

func (m *MessageBase) IsValid() bool {
	callTime := time.Now().UnixNano()
	defer messagesMessageBaseIsValid.Observe(float64(time.Now().UnixNano() - callTime))
	return m.Sigvalid
}

func (m *MessageBase) SetValid() {
	callTime := time.Now().UnixNano()
	defer messagesMessageBaseSetValid.Observe(float64(time.Now().UnixNano() - callTime))
	m.Sigvalid = true
}

// To suppress how many messages are sent to the NetworkInvalid Queue, we mark them, and only
// send them once.
func (m *MessageBase) MarkSentInvalid(b bool) {
	callTime := time.Now().UnixNano()
	defer messagesMessageBaseMarkSentInvalid.Observe(float64(time.Now().UnixNano() - callTime))
	m.MarkInvalid = b
}

func (m *MessageBase) SentInvlaid() bool {
	callTime := time.Now().UnixNano()
	defer messagesMessageBaseSentInvlaid.Observe(float64(time.Now().UnixNano() - callTime))
	return m.MarkInvalid
}

// Try and Resend.  Return true if we should keep the message, false if we should give up.
func (m *MessageBase) Resend(s interfaces.IState) (rtn bool) {
	callTime := time.Now().UnixNano()
	defer messagesMessageBaseResend.Observe(float64(time.Now().UnixNano() - callTime))
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
	callTime := time.Now().UnixNano()
	defer messagesMessageBaseExpire.Observe(float64(time.Now().UnixNano() - callTime))
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
	callTime := time.Now().UnixNano()
	defer messagesMessageBaseIsStalled.Observe(float64(time.Now().UnixNano() - callTime))
	return m.Stalled
}
func (m *MessageBase) SetStall(b bool) {
	callTime := time.Now().UnixNano()
	defer messagesMessageBaseSetStall.Observe(float64(time.Now().UnixNano() - callTime))
	m.Stalled = b
}

func (m *MessageBase) GetFullMsgHash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer messagesMessageBaseGetFullMsgHash.Observe(float64(time.Now().UnixNano() - callTime))
	if m.FullMsgHash == nil {
		m.FullMsgHash = primitives.NewZeroHash()
	}
	return m.FullMsgHash
}

func (m *MessageBase) SetFullMsgHash(hash interfaces.IHash) {
	callTime := time.Now().UnixNano()
	defer messagesMessageBaseSetFullMsgHash.Observe(float64(time.Now().UnixNano() - callTime))
	m.GetFullMsgHash().SetBytes(hash.Bytes())
}

func (m *MessageBase) GetOrigin() int {
	callTime := time.Now().UnixNano()
	defer messagesMessageBaseGetOrigin.Observe(float64(time.Now().UnixNano() - callTime))
	return m.Origin
}

func (m *MessageBase) SetOrigin(o int) {
	callTime := time.Now().UnixNano()
	defer messagesMessageBaseSetOrigin.Observe(float64(time.Now().UnixNano() - callTime))
	m.Origin = o
}

func (m *MessageBase) GetNetworkOrigin() string {
	callTime := time.Now().UnixNano()
	defer messagesMessageBaseGetNetworkOrigin.Observe(float64(time.Now().UnixNano() - callTime))
	return m.NetworkOrigin
}

func (m *MessageBase) SetNetworkOrigin(o string) {
	callTime := time.Now().UnixNano()
	defer messagesMessageBaseSetNetworkOrigin.Observe(float64(time.Now().UnixNano() - callTime))
	m.NetworkOrigin = o
}

// Returns true if this is a response to a peer to peer
// request.
func (m *MessageBase) IsPeer2Peer() bool {
	callTime := time.Now().UnixNano()
	defer messagesMessageBaseIsPeer2Peer.Observe(float64(time.Now().UnixNano() - callTime))
	return m.Peer2Peer
}

func (m *MessageBase) SetPeer2Peer(f bool) {
	callTime := time.Now().UnixNano()
	defer messagesMessageBaseSetPeer2Peer.Observe(float64(time.Now().UnixNano() - callTime))
	m.Peer2Peer = f
}

func (m *MessageBase) IsLocal() bool {
	callTime := time.Now().UnixNano()
	defer messagesMessageBaseIsLocal.Observe(float64(time.Now().UnixNano() - callTime))
	return m.LocalOnly
}

func (m *MessageBase) SetLocal(v bool) {
	callTime := time.Now().UnixNano()
	defer messagesMessageBaseSetLocal.Observe(float64(time.Now().UnixNano() - callTime))
	m.LocalOnly = v
}

func (m *MessageBase) GetLeaderChainID() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer messagesMessageBaseGetLeaderChainID.Observe(float64(time.Now().UnixNano() - callTime))
	if m.LeaderChainID == nil {
		m.LeaderChainID = primitives.NewZeroHash()
	}
	return m.LeaderChainID
}

func (m *MessageBase) SetLeaderChainID(hash interfaces.IHash) {
	callTime := time.Now().UnixNano()
	defer messagesMessageBaseSetLeaderChainID.Observe(float64(time.Now().UnixNano() - callTime))
	m.LeaderChainID = hash
}

func (m *MessageBase) GetVMIndex() (index int) {
	callTime := time.Now().UnixNano()
	defer messagesMessageBaseGetVMIndex.Observe(float64(time.Now().UnixNano() - callTime))
	index = m.VMIndex
	return
}

func (m *MessageBase) SetVMIndex(index int) {
	callTime := time.Now().UnixNano()
	defer messagesMessageBase.Observe(float64(time.Now().UnixNano() - callTime))
	m.VMIndex = index
}

func (m *MessageBase) GetVMHash() []byte {
	callTime := time.Now().UnixNano()
	defer messagesMessageBase.Observe(float64(time.Now().UnixNano() - callTime))
	return m.VMHash
}

func (m *MessageBase) SetVMHash(vmhash []byte) {
	callTime := time.Now().UnixNano()
	defer messagesMessageBase.Observe(float64(time.Now().UnixNano() - callTime))
	m.VMHash = vmhash
}

func (m *MessageBase) GetMinute() byte {
	callTime := time.Now().UnixNano()
	defer messagesMessageBaseGetMinute.Observe(float64(time.Now().UnixNano() - callTime))
	return m.Minute
}

func (m *MessageBase) SetMinute(minute byte) {
	callTime := time.Now().UnixNano()
	defer messagesMessageBaseSetMinute.Observe(float64(time.Now().UnixNano() - callTime))
	m.Minute = minute
}
