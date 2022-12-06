package adapter

import (
	"github.com/gomodule/redigo/redis"
	"github.com/gone-io/gone"
	"strconv"
)

func NewGeo() (gone.Angel, gone.GonerId) {
	return &geo{}, IdRedisGeo
}

type geo struct {
	gone.Flag
	redisPool *RedisPool `gone:"gone-redis-pool"`
}

func (g *geo) Start(cemetery gone.Cemetery) error {
	// do nothing
	return nil
}

func (g *geo) Stop(cemetery gone.Cemetery) error {
	// do nothing
	return nil
}

func (g *geo) GEOAdd(key, longitude, latitude string, member any) error {
	conn := g.redisPool.getConn()
	defer g.redisPool.CloseConn(conn)

	_, err := conn.Do("GEOADD", key, longitude, latitude, member)

	return err
}

func (g *geo) GEODist(key string, mem1, mem2 any, unit string) (float64, error) {
	conn := g.redisPool.getConn()
	defer g.redisPool.CloseConn(conn)

	return redis.Float64(conn.Do("GEODIST", key, mem1, mem2, unit))
}

func (g *geo) GEORadius(key, longitude, latitude string, radius float64, unit string) ([]GEOPos, error) {
	conn := g.redisPool.getConn()
	defer g.redisPool.CloseConn(conn)

	values, err := redis.Values(conn.Do("GEORADIUS", key, longitude, latitude, radius, unit, "WITHDIST", "WITHCOORD"))
	if err != nil {
		return nil, err
	}

	res := make([]GEOPos, 0)
	for _, val := range values {
		vals := val.([]interface{})
		var aGEOPos GEOPos
		for idx, v := range vals {
			switch idx {
			case 0:
				aGEOPos.Mem = string(v.([]byte))
			case 1:
				s := string(v.([]byte))
				aGEOPos.Distance, _ = strconv.ParseFloat(s, 64)
			case 2:
				minVal := v.([]interface{})
				if len(minVal) == 2 {
					long := string(minVal[0].([]byte))
					lat := string(minVal[1].([]byte))
					aGEOPos.Position.Longitude, _ = strconv.ParseFloat(long, 64)
					aGEOPos.Position.Latitude, _ = strconv.ParseFloat(lat, 64)
				}
			}
		}

		res = append(res, aGEOPos)
	}

	return res, nil
}
