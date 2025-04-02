package api

import (
	"context"

	"github.com/valkey-io/valkey-go"
)

type ValkeyClient interface {
	Get(key string) (string, error)
	Set(key string, value string) error
	Close() error
}

type valkeyClient struct {
	client valkey.Client
}

func NewValkeyClient() (*valkeyClient, error) {
	c, err := valkey.NewClient(valkey.ClientOption{
		InitAddress: []string{"127.0.0.1:6379"},
	})
	if err != nil {
		return nil, err
	}
	return &valkeyClient{
		client: c,
	}, nil
}

func (vc *valkeyClient) Get(ctx context.Context, key string) (string, error) {
	return vc.client.Do(ctx, vc.client.B().Get().Key(key).Build()).ToString()
}

func (vc *valkeyClient) Set(ctx context.Context, key string, value string) error {
	return vc.client.Do(ctx, vc.client.B().Set().Key(key).Value(value).Build()).Error()
}

func (vc *valkeyClient) Close() {
	vc.client.Close()
}
