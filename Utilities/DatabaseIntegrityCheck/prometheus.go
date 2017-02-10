package main

import "github.com/prometheus/client_golang/prometheus"

var (
	factomdmainmain = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_main_main_Summary",
		Help: "Summary of calls to  factomd_main_main",
	})
	factomdmainCheckDatabase = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_main_CheckDatabase_Summary",
		Help: "Summary of calls to  factomd_main_CheckDatabase",
	})
	factomdmainFetchBlockSet = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_main_FetchBlockSet_Summary",
		Help: "Summary of calls to  factomd_main_FetchBlockSet",
	})
)

func InitPrometheus() {
	prometheus.MustRegister(factomdmainmain)
	prometheus.MustRegister(factomdmainCheckDatabase)
	prometheus.MustRegister(factomdmainFetchBlockSet)
}
