package memstore

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/docker/distribution/uuid"
)


type News struct {
	ID uuid.UUID
	Topic string
	Language string
	Country string
	Author string
	Content string
	Keywords []string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type NewsStore interface {
	CreateNews(ctx context.Context, news *News) error
	UpdateNews(ctx context.Context, news *News) []*News
	DeleteNews(ctx context.Context, id string) error
	GetNews(ctx context.Context, id string) (*News, error)
	ListNews(ctx context.Context) []*News	
}

type NewsMemStore struct {
	lock sync.RWMutex
	news map[string]*News
}

func NewNewsMemStore() *NewsMemStore {
	return &NewsMemStore{
		lock: sync.RWMutex{},
		news: make(map[string]*News),
	}
}

func (s *NewsMemStore) CreateNews(ctx context.Context, news *News) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.news[news.ID.String()] = &News{
		ID: news.ID,
		Topic: news.Topic,
		Language: news.Language,
		Country: news.Country,
		Author: news.Author,
		Content: news.Content,
		Keywords: news.Keywords,
		CreatedAt: news.CreatedAt,
		UpdatedAt: news.UpdatedAt,
		DeletedAt: news.DeletedAt,
	}
	return nil
}

func (s *NewsMemStore) UpdateNews(ctx context.Context, news *News) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.news[news.ID.String()] = &News{
		ID: news.ID,
		Topic: news.Topic,
		Language: news.Language,
		Country: news.Country,
		Author: news.Author,
		Content: news.Content,
		Keywords: news.Keywords,
		CreatedAt: news.CreatedAt,
		UpdatedAt: news.UpdatedAt,
		DeletedAt: news.DeletedAt,
	}
	return nil
}

func (s *NewsMemStore) DeleteNews(ctx context.Context, id string) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.news, id)
	return nil
}

func (s *NewsMemStore) NewsExists(ctx context.Context, news *News) error {
	s.lock.RLock()
	defer s.lock.RUnlock()
	_, ok := s.news[news.ID.String()]
	if !ok {
		return errors.New("news not found")
	}
	return nil
}

func (s *NewsMemStore) GetNews(ctx context.Context, id string) (*News, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	news, ok := s.news[id]
	if !ok {
		return nil, errors.New("news not found")
	}
	return news, nil
}

func (s *NewsMemStore) ListNews(ctx context.Context) []*News {
	s.lock.RLock()
	defer s.lock.RUnlock()
	var newsList []*News
	for _, news := range s.news {
		newsList = append(newsList, &News{
			ID: news.ID,
			Topic: news.Topic,
			Language: news.Language,
			Country: news.Country,
			Author: news.Author,
			Content: news.Content,
			Keywords: news.Keywords,
			CreatedAt: news.CreatedAt,
			UpdatedAt: news.UpdatedAt,
		})
	}
	return newsList
}
