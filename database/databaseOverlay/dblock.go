// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Copyright (c) 2013-2014 Conformal Systems LLC.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package databaseOverlay

import (
	"sort"

	"github.com/FactomProject/factomd/common/directoryBlock"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
	"github.com/FactomProject/factomd/util"
	"time"
)

// ProcessDBlockBatche inserts the DBlock and update all it's dbentries in DB
func (db *Overlay) ProcessDBlockBatch(dblock interfaces.DatabaseBlockWithEntries) error {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddatabaseOverlayOverlayProcessDBlockBatch.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	err := db.ProcessBlockBatch(DIRECTORYBLOCK,
		DIRECTORYBLOCK_NUMBER,
		DIRECTORYBLOCK_SECONDARYINDEX, dblock)
	if err != nil {
		return err
	}

	return db.SaveIncludedInMultiFromBlock(dblock, false)
}

func (db *Overlay) ProcessDBlockBatchWithoutHead(dblock interfaces.DatabaseBlockWithEntries) error {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddatabaseOverlayOverlayProcessDBlockBatchWithoutHead.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	err := db.ProcessBlockBatchWithoutHead(DIRECTORYBLOCK,
		DIRECTORYBLOCK_NUMBER,
		DIRECTORYBLOCK_SECONDARYINDEX, dblock)
	if err != nil {
		return err
	}

	return db.SaveIncludedInMultiFromBlock(dblock, false)
}

func (db *Overlay) ProcessDBlockMultiBatch(dblock interfaces.DatabaseBlockWithEntries) error {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddatabaseOverlayOverlayProcessDBlockMultiBatch.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	err := db.ProcessBlockMultiBatch(DIRECTORYBLOCK,
		DIRECTORYBLOCK_NUMBER,
		DIRECTORYBLOCK_SECONDARYINDEX, dblock)
	if err != nil {
		return err
	}

	return db.SaveIncludedInMultiFromBlockMultiBatch(dblock, true)
}

// FetchHeightRange looks up a range of blocks by the start and ending
// heights.  Fetch is inclusive of the start height and exclusive of the
// ending height. To fetch all hashes from the start height until no
// more are present, use -1 as endHeight.
func (db *Overlay) FetchDBlockHeightRange(startHeight, endHeight int64) ([]interfaces.IHash, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddatabaseOverlayOverlayFetchDBlockHeightRange.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return db.FetchBlockIndexesInHeightRange(DIRECTORYBLOCK_NUMBER, startHeight, endHeight)
}

// FetchBlockHeightByKeyMR returns the block height for the given hash.  This is
// part of the database.Db interface implementation.
func (db *Overlay) FetchDBlockHeightByKeyMR(sha interfaces.IHash) (int64, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddatabaseOverlayOverlayFetchDBlockHeightByKeyMR.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	dblk, err := db.FetchDBlock(sha)
	if err != nil {
		return -1, err
	}

	var height int64 = -1
	if dblk != nil {
		height = int64(dblk.GetDatabaseHeight())
	}

	return height, nil
}

func (db *Overlay) FetchDBlock(hash interfaces.IHash) (interfaces.IDirectoryBlock, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddatabaseOverlayOverlayFetchDBlock.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	block, err := db.FetchDBlockByPrimary(hash)
	if err != nil {
		return nil, err
	}
	if block != nil {
		return block, nil
	}
	return db.FetchDBlockBySecondary(hash)
}

// FetchDBlock gets an entry by hash from the database.
func (db *Overlay) FetchDBlockByPrimary(keyMR interfaces.IHash) (interfaces.IDirectoryBlock, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddatabaseOverlayOverlayFetchDBlockByPrimary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	block, err := db.FetchBlock(DIRECTORYBLOCK, keyMR, new(directoryBlock.DirectoryBlock))
	if err != nil {
		return nil, err
	}
	if block == nil {
		return nil, nil
	}
	return block.(interfaces.IDirectoryBlock), nil
}

