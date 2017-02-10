package entryBlock

import "github.com/prometheus/client_golang/prometheus"

var (
	factomdentryBlockEBlockInit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlock_Init_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlock_Init",
	})
	factomdentryBlockEBlockGetEntryHashes = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlock_GetEntryHashes_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlock_GetEntryHashes",
	})
	factomdentryBlockEBlockGetEntrySigHashes = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlock_GetEntrySigHashes_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlock_GetEntrySigHashes",
	})
	factomdentryBlockEBlockNew = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlock_New_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlock_New",
	})
	factomdentryBlockEBlockGetWelds = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlock_GetWelds_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlock_GetWelds",
	})
	factomdentryBlockEBlockGetWeldHashes = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlock_GetWeldHashes_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlock_GetWeldHashes",
	})
	factomdentryBlockEBlockGetDatabaseHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlock_GetDatabaseHeight_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlock_GetDatabaseHeight",
	})
	factomdentryBlockEBlockGetChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlock_GetChainID_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlock_GetChainID",
	})
	factomdentryBlockEBlockGetHashOfChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlock_GetHashOfChainID_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlock_GetHashOfChainID",
	})
	factomdentryBlockEBlockGetHashOfChainIDHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlock_GetHashOfChainIDHash_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlock_GetHashOfChainIDHash",
	})
	factomdentryBlockEBlockDatabasePrimaryIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlock_DatabasePrimaryIndex_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlock_DatabasePrimaryIndex",
	})
	factomdentryBlockEBlockDatabaseSecondaryIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlock_DatabaseSecondaryIndex_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlock_DatabaseSecondaryIndex",
	})
	factomdentryBlockEBlockGetHeader = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlock_GetHeader_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlock_GetHeader",
	})
	factomdentryBlockEBlockGetBody = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlock_GetBody_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlock_GetBody",
	})
	factomdentryBlockEBlockAddEBEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlock_AddEBEntry_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlock_AddEBEntry",
	})
	factomdentryBlockEBlockAddEndOfMinuteMarker = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlock_AddEndOfMinuteMarker_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlock_AddEndOfMinuteMarker",
	})
	factomdentryBlockEBlockBuildHeader = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlock_BuildHeader_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlock_BuildHeader",
	})
	factomdentryBlockEBlockGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlock_GetHash_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlock_GetHash",
	})
	factomdentryBlockEBlockHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlock_Hash_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlock_Hash",
	})
	factomdentryBlockEBlockHeaderHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlock_HeaderHash_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlock_HeaderHash",
	})
	factomdentryBlockEBlockBodyKeyMR = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlock_BodyKeyMR_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlock_BodyKeyMR",
	})
	factomdentryBlockEBlockKeyMR = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlock_KeyMR_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlock_KeyMR",
	})
	factomdentryBlockEBlockMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlock_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlock_MarshalBinary",
	})
	factomdentryBlockUnmarshalEBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_UnmarshalEBlock_Summary",
		Help: "Summary of calls to  factomd_entryBlock_UnmarshalEBlock",
	})
	factomdentryBlockUnmarshalEBlockData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_UnmarshalEBlockData_Summary",
		Help: "Summary of calls to  factomd_entryBlock_UnmarshalEBlockData",
	})
	factomdentryBlockEBlockUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlock_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlock_UnmarshalBinaryData",
	})
	factomdentryBlockEBlockUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlock_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlock_UnmarshalBinary",
	})
	factomdentryBlockEBlockmarshalBodyBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlock_marshalBodyBinary_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlock_marshalBodyBinary",
	})
	factomdentryBlockEBlockunmarshalBodyBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlock_unmarshalBodyBinaryData_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlock_unmarshalBodyBinaryData",
	})
	factomdentryBlockEBlockunmarshalBodyBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlock_unmarshalBodyBinary_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlock_unmarshalBodyBinary",
	})
	factomdentryBlockEBlockJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlock_JSONByte_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlock_JSONByte",
	})
	factomdentryBlockEBlockJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlock_JSONString_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlock_JSONString",
	})
	factomdentryBlockEBlockString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlock_String_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlock_String",
	})
	factomdentryBlockNewEBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_NewEBlock_Summary",
		Help: "Summary of calls to  factomd_entryBlock_NewEBlock",
	})
	factomdentryBlockNewEBlockBody = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_NewEBlockBody_Summary",
		Help: "Summary of calls to  factomd_entryBlock_NewEBlockBody",
	})
	factomdentryBlockEBlockBodyMR = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlockBody_MR_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlockBody_MR",
	})
	factomdentryBlockEBlockBodyJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlockBody_JSONByte_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlockBody_JSONByte",
	})
	factomdentryBlockEBlockBodyJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlockBody_JSONString_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlockBody_JSONString",
	})
	factomdentryBlockEBlockBodyString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlockBody_String_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlockBody_String",
	})
	factomdentryBlockEBlockBodyGetEBEntries = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlockBody_GetEBEntries_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlockBody_GetEBEntries",
	})
	factomdentryBlockEBlockBodyAddEBEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlockBody_AddEBEntry_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlockBody_AddEBEntry",
	})
	factomdentryBlockEBlockBodyAddEndOfMinuteMarker = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlockBody_AddEndOfMinuteMarker_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlockBody_AddEndOfMinuteMarker",
	})
	factomdentryBlockEBlockHeaderInit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlockHeader_Init_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlockHeader_Init",
	})
	factomdentryBlockNewEBlockHeader = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_NewEBlockHeader_Summary",
		Help: "Summary of calls to  factomd_entryBlock_NewEBlockHeader",
	})
	factomdentryBlockEBlockHeaderJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlockHeader_JSONByte_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlockHeader_JSONByte",
	})
	factomdentryBlockEBlockHeaderJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlockHeader_JSONString_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlockHeader_JSONString",
	})
	factomdentryBlockEBlockHeaderString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlockHeader_String_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlockHeader_String",
	})
	factomdentryBlockEBlockHeaderGetChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlockHeader_GetChainID_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlockHeader_GetChainID",
	})
	factomdentryBlockEBlockHeaderSetChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlockHeader_SetChainID_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlockHeader_SetChainID",
	})
	factomdentryBlockEBlockHeaderGetBodyMR = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlockHeader_GetBodyMR_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlockHeader_GetBodyMR",
	})
	factomdentryBlockEBlockHeaderSetBodyMR = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlockHeader_SetBodyMR_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlockHeader_SetBodyMR",
	})
	factomdentryBlockEBlockHeaderGetPrevKeyMR = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlockHeader_GetPrevKeyMR_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlockHeader_GetPrevKeyMR",
	})
	factomdentryBlockEBlockHeaderSetPrevKeyMR = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlockHeader_SetPrevKeyMR_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlockHeader_SetPrevKeyMR",
	})
	factomdentryBlockEBlockHeaderGetPrevFullHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlockHeader_GetPrevFullHash_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlockHeader_GetPrevFullHash",
	})
	factomdentryBlockEBlockHeaderSetPrevFullHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlockHeader_SetPrevFullHash_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlockHeader_SetPrevFullHash",
	})
	factomdentryBlockEBlockHeaderGetEBSequence = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlockHeader_GetEBSequence_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlockHeader_GetEBSequence",
	})
	factomdentryBlockEBlockHeaderSetEBSequence = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlockHeader_SetEBSequence_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlockHeader_SetEBSequence",
	})
	factomdentryBlockEBlockHeaderGetDBHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlockHeader_GetDBHeight_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlockHeader_GetDBHeight",
	})
	factomdentryBlockEBlockHeaderSetDBHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlockHeader_SetDBHeight_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlockHeader_SetDBHeight",
	})
	factomdentryBlockEBlockHeaderGetEntryCount = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlockHeader_GetEntryCount_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlockHeader_GetEntryCount",
	})
	factomdentryBlockEBlockHeaderSetEntryCount = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlockHeader_SetEntryCount_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlockHeader_SetEntryCount",
	})
	factomdentryBlockEBlockHeaderMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlockHeader_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlockHeader_MarshalBinary",
	})
	factomdentryBlockEBlockHeaderUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlockHeader_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlockHeader_UnmarshalBinaryData",
	})
	factomdentryBlockEBlockHeaderUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_EBlockHeader_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_entryBlock_EBlockHeader_UnmarshalBinary",
	})
	factomdentryBlockEntryKSize = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_Entry_KSize_Summary",
		Help: "Summary of calls to  factomd_entryBlock_Entry_KSize",
	})
	factomdentryBlockEntryNew = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_Entry_New_Summary",
		Help: "Summary of calls to  factomd_entryBlock_Entry_New",
	})
	factomdentryBlockEntryGetDatabaseHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_Entry_GetDatabaseHeight_Summary",
		Help: "Summary of calls to  factomd_entryBlock_Entry_GetDatabaseHeight",
	})
	factomdentryBlockEntryGetWeld = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_Entry_GetWeld_Summary",
		Help: "Summary of calls to  factomd_entryBlock_Entry_GetWeld",
	})
	factomdentryBlockEntryGetWeldHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_Entry_GetWeldHash_Summary",
		Help: "Summary of calls to  factomd_entryBlock_Entry_GetWeldHash",
	})
	factomdentryBlockEntryGetChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_Entry_GetChainID_Summary",
		Help: "Summary of calls to  factomd_entryBlock_Entry_GetChainID",
	})
	factomdentryBlockEntryDatabasePrimaryIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_Entry_DatabasePrimaryIndex_Summary",
		Help: "Summary of calls to  factomd_entryBlock_Entry_DatabasePrimaryIndex",
	})
	factomdentryBlockEntryDatabaseSecondaryIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_Entry_DatabaseSecondaryIndex_Summary",
		Help: "Summary of calls to  factomd_entryBlock_Entry_DatabaseSecondaryIndex",
	})
	factomdentryBlockNewChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_NewChainID_Summary",
		Help: "Summary of calls to  factomd_entryBlock_NewChainID",
	})
	factomdentryBlockEntryGetContent = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_Entry_GetContent_Summary",
		Help: "Summary of calls to  factomd_entryBlock_Entry_GetContent",
	})
	factomdentryBlockEntryGetChainIDHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_Entry_GetChainIDHash_Summary",
		Help: "Summary of calls to  factomd_entryBlock_Entry_GetChainIDHash",
	})
	factomdentryBlockEntryExternalIDs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_Entry_ExternalIDs_Summary",
		Help: "Summary of calls to  factomd_entryBlock_Entry_ExternalIDs",
	})
	factomdentryBlockEntryIsValid = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_Entry_IsValid_Summary",
		Help: "Summary of calls to  factomd_entryBlock_Entry_IsValid",
	})
	factomdentryBlockEntryGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_Entry_GetHash_Summary",
		Help: "Summary of calls to  factomd_entryBlock_Entry_GetHash",
	})
	factomdentryBlockEntryMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_Entry_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_entryBlock_Entry_MarshalBinary",
	})
	factomdentryBlockEntryMarshalExtIDsBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_Entry_MarshalExtIDsBinary_Summary",
		Help: "Summary of calls to  factomd_entryBlock_Entry_MarshalExtIDsBinary",
	})
	factomdentryBlockUnmarshalEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_UnmarshalEntry_Summary",
		Help: "Summary of calls to  factomd_entryBlock_UnmarshalEntry",
	})
	factomdentryBlockEntryUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_Entry_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_entryBlock_Entry_UnmarshalBinaryData",
	})
	factomdentryBlockEntryUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_Entry_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_entryBlock_Entry_UnmarshalBinary",
	})
	factomdentryBlockEntryJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_Entry_JSONByte_Summary",
		Help: "Summary of calls to  factomd_entryBlock_Entry_JSONByte",
	})
	factomdentryBlockEntryJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_Entry_JSONString_Summary",
		Help: "Summary of calls to  factomd_entryBlock_Entry_JSONString",
	})
	factomdentryBlockEntryString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_Entry_String_Summary",
		Help: "Summary of calls to  factomd_entryBlock_Entry_String",
	})
	factomdentryBlockNewEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlock_NewEntry_Summary",
		Help: "Summary of calls to  factomd_entryBlock_NewEntry",
	})
)

