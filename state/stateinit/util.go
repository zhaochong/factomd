package stateinit

import (
	"fmt"
	"github.com/FactomProject/factomd/logger"
	"github.com/FactomProject/factomd/util"

	"sort"

	. "github.com/FactomProject/factomd/state/processState"

	. "github.com/FactomProject/factomd/common"
	. "github.com/FactomProject/factomd/common/AdminBlock"
	. "github.com/FactomProject/factomd/common/DirectoryBlock"
	. "github.com/FactomProject/factomd/common/EntryBlock"
	. "github.com/FactomProject/factomd/common/EntryCreditBlock"
	. "github.com/FactomProject/factomd/common/interfaces"
)

func exportDChain(ps *ProcessState, chain *DChain) {
	if len(chain.Blocks) == 0 || procLog.Level() < logger.Info {
		//log.Println("no blocks to save for chain: " + string (*chain.ChainID))
		return
	}

	// get all ecBlocks from db
	dBlocks, _ := ps.DB.FetchAllDBlocks()
	sort.Sort(util.ByDBlockIDAccending(dBlocks))

	for _, block := range dBlocks {

		data, err := block.MarshalBinary()
		if err != nil {
			panic(err)
		}

		strChainID := chain.ChainID.String()
		if fileNotExists(dataStorePath + strChainID) {
			err := os.MkdirAll(dataStorePath+strChainID, 0777)
			if err == nil {
				procLog.Info("Created directory " + dataStorePath + strChainID)
			} else {
				procLog.Error(err)
			}
		}
		err = ioutil.WriteFile(fmt.Sprintf(dataStorePath+strChainID+"/store.%09d.block", block.Header.DBHeight), data, 0777)
		if err != nil {
			panic(err)
		}
	}
}

func exportEChain(ps *ProcessState, chain *EChain) {
	if procLog.Level() < logger.Info {
		return
	}

	eBlocks, _ := ps.DB.FetchAllEBlocksByChain(chain.ChainID)
	sort.Sort(util.ByEBlockIDAccending(*eBlocks))

	for _, block := range *eBlocks {

		data, err := block.MarshalBinary()
		if err != nil {
			panic(err)
		}

		strChainID := chain.ChainID.String()
		if fileNotExists(dataStorePath + strChainID) {
			err := os.MkdirAll(dataStorePath+strChainID, 0777)
			if err == nil {
				procLog.Info("Created directory " + dataStorePath + strChainID)
			} else {
				procLog.Error(err)
			}
		}

		err = ioutil.WriteFile(fmt.Sprintf(dataStorePath+strChainID+"/store.%09d.%09d.block", block.Header.EBSequence, block.Header.DBHeight), data, 0777)
		if err != nil {
			panic(err)
		}
	}
}

func exportECChain(ps *ProcessState, chain *ECChain) {
	if procLog.Level() < logger.Info {
		return
	}
	// get all ecBlocks from db
	ecBlocks, _ := ps.DB.FetchAllECBlocks()
	sort.Sort(util.ByECBlockIDAccending(ecBlocks))

	for _, block := range ecBlocks {
		data, err := block.MarshalBinary()
		if err != nil {
			panic(err)
		}

		strChainID := chain.ChainID.String()
		if fileNotExists(dataStorePath + strChainID) {
			err := os.MkdirAll(dataStorePath+strChainID, 0777)
			if err == nil {
				procLog.Info("Created directory " + dataStorePath + strChainID)
			} else {
				procLog.Error(err)
			}
		}
		err = ioutil.WriteFile(fmt.Sprintf(dataStorePath+strChainID+"/store.%09d.block", block.Header.DBHeight), data, 0777)
		if err != nil {
			panic(err)
		}
	}
}

func exportAChain(ps *ProcessState, chain *AdminChain) {
	if procLog.Level() < logger.Info {
		return
	}
	// get all aBlocks from db
	aBlocks, _ := ps.DB.FetchAllABlocks()
	sort.Sort(util.ByABlockIDAccending(aBlocks))

	for _, block := range aBlocks {

		data, err := block.MarshalBinary()
		if err != nil {
			panic(err)
		}

		strChainID := chain.ChainID.String()
		if fileNotExists(dataStorePath + strChainID) {
			err := os.MkdirAll(dataStorePath+strChainID, 0777)
			if err == nil {
				procLog.Info("Created directory " + dataStorePath + strChainID)
			} else {
				procLog.Error(err)
			}
		}
		err = ioutil.WriteFile(fmt.Sprintf(dataStorePath+strChainID+"/store.%09d.block", block.Header.DBHeight), data, 0777)
		if err != nil {
			panic(err)
		}
	}
}

