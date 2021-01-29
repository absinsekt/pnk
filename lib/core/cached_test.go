package core

import (
	"reflect"
	"testing"
	"time"
)

var (
	fn1 = func() (interface{}, error) {
		return 1, nil
	}

	fn2 = func() (interface{}, error) {
		return 2, nil
	}
)

func TestGetCached(t *testing.T) {
	type args struct {
		key        string
		expiration time.Duration
		method     Method
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{name: "return from func1", args: args{key: "one", expiration: 2 * time.Second, method: fn1}, want: 1},
		{name: "return from cache1", args: args{key: "one", expiration: time.Second, method: fn2}, want: 1},
		{name: "return from func2", args: args{key: "one", expiration: 2 * time.Second, method: fn2}, want: 2},
		{name: "return from cache2", args: args{key: "one", expiration: time.Second, method: fn1}, want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := GetCached(tt.args.key, tt.args.expiration, tt.args.method); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCached() = %v, want %v", got, tt.want)
			}
		})

		time.Sleep(2 * time.Second)
	}
}
