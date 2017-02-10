package testHelper

import (
	"github.com/FactomProject/factomd/common/constants"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
	"time"
)

func NewRepeatingHash(b byte) interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdtestHelperNewRepeatingHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	tmp := make([]byte, constants.HASH_LENGTH)
	for i := range tmp {
		tmp[i] = b
	}
	return primitives.NewHash(tmp)
}
