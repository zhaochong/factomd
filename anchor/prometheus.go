// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package anchor

import "github.com/prometheus/client_golang/prometheus"

var (
	//anchorRecords variables
	AnchorErrors = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "factomd_wsapi_requests_v1",
		Help: "Number of api v1 service requests",
	})    
	AnchorStringCall = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_anchor_string_summary",
		Help: "Summary of factomd.anchor.String calls",
	})
	AnchorMarshal = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_anchor_marshal_summary",
		Help: "Summary of factomd.anchor.Marshal calls",
	})
	AnchorMarshalSign = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_anchor_marshalandsign_summary",
		Help: "Summary of factomd.anchor.MarshalAndSign calls",
	})
 	AnchorMarshalSignV2 = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_anchor_marshalandsign_v2_summary",
		Help: "Summary of factomd.anchor.MarshalAndSignV2 calls",
	})
	AnchorUnMarshal = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_anchor_unmarshal_summary",
		Help: "Summary of factomd.anchor.UnMarshal calls",
	})
   
)

func InitPrometheus() {
	//wsapi variables
	prometheus.MustRegister(AnchorErrors)
	prometheus.MustRegister(AnchorStringCall)
	prometheus.MustRegister(AnchorMarshal)
	prometheus.MustRegister(AnchorMarshalSign)
	prometheus.MustRegister(AnchorMarshalSignV2)
	prometheus.MustRegister(AnchorUnMarshal)


/* these are what is being added to each function to measure function time
    callTime := time.Now().UnixNano()

	runTime := time.Now().UnixNano() - callTime
	v1DBlockByHeightSummary.Observe(float64(runTime))
*/


}
