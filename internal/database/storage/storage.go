package storage

import (
	"context"
	"errors"

	"go.uber.org/zap"

	"github.com/blckvia/go-db/internal/database/storage/engine"
)

//go:generate mockgen -source=storage.go -destination=../../mock/storage_mock.go -package=mock

var (
	ErrorNotFound = errors.New("not found")
)

type IStorage interface {
	Set(ctx context.Context, key, value string) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
}

type Storage struct {
	engine engine.IEngine
	logger *zap.Logger
}

func NewStorage(engine engine.IEngine, logger *zap.Logger) *Storage {
	return &Storage{engine: engine, logger: logger}
}

func (s *Storage) Set(_ context.Context, key, value string) error {
	return s.engine.Set(key, value)
}

func (s *Storage) Get(_ context.Context, key string) (string, error) {
	return s.engine.Get(key)
}

func (s *Storage) Delete(_ context.Context, key string) error {
	return s.engine.Delete(key)
}
