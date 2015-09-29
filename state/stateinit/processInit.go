// Copyright 2015 FactomProject Authors. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package stateinit

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/FactomProject/factomd/anchor"
	"github.com/FactomProject/factomd/btcd/wire"
	fct "github.com/FactomProject/factomd/common/factoid"
	"github.com/FactomProject/factomd/common/factoid/block"
	"github.com/FactomProject/factomd/consensus"
	cp "github.com/FactomProject/factomd/controlpanel"
	"github.com/FactomProject/factomd/database"
	"github.com/FactomProject/factomd/logger"
	"github.com/FactomProject/factomd/util"
	"github.com/davecgh/go-spew/spew"
	"sort"
	"strconv"

	. "github.com/FactomProject/factomd/state/processState"

	. "github.com/FactomProject/factomd/common"
	. "github.com/FactomProject/factomd/common/AdminBlock"
	. "github.com/FactomProject/factomd/common/DirectoryBlock"
	. "github.com/FactomProject/factomd/common/EntryBlock"
	. "github.com/FactomProject/factomd/common/EntryCreditBlock"
	. "github.com/FactomProject/factomd/common/constants"
	. "github.com/FactomProject/factomd/common/interfaces"
	. "github.com/FactomProject/factomd/common/primitives"
)

// Get the configurations
func LoadConfigurations(cfg *util.FactomdConfig) {

	//setting the variables by the valued form the config file
	/*logLevel = cfg.Log.LogLevel
	dataStorePath = cfg.App.DataStorePath
	ldbpath = cfg.App.LdbPath
	directoryBlockInSeconds = cfg.App.DirectoryBlockInSeconds
	ps.NodeMode = cfg.App.NodeMode
	serverPrivKeyHex = cfg.App.ServerPrivKey

	FactomdUser = cfg.Btc.RpcUser
	FactomdPass = cfg.Btc.RpcPass*/
}

func InitProcessState(
	ldb database.Db,
	inMsgQ chan wire.FtmInternalMsg,
	outMsgQ chan wire.FtmInternalMsg,
	inCtlMsgQ chan wire.FtmInternalMsg,
	outCtlMsgQ chan wire.FtmInternalMsg) *ProcessState {
	ps := new(ProcessState)

	ps.CommitChainMap = make(map[string]*CommitChain, 0)
	ps.CommitEntryMap = make(map[string]*CommitEntry, 0)
	ps.ServerIndex = NewServerIndexNumber()

	ps.DB = ldb

	ps.InMsgQueue = inMsgQ
	ps.OutMsgQueue = outMsgQ

	ps.InCtlMsgQueue = inCtlMsgQ
	ps.OutCtlMsgQueue = outCtlMsgQ

	// init server private key or pub key
	initServerKeys(ps)

	// init mem pools
	ps.FMemPool = new(FtmMemPool)
	ps.FMemPool.Init_FtmMemPool()

	ps.FactoshisPerCredit = 666666 // .001 / .15 * 100000000 (assuming a Factoid is .15 cents, entry credit = .1 cents

	// init Directory Block Chain
	initDChain(ps)

	procLog.Info("Loaded ", ps.DChain.NextDBHeight, " Directory blocks for chain: "+ps.DChain.ChainID.String())

	// init Entry Credit Chain
	initECChain(ps)
	procLog.Info("Loaded ", ps.ECChain.NextBlockHeight, " Entry Credit blocks for chain: "+ps.ECChain.ChainID.String())

	// init Admin Chain
	initAChain(ps)
	procLog.Info("Loaded ", ps.AChain.NextBlockHeight, " Admin blocks for chain: "+ps.AChain.ChainID.String())

	initFctChain(ps)
	//ps.LoadState()
	procLog.Info("Loaded ", ps.FChain.NextBlockHeight, " factoid blocks for chain: "+ps.FChain.ChainID.String())

	//Init anchor for server
	if ps.NodeMode == SERVER_NODE {
		anchor.InitAnchor(ps.DB, ps.InMsgQueue, ps.ServerPrivKey)
	}
	// build the Genesis blocks if the current height is 0
	if ps.DChain.NextDBHeight == 0 && ps.NodeMode == SERVER_NODE {
		buildGenesisBlocks(ps)
	} else {
		// To be improved in milestone 2
		SignDirectoryBlock(ps)
	}

	// init process list manager
	initProcessListMgr(ps)

	// init Entry Chains
	initEChains(ps)
	for _, chain := range ps.ChainIDMap {
		initEChainFromDB(chain)

		procLog.Info("Loaded ", chain.NextBlockHeight, " blocks for chain: "+chain.ChainID.String())
	}

	// Validate all dir blocks
	err := validateDChain(ps.DChain)
	if err != nil {
		if ps.NodeMode == SERVER_NODE {
			panic("Error found in validating directory blocks: " + err.Error())
		} else {
			ps.DChain.IsValidated = false
		}
	}
	return ps
}

