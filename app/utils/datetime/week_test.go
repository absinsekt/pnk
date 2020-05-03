package datetime

import (
	"reflect"
	"testing"
	"time"
)

func TestGetWeekFirstDate(t *testing.T) {
	type args struct {
		currentDate time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "2020-04-06 — 2020-04-12",
			args: args{currentDate: time.Date(2020, 4, 12, 1, 22, 35, 0, time.UTC)},
			want: time.Date(2020, 4, 6, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "2020-03-16 — 2020-03-22",
			args: args{currentDate: time.Date(2020, 3, 17, 23, 59, 35, 0, time.UTC)},
			want: time.Date(2020, 3, 16, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "2020-04-13 — 2020-04-19",
			args: args{currentDate: time.Date(2020, 4, 19, 23, 59, 35, 0, time.UTC)},
			want: time.Date(2020, 4, 13, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "2020-02-24 — 2020-03-01",
			args: args{currentDate: time.Date(2020, 2, 29, 13, 28, 359, 0, time.UTC)},
			want: time.Date(2020, 2, 24, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetWeekFirstDate(tt.args.currentDate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWeekFirstDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
