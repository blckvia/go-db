package app

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/blckvia/go-db/internal/database"
	"github.com/blckvia/go-db/internal/database/compute"
	"github.com/blckvia/go-db/internal/database/storage"
	e "github.com/blckvia/go-db/internal/database/storage/engine"
)

type App struct {
	Logger   *zap.Logger
	Database *database.Database
}

func NewApp(ctx context.Context, logger *zap.Logger) *App {
	if err := InitConfig(); err != nil {
		logger.Fatal("error initializing configs", zap.Error(err))
	}

	engine := e.NewEngine()
	c := compute.NewCompute(logger)
	s := storage.NewStorage(engine, logger)
	db := database.NewDatabase(c, s, logger)

	app := &App{
		Logger:   logger,
		Database: db,
	}

	app.run(ctx)
	return app
}

func (a *App) run(ctx context.Context) {
	scanner := bufio.NewScanner(os.Stdin)
	a.Logger.Info("Welcome! Enter commands (SET, GET, DEL):")

	for {
		if !scanner.Scan() {
			break
		}

		query := scanner.Text()
		res := a.Database.Handle(ctx, query)
		a.Logger.Info("Query result", zap.String("query", query), zap.String("result", res))
		fmt.Println(res)
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	return viper.ReadInConfig()
}
