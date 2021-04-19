package utility

import (
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/labstack/gommon/log"

	"api.inno/glo-profile-service/model"
)

type RedisPool struct {
	rp *redis.Pool
}

func InitPool(host string, port string, password string, database string, maxIdle int, maxActive int) *RedisPool {
	redisCfg := &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.DialURL("redis://" + password +
				"@" + host + ":" + port +
				"/" + database)
			log.Info("Dial redis at " + "redis://" + password +
				"@" + host + ":" + port +
				"/" + database)

			if err != nil {
				log.Printf("ERROR: fail init redis: %s", err.Error())
			}
			return conn, err
		},
	}

	return &RedisPool{
		rp: redisCfg,
	}
}

func (r *RedisPool) SetRedis(key string, val string) error {
	// get conn and put back when exit from method
	conn := r.rp.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, val)
	if err != nil {
		log.Printf("ERROR: fail set key %s, val %s, error %s", key, val, err.Error())
		return err
	}

	return nil
}

func (r *RedisPool) GetRedis(key string) (string, error) {
	// get conn and put back when exit from method
	conn := r.rp.Get()
	defer conn.Close()

	s, err := redis.String(conn.Do("GET", key))
	if err != nil {
		log.Printf("ERROR: fail get key %s, error %s", key, err.Error())
		return "", err
	}

	return s, nil
}

func (r *RedisPool) Sadd(key string, val string) error {
	// get conn and put back when exit from method
	conn := r.rp.Get()
	defer conn.Close()

	_, err := conn.Do("SADD", key, val)
	if err != nil {
		log.Printf("ERROR: fail add val %s to set %s, error %s", val, key, err.Error())
		return err
	}

	return nil
}

func (r *RedisPool) Smembers(key string) ([]string, error) {
	// get conn and put back when exit from method
	conn := r.rp.Get()
	defer conn.Close()

	s, err := redis.Strings(conn.Do("SMEMBERS", key))
	if err != nil {
		log.Printf("ERROR: fail get set %s , error %s", key, err.Error())
		return nil, err
	}

	return s, nil
}

func (r *RedisPool) HSetRedis(key string, val interface{}) error {

	conn := r.rp.Get()
	defer conn.Close()
	_, err := conn.Do("HSET", redis.Args{}.Add(key).AddFlat(val)...)
	if err != nil {
		log.Printf("ERROR: fail set key %s, val %s, error %s", key, val, err.Error())
		return err
	}
	return nil
}

func (r *RedisPool) HGetRedis(key string) (interface{}, error) {

	conn := r.rp.Get()
	defer conn.Close()
	values, err := redis.Values(conn.Do("HGETALL", key))
	if err != nil {
		log.Printf("ERROR: fail get key %s, error %s", key, err.Error())
		return "", err
	}
	p := model.Profile{}
	redis.ScanStruct(values, &p)
	return p, nil
}
