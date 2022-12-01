package adapter

import (
	"github.com/gone-io/gone"
	"github.com/gone-io/gone/goner"
	gone_redis "gone-redis-main"
)

func Priest(cemetery gone.Cemetery) error {
	_ = goner.BasePriest(cemetery)
	_ = gone_redis.Priest(cemetery)

	if nil == cemetery.GetTomById(IdGoneRedisPool) {
		cemetery.Bury(NewRedisPool())
	}

	if nil == cemetery.GetTomById(IdRedisStr) {
		cemetery.Bury(NewStr())
	}

	return nil
}