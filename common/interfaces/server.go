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
type IServer interface {
	GetChainID() IHash
	String() string
	GetName() string
	IsOnline() bool
	SetOnline(bool)
	LeaderToReplace() IHash
	SetReplace(IHash)
}

type Server struct {
	ChainID IHash
	Name    string
	Online  bool
	Replace IHash
}

var _ IServer = (*Server)(nil)

func (s *Server) GetName() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdinterfacesServerGetName.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return s.Name
}

func (s *Server) GetChainID() IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdinterfacesServerGetChainID.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return s.ChainID
}

func (s *Server) String() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdinterfacesServerString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return fmt.Sprintf("%s %s", "Server:", s.GetChainID().String())
}

func (s *Server) IsOnline() bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdinterfacesServerIsOnline.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return s.Online
}

func (s *Server) SetOnline(o bool) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdinterfacesServerSetOnline.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	s.Online = o
}

func (s *Server) LeaderToReplace() IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdinterfacesServerLeaderToReplace.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return s.Replace
}

func (s *Server) SetReplace(h IHash) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdinterfacesServerSetReplace.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	s.Replace = h
}
