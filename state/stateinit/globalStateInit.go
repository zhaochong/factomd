package stateinit

import (
	. "github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/database/ldb"
	. "github.com/FactomProject/factomd/state"
	"github.com/FactomProject/factomd/util"
)

func InitState() *State {
	cfg := util.ReadConfig()
	LoadConfigurations(cfg)

	ldbpath := cfg.App.LdbPath
	boltDBpath := cfg.App.BoltDBPath

	db, factoidState := InitDB(ldbpath, boltDBpath)

	answer := new(DaemonState)
	answer.Cfg = cfg
	answer.DB = db
	answer.FactoidState = factoidState

	return answer
}

// Initialize the level db and share it with other components
func InitDB(boltDBpath, ldbpath string) (*ldb.LevelDb, IFactoidState) {

	//init factoid_bolt db
	fmt.Println("boltDBpath:", boltDBpath)
	factoidState := stateinit.NewFactoidState(boltDBpath + "factoid_bolt.db")

	db, err := ldb.OpenLevelDB(ldbpath, false)

	if err != nil {
		ftmdLog.Errorf("err opening db: %v\n", err)
	}

	if db == nil {
		ftmdLog.Info("Creating new db ...")
		db, err = ldb.OpenLevelDB(ldbpath, true)

		if err != nil {
			panic(err)
		}
	}
	ftmdLog.Info("Database started from: " + ldbpath)
	return db, factoidState
}
