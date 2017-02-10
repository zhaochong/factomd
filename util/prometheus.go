package util

import "github.com/prometheus/client_golang/prometheus"

var (
	factomdutilFactomdConfigString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_util_FactomdConfig_String_Summary",
		Help: "Summary of calls to  factomd_util_FactomdConfig_String",
	})
	factomdutilConfigFilename = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_util_ConfigFilename_Summary",
		Help: "Summary of calls to  factomd_util_ConfigFilename",
	})
	factomdutilGetConfigFilename = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_util_GetConfigFilename_Summary",
		Help: "Summary of calls to  factomd_util_GetConfigFilename",
	})
	factomdutilGetChangeAcksHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_util_GetChangeAcksHeight_Summary",
		Help: "Summary of calls to  factomd_util_GetChangeAcksHeight",
	})
	factomdutilReadConfig = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_util_ReadConfig_Summary",
		Help: "Summary of calls to  factomd_util_ReadConfig",
	})
	factomdutilGetHomeDir = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_util_GetHomeDir_Summary",
		Help: "Summary of calls to  factomd_util_GetHomeDir",
	})
	factomdutilTrace = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_util_Trace_Summary",
		Help: "Summary of calls to  factomd_util_Trace",
	})
	factomdutilEntryCost = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_util_EntryCost_Summary",
		Help: "Summary of calls to  factomd_util_EntryCost",
	})
	factomdutilIsInPendingEntryList = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_util_IsInPendingEntryList_Summary",
		Help: "Summary of calls to  factomd_util_IsInPendingEntryList",
	})
	factomdutilByDBlockIDAccendingLen = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_util_ByDBlockIDAccending_Len_Summary",
		Help: "Summary of calls to  factomd_util_ByDBlockIDAccending_Len",
	})
	factomdutilByDBlockIDAccendingLess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_util_ByDBlockIDAccending_Less_Summary",
		Help: "Summary of calls to  factomd_util_ByDBlockIDAccending_Less",
	})
	factomdutilByDBlockIDAccendingSwap = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_util_ByDBlockIDAccending_Swap_Summary",
		Help: "Summary of calls to  factomd_util_ByDBlockIDAccending_Swap",
	})
	factomdutilByECBlockIDAccendingLen = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_util_ByECBlockIDAccending_Len_Summary",
		Help: "Summary of calls to  factomd_util_ByECBlockIDAccending_Len",
	})
	factomdutilByECBlockIDAccendingLess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_util_ByECBlockIDAccending_Less_Summary",
		Help: "Summary of calls to  factomd_util_ByECBlockIDAccending_Less",
	})
	factomdutilByECBlockIDAccendingSwap = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_util_ByECBlockIDAccending_Swap_Summary",
		Help: "Summary of calls to  factomd_util_ByECBlockIDAccending_Swap",
	})
	factomdutilByABlockIDAccendingLen = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_util_ByABlockIDAccending_Len_Summary",
		Help: "Summary of calls to  factomd_util_ByABlockIDAccending_Len",
	})
	factomdutilByABlockIDAccendingLess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_util_ByABlockIDAccending_Less_Summary",
		Help: "Summary of calls to  factomd_util_ByABlockIDAccending_Less",
	})
	factomdutilByABlockIDAccendingSwap = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_util_ByABlockIDAccending_Swap_Summary",
		Help: "Summary of calls to  factomd_util_ByABlockIDAccending_Swap",
	})
	factomdutilByFBlockIDAccendingLen = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_util_ByFBlockIDAccending_Len_Summary",
		Help: "Summary of calls to  factomd_util_ByFBlockIDAccending_Len",
	})
	factomdutilByFBlockIDAccendingLess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_util_ByFBlockIDAccending_Less_Summary",
		Help: "Summary of calls to  factomd_util_ByFBlockIDAccending_Less",
	})
	factomdutilByFBlockIDAccendingSwap = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_util_ByFBlockIDAccending_Swap_Summary",
		Help: "Summary of calls to  factomd_util_ByFBlockIDAccending_Swap",
	})
	factomdutilByEBlockIDAccendingLen = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_util_ByEBlockIDAccending_Len_Summary",
		Help: "Summary of calls to  factomd_util_ByEBlockIDAccending_Len",
	})
	factomdutilByEBlockIDAccendingLess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_util_ByEBlockIDAccending_Less_Summary",
		Help: "Summary of calls to  factomd_util_ByEBlockIDAccending_Less",
	})
	factomdutilByEBlockIDAccendingSwap = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_util_ByEBlockIDAccending_Swap_Summary",
		Help: "Summary of calls to  factomd_util_ByEBlockIDAccending_Swap",
	})
	factomdutilByByteArrayLen = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_util_ByByteArray_Len_Summary",
		Help: "Summary of calls to  factomd_util_ByByteArray_Len",
	})
	factomdutilByByteArrayLess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_util_ByByteArray_Less_Summary",
		Help: "Summary of calls to  factomd_util_ByByteArray_Less",
	})
	factomdutilByByteArraySwap = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_util_ByByteArray_Swap_Summary",
		Help: "Summary of calls to  factomd_util_ByByteArray_Swap",
	})
	factomdutilByDirBlockInfoIDAccendingLen = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_util_ByDirBlockInfoIDAccending_Len_Summary",
		Help: "Summary of calls to  factomd_util_ByDirBlockInfoIDAccending_Len",
	})
	factomdutilByDirBlockInfoIDAccendingLess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_util_ByDirBlockInfoIDAccending_Less_Summary",
		Help: "Summary of calls to  factomd_util_ByDirBlockInfoIDAccending_Less",
	})
	factomdutilByDirBlockInfoIDAccendingSwap = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_util_ByDirBlockInfoIDAccending_Swap_Summary",
		Help: "Summary of calls to  factomd_util_ByDirBlockInfoIDAccending_Swap",
	})
	factomdutilByDirBlockInfoTimestampLen = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_util_ByDirBlockInfoTimestamp_Len_Summary",
		Help: "Summary of calls to  factomd_util_ByDirBlockInfoTimestamp_Len",
	})
	factomdutilByDirBlockInfoTimestampLess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_util_ByDirBlockInfoTimestamp_Less_Summary",
		Help: "Summary of calls to  factomd_util_ByDirBlockInfoTimestamp_Less",
	})
	factomdutilByDirBlockInfoTimestampSwap = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_util_ByDirBlockInfoTimestamp_Swap_Summary",
		Help: "Summary of calls to  factomd_util_ByDirBlockInfoTimestamp_Swap",
	})
)

