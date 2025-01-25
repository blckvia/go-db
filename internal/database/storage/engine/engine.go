package engine

import "fmt"

//go:generate mockgen -source=engine.go -destination=../../mock/engine_mock.go -package=mock

type IEngine interface {
	Set(key, value string) error
	Get(key string) (string, error)
	Delete(key string) error
}

type Engine struct {
	data map[string]string
}

func NewEngine() *Engine {
	return &Engine{data: make(map[string]string)}
}

func (e *Engine) Set(key string, value string) error {
	e.data[key] = value
	return nil
}

func (e *Engine) Get(key string) (string, error) {
	if value, ok := e.data[key]; ok {
		return value, nil
	}

	return "", fmt.Errorf("key: %s not found", key)
}

func (e *Engine) Delete(key string) error {
	if _, ok := e.data[key]; ok {
		delete(e.data, key)
		return nil
	}

	return fmt.Errorf("key: %s not found", key)
}
