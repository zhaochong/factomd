package databaseOverlay

import (
	"github.com/FactomProject/factomd/common/entryCreditBlock"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
	"github.com/FactomProject/factomd/util"
	"sort"
	"time"
)

// ProcessECBlockBatch inserts the ECBlock and update all it's cbentries in DB
func (db *Overlay) ProcessECBlockBatch(block interfaces.IEntryCreditBlock, checkForDuplicateEntries bool) error {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddatabaseOverlayOverlayProcessECBlockBatch.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	err := db.ProcessBlockBatch(ENTRYCREDITBLOCK,
		ENTRYCREDITBLOCK_NUMBER,
		ENTRYCREDITBLOCK_SECONDARYINDEX, block)
	if err != nil {
		return err
	}
	err = db.SaveIncludedInMultiFromBlock(block, false)
	if err != nil {
		return err
	}
	return db.SavePaidForMultiFromBlock(block, checkForDuplicateEntries)
}

func (db *Overlay) ProcessECBlockBatchWithoutHead(block interfaces.IEntryCreditBlock, checkForDuplicateEntries bool) error {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddatabaseOverlayOverlayProcessECBlockBatchWithoutHead.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	err := db.ProcessBlockBatchWithoutHead(ENTRYCREDITBLOCK,
		ENTRYCREDITBLOCK_NUMBER,
		ENTRYCREDITBLOCK_SECONDARYINDEX, block)
	if err != nil {
		return err
	}
	err = db.SaveIncludedInMultiFromBlock(block, false)
	if err != nil {
		return err
	}
	return db.SavePaidForMultiFromBlock(block, checkForDuplicateEntries)
}

func (db *Overlay) ProcessECBlockMultiBatch(block interfaces.IEntryCreditBlock, checkForDuplicateEntries bool) error {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddatabaseOverlayOverlayProcessECBlockMultiBatch.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	err := db.ProcessBlockMultiBatch(ENTRYCREDITBLOCK,
		ENTRYCREDITBLOCK_NUMBER,
		ENTRYCREDITBLOCK_SECONDARYINDEX, block)
	if err != nil {
		return err
	}
	err = db.SaveIncludedInMultiFromBlockMultiBatch(block, true)
	if err != nil {
		return err
	}
	return db.SavePaidForMultiFromBlockMultiBatch(block, checkForDuplicateEntries)
}

func (db *Overlay) FetchECBlock(hash interfaces.IHash) (interfaces.IEntryCreditBlock, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddatabaseOverlayOverlayFetchECBlock.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	block, err := db.FetchECBlockByPrimary(hash)
	if err != nil {
		return nil, err
	}
	if block != nil {
		return block, nil
	}
	return db.FetchECBlockBySecondary(hash)
}

// FetchECBlockByHeaderHash gets an Entry Credit block by hash from the database.
func (db *Overlay) FetchECBlockBySecondary(hash interfaces.IHash) (interfaces.IEntryCreditBlock, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddatabaseOverlayOverlayFetchECBlockBySecondary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	block, err := db.FetchBlockBySecondaryIndex(ENTRYCREDITBLOCK_SECONDARYINDEX, ENTRYCREDITBLOCK, hash, entryCreditBlock.NewECBlock())
	if err != nil {
		return nil, err
	}
	if block == nil {
		return nil, nil
	}
	return block.(interfaces.IEntryCreditBlock), nil
}

// FetchECBlockByHash gets an Entry Credit block by hash from the database.
func (db *Overlay) FetchECBlockByPrimary(hash interfaces.IHash) (interfaces.IEntryCreditBlock, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddatabaseOverlayOverlayFetchECBlockByPrimary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	block, err := db.FetchBlock(ENTRYCREDITBLOCK, hash, entryCreditBlock.NewECBlock())
	if err != nil {
		return nil, err
	}
	if block == nil {
		return nil, nil
	}
	return block.(interfaces.IEntryCreditBlock), nil
}

// FetchECBlockByHeight gets an entry credit block by height from the database.
func (db *Overlay) FetchECBlockByHeight(blockHeight uint32) (interfaces.IEntryCreditBlock, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddatabaseOverlayOverlayFetchECBlockByHeight.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	block, err := db.FetchBlockByHeight(ENTRYCREDITBLOCK_NUMBER, ENTRYCREDITBLOCK, blockHeight, entryCreditBlock.NewECBlock())
	if err != nil {
		return nil, err
	}
	if block == nil {
		return nil, nil
	}
	return block.(interfaces.IEntryCreditBlock), nil
}

// FetchAllECBlocks gets all of the entry credit blocks
func (db *Overlay) FetchAllECBlocks() ([]interfaces.IEntryCreditBlock, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddatabaseOverlayOverlayFetchAllECBlocks.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	list, err := db.FetchAllBlocksFromBucket(ENTRYCREDITBLOCK, entryCreditBlock.NewECBlock())
	if err != nil {
		return nil, err
	}
	return toECBlocksList(list), nil
}

func (db *Overlay) FetchAllECBlockKeys() ([]interfaces.IHash, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddatabaseOverlayOverlayFetchAllECBlockKeys.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return db.FetchAllBlockKeysFromBucket(ENTRYCREDITBLOCK)
}

func toECBlocksList(source []interfaces.BinaryMarshallableAndCopyable) []interfaces.IEntryCreditBlock {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddatabaseOverlaytoECBlocksList.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	answer := make([]interfaces.IEntryCreditBlock, len(source))
	for i, v := range source {
		answer[i] = v.(interfaces.IEntryCreditBlock)
	}
	sort.Sort(util.ByECBlockIDAccending(answer))
	return answer
}

func (db *Overlay) SaveECBlockHead(block interfaces.IEntryCreditBlock, checkForDuplicateEntries bool) error {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddatabaseOverlayOverlaySaveECBlockHead.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return db.ProcessECBlockBatch(block, checkForDuplicateEntries)
}

func (db *Overlay) FetchECBlockHead() (interfaces.IEntryCreditBlock, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddatabaseOverlayOverlayFetchECBlockHead.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	blk := entryCreditBlock.NewECBlock()
	block, err := db.FetchChainHeadByChainID(ENTRYCREDITBLOCK, primitives.NewHash(blk.GetChainID().Bytes()), blk)
	if err != nil {
		return nil, err
	}
	if block == nil {
		return nil, nil
	}
	return block.(interfaces.IEntryCreditBlock), nil
}
