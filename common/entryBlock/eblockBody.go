// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package entryBlock

import (
	"bytes"
	"fmt"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
)

// EBlockBody is the series of Hashes that form the Entry Block Body.
type EBlockBody struct {
	EBEntries []interfaces.IHash
}

var _ interfaces.Printable = (*EBlockBody)(nil)
var _ interfaces.IEBlockBody = (*EBlockBody)(nil)

// NewEBlockBody initalizes an empty Entry Block Body.
func NewEBlockBody() *EBlockBody {
	callTime := time.Now().UnixNano()
	defer eBlockBodyNewEBlockBody.Observe(float64(time.Now().UnixNano() - callTime))	
		e := new(EBlockBody)
	e.EBEntries = make([]interfaces.IHash, 0)
	return e
}

// MR calculates the Merkle Root of the Entry Block Body. See func
// primitives.BuildMerkleTreeStore(hashes []interfaces.IHash) (merkles []interfaces.IHash) in common/merkle.go.
func (e *EBlockBody) MR() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer eBlockBodyMR.Observe(float64(time.Now().UnixNano() - callTime))	
	mrs := primitives.BuildMerkleTreeStore(e.EBEntries)
	r := mrs[len(mrs)-1]
	return r
}

func (e *EBlockBody) JSONByte() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer eBlockBodyJSONByte.Observe(float64(time.Now().UnixNano() - callTime))	
	return primitives.EncodeJSON(e)
}

func (e *EBlockBody) JSONString() (string, error) {
	callTime := time.Now().UnixNano()
	defer eBlockBodyJSONString.Observe(float64(time.Now().UnixNano() - callTime))	
	return primitives.EncodeJSONString(e)
}

func (e *EBlockBody) JSONBuffer(b *bytes.Buffer) error {
	callTime := time.Now().UnixNano()
	defer eBlockBodyJSONBuffer.Observe(float64(time.Now().UnixNano() - callTime))	
	return primitives.EncodeJSONToBuffer(e, b)
}

func (e *EBlockBody) String() string {
	callTime := time.Now().UnixNano()
	defer eBlockBodyString.Observe(float64(time.Now().UnixNano() - callTime))	
	var out primitives.Buffer
	for _, eh := range e.EBEntries {
		out.WriteString(fmt.Sprintf("    %20s: %x\n", "Entry Hash", eh.Bytes()[:3]))
	}
	return (string)(out.DeepCopyBytes())
}

func (e *EBlockBody) GetEBEntries() []interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer eBlockBodyGetEBEntries.Observe(float64(time.Now().UnixNano() - callTime))	
	return e.EBEntries[:]
}
