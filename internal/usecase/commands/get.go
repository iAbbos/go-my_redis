package commands

import (
	"fmt"
	"github.com/iAbbos/go-my_redis/internal/entity"
	"github.com/iAbbos/go-my_redis/internal/pkg/storage/cache"
	"net"
)

func Get(c net.Conn, args entity.Array) error {
	if len(args) != 1 {
		return fmt.Errorf("incorrect number of arguments: %+v", args)
	}
	key, ok := args[0].(entity.BulkString)
	if !ok {
		return fmt.Errorf("key should be string: %+v", args[0])
	}
	response := entity.BulkString{IsNull: true}.Encode()
	if val, ok := cache.Get(key.Value); ok {
		response = entity.BulkString{
			Value:  val,
			IsNull: false,
		}.Encode()
	}
	if _, err := c.Write([]byte(response)); err != nil {
		return fmt.Errorf("unable to write response buffer: %w", err)
	}
	return nil
}
