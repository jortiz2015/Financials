package svc

import (
	"context"
	"fin/api"
	"fin/model"
	"log"
)

// svc implements the API interface
type Svc struct {
	log   *log.Logger
	store api.Store
	ds    api.DataSource
}

// Creates a new service to use API.
func NewSvc(l *log.Logger, store api.Store, ds api.DataSource) api.API {
	return &Svc{log: l, store: store, ds: ds}
}

func (svc *Svc) GetAnnualBalanceSheets(ctx context.Context, symbol string, limit int) ([]model.BalanceSheet, error) {
	bs, err := svc.store.GetAnnualBalanceSheets(ctx, symbol, limit)
	if err != nil {
		svc.log.Printf("Error getting balance sheets: %s", err)
		return []model.BalanceSheet{}, err
	}
	svc.log.Printf("Got %d balance sheets", len(bs))
	// Return balance sheet if store has it.
	if len(bs) > 0 {
		return bs, nil
	}

	// When there is no record in the store,
	// we need to try to fetch it from the data source.
	bs, err = svc.ds.GetAnnualBalanceSheets(ctx, symbol, limit)
	if err != nil {
		svc.log.Fatalf("Error getting balance sheets: %s", err)
		return []model.BalanceSheet{}, err
	}

	// If no balance sheet is found, return an empty balance sheet.
	if len(bs) == 0 {
		svc.log.Printf("No balance sheet found for %s", symbol)
		return []model.BalanceSheet{}, nil
	}

	svc.store.InsertBalanceSheets(ctx, symbol, bs)
	if err != nil {
		svc.log.Fatalf("Error inserting balance sheet to store: %s", err)
	}

	return bs, nil
}

func (svc *Svc) GetQuarterlyBalanceSheets(ctx context.Context, symbol string, limit int) ([]model.BalanceSheet, error) {
	bs, err := svc.store.GetQuarterlyBalanceSheets(ctx, symbol, limit)
	if err != nil {
		svc.log.Printf("Error getting balance sheets: %s", err)
		return []model.BalanceSheet{}, err
	}
	svc.log.Printf("Got %d balance sheets", len(bs))
	// Return balance sheet if store has it.
	if len(bs) > 0 {
		return bs, nil
	}

	// When there is no record in the store,
	// we need to try to fetch it from the data source.
	bs, err = svc.ds.GetQuarterlyBalanceSheets(ctx, symbol, limit)
	if err != nil {
		svc.log.Fatalf("Error getting balance sheets: %s", err)
		return []model.BalanceSheet{}, err
	}

	// If no balance sheet is found, return an empty balance sheet.
	if len(bs) == 0 {
		svc.log.Printf("No balance sheet found for %s", symbol)
		return []model.BalanceSheet{}, nil
	}

	svc.store.InsertBalanceSheets(ctx, symbol, bs)
	if err != nil {
		svc.log.Fatalf("Error inserting balance sheet to store: %s", err)
	}

	return bs, nil
}

func (svc *Svc) GetAnnualIncomeStatements(ctx context.Context, symbol string, limit int) ([]model.IncomeStatement, error) {
	is, err := svc.store.GetAnnualIncomeStatements(ctx, symbol, limit)
	if err != nil {
		svc.log.Printf("Error getting Income Statements: %s", err)
		return []model.IncomeStatement{}, err
	}
	svc.log.Printf("Got %d Income Statements", len(is))
	// Return Income Statements if store has it.
	if len(is) > 0 {
		return is, nil
	}

	// When there is no record in the store,
	// we need to try to fetch it from the data source.
	is, err = svc.ds.GetAnnualIncomeStatements(ctx, symbol, limit)
	if err != nil {
		svc.log.Fatalf("Error getting Income Statements: %s", err)
		return []model.IncomeStatement{}, err
	}

	// If no Income Statement is found, return an empty Income Statement.
	if len(is) == 0 {
		svc.log.Printf("No Income Statement found for %s", symbol)
		return []model.IncomeStatement{}, nil
	}

	svc.store.InsertIncomeStatements(ctx, symbol, is, "10-K")
	if err != nil {
		svc.log.Fatalf("Error inserting Income Statements to store: %s", err)
	}

	return is, nil
}

