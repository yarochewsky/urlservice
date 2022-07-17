package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"urlservice/counter"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func handleURL(logger *logrus.Entry) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ur urlRequest

		if err := json.NewDecoder(r.Body).Decode(&ur); err != nil {
			logger.WithError(err).Warn("failed to unmarshal request body")
			w.WriteHeader(http.StatusBadRequest)
			writeError(w, fmt.Errorf("invalid request body"))
			return
		}

		logger.WithField("url", ur.Target).Info("received request")

		resp, err := http.Get(ur.Target)
		if err != nil {
			logger.WithError(err).Warn("failed to load target")
			w.WriteHeader(http.StatusInternalServerError)
			writeError(w, fmt.Errorf("failed to load target"))
			return
		}
		defer resp.Body.Close()

		content, err := io.ReadAll(resp.Body)
		if err != nil {
			logger.WithError(err).Warn("failed to read content")
			w.WriteHeader(http.StatusInternalServerError)
			writeError(w, fmt.Errorf("failed to read content"))
			return
		}

		w.WriteHeader(http.StatusOK)
		writeSuccess(w, counter.MapFrequencies(string(content)))
	}
}

func writeSuccess(w http.ResponseWriter, words map[string]int) {
	r, _ := json.Marshal(&urlResponse{WordCount: words})
	w.Write(r)
}

func writeError(w http.ResponseWriter, err error) {
	r, _ := json.Marshal(&urlResponse{Error: err.Error()})
	w.Write(r)
}

type urlResponse struct {
	WordCount map[string]int `json:"words,omitempty"`
	Error     string         `json:"error,omitempty"`
}

type urlRequest struct {
	Target string `json:"target"`
}

func New(logger *logrus.Entry) http.Handler {
	r := mux.NewRouter()

	r.Handle("/url", handleURL(logger)).Methods("POST").Name("solve")

	return r
}
