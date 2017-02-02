package adminBlock

import (
	"bytes"
	"fmt"

	"github.com/FactomProject/factomd/common/constants"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
)

type EndOfMinuteEntry struct {
	MinuteNumber byte
}

var _ interfaces.Printable = (*EndOfMinuteEntry)(nil)
var _ interfaces.BinaryMarshallable = (*EndOfMinuteEntry)(nil)
var _ interfaces.IABEntry = (*EndOfMinuteEntry)(nil)

func (m *EndOfMinuteEntry) Type() byte {
	callTime := time.Now().UnixNano()
	defer entryEndOfMinuteType.Observe(float64(time.Now().UnixNano() - callTime))	
	return constants.TYPE_MINUTE_NUM
}

func (c *EndOfMinuteEntry) UpdateState(state interfaces.IState) error {
	callTime := time.Now().UnixNano()
	defer entryEndOfMinuteUpdateState.Observe(float64(time.Now().UnixNano() - callTime))	

	return nil
}

func NewEndOfMinuteEntry(minuteNumber byte) *EndOfMinuteEntry {
	callTime := time.Now().UnixNano()
	defer entryEndOfMinuteNewentryEndOfMinute.Observe(float64(time.Now().UnixNano() - callTime))	
	e := new(EndOfMinuteEntry)
	e.MinuteNumber = minuteNumber
	return e
}

func (e *EndOfMinuteEntry) MarshalBinary() (data []byte, err error) {
	callTime := time.Now().UnixNano()
	defer entryEndOfMinuteMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))	
	var buf primitives.Buffer

	buf.Write([]byte{e.Type()})
	buf.Write([]byte{e.MinuteNumber})

	return buf.DeepCopyBytes(), nil
}

func (e *EndOfMinuteEntry) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	callTime := time.Now().UnixNano()
	defer entryEndOfMinuteUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))	
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error unmarshalling End Of Minute: %v", r)
		}
	}()
	newData = data
	if newData[0] != e.Type() {
		return nil, fmt.Errorf("Invalid Entry type")
	}

	newData = newData[1:]
	e.MinuteNumber, newData = newData[0], newData[1:]

	return
}

func (e *EndOfMinuteEntry) UnmarshalBinary(data []byte) (err error) {
	callTime := time.Now().UnixNano()
	defer entryEndOfMinuteUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))	
	_, err = e.UnmarshalBinaryData(data)
	return
}

func (e *EndOfMinuteEntry) JSONByte() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer entryEndOfMinuteJSONByte.Observe(float64(time.Now().UnixNano() - callTime))	
	return primitives.EncodeJSON(e)
}

func (e *EndOfMinuteEntry) JSONString() (string, error) {
	callTime := time.Now().UnixNano()
	defer entryEndOfMinuteJSONString.Observe(float64(time.Now().UnixNano() - callTime))	
	return primitives.EncodeJSONString(e)
}

func (e *EndOfMinuteEntry) JSONBuffer(b *bytes.Buffer) error {
	callTime := time.Now().UnixNano()
	defer entryEndOfMinuteJSONBuffer.Observe(float64(time.Now().UnixNano() - callTime))	
	return primitives.EncodeJSONToBuffer(e, b)
}

func (e *EndOfMinuteEntry) String() string {
	callTime := time.Now().UnixNano()
	defer entryEndOfMinuteString.Observe(float64(time.Now().UnixNano() - callTime))	
	return fmt.Sprintf("    E: %35s -- %17s %d",
		"EndOfMinuteEntry",
		"Minute", e.MinuteNumber)
}

func (e *EndOfMinuteEntry) IsInterpretable() bool {
	return true
}

func (e *EndOfMinuteEntry) Interpret() string {
	return fmt.Sprintf("End of Minute %v", e.MinuteNumber)
}

func (e *EndOfMinuteEntry) Hash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer entryEndOfMinuteHash.Observe(float64(time.Now().UnixNano() - callTime))	
	bin, err := e.MarshalBinary()
	if err != nil {
		panic(err)
	}
	return primitives.Sha(bin)
}
