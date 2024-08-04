package config

import (
	"reflect"
	"testing"

	"github.com/imega/mt"
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
				AMQP: AMQPConf{
					DSN:         "amqp://guest:guest@amqp:5672/",
					ServiceName: "calendar",
					MtConfig: mt.Config{
						Services: map[string]mt.Service{
							"calendar": {
								Exchange: mt.Exchange{
									Name:    "amq.direct",
									Type:    "direct",
									Durable: true,
									Queue: mt.Queue{
										Name:    "calendar",
										Durable: true,
									},
									Binding: mt.Binding{
										Key: "calendar",
									},
								},
							},
						},
					},
				},
				Sheduler: ShedulerConf{
					Interval: 10,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewConfig(tt.args.filepath)

			if got.Error != nil {
				t.Errorf("NewConfig Error = %v", got.Error)
			}

			// генерируется уникальный consume tag. потому не совпадает
			// if !reflect.DeepEqual(got.AMQP, tt.want.AMQP) {
			// 	t.Errorf("NewConfig AMQP = \n%v, want \n%v", got.AMQP, tt.want.AMQP)
			// }

			if !reflect.DeepEqual(got.Grpc, tt.want.Grpc) {
				t.Errorf("NewConfig Grpc = %v, want %v", got.Grpc, tt.want.Grpc)
			}

			if !reflect.DeepEqual(got.DB, tt.want.DB) {
				t.Errorf("NewConfig DB = %v, want %v", got.DB, tt.want.DB)
			}

			if !reflect.DeepEqual(got.Server, tt.want.Server) {
				t.Errorf("NewConfig Server = %v, want %v", got.Server, tt.want.Server)
			}

			if !reflect.DeepEqual(got.Logger, tt.want.Logger) {
				t.Errorf("NewConfig Logger = %v, want %v", got.Logger, tt.want.Logger)
			}

			if !reflect.DeepEqual(got.Sheduler, tt.want.Sheduler) {
				t.Errorf("NewConfig Sheduler = %v, want %v", got.Sheduler, tt.want.Sheduler)
			}
		})
	}
}
