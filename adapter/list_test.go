package adapter

import (
	"github.com/gone-io/gone"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type listTester struct {
	gone.Flag
	l IList `gone:"gone-redis-list,test-str"`
}

func TestList(t *testing.T) {
	gone.Test(func(u *listTester) {
		l := u.l

		t.Run("t-BLPOP & BRPOP", func(t *testing.T) {
			firtsList := "list-one"
			secondList := "list-two"

			pop, err := l.BLPop(1, firtsList, secondList)
			time.Sleep(1 * time.Second)
			assert.NotNil(t, err)

			pop, err = l.BRPop(1, firtsList, secondList)
			time.Sleep(1 * time.Second)
			assert.NotNil(t, err)

			err = l.LPush(firtsList, "one")
			assert.Nil(t, err)
			err = l.RPush(firtsList, "two", "three", 4, 5)
			assert.Nil(t, err)

			err = l.LPush(secondList, "six")
			assert.Nil(t, err)
			err = l.RPush(secondList, "seven", "eight", 9, 10)
			assert.Nil(t, err)

			pop, err = l.BLPop(1, firtsList, secondList)
			time.Sleep(1 * time.Second)
			assert.Nil(t, err)
			assert.EqualValues(t, "one", pop[1])

			pop, err = l.BRPop(1, secondList, firtsList)
			time.Sleep(1 * time.Second)
			assert.Nil(t, err)
			assert.EqualValues(t, "10", pop[1])
		})

		t.Run("t-lindex & linsert & len & Lpop & rpop", func(t *testing.T) {
			lk := "list-three"

			err := l.LPush(lk, "one", 2, 2.3, "four")
			assert.Nil(t, err)

			err = l.LInsert(lk, After, "2", "2.21")
			assert.Nil(t, err)

			lLen, err := l.LLen(lk)
			assert.Nil(t, err)
			assert.EqualValues(t, 5, lLen)

			pops, err := l.LPop(lk, 2)
			assert.Nil(t, err)
			assert.EqualValues(t, "2.3", pops[1])

			pops, err = l.RPop(lk, 2)
			assert.Nil(t, err)
			assert.EqualValues(t, "2.21", pops[1])

			pops, err = l.RPop(lk, 10)
			assert.Nil(t, err)
		})

		t.Run("t-LpushX & rpushX & LRange", func(t *testing.T) {
			lk := "list-four"
			err := l.LPushX(lk, "one", 2)
			assert.Nil(t, err)

			err = l.RPushX(lk, "three", 4)
			assert.Nil(t, err)

			err = l.LPush(lk, "one", 2, 3, "IV", "V", "six")
			assert.Nil(t, err)

			ranges, err := l.LRange(lk, 0, 2)
			assert.Nil(t, err)
			assert.EqualValues(t, "V", ranges[1])

			lLen, err := l.LLen(lk)
			assert.Nil(t, err)
			assert.EqualValues(t, lLen, 6)
		})

		t.Run("t-Lrem & Lset & Ltrim", func(t *testing.T) {
			lk := "list-five"
			err := l.LPush(lk, "six", 5, 3, "four", 3, 3, 3, "II", "I", 0)
			assert.Nil(t, err)

			err = l.LRem(lk, 2, 3)
			assert.Nil(t, err)

			val, err := l.LIndex(lk, 4)
			assert.Nil(t, err)
			assert.EqualValues(t, val, "four")

			err = l.LRem(lk, -1, 3)
			assert.Nil(t, err)

			val, err = l.LIndex(lk, 5)
			assert.Nil(t, err)
			assert.EqualValues(t, "5", val)

			err = l.LSet(lk, 5, "five")
			assert.Nil(t, err)
			val, err = l.LIndex(lk, 5)
			assert.Nil(t, err)
			assert.EqualValues(t, "five", val)

			err = l.LTrim(lk, 0, 4)
			assert.Nil(t, err)
			ranges, err := l.LRange(lk, 0, -1)
			assert.Nil(t, err)
			assert.EqualValues(t, "four", ranges[len(ranges)-1])
		})
	}, func(cemetery gone.Cemetery) error {
		cemetery.Bury(&listTester{})
		return nil
	}, Priest)
}
