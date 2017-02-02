// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package specialEntries

import "github.com/prometheus/client_golang/prometheus"

var (
	//anchorRecords variables
	ferEntryErrors = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "factomd_ferEntry_error_responses",
		Help: "Number of error responses from ferEntry objects",
	})
	specialEntriesGetVersion = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_ferEntry_GetVersion_Summary",
		Help: "Summary of ferEntry.GetVersion call",
	})
	specialEntriesSetVersion = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_ferEntry_SetVersion_Summary",
		Help: "Summary of ferEntry.SetVersion call",
	})
	specialEntriesGetExpirationHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_ferEntry_GetExpirationHeight_Summary",
		Help: "Summary of ferEntry.GetExpirationHeight call",
	})
	specialEntriesSetExpirationHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_ferEntry_SetExpirationHeight_Summary",
		Help: "Summary of ferEntry.SetExpirationHeight call",
	})
	specialEntriesGetResidentHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_ferEntry_GetResidentHeight_Summary",
		Help: "Summary of ferEntry.GetResidentHeight call",
	})
	specialEntriesSetResidentHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_ferEntry_SetResidentHeight_Summary",
		Help: "Summary of ferEntry.SetResidentHeight call",
	})
	specialEntriesGetTargetActivationHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_ferEntry_GetTargetActivationHeight_Summary",
		Help: "Summary of ferEntry.GetTargetActivationHeight call",
	})
	specialEntriesSetTargetActivationHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_ferEntry_SetTargetActivationHeight_Summary",
		Help: "Summary of ferEntry.SetTargetActivationHeight call",
	})
	specialEntriesGetPriority = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_ferEntry_GetPriority_Summary",
		Help: "Summary of ferEntry.GetPriority call",
	})
	specialEntriesSetPriority = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_ferEntry_SetPriority_Summary",
		Help: "Summary of ferEntry.SetPriority call",
	})
	specialEntriesGetTargetPrice = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_ferEntry_GetTargetPrice_Summary",
		Help: "Summary of ferEntry.GetTargetPrice call",
	})
	specialEntriesSetTargetPrice = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_ferEntry_SetTargetPrice_Summary",
		Help: "Summary of ferEntry.SetTargetPrice call",
	})
	specialEntriesJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_ferEntry_JSONByte_Summary",
		Help: "Summary of ferEntry.JSONByte call",
	})
	specialEntriesJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_ferEntry_JSONString_Summary",
		Help: "Summary of ferEntry.JSONString call",
	})
	specialEntriesJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_ferEntry_JSONBuffer_Summary",
		Help: "Summary of ferEntry.JSONBuffer call",
	})
	specialEntriesString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_ferEntry_String_Summary",
		Help: "Summary of ferEntry.String call",
	})
	specialEntriesUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_ferEntry_UnmarshalBinaryData_Summary",
		Help: "Summary of ferEntry.UnmarshalBinaryData call",
	})
	specialEntriesMarshalBinary= prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_ferEntry_MarshalBinary_Summary",
		Help: "Summary of ferEntry.MarshalBinary call",
	})
	
)

func InitPrometheus() {
	//wsapi variables
	prometheus.MustRegister(ferEntryErrors)
	prometheus.MustRegister(specialEntriesGetVersion)
	prometheus.MustRegister(specialEntriesSetVersion)
	prometheus.MustRegister(specialEntriesGetExpirationHeight)
	prometheus.MustRegister(specialEntriesSetExpirationHeight)
	prometheus.MustRegister(specialEntriesGetResidentHeight)
	prometheus.MustRegister(specialEntriesSetResidentHeight)
	prometheus.MustRegister(specialEntriesGetTargetActivationHeight)
	prometheus.MustRegister(specialEntriesSetTargetActivationHeight)
	prometheus.MustRegister(specialEntriesGetPriority)
	prometheus.MustRegister(specialEntriesSetPriority)
	prometheus.MustRegister(specialEntriesGetTargetPrice)
	prometheus.MustRegister(specialEntriesSetTargetPrice)
	prometheus.MustRegister(specialEntriesJSONByte)
	prometheus.MustRegister(specialEntriesJSONString)
	prometheus.MustRegister(specialEntriesJSONBuffer)
	prometheus.MustRegister(specialEntriesString)
	prometheus.MustRegister(specialEntriesUnmarshalBinaryData)
    prometheus.MustRegister(specialEntriesMarshalBinary)
}
