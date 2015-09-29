package stateinit

import (
	. "github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/database"
	"github.com/FactomProject/factomd/database/ldb"
	. "github.com/FactomProject/factomd/state"
	"github.com/FactomProject/factomd/util"

	"fmt"

	"github.com/FactomProject/factomd/logger"
	"os"
)

func CreateLog(prefix string) *logger.FLogger {
	logcfg := util.ReadConfig().Log
	logPath := logcfg.LogPath
	logLevel := logcfg.LogLevel
	logfile, _ := os.OpenFile(logPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0660)
	return logger.New(logfile, logLevel, prefix)
}

// setup subsystem loggers
var ftmdLog = CreateLog("FTMD")
var procLog = CreateLog("PROC")

func InitState() *State {
	cfg := util.ReadConfig()
	LoadConfigurations(cfg)

	ldbpath := cfg.App.LdbPath
	boltDBpath := cfg.App.BoltDBPath

	db, factoidState := InitDB(ldbpath, boltDBpath)

	answer := new(State)
	answer.Cfg = cfg
	answer.DB = db
	answer.FactoidState = factoidState

	return answer
}

// Initialize the level db and share it with other components
func InitDB(boltDBpath, ldbpath string) (database.Db, IFactoidState) {

	//init factoid_bolt db
	fmt.Println("boltDBpath:", boltDBpath)
	factoidState := NewFactoidState(boltDBpath + "factoid_bolt.db")

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
