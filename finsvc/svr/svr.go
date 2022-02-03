package svr

import (
	"context"
	"fin/api"
	"fin/model"
	pb "fin/pb"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Svr struct {
	log *log.Logger
	svc api.API
}

func NewSvr(l *log.Logger, api api.API) *Svr {
	return &Svr{log: l, svc: api}
}

func (svr *Svr) GetAnnualBalanceSheets(ctx context.Context, request *pb.GetRequest) (*pb.BalanceSheets, error) {
	bs := pb.BalanceSheets{}
	symbol := request.GetSymbol()
	limit := int(request.GetLimit())

	mbs, err := svr.svc.GetAnnualBalanceSheets(ctx, symbol, limit)
	if err != nil {
		svr.log.Printf("Error fetching balance sheets from service: %s", err)
	}

	if len(mbs) == 0 {
		return &bs, nil
	}

	for _, mb := range mbs {
		bs.BalanceSheets = append(bs.BalanceSheets, svr.GetBalanceSheet(mb))
	}

	return &bs, nil
}

func (svc *Svr) GetBalanceSheet(bs model.BalanceSheet) *pb.BalanceSheet {
	b := pb.BalanceSheet{}
	b.ReportDate = timestamppb.New(bs.ReportDate)
	b.FilingType = bs.FilingType
	b.FiscalDate = timestamppb.New(bs.FiscalDate)
	b.FiscalQuarter = int32(bs.FiscalQuarter)
	b.FiscalYear = int32(bs.FiscalYear)
	b.Currency = bs.Currency
	b.CurrentCash = bs.CurrentCash
	b.ShortTermInvestments = bs.ShortTermInvestments
	b.Receivables = bs.Receivables
	b.Inventory = bs.Inventory
	b.OtherCurrentAssets = bs.OtherCurrentAssets
	b.CurrentAssets = bs.CurrentAssets
	b.LongTermInvestments = bs.LongTermInvestments
	b.PropertyPlantEquipment = bs.PropertyPlantEquipment
	b.Goodwill = bs.Goodwill
	b.IntangibleAssets = bs.IntangibleAssets
	b.OtherAssets = bs.OtherAssets
	b.TotalAssets = bs.TotalAssets
	b.AccountsPayable = bs.AccountsPayable
	b.CurrentLongTermDebt = bs.CurrentLongTermDebt
	b.OtherCurrentLiabilities = bs.OtherCurrentLiabilities
	b.TotalCurrentLiabilities = bs.TotalCurrentLiabilities
	b.LongTermDebt = bs.LongTermDebt
	b.OtherLiabilities = bs.OtherLiabilities
	b.MinorityInterest = bs.MinorityInterest
	b.TotalLiabilities = bs.TotalLiabilities
	b.CommonStock = bs.CommonStock
	b.RetainedEarnings = bs.RetainedEarnings
	b.TreasuryStock = bs.TreasuryStock
	b.CapitalSurplus = bs.CapitalSurplus
	b.ShareholderEquity = bs.ShareholderEquity
	b.NetTangibleAssets = bs.NetTangibleAssets

	return &b
}

func (svr *Svr) StartServer() {
	svr.log.Printf("Starting server...")

	gRPC := grpc.NewServer()
	pb.RegisterFinSvcServer(gRPC, svr)

	reflection.Register(gRPC)

	l, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("Error listening: %s", err)
		os.Exit(1)
	}
	gRPC.Serve(l)
}
