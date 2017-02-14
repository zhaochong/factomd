// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package state

import (
	"fmt"
	"github.com/FactomProject/factomd/common/constants"
	"time"
)

func (s *State) syncEntryBlocks() {
	// All done is true, and as long as it says true, we walk our bookmark forward.  Once we find something
	// missing, we stop moving the bookmark, and rely on caching to keep us from thrashing the disk as we
	// review the directory block over again the next time.

	for i := 0; i < 20 && s.EntryBlockDBHeightProcessing <= s.GetHighestSavedBlk() && len(s.MissingEntryBlocks) < 1000; i++ {
		db := s.GetDirectoryBlockByHeight(s.EntryBlockDBHeightProcessing)

		if db == nil {
			return
		}

		alldone := true
		for _, ebKeyMR := range db.GetEntryHashes()[3:] {
			// The first three entries (0,1,2) in every directory block are blocks we already have by
			// definition.  If we decide to not have Factoid blocks or Entry Credit blocks in some cases,
			// then this assumption might not hold.  But it does for now.

			eBlock, _ := s.DB.FetchEBlock(ebKeyMR)

			// Ask for blocks we don't have.
			if eBlock == nil {
				alldone = false
				// Check lists and not add if already there.
				addit := true
				for _, eb := range s.MissingEntryBlocks {
					if eb.ebhash.Fixed() == ebKeyMR.Fixed() {
						addit = false
						break
					}
				}

				if addit {
					s.MissingEntryBlocks = append(s.MissingEntryBlocks,
						MissingEntryBlock{ebhash: ebKeyMR, dbheight: s.EntryBlockDBHeightProcessing})
				}
				// Something missing, stop moving the bookmark.

				continue
			}
		}
		if !alldone {
			break
		}
	}
}

func (s *State) syncEntries(eights bool) {

	alldone := true

	var keep []MissingEntry
	for _, v := range s.MissingEntries {
		e, err := s.DB.FetchEntry(v.entryhash)
		if err != nil || e == nil {
			keep = append(keep, v)
		}
	}
	s.MissingEntries = keep

	for s.EntryBlockDBHeightProcessing <= s.GetHighestSavedBlk() && len(s.MissingEntries) < 1000 {

		db := s.GetDirectoryBlockByHeight(s.EntryBlockDBHeightProcessing)

		if db == nil {
			return
		}

		for _, ebKeyMR := range db.GetEntryHashes()[3:] {
			// The first three entries (0,1,2) in every directory block are blocks we already have by
			// definition.  If we decide to not have Factoid blocks or Entry Credit blocks in some cases,
			// then this assumption might not hold.  But it does for now.

			eBlock, _ := s.DB.FetchEBlock(ebKeyMR)

			// Dont have an eBlock?  Huh. We can go on, but we can't advance
			if eBlock == nil {
				alldone = false
				continue
			}

			for i, entryhash := range eBlock.GetEntryHashes() {
				if entryhash.IsMinuteMarker() {
					fmt.Printf("SyncSync %4d eb: %x Minute %x\n", i, eBlock.GetHash().Bytes(), entryhash.Bytes()[31:])
					continue
				}
				e, _ := s.DB.FetchEntry(entryhash)
				if e == nil {
					//Check lists and not add if already there.
					addit := true
					for _, e := range s.MissingEntries {
						if e.ebhash.Fixed() == entryhash.Fixed() {
							addit = false
							break
						}
					}

					if addit {
						var v MissingEntry

						v.dbheight = eBlock.GetHeader().GetDBHeight()
						v.entryhash = entryhash
						v.ebhash = ebKeyMR

						s.MissingEntries = append(s.MissingEntries, v)
					}
					// Something missing. stop moving the bookmark.
					alldone = false
				} else {
					if i == 0 {
						fmt.Printf("SyncSync dbht: %8d eb: %x cnt: %d\n",
							s.EntryBlockDBHeightProcessing,
							eBlock.GetHash().Bytes(),
							len(eBlock.GetEntryHashes()))
					}
					fmt.Printf("SyncSync %4d eb: %x e: %x\n", i, eBlock.GetHash().Bytes(), entryhash.Bytes())
				}
				// Save the entry hash, and remove from commits IF this hash is valid in this current timeframe.
				s.Replay.SetHashNow(constants.REVEAL_REPLAY, entryhash.Fixed(), db.GetTimestamp())
				// If the save worked, then remove any commit that might be around.
				if !s.Replay.IsHashUnique(constants.REVEAL_REPLAY, entryhash.Fixed()) {
					delete(s.Commits, entryhash.Fixed())
				}
			}
		}
		if alldone {
			s.EntryBlockDBHeightProcessing++
		} else {
			return
		}
	}
}

// This routine walks through the directory blocks, looking for missing entry blocks and entries.
// Once it finds something missing, it keeps that as a mark, and starts there the next time it is
// called.

func (s *State) catchupEBlocks() {

	for {
		// If we still have blocks that we are asking for, then let's not add to the list.
		s.MissingEntriesMutex.Lock()
		s.syncEntryBlocks()
		s.syncEntries(false)
		s.MissingEntriesMutex.Unlock()
		time.Sleep(10)
	}
}
