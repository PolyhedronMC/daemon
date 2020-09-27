package main

import (
	"github.com/polyhedronmc/daemon/config"
	"github.com/polyhedronmc/daemon/database"
	"github.com/polyhedronmc/daemon/docker"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Daemon struct{}

func main() {
	cfg := zap.Config{
		Level:    zap.NewAtomicLevelAt(zap.DebugLevel),
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			LevelKey:    "level",
			EncodeLevel: zapcore.LowercaseLevelEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, _ := cfg.Build()
	defer logger.Sync()
	log := logger.Sugar()

	log.Info("Starting Polyhedron daemon...")
	config := config.GetConfig()
	version := docker.DetectDocker(config)

	log.Infof("Docker %s", version)

	log.Info("Connecting to PostgreSQL...")
	database.Connect(config.Database)
}
