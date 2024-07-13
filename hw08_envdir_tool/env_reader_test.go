package main

import (
	"reflect"
	"testing"
)

func TestReadDir(t *testing.T) {
	type args struct {
		dir string
	}

	tests := []struct {
		name    string
		args    args
		want    Environment
		wantErr bool
	}{
		{
			name:    "expected dir return error",
			args:    args{dir: "qweqwe"},
			want:    nil,
			wantErr: true,
		},
		{
			name: "correct with testdata",
			args: args{dir: "testdata/env"},
			want: map[string]EnvValue{
				"BAR": {
					Value:      "bar",
					NeedRemove: false,
				},
				"EMPTY": {
					Value:      "",
					NeedRemove: false,
				},
				"FOO": {
					Value:      "   foo\nwith new line",
					NeedRemove: false,
				},
				"HELLO": {
					Value:      `"hello"`,
					NeedRemove: false,
				},
				"UNSET": {
					Value:      "",
					NeedRemove: true,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadDir(tt.args.dir)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadDir() = %v, want %v", got, tt.want)
			}
		})
	}
}
