package config

import (
	"testing"
)

func TestNewConfig(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name string
		args args
		want Config
	}{
		{
			name: "test get conf",
			args: args{filepath: "../../configs/config.json"},
			want: Config{
				Logger: LoggerConf{
					Level: "INFO",
				},
				Server: ServerConf{
					Host: "localhost",
					Port: 8080,
				},
				DB: DBConf{
					Use:      false,
					Host:     "localhost",
					Port:     5435,
					User:     "root",
					Password: "qwerty",
				},
				Grpc: GrpcConf{
					Host: ":50051",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConfig(tt.args.filepath); got != tt.want {
				t.Errorf("NewConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
