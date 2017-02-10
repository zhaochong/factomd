package dataDumpFormatting

import "github.com/prometheus/client_golang/prometheus"

var (
	factomddataDumpFormattingRawProcessList = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dataDumpFormatting_RawProcessList_Summary",
		Help: "Summary of calls to  factomd_dataDumpFormatting_RawProcessList",
	})
	factomddataDumpFormattingRawPrintMap = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dataDumpFormatting_RawPrintMap_Summary",
		Help: "Summary of calls to  factomd_dataDumpFormatting_RawPrintMap",
	})
	factomddataDumpFormattingIdentities = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dataDumpFormatting_Identities_Summary",
		Help: "Summary of calls to  factomd_dataDumpFormatting_Identities",
	})
	factomddataDumpFormattingAuthorities = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dataDumpFormatting_Authorities_Summary",
		Help: "Summary of calls to  factomd_dataDumpFormatting_Authorities",
	})
	factomddataDumpFormattingMyNodeInfo = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dataDumpFormatting_MyNodeInfo_Summary",
		Help: "Summary of calls to  factomd_dataDumpFormatting_MyNodeInfo",
	})
	factomddataDumpFormattingreturnStatString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dataDumpFormatting_returnStatString_Summary",
		Help: "Summary of calls to  factomd_dataDumpFormatting_returnStatString",
	})
	factomddataDumpFormattingShortSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dataDumpFormatting_ShortSummary_Summary",
		Help: "Summary of calls to  factomd_dataDumpFormatting_ShortSummary",
	})
	factomddataDumpFormattingmessageLists = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_dataDumpFormatting_messageLists_Summary",
		Help: "Summary of calls to  factomd_dataDumpFormatting_messageLists",
	})
)

func InitPrometheus() {
	prometheus.MustRegister(factomddataDumpFormattingRawProcessList)
	prometheus.MustRegister(factomddataDumpFormattingRawPrintMap)
	prometheus.MustRegister(factomddataDumpFormattingIdentities)
	prometheus.MustRegister(factomddataDumpFormattingAuthorities)
	prometheus.MustRegister(factomddataDumpFormattingMyNodeInfo)
	prometheus.MustRegister(factomddataDumpFormattingreturnStatString)
	prometheus.MustRegister(factomddataDumpFormattingShortSummary)
	prometheus.MustRegister(factomddataDumpFormattingmessageLists)
}
