// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package directoryBlock

import "github.com/prometheus/client_golang/prometheus"

var (
	//directoryBlock.go variables
	directoryBlockErrors = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "factomd_directoryBlock_error_responses",
		Help: "Number of error responses from directoryBlock objects",
	})
	directoryBlockSetEntryHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_SetEntryHash_Summary",
		Help: "Summary of directoryBlock.SetEntryHash call",
	})
	directoryBlockSetABlockHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_ABlockHash_Summary",
		Help: "Summary of directoryBlock.SetABlockHash call",
	})
	directoryBlockSetECBlockHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_SetECBlockHash_Summary",
		Help: "Summary of directoryBlock.SetECBlockHash call",
	})
	directoryBlockSetFBlockHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_SetFBlockHash_Summary",
		Help: "Summary of directoryBlock.SetFBlockHash call",
	})
	directoryBlockGetEntryHashes = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_GetEntryHashes_Summary",
		Help: "Summary of directoryBlock.GetEntryHashes call",
	})
	directoryBlockSort = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_Sort_Summary",
		Help: "Summary of directoryBlock.Sort call",
	})
	directoryBlockGetEntryHashesForBranch = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_GetEntryHashesForBranch_Summary",
		Help: "Summary of directoryBlock.GetEntryHashesForBranch call",
	})
	directoryBlockGetDBEntries = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_GetDBEntries_Summary",
		Help: "Summary of directoryBlock.GetDBEntries call",
	})
	directoryBlockGetEBlockDBEntries = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_GetEBlockDBEntries_Summary",
		Help: "Summary of directoryBlock.GetEBlockDBEntries call",
	})
	directoryBlockGetKeyMR = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_GetKeyMR_Summary",
		Help: "Summary of directoryBlock.GetKeyMR call",
	})
	directoryBlockGetHeader = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_GetHeader_Summary",
		Help: "Summary of directoryBlock.GetHeader call",
	})
	directoryBlockSetHeader = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_SetHeader_Summary",
		Help: "Summary of directoryBlock.SetHeader call",
	})
	directoryBlockSetDBEntries = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_SetDBEntries_Summary",
		Help: "Summary of directoryBlock.SetDBEntries call",
	})
	directoryBlockNew = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_New_Summary",
		Help: "Summary of directoryBlock.New call",
	})
	directoryBlockGetDatabaseHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_GetDatabaseHeight_Summary",
		Help: "Summary of directoryBlock.GetDatabaseHeight call",
	})
	directoryBlockGetChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_GetChainID_Summary",
		Help: "Summary of directoryBlock.GetChainID call",
	})
	directoryBlockDatabasePrimaryIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DatabasePrimaryIndex_Summary",
		Help: "Summary of directoryBlock.DatabasePrimaryIndex call",
	})

	directoryBlockDatabaseSecondaryIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DatabaseSecondaryIndex_Summary",
		Help: "Summary of directoryBlock.DatabaseSecondaryIndex call",
	})
	directoryBlockJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_JSONByte_Summary",
		Help: "Summary of directoryBlock.JSONByte call",
	})
	directoryBlockJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_JSONString_Summary",
		Help: "Summary of directoryBlock.JSONString call",
	})
	directoryBlockJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_JSONBuffer_Summary",
		Help: "Summary of directoryBlock.JSONBuffer call",
	})
	directoryBlockString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_String_Summary",
		Help: "Summary of directoryBlock.String call",
	})
	directoryBlockMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_MarshalBinary_Summary",
		Help: "Summary of directoryBlock.MarshalBinary call",
	})
	directoryBlockBuildBodyMR = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_BuildBodyMR_Summary",
		Help: "Summary of directoryBlock.BuildBodyMR call",
	})
	directoryBlockHeaderHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_HeaderHash_Summary",
		Help: "Summary of directoryBlock.HeaderHash call",
	})
	directoryBlockBodyKeyMR = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_BodyKeyMR_Summary",
		Help: "Summary of directoryBlock.BodyKeyMR call",
	})
	directoryBlockBuildKeyMerkleRoot = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_BuildKeyMerkleRoot_Summary",
		Help: "Summary of directoryBlock.BuildKeyMerkleRoot call",
	})
	directoryBlockUnmarshalDBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_UnmarshalDBlock_Summary",
		Help: "Summary of directoryBlock.UnmarshalDBlock call",
	})
	directoryBlockUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_UnmarshalBinaryData_Summary",
		Help: "Summary of directoryBlock.UnmarshalBinaryData call",
	})
	directoryBlockGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_GetTimestamp_Summary",
		Help: "Summary of directoryBlock.GetTimestamp call",
	})
	directoryBlockUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_UnmarshalBinary_Summary",
		Help: "Summary of directoryBlock.UnmarshalBinary call",
	})
	directoryBlockGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_GetHash_Summary",
		Help: "Summary of directoryBlock.GetHash call",
	})
	directoryBlockGetFullHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_GetFullHash_Summary",
		Help: "Summary of directoryBlock.GetFullHash call",
	})
	directoryBlockAddEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_AddEntry_Summary",
		Help: "Summary of directoryBlock.AddEntry call",
	})
	directoryBlockNewDirectoryBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_NewDirectoryBlock_Summary",
		Help: "Summary of directoryBlock.NewDirectoryBlock call",
	})
	directoryBlockCheckBlockPairIntegrity = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_CheckBlockPairIntegrity_Summary",
		Help: "Summary of directoryBlock.CheckBlockPairIntegrity call",
	})
	directoryBlockMarshalJSON = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_MarshalJSON_Summary",
		Help: "Summary of directoryBlock.MarshalJSON call",
	})

	//directoryBlockEntry.go variables
	directoryBlockEntryGetChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_GetChainID_Summary",
		Help: "Summary of directoryBlock.GetChainID call",
	})
	directoryBlockEntrySetChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_SetChainID_Summary",
		Help: "Summary of directoryBlock.SetChainID call",
	})
	directoryBlockEntryGetKeyMR = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_GetKeyMR_Summary",
		Help: "Summary of directoryBlock.GetKeyMR call",
	})
	directoryBlockEntrySetKeyMR = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_SetKeyMR_Summary",
		Help: "Summary of directoryBlock.SetKeyMR call",
	})
	directoryBlockEntryMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_MarshalBinary_Summary",
		Help: "Summary of directoryBlock.MarshalBinary call",
	})
	directoryBlockEntryUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_UnmarshalBinaryData_Summary",
		Help: "Summary of directoryBlock.UnmarshalBinaryData call",
	})
	directoryBlockEntryUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_UnmarshalBinary_Summary",
		Help: "Summary of directoryBlock.UnmarshalBinary call",
	})
	directoryBlockEntryShaHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_ShaHash_Summary",
		Help: "Summary of directoryBlock.ShaHash call",
	})
	directoryBlockEntryJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_JSONByte_Summary",
		Help: "Summary of directoryBlock.JSONByte call",
	})
	directoryBlockEntryJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_JSONString_Summary",
		Help: "Summary of directoryBlock.JSONString call",
	})
	directoryBlockEntryJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_JSONBuffer_Summary",
		Help: "Summary of directoryBlock.JSONBuffer call",
	})
	directoryBlockEntryString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_String_Summary",
		Help: "Summary of directoryBlock.String call",
	})

	// directoryBlockHeaders.go variables
	directoryBlockHeaderGetVersion = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_GetVersion_Summary",
		Help: "Summary of directoryBlock.GetVersion call",
	})
	directoryBlockHeaderSetVersion = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_SetVersion_Summary",
		Help: "Summary of directoryBlock.SetVersion call",
	})
	directoryBlockHeaderGetNetworkID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_GetNetworkID_Summary",
		Help: "Summary of directoryBlock.GetNetworkID call",
	})
	directoryBlockHeaderSetNetworkID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_SetNetworkID_Summary",
		Help: "Summary of directoryBlock.SetNetworkID call",
	})
	directoryBlockHeaderGetBodyMR = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_GetBodyMR_Summary",
		Help: "Summary of directoryBlock.GetBodyMR call",
	})
	directoryBlockHeaderSetBodyMR = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_SetBodyMR_Summary",
		Help: "Summary of directoryBlock.SetBodyMR call",
	})
	directoryBlockHeaderGetPrevKeyMR = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_GetPrevKeyMR_Summary",
		Help: "Summary of directoryBlock.GetPrevKeyMR call",
	})
	directoryBlockHeaderSetPrevKeyMR = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_SetPrevKeyMR_Summary",
		Help: "Summary of directoryBlock.SetPrevKeyMR call",
	})
	directoryBlockHeaderGetPrevFullHash= prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_GetPrevFullHash_Summary",
		Help: "Summary of directoryBlock.GetPrevFullHash call",
	})
	directoryBlockHeaderSetPrevFullHash= prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_SetPrevFullHash_Summary",
		Help: "Summary of directoryBlock.SetPrevFullHash call",
	})
	directoryBlockHeaderSetPrevFullHash= prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_SetPrevFullHash_Summary",
		Help: "Summary of directoryBlock.SetPrevFullHash call",
	})
	directoryBlockHeaderGetTimestamp= prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_GetTimestamp_Summary",
		Help: "Summary of directoryBlock.GetTimestamp call",
	})
	directoryBlockHeaderSetTimestamp= prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_SetTimestamp_Summary",
		Help: "Summary of directoryBlock.SetTimestamp call",
	})
	directoryBlockHeaderGetDBHeight= prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_GetDBHeight_Summary",
		Help: "Summary of directoryBlock.GetDBHeight call",
	})
	directoryBlockHeaderSetDBHeight= prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_SetDBHeight_Summary",
		Help: "Summary of directoryBlock.SetDBHeight call",
	})
	directoryBlockHeaderGetBlockCount= prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_GetBlockCount_Summary",
		Help: "Summary of directoryBlock.GetBlockCount call",
	})
	directoryBlockHeaderSetBlockCount= prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_SetBlockCount_Summary",
		Help: "Summary of directoryBlock.SetBlockCount call",
	})
	directoryBlockHeaderJSONByte= prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_JSONByte_Summary",
		Help: "Summary of directoryBlock.JSONByte call",
	})
	directoryBlockHeaderJSONString= prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_JSONString_Summary",
		Help: "Summary of directoryBlock.JSONString call",
	})
	directoryBlockHeaderJSONBuffer= prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_JSONBuffer_Summary",
		Help: "Summary of directoryBlock.JSONBuffer call",
	})
    directoryBlockHeaderString= prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_String_Summary",
		Help: "Summary of directoryBlock.String call",
	})
   directoryBlockHeaderMarshalBinary= prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_MarshalBinary_Summary",
		Help: "Summary of directoryBlock.MarshalBinary call",
	})
   directoryBlockHeaderUnmarshalBinaryData= prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_UnmarshalBinaryData_Summary",
		Help: "Summary of directoryBlock.UnmarshalBinaryData call",
	})
   directoryBlockHeaderUnmarshalBinary= prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_UnmarshalBinary_Summary",
		Help: "Summary of directoryBlock.UnmarshalBinary call",
	})
  directoryBlockHeaderMarshalJSON= prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_MarshalJSON_Summary",
		Help: "Summary of directoryBlock.MarshalJSON call",
	})
  directoryBlockHeaderNewDBlockHeader= prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DBlockHeader_Summary",
		Help: "Summary of directoryBlock.DBlockHeader call",
	})


)

