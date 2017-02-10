// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package random

import (
	"crypto/rand"
	"math"
	"math/big"
	"time"
)

func RandUInt64() uint64 {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdrandomRandUInt64.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return RandUInt64Between(0, math.MaxUint64)
}

func RandUInt64Between(min, max uint64) uint64 {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdrandomRandUInt64Between.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if min >= max {
		return 0
	}
	uint64max := big.NewInt(0)
	uint64max.SetUint64(max - min)
	randnum, _ := rand.Int(rand.Reader, uint64max)
	m := big.NewInt(0)
	m.SetUint64(min)
	randnum = randnum.Add(randnum, m)
	return randnum.Uint64()
}

func RandInt64() int64 {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdrandomRandInt64.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	int64max := big.NewInt(math.MaxInt64)
	randnum, _ := rand.Int(rand.Reader, int64max)
	return randnum.Int64()
}

func RandInt64Between(min, max int64) int64 {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdrandomRandInt64Between.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if min >= max {
		return 0
	}
	int64max := big.NewInt(max - min)
	randnum, _ := rand.Int(rand.Reader, int64max)
	m := big.NewInt(min)
	randnum = randnum.Add(randnum, m)
	return randnum.Int64()
}

func RandInt() int {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdrandomRandInt.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return int(RandInt64())
}

func RandIntBetween(min, max int) int {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdrandomRandIntBetween.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return int(RandInt64Between(int64(min), int64(max)))
}

func RandByteSlice() []byte {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdrandomRandByteSlice.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	l := RandInt() % 64
	answer := make([]byte, l)
	_, err := rand.Read(answer)
	if err != nil {
		return nil
	}
	return answer
}

func RandNonEmptyByteSlice() []byte {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdrandomRandNonEmptyByteSlice.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	l := RandInt()%63 + 1
	answer := make([]byte, l)
	_, err := rand.Read(answer)
	if err != nil {
		return nil
	}
	return answer
}

func RandByteSliceOfLen(l int) []byte {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdrandomRandByteSliceOfLen.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if l <= 0 {
		return nil
	}
	answer := make([]byte, l)
	_, err := rand.Read(answer)
	if err != nil {
		return nil
	}
	return answer
}

const StringAlphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"

func RandomString() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdrandomRandomString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	l := RandIntBetween(0, 128)
	answer := []byte{}
	for i := 0; i < l; i++ {
		answer = append(answer, StringAlphabet[RandIntBetween(0, len(StringAlphabet)-1)])
	}
	return string(answer)
}
