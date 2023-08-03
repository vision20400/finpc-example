package main

import (
	"database/sql"
	"fmt"
	"net"
	"os"
	"sync"

	// external packages
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"

	// project packages
	. "github.com/ghilbut/finpc/grpc"
	. "github.com/ghilbut/finpc/rest"
)

func main() {
	log.SetLevel(log.TraceLevel)

	db, err := OpenDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		port := 8080
		addr := fmt.Sprintf(":%d", port)

		rest := NewRestServer()

		log.Printf("run RESTful server on port %d", port)
		if err := fasthttp.ListenAndServe(addr, rest.Handler); err != nil {
			log.Fatalf("failed to run RESTful server: %v", err)
		}
	}()

	go func() {
		port := 9095
		addr := fmt.Sprintf(":%d", port)

		listen, err := net.Listen("tcp4", addr)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		grpc := NewGrpcServer(db)

		log.Printf("run gRPC server on port %d", port)
		if err := grpc.Serve(listen); err != nil {
			log.Fatalf("failed to run gRPC server: %v", err)
		}
	}()

	wg.Wait()
}

func OpenDatabase() (*sql.DB, error) {

	host := os.Getenv("PG_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("PG_PORT")
	if port == "" {
		port = "5432"
	}
	user := os.Getenv("PG_USER")
	if user == "" {
		user = "postgres"
	}
	pw := os.Getenv("PG_PASSWORD")
	if pw == "" {
		pw = "postgrespw"
	}
	db := os.Getenv("PG_DATABASE")
	if db == "" {
		db = "postgres"
	}
	ssl := os.Getenv("PG_SSLMODE")
	if ssl == "" {
		ssl = "disable"
	}

	log.Infoln("PG_HOST: ", host)
	log.Infoln("PG_PORT: ", port)
	log.Infoln("PG_USER: ", user)
	log.Infoln("PG_DATABASE: ", db)
	log.Infoln("PG_SSLMODE: ", ssl)

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, pw, db, ssl)
	return sql.Open("postgres", dsn)
}
