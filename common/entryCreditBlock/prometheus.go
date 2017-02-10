package entryCreditBlock

import "github.com/prometheus/client_golang/prometheus"

var (
	factomdentryCreditBlockCommitChainInit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_CommitChain_Init_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_CommitChain_Init",
	})
	factomdentryCreditBlockCommitChainString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_CommitChain_String_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_CommitChain_String",
	})
	factomdentryCreditBlockNewCommitChain = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_NewCommitChain_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_NewCommitChain",
	})
	factomdentryCreditBlockCommitChainGetEntryHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_CommitChain_GetEntryHash_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_CommitChain_GetEntryHash",
	})
	factomdentryCreditBlockCommitChainIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_CommitChain_IsSameAs_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_CommitChain_IsSameAs",
	})
	factomdentryCreditBlockCommitChainHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_CommitChain_Hash_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_CommitChain_Hash",
	})
	factomdentryCreditBlockCommitChainIsInterpretable = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_CommitChain_IsInterpretable_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_CommitChain_IsInterpretable",
	})
	factomdentryCreditBlockCommitChainInterpret = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_CommitChain_Interpret_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_CommitChain_Interpret",
	})
	factomdentryCreditBlockCommitChainCommitMsg = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_CommitChain_CommitMsg_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_CommitChain_CommitMsg",
	})
	factomdentryCreditBlockCommitChainGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_CommitChain_GetTimestamp_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_CommitChain_GetTimestamp",
	})
	factomdentryCreditBlockCommitChainIsValid = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_CommitChain_IsValid_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_CommitChain_IsValid",
	})
	factomdentryCreditBlockCommitChainGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_CommitChain_GetHash_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_CommitChain_GetHash",
	})
	factomdentryCreditBlockCommitChainGetSigHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_CommitChain_GetSigHash_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_CommitChain_GetSigHash",
	})
	factomdentryCreditBlockCommitChainMarshalBinarySig = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_CommitChain_MarshalBinarySig_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_CommitChain_MarshalBinarySig",
	})
	factomdentryCreditBlockCommitChainMarshalBinaryTransaction = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_CommitChain_MarshalBinaryTransaction_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_CommitChain_MarshalBinaryTransaction",
	})
	factomdentryCreditBlockCommitChainMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_CommitChain_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_CommitChain_MarshalBinary",
	})
	factomdentryCreditBlockCommitChainSign = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_CommitChain_Sign_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_CommitChain_Sign",
	})
	factomdentryCreditBlockCommitChainValidateSignatures = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_CommitChain_ValidateSignatures_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_CommitChain_ValidateSignatures",
	})
	factomdentryCreditBlockCommitChainECID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_CommitChain_ECID_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_CommitChain_ECID",
	})
	factomdentryCreditBlockCommitChainUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_CommitChain_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_CommitChain_UnmarshalBinaryData",
	})
	factomdentryCreditBlockCommitChainUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_CommitChain_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_CommitChain_UnmarshalBinary",
	})
	factomdentryCreditBlockCommitChainJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_CommitChain_JSONByte_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_CommitChain_JSONByte",
	})
	factomdentryCreditBlockCommitChainJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_CommitChain_JSONString_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_CommitChain_JSONString",
	})
	factomdentryCreditBlockCommitEntryString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_CommitEntry_String_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_CommitEntry_String",
	})
	factomdentryCreditBlockCommitEntryGetEntryHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_CommitEntry_GetEntryHash_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_CommitEntry_GetEntryHash",
	})
	factomdentryCreditBlockCommitEntryIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_CommitEntry_IsSameAs_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_CommitEntry_IsSameAs",
	})
	factomdentryCreditBlockNewCommitEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_NewCommitEntry_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_NewCommitEntry",
	})
	factomdentryCreditBlockCommitEntryHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_CommitEntry_Hash_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_CommitEntry_Hash",
	})
	factomdentryCreditBlockCommitEntryIsInterpretable = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_CommitEntry_IsInterpretable_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_CommitEntry_IsInterpretable",
	})
	factomdentryCreditBlockCommitEntryInterpret = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_CommitEntry_Interpret_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_CommitEntry_Interpret",
	})
	factomdentryCreditBlockCommitEntryCommitMsg = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_CommitEntry_CommitMsg_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_CommitEntry_CommitMsg",
	})
	factomdentryCreditBlockCommitEntryGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_CommitEntry_GetTimestamp_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_CommitEntry_GetTimestamp",
	})
	factomdentryCreditBlockCommitEntryInTime = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_CommitEntry_InTime_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_CommitEntry_InTime",
	})
	factomdentryCreditBlockCommitEntryIsValid = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_CommitEntry_IsValid_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_CommitEntry_IsValid",
	})
	factomdentryCreditBlockCommitEntryGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_CommitEntry_GetHash_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_CommitEntry_GetHash",
	})
	factomdentryCreditBlockCommitEntryGetSigHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_CommitEntry_GetSigHash_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_CommitEntry_GetSigHash",
	})
	factomdentryCreditBlockCommitEntryMarshalBinarySig = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_CommitEntry_MarshalBinarySig_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_CommitEntry_MarshalBinarySig",
	})
	factomdentryCreditBlockCommitEntryMarshalBinaryTransaction = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_CommitEntry_MarshalBinaryTransaction_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_CommitEntry_MarshalBinaryTransaction",
	})
	factomdentryCreditBlockCommitEntryMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_CommitEntry_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_CommitEntry_MarshalBinary",
	})
	factomdentryCreditBlockCommitEntrySign = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_CommitEntry_Sign_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_CommitEntry_Sign",
	})
	factomdentryCreditBlockCommitEntryValidateSignatures = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_CommitEntry_ValidateSignatures_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_CommitEntry_ValidateSignatures",
	})
	factomdentryCreditBlockCommitEntryECID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_CommitEntry_ECID_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_CommitEntry_ECID",
	})
	factomdentryCreditBlockCommitEntryUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_CommitEntry_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_CommitEntry_UnmarshalBinaryData",
	})
	factomdentryCreditBlockCommitEntryUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_CommitEntry_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_CommitEntry_UnmarshalBinary",
	})
	factomdentryCreditBlockCommitEntryJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_CommitEntry_JSONByte_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_CommitEntry_JSONByte",
	})
	factomdentryCreditBlockCommitEntryJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_CommitEntry_JSONString_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_CommitEntry_JSONString",
	})
	factomdentryCreditBlockECBlockInit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlock_Init_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlock_Init",
	})
	factomdentryCreditBlockECBlockUpdateState = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlock_UpdateState_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlock_UpdateState",
	})
	factomdentryCreditBlockECBlockString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlock_String_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlock_String",
	})
	factomdentryCreditBlockECBlockGetEntries = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlock_GetEntries_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlock_GetEntries",
	})
	factomdentryCreditBlockECBlockGetEntryByHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlock_GetEntryByHash_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlock_GetEntryByHash",
	})
	factomdentryCreditBlockECBlockGetEntryHashes = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlock_GetEntryHashes_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlock_GetEntryHashes",
	})
	factomdentryCreditBlockECBlockGetEntrySigHashes = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlock_GetEntrySigHashes_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlock_GetEntrySigHashes",
	})
	factomdentryCreditBlockECBlockGetBody = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlock_GetBody_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlock_GetBody",
	})
	factomdentryCreditBlockECBlockGetHeader = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlock_GetHeader_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlock_GetHeader",
	})
	factomdentryCreditBlockECBlockNew = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlock_New_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlock_New",
	})
	factomdentryCreditBlockECBlockGetDatabaseHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlock_GetDatabaseHeight_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlock_GetDatabaseHeight",
	})
	factomdentryCreditBlockECBlockGetChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlock_GetChainID_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlock_GetChainID",
	})
	factomdentryCreditBlockECBlockDatabasePrimaryIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlock_DatabasePrimaryIndex_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlock_DatabasePrimaryIndex",
	})
	factomdentryCreditBlockECBlockDatabaseSecondaryIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlock_DatabaseSecondaryIndex_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlock_DatabaseSecondaryIndex",
	})
	factomdentryCreditBlockECBlockAddEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlock_AddEntry_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlock_AddEntry",
	})
	factomdentryCreditBlockECBlockGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlock_GetHash_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlock_GetHash",
	})
	factomdentryCreditBlockECBlockGetFullHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlock_GetFullHash_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlock_GetFullHash",
	})
	factomdentryCreditBlockECBlockHeaderHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlock_HeaderHash_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlock_HeaderHash",
	})
	factomdentryCreditBlockECBlockMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlock_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlock_MarshalBinary",
	})
	factomdentryCreditBlockECBlockBuildHeader = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlock_BuildHeader_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlock_BuildHeader",
	})
	factomdentryCreditBlockUnmarshalECBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_UnmarshalECBlock_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_UnmarshalECBlock",
	})
	factomdentryCreditBlockECBlockUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlock_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlock_UnmarshalBinaryData",
	})
	factomdentryCreditBlockECBlockUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlock_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlock_UnmarshalBinary",
	})
	factomdentryCreditBlockECBlockmarshalBodyBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlock_marshalBodyBinary_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlock_marshalBodyBinary",
	})
	factomdentryCreditBlockECBlockunmarshalBodyBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlock_unmarshalBodyBinaryData_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlock_unmarshalBodyBinaryData",
	})
	factomdentryCreditBlockECBlockunmarshalBodyBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlock_unmarshalBodyBinary_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlock_unmarshalBodyBinary",
	})
	factomdentryCreditBlockECBlockJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlock_JSONByte_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlock_JSONByte",
	})
	factomdentryCreditBlockECBlockJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlock_JSONString_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlock_JSONString",
	})
	factomdentryCreditBlockNewECBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_NewECBlock_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_NewECBlock",
	})
	factomdentryCreditBlockNextECBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_NextECBlock_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_NextECBlock",
	})
	factomdentryCreditBlockCheckBlockPairIntegrity = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_CheckBlockPairIntegrity_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_CheckBlockPairIntegrity",
	})
	factomdentryCreditBlockECBlockBodyGetEntries = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlockBody_GetEntries_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlockBody_GetEntries",
	})
	factomdentryCreditBlockECBlockBodyAddEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlockBody_AddEntry_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlockBody_AddEntry",
	})
	factomdentryCreditBlockECBlockBodySetEntries = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlockBody_SetEntries_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlockBody_SetEntries",
	})
	factomdentryCreditBlockECBlockBodyJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlockBody_JSONByte_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlockBody_JSONByte",
	})
	factomdentryCreditBlockECBlockBodyJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlockBody_JSONString_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlockBody_JSONString",
	})
	factomdentryCreditBlockECBlockBodyString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlockBody_String_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlockBody_String",
	})
	factomdentryCreditBlockNewECBlockBody = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_NewECBlockBody_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_NewECBlockBody",
	})
	factomdentryCreditBlockECBlockHeaderInit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlockHeader_Init_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlockHeader_Init",
	})
	factomdentryCreditBlockECBlockHeaderString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlockHeader_String_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlockHeader_String",
	})
	factomdentryCreditBlockECBlockHeaderSetBodySize = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlockHeader_SetBodySize_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlockHeader_SetBodySize",
	})
	factomdentryCreditBlockECBlockHeaderGetBodySize = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlockHeader_GetBodySize_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlockHeader_GetBodySize",
	})
	factomdentryCreditBlockECBlockHeaderSetObjectCount = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlockHeader_SetObjectCount_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlockHeader_SetObjectCount",
	})
	factomdentryCreditBlockECBlockHeaderGetObjectCount = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlockHeader_GetObjectCount_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlockHeader_GetObjectCount",
	})
	factomdentryCreditBlockECBlockHeaderSetHeaderExpansionArea = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlockHeader_SetHeaderExpansionArea_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlockHeader_SetHeaderExpansionArea",
	})
	factomdentryCreditBlockECBlockHeaderGetHeaderExpansionArea = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlockHeader_GetHeaderExpansionArea_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlockHeader_GetHeaderExpansionArea",
	})
	factomdentryCreditBlockECBlockHeaderSetBodyHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlockHeader_SetBodyHash_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlockHeader_SetBodyHash",
	})
	factomdentryCreditBlockECBlockHeaderGetBodyHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlockHeader_GetBodyHash_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlockHeader_GetBodyHash",
	})
	factomdentryCreditBlockECBlockHeaderGetECChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlockHeader_GetECChainID_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlockHeader_GetECChainID",
	})
	factomdentryCreditBlockECBlockHeaderSetPrevHeaderHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlockHeader_SetPrevHeaderHash_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlockHeader_SetPrevHeaderHash",
	})
	factomdentryCreditBlockECBlockHeaderGetPrevHeaderHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlockHeader_GetPrevHeaderHash_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlockHeader_GetPrevHeaderHash",
	})
	factomdentryCreditBlockECBlockHeaderSetPrevFullHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlockHeader_SetPrevFullHash_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlockHeader_SetPrevFullHash",
	})
	factomdentryCreditBlockECBlockHeaderGetPrevFullHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlockHeader_GetPrevFullHash_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlockHeader_GetPrevFullHash",
	})
	factomdentryCreditBlockECBlockHeaderSetDBHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlockHeader_SetDBHeight_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlockHeader_SetDBHeight",
	})
	factomdentryCreditBlockECBlockHeaderGetDBHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlockHeader_GetDBHeight_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlockHeader_GetDBHeight",
	})
	factomdentryCreditBlockNewECBlockHeader = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_NewECBlockHeader_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_NewECBlockHeader",
	})
	factomdentryCreditBlockECBlockHeaderJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlockHeader_JSONByte_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlockHeader_JSONByte",
	})
	factomdentryCreditBlockECBlockHeaderJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlockHeader_JSONString_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlockHeader_JSONString",
	})
	factomdentryCreditBlockECBlockHeaderMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlockHeader_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlockHeader_MarshalBinary",
	})
	factomdentryCreditBlockECBlockHeaderUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlockHeader_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlockHeader_UnmarshalBinaryData",
	})
	factomdentryCreditBlockECBlockHeaderUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlockHeader_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlockHeader_UnmarshalBinary",
	})
	factomdentryCreditBlockECBlockHeaderMarshalJSON = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ECBlockHeader_MarshalJSON_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ECBlockHeader_MarshalJSON",
	})
	factomdentryCreditBlockIncreaseBalanceInit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_IncreaseBalance_Init_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_IncreaseBalance_Init",
	})
	factomdentryCreditBlockIncreaseBalanceString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_IncreaseBalance_String_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_IncreaseBalance_String",
	})
	factomdentryCreditBlockNewIncreaseBalance = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_NewIncreaseBalance_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_NewIncreaseBalance",
	})
	factomdentryCreditBlockIncreaseBalanceGetEntryHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_IncreaseBalance_GetEntryHash_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_IncreaseBalance_GetEntryHash",
	})
	factomdentryCreditBlockIncreaseBalanceHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_IncreaseBalance_Hash_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_IncreaseBalance_Hash",
	})
	factomdentryCreditBlockIncreaseBalanceGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_IncreaseBalance_GetHash_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_IncreaseBalance_GetHash",
	})
	factomdentryCreditBlockIncreaseBalanceGetSigHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_IncreaseBalance_GetSigHash_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_IncreaseBalance_GetSigHash",
	})
	factomdentryCreditBlockIncreaseBalanceECID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_IncreaseBalance_ECID_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_IncreaseBalance_ECID",
	})
	factomdentryCreditBlockIncreaseBalanceIsInterpretable = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_IncreaseBalance_IsInterpretable_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_IncreaseBalance_IsInterpretable",
	})
	factomdentryCreditBlockIncreaseBalanceInterpret = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_IncreaseBalance_Interpret_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_IncreaseBalance_Interpret",
	})
	factomdentryCreditBlockIncreaseBalanceMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_IncreaseBalance_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_IncreaseBalance_MarshalBinary",
	})
	factomdentryCreditBlockIncreaseBalanceUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_IncreaseBalance_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_IncreaseBalance_UnmarshalBinaryData",
	})
	factomdentryCreditBlockIncreaseBalanceUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_IncreaseBalance_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_IncreaseBalance_UnmarshalBinary",
	})
	factomdentryCreditBlockIncreaseBalanceJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_IncreaseBalance_JSONByte_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_IncreaseBalance_JSONByte",
	})
	factomdentryCreditBlockIncreaseBalanceJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_IncreaseBalance_JSONString_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_IncreaseBalance_JSONString",
	})
	factomdentryCreditBlockIncreaseBalanceGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_IncreaseBalance_GetTimestamp_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_IncreaseBalance_GetTimestamp",
	})
	factomdentryCreditBlockMinuteNumberHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_MinuteNumber_Hash_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_MinuteNumber_Hash",
	})
	factomdentryCreditBlockMinuteNumberGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_MinuteNumber_GetHash_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_MinuteNumber_GetHash",
	})
	factomdentryCreditBlockMinuteNumberGetSigHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_MinuteNumber_GetSigHash_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_MinuteNumber_GetSigHash",
	})
	factomdentryCreditBlockMinuteNumberGetEntryHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_MinuteNumber_GetEntryHash_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_MinuteNumber_GetEntryHash",
	})
	factomdentryCreditBlockMinuteNumberIsInterpretable = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_MinuteNumber_IsInterpretable_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_MinuteNumber_IsInterpretable",
	})
	factomdentryCreditBlockMinuteNumberInterpret = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_MinuteNumber_Interpret_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_MinuteNumber_Interpret",
	})
	factomdentryCreditBlockNewMinuteNumber = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_NewMinuteNumber_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_NewMinuteNumber",
	})
	factomdentryCreditBlockMinuteNumberECID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_MinuteNumber_ECID_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_MinuteNumber_ECID",
	})
	factomdentryCreditBlockMinuteNumberMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_MinuteNumber_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_MinuteNumber_MarshalBinary",
	})
	factomdentryCreditBlockMinuteNumberUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_MinuteNumber_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_MinuteNumber_UnmarshalBinaryData",
	})
	factomdentryCreditBlockMinuteNumberUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_MinuteNumber_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_MinuteNumber_UnmarshalBinary",
	})
	factomdentryCreditBlockMinuteNumberJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_MinuteNumber_JSONByte_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_MinuteNumber_JSONByte",
	})
	factomdentryCreditBlockMinuteNumberJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_MinuteNumber_JSONString_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_MinuteNumber_JSONString",
	})
	factomdentryCreditBlockMinuteNumberString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_MinuteNumber_String_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_MinuteNumber_String",
	})
	factomdentryCreditBlockMinuteNumberGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_MinuteNumber_GetTimestamp_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_MinuteNumber_GetTimestamp",
	})
	factomdentryCreditBlockServerIndexNumberString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ServerIndexNumber_String_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ServerIndexNumber_String",
	})
	factomdentryCreditBlockServerIndexNumberHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ServerIndexNumber_Hash_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ServerIndexNumber_Hash",
	})
	factomdentryCreditBlockServerIndexNumberGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ServerIndexNumber_GetHash_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ServerIndexNumber_GetHash",
	})
	factomdentryCreditBlockServerIndexNumberGetEntryHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ServerIndexNumber_GetEntryHash_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ServerIndexNumber_GetEntryHash",
	})
	factomdentryCreditBlockServerIndexNumberGetSigHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ServerIndexNumber_GetSigHash_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ServerIndexNumber_GetSigHash",
	})
	factomdentryCreditBlockServerIndexNumberIsInterpretable = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ServerIndexNumber_IsInterpretable_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ServerIndexNumber_IsInterpretable",
	})
	factomdentryCreditBlockServerIndexNumberInterpret = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ServerIndexNumber_Interpret_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ServerIndexNumber_Interpret",
	})
	factomdentryCreditBlockNewServerIndexNumber = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_NewServerIndexNumber_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_NewServerIndexNumber",
	})
	factomdentryCreditBlockNewServerIndexNumber2 = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_NewServerIndexNumber2_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_NewServerIndexNumber2",
	})
	factomdentryCreditBlockServerIndexNumberECID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ServerIndexNumber_ECID_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ServerIndexNumber_ECID",
	})
	factomdentryCreditBlockServerIndexNumberMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ServerIndexNumber_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ServerIndexNumber_MarshalBinary",
	})
	factomdentryCreditBlockServerIndexNumberUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ServerIndexNumber_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ServerIndexNumber_UnmarshalBinaryData",
	})
	factomdentryCreditBlockServerIndexNumberUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ServerIndexNumber_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ServerIndexNumber_UnmarshalBinary",
	})
	factomdentryCreditBlockServerIndexNumberJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ServerIndexNumber_JSONByte_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ServerIndexNumber_JSONByte",
	})
	factomdentryCreditBlockServerIndexNumberJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ServerIndexNumber_JSONString_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ServerIndexNumber_JSONString",
	})
	factomdentryCreditBlockServerIndexNumberGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ServerIndexNumber_GetTimestamp_Summary",
		Help: "Summary of calls to  factomd_entryCreditBlock_ServerIndexNumber_GetTimestamp",
	})
)

