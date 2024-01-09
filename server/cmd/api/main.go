// Nama Package
package main

//Package
import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Versi Server
const version = "1.0.0"

// create variable config
type config struct {
	port int
	env  string
}

// create AppStatus
type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

// create server aplikasi
type application struct {
	config config
	logger *log.Logger
}

// inisialiasi server
func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "Development", "Application Environment (development | production)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

	fmt.Println("Server Is Running")

	// http.HandleFunc("/status", func(rw http.ResponseWriter, r *http.Request) {
	// 	currentStatus := AppStatus{
	// 		Status:      "Online",
	// 		Environment: cfg.env,
	// 		Version:     version,
	// 	}
	// 	res, err := json.MarshalIndent(currentStatus, "", "\t")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
	// 		return
	// 	}
	// 	rw.Header().Set("Content-Type", "application/json")
	// 	rw.WriteHeader(http.StatusOK)
	// 	rw.Write(res)
	// })

	// err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.port), nil)

	serve := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	logger.Printf("Starting Server On Port", cfg.port)
	err := serve.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
