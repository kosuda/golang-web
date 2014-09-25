package db

import (
	"github.com/garyburd/redigo/redis"
	"github.com/kosuda/golang-web/config"
	"log"
	"strconv"
)

var _conn redis.Conn

func init() {
	c := config.Configuration.Redis
	if conn, err := redis.Dial("tcp", c.Host+":"+strconv.Itoa(c.Port)); err != nil {
		log.Fatal(err.Error())

	} else {
		log.Println("Connected to redis store...")
		_conn = conn

	}

}

// Read all records
func Read(key string, id string) (string, error) {
	res, err := redis.String(_conn.Do("HGET", key, id))

	if err != nil {
		return "", err
	}

	return res, nil

}

// ReadAll function
func ReadAll(key string) ([]interface{}, error) {
	res, err := redis.Values(_conn.Do("HGETALL", key))
	// if err = redis.ScanStruct(res, data); err != nil {
	// 	return err
	// }
	return res, err
}

// Write function set records
func Write(key string, id string, data interface{}) error {
	// if _, err := _conn.Do("HMSET", redis.Args{}.Add(key).AddFlat(data)...); err != nil {
	// 	return err
	// }
	if _, err := _conn.Do("HSET", key, id, data); err != nil {
		return err
	}

	return nil

}

// Clear function implements flushall command.
func Clear() {
	_, err := _conn.Do("FLUSHALL")

	if err != nil {
		log.Fatal(err.Error())
	}

}

// Empty check function
func Empty() bool {
	keys, err := redis.Strings(_conn.Do("keys", "*"))

	if err != nil {
		log.Fatal(err.Error())
	}

	return len(keys) == 0

}
