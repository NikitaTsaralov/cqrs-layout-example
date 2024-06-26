package server

import (
	"context"

	"cqrs-layout-example/config"

	"github.com/NikitaTsaralov/utils/connectors/jaeger"
	"github.com/NikitaTsaralov/utils/connectors/postgres"
	"github.com/NikitaTsaralov/utils/logger"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	cfg      *config.Config
	postgres *sqlx.DB
	tracer   *jaeger.Trace
}

func NewServer(cfg *config.Config) *Server {
	return &Server{cfg: cfg}
}

func (s *Server) Start() {
	logger.Infof("starting...")

	s.startDB()
	s.startJaeger()
}

func (s *Server) Stop(ctx context.Context) error {
	s.stopDB()
	s.stopJaeger(ctx)

	return nil
}

func (s *Server) startDB() {
	var err error

	s.postgres = postgres.New(s.cfg.Postgres)
	if err != nil {
		logger.Panicf("PostgreSQL init error: %s", err)
	}

	logger.Infof("postgres connected, status: %#v", s.postgres.Stats())
}

func (s *Server) stopDB() {
	err := s.postgres.Close()
	if err != nil {
		logger.Errorf("can't close postgres properly, err: %s", err.Error())
	}

	logger.Info("postgres closed")
}

func (s *Server) startJaeger() {
	s.tracer = jaeger.Start(s.cfg.Jaeger)

	logger.Info("jaeger started")
}

func (s *Server) stopJaeger(ctx context.Context) {
	err := s.tracer.Stop(ctx)
	if err != nil {
		logger.Errorf("can't close jaeger properly, err: %s", err.Error())
	}

	logger.Info("jaeger closed")
}
