package ApiCalls

import (
	"net/http"
	"time"

	"github.com/stonoy/pokedex/Internal/Cache"
)

type Client struct {
	httpClient  http.Client
	cacheMemory *Cache.Cache
}

func NewClient(timeout time.Duration, cacheTimeout time.Duration) *Client {
	return &Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cacheMemory: Cache.NewCache(cacheTimeout),
	}
}

func (c *Client) RemoveCache() {
	c.cacheMemory.StartReapingLoop()
}