func InitPrometheus() {
	prometheus.MustRegister(factomdutilFactomdConfigString)
	prometheus.MustRegister(factomdutilConfigFilename)
	prometheus.MustRegister(factomdutilGetConfigFilename)
	prometheus.MustRegister(factomdutilGetChangeAcksHeight)
	prometheus.MustRegister(factomdutilReadConfig)
	prometheus.MustRegister(factomdutilGetHomeDir)
	prometheus.MustRegister(factomdutilTrace)
	prometheus.MustRegister(factomdutilEntryCost)
	prometheus.MustRegister(factomdutilIsInPendingEntryList)
	prometheus.MustRegister(factomdutilByDBlockIDAccendingLen)
	prometheus.MustRegister(factomdutilByDBlockIDAccendingLess)
	prometheus.MustRegister(factomdutilByDBlockIDAccendingSwap)
	prometheus.MustRegister(factomdutilByECBlockIDAccendingLen)
	prometheus.MustRegister(factomdutilByECBlockIDAccendingLess)
	prometheus.MustRegister(factomdutilByECBlockIDAccendingSwap)
	prometheus.MustRegister(factomdutilByABlockIDAccendingLen)
	prometheus.MustRegister(factomdutilByABlockIDAccendingLess)
	prometheus.MustRegister(factomdutilByABlockIDAccendingSwap)
	prometheus.MustRegister(factomdutilByFBlockIDAccendingLen)
	prometheus.MustRegister(factomdutilByFBlockIDAccendingLess)
	prometheus.MustRegister(factomdutilByFBlockIDAccendingSwap)
	prometheus.MustRegister(factomdutilByEBlockIDAccendingLen)
	prometheus.MustRegister(factomdutilByEBlockIDAccendingLess)
	prometheus.MustRegister(factomdutilByEBlockIDAccendingSwap)
	prometheus.MustRegister(factomdutilByByteArrayLen)
	prometheus.MustRegister(factomdutilByByteArrayLess)
	prometheus.MustRegister(factomdutilByByteArraySwap)
	prometheus.MustRegister(factomdutilByDirBlockInfoIDAccendingLen)
	prometheus.MustRegister(factomdutilByDirBlockInfoIDAccendingLess)
	prometheus.MustRegister(factomdutilByDirBlockInfoIDAccendingSwap)
	prometheus.MustRegister(factomdutilByDirBlockInfoTimestampLen)
	prometheus.MustRegister(factomdutilByDirBlockInfoTimestampLess)
	prometheus.MustRegister(factomdutilByDirBlockInfoTimestampSwap)
}
