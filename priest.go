package gone_redis

import (
	"github.com/gone-io/gone"
	"github.com/gone-io/gone/goner"
)

func Priest(cemetery gone.Cemetery) error {
	_ = goner.BasePriest(cemetery)

	return nil
}
