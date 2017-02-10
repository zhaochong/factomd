// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package interfaces

import (
	"fmt"
	"time"
)

// This object will hold the public keys for servers that are not
// us, and maybe other information about servers.
type IFctServer interface {
	GetChainID() IHash
	String() string
	GetName() string
	IsOnline() bool
	SetOnline(bool)
	LeaderToReplace() IHash
	SetReplace(IHash)
}

type FctServer struct {
	ChainID IHash
	Name    string
	Online  bool
	Replace IHash
}

var _ IFctServer = (*FctServer)(nil)

func (s *FctServer) GetChainID() IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdinterfacesFctServerGetChainID.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return s.ChainID
}

func (s *FctServer) String() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdinterfacesFctServerString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return fmt.Sprintf("Server[:4]: %x", s.GetChainID().Bytes()[:10])
}

func (s *FctServer) GetName() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdinterfacesFctServerGetName.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return s.Name
}

func (s *FctServer) IsOnline() bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdinterfacesFctServerIsOnline.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return s.Online
}

func (s *FctServer) SetOnline(o bool) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdinterfacesFctServerSetOnline.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	s.Online = o
}

func (s *FctServer) LeaderToReplace() IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdinterfacesFctServerLeaderToReplace.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return s.Replace
}

func (s *FctServer) SetReplace(h IHash) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdinterfacesFctServerSetReplace.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	s.Replace = h
}
