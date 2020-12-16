package config

import (
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
)

func TestConfigureByEnv(t *testing.T) {
	tests := []struct {
		Scenario string
		Cfg      *AppConfig
		Want     string
	}{
		{
			Scenario: "Load local environment properties",
			Cfg: &AppConfig{
				EnvName: "local",
				Dir:     "test-fixtures",
			},
			Want: "properties.local.yml",
		},
		{
			Scenario: "Load dev environment properties",
			Cfg: &AppConfig{
				EnvName: "dev",
				Dir:     "test-fixtures",
			},
			Want: "properties.dev.yml",
		},
		{
			Scenario: "Load prod environment properties",
			Cfg: &AppConfig{
				EnvName: "prod",
				Dir:     "test-fixtures",
			},
			Want: "properties.prod.yml",
		},
	}
	for _, test := range tests {
		t.Run(test.Scenario, func(t *testing.T) {
			ConfigureApp(test.Cfg)
			pathOfFileUsed := viper.ConfigFileUsed()
			dir, fileUsed := filepath.Split(pathOfFileUsed)
			_, dirUsed := filepath.Split(filepath.Dir(dir))

			if dirUsed != test.Cfg.Dir {
				t.Errorf("Dir: Expected \"%s\", got \"%s\"", test.Cfg.Dir, dirUsed)
			}
			if fileUsed != test.Want {
				t.Errorf("File: Expected \"%s\", got \"%s\"", test.Want, fileUsed)
			}
		})
	}
}
