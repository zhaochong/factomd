package cacheDB_test

import (
	"crypto/rand"
	"fmt"
	"testing"
	"time"

	"github.com/FactomProject/factomd/common/interfaces"
	. "github.com/FactomProject/factomd/database/cacheDB"
)

type TestData struct {
	Str string
}

func (t *TestData) New() interfaces.BinaryMarshallableAndCopyable {
	return new(TestData)
}

func (t *TestData) MarshalBinary() ([]byte, error) {
	return []byte(t.Str), nil
}

func (t *TestData) UnmarshalBinaryData(data []byte) ([]byte, error) {
	t.Str = string(data)
	return nil, nil
}

func (t *TestData) UnmarshalBinary(data []byte) (err error) {
	_, err = t.UnmarshalBinaryData(data)
	return
}

var _ interfaces.BinaryMarshallable = (*TestData)(nil)

func TestPutGetDelete(t *testing.T) {
	m := new(CacheDB)

	key := []byte("key")
	bucket := []byte("bucket")

	test := new(TestData)
	test.Str = "testtest"

	err := m.Put(bucket, key, test)
	if err != nil {
		t.Errorf("%v", err)
	}

	resp, err := m.Get(bucket, key, new(TestData))
	if err != nil {
		t.Errorf("%v", err)
	}

	if resp == nil {
		t.Errorf("resp is nil")
	}

	if resp.(*TestData).Str != test.Str {
		t.Errorf("data mismatch")
	}

	err = m.Delete(bucket, key)
	if err != nil {
		t.Errorf("%v", err)
	}

	resp, err = m.Get(bucket, key, new(TestData))
	if err != nil {
		t.Errorf("%v", err)
	}
	if resp != nil {
		t.Errorf("resp is not nil while it should be")
	}
}

func TestParallelAccess(t *testing.T) {
	threads := 100
	m := new(CacheDB)
	c := make(chan int)
	closed := make(chan int, threads)
	for i := 0; i < threads; i++ {
		go func() {
			for {
				select {
				case <-c:
					closed <- 1
					return
				default:
					str := new(TestData)
					str.Str = fmt.Sprintf("%x", RandomHex(32))
					err := m.Put(RandomHex(5), RandomHex(5), str)
					if err != nil {
						t.Errorf("Got error - %v", err)
					}
					_, err = m.Get(RandomHex(5), RandomHex(5), str)
					if err != nil {
						t.Errorf("Got error - %v", err)
					}
				}
			}
		}()
	}
	time.Sleep(10 * time.Second)
	close(c)
	time.Sleep(1 * time.Second)
	for i := 0; i < threads; i++ {
		<-closed
	}
}

func RandomHex(length int) []byte {
	if length <= 0 {
		return nil
	}
	answer := make([]byte, length)
	_, err := rand.Read(answer)
	if err != nil {
		return nil
	}
	return answer
}
