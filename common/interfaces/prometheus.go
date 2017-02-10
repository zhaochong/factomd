package interfaces

import "github.com/prometheus/client_golang/prometheus"

var (
	factomdinterfacesFctServerGetChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_interfaces_FctServer_GetChainID_Summary",
		Help: "Summary of calls to  factomd_interfaces_FctServer_GetChainID",
	})
	factomdinterfacesFctServerString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_interfaces_FctServer_String_Summary",
		Help: "Summary of calls to  factomd_interfaces_FctServer_String",
	})
	factomdinterfacesFctServerGetName = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_interfaces_FctServer_GetName_Summary",
		Help: "Summary of calls to  factomd_interfaces_FctServer_GetName",
	})
	factomdinterfacesFctServerIsOnline = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_interfaces_FctServer_IsOnline_Summary",
		Help: "Summary of calls to  factomd_interfaces_FctServer_IsOnline",
	})
	factomdinterfacesFctServerSetOnline = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_interfaces_FctServer_SetOnline_Summary",
		Help: "Summary of calls to  factomd_interfaces_FctServer_SetOnline",
	})
	factomdinterfacesFctServerLeaderToReplace = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_interfaces_FctServer_LeaderToReplace_Summary",
		Help: "Summary of calls to  factomd_interfaces_FctServer_LeaderToReplace",
	})
	factomdinterfacesFctServerSetReplace = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_interfaces_FctServer_SetReplace_Summary",
		Help: "Summary of calls to  factomd_interfaces_FctServer_SetReplace",
	})
	factomdinterfacesServerGetName = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_interfaces_Server_GetName_Summary",
		Help: "Summary of calls to  factomd_interfaces_Server_GetName",
	})
	factomdinterfacesServerGetChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_interfaces_Server_GetChainID_Summary",
		Help: "Summary of calls to  factomd_interfaces_Server_GetChainID",
	})
	factomdinterfacesServerString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_interfaces_Server_String_Summary",
		Help: "Summary of calls to  factomd_interfaces_Server_String",
	})
	factomdinterfacesServerIsOnline = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_interfaces_Server_IsOnline_Summary",
		Help: "Summary of calls to  factomd_interfaces_Server_IsOnline",
	})
	factomdinterfacesServerSetOnline = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_interfaces_Server_SetOnline_Summary",
		Help: "Summary of calls to  factomd_interfaces_Server_SetOnline",
	})
	factomdinterfacesServerLeaderToReplace = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_interfaces_Server_LeaderToReplace_Summary",
		Help: "Summary of calls to  factomd_interfaces_Server_LeaderToReplace",
	})
	factomdinterfacesServerSetReplace = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_interfaces_Server_SetReplace_Summary",
		Help: "Summary of calls to  factomd_interfaces_Server_SetReplace",
	})
)

func InitPrometheus() {
	prometheus.MustRegister(factomdinterfacesFctServerGetChainID)
	prometheus.MustRegister(factomdinterfacesFctServerString)
	prometheus.MustRegister(factomdinterfacesFctServerGetName)
	prometheus.MustRegister(factomdinterfacesFctServerIsOnline)
	prometheus.MustRegister(factomdinterfacesFctServerSetOnline)
	prometheus.MustRegister(factomdinterfacesFctServerLeaderToReplace)
	prometheus.MustRegister(factomdinterfacesFctServerSetReplace)
	prometheus.MustRegister(factomdinterfacesServerGetName)
	prometheus.MustRegister(factomdinterfacesServerGetChainID)
	prometheus.MustRegister(factomdinterfacesServerString)
	prometheus.MustRegister(factomdinterfacesServerIsOnline)
	prometheus.MustRegister(factomdinterfacesServerSetOnline)
	prometheus.MustRegister(factomdinterfacesServerLeaderToReplace)
	prometheus.MustRegister(factomdinterfacesServerSetReplace)
}
