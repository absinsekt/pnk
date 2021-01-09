package templateset

import "testing"

func Test_buildTemplateName(t *testing.T) {
	type args struct {
		fullPath string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"empty", args{""}, "."},
		{"templates", args{"templates/template.html"}, "template.html"},
		{"root template", args{"template.html"}, "template.html"},
		{"shared template", args{"shared/template.html"}, "template.html"},
		{"not shared template", args{"somewhere/template.html"}, "somewhere/template.html"},
		{"not shared \"shared\" template", args{"somewhere/shared/template.html"}, "somewhere/shared/template.html"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTemplateName(tt.args.fullPath); got != tt.want {
				t.Errorf("buildTemplateName() = %v, want %v", got, tt.want)
			}
		})
	}
}
