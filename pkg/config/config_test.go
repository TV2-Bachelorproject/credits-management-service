package config_test

import (
	"testing"

	"github.com/TV2-Bachelorproject/server/pkg/config"
)

func TestDefaultConfig(t *testing.T) {
	cfg := config.Get()

	if cfg.DB.Host != "localhost" {
		t.Errorf("expected localhost; got %s", cfg.DB.Host)
	}

	if cfg.DB.Port != 5432 {
		t.Errorf("expected 5432; got %d", cfg.DB.Port)
	}

	if cfg.DB.User != "root" {
		t.Errorf("expected root; got %s", cfg.DB.User)
	}

	if cfg.DB.Password != "root" {
		t.Errorf("expected root; got %s", cfg.DB.Password)
	}

	if cfg.DB.Database != "root" {
		t.Errorf("expected root; got %s", cfg.DB.Database)
	}
}
