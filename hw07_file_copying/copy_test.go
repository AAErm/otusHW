package main

import (
	"testing"
)

func TestCopy(t *testing.T) {
	type args struct {
		fromPath string
		toPath   string
		offset   int64
		limit    int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "unsupported file if isDir",
			args: args{
				fromPath: "testdata",
				toPath:   "",
				offset:   0,
				limit:    0,
			},
			wantErr: ErrUnsupportedFile,
		},
		{
			name: "offset exceeds file size",
			args: args{
				fromPath: "testdata/input.txt",
				toPath:   "out.txt",
				offset:   9999999,
				limit:    0,
			},
			wantErr: ErrOffsetExceedsFileSize,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Copy(tt.args.fromPath, tt.args.toPath, tt.args.offset, tt.args.limit); err != nil && err != tt.wantErr {
				t.Errorf("Copy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_werwer(t *testing.T) {
	type args struct {
		in0 string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := werwer(tt.args.in0); (err != nil) != tt.wantErr {
				t.Errorf("werwer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
