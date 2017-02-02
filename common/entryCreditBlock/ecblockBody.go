// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package entryCreditBlock

import (
	"fmt"

	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
)

type ECBlockBody struct {
	Entries []interfaces.IECBlockEntry
}

var _ = fmt.Print
var _ interfaces.Printable = (*ECBlockBody)(nil)
var _ interfaces.IECBlockBody = (*ECBlockBody)(nil)

func (e *ECBlockBody) GetEntries() []interfaces.IECBlockEntry {
	callTime := time.Now().UnixNano()
	defer entryCreditBlockecblockBodyGetEntries.Observe(float64(time.Now().UnixNano() - callTime))	
	return e.Entries
}

func (e *ECBlockBody) AddEntry(entry interfaces.IECBlockEntry) {
	callTime := time.Now().UnixNano()
	defer entryCreditBlockecblockBodyAddEntry.Observe(float64(time.Now().UnixNano() - callTime))	
	e.Entries = append(e.Entries, entry)
}

func (e *ECBlockBody) SetEntries(entries []interfaces.IECBlockEntry) {
	callTime := time.Now().UnixNano()
	defer entryCreditBlockecblockBodySetEntries.Observe(float64(time.Now().UnixNano() - callTime))	
	e.Entries = entries
}

func (e *ECBlockBody) JSONByte() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer entryCreditBlockecblockBodyJSONByte.Observe(float64(time.Now().UnixNano() - callTime))	
	return primitives.EncodeJSON(e)
}

func (e *ECBlockBody) JSONString() (string, error) {
	callTime := time.Now().UnixNano()
	defer entryCreditBlockecblockBodyJSONString.Observe(float64(time.Now().UnixNano() - callTime))	
	return primitives.EncodeJSONString(e)
}

func (e *ECBlockBody) String() string {
	callTime := time.Now().UnixNano()
	defer entryCreditBlockecblockBodyString.Observe(float64(time.Now().UnixNano() - callTime))	
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
	callTime := time.Now().UnixNano()
	defer entryCreditBlockecblockBodyNewECBlockBody.Observe(float64(time.Now().UnixNano() - callTime))	
	b := new(ECBlockBody)
	b.Entries = make([]interfaces.IECBlockEntry, 0)
	return b
}
