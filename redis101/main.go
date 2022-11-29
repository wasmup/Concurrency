package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	var ctx = context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := client.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get(ctx, "key").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("key2", val2)
	}

	key := "key1:wait"

	val, err = client.Get(ctx, key).Result()
	if err == redis.Nil {
		fmt.Println(key + " does not exist")
	} else if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(key, val)
	}

	expiration := 2 * time.Second
	b, e := client.SetNX(ctx, key, true, expiration).Result()
	fmt.Println(b, e)

	val, err = client.Get(ctx, key).Result()
	if err == redis.Nil {
		fmt.Println(key + " does not exist")
	} else if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(key, val)
	}

	b, e = client.SetNX(ctx, key, false, expiration).Result()
	fmt.Println(b, e)

	val, err = client.Get(ctx, key).Result()
	if err == redis.Nil {
		fmt.Println(key + " does not exist")
	} else if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(key, val)
	}

}