// FetchDBlockByMR gets a directory block by merkle root from the database.
func (db *Overlay) FetchDBlockBySecondary(dBMR interfaces.IHash) (interfaces.IDirectoryBlock, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddatabaseOverlayOverlayFetchDBlockBySecondary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	block, err := db.FetchBlockBySecondaryIndex(DIRECTORYBLOCK_SECONDARYINDEX, DIRECTORYBLOCK, dBMR, new(directoryBlock.DirectoryBlock))
	if err != nil {
		return nil, err
	}
	if block == nil {
		return nil, nil
	}
	return block.(interfaces.IDirectoryBlock), nil
}

// FetchDBlockByHeight gets an directory block by height from the database.
func (db *Overlay) FetchDBlockByHeight(dBlockHeight uint32) (interfaces.IDirectoryBlock, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddatabaseOverlayOverlayFetchDBlockByHeight.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	block, err := db.FetchBlockByHeight(DIRECTORYBLOCK_NUMBER, DIRECTORYBLOCK, dBlockHeight, new(directoryBlock.DirectoryBlock))
	if err != nil {
		return nil, err
	}
	if block == nil {
		return nil, nil
	}
	return block.(interfaces.IDirectoryBlock), nil
}

// FetchDBKeyMRByHeight gets a dBlock KeyMR from the database.
func (db *Overlay) FetchDBKeyMRByHeight(dBlockHeight uint32) (interfaces.IHash, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddatabaseOverlayOverlayFetchDBKeyMRByHeight.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return db.FetchBlockIndexByHeight(DIRECTORYBLOCK_NUMBER, dBlockHeight)
}

// FetchDBKeyMRByHash gets a DBlock KeyMR by hash.
func (db *Overlay) FetchDBKeyMRByHash(hash interfaces.IHash) (interfaces.IHash, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddatabaseOverlayOverlayFetchDBKeyMRByHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return db.FetchPrimaryIndexBySecondaryIndex(DIRECTORYBLOCK_SECONDARYINDEX, hash)
}

// FetchAllDBlocks gets all of the fbInfo
func (db *Overlay) FetchAllDBlocks() ([]interfaces.IDirectoryBlock, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddatabaseOverlayOverlayFetchAllDBlocks.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	list, err := db.FetchAllBlocksFromBucket(DIRECTORYBLOCK, new(directoryBlock.DirectoryBlock))
	if err != nil {
		return nil, err
	}
	return toDBlocksList(list), nil
}

func (db *Overlay) FetchAllDBlockKeys() ([]interfaces.IHash, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddatabaseOverlayOverlayFetchAllDBlockKeys.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return db.FetchAllBlockKeysFromBucket(DIRECTORYBLOCK)
}

func toDBlocksList(source []interfaces.BinaryMarshallableAndCopyable) []interfaces.IDirectoryBlock {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddatabaseOverlaytoDBlocksList.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	answer := make([]interfaces.IDirectoryBlock, len(source))
	for i, v := range source {
		answer[i] = v.(interfaces.IDirectoryBlock)
	}
	sort.Sort(util.ByDBlockIDAccending(answer))
	return answer
}

func (db *Overlay) SaveDirectoryBlockHead(dblock interfaces.DatabaseBlockWithEntries) error {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddatabaseOverlayOverlaySaveDirectoryBlockHead.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return db.ProcessDBlockBatch(dblock)
}

func (db *Overlay) FetchDBlockHead() (interfaces.IDirectoryBlock, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddatabaseOverlayOverlayFetchDBlockHead.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return db.FetchDirectoryBlockHead()
}

func (db *Overlay) FetchDirectoryBlockHead() (interfaces.IDirectoryBlock, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddatabaseOverlayOverlayFetchDirectoryBlockHead.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	blk := new(directoryBlock.DirectoryBlock)
	block, err := db.FetchChainHeadByChainID(DIRECTORYBLOCK, primitives.NewHash(blk.GetChainID().Bytes()), blk)
	if err != nil {
		return nil, err
	}
	if block == nil {
		return nil, nil
	}
	return block.(interfaces.IDirectoryBlock), nil
}
