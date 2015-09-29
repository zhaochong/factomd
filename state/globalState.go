package state

import (
	. "github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/database"
	"github.com/FactomProject/factomd/util"
)

type State struct {
	FactoidState IFactoidState

	Cfg *util.FactomdConfig
	DB  database.Db // database
}
