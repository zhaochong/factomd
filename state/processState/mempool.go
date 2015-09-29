// Copyright 2015 FactomProject Authors. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package processState

import (
	"errors"
	"github.com/FactomProject/factomd/btcd/wire"
	. "github.com/FactomProject/factomd/common/constants"
	. "github.com/FactomProject/factomd/common/interfaces"
	"sync"
	"time"
)

// FtmMemPool is used as a source of factom transactions
// (CommitChain, RevealChain, CommitEntry, RevealEntry)
type FtmMemPool struct {
	sync.RWMutex
	pool        map[IHash]wire.Message
	orphans     map[IHash]wire.Message
	blockpool   map[string]wire.Message // to hold the blocks or entries downloaded from peers
	lastUpdated time.Time               // last time pool was updated
}

// Add a factom message to the orphan pool
func (mp *FtmMemPool) Init_FtmMemPool() error {

	mp.pool = make(map[IHash]wire.Message)
	mp.orphans = make(map[IHash]wire.Message)
	mp.blockpool = make(map[string]wire.Message)

	return nil
}

// Add a factom message to the  Mem pool
func (mp *FtmMemPool) AddMsg(msg wire.Message, hash IHash) error {

	if len(mp.pool) > MAX_TX_POOL_SIZE {
		return errors.New("Transaction mem pool exceeds the limit.")
	}

	mp.pool[hash] = msg

	return nil
}

// Add a factom message to the orphan pool
func (mp *FtmMemPool) AddOrphanMsg(msg wire.Message, hash IHash) error {

	if len(mp.orphans) > MAX_ORPHAN_SIZE {
		errors.New("Ophan mem pool exceeds the limit.")
	}

	mp.orphans[hash] = msg

	return nil
}

// Add a factom block message to the  Mem pool
func (mp *FtmMemPool) AddBlockMsg(msg wire.Message, hash string) error {

	if len(mp.blockpool) > MAX_BLK_POOL_SIZE {
		errors.New("Block mem pool exceeds the limit. Please restart.")
	}
	mp.Lock()
	mp.blockpool[hash] = msg
	mp.Unlock()

	return nil
}

// Delete a factom block message from the  Mem pool
func (mp *FtmMemPool) DeleteBlockMsg(hash string) error {

	if mp.blockpool[hash] != nil {
		mp.Lock()
		delete(mp.blockpool, hash)
		mp.Unlock()
	}

	return nil
}
