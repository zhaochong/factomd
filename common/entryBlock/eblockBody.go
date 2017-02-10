// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package entryBlock

import (
	"encoding/hex"
	"fmt"

	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
	"time"
)

// EBlockBody is the series of Hashes that form the Entry Block Body.
type EBlockBody struct {
	EBEntries []interfaces.IHash
}

var _ interfaces.Printable = (*EBlockBody)(nil)
var _ interfaces.IEBlockBody = (*EBlockBody)(nil)

// NewEBlockBody initalizes an empty Entry Block Body.
func NewEBlockBody() *EBlockBody {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryBlockNewEBlockBody.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	e := new(EBlockBody)
	e.EBEntries = make([]interfaces.IHash, 0)
	return e
}

// MR calculates the Merkle Root of the Entry Block Body. See func
// primitives.BuildMerkleTreeStore(hashes []interfaces.IHash) (merkles []interfaces.IHash) in common/merkle.go.
func (e *EBlockBody) MR() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryBlockEBlockBodyMR.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	mrs := primitives.BuildMerkleTreeStore(e.EBEntries)
	r := mrs[len(mrs)-1]
	return r
}

func (e *EBlockBody) JSONByte() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryBlockEBlockBodyJSONByte.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSON(e)
}

func (e *EBlockBody) JSONString() (string, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryBlockEBlockBodyJSONString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSONString(e)
}

func (e *EBlockBody) String() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryBlockEBlockBodyString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	var out primitives.Buffer
	for _, eh := range e.EBEntries {
		out.WriteString(fmt.Sprintf("    %20s: %x\n", "Entry Hash", eh.Bytes()[:3]))
	}
	return (string)(out.DeepCopyBytes())
}

func (e *EBlockBody) GetEBEntries() []interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryBlockEBlockBodyGetEBEntries.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return e.EBEntries[:]
}

// AddEBEntry creates a new Entry Block Entry from the provided Factom Entry
// and adds it to the Entry Block Body.
func (e *EBlockBody) AddEBEntry(entry interfaces.IHash) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryBlockEBlockBodyAddEBEntry.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	e.EBEntries = append(e.EBEntries, entry)
}

// AddEndOfMinuteMarker adds the End of Minute to the Entry Block. The End of
// Minut byte becomes the last byte in a 32 byte slice that is added to the
// Entry Block Body as an Entry Block Entry.
func (e *EBlockBody) AddEndOfMinuteMarker(m byte) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryBlockEBlockBodyAddEndOfMinuteMarker.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	// create a map of possible minute markers that may be found in the
	// EBlock Body
	mins := make(map[string]uint8)
	for i := byte(1); i <= 10; i++ {
		h := make([]byte, 32)
		h[len(h)-1] = i
		mins[hex.EncodeToString(h)] = i
	}

	// check if the previous entry is a minute marker and return without
	// writing if it is
	prevEntry := e.EBEntries[len(e.EBEntries)-1]
	if _, exist := mins[prevEntry.String()]; exist {
		return
	}

	h := make([]byte, 32)
	h[len(h)-1] = m
	hash := primitives.NewZeroHash()
	hash.SetBytes(h)

	e.AddEBEntry(hash)
}
