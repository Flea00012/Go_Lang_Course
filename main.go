package main

import (
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


