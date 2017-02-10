package controlPanel

import (
	"encoding/json"
	"time"
)

// Used to send a height as json struct
func HeightToJsonStruct(height uint32) []byte {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdcontrolPanelHeightToJsonStruct.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	jData, err := json.Marshal(struct{ Height uint32 }{height})
	if err != nil {
		return []byte(`{"Height":0}`)
	}
	return jData
}