func (svc *Svc) GetQuarterlyIncomeStatements(ctx context.Context, symbol string, limit int) ([]model.IncomeStatement, error) {
	is, err := svc.store.GetQuarterlyIncomeStatements(ctx, symbol, limit)
	if err != nil {
		svc.log.Printf("Error getting Income Statements: %s", err)
		return []model.IncomeStatement{}, err
	}
	svc.log.Printf("Got %d Income Statements", len(is))
	// Return Income Statements if store has it.
	if len(is) > 0 {
		return is, nil
	}

	// When there is no record in the store,
	// we need to try to fetch it from the data source.
	is, err = svc.ds.GetQuarterlyIncomeStatements(ctx, symbol, limit)
	if err != nil {
		svc.log.Fatalf("Error getting Income Statements: %s", err)
		return []model.IncomeStatement{}, err
	}

	// If no Income Statement is found, return an empty Income Statement.
	if len(is) == 0 {
		svc.log.Printf("No Income Statement found for %s", symbol)
		return []model.IncomeStatement{}, nil
	}

	svc.store.InsertIncomeStatements(ctx, symbol, is, "10-Q")
	if err != nil {
		svc.log.Fatalf("Error inserting Income Statements to store: %s", err)
	}

	return is, nil
}

func (svc *Svc) GetAnnualCashFlows(ctx context.Context, symbol string, limit int) ([]model.CashFlow, error) {
	cf, err := svc.store.GetAnnualCashFlows(ctx, symbol, limit)
	if err != nil {
		svc.log.Printf("Error getting Cash Flows: %s", err)
		return []model.CashFlow{}, err
	}
	svc.log.Printf("Got %d Cash Flows", len(cf))
	// Return Cash Flows if store has it.
	if len(cf) > 0 {
		return cf, nil
	}

	// When there is no record in the store,
	// we need to try to fetch it from the data source.
	cf, err = svc.ds.GetAnnualCashFlows(ctx, symbol, limit)
	if err != nil {
		svc.log.Fatalf("Error getting Cash Flows: %s", err)
		return []model.CashFlow{}, err
	}

	// If no Cash Flow is found, return an empty Cash Flow.
	if len(cf) == 0 {
		svc.log.Printf("No Cash Flows found for %s", symbol)
		return []model.CashFlow{}, nil
	}

	svc.store.InsertCashFlows(ctx, symbol, cf, "10-K")
	if err != nil {
		svc.log.Fatalf("Error inserting Cash Flows to store: %s", err)
	}

	return cf, nil
}

func (svc *Svc) GetQuarterlyCashFlows(ctx context.Context, symbol string, limit int) ([]model.CashFlow, error) {
	cf, err := svc.store.GetQuarterlyCashFlows(ctx, symbol, limit)
	if err != nil {
		svc.log.Printf("Error getting Cash Flows: %s", err)
		return []model.CashFlow{}, err
	}
	svc.log.Printf("Got %d Cash Flows", len(cf))
	// Return Cash Flows if store has it.
	if len(cf) > 0 {
		return cf, nil
	}

	// When there is no record in the store,
	// we need to try to fetch it from the data source.
	cf, err = svc.ds.GetQuarterlyCashFlows(ctx, symbol, limit)
	if err != nil {
		svc.log.Fatalf("Error getting Cash Flows: %s", err)
		return []model.CashFlow{}, err
	}

	// If no Cash Flow is found, return an empty Cash Flow.
	if len(cf) == 0 {
		svc.log.Printf("No Cash Flows found for %s", symbol)
		return []model.CashFlow{}, nil
	}

	svc.store.InsertCashFlows(ctx, symbol, cf, "10-Q")
	if err != nil {
		svc.log.Fatalf("Error inserting Cash Flows to store: %s", err)
	}

	return cf, nil
}
