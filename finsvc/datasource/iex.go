package datasource

import (
	"context"
	"fin/api"
	"fin/model"
	"fmt"
	"log"
	"os"
	"time"

	iex "github.com/goinvest/iexcloud/v2"
)

type IEX struct {
	log    log.Logger
	client *iex.Client
}

func NewIEX(log *log.Logger) api.DataSource {
	client := iex.NewClient(os.Getenv("IEX_KEY"), iex.WithBaseURL("https://sandbox.iexapis.com/stable"))
	return &IEX{log: *log, client: client}
}

func (i *IEX) GetAnnualBalanceSheets(ctx context.Context, symbol string, limit int) ([]model.BalanceSheet, error) {
	response, err := i.client.AnnualBalanceSheets(context.Background(), symbol, limit)
	if err != nil {
		log.Fatalf("Error getting balance sheets: %s", err)
	}
	bss := response.Statements
	balanceSheets := []model.BalanceSheet{}

	for _, b := range bss {
		bs := model.BalanceSheet{}
		bs.ReportDate = time.Time(b.ReportDate)
		bs.FilingType = b.FilingType
		bs.FiscalDate = time.Time(b.FiscalDate)
		bs.FiscalQuarter = b.FiscalQuarter
		bs.FiscalYear = b.FiscalYear
		bs.Currency = b.Currency
		bs.CurrentCash = b.CurrentCash
		bs.ShortTermInvestments = b.ShortTermInvestments
		bs.Receivables = b.Receivables
		bs.Inventory = b.Inventory
		bs.OtherCurrentAssets = b.OtherCurrentAssets
		bs.CurrentAssets = b.CurrentAssets
		bs.LongTermInvestments = b.LongTermInvestments
		bs.PropertyPlantEquipment = b.PropertyPlantEquipment
		bs.Goodwill = b.Goodwill
		bs.IntangibleAssets = b.IntangibleAssets
		bs.OtherAssets = b.OtherAssets
		bs.TotalAssets = b.TotalAssets
		bs.AccountsPayable = b.AccountsPayable
		bs.CurrentLongTermDebt = b.CurrentLongTermDebt
		bs.OtherCurrentLiabilities = b.OtherCurrentLiabilities
		bs.TotalCurrentLiabilities = b.TotalCurrentLiabilities
		bs.LongTermDebt = b.LongTermDebt
		bs.OtherLiabilities = b.OtherLiablities
		bs.MinorityInterest = b.MinorityInterest
		bs.TotalLiabilities = b.TotalLiabilities
		bs.CommonStock = b.CommonStock
		bs.RetainedEarnings = b.RetainedEarnings
		bs.TreasuryStock = b.TreasuryStock
		bs.CapitalSurplus = b.CapitalSurplus
		bs.ShareholderEquity = b.ShareholderEquity
		bs.NetTangibleAssets = b.NetTangibleAssets
		balanceSheets = append(balanceSheets, bs)
	}

	return balanceSheets, nil
}

func (i *IEX) GetAnnualIncomeStatements(ctx context.Context, symbol string, limit int) ([]model.IncomeStatement, error) {
	client := iex.NewClient(os.Getenv("IEX_KEY"), iex.WithBaseURL("https://sandbox.iexapis.com/stable"))
	response, err := client.AnnualIncomeStatements(context.Background(), symbol, limit)
	if err != nil {
		log.Fatalf("Error getting Income Statements: %s", err)
	}

	i.log.Printf("RESPONSE: %+v\tERROR: %+v", response, err)

	iss := response.Statements
	incomeStatements := []model.IncomeStatement{}

	for _, i := range iss {
		is := model.IncomeStatement{}
		is.ReportDate = time.Time(i.ReportDate)
		is.FiscalDate = time.Time(i.FiscalDate)
		is.Currency = i.Currency
		is.TotalRevenue = i.TotalRevenue
		is.CostOfRevenue = i.CostOfRevenue
		is.GrossProfit = i.GrossProfit
		is.ResearchAndDevelopment = i.ResearchAndDevelopment
		is.SellingGeneralAndAdmin = i.SellingGeneralAndAdmin
		is.OperatingExpense = i.OperatingExpense
		is.OperatingIncome = i.OperatingIncome
		is.OtherIncomeExpenseNet = i.OtherIncomeExpenseNet
		is.EBIT = i.EBIT
		is.InterestIncome = i.InterestIncome
		is.PretaxIncome = i.PretaxIncome
		is.IncomeTax = i.IncomeTax
		is.MinorityInterest = i.MinorityInterest
		is.NetIncome = i.NetIncome
		is.NetIncomeBasic = i.NetIncomeBasic
		incomeStatements = append(incomeStatements, is)
	}

	return incomeStatements, nil
}

func (i *IEX) GetQuote(symbol string) (iex.Quote, error) {
	q, err := i.client.Quote(context.Background(), "aapl")
	if err != nil {
		log.Fatalf("Error getting quote: %s", err)
	}
	fmt.Println("Company Name: ", q.CompanyName)
	fmt.Println("Quote: ", q.Close)

	return q, nil
}
