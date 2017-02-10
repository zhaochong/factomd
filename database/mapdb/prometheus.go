package mapdb

import "github.com/prometheus/client_golang/prometheus"

var (
	factomdmapdbClose = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_mapdb_Close(__Summary",
		Help: "Summary of calls to  factomd_mapdb_Close(_",
	})
	factomdmapdbMapDBListAllBuckets = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_mapdb_MapDB_ListAllBuckets_Summary",
		Help: "Summary of calls to  factomd_mapdb_MapDB_ListAllBuckets",
	})
	factomdmapdbMapDBTrim = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_mapdb_MapDB_Trim_Summary",
		Help: "Summary of calls to  factomd_mapdb_MapDB_Trim",
	})
	factomdmapdbMapDBcreateCache = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_mapdb_MapDB_createCache_Summary",
		Help: "Summary of calls to  factomd_mapdb_MapDB_createCache",
	})
	factomdmapdbMapDBInit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_mapdb_MapDB_Init_Summary",
		Help: "Summary of calls to  factomd_mapdb_MapDB_Init",
	})
	factomdmapdbMapDBPut = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_mapdb_MapDB_Put_Summary",
		Help: "Summary of calls to  factomd_mapdb_MapDB_Put",
	})
	factomdmapdbMapDBRawPut = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_mapdb_MapDB_RawPut_Summary",
		Help: "Summary of calls to  factomd_mapdb_MapDB_RawPut",
	})
	factomdmapdbMapDBrawPut = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_mapdb_MapDB_rawPut_Summary",
		Help: "Summary of calls to  factomd_mapdb_MapDB_rawPut",
	})
	factomdmapdbMapDBPutInBatch = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_mapdb_MapDB_PutInBatch_Summary",
		Help: "Summary of calls to  factomd_mapdb_MapDB_PutInBatch",
	})
	factomdmapdbMapDBGet = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_mapdb_MapDB_Get_Summary",
		Help: "Summary of calls to  factomd_mapdb_MapDB_Get",
	})
	factomdmapdbMapDBDelete = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_mapdb_MapDB_Delete_Summary",
		Help: "Summary of calls to  factomd_mapdb_MapDB_Delete",
	})
	factomdmapdbMapDBListAllKeys = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_mapdb_MapDB_ListAllKeys_Summary",
		Help: "Summary of calls to  factomd_mapdb_MapDB_ListAllKeys",
	})
	factomdmapdbMapDBGetAll = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_mapdb_MapDB_GetAll_Summary",
		Help: "Summary of calls to  factomd_mapdb_MapDB_GetAll",
	})
	factomdmapdbMapDBClear = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_mapdb_MapDB_Clear_Summary",
		Help: "Summary of calls to  factomd_mapdb_MapDB_Clear",
	})
	factomdmapdbMapDBDoesKeyExist = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_mapdb_MapDB_DoesKeyExist_Summary",
		Help: "Summary of calls to  factomd_mapdb_MapDB_DoesKeyExist",
	})
)

func InitPrometheus() {
	prometheus.MustRegister(factomdmapdbClose)
	prometheus.MustRegister(factomdmapdbMapDBListAllBuckets)
	prometheus.MustRegister(factomdmapdbMapDBTrim)
	prometheus.MustRegister(factomdmapdbMapDBcreateCache)
	prometheus.MustRegister(factomdmapdbMapDBInit)
	prometheus.MustRegister(factomdmapdbMapDBPut)
	prometheus.MustRegister(factomdmapdbMapDBRawPut)
	prometheus.MustRegister(factomdmapdbMapDBrawPut)
	prometheus.MustRegister(factomdmapdbMapDBPutInBatch)
	prometheus.MustRegister(factomdmapdbMapDBGet)
	prometheus.MustRegister(factomdmapdbMapDBDelete)
	prometheus.MustRegister(factomdmapdbMapDBListAllKeys)
	prometheus.MustRegister(factomdmapdbMapDBGetAll)
	prometheus.MustRegister(factomdmapdbMapDBClear)
	prometheus.MustRegister(factomdmapdbMapDBDoesKeyExist)
}
