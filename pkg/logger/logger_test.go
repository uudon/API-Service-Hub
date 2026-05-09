package logger

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func TestNew(t *testing.T) {
	logger := New()

	if logger == nil {
		t.Fatal("New() should return a non-nil logger")
	}
}

func TestNewWithLevel(t *testing.T) {
	tests := []struct {
		name  string
		level zerolog.Level
	}{
		{
			name:  "Debug level",
			level: zerolog.DebugLevel,
		},
		{
			name:  "Info level",
			level: zerolog.InfoLevel,
		},
		{
			name:  "Warn level",
			level: zerolog.WarnLevel,
		},
		{
			name:  "Error level",
			level: zerolog.ErrorLevel,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			zerolog.SetGlobalLevel(tt.level)
			logger := New()

			if logger == nil {
				t.Error("New() should return a non-nil logger")
			}
		})
	}
}

func TestLoggerOutput(t *testing.T) {
	var buf bytes.Buffer
	logger := zerolog.New(&buf)

	logger.Info().Msg("test message")

	output := buf.String()
	if output == "" {
		t.Error("Expected non-empty output")
	}

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		t.Errorf("Failed to parse JSON output: %v", err)
	}

	if result["level"] != "info" {
		t.Errorf("Expected level 'info', got '%v'", result["level"])
	}

	if result["message"] != "test message" {
		t.Errorf("Expected message 'test message', got '%v'", result["message"])
	}
}

func TestLoggerLevels(t *testing.T) {
	levels := []zerolog.Level{
		zerolog.InfoLevel,
		zerolog.WarnLevel,
		zerolog.ErrorLevel,
	}

	for _, level := range levels {
		t.Run(level.String(), func(t *testing.T) {
			var buf bytes.Buffer
			logger := zerolog.New(&buf).Level(level)

			logger.WithLevel(level).Msg("test")

			if buf.Len() == 0 {
				t.Errorf("Expected output for level %s", level)
			}
		})
	}
}


func TestGlobalLogger(t *testing.T) {
	originalLogger := log.Logger
	defer func() {
		log.Logger = originalLogger
	}()

	logger := New()
	log.Logger = *logger

	// Just verify we can set it without panicking
	_ = log.Logger
}
