package number

import (
	"fmt"
	error2 "service/internal/pkg/error"
	"time"

	xtremepkg "github.com/globalxtreme/go-core/v2/pkg"
	"github.com/gomodule/redigo/redis"
)

type ItemNumber struct{}

func (in ItemNumber) Item() string {
	return in.generate("ITM")
}

/** --- Unexported Functions --- */

func (in ItemNumber) generate(prefix string) string {
	date := time.Now().Format("20060102")
	key := fmt.Sprintf("number_counter:%s:%s", prefix, date)

	xtremepkg.InitRedisPool()
	conn := xtremepkg.RedisPool.Get()
	defer conn.Close()

	seq, err := redis.Int64(conn.Do("INCR", key))
	if err != nil {
		error2.ErrXtremeItemNumber(err.Error())
	}

	conn.Do("EXPIRE", key, 86400)

	return fmt.Sprintf("%s%s%07d", prefix, date, seq)
}
