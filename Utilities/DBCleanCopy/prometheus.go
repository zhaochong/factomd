package main

import "github.com/prometheus/client_golang/prometheus"

var (
	factomdmainmain = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_main_main_Summary",
		Help: "Summary of calls to  factomd_main_main",
	})
	factomdmainCopyDB = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_main_CopyDB_Summary",
		Help: "Summary of calls to  factomd_main_CopyDB",
	})
)

func InitPrometheus() {
	prometheus.MustRegister(factomdmainmain)
	prometheus.MustRegister(factomdmainCopyDB)
}
