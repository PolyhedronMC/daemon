package api

import (
	"net/http"
	"encoding/json"
	"github.com/polyhedronmc/daemon/constants"
)

type version struct {
	Version string `json:"version"`
}

// GetVersion GET /version
func GetVersion(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(version { Version: constants.DaemonVersion })
}


