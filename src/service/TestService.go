package service

import (
	studentMapper "../data/mysql/mapper/student"
	redisMapper "../data/redis/mapper"
	"fmt"
)

func Test() {
	aaa := studentMapper.GetById(5)
	fmt.Println(aaa)
}

func RedisTest() {
	redisMapper.Set("abc", "123")
	redisMapper.Set("efg", "456")
	redisMapper.Set("abc", "789")
	a := redisMapper.Get("abc")
	b := redisMapper.Get("efg")
	fmt.Println(a)
	fmt.Println(b)
}