// Initialize Directory Block Chain from database
func initDChain(ps *ProcessState) {
	ps.DChain = new(DChain)

	//Initialize the Directory Block Chain ID
	ps.DChain.ChainID = new(Hash)
	barray := D_CHAINID
	ps.DChain.ChainID.SetBytes(barray)

	// get all dBlocks from db
	dBlocks, _ := ps.DB.FetchAllDBlocks()
	sort.Sort(util.ByDBlockIDAccending(dBlocks))

	ps.DChain.Blocks = make([]*DirectoryBlock, len(dBlocks), len(dBlocks)+1)

	for i := 0; i < len(dBlocks); i = i + 1 {
		if dBlocks[i].Header.DBHeight != uint32(i) {
			panic("Error in initializing dChain:" + ps.DChain.ChainID.String())
		}
		dBlocks[i].Chain = ps.DChain
		dBlocks[i].IsSealed = true
		dBlocks[i].IsSavedInDB = true
		ps.DChain.Blocks[i] = &dBlocks[i]
	}

	// double check the block ids
	for i := 0; i < len(ps.DChain.Blocks); i = i + 1 {
		if uint32(i) != ps.DChain.Blocks[i].Header.DBHeight {
			panic(errors.New("BlockID does not equal index for chain:" + ps.DChain.ChainID.String() + " block:" + fmt.Sprintf("%v", ps.DChain.Blocks[i].Header.DBHeight)))
		}
	}

	//Create an empty block and append to the chain
	if len(ps.DChain.Blocks) == 0 {
		ps.DChain.NextDBHeight = 0
		ps.DChain.NextBlock, _ = CreateDBlock(ps.DChain, nil, 10)
	} else {
		ps.DChain.NextDBHeight = uint32(len(ps.DChain.Blocks))
		ps.DChain.NextBlock, _ = CreateDBlock(ps.DChain, ps.DChain.Blocks[len(ps.DChain.Blocks)-1], 10)
		// Update dir block height cache in db
		ps.DB.UpdateBlockHeightCache(ps.DChain.NextDBHeight-1, ps.DChain.NextBlock.Header.PrevLedgerKeyMR)
	}

	exportDChain(ps, ps.DChain)

	//Double check the sealed flag
	if ps.DChain.NextBlock.IsSealed == true {
		panic("ps.DChain.Blocks[ps.DChain.NextBlockID].IsSealed for chain:" + ps.DChain.ChainID.String())
	}

}

