package v1

import (
	"context"
	"fmt"
	"github.com/timickb/link-shortener/internal/config"
	"github.com/timickb/link-shortener/internal/delivery/grpc/v1/shortener"
	"github.com/timickb/link-shortener/internal/interfaces"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	shortener.UnimplementedShortenerServiceServer

	ctx       context.Context
	cfg       *config.AppConfig
	log       interfaces.Logger
	shortener interfaces.Shortener
}

func New(ctx context.Context,
	log interfaces.Logger,
	cfg *config.AppConfig,
	short interfaces.Shortener) *Server {
	return &Server{
		ctx:       ctx,
		cfg:       cfg,
		log:       log,
		shortener: short,
	}
}

func (s *Server) Run() error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", s.cfg.RPCPort))
	if err != nil {
		return err
	}

	srv := grpc.NewServer()
	shortener.RegisterShortenerServiceServer(srv, s)

	s.log.Infof("Listening RPC server on port %d", s.cfg.RPCPort)
	if err := srv.Serve(listen); err != nil {
		s.log.Errorf("Error serving RPC: %w", err)
		return err
	}

	return nil
}

func (s *Server) CreateLink(ctx context.Context, request *shortener.CreateShorteningRequest) (*shortener.CreateShorteningResponse, error) {
	s.log.Info("RPC CreateLink: request = ", request)
	short, err := s.shortener.CreateLink(request.GetUrl())
	if err != nil {
		s.log.Info("RPC CreateLink error: %w", err)
		return nil, err
	}

	return &shortener.CreateShorteningResponse{Short: short}, nil
}

func (s *Server) RestoreLink(ctx context.Context, request *shortener.RestoreRequest) (*shortener.RestoreResponse, error) {
	s.log.Info("RPC RestoreLink: request = ", request)
	original, err := s.shortener.RestoreLink(request.GetShort())
	if err != nil {
		s.log.Info("RPC RestoreLink error: %w", err)
		return nil, err
	}
	return &shortener.RestoreResponse{Original: original}, nil
}
