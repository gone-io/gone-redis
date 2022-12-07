package adapter

import (
	"fmt"
	"github.com/gone-io/gone"
	"github.com/stretchr/testify/assert"
	"testing"
)

type keyOpTester struct {
	gone.Flag
	k IKey `gone:"gone-redis-key,test-key"`
}

func TestKeyOp(t *testing.T) {
	gone.Test(func(u *keyOpTester) {
		k := u.k

		t.Run("t", func(t *testing.T) {

			keys, err := k.Keys("*")
			assert.Nil(t, err)
			fmt.Printf("%v \n", keys)

			existKey := "zkey-two"
			err = k.Expire("zkey-two", 1000)
			assert.Nil(t, err)

			err = k.ExpireAt(existKey, 1670425206)
			assert.Nil(t, err)

			err = k.Del(existKey)
			assert.Nil(t, err)
		})

	}, func(cemetery gone.Cemetery) error {
		cemetery.Bury(&keyOpTester{})
		return nil
	}, Priest)
}
