package grpc

import (
	"context"
	"errors"

	"github.com/docker/distribution/uuid"
	newsv1 "github.com/supLano/go-grpc-proto/api/news/v1"
	"github.com/supLano/go-grpc-proto/memstore"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	newsv1.UnimplementedNewsServiceServer
	store *memstore.NewsMemStore
}

func NewServer(store *memstore.NewsMemStore) *Server {
	return &Server{
		store: store,
	}
}

func (s *Server) CreateNews(ctx context.Context, req *newsv1.CreateNewsRequest) (*newsv1.CreateNewsResponse, error) {
	news, err := ParseNewsRequest(req)
	if err != nil {
		return nil, err
	}
    err = s.store.CreateNews(ctx, news)
	if err != nil {
		return nil, err
	}
	response , error := ParseNewsResponse(news)
	if error != nil {	
		return nil, error
	}
	return response, nil
}

// func (s *Server) UpdateNews(ctx context.Context, req *newsv1.UpdateNewsRequest) (*newsv1.UpdateNewsResponse, error) {
// 	return ParseNewsResponse(news), nil
// }

// func (s *Server) DeleteNews(ctx context.Context, req *newsv1.DeleteNewsRequest) (*newsv1.DeleteNewsResponse, error) {
// 	return ParseNewsResponse(news), nil
// }

func (s *Server) GetNews(ctx context.Context, req *newsv1.GetNewsRequest) (*newsv1.GetNewsResponse, error) {
	parsedID, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, errors.New("news id is invalid")
	}
	news, err := s.store.GetNews(ctx, parsedID.String())
	if err != nil {
		return nil, err
	}
	response, err := ParseNewsResponse(news)
	if err != nil {
		return nil, err
	}
	return &newsv1.GetNewsResponse{
		Id: response.Id,
		Topic: response.Topic,
		Language: response.Language,
		Country: response.Country,
		Author: response.Author,
		Content: response.Content,
		Keywords: response.Keywords,
		CreatedAt: response.CreatedAt,
		UpdatedAt: response.UpdatedAt,
		DeletedAt: response.DeletedAt,
	}, nil
}

func (s *Server) ListNews(ctx context.Context, req *newsv1.ListNewsRequest) (*newsv1.NewsResponse, error) {
	return &newsv1.NewsResponse{}, nil
}

func ParseNewsRequest(req *newsv1.CreateNewsRequest) (*memstore.News, error) {
	if req == nil {
		return nil, errors.New("news request is nil")
	}

	parsedID, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, errors.New("news id is invalid")
	}

	return &memstore.News{
		ID:        parsedID,
		Topic:     req.GetTopic(),
		Language:  req.GetLanguage(),
		Country:   req.GetCountry(),
		Author:    req.GetAuthor(),
		Content:   req.GetContent(),
		Keywords:  req.GetKeywords(),
	}, nil
}

func ParseNewsResponse(news *memstore.News) (*newsv1.CreateNewsResponse, error) {
	if news == nil {
		return nil, errors.New("news is nil")
	}

	return &newsv1.CreateNewsResponse{
		Id: news.ID.String(),
		Topic: news.Topic,
		Language: news.Language,
		Country: news.Country,
		Author: news.Author,
		Content: news.Content,
		Keywords: news.Keywords,	
		CreatedAt: timestamppb.New(news.CreatedAt),
		UpdatedAt: timestamppb.New(news.UpdatedAt),
		DeletedAt: timestamppb.New(news.DeletedAt),	
	}, nil
}
