package log

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"

	"testing"
	"time"
)

var LogLevel int

var TestLogger testing.TB

const (
	StandardLog = 0
	DebugLog    = 1
)

func init() {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdloginit.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	LogLevel = StandardLog
}

func SetTestLogger(tb testing.TB) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdlogSetTestLogger.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	TestLogger = tb
}

func UnsetTestLogger() {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdlogUnsetTestLogger.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	TestLogger = nil
}

func SetLevel(level string) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdlogSetLevel.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if strings.ToLower(level) == "debug" {
		LogLevel = DebugLog
	} else {
		LogLevel = StandardLog
	}
}

func debugPrefix() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdlogdebugPrefix.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	_, file, line, ok := runtime.Caller(3)
	if !ok {
		file = "???"
		line = 0
	}
	return file + ":" + strconv.Itoa(line) + " - "
}

func PrintStack() {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdlogPrintStack.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	debug.PrintStack()
}

func Fatal(str string, args ...interface{}) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdlogFatal.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	printf(LogLevel == DebugLog, str+"\n", args...)
	os.Exit(1)
}

func Print(str string) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdlogPrint.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	printf(LogLevel == DebugLog, str+"\n")
}

func Println(str string) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdlogPrintln.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	printf(LogLevel == DebugLog, str+"\n")
}

func Printf(format string, args ...interface{}) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdlogPrintf.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	printf(LogLevel == DebugLog, format, args...)
}

func Printfln(format string, args ...interface{}) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdlogPrintfln.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	printf(LogLevel == DebugLog, format+"\n", args...)
}

func Debug(format string, args ...interface{}) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdlogDebug.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	printf(true, format+"\n", args...)
}

func printf(debug bool, format string, args ...interface{}) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdlogprintf.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if debug {
		format = debugPrefix() + format
	}
	if len(args) > 0 {
		if TestLogger != nil {
			TestLogger.Logf(format, args...)
		} else {
			fmt.Printf(format, args...)
		}
	} else {
		if TestLogger != nil {
			TestLogger.Log(format)
		} else {
			fmt.Print(format)
		}
	}
}
