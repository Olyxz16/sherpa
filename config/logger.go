package config

import (
	"go.uber.org/zap"
)

func DefaultLogger() *zap.Logger {
    cfg := NewServerConfig() 
	var logger *zap.Logger
	var err error
    if cfg.Debug {
		logger, err = zap.NewDevelopment()
        if err != nil {
            panic("Cannot create development logger")
        }
    } else {
		logger, err = zap.NewProduction()
        if err != nil {
            panic("Cannot create development logger")
        }
    }
	zap.ReplaceGlobals(logger)
	return logger
}
