package numbers

import "testing"

func TestRobin(t *testing.T) {
	type args struct {
		value int64
		min   int64
		max   int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"3 6 2 = 4", args{min: 3, max: 6, value: 4}, 4},
		{"3 6 2 = 6", args{min: 3, max: 6, value: 2}, 6},
		{"3 6 9 = 5", args{min: 3, max: 6, value: 9}, 5},
		{"0 6 8 = 1", args{min: 0, max: 6, value: 8}, 1},
		{"6 6 7 = 6", args{min: 6, max: 6, value: 7}, 6},
		{"-2 3 7 = 1", args{min: -2, max: 3, value: 7}, 1},
		{"1 7 0 = 7", args{min: 1, max: 7, value: 0}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Robin(tt.args.value, tt.args.min, tt.args.max); got != tt.want {
				t.Errorf("Robin() = %v, want %v", got, tt.want)
			}
		})
	}
}
