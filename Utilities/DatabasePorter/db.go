package main

import (
	//"fmt"
	"os"

	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/database/databaseOverlay"
	"github.com/FactomProject/factomd/database/hybridDB"
	"github.com/FactomProject/factomd/database/mapdb"
	"github.com/FactomProject/factomd/util"
	"time"
)

//DBInit

func InitBolt(cfg *util.FactomdConfig) interfaces.DBOverlay {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmainInitBolt.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	//fmt.Println("InitBolt")
	path := cfg.App.BoltDBPath + "/"

	os.MkdirAll(path, 0777)
	dbase := hybridDB.NewBoltMapHybridDB(nil, path+"FactomBolt-Import.db")
	return databaseOverlay.NewOverlay(dbase)
}

func InitLevelDB(cfg *util.FactomdConfig) interfaces.DBOverlay {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmainInitLevelDB.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	//fmt.Println("InitLevelDB")
	path := cfg.App.LdbPath + "/" + "FactoidLevel-Import.db"

	dbase, err := hybridDB.NewLevelMapHybridDB(path, false)

	if err != nil || dbase == nil {
		dbase, err = hybridDB.NewLevelMapHybridDB(path, true)
		if err != nil {
			panic(err)
		}
	}

	return databaseOverlay.NewOverlay(dbase)
}

func InitMapDB(cfg *util.FactomdConfig) interfaces.DBOverlay {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmainInitMapDB.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	//fmt.Println("InitMapDB")
	dbase := new(mapdb.MapDB)
	dbase.Init(nil)
	return databaseOverlay.NewOverlay(dbase)
}
