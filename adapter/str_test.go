package adapter

import (
	"github.com/gone-io/gone"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

type RedisTester struct {
	gone.Flag
	s IStr `gone:"gone-redis-str,test-str"`
}

func TestHash(t *testing.T) {
	gone.Test(func(u *RedisTester) {
		s := u.s

		t.Run("append", func(t *testing.T) {
			n := rand.Intn(100)
			field := "point-a"
			err := s.Append(field, n)
			assert.Nil(t, err)

			a := "append content"
			err = s.Append(field, a)
			assert.Nil(t, err)
		})
	}, func(cemetery gone.Cemetery) error {
		cemetery.Bury(&RedisTester{})
		return nil
	}, Priest)
}
