// Nama Package
package main

// Package
import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"server/models"
	"time"

	// Import package db
	_ "github.com/go-sql-driver/mysql"
)

// Versi Server
const version = "1.0.0"

// create variable config
type config struct {
	port int
	env  string

	// create database
	db struct {
		dsn string
	}
	// end section create database
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
	models models.Models
}

// inisialisasi server
func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "Development", "Application Environment (development | production)")

	// create Database
	var dbUser, dbPassword string
	flag.StringVar(&dbUser, "dbuser", "root", "MySQL database user")
	flag.StringVar(&dbPassword, "dbpassword", "", "MySQL database password")

	flag.StringVar(&cfg.db.dsn, "dsn", fmt.Sprintf("%s:%s@tcp(localhost:3306)/id_gomoviereact", dbUser, dbPassword), "MySQL connection config")
	// end section create database
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// create Database
	db, err := openDB(cfg, logger)

	if err != nil {
		logger.Fatal(err)
		return
	}
	defer db.Close()
	// end section create database

	app := &application{
		config: cfg,
		logger: logger,
		models: models.NewModels(db),
	}

	fmt.Println("Server Is Running")

	serve := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	logger.Printf("Starting Server On Port %d\n", cfg.port)
	err = serve.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}

// create Database
func openDB(cfg config, logger *log.Logger) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.db.dsn)

	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err)
	}

	logger.Println("Connected to the MySQL database!")
	return db, nil
}
