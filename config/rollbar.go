package config

import (
	"github.com/rollbar/rollbar-go"
)

// InitRollbar initializes the Rollbar integration
func InitRollbar(cfg *Config) {
	rollbar.SetToken(cfg.RollbarToken)
	rollbar.SetEnvironment(cfg.AppEnv)
	rollbar.SetServerRoot(cfg.RollbarServerRoot)
}

// LogPanic logs a recovered panic
func LogPanic() {
	if err := recover(); err != nil {
		rollbar.LogPanic(err, true)
		panic(err)
	}
}
