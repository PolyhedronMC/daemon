package daemon

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"github.com/polyhedronmc/daemon/config"
	"github.com/polyhedronmc/daemon/docker"
	"github.com/polyhedronmc/daemon/database"
)

// Daemon The Polyhedron daemon.
type Daemon struct {
	Config config.DaemonConfig
	Log *zap.SugaredLogger
	// api ApiServer
}

// Create Create a new daemon instance.
func Create(config config.DaemonConfig) Daemon {
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

	return Daemon {
		Config: config,
		Log: logger.Sugar(),
	}
}

// Start Start this deamon instance.
func (d *Daemon) Start() {
	docker.SetupClient()

	version := docker.GetVersion()
	d.Log.Infof("Docker %s", version)
	d.Log.Info("Connecting to PostgreSQL")
	database.Connect(d.Config.Database)

	monitor := d.CreateMonitor("gay")
	monitor.Start()
}
