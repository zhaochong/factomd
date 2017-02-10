// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package wsapi

import (
	"github.com/FactomProject/factomd/common/primitives"
	"time"
)

/*
The error codes from and including -32768 to -32000 are reserved for pre-defined errors.
Any code within this range, but not defined explicitly below is reserved for future use.
The error codes are nearly the same as those suggested for XML-RPC at the following url:
http://xmlrpc-epi.sourceforge.net/specs/rfc.fault_codes.php

code				message						meaning
-32700				Parse error					Invalid JSON was received by the server.
												An error occurred on the server while parsing the JSON text.
-32600				Invalid Request				The JSON sent is not a valid Request object.
-32601				Method not found			The method does not exist / is not available.
-32602				Invalid params				Invalid method parameter(s).
-32603				Internal error				Internal JSON-RPC error.
-32000 to -32099	Server error				Reserved for implementation-defined server-errors.
*/

func NewParseError() *primitives.JSONError {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdwsapiNewParseError.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.NewJSONError(-32700, "Parse error", nil)
}
func NewInvalidRequestError() *primitives.JSONError {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdwsapiNewInvalidRequestError.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.NewJSONError(-32600, "Invalid Request", nil)
}
func NewMethodNotFoundError() *primitives.JSONError {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdwsapiNewMethodNotFoundError.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.NewJSONError(-32601, "Method not found", nil)
}
func NewInvalidParamsError() *primitives.JSONError {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdwsapiNewInvalidParamsError.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.NewJSONError(-32602, "Invalid params", nil)
}
func NewInternalError() *primitives.JSONError {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdwsapiNewInternalError.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.NewJSONError(-32603, "Internal error", nil)
}

/*******************************************************************/
func NewCustomInternalError(data interface{}) *primitives.JSONError {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdwsapiNewCustomInternalError.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.NewJSONError(-32603, "Internal error", data)
}
func NewCustomInvalidParamsError(data interface{}) *primitives.JSONError {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdwsapiNewCustomInvalidParamsError.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.NewJSONError(-32602, "Invalid params", data)
}

/*******************************************************************/

func NewInvalidAddressError() *primitives.JSONError {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdwsapiNewInvalidAddressError.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.NewJSONError(-32602, "Invalid params", "Invalid Address")
}
func NewUnableToDecodeTransactionError() *primitives.JSONError {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdwsapiNewUnableToDecodeTransactionError.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.NewJSONError(-32602, "Invalid params", "Unable to decode the transaction")
}
func NewInvalidTransactionError() *primitives.JSONError {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdwsapiNewInvalidTransactionError.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.NewJSONError(-32602, "Invalid params", "Invalid Transaction")
}
func NewInvalidHashError() *primitives.JSONError {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdwsapiNewInvalidHashError.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.NewJSONError(-32602, "Invalid params", "Invalid Hash")
}
func NewInvalidEntryError() *primitives.JSONError {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdwsapiNewInvalidEntryError.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.NewJSONError(-32602, "Invalid params", "Invalid Entry")
}
func NewInvalidCommitChainError() *primitives.JSONError {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdwsapiNewInvalidCommitChainError.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.NewJSONError(-32602, "Invalid params", "Invalid Commit Chain")
}
func NewInvalidCommitEntryError() *primitives.JSONError {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdwsapiNewInvalidCommitEntryError.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.NewJSONError(-32602, "Invalid params", "Invalid Commit Entry")
}
func NewInvalidDataPassedError() *primitives.JSONError {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdwsapiNewInvalidDataPassedError.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.NewJSONError(-32602, "Invalid params", "Invalid data passed")
}
func NewInternalDatabaseError() *primitives.JSONError {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdwsapiNewInternalDatabaseError.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.NewJSONError(-32603, "Internal error", "database error")
}

//http://www.jsonrpc.org/specification : -32000 to -32099 error codes are reserved for implementation-defined server-errors.
func NewBlockNotFoundError() *primitives.JSONError {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdwsapiNewBlockNotFoundError.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.NewJSONError(-32008, "Block not found", nil)
}
func NewEntryNotFoundError() *primitives.JSONError {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdwsapiNewEntryNotFoundError.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.NewJSONError(-32008, "Entry not found", nil)
}
func NewObjectNotFoundError() *primitives.JSONError {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdwsapiNewObjectNotFoundError.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.NewJSONError(-32008, "Object not found", nil)
}
func NewMissingChainHeadError() *primitives.JSONError {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdwsapiNewMissingChainHeadError.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.NewJSONError(-32009, "Missing Chain Head", nil)
}
func NewReceiptError() *primitives.JSONError {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdwsapiNewReceiptError.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.NewJSONError(-32010, "Receipt creation error", nil)
}
