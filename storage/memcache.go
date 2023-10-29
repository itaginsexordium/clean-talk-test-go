package storage

import (
    "github.com/bradfitz/gomemcache/memcache"
)

type MemcacheClient struct {
    client *memcache.Client
}

func NewMemcacheClient(servers []string) *MemcacheClient {
    return &MemcacheClient{
        client: memcache.New(servers...),
    }
}

func (mc *MemcacheClient) Set(key string, value []byte) error {
    return mc.client.Set(&memcache.Item{Key: key, Value: value , Expiration: 86400,})
}

func (mc *MemcacheClient) Get(key string) ([]byte, error) {
    item, err := mc.client.Get(key)
    if err != nil {
        return nil, err
    }
    return item.Value, nil
}