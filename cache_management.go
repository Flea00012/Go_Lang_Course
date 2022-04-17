package main

import (
	"fmt"
	"sync"
	"time"
)

type CacheObject struct {
	Value      string
	TimeToLive int64
}

func (cacheObject CacheObject) IfExpired() bool {
	if cacheObject.TimeToLive == 0 {
		return false
	}
	return time.Now().UnixNano() > cacheObject.TimeToLive
}

type Cache struct {
	objects map[string]CacheObject
	mutex   *sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		objects: make(map[string]CacheObject),
		mutex:   &sync.RWMutex{},
	}
}

func (cache Cache) GetObject(cacheKey string) string {
	cache.mutex.RLock()
	defer cache.mutex.RUnlock()
	var object CacheObject
	object = cache.objects[cacheKey]
	if object.IfExpired() {
		delete(cache.objects, cacheKey)
		return ""
	}
	return object.Value
}

func (cache Cache) SetValue(cacheKey string, cacheValue string, timeToLive time.Duration) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	cache.objects[cacheKey] = CacheObject{
		Value:      cacheValue,
		TimeToLive: time.Now().Add(timeToLive).UnixNano(),
	}
}

func main()  {
	var cache *Cache
	cache = NewCache()
	cache.SetValue("name", "jane doe", 200000)
	var name string
	name = cache.GetObject("name")
	fmt.Printf(name)
}