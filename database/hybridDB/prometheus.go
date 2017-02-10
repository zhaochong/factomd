package hybridDB

import "github.com/prometheus/client_golang/prometheus"

var (
	factomdhybridDBHybridDBListAllBuckets = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_hybridDB_HybridDB_ListAllBuckets_Summary",
		Help: "Summary of calls to  factomd_hybridDB_HybridDB_ListAllBuckets",
	})
	factomdhybridDBHybridDBTrim = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_hybridDB_HybridDB_Trim_Summary",
		Help: "Summary of calls to  factomd_hybridDB_HybridDB_Trim",
	})
	factomdhybridDBHybridDBClose = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_hybridDB_HybridDB_Close_Summary",
		Help: "Summary of calls to  factomd_hybridDB_HybridDB_Close",
	})
	factomdhybridDBNewLevelMapHybridDB = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_hybridDB_NewLevelMapHybridDB_Summary",
		Help: "Summary of calls to  factomd_hybridDB_NewLevelMapHybridDB",
	})
	factomdhybridDBNewBoltMapHybridDB = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_hybridDB_NewBoltMapHybridDB_Summary",
		Help: "Summary of calls to  factomd_hybridDB_NewBoltMapHybridDB",
	})
	factomdhybridDBHybridDBPut = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_hybridDB_HybridDB_Put_Summary",
		Help: "Summary of calls to  factomd_hybridDB_HybridDB_Put",
	})
	factomdhybridDBHybridDBPutInBatch = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_hybridDB_HybridDB_PutInBatch_Summary",
		Help: "Summary of calls to  factomd_hybridDB_HybridDB_PutInBatch",
	})
	factomdhybridDBHybridDBGet = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_hybridDB_HybridDB_Get_Summary",
		Help: "Summary of calls to  factomd_hybridDB_HybridDB_Get",
	})
	factomdhybridDBHybridDBDelete = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_hybridDB_HybridDB_Delete_Summary",
		Help: "Summary of calls to  factomd_hybridDB_HybridDB_Delete",
	})
	factomdhybridDBHybridDBListAllKeys = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_hybridDB_HybridDB_ListAllKeys_Summary",
		Help: "Summary of calls to  factomd_hybridDB_HybridDB_ListAllKeys",
	})
	factomdhybridDBHybridDBGetAll = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_hybridDB_HybridDB_GetAll_Summary",
		Help: "Summary of calls to  factomd_hybridDB_HybridDB_GetAll",
	})
	factomdhybridDBHybridDBClear = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_hybridDB_HybridDB_Clear_Summary",
		Help: "Summary of calls to  factomd_hybridDB_HybridDB_Clear",
	})
	factomdhybridDBHybridDBDoesKeyExist = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_hybridDB_HybridDB_DoesKeyExist_Summary",
		Help: "Summary of calls to  factomd_hybridDB_HybridDB_DoesKeyExist",
	})
)

func InitPrometheus() {
	prometheus.MustRegister(factomdhybridDBHybridDBListAllBuckets)
	prometheus.MustRegister(factomdhybridDBHybridDBTrim)
	prometheus.MustRegister(factomdhybridDBHybridDBClose)
	prometheus.MustRegister(factomdhybridDBNewLevelMapHybridDB)
	prometheus.MustRegister(factomdhybridDBNewBoltMapHybridDB)
	prometheus.MustRegister(factomdhybridDBHybridDBPut)
	prometheus.MustRegister(factomdhybridDBHybridDBPutInBatch)
	prometheus.MustRegister(factomdhybridDBHybridDBGet)
	prometheus.MustRegister(factomdhybridDBHybridDBDelete)
	prometheus.MustRegister(factomdhybridDBHybridDBListAllKeys)
	prometheus.MustRegister(factomdhybridDBHybridDBGetAll)
	prometheus.MustRegister(factomdhybridDBHybridDBClear)
	prometheus.MustRegister(factomdhybridDBHybridDBDoesKeyExist)
}
