package main

import "github.com/prometheus/client_golang/prometheus"

var (
	factomdmainGetDBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_main_GetDBlock_Summary",
		Help: "Summary of calls to  factomd_main_GetDBlock",
	})
	factomdmainGetABlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_main_GetABlock_Summary",
		Help: "Summary of calls to  factomd_main_GetABlock",
	})
	factomdmainGetECBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_main_GetECBlock_Summary",
		Help: "Summary of calls to  factomd_main_GetECBlock",
	})
	factomdmainGetFBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_main_GetFBlock_Summary",
		Help: "Summary of calls to  factomd_main_GetFBlock",
	})
	factomdmainGetEBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_main_GetEBlock_Summary",
		Help: "Summary of calls to  factomd_main_GetEBlock",
	})
	factomdmainGetEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_main_GetEntry_Summary",
		Help: "Summary of calls to  factomd_main_GetEntry",
	})
	factomdmainGetDBlockHead = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_main_GetDBlockHead_Summary",
		Help: "Summary of calls to  factomd_main_GetDBlockHead",
	})
	factomdmainGetRaw = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_main_GetRaw_Summary",
		Help: "Summary of calls to  factomd_main_GetRaw",
	})
	factomdmainInitBolt = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_main_InitBolt_Summary",
		Help: "Summary of calls to  factomd_main_InitBolt",
	})
	factomdmainInitLevelDB = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_main_InitLevelDB_Summary",
		Help: "Summary of calls to  factomd_main_InitLevelDB",
	})
	factomdmainInitMapDB = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_main_InitMapDB_Summary",
		Help: "Summary of calls to  factomd_main_InitMapDB",
	})
	factomdmainPrintoutLoop = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_main_PrintoutLoop_Summary",
		Help: "Summary of calls to  factomd_main_PrintoutLoop",
	})
	factomdmainmain = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_main_main_Summary",
		Help: "Summary of calls to  factomd_main_main",
	})
	factomdmainSaveBlocksLoop = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_main_SaveBlocksLoop_Summary",
		Help: "Summary of calls to  factomd_main_SaveBlocksLoop",
	})
	factomdmainSaveToDBLoop = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_main_SaveToDBLoop_Summary",
		Help: "Summary of calls to  factomd_main_SaveToDBLoop",
	})
	factomdmainCheckDatabaseForMissingEntries = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_main_CheckDatabaseForMissingEntries_Summary",
		Help: "Summary of calls to  factomd_main_CheckDatabaseForMissingEntries",
	})
	factomdmainCheckDBlockEntries = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_main_CheckDBlockEntries_Summary",
		Help: "Summary of calls to  factomd_main_CheckDBlockEntries",
	})
	factomdmainGetDBlockList = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_main_GetDBlockList_Summary",
		Help: "Summary of calls to  factomd_main_GetDBlockList",
	})
)

func InitPrometheus() {
	prometheus.MustRegister(factomdmainGetDBlock)
	prometheus.MustRegister(factomdmainGetABlock)
	prometheus.MustRegister(factomdmainGetECBlock)
	prometheus.MustRegister(factomdmainGetFBlock)
	prometheus.MustRegister(factomdmainGetEBlock)
	prometheus.MustRegister(factomdmainGetEntry)
	prometheus.MustRegister(factomdmainGetDBlockHead)
	prometheus.MustRegister(factomdmainGetRaw)
	prometheus.MustRegister(factomdmainInitBolt)
	prometheus.MustRegister(factomdmainInitLevelDB)
	prometheus.MustRegister(factomdmainInitMapDB)
	prometheus.MustRegister(factomdmainPrintoutLoop)
	prometheus.MustRegister(factomdmainmain)
	prometheus.MustRegister(factomdmainSaveBlocksLoop)
	prometheus.MustRegister(factomdmainSaveToDBLoop)
	prometheus.MustRegister(factomdmainCheckDatabaseForMissingEntries)
	prometheus.MustRegister(factomdmainCheckDBlockEntries)
	prometheus.MustRegister(factomdmainGetDBlockList)
}
