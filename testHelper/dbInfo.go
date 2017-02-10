package testHelper

import (
	"github.com/FactomProject/factomd/common/directoryBlock/dbInfo"
	"time"
)

func CreateTestDirBlockInfo(prev *dbInfo.DirBlockInfo) *dbInfo.DirBlockInfo {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdtestHelperCreateTestDirBlockInfo.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	dbi := dbInfo.NewDirBlockInfo()
	if prev == nil {
		dbi.DBHeight = 0
	} else {
		dbi.DBHeight = prev.DBHeight + 1
	}
	height := dbi.DBHeight

	dbi.DBHash.UnmarshalBinary(IntToByteSlice(int(height)))
	dbi.Timestamp = int64(height)
	dbi.BTCTxHash.UnmarshalBinary(IntToByteSlice(int(height)))
	dbi.BTCTxOffset = int32(int(height))
	dbi.BTCBlockHeight = int32(height)
	dbi.BTCBlockHash.UnmarshalBinary(IntToByteSlice(255 - int(height)))
	dbi.DBMerkleRoot.UnmarshalBinary(IntToByteSlice(255 - int(height)))
	dbi.BTCConfirmed = height%2 == 1

	return dbi
}

func IntToByteSlice(n int) []byte {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdtestHelperIntToByteSlice.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	answer := make([]byte, 32)
	for i := range answer {
		answer[i] = byte(n)
	}
	return answer
}
