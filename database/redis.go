package database

import (
	"github.com/garyburd/redigo/redis"
	"lab.castawaylabs.com/orderchef/util"
	"time"
)

// Redis pool
var Redis *redis.Pool

func init() {
	Redis = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 60 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialTimeout("tcp", util.Config.SessionDb, 100*time.Millisecond, 200*time.Millisecond, 800*time.Millisecond)
			if err != nil {
				return nil, err
			}

			return c, err
		},
		/*TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},*/
	}
}
