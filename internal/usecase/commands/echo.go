package commands

import (
	"fmt"
	"github.com/iAbbos/go-my_redis/internal/entity"
	"net"
)

func Echo(c net.Conn, args entity.Array) error {
	if len(args) != 1 {
		return fmt.Errorf("incorrect number of arguments: %+v", args)
	}
	echoed, ok := args[0].(entity.BulkString)
	if !ok {
		return fmt.Errorf("echoed should be string: %+v", args[0])
	}
	response := entity.SimpleString(echoed.Value).Encode()
	if _, err := c.Write([]byte(response)); err != nil {
		return fmt.Errorf("unable to write response buffer: %w", err)
	}
	return nil
}
