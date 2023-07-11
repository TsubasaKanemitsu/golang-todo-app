package config

import (
	"reflect"
	"testing"
)

func TestPrepare(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "sample test",
			want: "testdb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Prepare(); !reflect.DeepEqual(got.Postgres.DBName, tt.want) {
				t.Errorf("Prepare() = %v, want %v", got.Postgres.DBName, tt.want)
			}
		})
	}
}
