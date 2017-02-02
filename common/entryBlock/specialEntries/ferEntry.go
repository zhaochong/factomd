package specialEntries

import (
	"encoding/json"
	"fmt"

	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
)

var _ = fmt.Print

type FEREntry struct {
	Version                string `json:"version"`
	ExpirationHeight       uint32 `json:"expiration_height"`
	ResidentHeight         uint32 `json:"resident_height"`
	TargetActivationHeight uint32 `json:"target_activation_height"`
	Priority               uint32 `json:"priority"`
	TargetPrice            uint64 `json:"target_price"`
}

var _ interfaces.Printable = (*FEREntry)(nil)
var _ interfaces.BinaryMarshallable = (*FEREntry)(nil)
var _ interfaces.IFEREntry = (*FEREntry)(nil)

// Getter Version
func (this *FEREntry) GetVersion() string {
	callTime := time.Now().UnixNano()
	defer specialEntriesGetVersion.Observe(float64(time.Now().UnixNano() - callTime))	
	return this.Version 
}

// Setter Version
func (this *FEREntry) SetVersion(passedVersion string) interfaces.IFEREntry {
	callTime := time.Now().UnixNano()
	defer specialEntriesSetVersion.Observe(float64(time.Now().UnixNano() - callTime))	
	this.Version = passedVersion
	return this
}

// Getter ExpirationHeight
func (this *FEREntry) GetExpirationHeight() uint32 {
	callTime := time.Now().UnixNano()
	defer specialEntriesGetExpirationHeight.Observe(float64(time.Now().UnixNano() - callTime))	
	return this.ExpirationHeight
}

// Setter ExpirationHeight
func (this *FEREntry) SetExpirationHeight(passedExpirationHeight uint32) interfaces.IFEREntry {
	callTime := time.Now().UnixNano()
	defer specialEntriesSetExpirationHeight.Observe(float64(time.Now().UnixNano() - callTime))	
	this.ExpirationHeight = passedExpirationHeight
	return this
}

// Getter ResidentHeight
func (this *FEREntry) GetResidentHeight() uint32 {
	callTime := time.Now().UnixNano()
	defer specialEntriesGetResidentHeight.Observe(float64(time.Now().UnixNano() - callTime))	
	return this.ResidentHeight
}

// Setter ResidentHeight
func (this *FEREntry) SetResidentHeight(passedResidentHeight uint32) interfaces.IFEREntry {
	callTime := time.Now().UnixNano()
	defer specialEntriesSetResidentHeight.Observe(float64(time.Now().UnixNano() - callTime))	
	this.ResidentHeight = passedResidentHeight
	return this
}

// Getter TargetActivationHeight
func (this *FEREntry) GetTargetActivationHeight() uint32 {
	callTime := time.Now().UnixNano()
	defer specialEntriesGetTargetActivationHeight.Observe(float64(time.Now().UnixNano() - callTime))	
	return this.TargetActivationHeight
}

// Setter TargetActivationHeight
func (this *FEREntry) SetTargetActivationHeight(passedTargetActivationHeight uint32) interfaces.IFEREntry {
	callTime := time.Now().UnixNano()
	defer specialEntriesSetTargetActivationHeight.Observe(float64(time.Now().UnixNano() - callTime))	
	this.TargetActivationHeight = passedTargetActivationHeight
	return this
}

// Getter Priority
func (this *FEREntry) GetPriority() uint32 {
	callTime := time.Now().UnixNano()
	defer specialEntriesGetPriority.Observe(float64(time.Now().UnixNano() - callTime))	
	return this.Priority
}

// Setter Priority
func (this *FEREntry) SetPriority(passedPriority uint32) interfaces.IFEREntry {
	callTime := time.Now().UnixNano()
	defer specialEntriesGetPriority.Observe(float64(time.Now().UnixNano() - callTime))	
	this.Priority = passedPriority
	return this
}

// Getter TargetPrice
func (this *FEREntry) GetTargetPrice() uint64 {
	callTime := time.Now().UnixNano()
	defer specialEntriesGetTargetPrice.Observe(float64(time.Now().UnixNano() - callTime))	
	return this.TargetPrice
}

// Setter TargetPrice
func (this *FEREntry) SetTargetPrice(passedTargetPrice uint64) interfaces.IFEREntry {
	callTime := time.Now().UnixNano()
	defer specialEntriesSetTargetPrice.Observe(float64(time.Now().UnixNano() - callTime))	
	this.TargetPrice = passedTargetPrice
	return this
}

func (e *FEREntry) JSONByte() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer specialEntriesJSONByte.Observe(float64(time.Now().UnixNano() - callTime))	
	return primitives.EncodeJSON(e)
}

func (e *FEREntry) JSONString() (string, error) {
	callTime := time.Now().UnixNano()
	defer specialEntriesJSONString.Observe(float64(time.Now().UnixNano() - callTime))	
	return primitives.EncodeJSONString(e)
}

func (e *FEREntry) String() string {
	callTime := time.Now().UnixNano()
	defer specialEntriesString.Observe(float64(time.Now().UnixNano() - callTime))	
	str, _ := e.JSONString()
	return str
}

func (e *FEREntry) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	callTime := time.Now().UnixNano()
	defer specialEntriesUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))	
	return nil, json.Unmarshal(data, e)
}

func (e *FEREntry) UnmarshalBinary(data []byte) (err error) {
	callTime := time.Now().UnixNano()
	defer specialEntriesUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))	
	_, err = e.UnmarshalBinaryData(data)
	return
}

func (e *FEREntry) MarshalBinary() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer specialEntriesMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))	
	return json.Marshal(e)
}
