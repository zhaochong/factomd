package specialEntries

import "github.com/prometheus/client_golang/prometheus"

var (
	factomdspecialEntriesFEREntryGetVersion = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_specialEntries_FEREntry_GetVersion_Summary",
		Help: "Summary of calls to  factomd_specialEntries_FEREntry_GetVersion",
	})
	factomdspecialEntriesFEREntrySetVersion = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_specialEntries_FEREntry_SetVersion_Summary",
		Help: "Summary of calls to  factomd_specialEntries_FEREntry_SetVersion",
	})
	factomdspecialEntriesFEREntryGetExpirationHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_specialEntries_FEREntry_GetExpirationHeight_Summary",
		Help: "Summary of calls to  factomd_specialEntries_FEREntry_GetExpirationHeight",
	})
	factomdspecialEntriesFEREntrySetExpirationHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_specialEntries_FEREntry_SetExpirationHeight_Summary",
		Help: "Summary of calls to  factomd_specialEntries_FEREntry_SetExpirationHeight",
	})
	factomdspecialEntriesFEREntryGetResidentHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_specialEntries_FEREntry_GetResidentHeight_Summary",
		Help: "Summary of calls to  factomd_specialEntries_FEREntry_GetResidentHeight",
	})
	factomdspecialEntriesFEREntrySetResidentHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_specialEntries_FEREntry_SetResidentHeight_Summary",
		Help: "Summary of calls to  factomd_specialEntries_FEREntry_SetResidentHeight",
	})
	factomdspecialEntriesFEREntryGetTargetActivationHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_specialEntries_FEREntry_GetTargetActivationHeight_Summary",
		Help: "Summary of calls to  factomd_specialEntries_FEREntry_GetTargetActivationHeight",
	})
	factomdspecialEntriesFEREntrySetTargetActivationHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_specialEntries_FEREntry_SetTargetActivationHeight_Summary",
		Help: "Summary of calls to  factomd_specialEntries_FEREntry_SetTargetActivationHeight",
	})
	factomdspecialEntriesFEREntryGetPriority = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_specialEntries_FEREntry_GetPriority_Summary",
		Help: "Summary of calls to  factomd_specialEntries_FEREntry_GetPriority",
	})
	factomdspecialEntriesFEREntrySetPriority = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_specialEntries_FEREntry_SetPriority_Summary",
		Help: "Summary of calls to  factomd_specialEntries_FEREntry_SetPriority",
	})
	factomdspecialEntriesFEREntryGetTargetPrice = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_specialEntries_FEREntry_GetTargetPrice_Summary",
		Help: "Summary of calls to  factomd_specialEntries_FEREntry_GetTargetPrice",
	})
	factomdspecialEntriesFEREntrySetTargetPrice = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_specialEntries_FEREntry_SetTargetPrice_Summary",
		Help: "Summary of calls to  factomd_specialEntries_FEREntry_SetTargetPrice",
	})
	factomdspecialEntriesFEREntryJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_specialEntries_FEREntry_JSONByte_Summary",
		Help: "Summary of calls to  factomd_specialEntries_FEREntry_JSONByte",
	})
	factomdspecialEntriesFEREntryJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_specialEntries_FEREntry_JSONString_Summary",
		Help: "Summary of calls to  factomd_specialEntries_FEREntry_JSONString",
	})
	factomdspecialEntriesFEREntryString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_specialEntries_FEREntry_String_Summary",
		Help: "Summary of calls to  factomd_specialEntries_FEREntry_String",
	})
	factomdspecialEntriesFEREntryUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_specialEntries_FEREntry_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_specialEntries_FEREntry_UnmarshalBinaryData",
	})
	factomdspecialEntriesFEREntryUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_specialEntries_FEREntry_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_specialEntries_FEREntry_UnmarshalBinary",
	})
	factomdspecialEntriesFEREntryMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_specialEntries_FEREntry_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_specialEntries_FEREntry_MarshalBinary",
	})
)

func InitPrometheus() {
	prometheus.MustRegister(factomdspecialEntriesFEREntryGetVersion)
	prometheus.MustRegister(factomdspecialEntriesFEREntrySetVersion)
	prometheus.MustRegister(factomdspecialEntriesFEREntryGetExpirationHeight)
	prometheus.MustRegister(factomdspecialEntriesFEREntrySetExpirationHeight)
	prometheus.MustRegister(factomdspecialEntriesFEREntryGetResidentHeight)
	prometheus.MustRegister(factomdspecialEntriesFEREntrySetResidentHeight)
	prometheus.MustRegister(factomdspecialEntriesFEREntryGetTargetActivationHeight)
	prometheus.MustRegister(factomdspecialEntriesFEREntrySetTargetActivationHeight)
	prometheus.MustRegister(factomdspecialEntriesFEREntryGetPriority)
	prometheus.MustRegister(factomdspecialEntriesFEREntrySetPriority)
	prometheus.MustRegister(factomdspecialEntriesFEREntryGetTargetPrice)
	prometheus.MustRegister(factomdspecialEntriesFEREntrySetTargetPrice)
	prometheus.MustRegister(factomdspecialEntriesFEREntryJSONByte)
	prometheus.MustRegister(factomdspecialEntriesFEREntryJSONString)
	prometheus.MustRegister(factomdspecialEntriesFEREntryString)
	prometheus.MustRegister(factomdspecialEntriesFEREntryUnmarshalBinaryData)
	prometheus.MustRegister(factomdspecialEntriesFEREntryUnmarshalBinary)
	prometheus.MustRegister(factomdspecialEntriesFEREntryMarshalBinary)
}
