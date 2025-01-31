package compute

import (
	"errors"
	"strings"

	"go.uber.org/zap"
)

var (
	errInvalidQuery     = errors.New("empty command")
	errInvalidCommand   = errors.New("invalid command")
	errInvalidArguments = errors.New("invalid arguments")
)

//go:generate mockgen -source=compute.go -destination=../../mock/compute_mock.go -package=mock

type ICompute interface {
	Compute(command string) (Query, error)
}

type Compute struct {
	logger *zap.Logger
}

func NewCompute(logger *zap.Logger) (*Compute, error) {
	if logger == nil {
		return nil, errors.New("logger is invalid")
	}
	return &Compute{
		logger: logger,
	}, nil
}

func (c *Compute) Compute(cmd string) (Query, error) {
	parts := strings.Fields(cmd)
	if len(parts) == 0 {
		c.logger.Debug("empty command", zap.String("cmd", cmd))
		return Query{}, errInvalidQuery
	}

	command := parts[0]
	cmdId := commandNameToCommandID(command)

	if cmdId == UnknownCommandID {
		c.logger.Debug("unknown command", zap.String("cmd", cmd))
		return Query{}, errInvalidCommand
	}

	query := NewQuery(cmdId, parts[1:])
	argumentNumber := commandArgumentsNumber(cmdId)
	if len(query.arguments) != argumentNumber {
		c.logger.Debug("invalid arguments", zap.String("cmd", cmd))
		return Query{}, errInvalidArguments
	}

	return query, nil
}
