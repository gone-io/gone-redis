package adapter

import (
	"github.com/gomodule/redigo/redis"
	"github.com/gone-io/gone"
)

func NewStr() (gone.Angel, gone.GonerId) {
	return &str{}, IdRedisStr
}

type str struct {
	gone.Flag
	redisPool *RedisPool `gone:"gone-redis-pool"`
}

func (s *str) Start(gone.Cemetery) error {
	return nil
}

func (s *str) Stop(gone.Cemetery) error {
	return nil
}

func (s *str) Append(key string, value any) error {
	conn := s.redisPool.getConn()
	defer s.redisPool.CloseConn(conn)

	_, err := conn.Do("APPEND", key, value)

	return err
}

func (s *str) Decr(key string) error {
	conn := s.redisPool.getConn()
	defer s.redisPool.CloseConn(conn)

	_, err := conn.Do("DECR", key)

	return err
}

func (s *str) DecrBy(key string, decr int64) error {
	conn := s.redisPool.getConn()
	defer s.redisPool.CloseConn(conn)

	_, err := conn.Do("DECRBY", key, decr)

	return err
}

func (s *str) Get(key string) (string, error) {
	conn := s.redisPool.getConn()
	defer s.redisPool.CloseConn(conn)

	return redis.String(conn.Do("GET", key))
}

func (s *str) GetRange(key string, start, end int64) (string, error) {
	conn := s.redisPool.getConn()
	defer s.redisPool.CloseConn(conn)

	return redis.String(conn.Do("GETRANGE", key, start, end))
}

func (s *str) GetSet(key string, value any) (string, error) {
	conn := s.redisPool.getConn()
	defer s.redisPool.CloseConn(conn)

	return redis.String(conn.Do("GETSET", key, value))
}

func (s *str) Incr(key string) error {
	conn := s.redisPool.getConn()
	defer s.redisPool.CloseConn(conn)

	_, err := conn.Do("INCR", key)

	return err
}

func (s *str) IncrBy(key string, incr int64) error {
	conn := s.redisPool.getConn()
	defer s.redisPool.CloseConn(conn)

	_, err := conn.Do("INCRBY", key, incr)

	return err
}

func (s *str) IncrByFloat(key string, incrF float64) error {
	conn := s.redisPool.getConn()
	defer s.redisPool.CloseConn(conn)

	_, err := conn.Do("INCRBYFLOAT", key, incrF)

	return err
}

func (s *str) MGet(keys ...string) ([]string, error) {
	conn := s.redisPool.getConn()
	defer s.redisPool.CloseConn(conn)

	args := arrToArgs(keys)
	return redis.Strings(conn.Do("MGET", args...))
}

func (s *str) MSet(kvs map[string]any) error {
	conn := s.redisPool.getConn()
	defer s.redisPool.CloseConn(conn)

	args := mapToArgs(kvs)
	_, err := conn.Do("MSET", args...)

	return err
}

func (s *str) MSetNX(kvs map[string]any) (ok bool, err error) {
	conn := s.redisPool.getConn()
	defer s.redisPool.CloseConn(conn)

	args := mapToArgs(kvs)
	ok, err = redis.Bool(conn.Do("MSETNX", args...))

	return
}

func (s *str) Set(key string, value any) error {
	conn := s.redisPool.getConn()
	defer s.redisPool.CloseConn(conn)

	_, err := conn.Do("SET", key, value)

	return err
}

func (s *str) SetEX(key string, seconds int64, value any) error {
	conn := s.redisPool.getConn()
	defer s.redisPool.CloseConn(conn)

	_, err := conn.Do("SETEX", key, seconds, value)

	return err
}

func (s *str) SetNX(key string, value any) (ok bool, err error) {
	conn := s.redisPool.getConn()
	defer s.redisPool.CloseConn(conn)

	ok, err = redis.Bool(conn.Do("SETNX", key, value))

	return
}

func (s *str) SetRange(key string, offset int64, value any) error {
	conn := s.redisPool.getConn()
	defer s.redisPool.CloseConn(conn)

	_, err := conn.Do("SETRANGE", key, offset, value)

	return err
}

func (s *str) StrLen(key string) (int64, error) {
	conn := s.redisPool.getConn()
	defer s.redisPool.CloseConn(conn)

	return redis.Int64(conn.Do("STRLEN", key))
}
