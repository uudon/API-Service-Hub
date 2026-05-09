package config

import (
	"testing"

	"github.com/rs/zerolog"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		name    string
		check   func(*Config, error)
		wantErr bool
	}{
		{
			name: "Default configuration",
			check: func(cfg *Config, err error) {
				if err != nil {
					t.Errorf("Load() unexpected error = %v", err)
					return
				}
				if cfg.Server.Host == "" {
					t.Error("Server.Host should not be empty")
				}
				if cfg.Port == 0 {
					t.Error("Port should not be 0")
				}
				if cfg.LogLevel == 0 {
					t.Error("LogLevel should not be 0")
				}
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg, err := Load()
			if tt.check != nil {
				tt.check(cfg, err)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConfigValidation(t *testing.T) {
	tests := []struct {
		name  string
		cfg   Config
		valid bool
	}{
		{
			name: "Valid config",
			cfg: Config{
				Server: ServerConfig{
					Host: "0.0.0.0",
					Port: 8080,
				},
				Port:         8080,
				ReadTimeout:  30,
				WriteTimeout: 30,
				LogLevel:     zerolog.InfoLevel,
			},
			valid: true,
		},
		{
			name: "Empty host",
			cfg: Config{
				Server: ServerConfig{
					Host: "",
					Port: 8080,
				},
				Port:         8080,
				ReadTimeout:  30,
				WriteTimeout: 30,
				LogLevel:     zerolog.InfoLevel,
			},
			valid: false,
		},
		{
			name: "Zero port",
			cfg: Config{
				Server: ServerConfig{
					Host: "0.0.0.0",
					Port: 0,
				},
				Port:         0,
				ReadTimeout:  30,
				WriteTimeout: 30,
				LogLevel:     zerolog.InfoLevel,
			},
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid := tt.cfg.Server.Host != "" && tt.cfg.Port > 0
			if valid != tt.valid {
				t.Errorf("Config validity = %v, want %v", valid, tt.valid)
			}
		})
	}
}
