package state_test

import (
	"testing"

	. "github.com/FactomProject/factomd/state"
	"fmt"
	"time"
)

func TestSyncDBStates1(t *testing.T) {
	state			 := new(State)
	histories  := new(AskHistories)
	state.HighestSaved=100
	state.HighestKnown=800
	state.HighestDisk=300
	for i:=0; i<400; i++ {
		time.Sleep(500 * time.Millisecond)
		begin := FindBegin(state,histories,i)
		end := FindEnd(state,histories,begin)
		fmt.Println("begin/end",begin,end)
	}
	fmt.Println("Get(0)",histories.Get(0))
	histories.Trim(state.HighestKnown,state.HighestKnown)
	begin := FindBegin(state,histories,0)
	if begin > 0 {
		end := FindEnd(state, histories, begin)
		fmt.Println("begin/end",begin,end)
	}
}
