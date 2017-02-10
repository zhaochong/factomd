package main

import "github.com/prometheus/client_golang/prometheus"

var (
	factomdmainmain = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_main_main_Summary",
		Help: "Summary of calls to  factomd_main_main",
	})
)

func InitPrometheus() {
	prometheus.MustRegister(factomdmainmain)
}
