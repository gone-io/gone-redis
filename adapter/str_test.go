package adapter

import (
	"github.com/gomodule/redigo/redis"
	"github.com/gone-io/gone"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

type RedisTester struct {
	gone.Flag
	s IStr `gone:"gone-redis-str,test-str"`
}

func TestHash(t *testing.T) {
	gone.Test(func(u *RedisTester) {
		s := u.s

		t.Run("t-append", func(t *testing.T) {
			n := rand.Intn(100)
			key := "point-a"
			err := s.Append(key, n)
			assert.Nil(t, err)

			a := "append content"
			err = s.Append(key, a)
			assert.Nil(t, err)
		})

		t.Run("t-decr & t-decrby & Get & incr & incrBy", func(t *testing.T) {
			key := "aNum"
			err := s.Incr(key)
			assert.Nil(t, err)

			err = s.IncrBy(key, 10)
			assert.Nil(t, err)

			val, err := s.Get(key)
			assert.Nil(t, err)

			actualVal, _ := strconv.Atoi(val)
			assert.EqualValues(t, 11, actualVal)

			err = s.Decr(key)
			assert.Nil(t, err)

			err = s.DecrBy(key, 5)
			assert.Nil(t, err)

			val, err = s.Get(key)
			assert.Nil(t, err)

			actualVal, _ = strconv.Atoi(val)
			assert.EqualValues(t, 5, actualVal)
		})

		t.Run("t-GetRange & Set", func(t *testing.T) {
			key := "aStr"
			err := s.Set(key, "Hello World.")
			assert.Nil(t, err)

			val, err := s.GetRange(key, 0, 3)
			assert.Nil(t, err)
			assert.EqualValues(t, "Hell", val)

			val, err = s.GetRange(key, 8, 100)
			assert.Nil(t, err)
			val2, err := s.GetRange(key, 8, -1)
			assert.Nil(t, err)
			assert.EqualValues(t, val2, val)
			assert.EqualValues(t, "rld.", val)

			val, err = s.GetRange(key, 99, -1)
			assert.Nil(t, err)
			assert.EqualValues(t, "", val)
		})

		t.Run("t-GetRange & Set", func(t *testing.T) {
			key := "aStr"
			_ = s.Set(key, "Hello world.")

			oldVal, err := s.GetSet(key, "Hello golang")
			assert.Nil(t, err)
			assert.EqualValues(t, "Hello world.", oldVal)

			newVal, err := s.Get(key)
			assert.Nil(t, err)
			assert.EqualValues(t, "Hello golang", newVal)
		})

		t.Run("t-incrByFloat", func(t *testing.T) {
			key := "aFloat"
			// init
			_ = s.Set(key, 0)

			err := s.IncrByFloat(key, 3)
			assert.Nil(t, err)

			err = s.IncrByFloat(key, 0.3)
			assert.Nil(t, err)

			val, err := s.Get(key)
			assert.Nil(t, err)
			// always inaccurate
			// Expected :3.3
			// Actual   :3.29999999999999982
			assert.EqualValues(t, "3.3", val)
		})

		t.Run("t-MGET & MSET & MSETNX & SETNX & SETEX", func(t *testing.T) {
			kvs := make(map[string]any)
			kvs["firstKey"] = 1
			kvs["secondKey"] = "two"
			kvs["thirdKey"] = "3"
			kvs["fourthKey"] = "four"

			err := s.MSet(kvs)
			assert.Nil(t, err)

			vals, err := s.MGet("firstKey", "secondKey")
			assert.Nil(t, err)
			assert.EqualValues(t, "1", vals[0])
			assert.EqualValues(t, "two", vals[1])

			kvs["five"] = 5
			ok, err := s.MSetNX(kvs)
			assert.Nil(t, err)
			assert.EqualValues(t, false, ok)

			kvs = make(map[string]any)
			kvs["fifthKey"] = 5
			ok, err = s.MSetNX(kvs)
			assert.Nil(t, err)
			assert.EqualValues(t, true, ok)

			ok, err = s.SetNX("fifthKey", 55)
			assert.Nil(t, err)
			assert.EqualValues(t, ok, false)

			ok, err = s.SetNX("sixthKey", 6)
			assert.Nil(t, err)
			assert.EqualValues(t, ok, true)

			err = s.SetEX("seventhKey", 1, 7)
			assert.Nil(t, err)
			time.Sleep(1 * time.Second)

			_, err = s.Get("seventhKey")
			assert.NotNil(t, err)
			assert.EqualValues(t, redis.ErrNil, err)
		})

		t.Run("t-setRange & strLen", func(t *testing.T) {
			key := "aStr"
			err := s.Set(key, "hello, netty")
			assert.Nil(t, err)

			err = s.SetRange(key, 6, "python")
			assert.Nil(t, err)

			newVal, err := s.Get(key)
			assert.Nil(t, err)
			assert.EqualValues(t, "hello,python", newVal)

			strLen, err := s.StrLen(key)
			assert.Nil(t, err)
			assert.EqualValues(t, 12, strLen)
		})

	}, func(cemetery gone.Cemetery) error {
		cemetery.Bury(&RedisTester{})
		return nil
	}, Priest)
}
