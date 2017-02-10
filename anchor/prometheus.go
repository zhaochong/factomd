package anchor

import "github.com/prometheus/client_golang/prometheus"

var (
	factomdanchorAnchorRecordJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_anchor_AnchorRecord_JSONByte_Summary",
		Help: "Summary of calls to  factomd_anchor_AnchorRecord_JSONByte",
	})
	factomdanchorAnchorRecordJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_anchor_AnchorRecord_JSONString_Summary",
		Help: "Summary of calls to  factomd_anchor_AnchorRecord_JSONString",
	})
	factomdanchorAnchorRecordString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_anchor_AnchorRecord_String_Summary",
		Help: "Summary of calls to  factomd_anchor_AnchorRecord_String",
	})
	factomdanchorAnchorRecordMarshal = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_anchor_AnchorRecord_Marshal_Summary",
		Help: "Summary of calls to  factomd_anchor_AnchorRecord_Marshal",
	})
	factomdanchorAnchorRecordMarshalAndSign = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_anchor_AnchorRecord_MarshalAndSign_Summary",
		Help: "Summary of calls to  factomd_anchor_AnchorRecord_MarshalAndSign",
	})
	factomdanchorAnchorRecordMarshalAndSignV2 = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_anchor_AnchorRecord_MarshalAndSignV2_Summary",
		Help: "Summary of calls to  factomd_anchor_AnchorRecord_MarshalAndSignV2",
	})
	factomdanchorAnchorRecordUnmarshal = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_anchor_AnchorRecord_Unmarshal_Summary",
		Help: "Summary of calls to  factomd_anchor_AnchorRecord_Unmarshal",
	})
	factomdanchorUnmarshalAnchorRecord = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_anchor_UnmarshalAnchorRecord_Summary",
		Help: "Summary of calls to  factomd_anchor_UnmarshalAnchorRecord",
	})
	factomdanchorUnmarshalAndValidateAnchorRecord = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_anchor_UnmarshalAndValidateAnchorRecord_Summary",
		Help: "Summary of calls to  factomd_anchor_UnmarshalAndValidateAnchorRecord",
	})
	factomdanchorUnmarshalAndValidateAnchorRecordV2 = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_anchor_UnmarshalAndValidateAnchorRecordV2_Summary",
		Help: "Summary of calls to  factomd_anchor_UnmarshalAndValidateAnchorRecordV2",
	})
	factomdanchorUnmarshalAndValidateAnchorEntryAnyVersion = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_anchor_UnmarshalAndValidateAnchorEntryAnyVersion_Summary",
		Help: "Summary of calls to  factomd_anchor_UnmarshalAndValidateAnchorEntryAnyVersion",
	})
	factomdanchorCreateAnchorRecordFromDBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_anchor_CreateAnchorRecordFromDBlock_Summary",
		Help: "Summary of calls to  factomd_anchor_CreateAnchorRecordFromDBlock",
	})
)

func InitPrometheus() {
	prometheus.MustRegister(factomdanchorAnchorRecordJSONByte)
	prometheus.MustRegister(factomdanchorAnchorRecordJSONString)
	prometheus.MustRegister(factomdanchorAnchorRecordString)
	prometheus.MustRegister(factomdanchorAnchorRecordMarshal)
	prometheus.MustRegister(factomdanchorAnchorRecordMarshalAndSign)
	prometheus.MustRegister(factomdanchorAnchorRecordMarshalAndSignV2)
	prometheus.MustRegister(factomdanchorAnchorRecordUnmarshal)
	prometheus.MustRegister(factomdanchorUnmarshalAnchorRecord)
	prometheus.MustRegister(factomdanchorUnmarshalAndValidateAnchorRecord)
	prometheus.MustRegister(factomdanchorUnmarshalAndValidateAnchorRecordV2)
	prometheus.MustRegister(factomdanchorUnmarshalAndValidateAnchorEntryAnyVersion)
	prometheus.MustRegister(factomdanchorCreateAnchorRecordFromDBlock)
}
