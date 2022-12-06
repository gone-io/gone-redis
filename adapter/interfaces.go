package adapter

type ListPos string

const (
	Before ListPos = "BEFORE"
	After  ListPos = "AFTER"
)

type ZSetOrderStrategy string

const (
	ByScore ZSetOrderStrategy = "BYSCORE"
	ByLex   ZSetOrderStrategy = "BYLEX"
)

type IStr interface {

	//Append add value (as string) to the key, if key not exist, create it.
	Append(key string, value any) error

	//Decr Decrease 1 from value of the key, error will happen if value of the key is not a number.
	Decr(key string) error

	//DecrBy Decrease decr from value of the key, error will happen if value of the key is not a number.
	DecrBy(key string, decr int64) error

	//Get returns the value of the key.
	Get(key string) (string, error)

	//GetRange returns the substring of the string value stored at key.
	GetRange(key string, start, end int64) (string, error)

	//GetSet use new value to replace old value and return old value.
	GetSet(key string, value any) (string, error)

	//Incr increase 1 from value of the key, error will happen if value of the key is not a number.
	Incr(key string) error

	//IncrBy increase decr from value of the key, error will happen if value of the key is not a number.
	IncrBy(key string, incr int64) error

	//IncrByFloat Increment the string representing a floating point number stored at key by the specified increment.
	//By using a negative increment value, the result is that the value stored at the key is decremented.
	//[notice] Calculating floating point numbers is often inaccurate.
	IncrByFloat(key string, incrF float64) error

	//MGet returns the values of all specified keys. If key not exist, return nil.
	MGet(keys ...string) ([]string, error)

	//MSet sets the given keys to their respective values. the map means <key, value>.
	MSet(kvs map[string]any) error

	//MSetNX Sets the given keys to their respective values. MSETNX will not perform any operation at all even if just a single key already exists.
	MSetNX(kvs map[string]any) (bool, error)

	//ISet give the string value to the key， NX | XX and other operations is not supported. Use redigo if you want.
	Set(key string, value any) error

	//SetEX ISet key to hold the string value and set key to timeout after a given number of seconds.
	SetEX(key string, seconds int64, value any) error

	//SetNX ISet key to hold string value if key does not exist. In that case, it is equal to SET.  "SET if Not eXists".
	SetNX(key string, value any) (bool, error)

	//SetRange Overwrites part of the string stored at key, starting at the specified offset, for the entire length of value. if offset out of range, append value to the key.
	SetRange(key string, offset int64, value any) error

	//StrLen returns the length of the string value stored at key.
	StrLen(key string) (int64, error)
}

type IHash interface {

	//HGetAll both of the field and value will be return, index 0 is field1, index 1 is value1, etc.
	HGetAll(key string) ([]string, error)

	//HMGet returns a list of the fields got.
	HMGet(key string, fields ...any) ([]string, error)

	//HGet returns a list of the fields got.
	HGet(key string, field any) (string, error)

	HSet(key string, field string, value any) error

	//HSetNX it's ok when the field not exist.
	HSetNX(key string, field string, value any) (bool, error)

	//HMSet the given map means <field, value>.
	HMSet(key string, kvs map[string]any) error

	//HIncrBy increase the field value, if the key or field not exist, it will increase from zero and set the field-value of the key.
	HIncrBy(key string, field string, incr int64) error

	//HKeys returns all fields contained.
	HKeys(key string) ([]string, error)

	//HValues returns all values contained.
	HVals(key string) ([]string, error)

	//HLen returns the num of fields.
	HLen(key string) (int64, error)

	//HDel remove a field.
	HDel(key string, field string) error

	//HExists determine if the field exist.
	HExists(key string, field string) (bool, error)
}

type IList interface {

	//BLPop unit of timeout is second，From first element.
	//[NOTICE] The first element of the result is the list key witch popped ele, second is the value.
	BLPop(timeout int64, keys ...any) ([]string, error)

	//BRPop unit of timeout is second，From last element.
	//[NOTICE] The first element of the result is the list key witch popped ele, second is the value.
	BRPop(timeout int64, keys ...any) ([]string, error)

	//LIndex return the element at the index
	LIndex(key string, idx int64) (string, error)

	//LInsert if pivot not exists, err will be returned.
	LInsert(key string, pos ListPos, pivot, element any) error

	LLen(key string) (int64, error)

	//LPop remove first elements. if elements in list is not enough, no error will be returned.
	LPop(key string, count int64) ([]string, error)

	//RPop remove last elements. if elements in list is not enough, no error will be returned.
	RPop(key string, count int64) ([]string, error)

	//LPush push elements to head of the list.
	LPush(key string, elements ...any) error

	//RPush push elements to tail of the list.
	RPush(key string, elements ...any) error

	//LPushX it's ok only if key has already existed.
	LPushX(key string, elements ...any) error

	//RPushX it's ok only if key has already existed.
	RPushX(key string, elements ...any) error

	LRange(key string, start, stop int64) ([]string, error)

	//LRem count < 0, last -> first; count > 0 first -> last; count = 0, remove all elements equals element.
	LRem(key string, count int64, element any) error

	//LSet index out of range when index greater than len.
	LSet(key string, index int64, element any) error

	//LTrim remove elements out of [start, stop] index, not contains index of start and stop.
	LTrim(key string, start, stop int64) error
}

