// Copyright 2015 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cacheDB

import (
	"fmt"
	"time"

	"github.com/FactomProject/factomd/common/interfaces"

	"github.com/patrickmn/go-cache"
)

type CacheDB struct {
	Cache *cache.Cache
}

var _ interfaces.IDatabase = (*CacheDB)(nil)

func (CacheDB) Close() error {
	return nil
}

func (db *CacheDB) ListAllBuckets() ([][]byte, error) {
	return nil, fmt.Errorf("Function not suporter for cache")
}

func (db *CacheDB) Trim() {
	db.Init()
	db.Cache.Flush()
}

func (db *CacheDB) Init() {
	if db.Cache == nil {
		db.Cache = cache.New(30*time.Second, 10*time.Second)
	}
}

func (db *CacheDB) Put(bucket, key []byte, data interfaces.BinaryMarshallable) error {
	db.Init()
	return db.rawPut(bucket, key, data)
}

func (db *CacheDB) RawPut(bucket, key []byte, data interfaces.BinaryMarshallable) error {
	db.Init()
	return db.rawPut(bucket, key, data)
}

func (db *CacheDB) rawPut(bucket, key []byte, data interfaces.BinaryMarshallable) error {
	db.Init()

	k := CombineBucketAndKey(bucket, key)

	if data == nil {
		db.Cache.Delete(k)
		return nil
	}

	var hex []byte
	var err error
	if data != nil {
		hex, err = data.MarshalBinary()
		if err != nil {
			return err
		}
	}

	db.Cache.Set(k, hex, cache.DefaultExpiration)

	return nil
}

func (db *CacheDB) PutInBatch(records []interfaces.Record) error {
	db.Init()
	for _, v := range records {
		err := db.rawPut(v.Bucket, v.Key, v.Data)
		if err != nil {
			return err
		}
	}
	return nil
}

func (db *CacheDB) Get(bucket, key []byte, destination interfaces.BinaryMarshallable) (interfaces.BinaryMarshallable, error) {
	db.Init()

	k := CombineBucketAndKey(bucket, key)
	v, ok := db.Cache.Get(k)
	if ok == false {
		return nil, nil
	}
	if v == nil {
		return nil, nil
	}
	_, err := destination.UnmarshalBinaryData(v.([]byte))
	if err != nil {
		return nil, err
	}
	return destination, nil
}

func (db *CacheDB) Delete(bucket, key []byte) error {
	db.Init()

	k := CombineBucketAndKey(bucket, key)
	db.Cache.Delete(k)

	return nil
}

func (db *CacheDB) ListAllKeys(bucket []byte) ([][]byte, error) {
	return nil, fmt.Errorf("Function not suporter for cache")
}

func (db *CacheDB) GetAll(bucket []byte, sample interfaces.BinaryMarshallableAndCopyable) ([]interfaces.BinaryMarshallableAndCopyable, [][]byte, error) {
	return nil, nil, fmt.Errorf("Function not suporter for cache")
}

func (db *CacheDB) Clear(bucket []byte) error {
	db.Trim()
	return nil
}

func CombineBucketAndKey(bucket []byte, key []byte) string {
	ldbKey := ExtendBucket(bucket)
	ldbKey = append(ldbKey, key...)
	return fmt.Sprintf("%x", ldbKey)
}

func ExtendBucket(bucket []byte) []byte {
	return append(bucket, ';')
}
