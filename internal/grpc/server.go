package grpc

import (
	"context"

	newsv1 "github.com/supLano/go-grpc-proto/api/news/v1"
)

type Server struct {
	newsv1.UnimplementedNewsServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) CreateNews(ctx context.Context, req *newsv1.NewsRequest) (*newsv1.NewsResponse, error) {
	return &newsv1.NewsResponse{}, nil
}

func (s *Server) UpdateNews(ctx context.Context, req *newsv1.NewsRequest) (*newsv1.NewsResponse, error) {
	return &newsv1.NewsResponse{}, nil
}

func (s *Server) DeleteNews(ctx context.Context, req *newsv1.NewsRequest) (*newsv1.NewsResponse, error) {
	return &newsv1.NewsResponse{}, nil
}

func (s *Server) GetNews(ctx context.Context, req *newsv1.NewsRequest) (*newsv1.NewsResponse, error) {
	return &newsv1.NewsResponse{}, nil
}

func (s *Server) ListNews(ctx context.Context, req *newsv1.NewsRequest) (*newsv1.NewsResponse, error) {
	return &newsv1.NewsResponse{}, nil
}