// Initialize Entry Credit Block Chain from database
func initECChain(ps *ProcessState) {

	ps.ECreditMap = make(map[string]int32)

	//Initialize the Entry Credit Chain ID
	ps.ECChain = NewECChain()

	// get all ecBlocks from db
	ecBlocks, _ := ps.DB.FetchAllECBlocks()
	sort.Sort(util.ByECBlockIDAccending(ecBlocks))

	for i, v := range ecBlocks {
		if v.Header.DBHeight != uint32(i) {
			panic("Error in initializing dChain:" + ps.ECChain.ChainID.String() + " DBHeight:" + strconv.Itoa(int(v.Header.DBHeight)) + " i:" + strconv.Itoa(i))
		}

		// Calculate the EC balance for each account
		initializeECreditMap(ps, &v)
	}

	//Create an empty block and append to the chain
	if len(ecBlocks) == 0 || ps.DChain.NextDBHeight == 0 {
		ps.ECChain.NextBlockHeight = 0
		ps.ECChain.NextBlock = NewECBlock()
		ps.ECChain.NextBlock.AddEntry(ps.ServerIndex)
		for i := 0; i < 10; i++ {
			marker := NewMinuteNumber()
			marker.Number = uint8(i + 1)
			ps.ECChain.NextBlock.AddEntry(marker)
		}
	} else {
		// Entry Credit Chain should have the same height as the dir chain
		ps.ECChain.NextBlockHeight = ps.DChain.NextDBHeight
		var err error
		ps.ECChain.NextBlock, err = NextECBlock(&ecBlocks[ps.ECChain.NextBlockHeight-1])
		if err != nil {
			panic(err)
		}
	}

	// create a backup copy before processing entries
	copyCreditMap(ps.ECreditMap, ps.ECreditMapBackup)
	exportECChain(ps, ps.ECChain)

	// ONly for debugging
	if procLog.Level() > logger.Info {
		printCreditMap()
	}

}

// Initialize Admin Block Chain from database
func initAChain(ps *ProcessState) {

	//Initialize the Admin Chain ID
	ps.AChain = new(AdminChain)
	ps.AChain.ChainID = new(Hash)
	ps.AChain.ChainID.SetBytes(ADMIN_CHAINID)

	// get all aBlocks from db
	aBlocks, _ := ps.DB.FetchAllABlocks()
	sort.Sort(util.ByABlockIDAccending(aBlocks))

	// double check the block ids
	for i := 0; i < len(aBlocks); i = i + 1 {
		if uint32(i) != aBlocks[i].Header.DBHeight {
			panic(errors.New("BlockID does not equal index for chain:" + ps.AChain.ChainID.String() + " block:" + fmt.Sprintf("%v", aBlocks[i].Header.DBHeight)))
		}
		if !validateDBSignature(&aBlocks[i], ps.DChain) {
			panic(errors.New("No valid signature found in Admin Block = " + fmt.Sprintf("%s\n", spew.Sdump(aBlocks[i]))))
		}
	}

	//Create an empty block and append to the chain
	if len(aBlocks) == 0 || ps.DChain.NextDBHeight == 0 {
		ps.AChain.NextBlockHeight = 0
		ps.AChain.NextBlock, _ = CreateAdminBlock(ps.AChain, nil, 10)

	} else {
		// Entry Credit Chain should have the same height as the dir chain
		ps.AChain.NextBlockHeight = ps.DChain.NextDBHeight
		ps.AChain.NextBlock, _ = CreateAdminBlock(ps.AChain, &aBlocks[ps.AChain.NextBlockHeight-1], 10)
	}

	exportAChain(ps, ps.AChain)

}

// Initialize Factoid Block Chain from database
func initFctChain(ps *ProcessState) {

	//Initialize the Admin Chain ID
	ps.FChain = new(FctChain)
	ps.FChain.ChainID = new(Hash)
	ps.FChain.ChainID.SetBytes(FACTOID_CHAINID)

	// get all aBlocks from db
	fBlocks, _ := ps.DB.FetchAllFBlocks()
	sort.Sort(util.ByFBlockIDAccending(fBlocks))

	// double check the block ids
	for i := 0; i < len(fBlocks); i = i + 1 {
		if uint32(i) != fBlocks[i].GetDBHeight() {
			panic(errors.New("BlockID does not equal index for chain:" +
				ps.FChain.ChainID.String() + " block:" +
				fmt.Sprintf("%v", fBlocks[i].GetDBHeight())))
		} else {
			ps.FactoshisPerCredit = fBlocks[i].GetExchRate()
			ps.SetFactoshisPerEC(ps.FactoshisPerCredit)
			// initialize the FactoidState in sequence
			err := ps.AddTransactionBlock(fBlocks[i])
			if err != nil {
				panic("Failed to rebuild factoid state: " + err.Error())
			}
		}
	}

	//Create an empty block and append to the chain
	if len(fBlocks) == 0 || ps.DChain.NextDBHeight == 0 {
		ps.SetFactoshisPerEC(ps.FactoshisPerCredit)
		ps.FChain.NextBlockHeight = 0
		// func GetGenesisFBlock(ftime uint64, ExRate uint64, addressCnt int, Factoids uint64 ) IFBlock {
		//ps.FChain.NextBlock = block.GetGenesisFBlock(0, ps.FactoshisPerCredit, 10, 200000000000)
		ps.FChain.NextBlock = block.GetGenesisFBlock()
		gb := ps.FChain.NextBlock

		// If a client, this block is going to get downloaded and added.  Don't do it twice.
		if ps.NodeMode == SERVER_NODE {
			err := ps.AddTransactionBlock(gb)
			if err != nil {
				panic(err)
			}
		}

	} else {
		ps.FChain.NextBlockHeight = ps.DChain.NextDBHeight
		ps.ProcessEndOfBlock2(ps.DChain.NextDBHeight)
		ps.FChain.NextBlock = ps.GetCurrentBlock()
	}

	exportFctChain(ps, ps.FChain)

}

