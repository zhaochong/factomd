package adminBlock

import (
	"fmt"

	"github.com/FactomProject/factomd/common/constants"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
)

// DB Signature Entry -------------------------
type IncreaseServerCount struct {
	Amount byte
}

var _ interfaces.IABEntry = (*IncreaseServerCount)(nil)
var _ interfaces.BinaryMarshallable = (*IncreaseServerCount)(nil)

// Create a new DB Signature Entry
func NewIncreaseServerCount(num byte) (e *IncreaseServerCount) {
	callTime := time.Now().UnixNano()
	defer entryIncreaseServerCountNewentryIncreaseServerCount.Observe(float64(time.Now().UnixNano() - callTime))	
	e = new(IncreaseServerCount)
	e.Amount = num
	return
}

func (c *IncreaseServerCount) UpdateState(state interfaces.IState) error {
	return nil
}

func (e *IncreaseServerCount) Type() byte {
	callTime := time.Now().UnixNano()
	defer entryIncreaseServerCountType.Observe(float64(time.Now().UnixNano() - callTime))	
	return constants.TYPE_ADD_SERVER_COUNT
}

func (e *IncreaseServerCount) MarshalBinary() (data []byte, err error) {
	callTime := time.Now().UnixNano()
	defer entryIncreaseServerCountMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))	
	var buf primitives.Buffer

	buf.Write([]byte{e.Type()})
	buf.Write([]byte{e.Amount})

	return buf.DeepCopyBytes(), nil
}

func (e *IncreaseServerCount) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	callTime := time.Now().UnixNano()
	defer entryIncreaseServerCountUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))	
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error unmarshalling Entry Increase Server Count: %v", r)
		}
	}()

	newData = data
	if newData[0] != e.Type() {
		return nil, fmt.Errorf("Invalid Entry type")
	}
	newData = newData[1:]

	e.Amount, newData = newData[0], newData[1:]

	return
}

func (e *IncreaseServerCount) UnmarshalBinary(data []byte) (err error) {
	callTime := time.Now().UnixNano()
	defer entryIncreaseServerCountUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))	
	_, err = e.UnmarshalBinaryData(data)
	return
}

func (e *IncreaseServerCount) JSONByte() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer entryIncreaseServerCountJSONByte.Observe(float64(time.Now().UnixNano() - callTime))	
	return primitives.EncodeJSON(e)
}

func (e *IncreaseServerCount) JSONString() (string, error) {
	callTime := time.Now().UnixNano()
	defer entryIncreaseServerCountJSONString.Observe(float64(time.Now().UnixNano() - callTime))	
	return primitives.EncodeJSONString(e)
}

func (e *IncreaseServerCount) String() string {
	callTime := time.Now().UnixNano()
	defer entryIncreaseServerCountString.Observe(float64(time.Now().UnixNano() - callTime))	
	str := fmt.Sprintf("    E: %35s -- by %d", "Increase Server Count", e.Amount)
	return str
}

func (e *IncreaseServerCount) IsInterpretable() bool {
	return false
}

func (e *IncreaseServerCount) Interpret() string {
	return ""
}

func (e *IncreaseServerCount) Hash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer entryIncreaseServerCountHash.Observe(float64(time.Now().UnixNano() - callTime))	
	bin, err := e.MarshalBinary()
	if err != nil {
		panic(err)
	}
	return primitives.Sha(bin)
}
