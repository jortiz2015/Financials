package fin

import (
	"context"
	pb "finsvc/finsvc/pb"
	"log"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Fin struct {
	log log.Logger
}

func NewFin(log *log.Logger) *Fin {
	return &Fin{log: *log}
}

func (f *Fin) GetBalanceSheet(ctx context.Context, ticker *pb.Ticker) (*pb.BalanceSheet, error) {
	f.log.Println("GetBalanceSheet called: ", "ticker", ticker.GetSymbol())

	return &pb.BalanceSheet{ReportDate: timestamppb.New(time.Now()), FillingType: "10k",
		FiscalDate: timestamppb.New(time.Now()), FiscalQuarter: 1, FiscalYear: 2022,
		Currency: "USD", CurrentCash: 420.69, ShortTermInvestments: 420.69, Receivables: 420.69,
		Inventory: 420.69, OtherCurrentAssets: 420.69, CurrentAssets: 420.69,
		LongTermInvestments: 420.69, PropertyPlantEquipment: 420.69, Goodwill: 420.69,
		IntangibleAssets: 420.69, OtherAssets: 420.69, TotalAssets: 420.69, AccountsPayable: 420.69,
		CurrentLongTermDebt: 420.69, OtherCurrentLiabilities: 420.69, TotalCurrentLiabilities: 420.69,
		LongTermDebt: 420.69, OtherLiablities: 420.69, MinorityInterest: 420.69,
		TotalLiabilities: 420.69, CommonStock: 420.69, RetainedEarnings: 420.69,
		TreasuryStock: 420.69, CapitalSurplus: 420.69, ShareholderEquity: 420.69,
		NetTangibleAssets: 420.69}, nil
}
