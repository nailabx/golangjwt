package api

import (
	"database/sql"
	"github.com/nailabx/golangjwt/service/user"
	"log"
	"net/http"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()
	v1 := http.NewServeMux()
	v1.Handle("/api/v1/", http.StripPrefix("/api/v1", router))

	// Register user routes
	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(router)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}
