package leveldb

import "github.com/prometheus/client_golang/prometheus"

var (
	factomdleveldbLevelDBListAllBuckets = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_leveldb_LevelDB_ListAllBuckets_Summary",
		Help: "Summary of calls to  factomd_leveldb_LevelDB_ListAllBuckets",
	})
	factomdleveldbLevelDBTrim = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_leveldb_LevelDB_Trim_Summary",
		Help: "Summary of calls to  factomd_leveldb_LevelDB_Trim",
	})
	factomdleveldbLevelDBDelete = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_leveldb_LevelDB_Delete_Summary",
		Help: "Summary of calls to  factomd_leveldb_LevelDB_Delete",
	})
	factomdleveldbLevelDBClose = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_leveldb_LevelDB_Close_Summary",
		Help: "Summary of calls to  factomd_leveldb_LevelDB_Close",
	})
	factomdleveldbExtendBucket = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_leveldb_ExtendBucket_Summary",
		Help: "Summary of calls to  factomd_leveldb_ExtendBucket",
	})
	factomdleveldbCombineBucketAndKey = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_leveldb_CombineBucketAndKey_Summary",
		Help: "Summary of calls to  factomd_leveldb_CombineBucketAndKey",
	})
	factomdleveldbLevelDBGet = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_leveldb_LevelDB_Get_Summary",
		Help: "Summary of calls to  factomd_leveldb_LevelDB_Get",
	})
	factomdleveldbLevelDBPut = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_leveldb_LevelDB_Put_Summary",
		Help: "Summary of calls to  factomd_leveldb_LevelDB_Put",
	})
	factomdleveldbLevelDBPutInBatch = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_leveldb_LevelDB_PutInBatch_Summary",
		Help: "Summary of calls to  factomd_leveldb_LevelDB_PutInBatch",
	})
	factomdleveldbLevelDBClear = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_leveldb_LevelDB_Clear_Summary",
		Help: "Summary of calls to  factomd_leveldb_LevelDB_Clear",
	})
	factomdleveldbLevelDBListAllKeys = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_leveldb_LevelDB_ListAllKeys_Summary",
		Help: "Summary of calls to  factomd_leveldb_LevelDB_ListAllKeys",
	})
	factomdleveldbLevelDBGetAll = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_leveldb_LevelDB_GetAll_Summary",
		Help: "Summary of calls to  factomd_leveldb_LevelDB_GetAll",
	})
	factomdleveldbNewLevelDB = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_leveldb_NewLevelDB_Summary",
		Help: "Summary of calls to  factomd_leveldb_NewLevelDB",
	})
	factomdleveldbaddOneToByteArray = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_leveldb_addOneToByteArray_Summary",
		Help: "Summary of calls to  factomd_leveldb_addOneToByteArray",
	})
	factomdleveldbLevelDBDoesKeyExist = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_leveldb_LevelDB_DoesKeyExist_Summary",
		Help: "Summary of calls to  factomd_leveldb_LevelDB_DoesKeyExist",
	})
)

func InitPrometheus() {
	prometheus.MustRegister(factomdleveldbLevelDBListAllBuckets)
	prometheus.MustRegister(factomdleveldbLevelDBTrim)
	prometheus.MustRegister(factomdleveldbLevelDBDelete)
	prometheus.MustRegister(factomdleveldbLevelDBClose)
	prometheus.MustRegister(factomdleveldbExtendBucket)
	prometheus.MustRegister(factomdleveldbCombineBucketAndKey)
	prometheus.MustRegister(factomdleveldbLevelDBGet)
	prometheus.MustRegister(factomdleveldbLevelDBPut)
	prometheus.MustRegister(factomdleveldbLevelDBPutInBatch)
	prometheus.MustRegister(factomdleveldbLevelDBClear)
	prometheus.MustRegister(factomdleveldbLevelDBListAllKeys)
	prometheus.MustRegister(factomdleveldbLevelDBGetAll)
	prometheus.MustRegister(factomdleveldbNewLevelDB)
	prometheus.MustRegister(factomdleveldbaddOneToByteArray)
	prometheus.MustRegister(factomdleveldbLevelDBDoesKeyExist)
}
