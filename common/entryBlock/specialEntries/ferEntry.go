package specialEntries

import (
	"encoding/json"
	"fmt"

	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
	"time"
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
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdspecialEntriesFEREntryGetVersion.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return this.Version
}

// Setter Version
func (this *FEREntry) SetVersion(passedVersion string) interfaces.IFEREntry {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdspecialEntriesFEREntrySetVersion.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	this.Version = passedVersion
	return this
}

// Getter ExpirationHeight
func (this *FEREntry) GetExpirationHeight() uint32 {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdspecialEntriesFEREntryGetExpirationHeight.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return this.ExpirationHeight
}

// Setter ExpirationHeight
func (this *FEREntry) SetExpirationHeight(passedExpirationHeight uint32) interfaces.IFEREntry {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdspecialEntriesFEREntrySetExpirationHeight.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	this.ExpirationHeight = passedExpirationHeight
	return this
}

// Getter ResidentHeight
func (this *FEREntry) GetResidentHeight() uint32 {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdspecialEntriesFEREntryGetResidentHeight.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return this.ResidentHeight
}

// Setter ResidentHeight
func (this *FEREntry) SetResidentHeight(passedResidentHeight uint32) interfaces.IFEREntry {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdspecialEntriesFEREntrySetResidentHeight.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	this.ResidentHeight = passedResidentHeight
	return this
}

// Getter TargetActivationHeight
func (this *FEREntry) GetTargetActivationHeight() uint32 {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdspecialEntriesFEREntryGetTargetActivationHeight.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return this.TargetActivationHeight
}

// Setter TargetActivationHeight
func (this *FEREntry) SetTargetActivationHeight(passedTargetActivationHeight uint32) interfaces.IFEREntry {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdspecialEntriesFEREntrySetTargetActivationHeight.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	this.TargetActivationHeight = passedTargetActivationHeight
	return this
}

// Getter Priority
func (this *FEREntry) GetPriority() uint32 {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdspecialEntriesFEREntryGetPriority.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return this.Priority
}

// Setter Priority
func (this *FEREntry) SetPriority(passedPriority uint32) interfaces.IFEREntry {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdspecialEntriesFEREntrySetPriority.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	this.Priority = passedPriority
	return this
}

// Getter TargetPrice
func (this *FEREntry) GetTargetPrice() uint64 {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdspecialEntriesFEREntryGetTargetPrice.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return this.TargetPrice
}

// Setter TargetPrice
func (this *FEREntry) SetTargetPrice(passedTargetPrice uint64) interfaces.IFEREntry {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdspecialEntriesFEREntrySetTargetPrice.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	this.TargetPrice = passedTargetPrice
	return this
}

func (e *FEREntry) JSONByte() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdspecialEntriesFEREntryJSONByte.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSON(e)
}

func (e *FEREntry) JSONString() (string, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdspecialEntriesFEREntryJSONString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSONString(e)
}

func (e *FEREntry) String() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdspecialEntriesFEREntryString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	str, _ := e.JSONString()
	return str
}

func (e *FEREntry) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdspecialEntriesFEREntryUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return nil, json.Unmarshal(data, e)
}

func (e *FEREntry) UnmarshalBinary(data []byte) (err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdspecialEntriesFEREntryUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	_, err = e.UnmarshalBinaryData(data)
	return
}

func (e *FEREntry) MarshalBinary() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdspecialEntriesFEREntryMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return json.Marshal(e)
}
