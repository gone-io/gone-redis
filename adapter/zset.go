package adapter

import (
	"github.com/gomodule/redigo/redis"
	"github.com/gone-io/gone"
	"strconv"
)

func NewZSet() (gone.Angel, gone.GonerId) {
	return &zset{}, IdRedisZSet
}

type zset struct {
	gone.Flag
	redisPool *RedisPool `gone:"gone-redis-pool"`
}

func (z *zset) Start(gone.Cemetery) error {
	//do nothing
	return nil
}

func (z *zset) Stop(gone.Cemetery) error {
	//do nothing
	return nil
}

func (z *zset) ZAdd(key string, score int64, member any) error {
	conn := z.redisPool.getConn()
	defer z.redisPool.CloseConn(conn)

	_, err := conn.Do("ZADD", key, score, member)

	return err
}

func (z *zset) ZCard(key string) (int64, error) {
	conn := z.redisPool.getConn()
	defer z.redisPool.CloseConn(conn)

	return redis.Int64(conn.Do("ZCARD", key))
}

func (z *zset) ZCount(key string, min, max int64, includeMin, includeMax bool) (int64, error) {
	conn := z.redisPool.getConn()
	defer z.redisPool.CloseConn(conn)

	l := strconv.FormatInt(min, 64)
	h := strconv.FormatInt(max, 64)
	if !includeMin {
		l = "(" + l
	}

	if !includeMax {
		h = "(" + h
	}

	return redis.Int64(conn.Do("ZCOUNT", key, l, h))
}

func (z *zset) ZIncrBy(key string, incr int64, member any) error {
	conn := z.redisPool.getConn()
	defer z.redisPool.CloseConn(conn)

	_, err := conn.Do("ZINCRBY", key, incr, member)

	return err
}

func (z *zset) ZPopMax(key string, count int64) ([]string, error) {
	conn := z.redisPool.getConn()
	defer z.redisPool.CloseConn(conn)

	return redis.Strings(conn.Do("ZPOPMAX", key, count))
}

func (z *zset) ZPopMin(key string, count int64) ([]string, error) {
	conn := z.redisPool.getConn()
	defer z.redisPool.CloseConn(conn)

	return redis.Strings(conn.Do("ZPOPMIN", key, count))
}

func (z *zset) ZRangeWithLimit(key string, start, end int64, orderStrategy ZSetOrderStrategy, offset, count int64, withScore bool) ([]string, error) {
	conn := z.redisPool.getConn()
	defer z.redisPool.CloseConn(conn)

	if withScore {
		return redis.Strings(conn.Do("ZRANGE", key, start, end, orderStrategy, "LIMIT", offset, count, "WITHSCORES"))
	}

	return redis.Strings(conn.Do("ZRANGE", key, start, end, orderStrategy, "LIMIT", offset, count))
}

func (z *zset) ZRangeWithoutLimit(key string, start, end int64, orderStrategy ZSetOrderStrategy, withScore bool) ([]string, error) {
	conn := z.redisPool.getConn()
	defer z.redisPool.CloseConn(conn)

	if withScore {
		return redis.Strings(conn.Do("ZRANGE", key, start, end, orderStrategy, "WITHSCORES"))
	}

	return redis.Strings(conn.Do("ZRANGE", key, start, end, orderStrategy))
}

func (z *zset) ZRank(key string, member any) (int64, error) {
	conn := z.redisPool.getConn()
	defer z.redisPool.CloseConn(conn)

	return redis.Int64(conn.Do("ZRANK", key, member))
}

func (z *zset) ZRem(key string, members ...any) error {
	conn := z.redisPool.getConn()
	defer z.redisPool.CloseConn(conn)

	args := redis.Args{}.Add(key).Add(members...)

	_, err := conn.Do("ZREM", args...)

	return err
}

func (z *zset) ZRemRangeByRank(key string, start, end int64) error {
	conn := z.redisPool.getConn()
	defer z.redisPool.CloseConn(conn)

	_, err := conn.Do("ZREMRANGEBYRANK", key, start, end)

	return err
}

func (z *zset) ZRemRangeByScore(key string, min, max int64) error {
	conn := z.redisPool.getConn()
	defer z.redisPool.CloseConn(conn)

	_, err := conn.Do("ZREMRANGEBYSCORE", key, min, max)

	return err
}

func (z *zset) ZRevRange(key string, start, stop int64, withScores bool) ([]string, error) {
	conn := z.redisPool.getConn()
	defer z.redisPool.CloseConn(conn)

	if withScores {
		return redis.Strings(conn.Do("ZREVRANGE", key, start, stop, "WITHSCORES"))
	}

	return redis.Strings(conn.Do("ZREVRANGE", key, start, stop))
}

func (z *zset) ZRevRangeByScoreWithLimit(key string, max, min int64, withScores bool, offset, count int64) ([]string, error) {
	conn := z.redisPool.getConn()
	defer z.redisPool.CloseConn(conn)

	if withScores {
		return redis.Strings(conn.Do("ZREVRANGEBYSCORE", key, max, min, "LIMIT", offset, count, "WITHSCORES"))
	}

	return redis.Strings(conn.Do("ZREVRANGEBYSCORE", key, max, min, "LIMIT", offset, count))
}

func (z *zset) ZRevRangeByScoreWithoutLimit(key string, max, min int64, withScores bool) ([]string, error) {
	conn := z.redisPool.getConn()
	defer z.redisPool.CloseConn(conn)

	if withScores {
		return redis.Strings(conn.Do("ZREVRANGEBYSCORE", key, max, min, "WITHSCORES"))
	}

	return redis.Strings(conn.Do("ZREVRANGEBYSCORE", key, max, min))
}

func (z *zset) ZScore(key string, member any) (float64, error) {
	conn := z.redisPool.getConn()
	defer z.redisPool.CloseConn(conn)

	return redis.Float64(conn.Do("ZSCORE", key, member))
}
