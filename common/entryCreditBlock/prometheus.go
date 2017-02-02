// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package entryCreditBlock

import "github.com/prometheus/client_golang/prometheus"

var (
	//commitChain variables
	entryCreditBlockErrors = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "factomd_ferEntry_error_responses",
		Help: "Number of error responses from ferEntry objects",
	})
	entryCreditBlockCommitChainString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_commitChain_String_Summary",
		Help: "Summary of entryCreditBlock.commitChain.String call",
	})
	entryCreditBlockCommitChainNewCommitChain = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_commitChain_NewCommitChain_Summary",
		Help: "Summary of entryCreditBlock.commitChain.NewCommitChain call",
	})
	entryCreditBlockCommitChainGetEntryHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_commitChain_GetEntryHash_Summary",
		Help: "Summary of entryCreditBlock.commitChain.GetEntryHash call",
	})
	entryCreditBlockCommitChainIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_commitChain_IsSameAs_Summary",
		Help: "Summary of entryCreditBlock.commitChain.IsSameAs call",
	})
	entryCreditBlockCommitChainHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_commitChain_Hash_Summary",
		Help: "Summary of entryCreditBlock.commitChain.Hash call",
	})
	entryCreditBlockCommitChainCommitMsg = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_commitChain_CommitMsg_Summary",
		Help: "Summary of entryCreditBlock.commitChain.CommitMsg call",
	})
	entryCreditBlockCommitChainGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_commitChain_GetTimestamp_Summary",
		Help: "Summary of entryCreditBlock.commitChain.GetTimestamp call",
	})
	entryCreditBlockCommitChainIsValid = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_commitChain_IsValid_Summary",
		Help: "Summary of entryCreditBlock.commitChain.IsValid call",
	})
	entryCreditBlockCommitChainGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_commitChain_GetHash_Summary",
		Help: "Summary of entryCreditBlock.commitChain.GetHash call",
	})
	entryCreditBlockCommitChainMarshalBinarySig = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_commitChain_MarshalBinarySig_Summary",
		Help: "Summary of entryCreditBlock.commitChain.MarshalBinarySig call",
	})
	entryCreditBlockCommitChainMarshalBinaryTransaction = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_commitChain_MarshalBinaryTransaction_Summary",
		Help: "Summary of entryCreditBlock.commitChain.MarshalBinaryTransaction call",
	})
	entryCreditBlockCommitChainMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_commitChain_MarshalBinary_Summary",
		Help: "Summary of entryCreditBlock.commitChain.MarshalBinary call",
	})
	entryCreditBlockCommitChainSign = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_commitChain_Sign_Summary",
		Help: "Summary of entryCreditBlock.commitChain.Sign call",
	})
	entryCreditBlockCommitChainValidateSignatures = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_commitChain_ValidateSignatures_Summary",
		Help: "Summary of entryCreditBlock.commitChain.ValidateSignatures call",
	})
	entryCreditBlockCommitChainECID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_commitChain_ECID_Summary",
		Help: "Summary of entryCreditBlock.commitChain.ECID call",
	})
	entryCreditBlockCommitChainUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_commitChain_UnmarshalBinaryData_Summary",
		Help: "Summary of entryCreditBlock.commitChain.UnmarshalBinaryData call",
	})
	entryCreditBlockCommitChainUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_commitChain_UnmarshalBinary_Summary",
		Help: "Summary of entryCreditBlock.commitChain.UnmarshalBinary call",
	})
	entryCreditBlockCommitChainJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_commitChain_JSONByte_Summary",
		Help: "Summary of entryCreditBlock.commitChain.JSONByte call",
	})
	entryCreditBlockCommitChainJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_commitChain_JSONString_Summary",
		Help: "Summary of entryCreditBlock.commitChain.JSONString call",
	})
	entryCreditBlockCommitChainJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_commitChain_JSONBuffer_Summary",
		Help: "Summary of entryCreditBlock.commitChain.JSONBuffer call",
	})
	//c commit entry variables
	entryCreditBlockCommitEntryString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_commitEntry_String_Summary",
		Help: "Summary of entryCreditBlock.commitEntry.String call",
	})
	entryCreditBlockCommitEntryGetEntryHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_commitEntry_GetEntryHash_Summary",
		Help: "Summary of entryCreditBlock.commitEntry.GetEntryHash call",
	})
	entryCreditBlockCommitEntryIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_commitEntry_IsSameAs_Summary",
		Help: "Summary of entryCreditBlock.commitEntry.IsSameAs call",
	})
	entryCreditBlockCommitEntryNewCommitEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_commitEntry_NewCommitEntry_Summary",
		Help: "Summary of entryCreditBlock.commitEntry.NewCommitEntry call",
	})
	entryCreditBlockCommitEntryHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_commitEntry_Hash_Summary",
		Help: "Summary of entryCreditBlock.commitEntry.Hash call",
	})
	entryCreditBlockCommitEntryCommitMsg = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_commitEntry_CommitMsg_Summary",
		Help: "Summary of entryCreditBlock.commitEntry.CommitMsg call",
	})
	entryCreditBlockCommitEntryGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_commitEntry_GetTimestamp_Summary",
		Help: "Summary of entryCreditBlock.commitEntry.GetTimestamp call",
	})
	entryCreditBlockCommitEntryInTime = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_commitEntry_InTime_Summary",
		Help: "Summary of entryCreditBlock.commitEntry.InTime call",
	})
	entryCreditBlockCommitEntryIsValid = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_commitEntry_IsValid_Summary",
		Help: "Summary of entryCreditBlock.commitEntry.IsValid call",
	})
	entryCreditBlockCommitEntryGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_commitEntry_GetHash_Summary",
		Help: "Summary of entryCreditBlock.commitEntry.GetHash call",
	})
	entryCreditBlockCommitEntryGetSigHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_commitEntry_GetSigHash_Summary",
		Help: "Summary of entryCreditBlock.commitEntry.GetSigHash call",
	})
	entryCreditBlockCommitEntryMarshalBinarySig = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_commitEntry_MarshalBinarySig_Summary",
		Help: "Summary of entryCreditBlock.commitEntry.MarshalBinarySig call",
	})

	entryCreditBlockCommitEntryMarshalBinaryTransaction = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_commitEntry_MarshalBinaryTransaction_Summary",
		Help: "Summary of entryCreditBlock.commitEntry.MarshalBinaryTransaction call",
	})
	entryCreditBlockCommitEntryMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_commitEntry_MarshalBinary_Summary",
		Help: "Summary of entryCreditBlock.commitEntry.MarshalBinary call",
	})
	entryCreditBlockCommitEntrySign = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_commitEntry_Sign_Summary",
		Help: "Summary of entryCreditBlock.commitEntry.Sign call",
	})
	entryCreditBlockCommitEntryValidateSignatures = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_commitEntry_ValidateSignatures_Summary",
		Help: "Summary of entryCreditBlock.commitEntry.ValidateSignatures call",
	})
	entryCreditBlockCommitEntryECID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_commitEntry_ECID_Summary",
		Help: "Summary of entryCreditBlock.commitEntry.ECID call",
	})
	entryCreditBlockCommitEntryUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_commitEntry_UnmarshalBinaryData_Summary",
		Help: "Summary of entryCreditBlock.commitEntry.UnmarshalBinaryData call",
	})
	entryCreditBlockCommitEntryUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_commitEntry_UnmarshalBinary_Summary",
		Help: "Summary of entryCreditBlock.commitEntry.UnmarshalBinary call",
	})
	entryCreditBlockCommitEntryJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_commitEntry_JSONByte_Summary",
		Help: "Summary of entryCreditBlock.commitEntry.JSONByte call",
	})
	entryCreditBlockCommitEntryJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_commitEntry_JSONString_Summary",
		Help: "Summary of entryCreditBlock.commitEntry.JSONString call",
	})
	entryCreditBlockCommitEntryJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_commitEntry_JSONBuffer_Summary",
		Help: "Summary of entryCreditBlock.commitEntry.JSONBuffer call",
	})
	// ecblock.go variables
	entryCreditBlockecblockUpdateState = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblock_UpdateState_Summary",
		Help: "Summary of entryCreditBlock.ecblock.UpdateState call",
	})
	entryCreditBlockecblockString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblock_String_Summary",
		Help: "Summary of entryCreditBlock.ecblock.String call",
	})
	entryCreditBlockecblockGetEntries = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblock_GetEntries_Summary",
		Help: "Summary of entryCreditBlock.ecblock.GetEntries call",
	})
	entryCreditBlockecblockGetEntryByHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblock_GetEntryByHash_Summary",
		Help: "Summary of entryCreditBlock.ecblock.GetEntryByHash call",
	})
	entryCreditBlockecblockGetEntryHashes = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblock_GetEntryHashes_Summary",
		Help: "Summary of entryCreditBlock.ecblock.GetEntryHashes call",
	})
	entryCreditBlockecblockGetEntrySigHashes = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblock_GetEntrySigHashes_Summary",
		Help: "Summary of entryCreditBlock.ecblock.GetEntrySigHashes call",
	})
	entryCreditBlockecblockGetBody = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblock_GetBody_Summary",
		Help: "Summary of entryCreditBlock.ecblock.GetBody call",
	})
	entryCreditBlockecblockGetHeader = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblock_GetHeader_Summary",
		Help: "Summary of entryCreditBlock.ecblock.GetHeader call",
	})
	entryCreditBlockecblockNew = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblock_New_Summary",
		Help: "Summary of entryCreditBlock.ecblock.New call",
	})
	entryCreditBlockecblockGetDatabaseHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblock_GetDatabaseHeight_Summary",
		Help: "Summary of entryCreditBlock.ecblock.GetDatabaseHeight call",
	})
	entryCreditBlockecblockGetChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblock_GetChainID_Summary",
		Help: "Summary of entryCreditBlock.ecblock.GetChainID call",
	})
	entryCreditBlockecblockDatabasePrimaryIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblock_DatabasePrimaryIndex_Summary",
		Help: "Summary of entryCreditBlock.ecblock.DatabasePrimaryIndex call",
	})
	entryCreditBlockecblockDatabaseSecondaryIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblock_DatabaseSecondaryIndex_Summary",
		Help: "Summary of entryCreditBlock.ecblock.DatabaseSecondaryIndex call",
	})
	entryCreditBlockecblockAddEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblock_AddEntry_Summary",
		Help: "Summary of entryCreditBlock.ecblock.AddEntry call",
	})
	entryCreditBlockecblockGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblock_GetHash_Summary",
		Help: "Summary of entryCreditBlock.ecblock.GetHash call",
	})
	entryCreditBlockecblockGetFullHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblock_GetFullHash_Summary",
		Help: "Summary of entryCreditBlock.ecblock.GetFullHash call",
	})
	entryCreditBlockecblockHeaderHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblock_HeaderHash_Summary",
		Help: "Summary of entryCreditBlock.ecblock.HeaderHash call",
	})
	entryCreditBlockecblockMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblock_MarshalBinary_Summary",
		Help: "Summary of entryCreditBlock.ecblock.MarshalBinary call",
	})
	entryCreditBlockecblockBuildHeader = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblock_BuildHeader_Summary",
		Help: "Summary of entryCreditBlock.ecblock.BuildHeader call",
	})
	entryCreditBlockecblockUnmarshalECBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblock_UnmarshalECBlock_Summary",
		Help: "Summary of entryCreditBlock.ecblock.UnmarshalECBlock call",
	})
	entryCreditBlockecblockUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblock_UnmarshalBinaryData_Summary",
		Help: "Summary of entryCreditBlock.ecblock.UnmarshalBinaryData call",
	})
	entryCreditBlockecblockUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblock_UnmarshalBinary_Summary",
		Help: "Summary of entryCreditBlock.ecblock.UnmarshalBinary call",
	})
	entryCreditBlockecblockmarshalHeaderBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblock_marshalHeaderBinary_Summary",
		Help: "Summary of entryCreditBlock.ecblock.marshalHeaderBinary call",
	})
	entryCreditBlockecblockunmarshalBodyBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblock_unmarshalBodyBinaryData_Summary",
		Help: "Summary of entryCreditBlock.ecblock.unmarshalBodyBinaryData call",
	})
	entryCreditBlockecblockunmarshalBodyBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblock_unmarshalBodyBinary_Summary",
		Help: "Summary of entryCreditBlock.ecblock.unmarshalBodyBinary call",
	})
	entryCreditBlockecblockunmarshalHeaderBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblock_unmarshalHeaderBinary_Summary",
		Help: "Summary of entryCreditBlock.ecblock.unmarshalHeaderBinary call",
	})
	entryCreditBlockecblockJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblock_JSONByte_Summary",
		Help: "Summary of entryCreditBlock.ecblock.JSONByte call",
	})
	entryCreditBlockecblockJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblock_JSONString_Summary",
		Help: "Summary of entryCreditBlock.ecblock.JSONString call",
	})
	entryCreditBlockecblockJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblock_JSONBuffer_Summary",
		Help: "Summary of entryCreditBlock.ecblock.JSONBuffer call",
	})
	entryCreditBlockecblockNewECBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblock_NewECBlock_Summary",
		Help: "Summary of entryCreditBlock.ecblock.NewECBlock call",
	})
	entryCreditBlockecblockNextECBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblock_NextECBlock_Summary",
		Help: "Summary of entryCreditBlock.ecblock.NextECBlock call",
	})
	entryCreditBlockecblockCheckBlockPairIntegrity = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblock_CheckBlockPairIntegrity_Summary",
		Help: "Summary of entryCreditBlock.ecblock.CheckBlockPairIntegrity call",
	})
	// ecblockbody variables
	entryCreditBlockecblockBodyGetEntries = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblockBody_GetEntries_Summary",
		Help: "Summary of entryCreditBlock.ecblockBody.GetEntries call",
	})
	entryCreditBlockecblockBodyAddEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblockBody_AddEntry_Summary",
		Help: "Summary of entryCreditBlock.ecblockBody.AddEntry call",
	})
	entryCreditBlockecblockBodySetEntries = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblockBody_SetEntries_Summary",
		Help: "Summary of entryCreditBlock.ecblockBody.SetEntries call",
	})
	entryCreditBlockecblockBodyJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblockBody_JSONByte_Summary",
		Help: "Summary of entryCreditBlock.ecblockBody.JSONByte call",
	})
	entryCreditBlockecblockBodyJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblockBody_JSONString_Summary",
		Help: "Summary of entryCreditBlock.ecblockBody.JSONString call",
	})
	entryCreditBlockecblockBodyJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblockBody_JSONBuffer_Summary",
		Help: "Summary of entryCreditBlock.ecblockBody.JSONBuffer call",
	})
	entryCreditBlockecblockBodyString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblockBody_String_Summary",
		Help: "Summary of entryCreditBlock.ecblockBody.String call",
	})
	entryCreditBlockecblockBodyNewECBlockBody = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblockBody_NewECBlockBody_Summary",
		Help: "Summary of entryCreditBlock.ecblockBody.NewECBlockBody call",
	})
	// ecblockHead variables
	entryCreditBlockecblockHeadString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblockHead_String_Summary",
		Help: "Summary of entryCreditBlock.ecblockHead.String call",
	})
	entryCreditBlockecblockHeadSetBodySize = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblockHead_SetBodySize_Summary",
		Help: "Summary of entryCreditBlock.ecblockHead.SetBodySize call",
	})
	entryCreditBlockecblockHeadGetBodySize = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblockHead_GetBodySize_Summary",
		Help: "Summary of entryCreditBlock.ecblockHead.GetBodySize call",
	})
	entryCreditBlockecblockHeadSetObjectCount = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblockHead_SetObjectCount_Summary",
		Help: "Summary of entryCreditBlock.ecblockHead.SetObjectCount call",
	})
	entryCreditBlockecblockHeadGetObjectCount = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblockHead_GetObjectCount_Summary",
		Help: "Summary of entryCreditBlock.ecblockHead.GetObjectCount call",
	})
	entryCreditBlockecblockHeadSetHeaderExpansionArea = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblockHead_SetHeaderExpansionArea_Summary",
		Help: "Summary of entryCreditBlock.ecblockHead.SetHeaderExpansionArea call",
	})
	entryCreditBlockecblockHeadGetHeaderExpansionArea = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblockHead_GetHeaderExpansionArea_Summary",
		Help: "Summary of entryCreditBlock.ecblockHead.GetHeaderExpansionArea call",
	})
	entryCreditBlockecblockHeadSetBodyHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblockHead_SetBodyHash_Summary",
		Help: "Summary of entryCreditBlock.ecblockHead.SetBodyHash call",
	})
	entryCreditBlockecblockHeadGetBodyHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblockHead_GetBodyHash_Summary",
		Help: "Summary of entryCreditBlock.ecblockHead.GetBodyHash call",
	})
	entryCreditBlockecblockHeadGetECChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblockHead_GetECChainID_Summary",
		Help: "Summary of entryCreditBlock.ecblockHead.GetECChainID call",
	})
	entryCreditBlockecblockHeadSetPrevHeaderHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblockHead_SetPrevHeaderHash_Summary",
		Help: "Summary of entryCreditBlock.ecblockHead.SetPrevHeaderHash call",
	})
	entryCreditBlockecblockHeadGetPrevHeaderHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblockHead_GetPrevHeaderHash_Summary",
		Help: "Summary of entryCreditBlock.ecblockHead.GetPrevHeaderHash call",
	})
	entryCreditBlockecblockHeadSetPrevFullHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblockHead_SetPrevFullHash_Summary",
		Help: "Summary of entryCreditBlock.ecblockHead.SetPrevFullHash call",
	})
	entryCreditBlockecblockHeadGetPrevFullHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblockHead_GetPrevFullHash_Summary",
		Help: "Summary of entryCreditBlock.ecblockHead.GetPrevFullHash call",
	})
	entryCreditBlockecblockHeadGetDBHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblockHead_GetDBHeight_Summary",
		Help: "Summary of entryCreditBlock.ecblockHead.GetDBHeight call",
	})
	entryCreditBlockecblockHeadSetDBHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblockHead_SetDBHeight_Summary",
		Help: "Summary of entryCreditBlock.ecblockHead.SetDBHeight call",
	})
	entryCreditBlockecblockHeadNewECBlockHeader = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblockHead_NewECBlockHeader_Summary",
		Help: "Summary of entryCreditBlock.ecblockHead.NewECBlockHeader call",
	})
	entryCreditBlockecblockHeadJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblockHead_JSONByte_Summary",
		Help: "Summary of entryCreditBlock.ecblockHead.JSONByte call",
	})
	entryCreditBlockecblockHeadJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblockHead_JSONString_Summary",
		Help: "Summary of entryCreditBlock.ecblockHead.JSONString call",
	})
	entryCreditBlockecblockHeadJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblockHead_JSONBuffer_Summary",
		Help: "Summary of entryCreditBlock.ecblockHead.JSONBuffer call",
	})
	entryCreditBlockecblockHeadMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblockHead_MarshalBinary_Summary",
		Help: "Summary of entryCreditBlock.ecblockHead.MarshalBinary call",
	})
	entryCreditBlockecblockHeadUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblockHead_UnmarshalBinaryData_Summary",
		Help: "Summary of entryCreditBlock.ecblockHead.UnmarshalBinaryData call",
	})
	entryCreditBlockecblockHeadUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblockHead_UnmarshalBinary_Summary",
		Help: "Summary of entryCreditBlock.ecblockHead.UnmarshalBinary call",
	})
	entryCreditBlockecblockHeadMarshalJSON = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_ecblockHead_MarshalJSON_Summary",
		Help: "Summary of entryCreditBlock.ecblockHead.MarshalJSON call",
	})
	// increaseBalance variables
	entryCreditBlockincreaseBalanceString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_increaseBalance_String_Summary",
		Help: "Summary of entryCreditBlock.increaseBalance.String call",
	})
	entryCreditBlockincreaseBalanceNewIncreaseBalance = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_increaseBalance_NewIncreaseBalance_Summary",
		Help: "Summary of entryCreditBlock.increaseBalance.NewIncreaseBalance call",
	})
	entryCreditBlockincreaseBalanceHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_increaseBalance_Hash_Summary",
		Help: "Summary of entryCreditBlock.increaseBalance.Hash call",
	})
	entryCreditBlockincreaseBalanceGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_increaseBalance_GetHash_Summary",
		Help: "Summary of entryCreditBlock.increaseBalance.GetHash call",
	})
	entryCreditBlockincreaseBalanceECID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_increaseBalance_ECID_Summary",
		Help: "Summary of entryCreditBlock.increaseBalance.ECID call",
	})
	entryCreditBlockincreaseBalanceMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_increaseBalance_MarshalBinary_Summary",
		Help: "Summary of entryCreditBlock.increaseBalance.MarshalBinary call",
	})
	entryCreditBlockincreaseBalanceUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_increaseBalance_UnmarshalBinaryData_Summary",
		Help: "Summary of entryCreditBlock.increaseBalance.UnmarshalBinaryData call",
	})
	entryCreditBlockincreaseBalanceUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_increaseBalance_UnmarshalBinary_Summary",
		Help: "Summary of entryCreditBlock.increaseBalance.UnmarshalBinary call",
	})
	entryCreditBlockincreaseBalanceJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_increaseBalance_JSONByte_Summary",
		Help: "Summary of entryCreditBlock.increaseBalance.JSONByte call",
	})
	entryCreditBlockincreaseBalanceJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_increaseBalance_JSONString_Summary",
		Help: "Summary of entryCreditBlock.increaseBalance.JSONString call",
	})
	entryCreditBlockincreaseBalanceJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_increaseBalance_JSONBuffer_Summary",
		Help: "Summary of entryCreditBlock.increaseBalance.JSONBuffer call",
	})
	entryCreditBlockincreaseBalanceGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_increaseBalance_GetTimestamp_Summary",
		Help: "Summary of entryCreditBlock.increaseBalance.GetTimestamp call",
	})
	// minuteNumber variables
	entryCreditBlockminuteNumberHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_minuteNumber_Hash_Summary",
		Help: "Summary of entryCreditBlock.minuteNumber.Hash call",
	})
	entryCreditBlockminuteNumberGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_minuteNumber_GetHash_Summary",
		Help: "Summary of entryCreditBlock.minuteNumber.GetHash call",
	})
	entryCreditBlockminuteNumberGetSigHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_minuteNumber_GetSigHash_Summary",
		Help: "Summary of entryCreditBlock.minuteNumber.GetSigHash call",
	})
	entryCreditBlockminuteNumberGetEntryHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_minuteNumber_GetEntryHash_Summary",
		Help: "Summary of entryCreditBlock.minuteNumber.GetEntryHash call",
	})
	entryCreditBlockminuteNumberIsInterpretable = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_minuteNumber_IsInterpretable_Summary",
		Help: "Summary of entryCreditBlock.minuteNumber.IsInterpretable call",
	})
	entryCreditBlockminuteNumberInterpret = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_minuteNumber_Interpret_Summary",
		Help: "Summary of entryCreditBlock.minuteNumber.Interpret call",
	})
	entryCreditBlockminuteNumberNewMinuteNumber = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_minuteNumber_NewMinuteNumber_Summary",
		Help: "Summary of entryCreditBlock.minuteNumber.NewMinuteNumber call",
	})
	entryCreditBlockminuteNumberECID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_minuteNumber_ECID_Summary",
		Help: "Summary of entryCreditBlock.minuteNumber.ECID call",
	})
	entryCreditBlockminuteNumberMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_minuteNumber_MarshalBinary_Summary",
		Help: "Summary of entryCreditBlock.minuteNumber.MarshalBinary call",
	})
	entryCreditBlockminuteNumberUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_minuteNumber_UnmarshalBinaryData_Summary",
		Help: "Summary of entryCreditBlock.minuteNumber.UnmarshalBinaryData call",
	})
	entryCreditBlockminuteNumberUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_minuteNumber_UnmarshalBinary_Summary",
		Help: "Summary of entryCreditBlock.minuteNumber.UnmarshalBinary call",
	})
	entryCreditBlockminuteNumberJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_minuteNumber_JSONByte_Summary",
		Help: "Summary of entryCreditBlock.minuteNumber.JSONByte call",
	})
	entryCreditBlockminuteNumberJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_minuteNumber_JSONString_Summary",
		Help: "Summary of entryCreditBlock.minuteNumber.JSONString call",
	})
	entryCreditBlockminuteNumberJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_minuteNumber_JSONBuffer_Summary",
		Help: "Summary of entryCreditBlock.minuteNumber.JSONBuffer call",
	})
	entryCreditBlockminuteNumberString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_minuteNumber_String_Summary",
		Help: "Summary of entryCreditBlock.minuteNumber.String call",
	})
	//serverindexnumber.go variables
	entryCreditBlockServerIndexNumberString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_serverindexnumber_String_Summary",
		Help: "Summary of entryCreditBlock.serverindexnumber.String call",
	})
	entryCreditBlockServerIndexNumberHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_serverindexnumber_Hash_Summary",
		Help: "Summary of entryCreditBlock.serverindexnumber.Hash call",
	})
	entryCreditBlockServerIndexNumberGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_serverindexnumber_GetHash_Summary",
		Help: "Summary of entryCreditBlock.serverindexnumber.GetHash call",
	})
	entryCreditBlockServerIndexNumberInterpret = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_serverindexnumber_Interpret_Summary",
		Help: "Summary of entryCreditBlock.serverindexnumber.Interpret call",
	})
	entryCreditBlockServerIndexNumberNewServerIndexNumber = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_serverindexnumber_NewServerIndexNumber_Summary",
		Help: "Summary of entryCreditBlock.serverindexnumber.NewServerIndexNumber call",
	})
	entryCreditBlockServerIndexNumberNewServerIndexNumber2 = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_serverindexnumber_NewServerIndexNumber2_Summary",
		Help: "Summary of entryCreditBlock.serverindexnumber.NewServerIndexNumber2 call",
	})
	entryCreditBlockServerIndexNumberECID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_serverindexnumber_ECID_Summary",
		Help: "Summary of entryCreditBlock.serverindexnumber.ECID call",
	})
	entryCreditBlockServerIndexNumberMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_serverindexnumber_MarshalBinary_Summary",
		Help: "Summary of entryCreditBlock.serverindexnumber.MarshalBinary call",
	})
	entryCreditBlockServerIndexNumberUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_serverindexnumber_UnmarshalBinaryData_Summary",
		Help: "Summary of entryCreditBlock.serverindexnumber.UnmarshalBinaryData call",
	})
	entryCreditBlockServerIndexNumberUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_serverindexnumber_UnmarshalBinary_Summary",
		Help: "Summary of entryCreditBlock.serverindexnumber.UnmarshalBinary call",
	})
	entryCreditBlockServerIndexNumberJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_serverindexnumber_JSONByte_Summary",
		Help: "Summary of entryCreditBlock.serverindexnumber.JSONByte call",
	})
	entryCreditBlockServerIndexNumberJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_serverindexnumber_JSONString_Summary",
		Help: "Summary of entryCreditBlock.serverindexnumber.JSONString call",
	})
	entryCreditBlockServerIndexNumberJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryCreditBlock_serverindexnumber_JSONBuffer_Summary",
		Help: "Summary of entryCreditBlock.serverindexnumber.JSONBuffer call",
	})
)

