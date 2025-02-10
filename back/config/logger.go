package config

import (
	"go.uber.org/zap"
)

func DefaultLogger() *zap.Logger {
    cfg, err := NewServerConfig() 
    if err != nil {
        panic("Cannot create config")
    }
	var logger *zap.Logger
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
