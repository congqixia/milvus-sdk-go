package client

import (
	"context"
	"errors"
	"fmt"

	"github.com/milvus-io/milvus-sdk-go/v2/entity"
)

func (c *grpcClient) InsertByRows(ctx context.Context, rows []entity.Row) error {
	if c.service == nil {
		return ErrClientNotReady
	}
	if len(rows) == 0 {
		return errors.New("empty rows provided")
	}
	sch, err := entity.ParseSchema(rows[0])
	if err != nil {
		return err
	}

	has, err := c.HasCollection(ctx, sch.CollectionName)
	if err != nil {
		return err
	}
	if !has {
		return fmt.Errorf("collection %s  exist", sch.CollectionName)
	}

	return nil
}
