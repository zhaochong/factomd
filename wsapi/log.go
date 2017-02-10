// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package wsapi

import (
	"github.com/FactomProject/factomd/logger"
	"time"
)

// setup subsystem loggers
var (
	rpcLog    *logger.FLogger
	serverLog *logger.FLogger
	wsLog     *logger.FLogger
)

func InitLogs(logPath, logLevel string) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdwsapiInitLogs.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	rpcLog = logger.NewLogFromConfig(logPath, logLevel, "RPC")
	serverLog = logger.NewLogFromConfig(logPath, logLevel, "SERV")
	wsLog = logger.NewLogFromConfig(logPath, logLevel, "WSAPI")
}
