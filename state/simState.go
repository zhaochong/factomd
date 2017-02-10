// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package state

// This file is for the simulator to attach identities properly to the state.
// Each state has its own set of keys that need to match the ones in the
// identitiy to properly test identities/authorities
import (
	"github.com/FactomProject/factomd/common/primitives"
	"time"
)

func (s *State) SimSetNewKeys(p *primitives.PrivateKey) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdstateStateSimSetNewKeys.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	s.serverPrivKey = p
	s.serverPubKey = p.Pub
}

func (s *State) SimGetSigKey() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdstateStateSimGetSigKey.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return s.serverPrivKey.Pub.String()
}
