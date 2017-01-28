package state

import (
	"fmt"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/messages"
	"time"
)

var _ = fmt.Print

// Create a request for dbsates from begin to end
//
func Ask(state *State, begin uint32, end uint32) {
	if begin > 0 {
		begin--
	}
	msg := messages.NewDBStateMissing(state, uint32(begin), uint32(end))
	if msg != nil {
		msg.SendOut(state, msg)
		state.DBStateAskCnt++
	}
}

// Keep track of request history per dbstate
type AskHistory struct {
	LastRequest interfaces.Timestamp
	DBHeight    uint32
}

// All the dbstates we have asked for to date.
type AskHistories struct {
	Base    int
	History []*AskHistory
}

// Trim old histories.  We don't need them any more.
func (h *AskHistories) Trim(HighestSaved uint32, HighestKnown uint32) {

	for len(h.History) < int(HighestSaved)-int(h.Base) {
		h.History = append(h.History, nil)
	}
	if h.Base < int(HighestSaved) {
		cnt := int(HighestSaved) - h.Base
		if len(h.History) >= cnt {
			h.History = append(make([]*AskHistory,0), h.History[cnt:]...)
			h.Base = int(HighestSaved)
		}
	}
}

func (h *AskHistories) Get(DBHeight int) *AskHistory {

	index := DBHeight - h.Base

	if index < 0 {
		return nil
	}

	for index >= len(h.History) {
		h.History = append(h.History, nil)
	}

	if h.History[index] == nil {
		h.History[index] = new(AskHistory)
		h.History[index].DBHeight = uint32(DBHeight)
	}
	return h.History[index]
}

func FindBegin(state *State, histories *AskHistories, start int) int {

	// No point asking for dbstates we have saved to disk
	if start < int(state.HighestDisk) {
		start = int(state.HighestDisk)
	}

	begin := start
	now := state.GetTimestamp()
	for {
		h := histories.Get(begin)
		ir := begin - state.DBStatesReceivedBase
		switch {
		case begin > int(state.HighestKnown) :
			return -1
		case h == nil:
		case len(state.DBStatesReceived) > int(ir) && state.DBStatesReceived[ir] != nil:
		case h.LastRequest == nil:
			h.LastRequest = now
			return begin
		case now.GetTimeSeconds()-h.LastRequest.GetTimeSeconds() >= 3:
			h.LastRequest = now
			return begin
		}
		begin++
	}
}

func FindEnd(state *State, histories *AskHistories, begin int) int {

	// If no valid begin, there is no valid end
	if begin < 0 {
		return -1
	}

	now := state.GetTimestamp()
	end := begin+1
	for {
		h := histories.Get(end)
		ir := int(end) - state.DBStatesReceivedBase

		switch {
		case len(state.DBStatesReceived) > int(ir) && state.DBStatesReceived[ir] != nil:
			return end-1
		case end >= int(state.HighestKnown):
			return end
		case end-begin >= 200:
			return end
		case h.LastRequest == nil:
		case now.GetTimeSeconds()-h.LastRequest.GetTimeSeconds() < 3:
			return end-1
		}
		h.LastRequest = now
		end++
	}
	return end
}

// Does one check, and possible submission of a request.
func Step(state *State, histories *AskHistories) {


	// While ignoring, don't ask for dbstates
	if state.IgnoreMissing {
		return
	}

	// We only check if we need updates once every so often.
	if len(state.inMsgQueue) > 1000 {
		return
	}

	hs := state.HighestSaved
	hk := state.HighestKnown

	if state.CurrentMinute > 0 && hk-hs <= 1 {
		return
	} else if state.CurrentMinute == 0 && hk-hs <= 2 {
		return
	}

	//histories.Trim(hs, hk)

	begin := FindBegin(state, histories, int(hs+1))
	if begin <= 0 || uint32(begin) > hs+1000 {
		return
	}

	end := FindEnd(state, histories, begin)
	Ask(state, uint32(begin), uint32(end))

}

// Once a second at most, we check to see if we need to pull down some blocks to catch up.
func (list *DBStateList) Catchup() {

	histories := new(AskHistories)

	// Just let the system come up before we try to sync.
	time.Sleep(5 * time.Second)

	state := list.State
	state.DBStateMutex.Lock()

	for {
		// So we yeild so we don't lock up the system.
		state.DBStateMutex.Unlock()
		time.Sleep(500 * time.Millisecond)
		state.DBStateMutex.Lock()

		Step(list.State, histories)
	}
}