// Initialize Entry Block Chains from database
func initEChains(ps *ProcessState) {
	ps.ChainIDMap = make(map[string]*EChain)

	chains, err := ps.DB.FetchAllChains()

	if err != nil {
		panic(err)
	}

	for _, chain := range chains {
		var newChain = chain
		ps.ChainIDMap[newChain.ChainID.String()] = newChain
		exportEChain(ps, chain)
	}

}

// Re-calculate Entry Credit Balance Map with a new Entry Credit Block
func initializeECreditMap(ps *ProcessState, block *ECBlock) {
	for _, entry := range block.Body.Entries {
		// Only process: ECIDChainCommit, ECIDEntryCommit, ECIDBalanceIncrease
		switch entry.ECID() {
		case ECIDChainCommit:
			e := entry.(*CommitChain)
			ps.ECreditMap[string(e.ECPubKey[:])] -= int32(e.Credits)
			ps.UpdateECBalance(fct.NewAddress(e.ECPubKey[:]), int64(e.Credits))
		case ECIDEntryCommit:
			e := entry.(*CommitEntry)
			ps.ECreditMap[string(e.ECPubKey[:])] -= int32(e.Credits)
			ps.UpdateECBalance(fct.NewAddress(e.ECPubKey[:]), int64(e.Credits))
		case ECIDBalanceIncrease:
			e := entry.(*IncreaseBalance)
			ps.ECreditMap[string(e.ECPubKey[:])] += int32(e.NumEC)
			// Don't add the Increases to Factoid state, the Factoid processing will do that.
		case ECIDServerIndexNumber:
		case ECIDMinuteNumber:
		default:
			panic("Unknow entry type:" + string(entry.ECID()) + " for ECBlock:" + strconv.FormatUint(uint64(block.Header.DBHeight), 10))
		}
	}
}

// Initialize server private key and server public key for milestone 1
func initServerKeys(ps *ProcessState) {
	if ps.NodeMode == SERVER_NODE {
		var err error
		ps.ServerPrivKey, err = NewPrivateKeyFromHex(serverPrivKeyHex)
		if err != nil {
			panic("Cannot parse Server Private Key from configuration file: " + err.Error())
		}
	}

	ps.ServerPubKey = PubKeyFromString(SERVER_PUB_KEY)
}

// Initialize the process list manager with the proper dir block height
func initProcessListMgr(ps *ProcessState) {
	ps.PlMgr = consensus.NewProcessListMgr(ps.DChain.NextDBHeight, 1, 10, serverPrivKey)

}

