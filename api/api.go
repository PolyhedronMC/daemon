package api
 
import (
	"github.com/polyhedronmc/daemon/config"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap"
    "net/http"
	"github.com/gorilla/mux"
	"github.com/polyhedronmc/daemon/daemon"
)

// DaemonServer The Polyhedron daemon API server.
type DaemonServer struct {
	Daemon daemon.Daemon
	Config config.DaemonConfig
	Log *zap.SugaredLogger

	router *mux.Router
}

// Create Create a new API server instance.
func Create(config config.DaemonConfig) DaemonServer {
    router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", GetVersion)
	
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

	return DaemonServer {
		Config: config,
		Log: logger.Sugar(),
		router: router,
	}
}

// Listen Listen for requests.
func (s DaemonServer) Listen() {
	s.router.Use(createLogMiddleware(s))
	s.Log.Info("Starting HTTP server on port 8080")
	s.Log.Fatal(http.ListenAndServe(":8080", s.router))
}

