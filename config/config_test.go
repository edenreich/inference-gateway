package config_test

import (
	"context"
	"testing"
	"time"

	config "github.com/edenreich/inference-gateway/config"
	"github.com/sethvargo/go-envconfig"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		name          string
		env           map[string]string
		expectedCfg   config.Config
		expectedError string
	}{
		{
			name: "Success_AllEnvVariablesSet",
			env: map[string]string{
				"APPLICATION_NAME":        "inference-gateway",
				"ENABLE_TELEMETRY":        "true",
				"SERVER_READ_TIMEOUT":     "20s",
				"SERVER_WRITE_TIMEOUT":    "40s",
				"SERVER_IDLE_TIMEOUT":     "150s",
				"ENVIRONMENT":             "development",
				"SERVER_HOST":             "192.168.1.1",
				"SERVER_PORT":             "9090",
				"SERVER_TLS_CERT_PATH":    "/path/to/cert.pem",
				"SERVER_TLS_KEY_PATH":     "/path/to/key.pem",
				"OLLAMA_API_URL":          "http://ollama.local",
				"GROQ_API_KEY":            "groq123",
				"OPENAI_API_KEY":          "openai123",
				"GOOGLE_AISTUDIO_API_KEY": "google123",
			},
			expectedCfg: config.Config{
				ApplicationName:    "inference-gateway",
				EnableTelemetry:    true,
				ServerReadTimeout:  20 * time.Second,
				ServerWriteTimeout: 40 * time.Second,
				ServerIdleTimeout:  150 * time.Second,
				Environment:        "development",
				ServerHost:         "192.168.1.1",
				ServerPort:         "9090",
				ServerTLSCertPath:  "/path/to/cert.pem",
				ServerTLSKeyPath:   "/path/to/key.pem",
				OllamaAPIURL:       "http://ollama.local",
				GroqAPIURL:         "https://api.groq.com",
				GroqAPIKey:         "groq123",
				OpenaiAPIURL:       "https://api.openai.com",
				OpenaiAPIKey:       "openai123",
				GoogleAIStudioURL:  "https://generativelanguage.googleapis.com",
				GoogleAIStudioKey:  "google123",
			},
		},
		{
			name: "Success_Defaults",
			env:  map[string]string{},
			expectedCfg: config.Config{
				ApplicationName:    "inference-gateway",
				EnableTelemetry:    false,
				ServerReadTimeout:  5 * time.Second,
				ServerWriteTimeout: 10 * time.Second,
				ServerIdleTimeout:  120 * time.Second,
				Environment:        "production",
				ServerHost:         "127.0.0.1",
				ServerPort:         "8080",
				OllamaAPIURL:       "http://ollama:8080",
				GroqAPIURL:         "https://api.groq.com",
				OpenaiAPIURL:       "https://api.openai.com",
				GoogleAIStudioURL:  "https://generativelanguage.googleapis.com",
			},
		},
		{
			name: "Error_InvalidEnableTelemetry",
			env: map[string]string{
				"ENABLE_TELEMETRY": "notabool",
			},
			expectedError: `EnableTelemetry("notabool"): strconv.ParseBool: parsing "notabool": invalid syntax`,
		},
		{
			name: "Error_InvalidServerReadTimeout",
			env: map[string]string{
				"SERVER_READ_TIMEOUT": "invalid",
			},
			expectedError: `ServerReadTimeout("invalid"): time: invalid duration "invalid"`,
		},
		{
			name: "Error_InvalidServerWriteTimeout",
			env: map[string]string{
				"SERVER_WRITE_TIMEOUT": "invalid",
			},
			expectedError: `ServerWriteTimeout("invalid"): time: invalid duration "invalid"`,
		},
		{
			name: "Error_InvalidServerIdleTimeout",
			env: map[string]string{
				"SERVER_IDLE_TIMEOUT": "invalid",
			},
			expectedError: `ServerIdleTimeout("invalid"): time: invalid duration "invalid"`,
		},
		{
			name: "PartialEnvVariables",
			env: map[string]string{
				"ENABLE_TELEMETRY":    "true",
				"SERVER_READ_TIMEOUT": "25s",
				"ENVIRONMENT":         "development",
				"OLLAMA_API_URL":      "http://ollama.test",
			},
			expectedCfg: config.Config{
				ApplicationName:    "inference-gateway",
				EnableTelemetry:    true,
				ServerReadTimeout:  25 * time.Second,
				ServerWriteTimeout: 10 * time.Second,
				ServerIdleTimeout:  120 * time.Second,
				Environment:        "development",
				ServerHost:         "127.0.0.1",
				ServerPort:         "8080",
				OllamaAPIURL:       "http://ollama.test",
				GroqAPIURL:         "https://api.groq.com",
				OpenaiAPIURL:       "https://api.openai.com",
				GoogleAIStudioURL:  "https://generativelanguage.googleapis.com",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lookuper := envconfig.MapLookuper(tt.env)
			var cfg config.Config
			err := envconfig.ProcessWith(context.Background(), &envconfig.Config{
				Target:   &cfg,
				Lookuper: lookuper,
			})
			if tt.expectedError != "" {
				if err == nil {
					t.Fatalf("Expected error '%s', got nil", tt.expectedError)
				}
				if err.Error() != tt.expectedError {
					t.Errorf("Expected error '%s', got '%s'", tt.expectedError, err.Error())
				}
				return
			}

			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}

			if cfg != tt.expectedCfg {
				t.Errorf("Expected config %+v, got %+v", tt.expectedCfg, cfg)
			}
		})
	}
}