func exportFctChain(ps *ProcessState, chain *FctChain) {
	if procLog.Level() < logger.Info {
		return
	}
	// get all aBlocks from db
	FBlocks, _ := ps.DB.FetchAllFBlocks()
	sort.Sort(util.ByFBlockIDAccending(FBlocks))

	for _, block := range FBlocks {

		data, err := block.MarshalBinary()
		if err != nil {
			panic(err)
		}

		strChainID := chain.ChainID.String()
		if fileNotExists(dataStorePath + strChainID) {
			err := os.MkdirAll(dataStorePath+strChainID, 0777)
			if err == nil {
				procLog.Info("Created directory " + dataStorePath + strChainID)
			} else {
				procLog.Error(err)
			}
		}
		err = ioutil.WriteFile(fmt.Sprintf(dataStorePath+strChainID+"/store.%09d.block", block.GetDBHeight()), data, 0777)
		if err != nil {
			panic(err)
		}
	}
}

// to export individual block once at a time - for debugging ------------------------
func exportDBlock(ps *ProcessState, block *DirectoryBlock) {
	if block == nil || procLog.Level() < logger.Info {
		//log.Println("no blocks to save for chain: " + string (*chain.ChainID))
		return
	}

	data, err := block.MarshalBinary()
	if err != nil {
		panic(err)
	}

	strChainID := dchain.ChainID.String()
	if fileNotExists(dataStorePath + strChainID) {
		err := os.MkdirAll(dataStorePath+strChainID, 0777)
		if err == nil {
			procLog.Info("Created directory " + dataStorePath + strChainID)
		} else {
			procLog.Error(err)
		}
	}
	err = ioutil.WriteFile(fmt.Sprintf(dataStorePath+strChainID+"/store.%09d.block", block.Header.DBHeight), data, 0777)
	if err != nil {
		panic(err)
	}

}

func exportEBlock(ps *ProcessState, block *EBlock) {
	if block == nil || procLog.Level() < logger.Info {
		return
	}

	data, err := block.MarshalBinary()
	if err != nil {
		panic(err)
	}

	strChainID := block.Header.ChainID.String()
	if fileNotExists(dataStorePath + strChainID) {
		err := os.MkdirAll(dataStorePath+strChainID, 0777)
		if err == nil {
			procLog.Info("Created directory " + dataStorePath + strChainID)
		} else {
			procLog.Error(err)
		}
	}

	err = ioutil.WriteFile(fmt.Sprintf(dataStorePath+strChainID+"/store.%09d.%09d.block", block.Header.EBSequence, block.Header.DBHeight), data, 0777)
	if err != nil {
		panic(err)
	}

}

func exportECBlock(ps *ProcessState, block *ECBlock) {
	if block == nil || procLog.Level() < logger.Info {
		return
	}

	data, err := block.MarshalBinary()
	if err != nil {
		panic(err)
	}

	strChainID := block.Header.ECChainID.String()
	if fileNotExists(dataStorePath + strChainID) {
		err := os.MkdirAll(dataStorePath+strChainID, 0777)
		if err == nil {
			procLog.Info("Created directory " + dataStorePath + strChainID)
		} else {
			procLog.Error(err)
		}
	}
	err = ioutil.WriteFile(fmt.Sprintf(dataStorePath+strChainID+"/store.%09d.block", block.Header.DBHeight), data, 0777)
	if err != nil {
		panic(err)
	}

}

func exportABlock(ps *ProcessState, block *AdminBlock) {
	if block == nil || procLog.Level() < logger.Info {
		return
	}

	data, err := block.MarshalBinary()
	if err != nil {
		panic(err)
	}

	strChainID := block.Header.AdminChainID.String()
	if fileNotExists(dataStorePath + strChainID) {
		err := os.MkdirAll(dataStorePath+strChainID, 0777)
		if err == nil {
			procLog.Info("Created directory " + dataStorePath + strChainID)
		} else {
			procLog.Error(err)
		}
	}
	err = ioutil.WriteFile(fmt.Sprintf(dataStorePath+strChainID+"/store.%09d.block", block.Header.DBHeight), data, 0777)
	if err != nil {
		panic(err)
	}

}

func exportFctBlock(ps *ProcessState, block IFBlock) {
	if block == nil || procLog.Level() < logger.Info {
		return
	}

	data, err := block.MarshalBinary()
	if err != nil {
		panic(err)
	}

	strChainID := block.GetChainID().String()
	if fileNotExists(dataStorePath + strChainID) {
		err := os.MkdirAll(dataStorePath+strChainID, 0777)
		if err == nil {
			procLog.Info("Created directory " + dataStorePath + strChainID)
		} else {
			procLog.Error(err)
		}
	}
	err = ioutil.WriteFile(fmt.Sprintf(dataStorePath+strChainID+"/store.%09d.block", block.GetDBHeight()), data, 0777)
	if err != nil {
		panic(err)
	}

}
