package main

import (
	"github.com/gomodule/redigo/redis"
)

func main() {
	p := &redis.Pool{
		MaxActive: 10,
		MaxIdle:   3,
		Dial: func() (redis.Conn, error) {
			d, err := redis.Dial("tcp", "192.168.101.7:6379", redis.DialPassword(""))
			if err != nil {
				return nil, err
			}
			return d, nil
		},
	}
	c := p.Get()
	defer func(c redis.Conn) {
		err := c.Close()
		if err != nil {

		}
	}(c)
}
