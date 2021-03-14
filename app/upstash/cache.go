package facades

import (
	"github.com/go-redis/redis/v8"
	"log"
)

var redisInstance redis.UniversalClient

func getCache() redis.UniversalClient {
	if redisInstance == nil {
		option, err := redis.ParseURL("redis://:b45891f7c6844d5180e6621098d4bafd@us1-loyal-pipefish-32286.upstash.io:32286")
		if err != nil {
			log.Fatal(err)
		}

		redisInstance = redis.NewClient(option)
	}
	return redisInstance
}

func UseCache() redis.UniversalClient {
	return getCache()
}
