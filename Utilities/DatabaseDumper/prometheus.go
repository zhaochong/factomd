package main

import "github.com/prometheus/client_golang/prometheus"

var (
	factomdmainmain = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_main_main_Summary",
		Help: "Summary of calls to  factomd_main_main",
	})
	factomdmainExportDatabaseJSON = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_main_ExportDatabaseJSON_Summary",
		Help: "Summary of calls to  factomd_main_ExportDatabaseJSON",
	})
	factomdmainKeyToName = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_main_KeyToName_Summary",
		Help: "Summary of calls to  factomd_main_KeyToName",
	})
)

func InitPrometheus() {
	prometheus.MustRegister(factomdmainmain)
	prometheus.MustRegister(factomdmainExportDatabaseJSON)
	prometheus.MustRegister(factomdmainKeyToName)
}
