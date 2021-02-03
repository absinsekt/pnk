package core

import (
	"sync"
	"time"
)

// Method generic cacheable function
type Method func() (interface{}, error)

type cached struct {
	CreatedAt  time.Time
	Expiration time.Duration
	Result     interface{}
}

var (
	store = sync.Map{}
)

// GetCached returns cached result of a func
func GetCached(key string, expiration time.Duration, method Method) (interface{}, error) {
	var (
		result interface{}
		err    error
	)

	if val, ok := store.Load(key); ok {
		go flushExpired()
		return val.(cached).Result, nil
	}

	result, err = method()

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
