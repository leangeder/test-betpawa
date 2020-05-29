package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	formatter := &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "serverity",
			logrus.FieldKeyMsg:   "message",
			// logrus.FieldKeyFunc:  "function",
			// logrus.FieldKeyFile:  "file",
		},
	}
	formatter.TimestampFormat = "2006-01-02 15:04:05.000000"
	log.SetFormatter(formatter)
	log.SetOutput(os.Stderr)
	log.SetReportCaller(true)
	log.SetLevel(logrus.DebugLevel)
}

func main() {

	go func() {
		m := mux.NewRouter().StrictSlash(false)
		m.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Ok"))
		}).Methods(http.MethodGet)

		srv := &http.Server{
			Addr:         "0.0.0.0:8090",
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  120 * time.Second,
			Handler:      m,
		}
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal("Server Error: %v", err)
		}
	}()

	// Set Handler due to receive the Slack Dialog message
	m := mux.NewRouter().StrictSlash(false)
	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now().UTC()
		w.Write([]byte("Success!"))
		end := time.Now().UTC()
		latency := end.Sub(start)
		log.WithFields(
			logrus.Fields{
				"status":     http.StatusOK,
				"method":     r.Method,
				"path":       r.URL.Path,
				"ip":         r.Header.Get("X-FORWARDED-FOR"),
				"duration":   latency.String(),
				"user_agent": r.UserAgent(),
			},
		).Info()
	}).Methods(http.MethodGet)

	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      m,
	}

	log.Infof("Start to listen on 0.0.0.0:8080")
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("Server Error: %v", err)
	}
}
