package api

import (
	"net/http"
)

func createLogMiddleware(api DaemonServer) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			api.Log.Debugf("%s %s", r.Method, r.RequestURI)
			next.ServeHTTP(w, r)
		})
	}
}