func InitPrometheus() {
	//directoryBlock.go variables
	prometheus.MustRegister(directoryBlockErrors)
	prometheus.MustRegister(directoryBlockSetEntryHash)
	prometheus.MustRegister(directoryBlockSetABlockHash)
	prometheus.MustRegister(directoryBlockSetECBlockHash)
	prometheus.MustRegister(directoryBlockSetFBlockHash)
	prometheus.MustRegister(directoryBlockGetEntryHashes)
	prometheus.MustRegister(directoryBlockSort)
	prometheus.MustRegister(directoryBlockGetEntryHashesForBranch)
	prometheus.MustRegister(directoryBlockGetDBEntries)
	prometheus.MustRegister(directoryBlockGetEBlockDBEntries)
	prometheus.MustRegister(directoryBlockGetKeyMR)
	prometheus.MustRegister(directoryBlockGetHeader)
	prometheus.MustRegister(directoryBlockSetHeader)
	prometheus.MustRegister(directoryBlockSetDBEntries)
	prometheus.MustRegister(directoryBlockNew)
	prometheus.MustRegister(directoryBlockGetDatabaseHeight)
	prometheus.MustRegister(directoryBlockGetChainID)
	prometheus.MustRegister(directoryBlockDatabasePrimaryIndex)
	prometheus.MustRegister(directoryBlockDatabaseSecondaryIndex)
	prometheus.MustRegister(directoryBlockJSONByte)
	prometheus.MustRegister(directoryBlockJSONString)
	prometheus.MustRegister(directoryBlockJSONBuffer)
	prometheus.MustRegister(directoryBlockString)
	prometheus.MustRegister(directoryBlockMarshalBinary)
	prometheus.MustRegister(directoryBlockBuildBodyMR)
	prometheus.MustRegister(directoryBlockHeaderHash)
	prometheus.MustRegister(directoryBlockBodyKeyMR)
	prometheus.MustRegister(directoryBlockBuildKeyMerkleRoot)
	prometheus.MustRegister(directoryBlockUnmarshalDBlock)
	prometheus.MustRegister(directoryBlockUnmarshalBinaryData)
	prometheus.MustRegister(directoryBlockGetTimestamp)
	prometheus.MustRegister(directoryBlockUnmarshalBinary)
	prometheus.MustRegister(directoryBlockGetHash)
	prometheus.MustRegister(directoryBlockGetFullHash)
	prometheus.MustRegister(directoryBlockAddEntry)
	prometheus.MustRegister(directoryBlockNewDirectoryBlock)
	prometheus.MustRegister(directoryBlockCheckBlockPairIntegrity)
	prometheus.MustRegister(directoryBlockMarshalJSON)

	//directoryBlockEntry.go variables)
	prometheus.MustRegister(directoryBlockEntryGetChainID)
	prometheus.MustRegister(directoryBlockEntrySetChainID)
	prometheus.MustRegister(directoryBlockEntryGetKeyMR)
	prometheus.MustRegister(directoryBlockEntrySetKeyMR)
	prometheus.MustRegister(directoryBlockEntryMarshalBinary)
	prometheus.MustRegister(directoryBlockEntryUnmarshalBinaryData)
	prometheus.MustRegister(directoryBlockEntryUnmarshalBinary)
	prometheus.MustRegister(directoryBlockEntryShaHash)
	prometheus.MustRegister(directoryBlockEntryJSONByte)
	prometheus.MustRegister(directoryBlockEntryJSONString)
	prometheus.MustRegister(directoryBlockEntryJSONBuffer)
	prometheus.MustRegister(directoryBlockEntryString)

	// directoryBlockHeaders.go variables)
	prometheus.MustRegister(directoryBlockHeaderGetVersion)
	prometheus.MustRegister(directoryBlockHeaderSetVersion)
	prometheus.MustRegister(directoryBlockHeaderGetNetworkID)
	prometheus.MustRegister(directoryBlockHeaderSetNetworkID)
	prometheus.MustRegister(directoryBlockHeaderGetBodyMR)
	prometheus.MustRegister(directoryBlockHeaderSetBodyMR)
	prometheus.MustRegister(directoryBlockHeaderGetPrevKeyMR)
	prometheus.MustRegister(directoryBlockHeaderSetPrevKeyMR)
	prometheus.MustRegister(directoryBlockHeaderGetPrevFullHash)
	prometheus.MustRegister(directoryBlockHeaderSetPrevFullHash)
	prometheus.MustRegister(directoryBlockHeaderSetPrevFullHash)
	prometheus.MustRegister(directoryBlockHeaderGetTimestamp)
	prometheus.MustRegister(directoryBlockHeaderSetTimestamp)
	prometheus.MustRegister(directoryBlockHeaderGetDBHeight)
	prometheus.MustRegister(directoryBlockHeaderSetDBHeight)
	prometheus.MustRegister(directoryBlockHeaderGetBlockCount)
	prometheus.MustRegister(directoryBlockHeaderSetBlockCount)
	prometheus.MustRegister(directoryBlockHeaderJSONByte)
	prometheus.MustRegister(directoryBlockHeaderJSONString)
	prometheus.MustRegister(directoryBlockHeaderJSONBuffer)
	prometheus.MustRegister(directoryBlockHeaderString)
	prometheus.MustRegister(directoryBlockHeaderMarshalBinary)
	prometheus.MustRegister(directoryBlockHeaderUnmarshalBinaryData)
	prometheus.MustRegister(directoryBlockHeaderUnmarshalBinary)
	prometheus.MustRegister(directoryBlockHeaderMarshalJSON)
	prometheus.MustRegister(directoryBlockHeaderNewDBlockHeader)
}
