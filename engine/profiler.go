// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package engine

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

// StartProfiler runs the go pprof tool
// `go tool pprof http://localhost:6060/debug/pprof/profile`
// https://golang.org/pkg/net/http/pprof/
func StartProfiler() {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdengineStartProfiler.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	log.Println(http.ListenAndServe(fmt.Sprintf("localhost:%s", logPort), nil))
}
