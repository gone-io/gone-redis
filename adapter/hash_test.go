package adapter

import (
	"fmt"
	"github.com/gone-io/gone"
	"github.com/stretchr/testify/assert"
	"testing"
)

type hashTester struct {
	gone.Flag
	h IHash `gone:"gone-redis-hash, test-hash"`
}

func TestNewHash(t *testing.T) {
	gone.Test(func(u *hashTester) {
		h := u.h

		t.Run("t-HGETALL & HGET & HMGET & HSET & HMSET & HSETNX", func(t *testing.T) {
			hk := "hash-one"
			fields, err := h.HGetAll(hk)
			assert.Nil(t, err)
			assert.EqualValues(t, len(fields), 0)

			err = h.HSet(hk, "one", 1)
			assert.Nil(t, err)

			kvs := make(map[string]any)
			kvs["two"] = 2
			kvs["three"] = 3
			kvs["four"] = "IV"
			err = h.HMSet(hk, kvs)
			assert.Nil(t, err)

			ok, err := h.HSetNX(hk, "four", 4)
			assert.Nil(t, err)
			assert.EqualValues(t, false, ok)

			ok, err = h.HSetNX(hk, "five", "five")
			assert.Nil(t, err)
			assert.EqualValues(t, true, ok)

			fields, err = h.HGetAll(hk)
			assert.Nil(t, err)
			assert.EqualValues(t, 10, len(fields))

			val, err := h.HGet(hk, "three")
			assert.Nil(t, err)
			assert.EqualValues(t, "3", val)

			vals, err := h.HMGet(hk, "one", "two")
			assert.Nil(t, err)
			assert.EqualValues(t, "1", vals[0])
			assert.EqualValues(t, "2", vals[1])
		})

		t.Run("t-Hincrby & Hkeys & Hvalues & Hlen & Hdel & Hexists", func(t *testing.T) {
			hk := "hash-two"

			err := h.HIncrBy(hk, "field-one", 5)
			assert.Nil(t, err)

			err = h.HIncrBy(hk, "field-two", -5)
			assert.Nil(t, err)

			err = h.HIncrBy(hk, "field-one", -7)
			assert.Nil(t, err)

			val, err := h.HGet(hk, "field-one")
			assert.Nil(t, err)
			assert.EqualValues(t, "-2", val)

			keys, err := h.HKeys(hk)
			assert.Nil(t, err)
			assert.EqualValues(t, 2, len(keys))
			fmt.Printf("%v", keys)

			vals, err := h.HVals(hk)
			assert.Nil(t, err)
			assert.EqualValues(t, 2, len(keys))
			fmt.Printf("%v", vals)

			err = h.HDel(hk, "field-one")
			assert.Nil(t, err)

			keys, err = h.HKeys(hk)
			assert.Nil(t, err)
			assert.EqualValues(t, 1, len(keys))
			fmt.Printf("%v", keys)

			ok, err := h.HExists(hk, "field-one")
			assert.Nil(t, err)
			assert.EqualValues(t, false, ok)

			ok, err = h.HExists(hk, "field-two")
			assert.Nil(t, err)
			assert.EqualValues(t, true, ok)
		})

	}, func(cemetery gone.Cemetery) error {
		cemetery.Bury(&hashTester{})
		return nil
	}, Priest)
}
