package api

import (
	"context"
	"fin/model"
)

// Core API port for service.
type API interface {
	GetAnnualBalanceSheets(ctx context.Context, symbol string, limit int) ([]model.BalanceSheet, error)
	GetAnnualIncomeStatements(ctx context.Context, symbol string, limit int) ([]model.IncomeStatement, error)
	GetAnnualCashFlows(ctx context.Context, symbol string, limit int) ([]model.CashFlow, error)
}

// Interface to fetch data from a data source like IEX.
// This interface is used when new data is needed to be fetched
// from a data source then stored in a permanent storage.
type DataSource interface {
	GetAnnualBalanceSheets(ctx context.Context, symbol string, limit int) ([]model.BalanceSheet, error)
	GetAnnualIncomeStatements(ctx context.Context, symbol string, limit int) ([]model.IncomeStatement, error)
	GetAnnualCashFlows(ctx context.Context, symbol string, limit int) ([]model.CashFlow, error)
}

// Interface to fetch data from a data store like dynamoDB or in-memory cache.
type Store interface {
	GetAnnualBalanceSheets(ctx context.Context, symbol string, limit int) ([]model.BalanceSheet, error)
	InsertAnnualBalanceSheet(ctx context.Context, symbol string, bs *model.BalanceSheet) error
	InsertAnnualBalanceSheets(ctx context.Context, symbol string, bs []model.BalanceSheet) error
	GetAnnualIncomeStatements(ctx context.Context, symbol string, limit int) ([]model.IncomeStatement, error)
	InsertAnnualIncomeStatement(ctx context.Context, symbol string, is *model.IncomeStatement) error
	InsertAnnualIncomeStatements(ctx context.Context, symbol string, is []model.IncomeStatement) error
	GetAnnualCashFlows(ctx context.Context, symbol string, limit int) ([]model.CashFlow, error)
	InsertAnnualCashFlow(ctx context.Context, symbol string, cf *model.CashFlow) error
	InsertAnnualCashFlows(ctx context.Context, symbol string, cf []model.CashFlow) error
}
