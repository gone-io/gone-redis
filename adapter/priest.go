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

	if nil == cemetery.GetTomById(IdRedisList) {
		cemetery.Bury(NewList())
	}

	if nil == cemetery.GetTomById(IdRedisHash) {
		cemetery.Bury(NewHash())
	}

	if nil == cemetery.GetTomById(IdRedisSet) {
		cemetery.Bury(NewSet())
	}

	if nil == cemetery.GetTomById(IdRedisGeo) {
		cemetery.Bury(NewGeo())
	}

	if nil == cemetery.GetTomById(IdRedisZSet) {
		cemetery.Bury(NewZSet())
	}

	return nil
}
