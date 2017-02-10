// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

// logger is based on github.com/alexcesaro/log and
// github.com/alexcesaro/log/golog (MIT License)

package anchor

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
	"time"
)

//AnchorRecord is used to construct anchor chain
type AnchorRecord struct {
	AnchorRecordVer int
	DBHeight        uint32
	KeyMR           string
	RecordHeight    uint32 //the block height we intended to put the anchorrecod into

	Bitcoin  *BitcoinStruct  `json:",omitempty"`
	Ethereum *EthereumStruct `json:",omitempty"`
}

type BitcoinStruct struct {
	Address     string //"1HLoD9E4SDFFPDiYfNYnkBLQ85Y51J3Zb1",
	TXID        string //"9b0fc92260312ce44e74ef369f5c66bbb85848f2eddd5a7a1cde251e54ccfdd5", BTC Hash - in reverse byte order
	BlockHeight int32  //345678,
	BlockHash   string //"00000000000000000cc14eacfc7057300aea87bed6fee904fd8e1c1f3dc008d4", BTC Hash - in reverse byte order
	Offset      int32  //87
}

type EthereumStruct struct {
	Address     string //0x30aa981f6d2fce81083e584c8ee2f822b548752f
	TXID        string //0x50ea0effc383542811a58704a6d6842ed6d76439a2d942d941896ad097c06a78
	BlockHeight int64  //293003
	BlockHash   string //0x3b504616495fc9cf7be9b5b776692a9abbfb95491fa62abf62dcdf4d53ff5979
	Offset      int64  //0
}

var _ interfaces.Printable = (*AnchorRecord)(nil)
var _ interfaces.IAnchorRecord = (*AnchorRecord)(nil)

func (e *AnchorRecord) JSONByte() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdanchorAnchorRecordJSONByte.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSON(e)
}

func (e *AnchorRecord) JSONString() (string, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdanchorAnchorRecordJSONString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSONString(e)
}

func (e *AnchorRecord) String() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdanchorAnchorRecordString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	str, _ := e.JSONString()
	return str
}

func (ar *AnchorRecord) Marshal() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdanchorAnchorRecordMarshal.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	data, err := json.Marshal(ar)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (ar *AnchorRecord) MarshalAndSign(priv interfaces.Signer) ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdanchorAnchorRecordMarshalAndSign.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	data, err := ar.Marshal()
	if err != nil {
		return nil, err
	}
	sig := priv.Sign(data)
	return append(data, []byte(fmt.Sprintf("%x", sig.Bytes()))...), nil
}

func (ar *AnchorRecord) MarshalAndSignV2(priv interfaces.Signer) ([]byte, []byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdanchorAnchorRecordMarshalAndSignV2.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	data, err := ar.Marshal()
	if err != nil {
		return nil, nil, err
	}
	sig := priv.Sign(data)
	return data, sig.Bytes(), nil
}

func (ar *AnchorRecord) Unmarshal(data []byte) error {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdanchorAnchorRecordUnmarshal.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if len(data) == 0 {
		return fmt.Errorf("Invalid data passed")
	}
	str := string(data)
	end := strings.LastIndex(str, "}}")
	if end < 0 {
		return fmt.Errorf("Found no closing bracket in `%v`", str)
	}
	str = str[:end+2]
	err := json.Unmarshal([]byte(str), ar)
	if err != nil {
		return err
	}

	return nil
}

func UnmarshalAnchorRecord(data []byte) (*AnchorRecord, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdanchorUnmarshalAnchorRecord.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	ar := new(AnchorRecord)
	err := ar.Unmarshal(data)
	if err != nil {
		return nil, err
	}
	return ar, nil
}

func UnmarshalAndValidateAnchorRecord(data []byte, publicKeys []interfaces.Verifier) (*AnchorRecord, bool, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdanchorUnmarshalAndValidateAnchorRecord.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if len(data) == 0 {
		return nil, false, fmt.Errorf("Invalid data passed")
	}
	str := string(data)
	end := strings.LastIndex(str, "}}")
	if end < 0 {
		return nil, false, fmt.Errorf("Found no closing bracket in `%v`", str)
	}
	anchorStr := str[:end+2]
	signatureStr := str[end+2:]

	sig := new(primitives.ByteSliceSig)
	err := sig.UnmarshalText([]byte(signatureStr))
	if err != nil {
		return nil, false, err
	}
	fixed, err := sig.GetFixed()
	if err != nil {
		return nil, false, err
	}

	valid := false
	for _, publicKey := range publicKeys {
		valid = publicKey.Verify([]byte(anchorStr), &fixed)
		if valid == true {
			break
		}
	}
	if valid == false {
		return nil, false, nil
	}

	ar := new(AnchorRecord)
	err = ar.Unmarshal(data)
	if err != nil {
		return nil, false, err
	}
	return ar, true, nil
}

func UnmarshalAndValidateAnchorRecordV2(data []byte, extIDs [][]byte, publicKeys []interfaces.Verifier) (*AnchorRecord, bool, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdanchorUnmarshalAndValidateAnchorRecordV2.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if len(data) == 0 {
		return nil, false, fmt.Errorf("Invalid data passed")
	}
	if len(extIDs) != 1 {
		return nil, false, fmt.Errorf("Invalid External IDs passed")
	}

	sig := new(primitives.ByteSliceSig)
	sig.UnmarshalBinary(extIDs[0])
	fixed, err := sig.GetFixed()
	if err != nil {
		return nil, false, err
	}

	valid := false
	for _, publicKey := range publicKeys {
		valid = publicKey.Verify(data, &fixed)
		if valid == true {
			break
		}
	}
	if valid == false {
		return nil, false, nil
	}

	ar := new(AnchorRecord)
	err = ar.Unmarshal(data)
	if err != nil {
		return nil, false, err
	}
	return ar, true, nil
}

func UnmarshalAndValidateAnchorEntryAnyVersion(entry interfaces.IEBEntry, publicKeys []interfaces.Verifier) (*AnchorRecord, bool, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdanchorUnmarshalAndValidateAnchorEntryAnyVersion.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	ar, valid, err := UnmarshalAndValidateAnchorRecord(entry.GetContent(), publicKeys)
	if ar == nil {
		ar, valid, err = UnmarshalAndValidateAnchorRecordV2(entry.GetContent(), entry.ExternalIDs(), publicKeys)
		return ar, valid, err
	}
	return ar, valid, err
}

func CreateAnchorRecordFromDBlock(dBlock interfaces.IDirectoryBlock) *AnchorRecord {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdanchorCreateAnchorRecordFromDBlock.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	ar := new(AnchorRecord)
	ar.AnchorRecordVer = 1
	ar.DBHeight = dBlock.GetHeader().GetDBHeight()
	ar.KeyMR = dBlock.DatabasePrimaryIndex().String()
	ar.RecordHeight = ar.DBHeight
	return ar
}
