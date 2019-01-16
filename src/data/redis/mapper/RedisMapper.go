package mapper

import (
	"../../../common/conn"
	"../../../common/utils/LoggerUtils"
	"github.com/garyburd/redigo/redis"
)

func Get(k string) string {
	c := conn.RedisConn()
	value, _ := redis.String(c.Do("GET", k))
	return value
}

func Set(k string, v string) {
	c := conn.RedisConn()
	_, err := c.Do("SET", k, v)
	LoggerUtils.Error(err)
}