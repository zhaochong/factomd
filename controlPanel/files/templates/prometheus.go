package templates

import "github.com/prometheus/client_golang/prometheus"

var (
	factomdtemplatesServeHTTP = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_templates_ServeHTTP_Summary",
		Help: "Summary of calls to  factomd_templates_ServeHTTP",
	})
	factomdtemplatesOpen = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_templates_Open_Summary",
		Help: "Summary of calls to  factomd_templates_Open",
	})
	factomdtemplatesModTime = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_templates_ModTime_Summary",
		Help: "Summary of calls to  factomd_templates_ModTime",
	})
	factomdtemplatesHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_templates_Hash_Summary",
		Help: "Summary of calls to  factomd_templates_Hash",
	})
	factomdtemplatesOpenGlob = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_templates_OpenGlob_Summary",
		Help: "Summary of calls to  factomd_templates_OpenGlob",
	})
)

func InitPrometheus() {
	prometheus.MustRegister(factomdtemplatesServeHTTP)
	prometheus.MustRegister(factomdtemplatesOpen)
	prometheus.MustRegister(factomdtemplatesModTime)
	prometheus.MustRegister(factomdtemplatesHash)
	prometheus.MustRegister(factomdtemplatesOpenGlob)
}
