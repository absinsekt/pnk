package strings

import "testing"

func TestGenerateRandomString(t *testing.T) {
	type args struct {
		length int64
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"16", args{16}, 16, false},
		{"64", args{64}, 64, false},
		{"128", args{128}, 128, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateRandomString(tt.args.length)
			if len(got) != tt.want {
				t.Errorf("GenerateRandomString() length = %v, want %v", got, tt.want)
			}
		})
	}
}
