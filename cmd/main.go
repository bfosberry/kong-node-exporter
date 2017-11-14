package main

import (
	"fmt"
	"net/http"
	//"os"

	"github.com/Sirupsen/logrus"
	"github.com/bfosberry/kong-node-exporter/handlers"
	"github.com/bfosberry/kong-node-exporter/kong"
	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
)

// Config holds the application config loaded from the environment
type Config struct {
	KongIP   string `split_words:"true" required:"true"`
	KongPort int    `split_words:"true" required:"true"`
	Port     int    `default:"8080"`
}

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	log := logrus.WithFields(logrus.Fields{
		"module": "kong-node-exporter",
	})
	log.Info("startup")
	config, err := getConfig()
	if err != nil {
		log.Fatal(err)
	}

	kongClient := kong.NewHTTPClient("http", config.KongIP, config.KongPort)

	r := mux.NewRouter()
	r.HandleFunc("/health", handlers.Health)
	r.HandleFunc("/metrics", handlers.NewMetricsHandler(kongClient, log))

	http.Handle("/", r)
	log.WithField("port", config.Port).Info("listening")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil))
}

func getConfig() (*Config, error) {
	c := &Config{}
	err := envconfig.Process("exporter", c)
	return c, err
}
