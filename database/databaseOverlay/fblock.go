package databaseOverlay

import (
	"github.com/FactomProject/factomd/common/factoid"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
	"github.com/FactomProject/factomd/util"
	"sort"
	"time"
)

func (db *Overlay) ProcessFBlockBatch(block interfaces.DatabaseBlockWithEntries) error {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddatabaseOverlayOverlayProcessFBlockBatch.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	err := db.ProcessBlockBatch(FACTOIDBLOCK, FACTOIDBLOCK_NUMBER, FACTOIDBLOCK_SECONDARYINDEX, block)
	if err != nil {
		return err
	}
	return db.SaveIncludedInMultiFromBlock(block, false)
}

func (db *Overlay) ProcessFBlockBatchWithoutHead(block interfaces.DatabaseBlockWithEntries) error {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddatabaseOverlayOverlayProcessFBlockBatchWithoutHead.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	err := db.ProcessBlockBatchWithoutHead(FACTOIDBLOCK, FACTOIDBLOCK_NUMBER, FACTOIDBLOCK_SECONDARYINDEX, block)
	if err != nil {
		return err
	}
	return db.SaveIncludedInMultiFromBlock(block, false)
}

func (db *Overlay) ProcessFBlockMultiBatch(block interfaces.DatabaseBlockWithEntries) error {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddatabaseOverlayOverlayProcessFBlockMultiBatch.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	err := db.ProcessBlockMultiBatch(FACTOIDBLOCK, FACTOIDBLOCK_NUMBER, FACTOIDBLOCK_SECONDARYINDEX, block)
	if err != nil {
		return err
	}
	return db.SaveIncludedInMultiFromBlockMultiBatch(block, true)
}

func (db *Overlay) FetchFBlock(hash interfaces.IHash) (interfaces.IFBlock, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddatabaseOverlayOverlayFetchFBlock.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	block, err := db.FetchFBlockByPrimary(hash)
	if err != nil {
		return nil, err
	}
	if block != nil {
		return block, nil
	}
	return db.FetchFBlockBySecondary(hash)
}

func (db *Overlay) FetchFBlockBySecondary(hash interfaces.IHash) (interfaces.IFBlock, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddatabaseOverlayOverlayFetchFBlockBySecondary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	block, err := db.FetchBlockBySecondaryIndex(FACTOIDBLOCK_SECONDARYINDEX, FACTOIDBLOCK, hash, new(factoid.FBlock))
	if err != nil {
		return nil, err
	}
	if block == nil {
		return nil, nil
	}
	return block.(interfaces.IFBlock), nil
}

func (db *Overlay) FetchFBlockByPrimary(keyMR interfaces.IHash) (interfaces.IFBlock, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddatabaseOverlayOverlayFetchFBlockByPrimary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	block, err := db.FetchBlock(FACTOIDBLOCK, keyMR, new(factoid.FBlock))
	if err != nil {
		return nil, err
	}
	if block == nil {
		return nil, nil
	}
	return block.(interfaces.IFBlock), nil
}

// FetchFBlockByHeight gets a factoid block by height from the database.
func (db *Overlay) FetchFBlockByHeight(blockHeight uint32) (interfaces.IFBlock, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddatabaseOverlayOverlayFetchFBlockByHeight.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	block, err := db.FetchBlockByHeight(FACTOIDBLOCK_NUMBER, FACTOIDBLOCK, blockHeight, new(factoid.FBlock))
	if err != nil {
		return nil, err
	}
	if block == nil {
		return nil, nil
	}
	return block.(interfaces.IFBlock), nil
}

func (db *Overlay) FetchAllFBlocks() ([]interfaces.IFBlock, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddatabaseOverlayOverlayFetchAllFBlocks.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	list, err := db.FetchAllBlocksFromBucket(FACTOIDBLOCK, new(factoid.FBlock))
	if err != nil {
		return nil, err
	}
	return toFactoidList(list), nil
}

func (db *Overlay) FetchAllFBlockKeys() ([]interfaces.IHash, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddatabaseOverlayOverlayFetchAllFBlockKeys.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return db.FetchAllBlockKeysFromBucket(FACTOIDBLOCK)
}

func toFactoidList(source []interfaces.BinaryMarshallableAndCopyable) []interfaces.IFBlock {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddatabaseOverlaytoFactoidList.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	answer := make([]interfaces.IFBlock, len(source))
	for i, v := range source {
		answer[i] = v.(interfaces.IFBlock)
	}
	sort.Sort(util.ByFBlockIDAccending(answer))
	return answer
}

func (db *Overlay) SaveFactoidBlockHead(fblock interfaces.DatabaseBlockWithEntries) error {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddatabaseOverlayOverlaySaveFactoidBlockHead.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return db.ProcessFBlockBatch(fblock)
}

func (db *Overlay) FetchFBlockHead() (interfaces.IFBlock, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddatabaseOverlayOverlayFetchFBlockHead.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return db.FetchFactoidBlockHead()
}

func (db *Overlay) FetchFactoidBlockHead() (interfaces.IFBlock, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddatabaseOverlayOverlayFetchFactoidBlockHead.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	blk := new(factoid.FBlock)
	block, err := db.FetchChainHeadByChainID(FACTOIDBLOCK, primitives.NewHash(blk.GetChainID().Bytes()), blk)
	if err != nil {
		return nil, err
	}
	if block == nil {
		return nil, nil
	}
	return block.(interfaces.IFBlock), nil
}