func InitPrometheus() {
	//commitChain variables
	prometheus.MustRegister(entryCreditBlockErrors)
	prometheus.MustRegister(entryCreditBlockCommitChainString)
	prometheus.MustRegister(entryCreditBlockCommitChainNewCommitChain)
	prometheus.MustRegister(entryCreditBlockCommitChainGetEntryHash)
	prometheus.MustRegister(entryCreditBlockCommitChainIsSameAs)
	prometheus.MustRegister(entryCreditBlockCommitChainHash)
	prometheus.MustRegister(entryCreditBlockCommitChainCommitMsg)
	prometheus.MustRegister(entryCreditBlockCommitChainGetTimestamp)
	prometheus.MustRegister(entryCreditBlockCommitChainIsValid)
	prometheus.MustRegister(entryCreditBlockCommitChainGetHash)
	prometheus.MustRegister(entryCreditBlockCommitChainMarshalBinarySig)
	prometheus.MustRegister(entryCreditBlockCommitChainMarshalBinaryTransaction)
	prometheus.MustRegister(entryCreditBlockCommitChainMarshalBinary)
	prometheus.MustRegister(entryCreditBlockCommitChainSign)
	prometheus.MustRegister(entryCreditBlockCommitChainValidateSignatures)
	prometheus.MustRegister(entryCreditBlockCommitChainECID)
	prometheus.MustRegister(entryCreditBlockCommitChainUnmarshalBinaryData)
	prometheus.MustRegister(entryCreditBlockCommitChainUnmarshalBinary)
	prometheus.MustRegister(entryCreditBlockCommitChainJSONByte)
	prometheus.MustRegister(entryCreditBlockCommitChainJSONString)
	prometheus.MustRegister(entryCreditBlockCommitChainJSONBuffer)
	prometheus.MustRegister(entryCreditBlockCommitEntryString)
	prometheus.MustRegister(entryCreditBlockCommitEntryGetEntryHash)
	prometheus.MustRegister(entryCreditBlockCommitEntryIsSameAs)
	prometheus.MustRegister(entryCreditBlockCommitEntryNewCommitEntry)
	prometheus.MustRegister(entryCreditBlockCommitEntryHash)
	prometheus.MustRegister(entryCreditBlockCommitEntryCommitMsg)
	prometheus.MustRegister(entryCreditBlockCommitEntryGetTimestamp)
	prometheus.MustRegister(entryCreditBlockCommitEntryInTime)
	prometheus.MustRegister(entryCreditBlockCommitEntryIsValid)
	prometheus.MustRegister(entryCreditBlockCommitEntryGetHash)
	prometheus.MustRegister(entryCreditBlockCommitEntryGetSigHash)
	prometheus.MustRegister(entryCreditBlockCommitEntryMarshalBinarySig)
	prometheus.MustRegister(entryCreditBlockCommitEntryMarshalBinaryTransaction)
	prometheus.MustRegister(entryCreditBlockCommitEntryMarshalBinary)
	prometheus.MustRegister(entryCreditBlockCommitEntrySign)
	prometheus.MustRegister(entryCreditBlockCommitEntryValidateSignatures)
	prometheus.MustRegister(entryCreditBlockCommitEntryECID)
	prometheus.MustRegister(entryCreditBlockCommitEntryUnmarshalBinaryData)
	prometheus.MustRegister(entryCreditBlockCommitEntryUnmarshalBinary)
	prometheus.MustRegister(entryCreditBlockCommitEntryJSONByte)
	prometheus.MustRegister(entryCreditBlockCommitEntryJSONString)
	prometheus.MustRegister(entryCreditBlockCommitEntryJSONBuffer)
	prometheus.MustRegister(entryCreditBlockecblockUpdateState)
	prometheus.MustRegister(entryCreditBlockecblockString)
	prometheus.MustRegister(entryCreditBlockecblockGetEntries)
	prometheus.MustRegister(entryCreditBlockecblockGetEntryByHash)
	prometheus.MustRegister(entryCreditBlockecblockGetEntryHashes)
	prometheus.MustRegister(entryCreditBlockecblockGetEntrySigHashes)
	prometheus.MustRegister(entryCreditBlockecblockGetBody)
	prometheus.MustRegister(entryCreditBlockecblockGetHeader)
	prometheus.MustRegister(entryCreditBlockecblockNew)
	prometheus.MustRegister(entryCreditBlockecblockGetDatabaseHeight)
	prometheus.MustRegister(entryCreditBlockecblockGetChainID)
	prometheus.MustRegister(entryCreditBlockecblockDatabasePrimaryIndex)
	prometheus.MustRegister(entryCreditBlockecblockDatabaseSecondaryIndex)
	prometheus.MustRegister(entryCreditBlockecblockAddEntry)
	prometheus.MustRegister(entryCreditBlockecblockGetHash)
	prometheus.MustRegister(entryCreditBlockecblockGetFullHash)
	prometheus.MustRegister(entryCreditBlockecblockHeaderHash)
	prometheus.MustRegister(entryCreditBlockecblockMarshalBinary)
	prometheus.MustRegister(entryCreditBlockecblockBuildHeader)
	prometheus.MustRegister(entryCreditBlockecblockUnmarshalECBlock)
	prometheus.MustRegister(entryCreditBlockecblockUnmarshalBinaryData)
	prometheus.MustRegister(entryCreditBlockecblockUnmarshalBinary)
	prometheus.MustRegister(entryCreditBlockecblockmarshalHeaderBinary)
	prometheus.MustRegister(entryCreditBlockecblockunmarshalBodyBinaryData)
	prometheus.MustRegister(entryCreditBlockecblockunmarshalBodyBinary)
	prometheus.MustRegister(entryCreditBlockecblockunmarshalHeaderBinary)
	prometheus.MustRegister(entryCreditBlockecblockJSONByte)
	prometheus.MustRegister(entryCreditBlockecblockJSONString)
	prometheus.MustRegister(entryCreditBlockecblockJSONBuffer)
	prometheus.MustRegister(entryCreditBlockecblockNewECBlock)
	prometheus.MustRegister(entryCreditBlockecblockNextECBlock)
	prometheus.MustRegister(entryCreditBlockecblockCheckBlockPairIntegrity)
	// ecblockbody variables)
	prometheus.MustRegister(entryCreditBlockecblockBodyGetEntries)
	prometheus.MustRegister(entryCreditBlockecblockBodyAddEntry)
	prometheus.MustRegister(entryCreditBlockecblockBodySetEntries)
	prometheus.MustRegister(entryCreditBlockecblockBodyJSONByte)
	prometheus.MustRegister(entryCreditBlockecblockBodyJSONString)
	prometheus.MustRegister(entryCreditBlockecblockBodyJSONBuffer)
	prometheus.MustRegister(entryCreditBlockecblockBodyString)
	prometheus.MustRegister(entryCreditBlockecblockBodyNewECBlockBody)
	// ecblockHead variables)
	prometheus.MustRegister(entryCreditBlockecblockHeadString)
	prometheus.MustRegister(entryCreditBlockecblockHeadSetBodySize)
	prometheus.MustRegister(entryCreditBlockecblockHeadGetBodySize)
	prometheus.MustRegister(entryCreditBlockecblockHeadSetObjectCount)
	prometheus.MustRegister(entryCreditBlockecblockHeadGetObjectCount)
	prometheus.MustRegister(entryCreditBlockecblockHeadSetHeaderExpansionArea)
	prometheus.MustRegister(entryCreditBlockecblockHeadGetHeaderExpansionArea)
	prometheus.MustRegister(entryCreditBlockecblockHeadSetBodyHash)
	prometheus.MustRegister(entryCreditBlockecblockHeadGetBodyHash)
	prometheus.MustRegister(entryCreditBlockecblockHeadGetECChainID)
	prometheus.MustRegister(entryCreditBlockecblockHeadSetPrevHeaderHash)
	prometheus.MustRegister(entryCreditBlockecblockHeadGetPrevHeaderHash)
	prometheus.MustRegister(entryCreditBlockecblockHeadSetPrevFullHash)
	prometheus.MustRegister(entryCreditBlockecblockHeadGetPrevFullHash)
	prometheus.MustRegister(entryCreditBlockecblockHeadGetDBHeight)
	prometheus.MustRegister(entryCreditBlockecblockHeadSetDBHeight)
	prometheus.MustRegister(entryCreditBlockecblockHeadNewECBlockHeader)
	prometheus.MustRegister(entryCreditBlockecblockHeadJSONByte)
	prometheus.MustRegister(entryCreditBlockecblockHeadJSONString)
	prometheus.MustRegister(entryCreditBlockecblockHeadJSONBuffer)
	prometheus.MustRegister(entryCreditBlockecblockHeadMarshalBinary)
	prometheus.MustRegister(entryCreditBlockecblockHeadUnmarshalBinaryData)
	prometheus.MustRegister(entryCreditBlockecblockHeadUnmarshalBinary)
	prometheus.MustRegister(entryCreditBlockecblockHeadMarshalJSON)
	// increaseBalance variables)
	prometheus.MustRegister(entryCreditBlockincreaseBalanceString)
	prometheus.MustRegister(entryCreditBlockincreaseBalanceNewIncreaseBalance)
	prometheus.MustRegister(entryCreditBlockincreaseBalanceHash)
	prometheus.MustRegister(entryCreditBlockincreaseBalanceGetHash)
	prometheus.MustRegister(entryCreditBlockincreaseBalanceECID)
	prometheus.MustRegister(entryCreditBlockincreaseBalanceMarshalBinary)
	prometheus.MustRegister(entryCreditBlockincreaseBalanceUnmarshalBinaryData)
	prometheus.MustRegister(entryCreditBlockincreaseBalanceUnmarshalBinary)
	prometheus.MustRegister(entryCreditBlockincreaseBalanceJSONByte)
	prometheus.MustRegister(entryCreditBlockincreaseBalanceJSONString)
	prometheus.MustRegister(entryCreditBlockincreaseBalanceJSONBuffer)
	prometheus.MustRegister(entryCreditBlockincreaseBalanceGetTimestamp)
	// minuteNumber variables)
	prometheus.MustRegister(entryCreditBlockminuteNumberHash)
	prometheus.MustRegister(entryCreditBlockminuteNumberGetHash)
	prometheus.MustRegister(entryCreditBlockminuteNumberGetSigHash)
	prometheus.MustRegister(entryCreditBlockminuteNumberGetEntryHash)
	prometheus.MustRegister(entryCreditBlockminuteNumberIsInterpretable)
	prometheus.MustRegister(entryCreditBlockminuteNumberInterpret)
	prometheus.MustRegister(entryCreditBlockminuteNumberNewMinuteNumber)
	prometheus.MustRegister(entryCreditBlockminuteNumberECID)
	prometheus.MustRegister(entryCreditBlockminuteNumberMarshalBinary)
	prometheus.MustRegister(entryCreditBlockminuteNumberUnmarshalBinaryData)
	prometheus.MustRegister(entryCreditBlockminuteNumberUnmarshalBinary)
	prometheus.MustRegister(entryCreditBlockminuteNumberJSONByte)
	prometheus.MustRegister(entryCreditBlockminuteNumberJSONString)
	prometheus.MustRegister(entryCreditBlockminuteNumberJSONBuffer)
	prometheus.MustRegister(entryCreditBlockminuteNumberString)
	//serverindexnumber.go variables)
	prometheus.MustRegister(entryCreditBlockServerIndexNumberString)
	prometheus.MustRegister(entryCreditBlockServerIndexNumberHash)
	prometheus.MustRegister(entryCreditBlockServerIndexNumberGetHash)
	prometheus.MustRegister(entryCreditBlockServerIndexNumberInterpret)
	prometheus.MustRegister(entryCreditBlockServerIndexNumberNewServerIndexNumber)
	prometheus.MustRegister(entryCreditBlockServerIndexNumberNewServerIndexNumber2)
	prometheus.MustRegister(entryCreditBlockServerIndexNumberECID)
	prometheus.MustRegister(entryCreditBlockServerIndexNumberMarshalBinary)
	prometheus.MustRegister(entryCreditBlockServerIndexNumberUnmarshalBinaryData)
	prometheus.MustRegister(entryCreditBlockServerIndexNumberUnmarshalBinary)
	prometheus.MustRegister(entryCreditBlockServerIndexNumberJSONByte)
	prometheus.MustRegister(entryCreditBlockServerIndexNumberJSONString)
	prometheus.MustRegister(entryCreditBlockServerIndexNumberJSONBuffer)

}
