package api

import (
	"context"
	"finsvc/model"
)

// Core API port for service.
type API interface {
	GetBalanceSheet(ctx context.Context, symbol string, limit int) (model.BalanceSheet, error)
}

// Interface to fetch data from a data source like IEX.
// This interface is used when new data is needed to be fetched
// from a data source then stored in a permanent storage.
type DataSource interface {
	GetBalanceSheet(ctx context.Context, symbol string, limit int) (model.BalanceSheet, error)
}

// Interface to fetch data from a data store like dynamoDB or in-memory cache.
type Store interface {
	GetBalanceSheet(ctx context.Context, symbol string, limit int) (model.BalanceSheet, error)
	InsertBalanceSheet(ctx context.Context, symbol string, bs model.BalanceSheet) error
}
