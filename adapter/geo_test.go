package adapter

import (
	"fmt"
	"github.com/gone-io/gone"
	"github.com/stretchr/testify/assert"
	"testing"
)

type geoTester struct {
	gone.Flag
	g IGeo `gone:"gone-redis-geo, test-geo"`
}

func TestNewGeo(t *testing.T) {
	gone.Test(func(u *geoTester) {
		g := u.g
		t.Run("t-Geoadd & geoDist & geoRadius", func(t *testing.T) {
			gKey := "geoKey"

			err := g.GEOAdd(gKey, "122.2322113", "34.11123", "chengdu")
			assert.Nil(t, err)
			_ = g.GEOAdd(gKey, "142.2322113", "14.11123", "chongqing")
			_ = g.GEOAdd(gKey, "42.1123", "14.71123", "wuhan")
			_ = g.GEOAdd(gKey, "42.123", "14.71125", "changsha")

			res, err := g.GEORadius(gKey, "122.2322114", "33.123", 50550, "km")
			assert.Nil(t, err)
			fmt.Printf("%v", res)

			dist, err := g.GEODist(gKey, "chengdu", "chongqing", "m")
			assert.Nil(t, err)
			fmt.Printf("%f", dist)
		})

	}, func(cemetery gone.Cemetery) error {
		cemetery.Bury(&geoTester{})
		return nil
	}, Priest)
}
