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

	svc.store.InsertAnnualBalanceSheets(ctx, symbol, bs)
	if err != nil {
		svc.log.Fatalf("Error inserting balance sheet to store: %s", err)
	}

	return bs, nil
}
