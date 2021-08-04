package config

import (
	"testing"

	mocks "github.com/yanoandri/yano-golang-training-beginner/mocks/config"
)

func TestGetPostgresConnectionString(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{
			name: "success_get_postgres_connection_string",
			want: "host=127.0.0.1 port=5432 user=postgres dbname=payment password= sslmode=disable",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPostgresConnectionString("postgres", "", "payment", "127.0.0.1", "5432"); got != tt.want {
				t.Errorf("GetPostgresConnectionString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetupDB(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "success_run_setup_db",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conn := mocks.ConfigDB{}
			conn.On("SetupDB").Return(true)
			conn.SetupDB()
		})
	}
}
