package server

import (
	"errors"
	"fmt"
	"github.com/iAbbos/go-my_redis/internal/delivery/tcp/handler"
	"github.com/iAbbos/go-my_redis/internal/pkg/config"
	"io"
	"net"
	"os"
	"syscall"
)

type Server struct {
	Config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		Config: config,
	}
}

func (s *Server) Run() error {
	l, err := net.Listen("tcp", s.Config.Server.Host+s.Config.Server.Port)
	if err != nil {
		return fmt.Errorf("tsp server error on listen: %s %w", s.Config.Server.Port, err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			return fmt.Errorf("tcp server error on accept: %w", err)
		}
		go func() {
			fmt.Println("Accepted a new connection, handling now...")
			if err := handler.HandleConnection(conn); err != nil {
				switch {
				case errors.Is(err, net.ErrClosed),
					errors.Is(err, io.EOF),
					errors.Is(err, syscall.EPIPE):
					fmt.Println("Connect closing...")
				default:
					fmt.Println("error handling connection: ", err.Error())
					os.Exit(1)
				}
			}
		}()
	}
}
