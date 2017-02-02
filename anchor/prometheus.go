// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package anchor

import "github.com/prometheus/client_golang/prometheus"

var (
	//anchorRecords variables
	anchorErrors = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "factomd_anchor_error_responses",
		Help: "Number of error responses from anchor objects",
	})    
	anchorJSONByte = prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_anchor_JSONByte_Summary",
		Help: "Summary of anchor.JSONBytes call",
	})    
	anchorJSONString = prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_anchor_JSONString_Summary",
		Help: "Summary of anchor.JSONBytes call",
	})    
	anchorJSONBuffer = prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_anchor_JSONBuffer_Summary",
		Help: "Summary of anchor.JSONBuffer call",
	})    
	anchorStringCall = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_anchor_string_summary",
		Help: "Summary of factomd.anchor.String calls",
	})
	anchorMarshal = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_anchor_marshal_summary",
		Help: "Summary of factomd.anchor.Marshal calls",
	})
	anchorMarshalSign = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_anchor_marshalandsign_summary",
		Help: "Summary of factomd.anchor.MarshalAndSign calls",
	})
 	anchorMarshalSignV2 = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_anchor_marshalandsign_v2_summary",
		Help: "Summary of factomd.anchor.MarshalAndSignV2 calls",
	})
	anchorUnMarshal = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_anchor_unmarshal_summary",
		Help: "Summary of factomd.anchor.UnMarshal calls",
	})
   anchorUnmarshalAnchorRecord = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_anchor_UnmarshalAnchorRecord_summary",
		Help: "Summary of factomd.anchor.UnmarshalAnchorRecord calls",
	})
   anchorUnmarshalAndValidateAnchorRecord = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_anchor_UnmarshalAndValidateAnchorRecord_summary",
		Help: "Summary of factomd.anchor.UnmarshalAndValidateAnchorRecord calls",
	})
   anchorUnmarshalAndValidateAnchorRecordV2 = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_anchor_UnmarshalAndValidateAnchorRecordV2_summary",
		Help: "Summary of factomd.anchor.UnmarshalAndValidateAnchorRecordV2 calls",
	})
   anchorUnmarshalAndValidateAnchorEntryAnyVersion = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_anchor_UnmarshalAndValidateAnchorEntryAnyVersion_summary",
		Help: "Summary of factomd.anchor.UnmarshalAndValidateAnchorEntryAnyVersion calls",
	})
   anchorCreateAnchorRecordFromDBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_anchor_CreateAnchorRecordFromDBlock_summary",
		Help: "Summary of factomd.anchor.CreateAnchorRecordFromDBlock calls",
	})
   
)



func InitPrometheus() {
	//anchorRecord.go variables
	prometheus.MustRegister(AnchorErrors)
	prometheus.MustRegister(anchorJSONByte)
	prometheus.MustRegister(anchorJSONString)
	prometheus.MustRegister(anchorJSONBuffer)
	prometheus.MustRegister(anchorStringCall)
	prometheus.MustRegister(anchorMarshal)
	prometheus.MustRegister(anchorMarshalSign)
	prometheus.MustRegister(anchorMarshalSignV2)
	prometheus.MustRegister(anchorUnMarshal)
	prometheus.MustRegister(anchorUnmarshalAnchorRecord)
	prometheus.MustRegister(anchorUnmarshalAndValidateAnchorRecord)
	prometheus.MustRegister(anchorUnmarshalAndValidateAnchorRecordV2)
	prometheus.MustRegister(anchorUnmarshalAndValidateAnchorEntryAnyVersion)
	prometheus.MustRegister(anchorCreateAnchorRecordFromDBlock)
   

/* these are what is being added to each function to measure function time
    callTime := time.Now().UnixNano()

	runTime := time.Now().UnixNano() - callTime
	v1DBlockByHeightSummary.Observe(float64(runTime))
*/


}
