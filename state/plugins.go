package state

import (
	"fmt"
	"time"
	// "sync"

	//"github.com/FactomProject/factomd/common/entryBlock"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/messages"
)

type heightError struct {
	Err    error
	Height uint32
}

// Controls the flow of uploading torrents
type UploadController struct {
	// DO NOT USE THIS MAP OUTSIDE sortRequests()
	// It is not concurrency safe
	uploaded           map[uint32]struct{} // Map of uploaded heights
	requestUploadQueue chan uint32
	sendUploadQueue    chan uint32      // heights to be uploaded
	failedQueue        chan heightError // Channel of heights that failed to upload

	DBStateManager interfaces.IManagerController

	quit chan int
}

func NewUploadController(dbsm interfaces.IManagerController) *UploadController {
	u := new(UploadController)
	u.requestUploadQueue = make(chan uint32, 100000) // Channel used if torrents enabled. Queue of torrents to upload
	u.sendUploadQueue = make(chan uint32, 100000)
	u.failedQueue = make(chan heightError, 1000)

	u.uploaded = make(map[uint32]struct{})

	u.quit = make(chan int, 10)
	u.DBStateManager = dbsm

	return u
}

func (s *State) RunUploadController() {
	fmt.Println("Starting upload controller")
	go s.Uploader.sortRequests()
	go s.uploadBlocks()
	go s.Uploader.handleErrors()
	go s.Uploader.Status()
}

func (u *UploadController) Status() {
	for {
		select {
		case <-u.quit:
			u.quit <- 0
			return
		default:

			time.Sleep(2 * time.Second)
			fmt.Printf("[Uploader] Request: %d, Send: %d, failed: %d\n", len(u.requestUploadQueue), len(u.sendUploadQueue), len(u.failedQueue))
		}
	}
}

func (u *UploadController) Close() {
	u.quit <- 0
}

// sortRequests sorts through the inital requests to toss out repeats
func (u *UploadController) sortRequests() {
	for {
	backToTopSortRequests:
		select {
		// Avoid defering the lock, more overhead
		case h := <-u.requestUploadQueue:
			if _, ok := u.uploaded[h]; ok {
				// Already uploaded, toss out
				goto backToTopSortRequests
			}

			u.uploaded[h] = struct{}{}
			u.sendUploadQueue <- h
		case <-u.quit:
			u.quit <- 0
			return
		}
	}
}

func (s *State) uploadBlocks() {
	u := s.Uploader
	for {
	backToTopUploadBlocks:
		select {
		case <-u.quit:
			u.quit <- 0
			return
		default:
			readyFor := u.DBStateManager.RequestMoreUploads()
			// Need to check if we are able to upload any blocks. If we cannot, we will wait
			if readyFor == 0 { // Not ready for anything
				time.Sleep(1 * time.Second)
				goto backToTopUploadBlocks
			} else if readyFor < 0 {
				// This is a crash....
				return
			} else {
				// We can make some uploads. Only loop readyFor times
				for i := 0; i < readyFor; i++ {
					select { // We will block if nothing is in queue and chill here
					case h := <-u.sendUploadQueue:
						err := s.uploadDBState(h)
						if err != nil {
							u.failedQueue <- heightError{Height: h, Err: err}
						}
					case <-u.quit:
						u.quit <- 0
						return
					}
				}
			}
		}
	}
}

func (u *UploadController) handleErrors() {
	for {
		select {
		case <-u.quit:
			u.quit <- 0
			return
		case <-u.failedQueue:
			// TODO: Handle errors
		}
	}
}

/*****************
	State Calls
******************/

// Only called once to set the torrent flag.
func (s *State) SetUseTorrent(setVal bool) {
	s.useDBStateManager = setVal
}

func (s *State) UsingTorrent() bool {
	return s.useDBStateManager
}

/*
	msg, err := list.State.LoadDBState(uint32(dbheight))
		if err != nil {
			fmt.Println("[1] Error creating torrent in SaveDBStateToDB: " + err.Error())
		}
*/

// All calls get sent here and redirected into the uploadcontroller queue.
func (s *State) UploadDBState(dbheight uint32) {
	s.Uploader.requestUploadQueue <- dbheight
}

func (s *State) uploadDBState(height uint32) error {
	// Create the torrent
	if s.UsingTorrent() {
		msg, err := s.LoadDBState(height)
		if err != nil {
			return err
		}
		d := msg.(*messages.DBStateMsg)
		//fmt.Printf("Uploading DBState %d, Sigs: %d\n", d.DirectoryBlock.GetDatabaseHeight(), len(d.SignatureList.List))
		block := NewWholeBlock()
		block.DBlock = d.DirectoryBlock
		block.ABlock = d.AdminBlock
		block.FBlock = d.FactoidBlock
		block.ECBlock = d.EntryCreditBlock

		eHashes := make([]interfaces.IHash, 0)
		for _, e := range d.EBlocks {
			block.AddEblock(e)
			for _, eh := range e.GetEntryHashes() {
				eHashes = append(eHashes, eh)
			}
		}

		if len(eHashes) == 0 {
			// No hashes in the msg. Possibly not make torrent?
			// If we only use torrents for entry syncing, then no need
			// to make this torrent
		}

		for _, e := range eHashes {
			if e.String()[:62] != "00000000000000000000000000000000000000000000000000000000000000" {
				//} else {
				ent, err := s.DB.FetchEntry(e)
				if err != nil {
					return fmt.Errorf("[2] Error creating torrent in SaveDBStateToDB: " + err.Error())
				}
				block.AddIEBEntry(ent)
			}
		}

		if len(d.SignatureList.List) == 0 {
			return fmt.Errorf("No signatures given, signatures must be in to be able to torrent")
		}
		block.SigList = d.SignatureList.List

		data, err := block.MarshalBinary()
		if err != nil {
			return fmt.Errorf("[3] Error creating torrent in SaveDBStateToDB: " + err.Error())

		}

		if s.IsLeader() {
			s.DBStateManager.UploadDBStateBytes(data, true)
		} else {
			s.DBStateManager.UploadDBStateBytes(data, false)
		}
	}
	return nil
}

func (s *State) GetMissingDBState(height uint32) error {
	return s.DBStateManager.RetrieveDBStateByHeight(height)
}
