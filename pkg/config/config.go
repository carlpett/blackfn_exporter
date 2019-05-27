package config

import (
	"os"

	"github.com/go-kit/kit/log"
	"github.com/prometheus/blackbox_exporter/config"
	"github.com/prometheus/common/promlog"
)

var (
	Config *config.Config
	Logger log.Logger
)

func init() {
	sc := &config.SafeConfig{C: &config.Config{}}
	configFile := "blackbox_exporter.yml"
	if f := os.Getenv("CONFIG_FILE"); f != "" {
		configFile = f
	}
	if err := sc.ReloadConfig(configFile); err != nil {
		panic(err)
	}
	Config = sc.C

	logLevel := "info"
	if lvl := os.Getenv("LOG_LEVEL"); lvl != "" {
		logLevel = lvl
	}
	al := promlog.AllowedLevel{}
	al.Set(logLevel)
	Logger = promlog.New(al)
}
