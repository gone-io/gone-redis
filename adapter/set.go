package adapter

import (
	"github.com/gomodule/redigo/redis"
	"github.com/gone-io/gone"
)

func NewSet() (gone.Angel, gone.GonerId) {
	return &set{}, IdRedisSet
}

type set struct {
	gone.Flag
	redisPool *RedisPool `gone:"gone-redis-pool"`
}

func (s set) Start(cemetery gone.Cemetery) error {
	//do nothing
	return nil
}

func (s set) Stop(cemetery gone.Cemetery) error {
	//do nothing
	return nil
}

func (s set) SAdd(key string, value ...any) error {
	conn := s.redisPool.getConn()
	defer s.redisPool.CloseConn(conn)

	args := redis.Args{}.Add(key).Add(value...)
	_, err := conn.Do("SADD", args...)

	return err
}

func (s set) SCard(key string) (int64, error) {
	conn := s.redisPool.getConn()
	defer s.redisPool.CloseConn(conn)

	return redis.Int64(conn.Do("SCARD", key))
}

func (s set) SDiff(baseKey string, keys ...string) ([]string, error) {
	conn := s.redisPool.getConn()
	defer s.redisPool.CloseConn(conn)

	args := redis.Args{}.Add(baseKey).Add(arrToArgs(keys)...)

	return redis.Strings(conn.Do("SDIFF", args...))
}

func (s set) SDiffStore(desKey, baseKey string, keys ...string) error {
	conn := s.redisPool.getConn()
	defer s.redisPool.CloseConn(conn)

	args := redis.Args{}.Add(desKey).Add(baseKey).Add(arrToArgs(keys)...)

	_, err := conn.Do("SDIFFSTORE", args...)

	return err
}

func (s set) SInter(keys ...string) ([]string, error) {
	conn := s.redisPool.getConn()
	defer s.redisPool.CloseConn(conn)

	args := redis.Args{}.Add(arrToArgs(keys)...)

	return redis.Strings(conn.Do("SINTER", args...))
}

func (s set) SInterStore(desKey string, keys ...string) error {
	conn := s.redisPool.getConn()
	defer s.redisPool.CloseConn(conn)

	args := redis.Args{}.Add(desKey).Add(arrToArgs(keys)...)

	_, err := conn.Do("SINTERSTORE", args...)

	return err
}

func (s set) SIsMember(key string, mem any) (bool, error) {
	conn := s.redisPool.getConn()
	defer s.redisPool.CloseConn(conn)

	return redis.Bool(conn.Do("SISMEMBER", key, mem))
}

func (s set) SMembers(key string) ([]string, error) {
	conn := s.redisPool.getConn()
	defer s.redisPool.CloseConn(conn)

	return redis.Strings(conn.Do("SMEMBERS", key))
}

func (s set) SMove(srcKey, desKey string, member any) error {
	conn := s.redisPool.getConn()
	defer s.redisPool.CloseConn(conn)

	_, err := conn.Do("SMOVE", srcKey, desKey, member)
	return err
}

func (s set) SPop(key string, count int64) ([]string, error) {
	conn := s.redisPool.getConn()
	defer s.redisPool.CloseConn(conn)

	return redis.Strings(conn.Do("SPOP", key, count))
}

func (s set) SRandMember(key string, count int64) ([]string, error) {
	conn := s.redisPool.getConn()
	defer s.redisPool.CloseConn(conn)

	return redis.Strings(conn.Do("SRANDMEMBER", key, count))
}

func (s set) SRem(key string, members ...any) error {
	conn := s.redisPool.getConn()
	defer s.redisPool.CloseConn(conn)

	args := redis.Args{}.Add(key).Add(members...)
	_, err := conn.Do("SREM", args...)

	return err
}

func (s set) SUnion(keys ...string) ([]string, error) {
	conn := s.redisPool.getConn()
	defer s.redisPool.CloseConn(conn)

	args := redis.Args{}.Add(arrToArgs(keys)...)

	return redis.Strings(conn.Do("SUNION", args...))
}

func (s set) SUnionStore(desKey string, keys ...string) error {
	conn := s.redisPool.getConn()
	defer s.redisPool.CloseConn(conn)

	args := redis.Args{}.Add(desKey).Add(arrToArgs(keys))
	_, err := conn.Do("SUNIONSTORE", args...)

	return err
}
