package adapter

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func mapToArgs(kvs map[string]any) redis.Args {
	args := redis.Args{}
	for k, v := range kvs {
		args = args.Add(k).Add(fmt.Sprintf("%v", v))
	}

	return args
}

func toArgs(ks ...any) redis.Args {
	args := redis.Args{}
	for _, k := range ks {
		args = args.Add(k)
	}

	return args
}

func arrToArgs(ks []string) redis.Args {
	args := redis.Args{}
	for _, k := range ks {
		args = args.Add(k)
	}

	return args
}
