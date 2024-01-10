// Nama Package
package main

//Package
import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	//import package db

	_ "github.com/lib/pq"
)

// Versi Server
const version = "1.0.0"

// create variable config
type config struct {
	port int
	env  string

	//create database
	db struct {
		dsn string
	}
	//end section crete database
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

	//create Database
	flag.StringVar(&cfg.db.dsn, "dsn", "postgres://postgres@localhost/movie?sslmode=disable", "postgres connection config")
	//end section crete database
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	//create Database
	db, err := openDB(cfg)

	if err != nil {
		logger.Fatal(err)

	}
	defer db.Close()
	//end section crete database
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
	err = serve.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}

// create Database
func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)

	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = db.PingContext(ctx)

	if err != nil {
		return nil, err
	}
	return db, nil
}

//end section crete database
