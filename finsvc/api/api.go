package api

import (
	"context"
	"fin/model"
)

// Core API port for service.
type API interface {
	GetAnnualBalanceSheets(ctx context.Context, symbol string, limit int) ([]model.BalanceSheet, error)
	GetQuarterlyBalanceSheets(ctx context.Context, symbol string, limit int) ([]model.BalanceSheet, error)
	GetAnnualIncomeStatements(ctx context.Context, symbol string, limit int) ([]model.IncomeStatement, error)
	GetQuarterlyIncomeStatements(ctx context.Context, symbol string, limit int) ([]model.IncomeStatement, error)
	GetAnnualCashFlows(ctx context.Context, symbol string, limit int) ([]model.CashFlow, error)
	GetQuarterlyCashFlows(ctx context.Context, symbol string, limit int) ([]model.CashFlow, error)
	GetAnnualFinancials(ctx context.Context, symbol string, limit int) (model.Financials, error)
	GetQuarterlyFinancials(ctx context.Context, symbol string, limit int) (model.Financials, error)
}

// Interface to fetch data from a data source like IEX.
// This interface is used when new data is needed to be fetched
// from a data source then stored in a permanent storage.
type DataSource interface {
	GetAnnualBalanceSheets(ctx context.Context, symbol string, limit int) ([]model.BalanceSheet, error)
	GetQuarterlyBalanceSheets(ctx context.Context, symbol string, limit int) ([]model.BalanceSheet, error)
	GetAnnualIncomeStatements(ctx context.Context, symbol string, limit int) ([]model.IncomeStatement, error)
	GetQuarterlyIncomeStatements(ctx context.Context, symbol string, limit int) ([]model.IncomeStatement, error)
	GetAnnualCashFlows(ctx context.Context, symbol string, limit int) ([]model.CashFlow, error)
	GetQuarterlyCashFlows(ctx context.Context, symbol string, limit int) ([]model.CashFlow, error)
}

// Interface to fetch data from a data store like dynamoDB or in-memory cache.
type Store interface {
	GetAnnualBalanceSheets(ctx context.Context, symbol string, limit int) ([]model.BalanceSheet, error)
	GetQuarterlyBalanceSheets(ctx context.Context, symbol string, limit int) ([]model.BalanceSheet, error)
	InsertBalanceSheet(ctx context.Context, symbol string, bs *model.BalanceSheet) error
	InsertBalanceSheets(ctx context.Context, symbol string, bs []model.BalanceSheet) error
	GetAnnualIncomeStatements(ctx context.Context, symbol string, limit int) ([]model.IncomeStatement, error)
	GetQuarterlyIncomeStatements(ctx context.Context, symbol string, limit int) ([]model.IncomeStatement, error)
	InsertIncomeStatement(ctx context.Context, symbol string, is *model.IncomeStatement, filing string) error
	InsertIncomeStatements(ctx context.Context, symbol string, is []model.IncomeStatement, filing string) error
	GetAnnualCashFlows(ctx context.Context, symbol string, limit int) ([]model.CashFlow, error)
	GetQuarterlyCashFlows(ctx context.Context, symbol string, limit int) ([]model.CashFlow, error)
	InsertCashFlow(ctx context.Context, symbol string, cf *model.CashFlow, filing string) error
	InsertCashFlows(ctx context.Context, symbol string, cf []model.CashFlow, filing string) error
}
