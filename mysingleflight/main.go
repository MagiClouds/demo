package main

import (
	"golang.org/x/sync/singleflight"
	"log"
	"time"
)

func main() {

	var singleSetCache singleflight.Group

	getAndSetCache := func(requestID int, cacheKey string) (string, error) {

		log.Printf("request %v start to get and set cache...", requestID)

		value, _, _ := singleSetCache.Do(cacheKey, func() (ret interface{}, err error) {
			//do的入参key，可以直接使用缓存的key，这样同一个缓存，只有一个协程会去读DB

			log.Printf("request %v is setting cache...", requestID)

			time.Sleep(3 * time.Second)

			log.Printf("request %v set cache success!", requestID)

			return "VALUE", nil

		})



		return value.(string), nil

	}

	cacheKey := "cacheKey"

	for i := 1; i < 10; i++ { //模拟多个协程同时请求

		go func(requestID int) {

			value, _ := getAndSetCache(requestID, cacheKey)

			log.Printf("request %v get value: %v", requestID, value)

		}(i)

	}

	time.Sleep(5 * time.Second)

	value, _ := getAndSetCache(11, cacheKey)

	log.Printf("request %v get value: %v", 11, value)

}

//
//var sf singleflight.Group
//
//func OnceInRedis(key string, expiration time.Duration, f func() (interface{}, error), vo interface{}) error {
//	do, err, _ := sf.Do(key, func() (i interface{}, err error) {
//		result, err := redisCache.Get(key).Bytes()
//		if err != nil && err != redis.Nil {
//			return nil, err
//		}
//
//		if err != redis.Nil {
//			return result, nil
//		}
//
//		data, err := f()
//		if err != nil {
//			return nil, err
//		}
//
//		result, _ = json.Marshal(data)
//
//		if err = redisCache.Set(key, result, expiration).Err(); err != nil {
//			return nil, err
//		}
//
//		return result, nil
//	})
//	if err != nil {
//		return err
//	}
//
//	return json.Unmarshal(do.([]byte), vo)
//}
