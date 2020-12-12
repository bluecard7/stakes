package config

import (
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
)

func TestConfigureByEnv(t *testing.T) {
	tests := []struct {
		Cfg  *AppConfig
		Want string
	}{
		{
			Cfg: &AppConfig{
				EnvName: "local",
				Dir:     "test-fixtures",
			},
			Want: "properties.local.yml",
		},
		{
			Cfg: &AppConfig{
				EnvName: "dev",
				Dir:     "test-fixtures",
			},
			Want: "properties.dev.yml",
		},
		{
			Cfg: &AppConfig{
				EnvName: "prod",
				Dir:     "test-fixtures",
			},
			Want: "properties.prod.yml",
		},
	}
	for _, test := range tests {
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
	}
}
