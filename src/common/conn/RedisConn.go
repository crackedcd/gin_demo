package conn

import (
	"../config"
	"../utils/LoggerUtils"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var redisPoll chan redis.Conn

func RedisConn() redis.Conn {
	host, port, auth, db := config.GetRedisConfig()
	return initRedis(host, port, auth, db)
}

func putRedis(conn redis.Conn) {
	if redisPoll == nil {
		redisPoll = make(chan redis.Conn, maxPoolSize)
	}
	if len(redisPoll) >= maxPoolSize {
		_ = conn.Close()
		return
	}
	redisPoll <- conn
}

func initRedis(host string, port int, auth string, db int) redis.Conn {
	address := fmt.Sprintf("%s:%d", host, port)
	if len(redisPoll) == 0 {
		// 如果长度为0，就定义一个redis.Conn类型长度为maxPoolSize的channel
		redisPoll = make(chan redis.Conn, maxPoolSize)
		go func() {
			for i := 0; i < maxPoolSize/2; i++ {
				conn, connErr := redis.Dial("tcp", address)
				if conn != nil {
					authErr := conn.Send("auth", auth)
					LoggerUtils.Error(authErr)
					selectErr := conn.Send("select", db)
					LoggerUtils.Error(selectErr)
					putRedis(conn)
				} else {
					LoggerUtils.Error(connErr)
				}
			}
		}()
	}
	return <- redisPoll
}
