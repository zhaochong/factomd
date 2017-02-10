package boltdb

import "github.com/prometheus/client_golang/prometheus"

var (
	factomdboltdbNewBoltDB = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_boltdb_NewBoltDB_Summary",
		Help: "Summary of calls to  factomd_boltdb_NewBoltDB",
	})
	factomdboltdbBoltDBListAllBuckets = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_boltdb_BoltDB_ListAllBuckets_Summary",
		Help: "Summary of calls to  factomd_boltdb_BoltDB_ListAllBuckets",
	})
	factomdboltdbBoltDBDelete = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_boltdb_BoltDB_Delete_Summary",
		Help: "Summary of calls to  factomd_boltdb_BoltDB_Delete",
	})
	factomdboltdbBoltDBTrim = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_boltdb_BoltDB_Trim_Summary",
		Help: "Summary of calls to  factomd_boltdb_BoltDB_Trim",
	})
	factomdboltdbBoltDBClose = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_boltdb_BoltDB_Close_Summary",
		Help: "Summary of calls to  factomd_boltdb_BoltDB_Close",
	})
	factomdboltdbBoltDBGet = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_boltdb_BoltDB_Get_Summary",
		Help: "Summary of calls to  factomd_boltdb_BoltDB_Get",
	})
	factomdboltdbBoltDBPut = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_boltdb_BoltDB_Put_Summary",
		Help: "Summary of calls to  factomd_boltdb_BoltDB_Put",
	})
	factomdboltdbBoltDBPutInBatch = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_boltdb_BoltDB_PutInBatch_Summary",
		Help: "Summary of calls to  factomd_boltdb_BoltDB_PutInBatch",
	})
	factomdboltdbBoltDBClear = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_boltdb_BoltDB_Clear_Summary",
		Help: "Summary of calls to  factomd_boltdb_BoltDB_Clear",
	})
	factomdboltdbBoltDBListAllKeys = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_boltdb_BoltDB_ListAllKeys_Summary",
		Help: "Summary of calls to  factomd_boltdb_BoltDB_ListAllKeys",
	})
	factomdboltdbBoltDBGetAll = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_boltdb_BoltDB_GetAll_Summary",
		Help: "Summary of calls to  factomd_boltdb_BoltDB_GetAll",
	})
	factomdboltdbBoltDBInit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_boltdb_BoltDB_Init_Summary",
		Help: "Summary of calls to  factomd_boltdb_BoltDB_Init",
	})
	factomdboltdbBoltDBDoesKeyExist = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_boltdb_BoltDB_DoesKeyExist_Summary",
		Help: "Summary of calls to  factomd_boltdb_BoltDB_DoesKeyExist",
	})
)

func InitPrometheus() {
	prometheus.MustRegister(factomdboltdbNewBoltDB)
	prometheus.MustRegister(factomdboltdbBoltDBListAllBuckets)
	prometheus.MustRegister(factomdboltdbBoltDBDelete)
	prometheus.MustRegister(factomdboltdbBoltDBTrim)
	prometheus.MustRegister(factomdboltdbBoltDBClose)
	prometheus.MustRegister(factomdboltdbBoltDBGet)
	prometheus.MustRegister(factomdboltdbBoltDBPut)
	prometheus.MustRegister(factomdboltdbBoltDBPutInBatch)
	prometheus.MustRegister(factomdboltdbBoltDBClear)
	prometheus.MustRegister(factomdboltdbBoltDBListAllKeys)
	prometheus.MustRegister(factomdboltdbBoltDBGetAll)
	prometheus.MustRegister(factomdboltdbBoltDBInit)
	prometheus.MustRegister(factomdboltdbBoltDBDoesKeyExist)
}