// Initialize the entry chains in memory from db
func initEChainFromDB(chain *EChain) {

	eBlocks, _ := ps.DB.FetchAllEBlocksByChain(chain.ChainID)
	sort.Sort(util.ByEBlockIDAccending(*eBlocks))

	for i := 0; i < len(*eBlocks); i = i + 1 {
		if uint32(i) != (*eBlocks)[i].Header.EBSequence {
			panic(errors.New("BlockID does not equal index for chain:" + chain.ChainID.String() + " block:" + fmt.Sprintf("%v", (*eBlocks)[i].Header.EBSequence)))
		}
	}

	var err error
	if len(*eBlocks) == 0 {
		chain.NextBlockHeight = 0
		chain.NextBlock, err = MakeEBlock(chain, nil)
		if err != nil {
			panic(err)
		}
	} else {
		chain.NextBlockHeight = uint32(len(*eBlocks))
		chain.NextBlock, err = MakeEBlock(chain, &(*eBlocks)[len(*eBlocks)-1])
		if err != nil {
			panic(err)
		}
	}

	// Initialize chain with the first entry (Name and rules) for non-server mode
	if ps.NodeMode != SERVER_NODE && chain.FirstEntry == nil && len(*eBlocks) > 0 {
		chain.FirstEntry, _ = ps.DB.FetchEntryByHash((*eBlocks)[0].Body.EBEntries[0])
		if chain.FirstEntry != nil {
			ps.DB.InsertChain(chain)
		}
	}

}

// Validate dir chain from genesis block
func validateDChain(c *DChain) error {

	if ps.NodeMode != SERVER_NODE && len(c.Blocks) == 0 {
		return nil
	}

	if uint32(len(c.Blocks)) != c.NextDBHeight {
		return errors.New("Dir chain has an un-expected Next Block ID: " + strconv.Itoa(int(c.NextDBHeight)))
	}

	//prevMR and prevBlkHash are used to validate against the block next in the chain
	prevMR, prevBlkHash, err := validateDBlock(c, c.Blocks[0])
	if err != nil {
		return err
	}

	//validate the genesis block
	//prevBlkHash is the block hash for c.Blocks[0]
	if prevBlkHash == nil || prevBlkHash.String() != GENESIS_DIR_BLOCK_HASH {

		str := fmt.Sprintf("<pre>" +
			"Expected: " + GENESIS_DIR_BLOCK_HASH + "<br>" +
			"Found:    " + prevBlkHash.String() + "</pre><br><br>")
		cp.CP.AddUpdate(
			"GenHash",                    // tag
			"warning",                    // Category
			"Genesis Hash doesn't match", // Title
			str, // Message
			0)
		// panic for Milestone 1
		panic("Genesis Block wasn't as expected:\n" +
			"    Expected: " + GENESIS_DIR_BLOCK_HASH + "\n" +
			"    Found:    " + prevBlkHash.String())
	}

	for i := 1; i < len(c.Blocks); i++ {
		if !prevBlkHash.IsSameAs(c.Blocks[i].Header.PrevLedgerKeyMR) {
			return errors.New("Previous block hash not matching for Dir block: " + strconv.Itoa(i))
		}
		if !prevMR.IsSameAs(c.Blocks[i].Header.PrevKeyMR) {
			return errors.New("Previous merkle root not matching for Dir block: " + strconv.Itoa(i))
		}
		mr, dblkHash, err := validateDBlock(c, c.Blocks[i])
		if err != nil {
			c.Blocks[i].IsValidated = false
			return err
		}

		prevMR = mr
		prevBlkHash = dblkHash
		c.Blocks[i].IsValidated = true
	}

	return nil
}

// Validate a dir block
func validateDBlock(c *DChain, b *DirectoryBlock) (merkleRoot IHash, dbHash IHash, err error) {

	bodyMR, err := b.BuildBodyMR()
	if err != nil {
		return nil, nil, err
	}

	if !b.Header.BodyMR.IsSameAs(bodyMR) {
		return nil, nil, errors.New("Invalid body MR for dir block: " + string(b.Header.DBHeight))
	}

	for _, dbEntry := range b.DBEntries {
		switch dbEntry.ChainID.String() {
		case ps.ECChain.ChainID.String():
			err := validateCBlockByMR(dbEntry.KeyMR)
			if err != nil {
				return nil, nil, err
			}
		case ps.AChain.ChainID.String():
			err := validateABlockByMR(dbEntry.KeyMR)
			if err != nil {
				return nil, nil, err
			}
		case wire.FChainID.String():
			err := validateFBlockByMR(dbEntry.KeyMR)
			if err != nil {
				return nil, nil, err
			}
		default:
			err := validateEBlockByMR(dbEntry.ChainID, dbEntry.KeyMR)
			if err != nil {
				return nil, nil, err
			}
		}
	}

	b.DBHash, _ = CreateHash(b)
	b.BuildKeyMerkleRoot()

	return b.KeyMR, b.DBHash, nil
}

