package main

import (
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
	svr := svr.NewSvr(log, svc)
	svr.StartServer()
}
