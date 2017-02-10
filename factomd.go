// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"github.com/FactomProject/factomd/engine"
	"time"
)

func main() {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdmainmain.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	// uncomment StartProfiler() to run the pprof tool (for testing)
	engine.Factomd()
}