// Validate Entry Credit Block by merkle root
func validateCBlockByMR(mr IHash) error {
	cb, _ := ps.DB.FetchECBlockByHash(mr)

	if cb == nil {
		return errors.New("Entry Credit block not found in db for merkle root: " + mr.String())
	}

	return nil
}

// Validate Admin Block by merkle root
func validateABlockByMR(mr IHash) error {
	b, _ := ps.DB.FetchABlockByHash(mr)

	if b == nil {
		return errors.New("Admin block not found in db for merkle root: " + mr.String())
	}

	return nil
}

// Validate FBlock by merkle root
func validateFBlockByMR(mr IHash) error {
	b, _ := ps.DB.FetchFBlockByHash(mr)

	if b == nil {
		return errors.New("Factoid block not found in db for merkle root: \n" + mr.String())
	}

	// check that we used the KeyMR to store the block...
	if !bytes.Equal(b.GetKeyMR().Bytes(), mr.Bytes()) {
		return fmt.Errorf("Factoid block match failure: block %d \n%s\n%s",
			b.GetDBHeight(),
			"Key in the database:   "+mr.String(),
			"Hash of the blk found: "+b.GetKeyMR().String())
	}

	return nil
}

// Validate Entry Block by merkle root
func validateEBlockByMR(cid IHash, mr IHash) error {

	eb, err := ps.DB.FetchEBlockByMR(mr)
	if err != nil {
		return err
	}

	if eb == nil {
		return errors.New("Entry block not found in db for merkle root: " + mr.String())
	}
	keyMR, err := eb.KeyMR()
	if err != nil {
		return err
	}
	if !mr.IsSameAs(keyMR) {
		return errors.New("Entry block's merkle root does not match with: " + mr.String())
	}

	for _, ebEntry := range eb.Body.EBEntries {
		if !bytes.Equal(ebEntry.Bytes()[:31], ZERO_HASH[:31]) {
			entry, _ := ps.DB.FetchEntryByHash(ebEntry)
			if entry == nil {
				return errors.New("Entry not found in db for entry hash: " + ebEntry.String())
			}
		} // Else ... we could do a bit more validation of the minute markers.
	}

	return nil
}

// build Genesis blocks
func buildGenesisBlocks(ps *ProcessState) error {
	//Set the timestamp for the genesis block
	t, err := time.Parse(time.RFC3339, GENESIS_BLK_TIMESTAMP)
	if err != nil {
		panic("Not able to parse the genesis block time stamp")
	}
	ps.DChain.NextBlock.Header.Timestamp = uint32(t.Unix() / 60)

	// Allocate the first two dbentries for ECBlock and Factoid block
	ps.DChain.AddDBEntry(&DBEntry{}) // AdminBlock
	ps.DChain.AddDBEntry(&DBEntry{}) // ECBlock
	ps.DChain.AddDBEntry(&DBEntry{}) // Factoid block

	// Entry Credit Chain
	cBlock := newEntryCreditBlock(ecchain)
	procLog.Debugf("buildGenesisBlocks: cBlock=%s\n", spew.Sdump(cBlock))
	ps.DChain.AddECBlockToDBEntry(cBlock)
	exportECChain(ecchain)

	// Admin chain
	aBlock := newAdminBlock(achain)
	procLog.Debugf("buildGenesisBlocks: aBlock=%s\n", spew.Sdump(aBlock))
	ps.DChain.AddABlockToDBEntry(aBlock)
	exportAChain(achain)

	// factoid Genesis Address
	//fchain.NextBlock = block.GetGenesisFBlock(0, ps.FactoshisPerCredit, 10, 200000000000)
	fchain.NextBlock = block.GetGenesisFBlock()
	FBlock := newFactoidBlock(fchain)
	ps.DChain.AddFBlockToDBEntry(FBlock)
	exportFctChain(fchain)

	// Directory Block chain
	procLog.Debug("in buildGenesisBlocks")
	dbBlock := newDirectoryBlock(dchain)

	// Check block hash if genesis block
	if dbBlock.DBHash.String() != GENESIS_DIR_BLOCK_HASH {
		//Panic for Milestone 1
		panic("\nGenesis block hash expected: " + GENESIS_DIR_BLOCK_HASH +
			"\nGenesis block hash found:    " + dbBlock.DBHash.String() + "\n")
	}

	exportDChain(dchain)

	// place an anchor into btc
	placeAnchor(dbBlock)

	return nil
}

