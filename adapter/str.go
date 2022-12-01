package adapter

import "github.com/gone-io/gone"

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
	return nil
}

func (s *str) DecrBy(key string, decr int64) error {
	return nil
}

func (s *str) Get(key string) (string, error) {
	return "", nil
}

func (s *str) GetRange(key string, start, end int64) (string, error) {
	return "", nil
}

func (s *str) GetSet(key string, value any) (string, error) {
	return "", nil
}

func (s *str) Incr(key string) error {
	return nil
}

func (s *str) IncrBy(key string, incr int64) error {
	return nil
}

func (s *str) IncrByFloat(key string, incrF float64) error {
	return nil
}

func (s *str) MGet(keys ...string) ([]string, error) {
	return nil, nil
}

func (s *str) MSet(kvs map[string]any) error {
	return nil
}

func (s *str) MSetNX(kvs map[string]any) error {
	return nil
}

func (s *str) Set(key string, value any) error {
	return nil
}

func (s *str) SetEX(key string, seconds int64, value any) error {
	return nil
}

func (s *str) SetNX(key string, value any) error {
	return nil
}

func (s *str) SetRange(key string, offset int64, value any) error {
	return nil
}

func (s *str) StrLen(key string) (int64, error) {
	return 0, nil
}
