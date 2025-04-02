package main

import (
	"context"
	"fmt"

	"github.com/my-valkey/api"
)

func main() {
	client, err := api.NewValkeyClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	ctx := context.Background()
	key := "test_key"
	value := "test_value"
	if err = client.Set(ctx, key, value); err != nil {
		panic(err)
	}

	fmt.Printf("SET: key=%s, value=%s\n", key, value)

	res, err := client.Get(ctx, key)
	if err != nil {
		panic(err)
	}

	fmt.Printf("GET: key=%s, value=%s\n", key, res)
}
