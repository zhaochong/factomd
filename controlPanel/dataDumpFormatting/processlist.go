package dataDumpFormatting

import (
	"github.com/FactomProject/factomd/state"
	"time"
)

func RawProcessList(copyDS state.DisplayState) string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddataDumpFormattingRawProcessList.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	/*if st.IdentityChainID == nil {
		return ""
	}
	b := st.CurrentNodeHeight
	pl := st.ProcessLists.Get(b)
	if pl == nil {
		return ""
	}
	return pl.String()*/
	return "Currently undergoing concurrency fixes."
}

func RawPrintMap(copyDS state.DisplayState) string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddataDumpFormattingRawPrintMap.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return copyDS.PrintMap
}
