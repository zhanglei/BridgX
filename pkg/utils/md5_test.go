package utils

import "testing"

func TestBase64Md5(t *testing.T) {
	type args struct {
		params string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1",
			args{params: "123456"},
			"87d9bb400c0634691f0e3baaf1e2fd0d"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Base64Md5(tt.args.params); got != tt.want {
				t.Errorf("Base64Md5() = %v, want %v", got, tt.want)
			}
		})
	}
}
