package random

import "github.com/prometheus/client_golang/prometheus"

var (
	factomdrandomRandUInt64 = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_random_RandUInt64_Summary",
		Help: "Summary of calls to  factomd_random_RandUInt64",
	})
	factomdrandomRandUInt64Between = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_random_RandUInt64Between_Summary",
		Help: "Summary of calls to  factomd_random_RandUInt64Between",
	})
	factomdrandomRandInt64 = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_random_RandInt64_Summary",
		Help: "Summary of calls to  factomd_random_RandInt64",
	})
	factomdrandomRandInt64Between = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_random_RandInt64Between_Summary",
		Help: "Summary of calls to  factomd_random_RandInt64Between",
	})
	factomdrandomRandInt = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_random_RandInt_Summary",
		Help: "Summary of calls to  factomd_random_RandInt",
	})
	factomdrandomRandIntBetween = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_random_RandIntBetween_Summary",
		Help: "Summary of calls to  factomd_random_RandIntBetween",
	})
	factomdrandomRandByteSlice = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_random_RandByteSlice_Summary",
		Help: "Summary of calls to  factomd_random_RandByteSlice",
	})
	factomdrandomRandNonEmptyByteSlice = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_random_RandNonEmptyByteSlice_Summary",
		Help: "Summary of calls to  factomd_random_RandNonEmptyByteSlice",
	})
	factomdrandomRandByteSliceOfLen = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_random_RandByteSliceOfLen_Summary",
		Help: "Summary of calls to  factomd_random_RandByteSliceOfLen",
	})
	factomdrandomRandomString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_random_RandomString_Summary",
		Help: "Summary of calls to  factomd_random_RandomString",
	})
)

func InitPrometheus() {
	prometheus.MustRegister(factomdrandomRandUInt64)
	prometheus.MustRegister(factomdrandomRandUInt64Between)
	prometheus.MustRegister(factomdrandomRandInt64)
	prometheus.MustRegister(factomdrandomRandInt64Between)
	prometheus.MustRegister(factomdrandomRandInt)
	prometheus.MustRegister(factomdrandomRandIntBetween)
	prometheus.MustRegister(factomdrandomRandByteSlice)
	prometheus.MustRegister(factomdrandomRandNonEmptyByteSlice)
	prometheus.MustRegister(factomdrandomRandByteSliceOfLen)
	prometheus.MustRegister(factomdrandomRandomString)
}
