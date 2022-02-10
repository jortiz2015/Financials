package main

import (
	"context"
	"fin/datasource"
	"fin/store"
	"fin/svc"
	"fin/svr"
	"log"
)

func main() {
	log := log.Default()
	ds := datasource.NewIEX(log)
	store, _ := store.NewDynamo(log)

	svc := svc.NewSvc(log, store, ds)
	is, err := svc.GetAnnualIncomeStatements(context.Background(), "AAPL", 1)
	if err != nil {
		log.Fatalf("Error getting Income Statements: %s", err)
	}
	log.Println(is)
	svr := svr.NewSvr(log, svc)
	svr.StartServer()
}
