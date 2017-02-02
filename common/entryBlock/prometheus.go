// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package entryBlock

import "github.com/prometheus/client_golang/prometheus"

var (
	//eblock.go variables
	entryBlockErrors = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "factomd_entryBlock_error_responses",
		Help: "Number of error responses from entryBlock objects",
	})
	entryBlockGetEntryHashes = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBloc_GetEntryHashes_Summary",
		Help: "Summary of entryBloc.GetEntryHashes call",
	})
	entryBlockGetEntrySigHashes = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBloc_GetEntrySigHashes_Summary",
		Help: "Summary of entryBloc.GetEntrySigHashes call",
	})
	entryBlockNew = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBloc_New_Summary",
		Help: "Summary of entryBloc.New call",
	})
	entryBlockGetWelds = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBloc_GetWelds_Summary",
		Help: "Summary of entryBloc.GetWelds call",
	})
	entryBlockGetWeldHashes = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBloc_GetWeldHashes_Summary",
		Help: "Summary of entryBloc.GetWeldHashes call",
	})
	entryBlockGetDatabaseHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBloc_GetDatabaseHeight_Summary",
		Help: "Summary of entryBloc.GetDatabaseHeight call",
	})
	entryBlockGetChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBloc_GetChainID_Summary",
		Help: "Summary of entryBloc.GetChainID call",
	})
	entryBlockGetHashOfChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBloc_GetHashOfChainID_Summary",
		Help: "Summary of entryBloc.GetChainID call",
	})
	entryBlockGetHashOfChainIDHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBloc_GetHashOfChainIDHash_Summary",
		Help: "Summary of entryBloc.GetChainIDHash call",
	})
	entryBlockDatabasePrimaryIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBloc_DatabasePrimaryIndex_Summary",
		Help: "Summary of entryBloc.DatabasePrimaryIndex call",
	})
	entryBlockDatabaseSecondaryIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBloc_DatabaseSecondaryIndex_Summary",
		Help: "Summary of entryBloc.DatabaseSecondaryIndex call",
	})
	entryBlockGetHeader = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBloc_GetHeader_Summary",
		Help: "Summary of entryBloc.GetHeader  call",
	})
	entryBlockGetBody = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBloc_GetBody_Summary",
		Help: "Summary of entryBloc.GetBody  call",
	})
	entryBlockAddEBEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBloc_AddEBEntry_Summary",
		Help: "Summary of entryBloc.AddEBEntry  call",
	})
	entryBlockAddEndOfMinuteMarker = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBloc_AddEndOfMinuteMarker_Summary",
		Help: "Summary of entryBloc.AddEndOfMinuteMarker call",
	})
	entryBlockBuildHeader = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBloc_BuildHeader_Summary",
		Help: "Summary of entryBloc.BuildHeader call",
	})
	entryBlockGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBloc_GetHash_Summary",
		Help: "Summary of entryBloc.GetHash call",
	})
	entryBlockHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBloc_Hash_Summary",
		Help: "Summary of entryBloc.Hash call",
	})
	entryBlockHeaderHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBloc_HeaderHash_Summary",
		Help: "Summary of entryBloc.HeaderHash call",
	})
	entryBlockBodyKeyMR = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBloc_BodyKeyMR_Summary",
		Help: "Summary of entryBloc.BodyKeyMR call",
	})
	entryBlockKeyMR = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBloc_KeyMR_Summary",
		Help: "Summary of entryBloc.KeyMR call",
	})
	entryBlockMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBloc_MarshalBinary_Summary",
		Help: "Summary of entryBloc.MarshalBinary call",
	})
	entryBlockUnmarshalEBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBloc_UnmarshalEBlock_Summary",
		Help: "Summary of entryBloc.UnmarshalEBlock call",
	})
	entryBlockUnmarshalEBlockData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBloc_UnmarshalEBlockData_Summary",
		Help: "Summary of entryBloc.UnmarshalEBlockData call",
	})
	entryBlockUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBloc_UnmarshalBinaryData_Summary",
		Help: "Summary of entryBloc.UnmarshalBinaryData call",
	})
	entryBlockUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBloc_UnmarshalBinary_Summary",
		Help: "Summary of entryBloc.UnmarshalBinary call",
	})
	entryBlockmarshalBodyBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBloc_marshalBodyBinary_Summary",
		Help: "Summary of entryBloc.marshalBodyBinary call",
	})
	entryBlockunmarshalBodyBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBloc_unmarshalBodyBinaryData_Summary",
		Help: "Summary of entryBloc.marshalBodyBinary call",
	})
	entryBlockunmarshalBodyBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBloc_unmarshalBodyBinary_Summary",
		Help: "Summary of entryBloc.unmarshalBodyBinary call",
	})
	entryBlockJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBloc_JSONByte_Summary",
		Help: "Summary of entryBloc.JSONByte call",
	})
	entryBlockJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBloc_JSONString_Summary",
		Help: "Summary of entryBloc.JSONString call",
	})
	entryBlockJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBloc_JSONBuffer_Summary",
		Help: "Summary of entryBloc.JSONBuffer call",
	})
	entryBlockString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBloc_String_Summary",
		Help: "Summary of entryBloc.String call",
	})
	entryBlockNewEBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBloc_NewEBlock_Summary",
		Help: "Summary of entryBloc.NewEBlock call",
	})
	//eblockBody.go variables
	eBlockBodyNewEBlockBody = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBloc_NewEBlockBody_Summary",
		Help: "Summary of entryBloc.NewEBlockBody call",
	})
	eBlockBodyMR = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBloc_MR_Summary",
		Help: "Summary of entryBloc.MR call",
	})
	eBlockBodyJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBloc_JSONByte_Summary",
		Help: "Summary of entryBloc.JSONByte call",
	})
	eBlockBodyJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBloc_JSONString_Summary",
		Help: "Summary of entryBloc.JSONString call",
	})
	eBlockBodyJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBloc_JSONBuffer_Summary",
		Help: "Summary of entryBloc.JSONBuffer call",
	})
	eBlockBodyString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBloc_String_Summary",
		Help: "Summary of entryBloc.String call",
	})
	eBlockBodyGetEBEntries = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBloc_GetEBEntries_Summary",
		Help: "Summary of entryBloc.GetEBEntries call",
	})
	//eblockHeader variables
	eBlockHeaderNewEBlockHeader = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlockhead_NewEBlockHeader_Summary",
		Help: "Summary of entryBlockhead.NewEBlockHeader call",
	})
	eBlockHeaderJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlockhead_JSONByte_Summary",
		Help: "Summary of entryBlockhead.JSONByte call",
	})
	eBlockHeaderJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlockhead_JSONString_Summary",
		Help: "Summary of entryBlockhead.JSONString call",
	})
	eBlockHeaderJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlockhead_JSONBuffer_Summary",
		Help: "Summary of entryBlockhead.JSONBuffer call",
	})
	eBlockHeaderString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlockhead_String_Summary",
		Help: "Summary of entryBlockhead.String call",
	})
	eBlockHeaderGetChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlockhead_GetChainID_Summary",
		Help: "Summary of entryBlockhead.GetChainID call",
	})
	eBlockHeaderGetBodyMR = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlockhead_GetBodyMR_Summary",
		Help: "Summary of entryBlockhead.GetBodyMR call",
	})
	eBlockHeaderSetBodyMR = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlockhead_SetBodyMR_Summary",
		Help: "Summary of entryBlockhead.SetBodyMR call",
	})
	eBlockHeaderSetPrevKeyMR = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlockhead_SetPrevKeyMR_Summary",
		Help: "Summary of entryBlockhead.SetPrevKeyMR call",
	})
	eBlockHeaderGetPrevFullHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlockhead_GetPrevFullHash_Summary",
		Help: "Summary of entryBlockhead.GetPrevFullHash call",
	})
	eBlockHeaderSetPrevFullHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlockhead_SetPrevFullHash_Summary",
		Help: "Summary of entryBlockhead.SetPrevFullHash call",
	})
	eBlockHeaderGetEBSequence = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlockhead_GetEBSequence_Summary",
		Help: "Summary of entryBlockhead.GetEBSequence call",
	})
	eBlockHeaderSetEBSequence = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlockhead_SetEBSequence_Summary",
		Help: "Summary of entryBlockhead.SetEBSequence call",
	})
	eBlockHeaderGetDBHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlockhead_GetDBHeight_Summary",
		Help: "Summary of entryBlockhead.GetDBHeight call",
	})
	eBlockHeaderSetDBHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlockhead_SetDBHeight_Summary",
		Help: "Summary of entryBlockhead.SetDBHeight call",
	})
	eBlockHeaderGetEntryCount = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlockhead_GetEntryCount_Summary",
		Help: "Summary of entryBlockhead.GetEntryCount call",
	})
	eBlockHeaderSetEntryCount = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlockhead_SetEntryCount_Summary",
		Help: "Summary of entryBlockhead.SetEntryCount call",
	})
	eBlockHeaderMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlockhead_MarshalBinary_Summary",
		Help: "Summary of entryBlockhead.MarshalBinary call",
	})
	eBlockHeaderUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlockhead_UnmarshalBinaryData_Summary",
		Help: "Summary of entryBlockhead.UnmarshalBinaryData call",
	})
	eBlockHeaderUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entryBlockhead_UnmarshalBinary_Summary",
		Help: "Summary of entryBlockhead.UnmarshalBinary call",
	})
	//entry.go variables
	entryKSize = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entry_KSize_Summary",
		Help: "Summary of entry.KSize call",
	})
	entryNew = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entry_New_Summary",
		Help: "Summary of entry.New call",
	})
	entryGetDatabaseHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entry_GetDatabaseHeight_Summary",
		Help: "Summary of entry.GetDatabaseHeight call",
	})
	entryGetWeld = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entry_GetWeld_Summary",
		Help: "Summary of entry.GetWeld call",
	})
	entryGetWeldHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entry_GetWeldHash_Summary",
		Help: "Summary of entry.GetWeldHash call",
	})
	entryGetChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entry_GetChainID_Summary",
		Help: "Summary of entry.GetChainID call",
	})
	entryDatabasePrimaryIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entry_DatabasePrimaryIndex_Summary",
		Help: "Summary of entry.DatabasePrimaryIndex call",
	})
	entryDatabaseSecondaryIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entry_DatabaseSecondaryIndex_Summary",
		Help: "Summary of entry.DatabaseSecondaryIndex call",
	})
	entryNewChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entry_NewChainID_Summary",
		Help: "Summary of entry.NewChainID call",
	})
	entryGetContent = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entry_GetContent_Summary",
		Help: "Summary of entry.GetContent call",
	})
	entryGetChainIDHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entry_GetChainIDHash_Summary",
		Help: "Summary of entry.GetChainIDHash call",
	})
	entryExternalIDs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entry_ExternalIDs_Summary",
		Help: "Summary of entry.ExternalIDs call",
	})
	entryIsValid = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entry_IsValid_Summary",
		Help: "Summary of entry.IsValid call",
	})
	entryGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entry_GetHash_Summary",
		Help: "Summary of entry.GetHash call",
	})
	entryMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entry_MarshalBinary_Summary",
		Help: "Summary of entry.MarshalBinary call",
	})
	entryMarshalExtIDsBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entry_MarshalExtIDsBinary_Summary",
		Help: "Summary of entry.MarshalExtIDsBinary call",
	})
	entryUnmarshalEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entry_UnmarshalEntry_Summary",
		Help: "Summary of entry.UnmarshalEntry call",
	})
	entryUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entry_UnmarshalBinaryData_Summary",
		Help: "Summary of entry.UnmarshalBinaryData call",
	})
	entryUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entry_UnmarshalBinary_Summary",
		Help: "Summary of entry.UnmarshalBinary call",
	})
	entryJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entry_JSONByte_Summary",
		Help: "Summary of entry.JSONByte call",
	})
	entryJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entry_JSONString_Summary",
		Help: "Summary of entry.JSONString call",
	})
	entryJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entry_JSONBuffer_Summary",
		Help: "Summary of entry.JSONBuffer call",
	})
	entryString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entry_String_Summary",
		Help: "Summary of entry.String call",
	})
	entryNewEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_entry_NewEntry_Summary",
		Help: "Summary of entry.NewEntry call",
	})
)

