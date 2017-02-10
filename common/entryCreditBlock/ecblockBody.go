// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package entryCreditBlock

import (
	"fmt"

	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
	"time"
)

type ECBlockBody struct {
	Entries []interfaces.IECBlockEntry
}

var _ = fmt.Print
var _ interfaces.Printable = (*ECBlockBody)(nil)
var _ interfaces.IECBlockBody = (*ECBlockBody)(nil)

func (e *ECBlockBody) GetEntries() []interfaces.IECBlockEntry {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockECBlockBodyGetEntries.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return e.Entries
}

func (e *ECBlockBody) AddEntry(entry interfaces.IECBlockEntry) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockECBlockBodyAddEntry.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	e.Entries = append(e.Entries, entry)
}

func (e *ECBlockBody) SetEntries(entries []interfaces.IECBlockEntry) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockECBlockBodySetEntries.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	e.Entries = entries
}

func (e *ECBlockBody) JSONByte() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockECBlockBodyJSONByte.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSON(e)
}

func (e *ECBlockBody) JSONString() (string, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockECBlockBodyJSONString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSONString(e)
}

func (e *ECBlockBody) String() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockECBlockBodyString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	var out primitives.Buffer
	for _, v := range e.Entries {
		out.WriteString(v.String())
	}
	return string(out.DeepCopyBytes())
}

/*******************************************************
 * Support Functions
 *******************************************************/

func NewECBlockBody() interfaces.IECBlockBody {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockNewECBlockBody.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	b := new(ECBlockBody)
	b.Entries = make([]interfaces.IECBlockEntry, 0)
	return b
}
