package main

import "github.com/prometheus/client_golang/prometheus"

var (
	factomdmainInitNetwork = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_main_InitNetwork_Summary",
		Help: "Summary of calls to  factomd_main_InitNetwork",
	})
	factomdmainMsgIsNew = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_main_MsgIsNew_Summary",
		Help: "Summary of calls to  factomd_main_MsgIsNew",
	})
	factomdmainSetMsg = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_main_SetMsg_Summary",
		Help: "Summary of calls to  factomd_main_SetMsg",
	})
	factomdmainlisten = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_main_listen_Summary",
		Help: "Summary of calls to  factomd_main_listen",
	})
	factomdmainmain = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_main_main_Summary",
		Help: "Summary of calls to  factomd_main_main",
	})
)

func InitPrometheus() {
	prometheus.MustRegister(factomdmainInitNetwork)
	prometheus.MustRegister(factomdmainMsgIsNew)
	prometheus.MustRegister(factomdmainSetMsg)
	prometheus.MustRegister(factomdmainlisten)
	prometheus.MustRegister(factomdmainmain)
}
