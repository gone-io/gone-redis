package adapter

import (
	"github.com/gomodule/redigo/redis"
	"github.com/gone-io/gone"
	"github.com/gone-io/gone/goner/logrus"
	"sync"
)

func NewRedisPool() (gone.Angel, gone.GonerId) {
	return &RedisPool{}, IdGoneRedisPool
}

type RedisPool struct {
	gone.Flag
	logrus.Logger `gone:"gone-logger"`
	pool          *redis.Pool
	keyPrefix     string `gone:"config,redis.key.prefix"`

	server    string `gone:"config,redis.server"`
	password  string `gone:"config,redis.password"`
	maxIdle   int    `gone:"config,redis.max-idle,default=2"`
	maxActive int    `gone:"config,redis.max-active,default=10"`
	dbIndex   int    `gone:"config,redis.db,default=0"`

	once sync.Once
}

func (r *RedisPool) connect() {
	r.once.Do(func() {
		r.pool = &redis.Pool{
			MaxIdle:   r.maxIdle,   /*最大的空闲连接数*/
			MaxActive: r.maxActive, /*最大的激活连接数*/
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial(
					"tcp",
					r.server,
					redis.DialPassword(r.password),
					redis.DialDatabase(r.dbIndex),
				)
				if err != nil {
					return nil, err
				}
				return c, nil
			},
		}

		_, err := r.pool.Get().Do("ping")
		if err != nil {
			panic(err)
		}
	})
}

func (r *RedisPool) Start(gone.Cemetery) error {
	r.connect()
	return nil
}

func (r *RedisPool) Stop(gone.Cemetery) error {
	return r.pool.Close()
}

func (r *RedisPool) getConn() redis.Conn {
	r.connect()
	return r.pool.Get()
}

func (r *RedisPool) CloseConn(conn redis.Conn) {
	err := conn.Close()
	if err != nil {
		r.Logger.Errorf("close redis connection failed. err: %s", err.Error())
	}
}
