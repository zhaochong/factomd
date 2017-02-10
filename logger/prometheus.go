package logger

import "github.com/prometheus/client_golang/prometheus"

var (
	factomdloggerNewLogFromConfig = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_logger_NewLogFromConfig_Summary",
		Help: "Summary of calls to  factomd_logger_NewLogFromConfig",
	})
	factomdloggerNew = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_logger_New_Summary",
		Help: "Summary of calls to  factomd_logger_New",
	})
	factomdloggerFLoggerLevel = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_logger_FLogger_Level_Summary",
		Help: "Summary of calls to  factomd_logger_FLogger_Level",
	})
	factomdloggerFLoggerEmergency = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_logger_FLogger_Emergency_Summary",
		Help: "Summary of calls to  factomd_logger_FLogger_Emergency",
	})
	factomdloggerFLoggerEmergencyf = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_logger_FLogger_Emergencyf_Summary",
		Help: "Summary of calls to  factomd_logger_FLogger_Emergencyf",
	})
	factomdloggerFLoggerAlert = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_logger_FLogger_Alert_Summary",
		Help: "Summary of calls to  factomd_logger_FLogger_Alert",
	})
	factomdloggerFLoggerAlertf = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_logger_FLogger_Alertf_Summary",
		Help: "Summary of calls to  factomd_logger_FLogger_Alertf",
	})
	factomdloggerFLoggerCritical = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_logger_FLogger_Critical_Summary",
		Help: "Summary of calls to  factomd_logger_FLogger_Critical",
	})
	factomdloggerFLoggerCriticalf = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_logger_FLogger_Criticalf_Summary",
		Help: "Summary of calls to  factomd_logger_FLogger_Criticalf",
	})
	factomdloggerFLoggerError = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_logger_FLogger_Error_Summary",
		Help: "Summary of calls to  factomd_logger_FLogger_Error",
	})
	factomdloggerFLoggerErrorf = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_logger_FLogger_Errorf_Summary",
		Help: "Summary of calls to  factomd_logger_FLogger_Errorf",
	})
	factomdloggerFLoggerWarning = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_logger_FLogger_Warning_Summary",
		Help: "Summary of calls to  factomd_logger_FLogger_Warning",
	})
	factomdloggerFLoggerWarningf = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_logger_FLogger_Warningf_Summary",
		Help: "Summary of calls to  factomd_logger_FLogger_Warningf",
	})
	factomdloggerFLoggerNotice = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_logger_FLogger_Notice_Summary",
		Help: "Summary of calls to  factomd_logger_FLogger_Notice",
	})
	factomdloggerFLoggerNoticef = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_logger_FLogger_Noticef_Summary",
		Help: "Summary of calls to  factomd_logger_FLogger_Noticef",
	})
	factomdloggerFLoggerInfo = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_logger_FLogger_Info_Summary",
		Help: "Summary of calls to  factomd_logger_FLogger_Info",
	})
	factomdloggerFLoggerInfof = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_logger_FLogger_Infof_Summary",
		Help: "Summary of calls to  factomd_logger_FLogger_Infof",
	})
	factomdloggerFLoggerDebug = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_logger_FLogger_Debug_Summary",
		Help: "Summary of calls to  factomd_logger_FLogger_Debug",
	})
	factomdloggerFLoggerDebugf = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_logger_FLogger_Debugf_Summary",
		Help: "Summary of calls to  factomd_logger_FLogger_Debugf",
	})
	factomdloggerFLoggerwrite = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_logger_FLogger_write_Summary",
		Help: "Summary of calls to  factomd_logger_FLogger_write",
	})
	factomdloggerlevelFromString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_logger_levelFromString_Summary",
		Help: "Summary of calls to  factomd_logger_levelFromString",
	})
)

func InitPrometheus() {
	prometheus.MustRegister(factomdloggerNewLogFromConfig)
	prometheus.MustRegister(factomdloggerNew)
	prometheus.MustRegister(factomdloggerFLoggerLevel)
	prometheus.MustRegister(factomdloggerFLoggerEmergency)
	prometheus.MustRegister(factomdloggerFLoggerEmergencyf)
	prometheus.MustRegister(factomdloggerFLoggerAlert)
	prometheus.MustRegister(factomdloggerFLoggerAlertf)
	prometheus.MustRegister(factomdloggerFLoggerCritical)
	prometheus.MustRegister(factomdloggerFLoggerCriticalf)
	prometheus.MustRegister(factomdloggerFLoggerError)
	prometheus.MustRegister(factomdloggerFLoggerErrorf)
	prometheus.MustRegister(factomdloggerFLoggerWarning)
	prometheus.MustRegister(factomdloggerFLoggerWarningf)
	prometheus.MustRegister(factomdloggerFLoggerNotice)
	prometheus.MustRegister(factomdloggerFLoggerNoticef)
	prometheus.MustRegister(factomdloggerFLoggerInfo)
	prometheus.MustRegister(factomdloggerFLoggerInfof)
	prometheus.MustRegister(factomdloggerFLoggerDebug)
	prometheus.MustRegister(factomdloggerFLoggerDebugf)
	prometheus.MustRegister(factomdloggerFLoggerwrite)
	prometheus.MustRegister(factomdloggerlevelFromString)
}
