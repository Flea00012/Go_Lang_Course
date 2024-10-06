package algos

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
	queue   []CacheObject
	size    int
}

func NewCache(i int) *Cache {
	return &Cache{
		objects: make(map[string]CacheObject),
		mutex:   &sync.RWMutex{},
		queue:   make([]CacheObject, i),
		size:    i,
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
	cache.queue = append(cache.queue, object)
	return object.Value
}

func (cache Cache) SetValue(cacheKey string, cacheValue string, timeToLive time.Duration) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	for len(cache.objects) < cache.size {
		c := CacheObject{
			Value:      cacheValue,
			TimeToLive: time.Now().Add(timeToLive).UnixNano(),
		}
		cache.objects[cacheKey] = c
		cache.queue = append(cache.queue, c)
	}
}
