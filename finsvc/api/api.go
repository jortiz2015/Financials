package api

import (
	"context"
	"fin/model"
)

// Core API port for service.
type API interface {
	GetAnnualBalanceSheet(ctx context.Context, symbol string, limit int) ([]model.BalanceSheet, error)
}

// Interface to fetch data from a data source like IEX.
// This interface is used when new data is needed to be fetched
// from a data source then stored in a permanent storage.
type DataSource interface {
	GetAnnualBalanceSheet(ctx context.Context, symbol string, limit int) ([]model.BalanceSheet, error)
}

// Interface to fetch data from a data store like dynamoDB or in-memory cache.
type Store interface {
	GetAnnualBalanceSheet(ctx context.Context, symbol string, limit int) ([]model.BalanceSheet, error)
	InsertAnnualBalanceSheet(ctx context.Context, symbol string, bs *model.BalanceSheet) error
}
