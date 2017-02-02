// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package dbInfo

import "github.com/prometheus/client_golang/prometheus"

var (
	//anchorRecords variables
	dbInfoErrors = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "factomd_anchor_error_responses",
		Help: "Number of error responses from dbInfo objects",
	})
	dbInfoNewDirBlockInfo = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_NewDirBlockInfo_Summary",
		Help: "Summary of dbInfo.NewDirBlockInfo call",
	})
	dbInfoString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_String_Summary",
		Help: "Summary of dbInfo.String call",
	})
	dbInfoJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_JSONByte_Summary",
		Help: "Summary of dbInfo.JSONByte call",
	})
	dbInfoJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_JSONString_Summary",
		Help: "Summary of dbInfo.JSONString call",
	})
	dbInfoJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_JSONBuffer_Summary",
		Help: "Summary of dbInfo.JSONBuffer call",
	})
	dbInfoNew = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_New_Summary",
		Help: "Summary of dbInfo.New call",
	})
	dbInfoGetDatabaseHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_GetDatabaseHeight_Summary",
		Help: "Summary of dbInfo.GetDatabaseHeight call",
	})
	dbInfoGetDBHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_GetDBHeight_Summary",
		Help: "Summary of dbInfo.GetDBHeight call",
	})
	dbInfoGetBTCConfirmed = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_GetBTCConfirmed_Summary",
		Help: "Summary of dbInfo.GetBTCConfirmed call",
	})
	dbInfoGetChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_GetChainID_Summary",
		Help: "Summary of dbInfo.GetChainID call",
	})
	dbInfoDatabasePrimaryIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_DatabasePrimaryIndex_Summary",
		Help: "Summary of dbInfo.DatabasePrimaryIndex call",
	})
	dbInfoDatabaseSecondaryIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_DatabaseSecondaryIndex_Summary",
		Help: "Summary of dbInfo.DatabaseSecondaryIndex call",
	})
	dbInfoGetDBMerkleRoot = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_GetDBMerkleRoot_Summary",
		Help: "Summary of dbInfo.GetDBMerkleRoot call",
	})
	dbInfoGetBTCTxHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_GetBTCTxHash_Summary",
		Help: "Summary of dbInfo.GetBTCTxHash call",
	})
	dbInfoGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_GetTimestamp_Summary",
		Help: "Summary of dbInfo.GetTimestamp call",
	})
	dbInfoGetBTCBlockHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_GetBTCBlockHeight_Summary",
		Help: "Summary of dbInfo.GetBTCBlockHeight call",
	})
	dbInfoMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_MarshalBinary_Summary",
		Help: "Summary of dbInfo.MarshalBinary call",
	})
	dbInfoUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_UnmarshalBinaryData_Summary",
		Help: "Summary of dbInfo.UnmarshalBinaryData call",
	})
	dbInfoUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_UnmarshalBinary_Summary",
		Help: "Summary of dbInfo.UnmarshalBinary call",
	})
	dbInfoSetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_SetTimestamp_Summary",
		Help: "Summary of dbInfo.SetTimestamp call",
	})
	dbInfonewDirBlockInfoCopyFromDBI = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_newDirBlockInfoCopyFromDBI_Summary",
		Help: "Summary of dbInfo.newDirBlockInfoCopyFromDBI call",
	})
	dbInfonewDirBlockInfoCopy = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_newDirBlockInfoCopy_Summary",
		Help: "Summary of dbInfo.newDirBlockInfoCopy call",
	})
	dbInfoparseDirBlockInfoCopy = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_parseDirBlockInfoCopy_Summary",
		Help: "Summary of dbInfo.parseDirBlockInfoCopy call",
	})
	dbInfoNewDirBlockInfoFromDirBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_NewDirBlockInfoFromDirBlock_Summary",
		Help: "Summary of dbInfo.NewDirBlockInfoFromDirBlock call",
	})
)

func InitPrometheus() {
	//wsapi variables
	prometheus.MustRegister(AnchorErrors)
	prometheus.MustRegister(dbInfoErrors)
	prometheus.MustRegister(dbInfoNewDirBlockInfo)
	prometheus.MustRegister(dbInfoString)
	prometheus.MustRegister(dbInfoJSONByte)
	prometheus.MustRegister(dbInfoJSONString)
	prometheus.MustRegister(dbInfoJSONBuffer)
	prometheus.MustRegister(dbInfoNew)
	prometheus.MustRegister(dbInfoGetDatabaseHeight)
	prometheus.MustRegister(dbInfoGetDBHeight)
	prometheus.MustRegister(dbInfoGetBTCConfirmed)
	prometheus.MustRegister(dbInfoGetChainID)
	prometheus.MustRegister(dbInfoDatabasePrimaryIndex)
	prometheus.MustRegister(dbInfoDatabaseSecondaryIndex)
	prometheus.MustRegister(dbInfoGetDBMerkleRoot)
	prometheus.MustRegister(dbInfoGetBTCTxHash)
	prometheus.MustRegister(dbInfoGetTimestamp)
	prometheus.MustRegister(dbInfoGetBTCBlockHeight)
	prometheus.MustRegister(dbInfoMarshalBinary)
	prometheus.MustRegister(dbInfoUnmarshalBinaryData)
	prometheus.MustRegister(dbInfoUnmarshalBinaryData)
	prometheus.MustRegister(dbInfoUnmarshalBinary)
	prometheus.MustRegister(dbInfoSetTimestamp)
	prometheus.MustRegister(dbInfonewDirBlockInfoCopyFromDBI)
	prometheus.MustRegister(dbInfonewDirBlockInfoCopy)
	prometheus.MustRegister(dbInfoparseDirBlockInfoCopy)
	prometheus.MustRegister(dbInfoNewDirBlockInfoFromDirBlock)
}
