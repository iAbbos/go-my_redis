package app

import (
	tcpserver "github.com/iAbbos/go-my_redis/internal/delivery/tcp/server"
	configpkg "github.com/iAbbos/go-my_redis/internal/pkg/config"
)

type App struct {
	Config *configpkg.Config
}

func NewApp(config *configpkg.Config) (*App, error) {
	return &App{
		Config: config,
	}, nil
}

func (a *App) Run() error {
	server := tcpserver.NewServer(a.Config)
	err := server.Run()
	if err != nil {
		return err
	}
	return nil
}
