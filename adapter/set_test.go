package adapter

import (
	"fmt"
	"github.com/gone-io/gone"
	"github.com/stretchr/testify/assert"
	"testing"
)

type setTester struct {
	gone.Flag
	s ISet `gone:"gone-redis-set, test-hash"`
}

func TestNewSet(t *testing.T) {
	gone.Test(func(u *setTester) {
		s := u.s

		t.Run("tSadd & sCard & sDiff & sDiffStore & sInter & sInterStore", func(t *testing.T) {
			setKey1 := "set-one"
			setKey2 := "set-two"
			setKey3 := "set-three"
			setKey4 := "set-four"
			setKey5 := "set-five"

			err := s.SAdd(setKey1, 1, "two", "3", 4, 5)
			assert.Nil(t, err)

			err = s.SAdd(setKey2, "one", "two", 3, "four", 5)
			assert.Nil(t, err)

			err = s.SAdd(setKey3, "one", "tw", 3, "four", 5)
			assert.Nil(t, err)

			num, err := s.SCard(setKey1)
			assert.Nil(t, err)
			assert.EqualValues(t, 5, num)

			diffs, err := s.SDiff(setKey1, setKey2, setKey3)
			assert.Nil(t, err)
			assert.EqualValues(t, 2, len(diffs))
			fmt.Printf("%v \n", diffs)

			err = s.SDiffStore(setKey4, setKey1, setKey2, setKey3)
			assert.Nil(t, err)

			num, err = s.SCard(setKey4)
			assert.EqualValues(t, 2, num)

			inter, err := s.SInter(setKey1, setKey2, setKey3)
			assert.Nil(t, err)
			assert.EqualValues(t, 2, len(inter))
			fmt.Printf("%v \n", inter)

			err = s.SInterStore(setKey1, setKey2, setKey5)
			assert.Nil(t, err)

			num, err = s.SCard(setKey4)
			assert.EqualValues(t, 2, num)
		})

		t.Run("t-Sismember & smembers & smove & spop & srandmember & srem & sunion & sunionstore", func(t *testing.T) {
			setKey6 := "set-six"
			setKey7 := "set-seven"
			setKey8 := "set-eight"

			err := s.SAdd(setKey6, 1, "two", "3", 4, 5)
			assert.Nil(t, err)

			err = s.SAdd(setKey7, "one", "two", 3, "four", 5)
			assert.Nil(t, err)

			isMember, err := s.SIsMember(setKey6, "two")
			assert.Nil(t, err)
			assert.EqualValues(t, isMember, true)

			isMember, err = s.SIsMember(setKey6, "2")
			assert.Nil(t, err)
			assert.EqualValues(t, isMember, false)

			members, err := s.SMembers(setKey6)
			assert.Nil(t, err)
			assert.EqualValues(t, 5, len(members))
			fmt.Printf("%v \n", members)

			err = s.SMove(setKey6, setKey7, "1")
			assert.Nil(t, err)

			len6, err := s.SCard(setKey6)
			assert.Nil(t, err)
			assert.EqualValues(t, 4, len6)
			len7, err := s.SCard(setKey7)
			assert.Nil(t, err)
			assert.EqualValues(t, 6, len7)

			pop, err := s.SPop(setKey7, 1)
			assert.Nil(t, err)
			len7, err = s.SCard(setKey7)
			assert.EqualValues(t, 5, len7)
			fmt.Printf("%v \n", pop)

			members, err = s.SRandMember(setKey7, 3)
			assert.Nil(t, err)
			assert.EqualValues(t, 3, len(members))
			fmt.Printf("%v \n", members)
			len7, err = s.SCard(setKey7)
			assert.EqualValues(t, 5, len7)

			err = s.SRem(setKey7, "two", "3")
			assert.Nil(t, err)
			len7, err = s.SCard(setKey7)
			assert.EqualValues(t, 3, len7)

			_, err = s.SUnion(setKey6, setKey7)
			assert.Nil(t, err)

			err = s.SUnionStore(setKey8, setKey6, setKey7)
			assert.Nil(t, err)
		})

	}, func(cemetery gone.Cemetery) error {
		cemetery.Bury(&setTester{})
		return nil
	}, Priest)
}
