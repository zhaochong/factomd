// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Transaction Address for a factoid transaction.   contains an amount
// and the address.  Our inputs spec how much is going into a transaction
// and our outputs spec how much is going out of a transaction.  This
// avoids having to have extra outputs to deal with change.
//

package factoid

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
	"github.com/FactomProject/factomd/common/primitives/random"
	"time"
)

type TransAddress struct {
	Amount  uint64              `json:"amount"`
	Address interfaces.IAddress `json:"address"`
	// Not marshalled
	UserAddress string `json:"useraddress"`
}

var _ interfaces.ITransAddress = (*TransAddress)(nil)

func RandomTransAddress() interfaces.ITransAddress {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidRandomTransAddress.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	ta := new(TransAddress)
	ta.Address = RandomAddress()
	ta.Amount = random.RandUInt64()
	ta.UserAddress = random.RandomString()
	return ta
}

func (t *TransAddress) SetUserAddress(v string) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidTransAddressSetUserAddress.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	t.UserAddress = v
}

func (t *TransAddress) GetUserAddress() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidTransAddressGetUserAddress.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return t.UserAddress
}

func (t *TransAddress) UnmarshalBinary(data []byte) error {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidTransAddressUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	_, err := t.UnmarshalBinaryData(data)
	return err
}

func (e *TransAddress) JSONByte() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidTransAddressJSONByte.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSON(e)
}

func (e *TransAddress) JSONString() (string, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidTransAddressJSONString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSONString(e)
}

func (t *TransAddress) String() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidTransAddressString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	str, _ := t.JSONString()
	return str
}

func (t *TransAddress) IsSameAs(add interfaces.ITransAddress) bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidTransAddressIsSameAs.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if t.GetAmount() != add.GetAmount() {
		return false
	}
	if t.GetAddress().IsSameAs(add.GetAddress()) == false {
		return false
	}
	return true
}

func (t *TransAddress) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidTransAddressUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if len(data) < 36 {
		return nil, fmt.Errorf("Data source too short to UnmarshalBinary() an address: %d", len(data))
	}

	t.Amount, data = primitives.DecodeVarInt(data)
	t.Address = new(Address)

	data, err = t.Address.UnmarshalBinaryData(data)

	return data, err
}

// MarshalBinary.  'nuff said
func (a TransAddress) MarshalBinary() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidTransAddressMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	var out primitives.Buffer

	err := primitives.EncodeVarInt(&out, a.Amount)
	if err != nil {
		return nil, err
	}
	data, err := a.Address.MarshalBinary()
	out.Write(data)

	return out.DeepCopyBytes(), err
}

// Accessor. Default to a zero length string.  This is a debug
// thing for looking out what we have built. Used by
// CustomMarshalText
func (ta TransAddress) GetName() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidTransAddressGetName.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return ""
}

// Accessor.  Get the amount with this address.
func (ta TransAddress) GetAmount() uint64 {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidTransAddressGetAmount.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return ta.Amount
}

// Accessor.  Get the amount with this address.
func (ta *TransAddress) SetAmount(amount uint64) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidTransAddressSetAmount.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	ta.Amount = amount
}

// Accessor.  Get the raw address.  Could be an actual address,
// or a hash of an authorization block.  See authorization.go
func (ta TransAddress) GetAddress() interfaces.IAddress {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidTransAddressGetAddress.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return ta.Address
}

// Accessor.  Get the raw address.  Could be an actual address,
// or a hash of an authorization block.  See authorization.go
func (ta *TransAddress) SetAddress(address interfaces.IAddress) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidTransAddressSetAddress.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	ta.Address = address
}

// Make this into somewhat readable text.
func (ta TransAddress) CustomMarshalTextAll(fct bool, label string) ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidTransAddressCustomMarshalTextAll.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	var out primitives.Buffer
	out.WriteString(fmt.Sprintf("   %8s:", label))
	v := primitives.ConvertDecimalToPaddedString(ta.Amount)
	fill := 8 - len(v) + strings.Index(v, ".") + 1
	fstr := fmt.Sprintf("%%%vs%%%vs ", 18-fill, fill)
	out.WriteString(fmt.Sprintf(fstr, v, ""))
	if fct {
		out.WriteString(primitives.ConvertFctAddressToUserStr(ta.Address))
	} else {
		out.WriteString(primitives.ConvertECAddressToUserStr(ta.Address))
	}
	str := fmt.Sprintf("\n                  %016x %038s\n\n", ta.Amount, string(hex.EncodeToString(ta.GetAddress().Bytes())))
	out.WriteString(str)
	return out.DeepCopyBytes(), nil
}

func (ta TransAddress) CustomMarshalText2(label string) ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidTransAddressCustomMarshalText2.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return ta.CustomMarshalTextAll(true, label)
}

func (ta TransAddress) CustomMarshalTextEC2(label string) ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidTransAddressCustomMarshalTextEC2.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return ta.CustomMarshalTextAll(false, label)
}

func (ta TransAddress) CustomMarshalTextInput() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidTransAddressCustomMarshalTextInput.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return ta.CustomMarshalText2("input")
}

func (ta TransAddress) StringInput() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidTransAddressStringInput.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	b, _ := ta.CustomMarshalTextInput()
	return string(b)
}

func (ta TransAddress) CustomMarshalTextOutput() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidTransAddressCustomMarshalTextOutput.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return ta.CustomMarshalText2("output")
}

func (ta TransAddress) StringOutput() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidTransAddressStringOutput.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	b, _ := ta.CustomMarshalTextOutput()
	return string(b)
}

func (ta TransAddress) CustomMarshalTextECOutput() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidTransAddressCustomMarshalTextECOutput.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return ta.CustomMarshalTextEC2("ecoutput")
}

func (ta TransAddress) StringECOutput() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidTransAddressStringECOutput.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	b, _ := ta.CustomMarshalTextECOutput()
	return string(b)
}

/******************************
 * Helper functions
 ******************************/

func NewOutECAddress(address interfaces.IAddress, amount uint64) interfaces.ITransAddress {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidNewOutECAddress.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	ta := new(TransAddress)
	ta.Amount = amount
	ta.Address = address
	ta.UserAddress = primitives.ConvertECAddressToUserStr(address)
	return ta
}

func NewOutAddress(address interfaces.IAddress, amount uint64) interfaces.ITransAddress {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidNewOutAddress.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	ta := new(TransAddress)
	ta.Amount = amount
	ta.Address = address
	ta.UserAddress = primitives.ConvertFctAddressToUserStr(address)
	return ta
}

func NewInAddress(address interfaces.IAddress, amount uint64) interfaces.ITransAddress {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidNewInAddress.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	ta := new(TransAddress)
	ta.Amount = amount
	ta.Address = address
	//  at this point we know this address is an EC address.
	//  so fill useraddress with a factoid formatted human readable address
	ta.UserAddress = primitives.ConvertFctAddressToUserStr(address)
	return ta
}
