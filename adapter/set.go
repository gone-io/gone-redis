package adapter

import "github.com/gone-io/gone"

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

	return nil
}

func (s set) SCard(key string) (int64, error) {
	return 0, nil
}

func (s set) SDiff(baseKey string, keys ...string) ([]string, error) {
	return nil, nil
}

func (s set) SDiffStore(desKey, baseKey string, keys ...string) error {
	return nil
}

func (s set) SInter(keys ...string) ([]string, error) {
	return nil, nil
}

func (s set) SInterStore(desKey string, keys ...string) error {
	return nil
}

func (s set) SIsMember(key string, mem any) (bool, error) {
	return false, nil
}

func (s set) SMembers(key string) ([]string, error) {
	return nil, nil
}

func (s set) SMove(srcKey, desKey string, member any) error {
	return nil
}

func (s set) SPop(key string, count int64) ([]string, error) {
	return nil, nil
}

func (s set) SRandMember(key string, count int64) ([]string, error) {
	return nil, nil
}

func (s set) SRem(key string, members ...any) error {
	return nil
}

func (s set) SUnion(keys ...string) ([]string, error) {
	return nil, nil
}

func (s set) SUnionStore(desKey string, keys ...string) error {
	return nil
}
