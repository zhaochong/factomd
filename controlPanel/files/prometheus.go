package files

import "github.com/prometheus/client_golang/prometheus"

var (
	factomdfilesOpen = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_files_Open_Summary",
		Help: "Summary of calls to  factomd_files_Open",
	})
	factomdfilesModTime = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_files_ModTime_Summary",
		Help: "Summary of calls to  factomd_files_ModTime",
	})
	factomdfilesHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_files_Hash_Summary",
		Help: "Summary of calls to  factomd_files_Hash",
	})
	factomdfilesOpenGlob = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_files_OpenGlob_Summary",
		Help: "Summary of calls to  factomd_files_OpenGlob",
	})
	factomdfilesCustomParseGlob = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_files_CustomParseGlob_Summary",
		Help: "Summary of calls to  factomd_files_CustomParseGlob",
	})
	factomdfilesCustomParseFile = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_files_CustomParseFile_Summary",
		Help: "Summary of calls to  factomd_files_CustomParseFile",
	})
	factomdfilesparseData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_files_parseData_Summary",
		Help: "Summary of calls to  factomd_files_parseData",
	})
)

func InitPrometheus() {
	prometheus.MustRegister(factomdfilesOpen)
	prometheus.MustRegister(factomdfilesModTime)
	prometheus.MustRegister(factomdfilesHash)
	prometheus.MustRegister(factomdfilesOpenGlob)
	prometheus.MustRegister(factomdfilesCustomParseGlob)
	prometheus.MustRegister(factomdfilesCustomParseFile)
	prometheus.MustRegister(factomdfilesparseData)
}
