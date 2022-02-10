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