func InitPrometheus() {
	prometheus.MustRegister(factomdentryCreditBlockCommitChainInit)
	prometheus.MustRegister(factomdentryCreditBlockCommitChainString)
	prometheus.MustRegister(factomdentryCreditBlockNewCommitChain)
	prometheus.MustRegister(factomdentryCreditBlockCommitChainGetEntryHash)
	prometheus.MustRegister(factomdentryCreditBlockCommitChainIsSameAs)
	prometheus.MustRegister(factomdentryCreditBlockCommitChainHash)
	prometheus.MustRegister(factomdentryCreditBlockCommitChainIsInterpretable)
	prometheus.MustRegister(factomdentryCreditBlockCommitChainInterpret)
	prometheus.MustRegister(factomdentryCreditBlockCommitChainCommitMsg)
	prometheus.MustRegister(factomdentryCreditBlockCommitChainGetTimestamp)
	prometheus.MustRegister(factomdentryCreditBlockCommitChainIsValid)
	prometheus.MustRegister(factomdentryCreditBlockCommitChainGetHash)
	prometheus.MustRegister(factomdentryCreditBlockCommitChainGetSigHash)
	prometheus.MustRegister(factomdentryCreditBlockCommitChainMarshalBinarySig)
	prometheus.MustRegister(factomdentryCreditBlockCommitChainMarshalBinaryTransaction)
	prometheus.MustRegister(factomdentryCreditBlockCommitChainMarshalBinary)
	prometheus.MustRegister(factomdentryCreditBlockCommitChainSign)
	prometheus.MustRegister(factomdentryCreditBlockCommitChainValidateSignatures)
	prometheus.MustRegister(factomdentryCreditBlockCommitChainECID)
	prometheus.MustRegister(factomdentryCreditBlockCommitChainUnmarshalBinaryData)
	prometheus.MustRegister(factomdentryCreditBlockCommitChainUnmarshalBinary)
	prometheus.MustRegister(factomdentryCreditBlockCommitChainJSONByte)
	prometheus.MustRegister(factomdentryCreditBlockCommitChainJSONString)
	prometheus.MustRegister(factomdentryCreditBlockCommitEntryString)
	prometheus.MustRegister(factomdentryCreditBlockCommitEntryGetEntryHash)
	prometheus.MustRegister(factomdentryCreditBlockCommitEntryIsSameAs)
	prometheus.MustRegister(factomdentryCreditBlockNewCommitEntry)
	prometheus.MustRegister(factomdentryCreditBlockCommitEntryHash)
	prometheus.MustRegister(factomdentryCreditBlockCommitEntryIsInterpretable)
	prometheus.MustRegister(factomdentryCreditBlockCommitEntryInterpret)
	prometheus.MustRegister(factomdentryCreditBlockCommitEntryCommitMsg)
	prometheus.MustRegister(factomdentryCreditBlockCommitEntryGetTimestamp)
	prometheus.MustRegister(factomdentryCreditBlockCommitEntryInTime)
	prometheus.MustRegister(factomdentryCreditBlockCommitEntryIsValid)
	prometheus.MustRegister(factomdentryCreditBlockCommitEntryGetHash)
	prometheus.MustRegister(factomdentryCreditBlockCommitEntryGetSigHash)
	prometheus.MustRegister(factomdentryCreditBlockCommitEntryMarshalBinarySig)
	prometheus.MustRegister(factomdentryCreditBlockCommitEntryMarshalBinaryTransaction)
	prometheus.MustRegister(factomdentryCreditBlockCommitEntryMarshalBinary)
	prometheus.MustRegister(factomdentryCreditBlockCommitEntrySign)
	prometheus.MustRegister(factomdentryCreditBlockCommitEntryValidateSignatures)
	prometheus.MustRegister(factomdentryCreditBlockCommitEntryECID)
	prometheus.MustRegister(factomdentryCreditBlockCommitEntryUnmarshalBinaryData)
	prometheus.MustRegister(factomdentryCreditBlockCommitEntryUnmarshalBinary)
	prometheus.MustRegister(factomdentryCreditBlockCommitEntryJSONByte)
	prometheus.MustRegister(factomdentryCreditBlockCommitEntryJSONString)
	prometheus.MustRegister(factomdentryCreditBlockECBlockInit)
	prometheus.MustRegister(factomdentryCreditBlockECBlockUpdateState)
	prometheus.MustRegister(factomdentryCreditBlockECBlockString)
	prometheus.MustRegister(factomdentryCreditBlockECBlockGetEntries)
	prometheus.MustRegister(factomdentryCreditBlockECBlockGetEntryByHash)
	prometheus.MustRegister(factomdentryCreditBlockECBlockGetEntryHashes)
	prometheus.MustRegister(factomdentryCreditBlockECBlockGetEntrySigHashes)
	prometheus.MustRegister(factomdentryCreditBlockECBlockGetBody)
	prometheus.MustRegister(factomdentryCreditBlockECBlockGetHeader)
	prometheus.MustRegister(factomdentryCreditBlockECBlockNew)
	prometheus.MustRegister(factomdentryCreditBlockECBlockGetDatabaseHeight)
	prometheus.MustRegister(factomdentryCreditBlockECBlockGetChainID)
	prometheus.MustRegister(factomdentryCreditBlockECBlockDatabasePrimaryIndex)
	prometheus.MustRegister(factomdentryCreditBlockECBlockDatabaseSecondaryIndex)
	prometheus.MustRegister(factomdentryCreditBlockECBlockAddEntry)
	prometheus.MustRegister(factomdentryCreditBlockECBlockGetHash)
	prometheus.MustRegister(factomdentryCreditBlockECBlockGetFullHash)
	prometheus.MustRegister(factomdentryCreditBlockECBlockHeaderHash)
	prometheus.MustRegister(factomdentryCreditBlockECBlockMarshalBinary)
	prometheus.MustRegister(factomdentryCreditBlockECBlockBuildHeader)
	prometheus.MustRegister(factomdentryCreditBlockUnmarshalECBlock)
	prometheus.MustRegister(factomdentryCreditBlockECBlockUnmarshalBinaryData)
	prometheus.MustRegister(factomdentryCreditBlockECBlockUnmarshalBinary)
	prometheus.MustRegister(factomdentryCreditBlockECBlockmarshalBodyBinary)
	prometheus.MustRegister(factomdentryCreditBlockECBlockunmarshalBodyBinaryData)
	prometheus.MustRegister(factomdentryCreditBlockECBlockunmarshalBodyBinary)
	prometheus.MustRegister(factomdentryCreditBlockECBlockJSONByte)
	prometheus.MustRegister(factomdentryCreditBlockECBlockJSONString)
	prometheus.MustRegister(factomdentryCreditBlockNewECBlock)
	prometheus.MustRegister(factomdentryCreditBlockNextECBlock)
	prometheus.MustRegister(factomdentryCreditBlockCheckBlockPairIntegrity)
	prometheus.MustRegister(factomdentryCreditBlockECBlockBodyGetEntries)
	prometheus.MustRegister(factomdentryCreditBlockECBlockBodyAddEntry)
	prometheus.MustRegister(factomdentryCreditBlockECBlockBodySetEntries)
	prometheus.MustRegister(factomdentryCreditBlockECBlockBodyJSONByte)
	prometheus.MustRegister(factomdentryCreditBlockECBlockBodyJSONString)
	prometheus.MustRegister(factomdentryCreditBlockECBlockBodyString)
	prometheus.MustRegister(factomdentryCreditBlockNewECBlockBody)
	prometheus.MustRegister(factomdentryCreditBlockECBlockHeaderInit)
	prometheus.MustRegister(factomdentryCreditBlockECBlockHeaderString)
	prometheus.MustRegister(factomdentryCreditBlockECBlockHeaderSetBodySize)
	prometheus.MustRegister(factomdentryCreditBlockECBlockHeaderGetBodySize)
	prometheus.MustRegister(factomdentryCreditBlockECBlockHeaderSetObjectCount)
	prometheus.MustRegister(factomdentryCreditBlockECBlockHeaderGetObjectCount)
	prometheus.MustRegister(factomdentryCreditBlockECBlockHeaderSetHeaderExpansionArea)
	prometheus.MustRegister(factomdentryCreditBlockECBlockHeaderGetHeaderExpansionArea)
	prometheus.MustRegister(factomdentryCreditBlockECBlockHeaderSetBodyHash)
	prometheus.MustRegister(factomdentryCreditBlockECBlockHeaderGetBodyHash)
	prometheus.MustRegister(factomdentryCreditBlockECBlockHeaderGetECChainID)
	prometheus.MustRegister(factomdentryCreditBlockECBlockHeaderSetPrevHeaderHash)
	prometheus.MustRegister(factomdentryCreditBlockECBlockHeaderGetPrevHeaderHash)
	prometheus.MustRegister(factomdentryCreditBlockECBlockHeaderSetPrevFullHash)
	prometheus.MustRegister(factomdentryCreditBlockECBlockHeaderGetPrevFullHash)
	prometheus.MustRegister(factomdentryCreditBlockECBlockHeaderSetDBHeight)
	prometheus.MustRegister(factomdentryCreditBlockECBlockHeaderGetDBHeight)
	prometheus.MustRegister(factomdentryCreditBlockNewECBlockHeader)
	prometheus.MustRegister(factomdentryCreditBlockECBlockHeaderJSONByte)
	prometheus.MustRegister(factomdentryCreditBlockECBlockHeaderJSONString)
	prometheus.MustRegister(factomdentryCreditBlockECBlockHeaderMarshalBinary)
	prometheus.MustRegister(factomdentryCreditBlockECBlockHeaderUnmarshalBinaryData)
	prometheus.MustRegister(factomdentryCreditBlockECBlockHeaderUnmarshalBinary)
	prometheus.MustRegister(factomdentryCreditBlockECBlockHeaderMarshalJSON)
	prometheus.MustRegister(factomdentryCreditBlockIncreaseBalanceInit)
	prometheus.MustRegister(factomdentryCreditBlockIncreaseBalanceString)
	prometheus.MustRegister(factomdentryCreditBlockNewIncreaseBalance)
	prometheus.MustRegister(factomdentryCreditBlockIncreaseBalanceGetEntryHash)
	prometheus.MustRegister(factomdentryCreditBlockIncreaseBalanceHash)
	prometheus.MustRegister(factomdentryCreditBlockIncreaseBalanceGetHash)
	prometheus.MustRegister(factomdentryCreditBlockIncreaseBalanceGetSigHash)
	prometheus.MustRegister(factomdentryCreditBlockIncreaseBalanceECID)
	prometheus.MustRegister(factomdentryCreditBlockIncreaseBalanceIsInterpretable)
	prometheus.MustRegister(factomdentryCreditBlockIncreaseBalanceInterpret)
	prometheus.MustRegister(factomdentryCreditBlockIncreaseBalanceMarshalBinary)
	prometheus.MustRegister(factomdentryCreditBlockIncreaseBalanceUnmarshalBinaryData)
	prometheus.MustRegister(factomdentryCreditBlockIncreaseBalanceUnmarshalBinary)
	prometheus.MustRegister(factomdentryCreditBlockIncreaseBalanceJSONByte)
	prometheus.MustRegister(factomdentryCreditBlockIncreaseBalanceJSONString)
	prometheus.MustRegister(factomdentryCreditBlockIncreaseBalanceGetTimestamp)
	prometheus.MustRegister(factomdentryCreditBlockMinuteNumberHash)
	prometheus.MustRegister(factomdentryCreditBlockMinuteNumberGetHash)
	prometheus.MustRegister(factomdentryCreditBlockMinuteNumberGetSigHash)
	prometheus.MustRegister(factomdentryCreditBlockMinuteNumberGetEntryHash)
	prometheus.MustRegister(factomdentryCreditBlockMinuteNumberIsInterpretable)
	prometheus.MustRegister(factomdentryCreditBlockMinuteNumberInterpret)
	prometheus.MustRegister(factomdentryCreditBlockNewMinuteNumber)
	prometheus.MustRegister(factomdentryCreditBlockMinuteNumberECID)
	prometheus.MustRegister(factomdentryCreditBlockMinuteNumberMarshalBinary)
	prometheus.MustRegister(factomdentryCreditBlockMinuteNumberUnmarshalBinaryData)
	prometheus.MustRegister(factomdentryCreditBlockMinuteNumberUnmarshalBinary)
	prometheus.MustRegister(factomdentryCreditBlockMinuteNumberJSONByte)
	prometheus.MustRegister(factomdentryCreditBlockMinuteNumberJSONString)
	prometheus.MustRegister(factomdentryCreditBlockMinuteNumberString)
	prometheus.MustRegister(factomdentryCreditBlockMinuteNumberGetTimestamp)
	prometheus.MustRegister(factomdentryCreditBlockServerIndexNumberString)
	prometheus.MustRegister(factomdentryCreditBlockServerIndexNumberHash)
	prometheus.MustRegister(factomdentryCreditBlockServerIndexNumberGetHash)
	prometheus.MustRegister(factomdentryCreditBlockServerIndexNumberGetEntryHash)
	prometheus.MustRegister(factomdentryCreditBlockServerIndexNumberGetSigHash)
	prometheus.MustRegister(factomdentryCreditBlockServerIndexNumberIsInterpretable)
	prometheus.MustRegister(factomdentryCreditBlockServerIndexNumberInterpret)
	prometheus.MustRegister(factomdentryCreditBlockNewServerIndexNumber)
	prometheus.MustRegister(factomdentryCreditBlockNewServerIndexNumber2)
	prometheus.MustRegister(factomdentryCreditBlockServerIndexNumberECID)
	prometheus.MustRegister(factomdentryCreditBlockServerIndexNumberMarshalBinary)
	prometheus.MustRegister(factomdentryCreditBlockServerIndexNumberUnmarshalBinaryData)
	prometheus.MustRegister(factomdentryCreditBlockServerIndexNumberUnmarshalBinary)
	prometheus.MustRegister(factomdentryCreditBlockServerIndexNumberJSONByte)
	prometheus.MustRegister(factomdentryCreditBlockServerIndexNumberJSONString)
	prometheus.MustRegister(factomdentryCreditBlockServerIndexNumberGetTimestamp)
}
