package user

import (
	"testing"
)

func Test_createUserTable(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"create", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := createUserTable(); (err != nil) != tt.wantErr {
				t.Errorf("createUserTable() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_dropUserTable(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"drop", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := dropUserTable(); (err != nil) != tt.wantErr {
				t.Errorf("dropUserTable() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_createSuperUser(t *testing.T) {
	type args struct {
		username    string
		password    string
		firstName   string
		lastName    string
		email       string
		isStaff     bool
		isSuperuser bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"create superuser", args{
			"ramone",
			"petsematary",
			"Dee Dee",
			"Ramone",
			"ramone@mail.punk",
			true,
			true,
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateUser(tt.args.username, tt.args.password, tt.args.firstName, tt.args.lastName, tt.args.email, tt.args.isStaff, tt.args.isSuperuser); (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
