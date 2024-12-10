package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"users/config"
	postgresAdapter "users/internal/adapter/postgres"
	"users/internal/adapter/repository/zbank"
	"users/internal/infrastructure/database"
	"users/internal/infrastructure/database/postgres"
	"users/internal/ports/repository"
	"users/internal/usecase"
	zbankUC "users/internal/usecase/zbank"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	cfg *config.Config

	zbankDB *postgres.Postgres

	// repositories
	moneyRepository repository.MoneyRepository

	// services
	moneyUseCase usecase.MoneyUseCase

	router *chi.Mux
	server *http.Server
}

func New(cfg *config.Config) (*Server, error) {
	s := &Server{
		cfg: cfg,
	}

	if err := s.init(); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Server) init() error {
	if err := s.initDB(); err != nil {
		return fmt.Errorf("init db: %v", err)
	}
	if err := database.MigrateZBankDB(s.zbankDB); err != nil {
		return fmt.Errorf("migrate static db: %v", err)
	}

	s.initRepositories()
	s.initUseCases()
	s.initRouter()
	s.initHTTPServer()

	return nil
}

func (s *Server) initDB() error {
	var err error

	s.zbankDB, err = postgresAdapter.Connect(s.cfg.Server.StaticData.Connection)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) initRepositories() {
	s.moneyRepository = zbank.NewMoneyRepository(s.zbankDB)
}

func (s *Server) initUseCases() {
	s.moneyUseCase = zbankUC.NewMoneyUseCase(s.moneyRepository)
}

func (s *Server) initHTTPServer() {
	s.server = &http.Server{
		Addr:         fmt.Sprintf("%s:%s", s.cfg.Server.Addr, s.cfg.Server.Port),
		Handler:      s.router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

func (s *Server) Run() {
	log.Println("Server started")

	if err := s.server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
