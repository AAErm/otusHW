package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	type args struct {
		fromPath string
		toPath   string
		offset   int64
		limit    int64
	}
	tests := []struct {
		name        string
		args        args
		expectedErr error
	}{
		{
			name: "unsupported file if isDir",
			args: args{
				fromPath: "testdata",
				toPath:   "",
				offset:   0,
				limit:    0,
			},
			expectedErr: ErrUnsupportedFile,
		},
		{
			name: "offset exceeds file size",
			args: args{
				fromPath: "testdata/input.txt",
				toPath:   "out.txt",
				offset:   9999999,
				limit:    0,
			},
			expectedErr: ErrOffsetExceedsFileSize,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Copy(tt.args.fromPath, tt.args.toPath, tt.args.offset, tt.args.limit)
			require.Error(t, tt.expectedErr, err)
		})
	}
}
