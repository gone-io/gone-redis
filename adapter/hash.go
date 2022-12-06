package adapter

import (
	"github.com/gomodule/redigo/redis"
	"github.com/gone-io/gone"
)

func NewHash() (gone.Angel, gone.GonerId) {
	return &hash{}, IdRedisHash
}

type hash struct {
	gone.Flag
	redisPool *RedisPool `gone:"gone-redis-pool"`
}

func (h hash) Start(cemetery gone.Cemetery) error {
	//do nothing
	return nil
}

func (h hash) Stop(cemetery gone.Cemetery) error {
	//do nothing
	return nil
}

func (h hash) HGetAll(key string) ([]string, error) {
	conn := h.redisPool.getConn()
	defer h.redisPool.CloseConn(conn)

	return redis.Strings(conn.Do("HGETALL", key))
}

func (h hash) HGet(key string, field any) (string, error) {
	conn := h.redisPool.getConn()
	defer h.redisPool.CloseConn(conn)

	return redis.String(conn.Do("HGET", key, field))
}

func (h hash) HMGet(key string, fields ...any) ([]string, error) {
	conn := h.redisPool.getConn()
	defer h.redisPool.CloseConn(conn)

	m := make([]any, 0)
	m = append(m, key)
	m = append(m, fields...)

	args := toArgs(m...)

	return redis.Strings(conn.Do("HMGET", args...))
}

func (h hash) HSet(key string, field string, value any) error {
	conn := h.redisPool.getConn()
	defer h.redisPool.CloseConn(conn)

	_, err := conn.Do("HSET", key, field, value)

	return err
}

func (h hash) HSetNX(key string, field string, value any) (bool, error) {
	conn := h.redisPool.getConn()
	defer h.redisPool.CloseConn(conn)

	return redis.Bool(conn.Do("HSETNX", key, field, value))
}

func (h hash) HMSet(key string, kvs map[string]any) error {
	conn := h.redisPool.getConn()
	defer h.redisPool.CloseConn(conn)

	m := make([]any, 0)
	m = append(m, key)
	m = append(m, mapToArgs(kvs)...)

	args := toArgs(m...)

	_, err := conn.Do("HMSET", args...)

	return err
}

func (h hash) HIncrBy(key string, field string, incr int64) error {
	conn := h.redisPool.getConn()
	defer h.redisPool.CloseConn(conn)

	_, err := conn.Do("HINCRBY", key, field, incr)

	return err
}

func (h hash) HKeys(key string) ([]string, error) {
	conn := h.redisPool.getConn()
	defer h.redisPool.CloseConn(conn)

	return redis.Strings(conn.Do("HKEYS", key))
}

func (h hash) HVals(key string) ([]string, error) {
	conn := h.redisPool.getConn()
	defer h.redisPool.CloseConn(conn)

	return redis.Strings(conn.Do("HVALS", key))
}

func (h hash) HLen(key string) (int64, error) {
	conn := h.redisPool.getConn()
	defer h.redisPool.CloseConn(conn)

	return redis.Int64(conn.Do("HLEN", key))
}

func (h hash) HDel(key string, field string) error {
	conn := h.redisPool.getConn()
	defer h.redisPool.CloseConn(conn)

	_, err := conn.Do("HDEL", key, field)

	return err
}

func (h hash) HExists(key string, field string) (bool, error) {
	conn := h.redisPool.getConn()
	defer h.redisPool.CloseConn(conn)

	return redis.Bool(conn.Do("HEXISTS", key, field))
}
