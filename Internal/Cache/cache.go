package Cache

import (
	"fmt"
	"time"
)

type Cache struct {
	duration       time.Duration
	cachedArea     map[string]cacheAreaEntry
	cachedPokemons map[string]cachePokemonsEntry
}

type cacheAreaEntry struct {
	entryTime time.Time
	data      []byte
}

type cachePokemonsEntry struct {
	entryTime time.Time
	data      []byte
}

func NewCache(t time.Duration) *Cache {
	return &Cache{duration: t, cachedArea: map[string]cacheAreaEntry{}, cachedPokemons: map[string]cachePokemonsEntry{}}
}

func (c *Cache) AddArea(s string, val []byte) {

	c.cachedArea[s] = cacheAreaEntry{
		entryTime: time.Now(),
		data:      val,
	}
	fmt.Println("saved areas to cache...")
}

func (c *Cache) AddPokemons(s string, val []byte) {

	c.cachedPokemons[s] = cachePokemonsEntry{
		entryTime: time.Now(),
		data:      val,
	}
	fmt.Println("saved pokemons to cache...")
}

func (c *Cache) GetArea(s string) ([]byte, bool) {
	val, ok := c.cachedArea[s]
	if !ok {
		return nil, false
	}
	return val.data, true

}

func (c *Cache) GetPokemons(s string) ([]byte, bool) {
	val, ok := c.cachedPokemons[s]
	if !ok {
		return nil, false
	}
	return val.data, true

}

func (c *Cache) StartReapingLoop() {
	ticker := time.NewTicker(c.duration)
	for range ticker.C {
		// will go forever in interval of ticker
		c.ReapCacheAreaEnties()
		c.ReapCachePokemonsEnties()
	}
}

func (c *Cache) ReapCacheAreaEnties() {
	lastStoringTime := time.Now().UTC().Add(-c.duration)
	for k, v := range c.cachedArea {
		if v.entryTime.Before(lastStoringTime) {
			// fmt.Printf("\n %v\n", "deleted")
			delete(c.cachedArea, k)
		}
	}
}

func (c *Cache) ReapCachePokemonsEnties() {
	lastStoringTime := time.Now().UTC().Add(-c.duration)
	for k, v := range c.cachedPokemons {
		if v.entryTime.Before(lastStoringTime) {
			// fmt.Printf("\n %v\n", "deleted")
			delete(c.cachedPokemons, k)
		}
	}
}
