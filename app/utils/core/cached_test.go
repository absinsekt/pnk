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
		key       string
		frequency time.Duration
		method    Method
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{name: "return from func", args: args{key: "one", frequency: time.Second * 2, method: fn1}, want: 1},
		{name: "return from cache", args: args{key: "one", frequency: time.Second, method: fn2}, want: 1},
		{name: "return from func", args: args{key: "one", frequency: time.Second, method: fn2}, want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := GetCached(tt.args.key, tt.args.frequency, tt.args.method); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCached() = %v, want %v", got, tt.want)
			}
		})

		time.Sleep(time.Second)
	}
}