type ISet interface {
	SAdd(key string, value ...any) error

	//SCard returns num of members at the set.
	SCard(key string) (int64, error)

	//SDiff returns the members that baseKey contains and other keys not contain.
	SDiff(baseKey string, keys ...string) ([]string, error)

	//SDiffStore similar with SDiff, but the result will be stored at desKey. If desKey exist, overwrite it.
	SDiffStore(desKey, baseKey string, keys ...string) error

	//SInter returns the members of the set resulting from the intersection of all the given sets.
	SInter(keys ...string) ([]string, error)

	//SInterStore similar with SInter, but the result will be stored at desKey. If desKey exist, overwrite it.
	SInterStore(desKey string, keys ...string) error

	//SIsMember returns if member is a member of the set stored at key.
	SIsMember(key string, mem any) (bool, error)

	// SMembers returns all the members of the set value stored at key.
	SMembers(key string) ([]string, error)

	//SMove move member from srcKey to desKey. If srcKey or member not exist, do nothing; If descKey not exist, create it.
	SMove(srcKey, desKey string, member any) error

	//SPop removes and returns one or more random members from the set value store at key.
	SPop(key string, count int64) ([]string, error)

	//SRandMember returns a random element from the set value stored at key, not remove. When count > 0,
	//distinct elements returned, otherwise, elements returned not distinct.
	SRandMember(key string, count int64) ([]string, error)

	//SRem remove the members given.
	SRem(key string, members ...any) error

	//SUnion returns the members of the set resulting from the union of all the given sets.
	SUnion(keys ...string) ([]string, error)

	//SUnionStore stores the members of the set resulting from the union of all the given sets to desKey.
	SUnionStore(desKey string, keys ...string) error
}

type IZSet interface {

	//ZAdd some params like XX | NX, GT | LT, CH, INCR not support in this api, and score only support Integer, use redigo if you need them.
	ZAdd(key string, score int64, member any) error

	//ZCard returns the sorted set cardinality (number of elements) of the sorted set stored at key.
	ZCard(key string) (int64, error)

	//ZCount returns the number of elements in the sorted set at key with a score between min and max. if includeMin is false,
	//min -> "(min", if includeMax is false, max -> "(max".
	ZCount(key string, min, max int64, includeMin, includeMax bool) (int64, error)

	//ZIncrBy increments the score of member in the sorted set stored at key by increment. if key or member not exist,
	//a mew zset or member will be created and increase from 0.0.
	ZIncrBy(key string, incr int64, member any) error

	//ZPopMax removes and returns up to count members with the highest scores in the sorted set stored at key.  index 0 is member0, index 1 is score0, etc. [since 5.0.0]
	ZPopMax(key string, count int64) ([]string, error)

	//ZPopMin removes and returns up to count members with the lowest scores in the sorted set stored at key.  index 0 is member0, index 1 is score0, etc. [since 5.0.0]
	ZPopMin(key string, count int64) ([]string, error)

	//ZRangeWithLimit returns the specified range of elements in the sorted set stored at <key>, used [ZRANGE] command.
	ZRangeWithLimit(key string, start, end int64, orderStrategy ZSetOrderStrategy, offset, count int64, withScore bool) ([]string, error)

	//ZRangeWithoutLimit returns the specified range of elements in the sorted set stored at <key>, used [ZRANGE] command.
	ZRangeWithoutLimit(key string, start, end int64, orderStrategy ZSetOrderStrategy, withScore bool) ([]string, error)

	//ZRank returns the rank of member in the sorted set stored at key, with the scores ordered from low to high.
	ZRank(key string, member any) (int64, error)

	//ZRem removes the specified members from the sorted set stored at key. Not existing members are ignored.
	ZRem(key string, members ...any) error

	//ZRemRangeByRank removes all elements in the sorted set stored at key with rank between start and stop. if start and end GE 0, low -> high; else, from high to low
	ZRemRangeByRank(key string, start, end int64) error

	//ZRemRangeByScore removes all elements in the sorted set stored at key with a score between min and max (inclusive).
	ZRemRangeByScore(key string, min, max int64) error

	//ZRevRange returns the specified range of elements in the sorted set stored at key. The elements are considered to be ordered from the highest to the lowest score.
	ZRevRange(key string, start, stop int64, withScores bool) ([]string, error)

	//ZRevRangeByScoreWithLimit returns all the elements in the sorted set at key with a score between max and min
	//(including elements with score equal to max or min), and use offset and count paging. Used [ZREVRANGEBYSCORE] command
	ZRevRangeByScoreWithLimit(key string, max, min int64, withScores bool, offset, count int64)

	//ZRevRangeByScoreWithoutLimit returns all the elements in the sorted set at key with a score between max and min
	//(including elements with score equal to max or min). Used [ZREVRANGEBYSCORE] command
	ZRevRangeByScoreWithoutLimit(key string, max, min int64, withScores bool)

	//ZScore returns the score of member in the sorted set at key
	ZScore(key string, member any) (float64, error)
}

type IGeo interface {

	//in a way that makes it possible to query the items with the GEOSEARCH command.
	GEOAdd(key, longitude, latitude string, member any) error

	//GEODist Return the distance between two members in the geospatial index represented by the sorted set.
	// default unit is "M", unit => [M | KM | FT | MI]
	GEODist(key string, mem1, mem2 any, unit string) (float64, error)

	//GEORadius Return the members of a sorted set populated with geospatial information using GEOADD,
	//which are within the borders of the area specified with the center location and the maximum distance from the center (the radius).
	GEORadius(key, longitude, latitude string, radius float64, unit string) ([]GEOPos, error)
}
