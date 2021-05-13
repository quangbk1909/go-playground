package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main() {
	fmt.Println("Go redis tutorial")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	user := User{
		Name: "Huy Quang",
		Age:  19,
	}

	byte, _ := json.Marshal(user)
	err := client.Set(context.Background(),"user", byte, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	val, err := client.Get(context.Background(),"user").Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(val)

}

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
}
