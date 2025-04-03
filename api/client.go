package api

import (
	"context"

	"github.com/valkey-io/valkey-go"
)

type ValkeyClient interface {
	Get(key string) (string, error)
	Set(key string, value string) error
	SetWithExpiry(key string, value string, seconds int64) error
	Delete(key string) error
	Exists(key string) (bool, error)
	Expire(key string, seconds int64) error
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

func (vc *valkeyClient) SetWithExpiry(ctx context.Context, key string, value string, seconds int64) error {
	return vc.client.Do(ctx, vc.client.B().Set().Key(key).Value(value).ExSeconds(seconds).Build()).Error()
}

func (vc *valkeyClient) Delete(ctx context.Context, key string) error {
	return vc.client.Do(ctx, vc.client.B().Del().Key(key).Build()).Error()
}

func (vc *valkeyClient) Exists(ctx context.Context, key string) (bool, error) {
	return vc.client.Do(ctx, vc.client.B().Exists().Key(key).Build()).AsBool()
}

func (vc *valkeyClient) Expire(ctx context.Context, key string, seconds int64) error {
	return vc.client.Do(ctx, vc.client.B().Expire().Key(key).Seconds(seconds).Build()).Error()
}

func (vc *valkeyClient) Close() {
	vc.client.Close()
}
