package adapter

import (
	"github.com/gomodule/redigo/redis"
	"github.com/gone-io/gone"
)

func NewKeyOp() (gone.Angel, gone.GonerId) {
	return &keyOp{}, IdRedisKey
}

type keyOp struct {
	gone.Flag
	redisPool *RedisPool `gone:"gone-redis-pool"`
}

func (k *keyOp) Start(gone.Cemetery) error {
	return nil
}

func (k *keyOp) Stop(gone.Cemetery) error {
	return nil
}

func (k *keyOp) Keys(pattern string) ([]string, error) {
	conn := k.redisPool.getConn()
	defer k.redisPool.CloseConn(conn)

	return redis.Strings(conn.Do("KEYS", pattern))
}

func (k *keyOp) Expire(key string, sec int64) error {
	conn := k.redisPool.getConn()
	defer k.redisPool.CloseConn(conn)

	_, err := conn.Do("EXPIRE", key, sec)

	return err
}

func (k *keyOp) ExpireAt(key string, timestamp int64) error {
	conn := k.redisPool.getConn()
	defer k.redisPool.CloseConn(conn)

	_, err := conn.Do("EXPIREAT", key, timestamp)

	return err
}

func (k *keyOp) PExpire(key string, mill int64) error {
	conn := k.redisPool.getConn()
	defer k.redisPool.CloseConn(conn)

	_, err := conn.Do("PEXPIRE", key, mill)

	return err
}

func (k *keyOp) PExpireAt(key string, millTimestamp int64) error {
	conn := k.redisPool.getConn()
	defer k.redisPool.CloseConn(conn)

	_, err := conn.Do("PEXPIREAT", key, millTimestamp)

	return err
}

func (k *keyOp) Ttl(key string) (int64, error) {
	conn := k.redisPool.getConn()
	defer k.redisPool.CloseConn(conn)

	return redis.Int64(conn.Do("TTL", key))
}

func (k *keyOp) PTtl(key string) (int64, error) {
	conn := k.redisPool.getConn()
	defer k.redisPool.CloseConn(conn)

	return redis.Int64(conn.Do("PTTL", key))
}

func (k *keyOp) Del(keys ...string) error {
	conn := k.redisPool.getConn()
	defer k.redisPool.CloseConn(conn)

	args := redis.Args{}.Add(arrToArgs(keys)...)

	_, err := conn.Do("DEL", args...)

	return err
}
