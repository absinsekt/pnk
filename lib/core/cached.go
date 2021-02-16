package core

import (
	"sync"
	"time"
)

// CacheMethod generic cacheable function
type CacheMethod func() (interface{}, error)

type cached struct {
	CreatedAt  time.Time
	Expiration time.Duration
	Result     interface{}
}

var (
	store = sync.Map{}
)

// GetCached returns cached result of a func
func GetCached(key string, expiration time.Duration, cache CacheMethod) (interface{}, error) {
	var (
		result interface{}
		err    error
	)

	if Config.CacheEnabled {
		if val, ok := store.Load(key); ok {
			go flushExpired()
			return val.(cached).Result, nil
		}
	}

	result, err = cache()

	if err == nil {
		store.Store(key, cached{
			CreatedAt:  time.Now(),
			Expiration: expiration,
			Result:     result,
		})

		return result, nil
	}

	return nil, err
}

func flushExpired() {
	store.Range(func(key, value interface{}) bool {
		val := value.(cached)
		isOutdated := val.CreatedAt.Add(val.Expiration).Before(time.Now())

		if isOutdated {
			store.Delete(key)
		}

		return true
	})
}
