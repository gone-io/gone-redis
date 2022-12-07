package adapter

import (
	"fmt"
	"github.com/gone-io/gone"
	"github.com/stretchr/testify/assert"
	"testing"
)

type zsetTester struct {
	gone.Flag
	z IZSet `gone:"gone-redis-zset,test-str"`
}

func TestZset(t *testing.T) {
	gone.Test(func(u *zsetTester) {
		z := u.z

		t.Run("t-Zadd & zcard & zscore & zcount & zincrby & zpopmax & zpopmin & zrangewithlimit & zrangwithoutlimit", func(t *testing.T) {
			zkey := "zkey-one"
			err := z.ZAdd(zkey, 100, "mem1")
			assert.Nil(t, err)

			err = z.ZAdd(zkey, 200, "mem2")
			assert.Nil(t, err)

			err = z.ZAdd(zkey, 300, "mem3")
			assert.Nil(t, err)

			err = z.ZAdd(zkey, 400, "mem4")
			assert.Nil(t, err)

			err = z.ZAdd(zkey, 500, "mem5")
			assert.Nil(t, err)

			err = z.ZAdd(zkey, 600, "mem6")
			assert.Nil(t, err)

			err = z.ZAdd(zkey, 700, "mem7")
			assert.Nil(t, err)

			cnt, err := z.ZCard(zkey)
			assert.Nil(t, err)
			assert.EqualValues(t, 7, cnt)

			score, err := z.ZScore(zkey, "mem1")
			assert.Nil(t, err)
			assert.EqualValues(t, 100, score)

			err = z.ZIncrBy(zkey, 22, "mem1")
			assert.Nil(t, err)

			score, err = z.ZScore(zkey, "mem1")
			assert.Nil(t, err)
			assert.EqualValues(t, 122, score)

			max, err := z.ZPopMax(zkey, 2)
			assert.Nil(t, err)
			assert.EqualValues(t, 4, len(max))
			fmt.Printf("%v \n", max)

			min, err := z.ZPopMin(zkey, 2)
			assert.Nil(t, err)
			assert.EqualValues(t, 4, len(min))
			fmt.Printf("%v \n", min)

			res, err := z.ZRangeWithLimit(zkey, 100, 300, ByScore, 0, 2, true)
			assert.Nil(t, err)
			fmt.Printf("%v \n", res)

			res, err = z.ZRangeWithoutLimit(zkey, 100, 300, ByScore, true)
			assert.Nil(t, err)
			fmt.Printf("%v \n", res)
		})

		t.Run("t-Zrank & zrem & zremrangebyrank & zremrangebyscore & zrevrange & zrevrangebyscorewithlimit & zrevrangebyscorewithoutlimit", func(t *testing.T) {
			zkey := "zkey-two"
			err := z.ZAdd(zkey, 100, "mem1")
			assert.Nil(t, err)

			err = z.ZAdd(zkey, 200, "mem2")
			assert.Nil(t, err)

			err = z.ZAdd(zkey, 300, "mem3")
			assert.Nil(t, err)

			err = z.ZAdd(zkey, 400, "mem4")
			assert.Nil(t, err)

			err = z.ZAdd(zkey, 500, "mem5")
			assert.Nil(t, err)

			err = z.ZAdd(zkey, 600, "mem6")
			assert.Nil(t, err)

			err = z.ZAdd(zkey, 700, "mem7")
			assert.Nil(t, err)

			rank, err := z.ZRank(zkey, "mem6")
			assert.Nil(t, err)
			assert.EqualValues(t, 5, rank)

			err = z.ZRemRangeByRank(zkey, 6, 7)
			assert.Nil(t, err)
			res, err := z.ZRangeWithoutLimit(zkey, 0, 1000, ByScore, true)
			assert.Nil(t, err)
			fmt.Printf("%v \n", res)

			err = z.ZRemRangeByScore(zkey, 400, 499)
			assert.Nil(t, err)
			res, err = z.ZRangeWithoutLimit(zkey, 0, 1000, ByScore, true)
			assert.Nil(t, err)
			fmt.Printf("%v \n", res)

			res, err = z.ZRevRange(zkey, 100, 599, true)
			assert.Nil(t, err)
			fmt.Printf("%v \n", res)

			res, err = z.ZRevRangeByScoreWithLimit(zkey, 700, 500, true, 1, 1)
			assert.Nil(t, err)
			fmt.Printf("%v \n", res)

			res, err = z.ZRevRangeByScoreWithoutLimit(zkey, 700, 500, true)
			assert.Nil(t, err)
			fmt.Printf("%v \n", res)
		})

	}, func(cemetery gone.Cemetery) error {
		cemetery.Bury(&zsetTester{})
		return nil
	}, Priest)
}
