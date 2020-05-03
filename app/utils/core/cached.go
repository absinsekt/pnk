package core

import "time"

type Method func() (interface{}, error)

type cached struct {
	Frequency    time.Duration
	Result       interface{}
	LastModified time.Time
}

var (
	store = make(map[string]cached)
)

// GetCached returns cached result of a func
func GetCached(key string, frequency time.Duration, method Method) (interface{}, error) {
	var (
		result interface{}
		err    error
	)

	isOutdated := store[key].LastModified.Add(store[key].Frequency).Before(time.Now())

	if store[key].Result == nil || isOutdated {
		result, err = method()

		if err == nil {
			store[key] = cached{
				Frequency:    frequency,
				Result:       result,
				LastModified: time.Now(),
			}

			return store[key].Result, nil
		} else {
			return nil, err
		}
	}

	return store[key].Result, nil
}
