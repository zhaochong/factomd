package directoryBlock

import "github.com/prometheus/client_golang/prometheus"

var (
	factomddirectoryBlockDirectoryBlockInit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DirectoryBlock_Init_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DirectoryBlock_Init",
	})
	factomddirectoryBlockDirectoryBlockSetEntryHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DirectoryBlock_SetEntryHash_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DirectoryBlock_SetEntryHash",
	})
	factomddirectoryBlockDirectoryBlockSetABlockHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DirectoryBlock_SetABlockHash_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DirectoryBlock_SetABlockHash",
	})
	factomddirectoryBlockDirectoryBlockSetECBlockHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DirectoryBlock_SetECBlockHash_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DirectoryBlock_SetECBlockHash",
	})
	factomddirectoryBlockDirectoryBlockSetFBlockHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DirectoryBlock_SetFBlockHash_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DirectoryBlock_SetFBlockHash",
	})
	factomddirectoryBlockDirectoryBlockGetEntryHashes = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DirectoryBlock_GetEntryHashes_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DirectoryBlock_GetEntryHashes",
	})
	factomddirectoryBlockDirectoryBlockGetEntrySigHashes = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DirectoryBlock_GetEntrySigHashes_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DirectoryBlock_GetEntrySigHashes",
	})
	factomddirectoryBlockDirectoryBlockSort = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DirectoryBlock_Sort_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DirectoryBlock_Sort",
	})
	factomddirectoryBlockDirectoryBlockGetEntryHashesForBranch = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DirectoryBlock_GetEntryHashesForBranch_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DirectoryBlock_GetEntryHashesForBranch",
	})
	factomddirectoryBlockDirectoryBlockGetDBEntries = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DirectoryBlock_GetDBEntries_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DirectoryBlock_GetDBEntries",
	})
	factomddirectoryBlockDirectoryBlockGetEBlockDBEntries = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DirectoryBlock_GetEBlockDBEntries_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DirectoryBlock_GetEBlockDBEntries",
	})
	factomddirectoryBlockDirectoryBlockGetKeyMR = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DirectoryBlock_GetKeyMR_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DirectoryBlock_GetKeyMR",
	})
	factomddirectoryBlockDirectoryBlockGetHeader = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DirectoryBlock_GetHeader_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DirectoryBlock_GetHeader",
	})
	factomddirectoryBlockDirectoryBlockSetHeader = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DirectoryBlock_SetHeader_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DirectoryBlock_SetHeader",
	})
	factomddirectoryBlockDirectoryBlockSetDBEntries = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DirectoryBlock_SetDBEntries_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DirectoryBlock_SetDBEntries",
	})
	factomddirectoryBlockDirectoryBlockNew = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DirectoryBlock_New_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DirectoryBlock_New",
	})
	factomddirectoryBlockDirectoryBlockGetDatabaseHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DirectoryBlock_GetDatabaseHeight_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DirectoryBlock_GetDatabaseHeight",
	})
	factomddirectoryBlockDirectoryBlockGetChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DirectoryBlock_GetChainID_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DirectoryBlock_GetChainID",
	})
	factomddirectoryBlockDirectoryBlockDatabasePrimaryIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DirectoryBlock_DatabasePrimaryIndex_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DirectoryBlock_DatabasePrimaryIndex",
	})
	factomddirectoryBlockDirectoryBlockDatabaseSecondaryIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DirectoryBlock_DatabaseSecondaryIndex_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DirectoryBlock_DatabaseSecondaryIndex",
	})
	factomddirectoryBlockDirectoryBlockJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DirectoryBlock_JSONByte_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DirectoryBlock_JSONByte",
	})
	factomddirectoryBlockDirectoryBlockJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DirectoryBlock_JSONString_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DirectoryBlock_JSONString",
	})
	factomddirectoryBlockDirectoryBlockString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DirectoryBlock_String_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DirectoryBlock_String",
	})
	factomddirectoryBlockDirectoryBlockMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DirectoryBlock_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DirectoryBlock_MarshalBinary",
	})
	factomddirectoryBlockDirectoryBlockBuildBodyMR = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DirectoryBlock_BuildBodyMR_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DirectoryBlock_BuildBodyMR",
	})
	factomddirectoryBlockDirectoryBlockHeaderHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DirectoryBlock_HeaderHash_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DirectoryBlock_HeaderHash",
	})
	factomddirectoryBlockDirectoryBlockBodyKeyMR = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DirectoryBlock_BodyKeyMR_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DirectoryBlock_BodyKeyMR",
	})
	factomddirectoryBlockDirectoryBlockBuildKeyMerkleRoot = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DirectoryBlock_BuildKeyMerkleRoot_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DirectoryBlock_BuildKeyMerkleRoot",
	})
	factomddirectoryBlockUnmarshalDBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_UnmarshalDBlock_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_UnmarshalDBlock",
	})
	factomddirectoryBlockDirectoryBlockUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DirectoryBlock_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DirectoryBlock_UnmarshalBinaryData",
	})
	factomddirectoryBlockDirectoryBlockGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DirectoryBlock_GetTimestamp_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DirectoryBlock_GetTimestamp",
	})
	factomddirectoryBlockDirectoryBlockUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DirectoryBlock_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DirectoryBlock_UnmarshalBinary",
	})
	factomddirectoryBlockDirectoryBlockGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DirectoryBlock_GetHash_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DirectoryBlock_GetHash",
	})
	factomddirectoryBlockDirectoryBlockGetFullHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DirectoryBlock_GetFullHash_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DirectoryBlock_GetFullHash",
	})
	factomddirectoryBlockDirectoryBlockAddEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DirectoryBlock_AddEntry_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DirectoryBlock_AddEntry",
	})
	factomddirectoryBlockNewDirectoryBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_NewDirectoryBlock_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_NewDirectoryBlock",
	})
	factomddirectoryBlockCheckBlockPairIntegrity = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_CheckBlockPairIntegrity_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_CheckBlockPairIntegrity",
	})
	factomddirectoryBlockDirectoryBlockMarshalJSON = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DirectoryBlock_MarshalJSON_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DirectoryBlock_MarshalJSON",
	})
	factomddirectoryBlockDBEntryGetChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DBEntry_GetChainID_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DBEntry_GetChainID",
	})
	factomddirectoryBlockDBEntrySetChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DBEntry_SetChainID_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DBEntry_SetChainID",
	})
	factomddirectoryBlockDBEntryGetKeyMR = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DBEntry_GetKeyMR_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DBEntry_GetKeyMR",
	})
	factomddirectoryBlockDBEntrySetKeyMR = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DBEntry_SetKeyMR_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DBEntry_SetKeyMR",
	})
	factomddirectoryBlockDBEntryMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DBEntry_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DBEntry_MarshalBinary",
	})
	factomddirectoryBlockDBEntryUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DBEntry_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DBEntry_UnmarshalBinaryData",
	})
	factomddirectoryBlockDBEntryUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DBEntry_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DBEntry_UnmarshalBinary",
	})
	factomddirectoryBlockDBEntryShaHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DBEntry_ShaHash_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DBEntry_ShaHash",
	})
	factomddirectoryBlockDBEntryJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DBEntry_JSONByte_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DBEntry_JSONByte",
	})
	factomddirectoryBlockDBEntryJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DBEntry_JSONString_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DBEntry_JSONString",
	})
	factomddirectoryBlockDBEntryString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DBEntry_String_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DBEntry_String",
	})
	factomddirectoryBlockDBlockHeaderInit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DBlockHeader_Init_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DBlockHeader_Init",
	})
	factomddirectoryBlockDBlockHeaderGetVersion = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DBlockHeader_GetVersion_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DBlockHeader_GetVersion",
	})
	factomddirectoryBlockDBlockHeaderSetVersion = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DBlockHeader_SetVersion_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DBlockHeader_SetVersion",
	})
	factomddirectoryBlockDBlockHeaderGetNetworkID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DBlockHeader_GetNetworkID_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DBlockHeader_GetNetworkID",
	})
	factomddirectoryBlockDBlockHeaderSetNetworkID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DBlockHeader_SetNetworkID_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DBlockHeader_SetNetworkID",
	})
	factomddirectoryBlockDBlockHeaderGetBodyMR = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DBlockHeader_GetBodyMR_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DBlockHeader_GetBodyMR",
	})
	factomddirectoryBlockDBlockHeaderSetBodyMR = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DBlockHeader_SetBodyMR_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DBlockHeader_SetBodyMR",
	})
	factomddirectoryBlockDBlockHeaderGetPrevKeyMR = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DBlockHeader_GetPrevKeyMR_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DBlockHeader_GetPrevKeyMR",
	})
	factomddirectoryBlockDBlockHeaderSetPrevKeyMR = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DBlockHeader_SetPrevKeyMR_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DBlockHeader_SetPrevKeyMR",
	})
	factomddirectoryBlockDBlockHeaderGetPrevFullHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DBlockHeader_GetPrevFullHash_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DBlockHeader_GetPrevFullHash",
	})
	factomddirectoryBlockDBlockHeaderSetPrevFullHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DBlockHeader_SetPrevFullHash_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DBlockHeader_SetPrevFullHash",
	})
	factomddirectoryBlockDBlockHeaderGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DBlockHeader_GetTimestamp_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DBlockHeader_GetTimestamp",
	})
	factomddirectoryBlockDBlockHeaderSetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DBlockHeader_SetTimestamp_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DBlockHeader_SetTimestamp",
	})
	factomddirectoryBlockDBlockHeaderGetDBHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DBlockHeader_GetDBHeight_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DBlockHeader_GetDBHeight",
	})
	factomddirectoryBlockDBlockHeaderSetDBHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DBlockHeader_SetDBHeight_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DBlockHeader_SetDBHeight",
	})
	factomddirectoryBlockDBlockHeaderGetBlockCount = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DBlockHeader_GetBlockCount_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DBlockHeader_GetBlockCount",
	})
	factomddirectoryBlockDBlockHeaderSetBlockCount = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DBlockHeader_SetBlockCount_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DBlockHeader_SetBlockCount",
	})
	factomddirectoryBlockDBlockHeaderJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DBlockHeader_JSONByte_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DBlockHeader_JSONByte",
	})
	factomddirectoryBlockDBlockHeaderJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DBlockHeader_JSONString_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DBlockHeader_JSONString",
	})
	factomddirectoryBlockDBlockHeaderString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DBlockHeader_String_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DBlockHeader_String",
	})
	factomddirectoryBlockDBlockHeaderMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DBlockHeader_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DBlockHeader_MarshalBinary",
	})
	factomddirectoryBlockDBlockHeaderUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DBlockHeader_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DBlockHeader_UnmarshalBinaryData",
	})
	factomddirectoryBlockDBlockHeaderUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DBlockHeader_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DBlockHeader_UnmarshalBinary",
	})
	factomddirectoryBlockDBlockHeaderMarshalJSON = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_DBlockHeader_MarshalJSON_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_DBlockHeader_MarshalJSON",
	})
	factomddirectoryBlockNewDBlockHeader = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_directoryBlock_NewDBlockHeader_Summary",
		Help: "Summary of calls to  factomd_directoryBlock_NewDBlockHeader",
	})
)

