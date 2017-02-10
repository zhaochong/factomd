package dbInfo

import "github.com/prometheus/client_golang/prometheus"

var (
	factomddbInfoDirBlockInfoInit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_DirBlockInfo_Init_Summary",
		Help: "Summary of calls to  factomd_dbInfo_DirBlockInfo_Init",
	})
	factomddbInfoNewDirBlockInfo = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_NewDirBlockInfo_Summary",
		Help: "Summary of calls to  factomd_dbInfo_NewDirBlockInfo",
	})
	factomddbInfoDirBlockInfoString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_DirBlockInfo_String_Summary",
		Help: "Summary of calls to  factomd_dbInfo_DirBlockInfo_String",
	})
	factomddbInfoDirBlockInfoJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_DirBlockInfo_JSONByte_Summary",
		Help: "Summary of calls to  factomd_dbInfo_DirBlockInfo_JSONByte",
	})
	factomddbInfoDirBlockInfoJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_DirBlockInfo_JSONString_Summary",
		Help: "Summary of calls to  factomd_dbInfo_DirBlockInfo_JSONString",
	})
	factomddbInfoDirBlockInfoNew = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_DirBlockInfo_New_Summary",
		Help: "Summary of calls to  factomd_dbInfo_DirBlockInfo_New",
	})
	factomddbInfoDirBlockInfoGetDatabaseHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_DirBlockInfo_GetDatabaseHeight_Summary",
		Help: "Summary of calls to  factomd_dbInfo_DirBlockInfo_GetDatabaseHeight",
	})
	factomddbInfoDirBlockInfoGetDBHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_DirBlockInfo_GetDBHeight_Summary",
		Help: "Summary of calls to  factomd_dbInfo_DirBlockInfo_GetDBHeight",
	})
	factomddbInfoDirBlockInfoGetBTCConfirmed = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_DirBlockInfo_GetBTCConfirmed_Summary",
		Help: "Summary of calls to  factomd_dbInfo_DirBlockInfo_GetBTCConfirmed",
	})
	factomddbInfoDirBlockInfoGetChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_DirBlockInfo_GetChainID_Summary",
		Help: "Summary of calls to  factomd_dbInfo_DirBlockInfo_GetChainID",
	})
	factomddbInfoDirBlockInfoDatabasePrimaryIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_DirBlockInfo_DatabasePrimaryIndex_Summary",
		Help: "Summary of calls to  factomd_dbInfo_DirBlockInfo_DatabasePrimaryIndex",
	})
	factomddbInfoDirBlockInfoDatabaseSecondaryIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_DirBlockInfo_DatabaseSecondaryIndex_Summary",
		Help: "Summary of calls to  factomd_dbInfo_DirBlockInfo_DatabaseSecondaryIndex",
	})
	factomddbInfoDirBlockInfoGetDBMerkleRoot = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_DirBlockInfo_GetDBMerkleRoot_Summary",
		Help: "Summary of calls to  factomd_dbInfo_DirBlockInfo_GetDBMerkleRoot",
	})
	factomddbInfoDirBlockInfoGetBTCTxHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_DirBlockInfo_GetBTCTxHash_Summary",
		Help: "Summary of calls to  factomd_dbInfo_DirBlockInfo_GetBTCTxHash",
	})
	factomddbInfoDirBlockInfoGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_DirBlockInfo_GetTimestamp_Summary",
		Help: "Summary of calls to  factomd_dbInfo_DirBlockInfo_GetTimestamp",
	})
	factomddbInfoDirBlockInfoGetBTCBlockHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_DirBlockInfo_GetBTCBlockHeight_Summary",
		Help: "Summary of calls to  factomd_dbInfo_DirBlockInfo_GetBTCBlockHeight",
	})
	factomddbInfoDirBlockInfoMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_DirBlockInfo_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_dbInfo_DirBlockInfo_MarshalBinary",
	})
	factomddbInfoDirBlockInfoUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_DirBlockInfo_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_dbInfo_DirBlockInfo_UnmarshalBinaryData",
	})
	factomddbInfoDirBlockInfoUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_DirBlockInfo_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_dbInfo_DirBlockInfo_UnmarshalBinary",
	})
	factomddbInfoDirBlockInfoSetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_DirBlockInfo_SetTimestamp_Summary",
		Help: "Summary of calls to  factomd_dbInfo_DirBlockInfo_SetTimestamp",
	})
	factomddbInfonewDirBlockInfoCopyFromDBI = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_newDirBlockInfoCopyFromDBI_Summary",
		Help: "Summary of calls to  factomd_dbInfo_newDirBlockInfoCopyFromDBI",
	})
	factomddbInfonewDirBlockInfoCopy = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_newDirBlockInfoCopy_Summary",
		Help: "Summary of calls to  factomd_dbInfo_newDirBlockInfoCopy",
	})
	factomddbInfoDirBlockInfoparseDirBlockInfoCopy = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_DirBlockInfo_parseDirBlockInfoCopy_Summary",
		Help: "Summary of calls to  factomd_dbInfo_DirBlockInfo_parseDirBlockInfoCopy",
	})
	factomddbInfoNewDirBlockInfoFromDirBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dbInfo_NewDirBlockInfoFromDirBlock_Summary",
		Help: "Summary of calls to  factomd_dbInfo_NewDirBlockInfoFromDirBlock",
	})
)

func InitPrometheus() {
	prometheus.MustRegister(factomddbInfoDirBlockInfoInit)
	prometheus.MustRegister(factomddbInfoNewDirBlockInfo)
	prometheus.MustRegister(factomddbInfoDirBlockInfoString)
	prometheus.MustRegister(factomddbInfoDirBlockInfoJSONByte)
	prometheus.MustRegister(factomddbInfoDirBlockInfoJSONString)
	prometheus.MustRegister(factomddbInfoDirBlockInfoNew)
	prometheus.MustRegister(factomddbInfoDirBlockInfoGetDatabaseHeight)
	prometheus.MustRegister(factomddbInfoDirBlockInfoGetDBHeight)
	prometheus.MustRegister(factomddbInfoDirBlockInfoGetBTCConfirmed)
	prometheus.MustRegister(factomddbInfoDirBlockInfoGetChainID)
	prometheus.MustRegister(factomddbInfoDirBlockInfoDatabasePrimaryIndex)
	prometheus.MustRegister(factomddbInfoDirBlockInfoDatabaseSecondaryIndex)
	prometheus.MustRegister(factomddbInfoDirBlockInfoGetDBMerkleRoot)
	prometheus.MustRegister(factomddbInfoDirBlockInfoGetBTCTxHash)
	prometheus.MustRegister(factomddbInfoDirBlockInfoGetTimestamp)
	prometheus.MustRegister(factomddbInfoDirBlockInfoGetBTCBlockHeight)
	prometheus.MustRegister(factomddbInfoDirBlockInfoMarshalBinary)
	prometheus.MustRegister(factomddbInfoDirBlockInfoUnmarshalBinaryData)
	prometheus.MustRegister(factomddbInfoDirBlockInfoUnmarshalBinary)
	prometheus.MustRegister(factomddbInfoDirBlockInfoSetTimestamp)
	prometheus.MustRegister(factomddbInfonewDirBlockInfoCopyFromDBI)
	prometheus.MustRegister(factomddbInfonewDirBlockInfoCopy)
	prometheus.MustRegister(factomddbInfoDirBlockInfoparseDirBlockInfoCopy)
	prometheus.MustRegister(factomddbInfoNewDirBlockInfoFromDirBlock)
}
