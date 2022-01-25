package main

import (
	"context"
	fin "finsvc/finsvc/fin"
	pb "finsvc/finsvc/pb"
	"fmt"
	"log"
	"net"
	"os"

	iex "github.com/goinvest/iexcloud/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	client := iex.NewClient("Tsk_18fb2a95ff9241c7a7fbb2ff127abc75", iex.WithBaseURL("https://sandbox.iexapis.com/stable"))

	bs, err := client.AnnualBalanceSheets(context.Background(), "aapl", 1000)
	if err != nil {
		log.Fatalf("Error getting balance sheets: %s", err)
	}
	fmt.Println("Number of Balance sheets: ", len(bs.Statements))
	fmt.Println("Balance sheets: ", bs.Statements[0].Currency)
	fmt.Println("Balance sheets Report Date: ", bs.Statements[0].ReportDate)

	q, err := client.Quote(context.Background(), "aapl")
	if err != nil {
		log.Fatalf("Error getting quote: %s", err)
	}
	fmt.Println("Company Name: ", q.CompanyName)
	fmt.Println("Quote: ", q.Close)
	aapl := pb.Ticker{Symbol: "aapl"}
	fmt.Println("Ticker: ", aapl.GetSymbol())

	log := log.Default()
	svr := grpc.NewServer()
	fsvc := fin.NewFin(log)
	pb.RegisterFinSvcServer(svr, fsvc)

	reflection.Register(svr)

	l, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("Error listening: %s", err)
		os.Exit(1)
	}
	svr.Serve(l)
}
