package log

import "github.com/prometheus/client_golang/prometheus"

var (
	factomdloginit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_log_init_Summary",
		Help: "Summary of calls to  factomd_log_init",
	})
	factomdlogSetTestLogger = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_log_SetTestLogger_Summary",
		Help: "Summary of calls to  factomd_log_SetTestLogger",
	})
	factomdlogUnsetTestLogger = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_log_UnsetTestLogger_Summary",
		Help: "Summary of calls to  factomd_log_UnsetTestLogger",
	})
	factomdlogSetLevel = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_log_SetLevel_Summary",
		Help: "Summary of calls to  factomd_log_SetLevel",
	})
	factomdlogdebugPrefix = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_log_debugPrefix_Summary",
		Help: "Summary of calls to  factomd_log_debugPrefix",
	})
	factomdlogPrintStack = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_log_PrintStack_Summary",
		Help: "Summary of calls to  factomd_log_PrintStack",
	})
	factomdlogFatal = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_log_Fatal_Summary",
		Help: "Summary of calls to  factomd_log_Fatal",
	})
	factomdlogPrint = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_log_Print_Summary",
		Help: "Summary of calls to  factomd_log_Print",
	})
	factomdlogPrintln = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_log_Println_Summary",
		Help: "Summary of calls to  factomd_log_Println",
	})
	factomdlogPrintf = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_log_Printf_Summary",
		Help: "Summary of calls to  factomd_log_Printf",
	})
	factomdlogPrintfln = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_log_Printfln_Summary",
		Help: "Summary of calls to  factomd_log_Printfln",
	})
	factomdlogDebug = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_log_Debug_Summary",
		Help: "Summary of calls to  factomd_log_Debug",
	})
	factomdlogprintf = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_log_printf_Summary",
		Help: "Summary of calls to  factomd_log_printf",
	})
)

func InitPrometheus() {
	prometheus.MustRegister(factomdloginit)
	prometheus.MustRegister(factomdlogSetTestLogger)
	prometheus.MustRegister(factomdlogUnsetTestLogger)
	prometheus.MustRegister(factomdlogSetLevel)
	prometheus.MustRegister(factomdlogdebugPrefix)
	prometheus.MustRegister(factomdlogPrintStack)
	prometheus.MustRegister(factomdlogFatal)
	prometheus.MustRegister(factomdlogPrint)
	prometheus.MustRegister(factomdlogPrintln)
	prometheus.MustRegister(factomdlogPrintf)
	prometheus.MustRegister(factomdlogPrintfln)
	prometheus.MustRegister(factomdlogDebug)
	prometheus.MustRegister(factomdlogprintf)
}
