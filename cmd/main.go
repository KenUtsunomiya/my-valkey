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

	exists, err := client.Exists(ctx, key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("EXISTS: key=%s, exists=%t\n", key, exists)

	if err = client.Set(ctx, key, value); err != nil {
		panic(err)
	}
	fmt.Printf("SET: key=%s, value=%s\n", key, value)

	exists, err = client.Exists(ctx, key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("EXISTS: key=%s, exists=%t\n", key, exists)

	res, err := client.Get(ctx, key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("GET: key=%s, value=%s\n", key, res)

	if err = client.Delete(ctx, key); err != nil {
		panic(err)
	}
	fmt.Printf("DELETE: key=%s\n", key)
}