func InitPrometheus() {
	//wsapi variables
	prometheus.MustRegister(entryBlockErrors)
	prometheus.MustRegister(entryBlockGetEntryHashes)
	prometheus.MustRegister(entryBlockGetEntrySigHashes)
	prometheus.MustRegister(entryBlockNew)
	prometheus.MustRegister(entryBlockGetWelds)
	prometheus.MustRegister(entryBlockGetWeldHashes)
	prometheus.MustRegister(entryBlockGetDatabaseHeight)
	prometheus.MustRegister(entryBlockGetChainID)
	prometheus.MustRegister(entryBlockGetHashOfChainID)
	prometheus.MustRegister(entryBlockGetHashOfChainIDHash)
	prometheus.MustRegister(entryBlockDatabasePrimaryIndex)
	prometheus.MustRegister(entryBlockDatabaseSecondaryIndex)
	prometheus.MustRegister(entryBlockGetHeader)
	prometheus.MustRegister(entryBlockGetBody)
	prometheus.MustRegister(entryBlockAddEBEntry)
	prometheus.MustRegister(entryBlockAddEndOfMinuteMarker)
	prometheus.MustRegister(entryBlockBuildHeader)
	prometheus.MustRegister(entryBlockGetHash)
	prometheus.MustRegister(entryBlockHash)
	prometheus.MustRegister(entryBlockHeaderHash)
	prometheus.MustRegister(entryBlockBodyKeyMR)
	prometheus.MustRegister(entryBlockKeyMR)
	prometheus.MustRegister(entryBlockMarshalBinary)
	prometheus.MustRegister(entryBlockUnmarshalEBlock)
	prometheus.MustRegister(entryBlockUnmarshalEBlockData)
	prometheus.MustRegister(entryBlockUnmarshalBinaryData)
	prometheus.MustRegister(entryBlockUnmarshalBinary)
	prometheus.MustRegister(entryBlockmarshalBodyBinary)
	prometheus.MustRegister(entryBlockunmarshalBodyBinaryData)
	prometheus.MustRegister(entryBlockunmarshalBodyBinary)
	prometheus.MustRegister(entryBlockJSONByte)
	prometheus.MustRegister(entryBlockJSONString)
	prometheus.MustRegister(entryBlockJSONBuffer)
	prometheus.MustRegister(entryBlockString)
	prometheus.MustRegister(entryBlockNewEBlock)
	prometheus.MustRegister(eBlockBodyNewEBlockBody)
	prometheus.MustRegister(eBlockBodyMR)
	prometheus.MustRegister(eBlockBodyJSONByte)
	prometheus.MustRegister(eBlockBodyJSONString)
	prometheus.MustRegister(eBlockBodyJSONBuffer)
	prometheus.MustRegister(eBlockBodyString)
	prometheus.MustRegister(eBlockBodyGetEBEntries)
	//eblockHeader variables)
	prometheus.MustRegister(eBlockHeaderNewEBlockHeader)
	prometheus.MustRegister(eBlockHeaderJSONByte)
	prometheus.MustRegister(eBlockHeaderJSONString)
	prometheus.MustRegister(eBlockHeaderJSONBuffer)
	prometheus.MustRegister(eBlockHeaderString)
	prometheus.MustRegister(eBlockHeaderGetChainID)
	prometheus.MustRegister(eBlockHeaderGetBodyMR)
	prometheus.MustRegister(eBlockHeaderSetBodyMR)
	prometheus.MustRegister(eBlockHeaderSetPrevKeyMR)
	prometheus.MustRegister(eBlockHeaderGetPrevFullHash)
	prometheus.MustRegister(eBlockHeaderSetPrevFullHash)
	prometheus.MustRegister(eBlockHeaderGetEBSequence)
	prometheus.MustRegister(eBlockHeaderSetEBSequence)
	prometheus.MustRegister(eBlockHeaderGetDBHeight)
	prometheus.MustRegister(eBlockHeaderSetDBHeight)
	prometheus.MustRegister(eBlockHeaderGetEntryCount)
	prometheus.MustRegister(eBlockHeaderSetEntryCount)
	prometheus.MustRegister(eBlockHeaderMarshalBinary)
	prometheus.MustRegister(eBlockHeaderUnmarshalBinaryData)
	prometheus.MustRegister(eBlockHeaderUnmarshalBinary)
	//entry.go variables)
	prometheus.MustRegister(entryKSize)
	prometheus.MustRegister(entryNew)
	prometheus.MustRegister(entryGetDatabaseHeight)
	prometheus.MustRegister(entryGetWeld)
	prometheus.MustRegister(entryGetWeldHash)
	prometheus.MustRegister(entryGetChainID)
	prometheus.MustRegister(entryDatabasePrimaryIndex)
	prometheus.MustRegister(entryDatabaseSecondaryIndex)
	prometheus.MustRegister(entryNewChainID)
	prometheus.MustRegister(entryGetContent)
	prometheus.MustRegister(entryGetChainIDHash)
	prometheus.MustRegister(entryExternalIDs)
	prometheus.MustRegister(entryIsValid)
	prometheus.MustRegister(entryGetHash)
	prometheus.MustRegister(entryMarshalBinary)
	prometheus.MustRegister(entryMarshalExtIDsBinary)
	prometheus.MustRegister(entryUnmarshalEntry)
	prometheus.MustRegister(entryUnmarshalBinaryData)
	prometheus.MustRegister(entryUnmarshalBinary)
	prometheus.MustRegister(entryJSONByte)
	prometheus.MustRegister(entryJSONString)
	prometheus.MustRegister(entryJSONBuffer)
	prometheus.MustRegister(entryString)
	prometheus.MustRegister(entryNewEntry)
}
