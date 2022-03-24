package config_test

import (
	"path/filepath"
	"reflect"
	"testing"

	"polygon.am/core/pkg/config"
	"polygon.am/core/pkg/types"
	"polygon.am/core/pkg/util"
)

func TestParseConfig(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    *types.Config
		wantErr bool
	}{
		{
			name: "ParseConfig() must return correct struct",
			args: args{
				path: util.AssumeNoError(filepath.Abs("./tests/.conf-test.yaml")),
			},
			want: &types.Config{
				Polygon: types.Polygon{
					Addr: "127.0.0.1:1234",
				},
				Databases: types.Databases{
					Redis:    "redis://testing:testing@127.0.0.1:1234/",
					Postgres: "postgres://testing:testing@127.0.0.1:1234/",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := config.ParseConfig(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
