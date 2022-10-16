package api

import (
	"context"

	"devm/indexer/pkg/chain"
)

// StoreWriter
type StoreWriter interface {
	Ping() error
	HasScanned(ctx context.Context, id int64) bool
	SaveBlock(ctx context.Context, block *chain.Block) error
	SaveTx(ctx context.Context, tx *chain.Tx) error
}
