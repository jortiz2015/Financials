package model

import (
	"time"
)

// Main BalanceSheet model.
type BalanceSheet struct {
	ReportDate              time.Time `json:"reportDate" dynamodbav:"reportDate"`
	FilingType              string    `json:"filingType" dynamodbav:"filingType"`
	FiscalDate              time.Time `json:"fiscalDate" dynamodbav:"fiscalDate"`
	FiscalQuarter           int       `json:"fiscalQuarter" dynamodbav:"fiscalQuarter"`
	FiscalYear              int       `json:"fiscalYear" dynamodbav:"fiscalYear"`
	Currency                string    `json:"currency" dynamodbav:"currency"`
	CurrentCash             float64   `json:"currentCash" dynamodbav:"currentCash"`
	ShortTermInvestments    float64   `json:"shortTermInvestments" dynamodbav:"shortTermInvestments"`
	Receivables             float64   `json:"receivables" dynamodbav:"receivables"`
	Inventory               float64   `json:"inventory" dynamodbav:"inventory"`
	OtherCurrentAssets      float64   `json:"otherCurrentAssets" dynamodbav:"otherCurrentAssets"`
	CurrentAssets           float64   `json:"currentAssets" dynamodbav:"currentAssets"`
	LongTermInvestments     float64   `json:"longTermInvestments" dynamodbav:"longTermInvestments"`
	PropertyPlantEquipment  float64   `json:"propertyPlantEquipment" dynamodbav:"propertyPlantEquipment"`
	Goodwill                float64   `json:"goodwill" dynamodbav:"goodwill"`
	IntangibleAssets        float64   `json:"intangibleAssets" dynamodbav:"intangibleAssets"`
	OtherAssets             float64   `json:"otherAssets" dynamodbav:"otherAssets"`
	TotalAssets             float64   `json:"totalAssets" dynamodbav:"totalAssets"`
	AccountsPayable         float64   `json:"accountsPayable" dynamodbav:"accountsPayable"`
	CurrentLongTermDebt     float64   `json:"currentLongTermDebt" dynamodbav:"currentLongTermDebt"`
	OtherCurrentLiabilities float64   `json:"otherCurrentLiabilities" dynamodbav:"otherCurrentLiabilities"`
	TotalCurrentLiabilities float64   `json:"totalCurrentLiabilities" dynamodbav:"totalCurrentLiabilities"`
	LongTermDebt            float64   `json:"longTermDebt" dynamodbav:"longTermDebt"`
	OtherLiabilities        float64   `json:"otherLiabilities" dynamodbav:"otherLiabilities"`
	MinorityInterest        float64   `json:"minorityInterest" dynamodbav:"minorityInterest"`
	TotalLiabilities        float64   `json:"totalLiabilities" dynamodbav:"totalLiabilities"`
	CommonStock             float64   `json:"commonStock" dynamodbav:"commonStock"`
	RetainedEarnings        float64   `json:"retainedEarnings" dynamodbav:"retainedEarnings"`
	TreasuryStock           float64   `json:"treasuryStock" dynamodbav:"treasuryStock"`
	CapitalSurplus          float64   `json:"capitalSurplus" dynamodbav:"capitalSurplus"`
	ShareholderEquity       float64   `json:"shareholderEquity" dynamodbav:"shareholderEquity"`
	NetTangibleAssets       float64   `json:"netTangibleAssets" dynamodbav:"netTangibleAssets"`
}

// IncomeStatement models one income statement.
type IncomeStatement struct {
	ReportDate             time.Time `json:"reportDate" dynamodbav:"reportDate"`
	FiscalDate             time.Time `json:"fiscalDate" dynamodbav:"fiscalDate"`
	Currency               string    `json:"currency" dynamodbav:"currency"`
	TotalRevenue           float64   `json:"totalRevenue" dynamodbav:"totalRevenue"`
	CostOfRevenue          float64   `json:"costOfRevenue" dynamodbav:"costOfRevenue"`
	GrossProfit            float64   `json:"grossProfit" dynamodbav:"grossProfit"`
	ResearchAndDevelopment float64   `json:"researchAndDevelopment" dynamodbav:"researchAndDevelopment"`
	SellingGeneralAndAdmin float64   `json:"sellingGeneralAndAdmin" dynamodbav:"sellingGeneralAndAdmin"`
	OperatingExpense       float64   `json:"operatingExpense" dynamodbav:"operatingExpense"`
	OperatingIncome        float64   `json:"operatingIncome" dynamodbav:"operatingIncome"`
	OtherIncomeExpenseNet  float64   `json:"otherIncomeExpenseNet" dynamodbav:"otherIncomeExpenseNet"`
	EBIT                   float64   `json:"ebit" dynamodbav:"ebit"`
	InterestIncome         float64   `json:"interestIncome" dynamodbav:"interestIncome"`
	PretaxIncome           float64   `json:"pretaxIncome" dynamodbav:"pretaxIncome"`
	IncomeTax              float64   `json:"incomeTax" dynamodbav:"incomeTax"`
	MinorityInterest       float64   `json:"minorityInterest" dynamodbav:"minorityInterest"`
	NetIncome              float64   `json:"netIncome" dynamodbav:"netIncome"`
	NetIncomeBasic         float64   `json:"netIncomeBasic" dynamodbav:"netIncomeBasic"`
}

// CashFlow models one cash flow statement.
type CashFlow struct {
	ReportDate              time.Time `json:"reportDate" dynamodbav:"reportDate"`
	FiscalDate              time.Time `json:"fiscalDate" dynamodbav:"fiscalDate"`
	Currency                string    `json:"currency" dynamodbav:"currency"`
	NetIncome               float64   `json:"netIncome" dynamodbav:"netIncome"`
	Depreciation            float64   `json:"depreciation" dynamodbav:"depreciation"`
	ChangesInReceivables    float64   `json:"changesInReceivables" dynamodbav:"changesInReceivables"`
	ChangesInInventories    float64   `json:"changesInInventories" dynamodbav:"changesInInventories"`
	CashChange              float64   `json:"cashChange" dynamodbav:"cashChange"`
	CashFlow                float64   `json:"cashFlow" dynamodbav:"cashFlow"`
	CapitalExpenditures     float64   `json:"capitalExpenditures" dynamodbav:"capitalExpenditures"`
	Investments             float64   `json:"investments" dynamodbav:"investments"`
	InvestingActivityOther  float64   `json:"investingActivityOther" dynamodbav:"investingActivityOther"`
	TotalInvestingCashFlows float64   `json:"totalInvestingCashFlows" dynamodbav:"totalInvestingCashFlows"`
	DividendsPaid           float64   `json:"dividendsPaid" dynamodbav:"dividendsPaid"`
	NetBorrowings           float64   `json:"netBorrowings" dynamodbav:"netBorrowings"`
	OtherFinancingCashFlows float64   `json:"otherFinancingCashFlows" dynamodbav:"otherFinancingCashFlows"`
	CashFlowFinancing       float64   `json:"cashFlowFinancing" dynamodbav:"cashFlowFinancing"`
	ExchangeRateEffect      float64   `json:"exchangeRateEffect" dynamodbav:"exchangeRateEffect"`
}

type Financials struct {
	BalanceSheets    []BalanceSheet    `json:"balanceSheets" dynamodbav:"balanceSheets"`
	CashFlows        []CashFlow        `json:"cashFlows" dynamodbav:"cashFlows"`
	IncomeStatements []IncomeStatement `json:"incomeStatements" dynamodbav:"incomeStatements"`
}
