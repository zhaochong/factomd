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
		// We can adjust the sleep at the end depending on what we do in the loop
		// this pass
		rightDuration := time.Duration(time.Second * 1)

		// How many blocks ahead of the current we should request from the plugin
		allowed := 5000

		// What is the database at
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

		// If the network is at block 0, we aren't on the network
		if upper == 0 {
			time.Sleep(5 * time.Second)
			continue
		}

		// Synced up, sleep for awhile
		if lower == upper {
			rightDuration = time.Duration(20 * time.Second)
		}

		// Prometheus
		stateTorrentSyncingLower.Set(float64(lower))
		stateTorrentSyncingUpper.Set(float64(upper))

		// What is the end height we request
		max := lower + uint32(allowed)
		if upper < max {
			rightDuration = time.Duration(5 * time.Second)
			max = upper
		}

		var u uint32 = 0
		// The torrent plugin handles dealing with lots of heights. It has it's own queueing system, so
		// we can spam and repeat heights
		for u = lower; u < max; u++ {
			// Plugin handles repeat requests
			err := s.DBStateManager.RetrieveDBStateByHeight(u)
			if err != nil {
				if s.DBStateManager.Alive() == nil { // Some errors come from a plugin crash (like when you ctrl+c)
					log.Printf("[TorrentSync] Error while retrieving height %d by torrent, %s", u, err.Error())
				} else {
					// Connection to plugin lost, exit as it won't return
					return fmt.Errorf("Torrent plugin stopped")
				}
			}
		}

		// This tells our plugin to ignore any heights below this for retrieval
		s.DBStateManager.CompletedHeightTo(s.EntryDBHeightComplete)

		// Request the 2nd pass too, plugin will handle any repeats, and this makes our second pass catch up
		s.DBStateManager.RetrieveDBStateByHeight(s.EntryDBHeightComplete + 1)

		time.Sleep(rightDuration)
	}
}
