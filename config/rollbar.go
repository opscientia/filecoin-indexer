package config

import (
	"path"
	"runtime"

	"github.com/rollbar/rollbar-go"
)

// InitRollbar initializes the Rollbar integration
func InitRollbar(cfg *Config) {
	rollbar.SetToken(cfg.RollbarToken)
	rollbar.SetEnvironment(cfg.RollbarEnv)
	rollbar.SetServerRoot(serverRoot())
}

func serverRoot() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return ""
	}
	return path.Join(filename, "../..")
}

// LogPanic logs a recovered panic
func LogPanic() {
	if err := recover(); err != nil {
		rollbar.LogPanic(err, true)
		panic(err)
	}
}
