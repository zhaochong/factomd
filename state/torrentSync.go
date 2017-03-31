package state

import (
	"fmt"
	"log"
	"time"
)

// StartTorrentSyncing is an endless loop that uses torrents to sync missing blocks
// It will grab any block higher than the highest dblock saved in the database up
// to the highest known block.
func (s *State) StartTorrentSyncing() error {
	if !s.UsingTorrent() {
		return fmt.Errorf("State is not using torrents, yet torrent sync was called")
	}

	for {
		// How many requests we can send to the plugin
		allowed := 2500

		dblock, err := s.DB.FetchDBlockHead()
		if err != nil || dblock == nil {
			if err != nil {
				log.Printf("[TorrentSync] Error while retrieving dblock head, %s", err.Error())
			}
			time.Sleep(5 * time.Second) // To prevent error spam
			continue
		}

		// Range of heights to request
		lower := dblock.GetDatabaseHeight()
		upper := s.GetHighestKnownBlock()

		if upper == 0 {
			time.Sleep(5 * time.Second)
			continue
		}

		// Prometheus
		stateTorrentSyncingLower.Set(float64(lower))
		stateTorrentSyncingUpper.Set(float64(upper))

		max := lower + uint32(allowed)
		if upper < max {
			max = upper
		}
		var u uint32 = 0

		totalNeed := 0
		// The torrent plugin handles dealing with lots of heights. It has it's own queueing system, so
		// we can spam. The only things we have to be concerned about is overloading it's queueing system
		for u = lower; u < max; u++ {
			totalNeed++
			err := s.DBStateManager.RetrieveDBStateByHeight(u)
			if err != nil {
				log.Printf("[TorrentSync] Error while retrieving height %d by torrent, %s", u, err.Error())
			}
		}

		//if upper-lower > 100 {
		//	s.DBStateManager.CompletedHeightTo(lower)
		//} else {
		s.DBStateManager.CompletedHeightTo(s.EntryDBHeightComplete)
		//}

		time.Sleep(5 * time.Second)
	}
}
