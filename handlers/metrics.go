package handlers

import (
	"io"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/bfosberry/kong-node-exporter/kong"
)

// NewMetricsHandler creates a new http handler to return metrics
// from the provided kong client
func NewMetricsHandler(kongClient kong.Client, log *logrus.Entry) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := kongClient.GetStatus()
		if err != nil {
			w.WriteHeader(500)
			log.Error(err)
			return
		}

		metricsReader, err := resp.MetricsReader()
		if err != nil {
			w.WriteHeader(500)
			log.Error(err)
			return
		}
		if _, err := io.Copy(w, metricsReader); err != nil {
			log.Error(err)
		}
	}
}
