package database

import (
	"context"
	"errors"
	"fmt"

	"go.uber.org/zap"

	"github.com/blckvia/go-db/internal/database/compute"
	"github.com/blckvia/go-db/internal/database/storage"
)

type Database struct {
	compute compute.ICompute
	storage storage.IStorage
	logger  *zap.Logger
}

func NewDatabase(compute compute.ICompute, storage storage.IStorage, logger *zap.Logger) *Database {
	return &Database{
		compute: compute,
		storage: storage,
		logger:  logger,
	}
}

func (d *Database) Handle(ctx context.Context, cmd string) string {
	d.logger.Debug("Handling cmd", zap.String("query", cmd))
	query, err := d.compute.Compute(cmd)
	if err != nil {
		return fmt.Sprintf("error computing query: %s", err)
	}

	switch query.CommandID() {
	case compute.SetCommandID:
		return d.HandleSetCommand(ctx, query)
	case compute.GetCommandID:
		return d.HandleGetCommand(ctx, query)
	case compute.DelCommandID:
		return d.HandleDelCommand(ctx, query)
	default:
		d.logger.Error("Unknown command", zap.Int("cmdId", query.CommandID()))
		return fmt.Sprintf("unknown command: %v", query.CommandID())
	}
}

func (d *Database) HandleSetCommand(ctx context.Context, query compute.Query) string {
	args := query.Arguments()
	if err := d.storage.Set(ctx, args[0], args[1]); err != nil {
		return fmt.Sprintf("error setting value: %s", err)
	}

	return "OK"
}

func (d *Database) HandleGetCommand(ctx context.Context, query compute.Query) string {
	args := query.Arguments()
	val, err := d.storage.Get(ctx, args[0])
	if errors.Is(err, storage.ErrorNotFound) {
		return "not found"
	} else if err != nil {
		return fmt.Sprintf("error getting value: %s", err)
	}

	return fmt.Sprintf("value: %s", val)
}

func (d *Database) HandleDelCommand(ctx context.Context, query compute.Query) string {
	args := query.Arguments()
	if err := d.storage.Delete(ctx, args[0]); err != nil {
		return fmt.Sprintf("error deleting value: %s", err)
	}

	return "OK"
}
