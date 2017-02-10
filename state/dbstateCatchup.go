package state

import "github.com/FactomProject/factomd/common/messages"

// Once a second at most, we check to see if we need to pull down some blocks to catch up.
func (list *DBStateList) Catchup(justDoIt bool) {
	// We only check if we need updates once every so often.

	if len(list.State.inMsgQueue) > 1000 {
		// If we are behind the curve in processing messages, dump all the dbstates from holding.
		for k := range list.State.Holding {
			if _, ok := list.State.Holding[k].(*messages.DBStateMsg); ok {
				delete(list.State.Holding, k)
			}
		}
		return
	}

	now := list.State.GetTimestamp()

	hs := int(list.State.GetHighestSavedBlk())
	hk := int(list.State.GetHighestKnownBlock())
	begin := hs + 1
	end := hk

	ask := func() {
		if list.TimeToAsk != nil && hk-hs > 4 && now.GetTime().After(list.TimeToAsk.GetTime()) {

			for i, v := range list.State.DBStatesReceived {
				ix := i + list.State.DBStatesReceivedBase
				if ix <= hs {
					continue
				}
				if ix >= hk {
					return
				}
				if v != nil {
					begin = ix
				} else {
					continue
				}
			}

			// Don't ask for more than we already have.
			for i, v := range list.State.DBStatesReceived {
				ix := i + list.State.DBStatesReceivedBase
				if ix <= begin {
					continue
				}
				if ix >= end {
					break
				}
				if v != nil {
					end = ix + 1
					if begin > end {
						return
					}
					break
				}
			}

			msg := messages.NewDBStateMissing(list.State, uint32(begin), uint32(end+3))

			if msg != nil {
				//		list.State.RunLeader = false
				//		list.State.StartDelay = list.State.GetTimestamp().GetTimeMilli()
				msg.SendOut(list.State, msg)
				list.State.DBStateAskCnt++
				list.TimeToAsk.SetTimeSeconds(now.GetTimeSeconds() + 3)
				list.LastBegin = begin
				list.LastEnd = end + 3
			}
		}
	}

	if end-begin > 200 {
		end = begin + 200
	}

	if end+3 > begin && justDoIt {
		ask()
		return
	}

	// return if we are caught up, and clear our timer
	if end-begin <= 1 {
		list.TimeToAsk = nil
		return
	}

	// First Ask.  Because the timer is nil!
	if list.TimeToAsk == nil {
		// Okay, have nothing in play, so wait a bit just in case.
		list.TimeToAsk = list.State.GetTimestamp()
		list.TimeToAsk.SetTimeSeconds(now.GetTimeSeconds() + 5)
		list.LastBegin = begin
		list.LastEnd = end
		return
	}

	if list.TimeToAsk.GetTime().Before(now.GetTime()) {
		ask()
		return
	}

}
