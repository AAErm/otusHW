package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		// uncomment if task with asterisk completed
		// {input: `qwe\4\5`, expected: `qwe45`},
		// {input: `qwe\45`, expected: `qwe44444`},
		// {input: `qwe\\5`, expected: `qwe\\\\\`},
		// {input: `qwe\\\3`, expected: `qwe\3`},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b"}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}

func Test_nextIsInt(t *testing.T) {
	type args struct {
		s string
		i int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "EOF",
			args: args{
				s: "qw4",
				i: 2,
			},
			want: false,
		},
		{
			name: "next is int",
			args: args{
				s: "qw4",
				i: 1,
			},
			want: true,
		},
		{
			name: "next is not int",
			args: args{
				s: "qw4",
				i: 0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextIsInt(tt.args.s, tt.args.i); got != tt.want {
				t.Errorf("nextIsNotInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_prevIsRepeated(t *testing.T) {
	type args struct {
		s string
		i int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "prev is empty",
			args: args{
				s: "4wa",
				i: 0,
			},
			want: false,
		},
		{
			name: "prev is int",
			args: args{
				s: "4wa",
				i: 1,
			},
			want: false,
		},
		{
			name: "prev is not int",
			args: args{
				s: "4wa",
				i: 2,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prevIsRepeated(tt.args.s, tt.args.i); got != tt.want {
				t.Errorf("prevIsRepeated() = %v, want %v", got, tt.want)
			}
		})
	}
}