func InitPrometheus() {
	prometheus.MustRegister(factomddirectoryBlockDirectoryBlockInit)
	prometheus.MustRegister(factomddirectoryBlockDirectoryBlockSetEntryHash)
	prometheus.MustRegister(factomddirectoryBlockDirectoryBlockSetABlockHash)
	prometheus.MustRegister(factomddirectoryBlockDirectoryBlockSetECBlockHash)
	prometheus.MustRegister(factomddirectoryBlockDirectoryBlockSetFBlockHash)
	prometheus.MustRegister(factomddirectoryBlockDirectoryBlockGetEntryHashes)
	prometheus.MustRegister(factomddirectoryBlockDirectoryBlockGetEntrySigHashes)
	prometheus.MustRegister(factomddirectoryBlockDirectoryBlockSort)
	prometheus.MustRegister(factomddirectoryBlockDirectoryBlockGetEntryHashesForBranch)
	prometheus.MustRegister(factomddirectoryBlockDirectoryBlockGetDBEntries)
	prometheus.MustRegister(factomddirectoryBlockDirectoryBlockGetEBlockDBEntries)
	prometheus.MustRegister(factomddirectoryBlockDirectoryBlockGetKeyMR)
	prometheus.MustRegister(factomddirectoryBlockDirectoryBlockGetHeader)
	prometheus.MustRegister(factomddirectoryBlockDirectoryBlockSetHeader)
	prometheus.MustRegister(factomddirectoryBlockDirectoryBlockSetDBEntries)
	prometheus.MustRegister(factomddirectoryBlockDirectoryBlockNew)
	prometheus.MustRegister(factomddirectoryBlockDirectoryBlockGetDatabaseHeight)
	prometheus.MustRegister(factomddirectoryBlockDirectoryBlockGetChainID)
	prometheus.MustRegister(factomddirectoryBlockDirectoryBlockDatabasePrimaryIndex)
	prometheus.MustRegister(factomddirectoryBlockDirectoryBlockDatabaseSecondaryIndex)
	prometheus.MustRegister(factomddirectoryBlockDirectoryBlockJSONByte)
	prometheus.MustRegister(factomddirectoryBlockDirectoryBlockJSONString)
	prometheus.MustRegister(factomddirectoryBlockDirectoryBlockString)
	prometheus.MustRegister(factomddirectoryBlockDirectoryBlockMarshalBinary)
	prometheus.MustRegister(factomddirectoryBlockDirectoryBlockBuildBodyMR)
	prometheus.MustRegister(factomddirectoryBlockDirectoryBlockHeaderHash)
	prometheus.MustRegister(factomddirectoryBlockDirectoryBlockBodyKeyMR)
	prometheus.MustRegister(factomddirectoryBlockDirectoryBlockBuildKeyMerkleRoot)
	prometheus.MustRegister(factomddirectoryBlockUnmarshalDBlock)
	prometheus.MustRegister(factomddirectoryBlockDirectoryBlockUnmarshalBinaryData)
	prometheus.MustRegister(factomddirectoryBlockDirectoryBlockGetTimestamp)
	prometheus.MustRegister(factomddirectoryBlockDirectoryBlockUnmarshalBinary)
	prometheus.MustRegister(factomddirectoryBlockDirectoryBlockGetHash)
	prometheus.MustRegister(factomddirectoryBlockDirectoryBlockGetFullHash)
	prometheus.MustRegister(factomddirectoryBlockDirectoryBlockAddEntry)
	prometheus.MustRegister(factomddirectoryBlockNewDirectoryBlock)
	prometheus.MustRegister(factomddirectoryBlockCheckBlockPairIntegrity)
	prometheus.MustRegister(factomddirectoryBlockDirectoryBlockMarshalJSON)
	prometheus.MustRegister(factomddirectoryBlockDBEntryGetChainID)
	prometheus.MustRegister(factomddirectoryBlockDBEntrySetChainID)
	prometheus.MustRegister(factomddirectoryBlockDBEntryGetKeyMR)
	prometheus.MustRegister(factomddirectoryBlockDBEntrySetKeyMR)
	prometheus.MustRegister(factomddirectoryBlockDBEntryMarshalBinary)
	prometheus.MustRegister(factomddirectoryBlockDBEntryUnmarshalBinaryData)
	prometheus.MustRegister(factomddirectoryBlockDBEntryUnmarshalBinary)
	prometheus.MustRegister(factomddirectoryBlockDBEntryShaHash)
	prometheus.MustRegister(factomddirectoryBlockDBEntryJSONByte)
	prometheus.MustRegister(factomddirectoryBlockDBEntryJSONString)
	prometheus.MustRegister(factomddirectoryBlockDBEntryString)
	prometheus.MustRegister(factomddirectoryBlockDBlockHeaderInit)
	prometheus.MustRegister(factomddirectoryBlockDBlockHeaderGetVersion)
	prometheus.MustRegister(factomddirectoryBlockDBlockHeaderSetVersion)
	prometheus.MustRegister(factomddirectoryBlockDBlockHeaderGetNetworkID)
	prometheus.MustRegister(factomddirectoryBlockDBlockHeaderSetNetworkID)
	prometheus.MustRegister(factomddirectoryBlockDBlockHeaderGetBodyMR)
	prometheus.MustRegister(factomddirectoryBlockDBlockHeaderSetBodyMR)
	prometheus.MustRegister(factomddirectoryBlockDBlockHeaderGetPrevKeyMR)
	prometheus.MustRegister(factomddirectoryBlockDBlockHeaderSetPrevKeyMR)
	prometheus.MustRegister(factomddirectoryBlockDBlockHeaderGetPrevFullHash)
	prometheus.MustRegister(factomddirectoryBlockDBlockHeaderSetPrevFullHash)
	prometheus.MustRegister(factomddirectoryBlockDBlockHeaderGetTimestamp)
	prometheus.MustRegister(factomddirectoryBlockDBlockHeaderSetTimestamp)
	prometheus.MustRegister(factomddirectoryBlockDBlockHeaderGetDBHeight)
	prometheus.MustRegister(factomddirectoryBlockDBlockHeaderSetDBHeight)
	prometheus.MustRegister(factomddirectoryBlockDBlockHeaderGetBlockCount)
	prometheus.MustRegister(factomddirectoryBlockDBlockHeaderSetBlockCount)
	prometheus.MustRegister(factomddirectoryBlockDBlockHeaderJSONByte)
	prometheus.MustRegister(factomddirectoryBlockDBlockHeaderJSONString)
	prometheus.MustRegister(factomddirectoryBlockDBlockHeaderString)
	prometheus.MustRegister(factomddirectoryBlockDBlockHeaderMarshalBinary)
	prometheus.MustRegister(factomddirectoryBlockDBlockHeaderUnmarshalBinaryData)
	prometheus.MustRegister(factomddirectoryBlockDBlockHeaderUnmarshalBinary)
	prometheus.MustRegister(factomddirectoryBlockDBlockHeaderMarshalJSON)
	prometheus.MustRegister(factomddirectoryBlockNewDBlockHeader)
}
