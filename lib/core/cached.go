package core

import (
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
	store = make(map[string]cached)
)

// GetCached returns cached result of a func
func GetCached(key string, expiration time.Duration, method Method) (interface{}, error) {
	var (
		result interface{}
		err    error
	)

	go flushExpired()

	if store[key].Result == nil {
		result, err = method()

		if err == nil {
			store[key] = cached{
				CreatedAt:  time.Now(),
				Expiration: expiration,
				Result:     result,
			}

			return store[key].Result, nil
		}

		return nil, err
	}

	return store[key].Result, nil
}

func flushExpired() {
	for key, value := range store {
		isOutdated := value.CreatedAt.Add(store[key].Expiration).Before(time.Now())

		if isOutdated {
			delete(store, key)
		}
	}
}
