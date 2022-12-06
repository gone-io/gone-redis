package adapter

import (
	"github.com/gomodule/redigo/redis"
	"github.com/gone-io/gone"
)

func NewList() (gone.Angel, gone.GonerId) {
	return &list{}, IdRedisList
}

type list struct {
	gone.Flag
	redisPool *RedisPool `gone:"gone-redis-pool"`
}

func (l *list) Start(cemetery gone.Cemetery) error {
	// do nothing
	return nil
}

func (l *list) Stop(cemetery gone.Cemetery) error {
	// do nothing
	return nil
}

func (l *list) BLPop(timeout int64, keys ...any) ([]string, error) {
	conn := l.redisPool.getConn()
	defer l.redisPool.CloseConn(conn)

	m := make([]any, 0)
	m = append(m, keys...)
	m = append(m, timeout)

	args := toArgs(m...)

	return redis.Strings(conn.Do("BLPOP", args...))
}

func (l *list) BRPop(timeout int64, keys ...any) ([]string, error) {
	conn := l.redisPool.getConn()
	defer l.redisPool.CloseConn(conn)

	m := make([]any, 0)
	m = append(m, keys...)
	m = append(m, timeout)

	args := toArgs(m...)

	return redis.Strings(conn.Do("BRPOP", args...))
}

func (l *list) LIndex(key string, idx int64) (string, error) {
	conn := l.redisPool.getConn()
	defer l.redisPool.CloseConn(conn)

	return redis.String(conn.Do("LINDEX", key, idx))
}

func (l *list) LInsert(key string, pos ListPos, pivot, element any) error {
	conn := l.redisPool.getConn()
	defer l.redisPool.CloseConn(conn)

	_, err := conn.Do("LINSERT", key, pos, pivot, element)

	return err
}

func (l *list) LLen(key string) (int64, error) {
	conn := l.redisPool.getConn()
	defer l.redisPool.CloseConn(conn)

	return redis.Int64(conn.Do("LLEN", key))
}

func (l *list) LPop(key string, count int64) ([]string, error) {
	conn := l.redisPool.getConn()
	defer l.redisPool.CloseConn(conn)

	return redis.Strings(conn.Do("LPOP", key, count))
}

func (l *list) RPop(key string, count int64) ([]string, error) {
	conn := l.redisPool.getConn()
	defer l.redisPool.CloseConn(conn)

	return redis.Strings(conn.Do("RPOP", key, count))
}

func (l *list) LPush(key string, elements ...any) error {
	conn := l.redisPool.getConn()
	defer l.redisPool.CloseConn(conn)

	m := make([]any, 0)
	m = append(m, key)
	m = append(m, elements...)

	args := toArgs(m...)

	_, err := conn.Do("LPUSH", args...)

	return err
}

func (l *list) RPush(key string, elements ...any) error {
	conn := l.redisPool.getConn()
	defer l.redisPool.CloseConn(conn)

	m := make([]any, 0)
	m = append(m, key)
	m = append(m, elements...)

	args := toArgs(m...)

	_, err := conn.Do("RPUSH", args...)

	return err
}

func (l *list) LPushX(key string, elements ...any) error {
	conn := l.redisPool.getConn()
	defer l.redisPool.CloseConn(conn)

	m := make([]any, 0)
	m = append(m, key)
	m = append(m, elements...)

	args := toArgs(m...)

	_, err := conn.Do("LPUSHX", args...)

	return err
}

func (l *list) RPushX(key string, elements ...any) error {
	conn := l.redisPool.getConn()
	defer l.redisPool.CloseConn(conn)

	m := make([]any, 0)
	m = append(m, key)
	m = append(m, elements...)

	args := toArgs(m...)

	_, err := conn.Do("RPUSHX", args...)

	return err
}

func (l *list) LRange(key string, start, stop int64) ([]string, error) {
	conn := l.redisPool.getConn()
	defer l.redisPool.CloseConn(conn)

	return redis.Strings(conn.Do("LRANGE", key, start, stop))
}

func (l *list) LRem(key string, count int64, element any) error {
	conn := l.redisPool.getConn()
	defer l.redisPool.CloseConn(conn)

	_, err := conn.Do("LREM", key, count, element)

	return err
}

func (l *list) LSet(key string, index int64, element any) error {
	conn := l.redisPool.getConn()
	defer l.redisPool.CloseConn(conn)

	_, err := conn.Do("LSET", key, index, element)

	return err
}

func (l *list) LTrim(key string, start, stop int64) error {
	conn := l.redisPool.getConn()
	defer l.redisPool.CloseConn(conn)

	_, err := conn.Do("LTRIM", key, start, stop)

	return err
}
