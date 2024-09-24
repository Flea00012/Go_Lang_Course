package algos

import (
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCache(t *testing.T) {
	var wg sync.WaitGroup
	var cache *Cache
	cache = NewCache()

	for i := 0; i <= 6; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s := "name-" + strconv.Itoa(i)
			cache.SetValue(s, "john smith-"+strconv.Itoa(i), 200000000)
		}()
	}

	wg.Wait()

	var name string
	name = cache.GetObject("name-1")
	assert.Equal(t, "john smith-1", name)
	name = cache.GetObject("name-6")
	assert.Equal(t, "john smith-6", name)
}
