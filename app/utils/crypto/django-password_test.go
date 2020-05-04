package crypto

import (
	"testing"
)

func TestGetDjangoPasswordHash(t *testing.T) {
	type args struct {
		password   string
		salt       string
		iterations int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"piskapupiska", args{"piskapupiska", "7Hv5Dh0m62Wq", 30000}, "pbkdf2_sha256$30000$7Hv5Dh0m62Wq$iUw/yF3HsSyXQv/keTC9QRsj8KxbNUguXEkJDGbl5ns="},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDjangoPasswordHash(tt.args.password, tt.args.salt, tt.args.iterations); got != tt.want {
				t.Errorf("GetDjangoPasswordHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDjangoPasswordEquals(t *testing.T) {
	type args struct {
		input string
		eq    string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"piskapupiska", args{"piskapupiska", "pbkdf2_sha256$30000$7Hv5Dh0m62Wq$iUw/yF3HsSyXQv/keTC9QRsj8KxbNUguXEkJDGbl5ns="}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DjangoPasswordEquals(tt.args.input, tt.args.eq); got != tt.want {
				t.Errorf("DjangoPasswordEquals() = %v, want %v", got, tt.want)
			}
		})
	}
}
