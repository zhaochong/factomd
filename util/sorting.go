package util

import (
	//"github.com/FactomProject/factomd/common/directoryBlock/dbInfo"
	"github.com/FactomProject/factomd/common/interfaces"

	"bytes"
	"time"
)

//------------------------------------------------
// DBlock array sorting implementation - accending
type ByDBlockIDAccending []interfaces.IDirectoryBlock

func (f ByDBlockIDAccending) Len() int {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdutilByDBlockIDAccendingLen.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return len(f)
}
func (f ByDBlockIDAccending) Less(i, j int) bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdutilByDBlockIDAccendingLess.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return f[i].GetHeader().GetDBHeight() < f[j].GetHeader().GetDBHeight()
}
func (f ByDBlockIDAccending) Swap(i, j int) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdutilByDBlockIDAccendingSwap.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	f[i], f[j] = f[j], f[i]
}

//------------------------------------------------
// CBlock array sorting implementation - accending
type ByECBlockIDAccending []interfaces.IEntryCreditBlock

func (f ByECBlockIDAccending) Len() int {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdutilByECBlockIDAccendingLen.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return len(f)
}
func (f ByECBlockIDAccending) Less(i, j int) bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdutilByECBlockIDAccendingLess.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return f[i].GetHeader().GetDBHeight() < f[j].GetHeader().GetDBHeight()
}
func (f ByECBlockIDAccending) Swap(i, j int) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdutilByECBlockIDAccendingSwap.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	f[i], f[j] = f[j], f[i]
}

//------------------------------------------------
// ABlock array sorting implementation - accending
type ByABlockIDAccending []interfaces.IAdminBlock

func (f ByABlockIDAccending) Len() int {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdutilByABlockIDAccendingLen.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return len(f)
}
func (f ByABlockIDAccending) Less(i, j int) bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdutilByABlockIDAccendingLess.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return f[i].GetHeader().GetDBHeight() < f[j].GetHeader().GetDBHeight()
}
func (f ByABlockIDAccending) Swap(i, j int) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdutilByABlockIDAccendingSwap.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	f[i], f[j] = f[j], f[i]
}

//------------------------------------------------
// ABlock array sorting implementation - accending
type ByFBlockIDAccending []interfaces.IFBlock

func (f ByFBlockIDAccending) Len() int {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdutilByFBlockIDAccendingLen.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return len(f)
}
func (f ByFBlockIDAccending) Less(i, j int) bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdutilByFBlockIDAccendingLess.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return f[i].GetDBHeight() < f[j].GetDBHeight()
}
func (f ByFBlockIDAccending) Swap(i, j int) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdutilByFBlockIDAccendingSwap.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	f[i], f[j] = f[j], f[i]
}

//------------------------------------------------
// EBlock array sorting implementation - accending
type ByEBlockIDAccending []interfaces.IEntryBlock

func (f ByEBlockIDAccending) Len() int {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdutilByEBlockIDAccendingLen.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return len(f)
}
func (f ByEBlockIDAccending) Less(i, j int) bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdutilByEBlockIDAccendingLess.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return f[i].GetHeader().GetEBSequence() < f[j].GetHeader().GetEBSequence()
}
func (f ByEBlockIDAccending) Swap(i, j int) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdutilByEBlockIDAccendingSwap.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	f[i], f[j] = f[j], f[i]
}

//------------------------------------------------
// Byte array sorting - ascending
type ByByteArray [][]byte

func (f ByByteArray) Len() int {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdutilByByteArrayLen.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return len(f)
}
func (f ByByteArray) Less(i, j int) bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdutilByByteArrayLess.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return bytes.Compare(f[i], f[j]) < 0
}
func (f ByByteArray) Swap(i, j int) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdutilByByteArraySwap.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	f[i], f[j] = f[j], f[i]
}

//------------------------------------------------
// DirBlock Info array sorting implementation - accending
type ByDirBlockInfoIDAccending []interfaces.IDirBlockInfo

func (f ByDirBlockInfoIDAccending) Len() int {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdutilByDirBlockInfoIDAccendingLen.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return len(f)
}
func (f ByDirBlockInfoIDAccending) Less(i, j int) bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdutilByDirBlockInfoIDAccendingLess.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return f[i].GetDBHeight() < f[j].GetDBHeight()
}
func (f ByDirBlockInfoIDAccending) Swap(i, j int) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdutilByDirBlockInfoIDAccendingSwap.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	f[i], f[j] = f[j], f[i]
}

// ByDirBlockInfoTimestamp defines the methods needed to satisify sort.Interface to
// sort a slice of DirBlockInfo by their Timestamp.
type ByDirBlockInfoTimestamp []interfaces.IDirBlockInfo

func (u ByDirBlockInfoTimestamp) Len() int {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdutilByDirBlockInfoTimestampLen.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////
	return len(u)
}
func (u ByDirBlockInfoTimestamp) Less(i, j int) bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdutilByDirBlockInfoTimestampLess.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if u[i].GetTimestamp() == u[j].GetTimestamp() {
		return u[i].GetDBHeight() < u[j].GetDBHeight()
	}
	return u[i].GetTimestamp().GetTimeMilli() < u[j].GetTimestamp().GetTimeMilli()
}
func (u ByDirBlockInfoTimestamp) Swap(i, j int) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdutilByDirBlockInfoTimestampSwap.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////
	u[i], u[j] = u[j], u[i]
}
