package utils

import (
	"context"
	"testing"
)

func TestLarkAlarm(t *testing.T) {
	type args struct {
		ctx    context.Context
		hookID string
		title  string
		text   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "lark msg test",
			args: args{
				ctx:    context.Background(),
				hookID: "ec3a707a-0d56-491c-aec9-d6d594058171",
				title:  "扩缩容信息",
				text:   "扩容100台，成功99台，失败1台",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := LarkAlarm(tt.args.ctx, tt.args.hookID, tt.args.title, tt.args.text); (err != nil) != tt.wantErr {
				t.Errorf("LarkAlarm() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLocalIp(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "获取本地 ip",
			want:    "10.5.101.84",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LocalIp()
			if (err != nil) != tt.wantErr {
				t.Errorf("LocalIp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("LocalIp() got = %v, want %v", got, tt.want)
			}
		})
	}
}
