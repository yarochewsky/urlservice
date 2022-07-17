package health

import (
	"github.com/gorilla/mux"

	"net/http"
	realhttp "net/http"
)

func handleHealth() realhttp.HandlerFunc {
	return func(w realhttp.ResponseWriter, r *realhttp.Request) {
		w.WriteHeader(http.StatusOK)
	}
}

func New() realhttp.Handler {
	r := mux.NewRouter()

	r.Handle("/health", handleHealth()).Name("health")

	return r
}