// build blocks from all process lists
func buildBlocks() error {

	// Allocate the first three dbentries for Admin block, ECBlock and Factoid block
	ps.DChain.AddDBEntry(&DBEntry{}) // AdminBlock
	ps.DChain.AddDBEntry(&DBEntry{}) // ECBlock
	ps.DChain.AddDBEntry(&DBEntry{}) // factoid

	if plMgr != nil && plMgr.MyProcessList.IsValid() {
		buildFromProcessList(plMgr.MyProcessList)
	}

	// Entry Credit Chain
	ecBlock := newEntryCreditBlock(ecchain)
	ps.DChain.AddECBlockToDBEntry(ecBlock)
	exportECBlock(ecBlock)

	// Admin chain
	aBlock := newAdminBlock(achain)

	ps.DChain.AddABlockToDBEntry(aBlock)
	exportABlock(aBlock)

	// Factoid chain
	fBlock := newFactoidBlock(fchain)

	ps.DChain.AddFBlockToDBEntry(fBlock)
	exportFctBlock(fBlock)

	// sort the echains by chain id
	var keys []string
	for k := range chainIDMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Entry Chains
	for _, k := range keys {
		chain := chainIDMap[k]
		eblock := newEntryBlock(chain)
		if eblock != nil {
			ps.DChain.AddEBlockToDBEntry(eblock)
		}
		exportEBlock(eblock)
	}

	// Directory Block chain
	procLog.Debug("in buildBlocks")
	dbBlock := newDirectoryBlock(dchain)

	// Generate the inventory vector and relay it.
	binary, _ := dbBlock.MarshalBinary()
	commonHash := Sha(binary)
	hash, _ := NewShaHash(commonHash.Bytes())
	outMsgQueue <- (&wire.MsgInt_DirBlock{hash})

	// Update dir block height cache in db
	db.UpdateBlockHeightCache(dbBlock.Header.DBHeight, commonHash)
	db.UpdateNextBlockHeightCache(ps.DChain.NextDBHeight)

	exportDBlock(dbBlock)

	// re-initialize the process lit manager
	initProcessListMgr()

	// Initialize timer for the new dblock
	if ps.NodeMode == SERVER_NODE {
		timer := &BlockTimer{
			nextDBlockHeight: ps.DChain.NextDBHeight,
			inCtlMsgQueue:    inCtlMsgQueue,
		}
		go timer.StartBlockTimer()
	}

	// place an anchor into btc
	placeAnchor(dbBlock)

	return nil
}

// Sign the directory block
func SignDirectoryBlock(ps *ProcessState) error {
	// Only Servers can write the anchor to Bitcoin network
	if ps.NodeMode == SERVER_NODE && ps.DChain.NextDBHeight > 0 {
		// get the previous directory block from db
		dbBlock, _ := ps.DB.FetchDBlockByHeight(dchain.NextDBHeight - 1)
		dbHeaderBytes, _ := dbBlock.Header.MarshalBinary()
		identityChainID := NewZeroHash() // 0 ID for milestone 1
		sig := serverPrivKey.Sign(dbHeaderBytes)
		ps.AChain.NextBlock.AddABEntry(NewDBSignatureEntry(identityChainID, sig))
	}
	return nil
}