func InitPrometheus() {
	prometheus.MustRegister(factomdentryBlockEBlockInit)
	prometheus.MustRegister(factomdentryBlockEBlockGetEntryHashes)
	prometheus.MustRegister(factomdentryBlockEBlockGetEntrySigHashes)
	prometheus.MustRegister(factomdentryBlockEBlockNew)
	prometheus.MustRegister(factomdentryBlockEBlockGetWelds)
	prometheus.MustRegister(factomdentryBlockEBlockGetWeldHashes)
	prometheus.MustRegister(factomdentryBlockEBlockGetDatabaseHeight)
	prometheus.MustRegister(factomdentryBlockEBlockGetChainID)
	prometheus.MustRegister(factomdentryBlockEBlockGetHashOfChainID)
	prometheus.MustRegister(factomdentryBlockEBlockGetHashOfChainIDHash)
	prometheus.MustRegister(factomdentryBlockEBlockDatabasePrimaryIndex)
	prometheus.MustRegister(factomdentryBlockEBlockDatabaseSecondaryIndex)
	prometheus.MustRegister(factomdentryBlockEBlockGetHeader)
	prometheus.MustRegister(factomdentryBlockEBlockGetBody)
	prometheus.MustRegister(factomdentryBlockEBlockAddEBEntry)
	prometheus.MustRegister(factomdentryBlockEBlockAddEndOfMinuteMarker)
	prometheus.MustRegister(factomdentryBlockEBlockBuildHeader)
	prometheus.MustRegister(factomdentryBlockEBlockGetHash)
	prometheus.MustRegister(factomdentryBlockEBlockHash)
	prometheus.MustRegister(factomdentryBlockEBlockHeaderHash)
	prometheus.MustRegister(factomdentryBlockEBlockBodyKeyMR)
	prometheus.MustRegister(factomdentryBlockEBlockKeyMR)
	prometheus.MustRegister(factomdentryBlockEBlockMarshalBinary)
	prometheus.MustRegister(factomdentryBlockUnmarshalEBlock)
	prometheus.MustRegister(factomdentryBlockUnmarshalEBlockData)
	prometheus.MustRegister(factomdentryBlockEBlockUnmarshalBinaryData)
	prometheus.MustRegister(factomdentryBlockEBlockUnmarshalBinary)
	prometheus.MustRegister(factomdentryBlockEBlockmarshalBodyBinary)
	prometheus.MustRegister(factomdentryBlockEBlockunmarshalBodyBinaryData)
	prometheus.MustRegister(factomdentryBlockEBlockunmarshalBodyBinary)
	prometheus.MustRegister(factomdentryBlockEBlockJSONByte)
	prometheus.MustRegister(factomdentryBlockEBlockJSONString)
	prometheus.MustRegister(factomdentryBlockEBlockString)
	prometheus.MustRegister(factomdentryBlockNewEBlock)
	prometheus.MustRegister(factomdentryBlockNewEBlockBody)
	prometheus.MustRegister(factomdentryBlockEBlockBodyMR)
	prometheus.MustRegister(factomdentryBlockEBlockBodyJSONByte)
	prometheus.MustRegister(factomdentryBlockEBlockBodyJSONString)
	prometheus.MustRegister(factomdentryBlockEBlockBodyString)
	prometheus.MustRegister(factomdentryBlockEBlockBodyGetEBEntries)
	prometheus.MustRegister(factomdentryBlockEBlockBodyAddEBEntry)
	prometheus.MustRegister(factomdentryBlockEBlockBodyAddEndOfMinuteMarker)
	prometheus.MustRegister(factomdentryBlockEBlockHeaderInit)
	prometheus.MustRegister(factomdentryBlockNewEBlockHeader)
	prometheus.MustRegister(factomdentryBlockEBlockHeaderJSONByte)
	prometheus.MustRegister(factomdentryBlockEBlockHeaderJSONString)
	prometheus.MustRegister(factomdentryBlockEBlockHeaderString)
	prometheus.MustRegister(factomdentryBlockEBlockHeaderGetChainID)
	prometheus.MustRegister(factomdentryBlockEBlockHeaderSetChainID)
	prometheus.MustRegister(factomdentryBlockEBlockHeaderGetBodyMR)
	prometheus.MustRegister(factomdentryBlockEBlockHeaderSetBodyMR)
	prometheus.MustRegister(factomdentryBlockEBlockHeaderGetPrevKeyMR)
	prometheus.MustRegister(factomdentryBlockEBlockHeaderSetPrevKeyMR)
	prometheus.MustRegister(factomdentryBlockEBlockHeaderGetPrevFullHash)
	prometheus.MustRegister(factomdentryBlockEBlockHeaderSetPrevFullHash)
	prometheus.MustRegister(factomdentryBlockEBlockHeaderGetEBSequence)
	prometheus.MustRegister(factomdentryBlockEBlockHeaderSetEBSequence)
	prometheus.MustRegister(factomdentryBlockEBlockHeaderGetDBHeight)
	prometheus.MustRegister(factomdentryBlockEBlockHeaderSetDBHeight)
	prometheus.MustRegister(factomdentryBlockEBlockHeaderGetEntryCount)
	prometheus.MustRegister(factomdentryBlockEBlockHeaderSetEntryCount)
	prometheus.MustRegister(factomdentryBlockEBlockHeaderMarshalBinary)
	prometheus.MustRegister(factomdentryBlockEBlockHeaderUnmarshalBinaryData)
	prometheus.MustRegister(factomdentryBlockEBlockHeaderUnmarshalBinary)
	prometheus.MustRegister(factomdentryBlockEntryKSize)
	prometheus.MustRegister(factomdentryBlockEntryNew)
	prometheus.MustRegister(factomdentryBlockEntryGetDatabaseHeight)
	prometheus.MustRegister(factomdentryBlockEntryGetWeld)
	prometheus.MustRegister(factomdentryBlockEntryGetWeldHash)
	prometheus.MustRegister(factomdentryBlockEntryGetChainID)
	prometheus.MustRegister(factomdentryBlockEntryDatabasePrimaryIndex)
	prometheus.MustRegister(factomdentryBlockEntryDatabaseSecondaryIndex)
	prometheus.MustRegister(factomdentryBlockNewChainID)
	prometheus.MustRegister(factomdentryBlockEntryGetContent)
	prometheus.MustRegister(factomdentryBlockEntryGetChainIDHash)
	prometheus.MustRegister(factomdentryBlockEntryExternalIDs)
	prometheus.MustRegister(factomdentryBlockEntryIsValid)
	prometheus.MustRegister(factomdentryBlockEntryGetHash)
	prometheus.MustRegister(factomdentryBlockEntryMarshalBinary)
	prometheus.MustRegister(factomdentryBlockEntryMarshalExtIDsBinary)
	prometheus.MustRegister(factomdentryBlockUnmarshalEntry)
	prometheus.MustRegister(factomdentryBlockEntryUnmarshalBinaryData)
	prometheus.MustRegister(factomdentryBlockEntryUnmarshalBinary)
	prometheus.MustRegister(factomdentryBlockEntryJSONByte)
	prometheus.MustRegister(factomdentryBlockEntryJSONString)
	prometheus.MustRegister(factomdentryBlockEntryString)
	prometheus.MustRegister(factomdentryBlockNewEntry)
}
