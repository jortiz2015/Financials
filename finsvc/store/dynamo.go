package store

import (
	"context"
	"fin/api"
	"fin/model"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"

	//"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	//"encoding/json"
)

type Dynamo struct {
	log    log.Logger
	client *dynamodb.Client
}

func NewDynamo(log *log.Logger) (api.Store, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Using the Config value, create the DynamoDB client
	client := dynamodb.NewFromConfig(cfg)
	return &Dynamo{log: *log, client: client}, err
}

func (d *Dynamo) GetAnnualBalanceSheets(ctx context.Context, symbol string, limit int) ([]model.BalanceSheet, error) {
	return d.getBalanceSheets(ctx, symbol, "10-K", limit)
}

func (d *Dynamo) GetQuarterlyBalanceSheets(ctx context.Context, symbol string, limit int) ([]model.BalanceSheet, error) {
	return d.getBalanceSheets(ctx, symbol, "10-Q", limit)
}

func (d *Dynamo) getBalanceSheets(ctx context.Context, symbol string, filing string, limit int) ([]model.BalanceSheet, error) {
	_limit := int32(limit)
	pkFilter := fmt.Sprintf("SYMBOL#%s", symbol)
	skFilter := fmt.Sprintf("STATEMENT#BALANCESHEET#FILING#%s", filing)

	keyCond := expression.KeyAnd(
		expression.Key("pk").Equal(expression.Value(pkFilter)),
		expression.KeyBeginsWith(expression.Key("sk"), skFilter),
	)

	expr, err := expression.NewBuilder().WithKeyCondition(keyCond).Build()
	if err != nil {
		d.log.Fatalf("unable to build expression, %v", err)
	}
	out, err := d.client.Query(ctx, &dynamodb.QueryInput{
		TableName:                 aws.String("Financials"),
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		ScanIndexForward:          aws.Bool(false), // Descending
		Limit:                     &_limit,
	})

	if err != nil {
		d.log.Fatalf("Unable to fetch Annual Balance Sheet for symbol: %s, %v", symbol, err)
		return []model.BalanceSheet{}, err
	}

	d.log.Println("QUERYING BALANCE SHEET ", out.Count)

	balanceSheets := []model.BalanceSheet{}
	//attributevalue.UnmarshalListOfMaps(out.Items, &balanceSheets)

	for i := 0; i < int(out.Count); i++ {
		bs := model.BalanceSheet{}
		bs.ReportDate, _ = time.Parse("2006-01-02", out.Items[i]["reportDate"].(*types.AttributeValueMemberS).Value)
		bs.FilingType = out.Items[i]["filingType"].(*types.AttributeValueMemberS).Value
		bs.FiscalDate, _ = time.Parse("2006-01-02", (out.Items[i]["fiscalDate"].(*types.AttributeValueMemberS).Value))
		fq, _ := strconv.ParseInt(out.Items[i]["fiscalQuarter"].(*types.AttributeValueMemberN).Value, 10, 32)
		bs.FiscalQuarter = int(fq)
		fy, _ := strconv.ParseInt(out.Items[i]["fiscalYear"].(*types.AttributeValueMemberN).Value, 10, 32)
		bs.FiscalYear = int(fy)
		bs.Currency = out.Items[i]["currency"].(*types.AttributeValueMemberS).Value
		bs.CurrentCash, _ = strconv.ParseFloat(out.Items[i]["currentCash"].(*types.AttributeValueMemberN).Value, 64)
		bs.ShortTermInvestments, _ = strconv.ParseFloat(out.Items[i]["shortTermInvestments"].(*types.AttributeValueMemberN).Value, 64)
		bs.Receivables, _ = strconv.ParseFloat(out.Items[i]["receivables"].(*types.AttributeValueMemberN).Value, 64)
		bs.Inventory, _ = strconv.ParseFloat(out.Items[i]["inventory"].(*types.AttributeValueMemberN).Value, 64)
		bs.OtherCurrentAssets, _ = strconv.ParseFloat(out.Items[i]["otherCurrentAssets"].(*types.AttributeValueMemberN).Value, 64)
		bs.CurrentAssets, _ = strconv.ParseFloat(out.Items[i]["currentAssets"].(*types.AttributeValueMemberN).Value, 64)
		bs.LongTermInvestments, _ = strconv.ParseFloat(out.Items[i]["longTermInvestments"].(*types.AttributeValueMemberN).Value, 64)
		bs.PropertyPlantEquipment, _ = strconv.ParseFloat(out.Items[i]["propertyPlantEquipment"].(*types.AttributeValueMemberN).Value, 64)
		bs.Goodwill, _ = strconv.ParseFloat(out.Items[i]["goodwill"].(*types.AttributeValueMemberN).Value, 64)
		bs.IntangibleAssets, _ = strconv.ParseFloat(out.Items[i]["intangibleAssets"].(*types.AttributeValueMemberN).Value, 64)
		bs.OtherAssets, _ = strconv.ParseFloat(out.Items[i]["otherAssets"].(*types.AttributeValueMemberN).Value, 64)
		bs.TotalAssets, _ = strconv.ParseFloat(out.Items[i]["totalAssets"].(*types.AttributeValueMemberN).Value, 64)
		bs.AccountsPayable, _ = strconv.ParseFloat(out.Items[i]["accountsPayable"].(*types.AttributeValueMemberN).Value, 64)
		bs.CurrentLongTermDebt, _ = strconv.ParseFloat(out.Items[i]["currentLongTermDebt"].(*types.AttributeValueMemberN).Value, 64)
		bs.OtherCurrentLiabilities, _ = strconv.ParseFloat(out.Items[i]["otherCurrentLiabilities"].(*types.AttributeValueMemberN).Value, 64)
		bs.TotalCurrentLiabilities, _ = strconv.ParseFloat(out.Items[i]["totalCurrentLiabilities"].(*types.AttributeValueMemberN).Value, 64)
		bs.LongTermDebt, _ = strconv.ParseFloat(out.Items[i]["longTermDebt"].(*types.AttributeValueMemberN).Value, 64)
		bs.OtherLiabilities, _ = strconv.ParseFloat(out.Items[i]["otherLiabilities"].(*types.AttributeValueMemberN).Value, 64)
		bs.MinorityInterest, _ = strconv.ParseFloat(out.Items[i]["minorityInterest"].(*types.AttributeValueMemberN).Value, 64)
		bs.TotalLiabilities, _ = strconv.ParseFloat(out.Items[i]["totalLiabilities"].(*types.AttributeValueMemberN).Value, 64)
		bs.CommonStock, _ = strconv.ParseFloat(out.Items[i]["commonStock"].(*types.AttributeValueMemberN).Value, 64)
		bs.RetainedEarnings, _ = strconv.ParseFloat(out.Items[i]["retainedEarnings"].(*types.AttributeValueMemberN).Value, 64)
		bs.TreasuryStock, _ = strconv.ParseFloat(out.Items[i]["treasuryStock"].(*types.AttributeValueMemberN).Value, 64)
		bs.CapitalSurplus, _ = strconv.ParseFloat(out.Items[i]["capitalSurplus"].(*types.AttributeValueMemberN).Value, 64)
		bs.ShareholderEquity, _ = strconv.ParseFloat(out.Items[i]["shareholderEquity"].(*types.AttributeValueMemberN).Value, 64)
		bs.NetTangibleAssets, _ = strconv.ParseFloat(out.Items[i]["netTangibleAssets"].(*types.AttributeValueMemberN).Value, 64)

		balanceSheets = append(balanceSheets, bs)
	}

	return balanceSheets, nil
}

func (d *Dynamo) InsertBalanceSheet(ctx context.Context, symbol string, bs *model.BalanceSheet) error {
	item := d.MarshalListOfMapsBS(symbol, bs)
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String("Financials"),
	}

	_, err := d.client.PutItem(ctx, input)
	if err != nil {
		d.log.Fatalf("Unable to insert Balance Sheet for symbol: %s, %v", symbol, err)
	}

	return err
}

func (d *Dynamo) InsertBalanceSheets(ctx context.Context, symbol string, bs []model.BalanceSheet) error {
	input := &dynamodb.BatchWriteItemInput{
		RequestItems: d.GetWriteRequestItemsBS(symbol, bs),
	}

	_, err := d.client.BatchWriteItem(ctx, input)
	if err != nil {
		d.log.Fatalf("Unable to insert bulk Balance Sheet for symbol: %s, %v", symbol, err)
	}

	return nil
}

// Function to create a map[string][]types.WriteRequest of Put Requests to insert array of model.BalanceSheet
func (d *Dynamo) GetWriteRequestItemsBS(symbol string, bs []model.BalanceSheet) map[string][]types.WriteRequest {
	requestItems := map[string][]types.WriteRequest{}
	requestItems["Financials"] = []types.WriteRequest{}

	for i := 0; i < len(bs); i++ {
		requestItems["Financials"] = append(requestItems["Financials"], types.WriteRequest{
			PutRequest: &types.PutRequest{
				Item: d.MarshalListOfMapsBS(symbol, &bs[i]),
			},
		})
	}

	return requestItems
}

func (d *Dynamo) MarshalListOfMapsBS(symbol string, bs *model.BalanceSheet) map[string]types.AttributeValue {
	pk := fmt.Sprintf("SYMBOL#%s", symbol)
	sk := fmt.Sprintf("STATEMENT#BALANCESHEET#FILING#%s#FISCALDATE#%s", bs.FilingType, bs.FiscalDate.Format("2006-01-02"))
	d.log.Printf("Inserting bs into dynamo. pk: %s\tsk: %s\n", pk, sk)

	item := map[string]types.AttributeValue{
		"pk":                      &types.AttributeValueMemberS{Value: pk},
		"sk":                      &types.AttributeValueMemberS{Value: sk},
		"reportDate":              &types.AttributeValueMemberS{Value: bs.ReportDate.Format("2006-01-02")},
		"filingType":              &types.AttributeValueMemberS{Value: bs.FilingType},
		"fiscalDate":              &types.AttributeValueMemberS{Value: bs.FiscalDate.Format("2006-01-02")},
		"fiscalQuarter":           &types.AttributeValueMemberN{Value: strconv.FormatInt(int64(bs.FiscalQuarter), 10)},
		"fiscalYear":              &types.AttributeValueMemberN{Value: strconv.FormatInt(int64(bs.FiscalYear), 10)},
		"currency":                &types.AttributeValueMemberS{Value: bs.Currency},
		"currentCash":             &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.CurrentCash, 'f', -1, 64)},
		"shortTermInvestments":    &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.ShortTermInvestments, 'f', -1, 64)},
		"receivables":             &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.Receivables, 'f', -1, 64)},
		"inventory":               &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.Inventory, 'f', -1, 64)},
		"otherCurrentAssets":      &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.OtherCurrentAssets, 'f', -1, 64)},
		"currentAssets":           &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.CurrentAssets, 'f', -1, 64)},
		"longTermInvestments":     &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.LongTermInvestments, 'f', -1, 64)},
		"propertyPlantEquipment":  &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.PropertyPlantEquipment, 'f', -1, 64)},
		"goodwill":                &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.Goodwill, 'f', -1, 64)},
		"intangibleAssets":        &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.IntangibleAssets, 'f', -1, 64)},
		"otherAssets":             &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.OtherAssets, 'f', -1, 64)},
		"totalAssets":             &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.TotalAssets, 'f', -1, 64)},
		"accountsPayable":         &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.AccountsPayable, 'f', -1, 64)},
		"currentLongTermDebt":     &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.CurrentLongTermDebt, 'f', -1, 64)},
		"otherCurrentLiabilities": &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.OtherCurrentLiabilities, 'f', -1, 64)},
		"totalCurrentLiabilities": &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.TotalCurrentLiabilities, 'f', -1, 64)},
		"longTermDebt":            &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.LongTermDebt, 'f', -1, 64)},
		"otherLiabilities":        &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.OtherLiabilities, 'f', -1, 64)},
		"minorityInterest":        &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.MinorityInterest, 'f', -1, 64)},
		"totalLiabilities":        &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.TotalLiabilities, 'f', -1, 64)},
		"commonStock":             &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.CommonStock, 'f', -1, 64)},
		"retainedEarnings":        &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.RetainedEarnings, 'f', -1, 64)},
		"treasuryStock":           &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.TreasuryStock, 'f', -1, 64)},
		"capitalSurplus":          &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.CapitalSurplus, 'f', -1, 64)},
		"shareholderEquity":       &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.ShareholderEquity, 'f', -1, 64)},
		"netTangibleAssets":       &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.NetTangibleAssets, 'f', -1, 64)},
	}

	return item
}

func (d *Dynamo) GetAnnualIncomeStatements(ctx context.Context, symbol string, limit int) ([]model.IncomeStatement, error) {
	return d.getIncomeStatements(ctx, symbol, "10-K", limit)
}

func (d *Dynamo) GetQuarterlyIncomeStatements(ctx context.Context, symbol string, limit int) ([]model.IncomeStatement, error) {
	return d.getIncomeStatements(ctx, symbol, "10-Q", limit)
}

func (d *Dynamo) getIncomeStatements(ctx context.Context, symbol string, filing string, limit int) ([]model.IncomeStatement, error) {
	_limit := int32(limit)
	pkFilter := fmt.Sprintf("SYMBOL#%s", symbol)
	skFilter := fmt.Sprintf("STATEMENT#INCOME#FILING#%s", filing)

	keyCond := expression.KeyAnd(
		expression.Key("pk").Equal(expression.Value(pkFilter)),
		expression.KeyBeginsWith(expression.Key("sk"), skFilter),
	)

	expr, err := expression.NewBuilder().WithKeyCondition(keyCond).Build()
	if err != nil {
		d.log.Fatalf("unable to build expression, %v", err)
	}
	out, err := d.client.Query(ctx, &dynamodb.QueryInput{
		TableName:                 aws.String("Financials"),
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		ScanIndexForward:          aws.Bool(false), // Descending
		Limit:                     &_limit,
	})

	if err != nil {
		d.log.Fatalf("Unable to fetch Income Statements for symbol: %s, %v", symbol, err)
		return []model.IncomeStatement{}, err
	}

	d.log.Println("QUERYING INCOME STATEMENTS", out.Count)

	incomeStatements := []model.IncomeStatement{}
	//attributevalue.UnmarshalListOfMaps(out.Items, &balanceSheets)

	for i := 0; i < int(out.Count); i++ {
		is := model.IncomeStatement{}
		is.ReportDate, _ = time.Parse("2006-01-02", out.Items[i]["reportDate"].(*types.AttributeValueMemberS).Value)
		is.FiscalDate, _ = time.Parse("2006-01-02", (out.Items[i]["fiscalDate"].(*types.AttributeValueMemberS).Value))
		is.Currency = out.Items[i]["currency"].(*types.AttributeValueMemberS).Value
		is.TotalRevenue, _ = strconv.ParseFloat(out.Items[i]["totalRevenue"].(*types.AttributeValueMemberN).Value, 64)
		is.CostOfRevenue, _ = strconv.ParseFloat(out.Items[i]["costOfRevenue"].(*types.AttributeValueMemberN).Value, 64)
		is.GrossProfit, _ = strconv.ParseFloat(out.Items[i]["grossProfit"].(*types.AttributeValueMemberN).Value, 64)
		is.ResearchAndDevelopment, _ = strconv.ParseFloat(out.Items[i]["researchAndDevelopment"].(*types.AttributeValueMemberN).Value, 64)
		is.SellingGeneralAndAdmin, _ = strconv.ParseFloat(out.Items[i]["sellingGeneralAndAdmin"].(*types.AttributeValueMemberN).Value, 64)
		is.OperatingExpense, _ = strconv.ParseFloat(out.Items[i]["operatingExpense"].(*types.AttributeValueMemberN).Value, 64)
		is.OperatingIncome, _ = strconv.ParseFloat(out.Items[i]["operatingIncome"].(*types.AttributeValueMemberN).Value, 64)
		is.OtherIncomeExpenseNet, _ = strconv.ParseFloat(out.Items[i]["otherIncomeExpenseNet"].(*types.AttributeValueMemberN).Value, 64)
		is.EBIT, _ = strconv.ParseFloat(out.Items[i]["ebit"].(*types.AttributeValueMemberN).Value, 64)
		is.InterestIncome, _ = strconv.ParseFloat(out.Items[i]["interestIncome"].(*types.AttributeValueMemberN).Value, 64)
		is.PretaxIncome, _ = strconv.ParseFloat(out.Items[i]["pretaxIncome"].(*types.AttributeValueMemberN).Value, 64)
		is.IncomeTax, _ = strconv.ParseFloat(out.Items[i]["incomeTax"].(*types.AttributeValueMemberN).Value, 64)
		is.MinorityInterest, _ = strconv.ParseFloat(out.Items[i]["minorityInterest"].(*types.AttributeValueMemberN).Value, 64)
		is.NetIncome, _ = strconv.ParseFloat(out.Items[i]["netIncome"].(*types.AttributeValueMemberN).Value, 64)
		is.NetIncomeBasic, _ = strconv.ParseFloat(out.Items[i]["netIncomeBasic"].(*types.AttributeValueMemberN).Value, 64)

		incomeStatements = append(incomeStatements, is)
	}

	return incomeStatements, nil
}

func (d *Dynamo) InsertIncomeStatement(ctx context.Context, symbol string, is *model.IncomeStatement, filing string) error {
	item := d.MarshalListOfMapsIS(symbol, is, filing)
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String("Financials"),
	}

	_, err := d.client.PutItem(ctx, input)
	if err != nil {
		d.log.Fatalf("Unable to insert Income Statement for symbol: %s, %v", symbol, err)
	}

	return err
}

func (d *Dynamo) InsertIncomeStatements(ctx context.Context, symbol string, is []model.IncomeStatement, filing string) error {
	input := &dynamodb.BatchWriteItemInput{
		RequestItems: d.GetWriteRequestItemsIS(symbol, is, filing),
	}

	_, err := d.client.BatchWriteItem(ctx, input)
	if err != nil {
		d.log.Fatalf("Unable to insert bulk Income Statements for symbol: %s, %v", symbol, err)
	}

	return nil
}

// Function to create a map[string][]types.WriteRequest of Put Requests to insert array of model.BalanceSheet
func (d *Dynamo) GetWriteRequestItemsIS(symbol string, is []model.IncomeStatement, filing string) map[string][]types.WriteRequest {
	requestItems := map[string][]types.WriteRequest{}
	requestItems["Financials"] = []types.WriteRequest{}

	for i := 0; i < len(is); i++ {
		requestItems["Financials"] = append(requestItems["Financials"], types.WriteRequest{
			PutRequest: &types.PutRequest{
				Item: d.MarshalListOfMapsIS(symbol, &is[i], filing),
			},
		})
	}

	return requestItems
}

func (d *Dynamo) MarshalListOfMapsIS(symbol string, is *model.IncomeStatement, filing string) map[string]types.AttributeValue {
	pk := fmt.Sprintf("SYMBOL#%s", symbol)
	sk := fmt.Sprintf("STATEMENT#INCOME#FILING#%s#FISCALDATE#%s", filing, is.FiscalDate.Format("2006-01-02"))
	d.log.Printf("Inserting is into dynamo. pk: %s\tsk: %s\n", pk, sk)

	item := map[string]types.AttributeValue{
		"pk":                     &types.AttributeValueMemberS{Value: pk},
		"sk":                     &types.AttributeValueMemberS{Value: sk},
		"reportDate":             &types.AttributeValueMemberS{Value: is.ReportDate.Format("2006-01-02")},
		"fiscalDate":             &types.AttributeValueMemberS{Value: is.FiscalDate.Format("2006-01-02")},
		"currency":               &types.AttributeValueMemberS{Value: is.Currency},
		"totalRevenue":           &types.AttributeValueMemberN{Value: strconv.FormatFloat(is.TotalRevenue, 'f', -1, 64)},
		"costOfRevenue":          &types.AttributeValueMemberN{Value: strconv.FormatFloat(is.CostOfRevenue, 'f', -1, 64)},
		"grossProfit":            &types.AttributeValueMemberN{Value: strconv.FormatFloat(is.GrossProfit, 'f', -1, 64)},
		"researchAndDevelopment": &types.AttributeValueMemberN{Value: strconv.FormatFloat(is.ResearchAndDevelopment, 'f', -1, 64)},
		"sellingGeneralAndAdmin": &types.AttributeValueMemberN{Value: strconv.FormatFloat(is.SellingGeneralAndAdmin, 'f', -1, 64)},
		"operatingExpense":       &types.AttributeValueMemberN{Value: strconv.FormatFloat(is.OperatingExpense, 'f', -1, 64)},
		"operatingIncome":        &types.AttributeValueMemberN{Value: strconv.FormatFloat(is.OperatingIncome, 'f', -1, 64)},
		"otherIncomeExpenseNet":  &types.AttributeValueMemberN{Value: strconv.FormatFloat(is.OtherIncomeExpenseNet, 'f', -1, 64)},
		"ebit":                   &types.AttributeValueMemberN{Value: strconv.FormatFloat(is.EBIT, 'f', -1, 64)},
		"interestIncome":         &types.AttributeValueMemberN{Value: strconv.FormatFloat(is.InterestIncome, 'f', -1, 64)},
		"pretaxIncome":           &types.AttributeValueMemberN{Value: strconv.FormatFloat(is.PretaxIncome, 'f', -1, 64)},
		"incomeTax":              &types.AttributeValueMemberN{Value: strconv.FormatFloat(is.IncomeTax, 'f', -1, 64)},
		"minorityInterest":       &types.AttributeValueMemberN{Value: strconv.FormatFloat(is.MinorityInterest, 'f', -1, 64)},
		"netIncome":              &types.AttributeValueMemberN{Value: strconv.FormatFloat(is.NetIncome, 'f', -1, 64)},
		"netIncomeBasic":         &types.AttributeValueMemberN{Value: strconv.FormatFloat(is.NetIncomeBasic, 'f', -1, 64)},
	}

	return item
}

// CASH FLOW FUNCTIONS

func (d *Dynamo) GetAnnualCashFlows(ctx context.Context, symbol string, limit int) ([]model.CashFlow, error) {
	return d.getCashFlows(ctx, symbol, "10-K", limit)
}

func (d *Dynamo) GetQuarterlyCashFlows(ctx context.Context, symbol string, limit int) ([]model.CashFlow, error) {
	return d.getCashFlows(ctx, symbol, "10-Q", limit)
}

func (d *Dynamo) getCashFlows(ctx context.Context, symbol string, filing string, limit int) ([]model.CashFlow, error) {
	_limit := int32(limit)
	pkFilter := fmt.Sprintf("SYMBOL#%s", symbol)
	skFilter := fmt.Sprintf("STATEMENT#CASHFLOW#FILING#%s", filing)

	keyCond := expression.KeyAnd(
		expression.Key("pk").Equal(expression.Value(pkFilter)),
		expression.KeyBeginsWith(expression.Key("sk"), skFilter),
	)

	expr, err := expression.NewBuilder().WithKeyCondition(keyCond).Build()
	if err != nil {
		d.log.Fatalf("unable to build expression, %v", err)
	}
	out, err := d.client.Query(ctx, &dynamodb.QueryInput{
		TableName:                 aws.String("Financials"),
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		ScanIndexForward:          aws.Bool(false), // Descending
		Limit:                     &_limit,
	})

	if err != nil {
		d.log.Fatalf("Unable to fetch Annual Cash Flows for symbol: %s, %v", symbol, err)
		return []model.CashFlow{}, err
	}

	d.log.Println("QUERYING CASH FLOWS", out.Count)

	cashFlows := []model.CashFlow{}
	//attributevalue.UnmarshalListOfMaps(out.Items, &balanceSheets)

	for i := 0; i < int(out.Count); i++ {
		cf := model.CashFlow{}
		cf.ReportDate, _ = time.Parse("2006-01-02", out.Items[i]["reportDate"].(*types.AttributeValueMemberS).Value)
		cf.FiscalDate, _ = time.Parse("2006-01-02", (out.Items[i]["fiscalDate"].(*types.AttributeValueMemberS).Value))
		cf.Currency = out.Items[i]["currency"].(*types.AttributeValueMemberS).Value
		cf.NetIncome, _ = strconv.ParseFloat(out.Items[i]["netIncome"].(*types.AttributeValueMemberN).Value, 64)
		cf.Depreciation, _ = strconv.ParseFloat(out.Items[i]["depreciation"].(*types.AttributeValueMemberN).Value, 64)
		cf.ChangesInReceivables, _ = strconv.ParseFloat(out.Items[i]["changesInReceivables"].(*types.AttributeValueMemberN).Value, 64)
		cf.ChangesInInventories, _ = strconv.ParseFloat(out.Items[i]["changesInInventories"].(*types.AttributeValueMemberN).Value, 64)
		cf.CashChange, _ = strconv.ParseFloat(out.Items[i]["cashChange"].(*types.AttributeValueMemberN).Value, 64)
		cf.CashFlow, _ = strconv.ParseFloat(out.Items[i]["cashFlow"].(*types.AttributeValueMemberN).Value, 64)
		cf.CapitalExpenditures, _ = strconv.ParseFloat(out.Items[i]["capitalExpenditures"].(*types.AttributeValueMemberN).Value, 64)
		cf.Investments, _ = strconv.ParseFloat(out.Items[i]["investments"].(*types.AttributeValueMemberN).Value, 64)
		cf.InvestingActivityOther, _ = strconv.ParseFloat(out.Items[i]["investingActivityOther"].(*types.AttributeValueMemberN).Value, 64)
		cf.TotalInvestingCashFlows, _ = strconv.ParseFloat(out.Items[i]["totalInvestingCashFlows"].(*types.AttributeValueMemberN).Value, 64)
		cf.DividendsPaid, _ = strconv.ParseFloat(out.Items[i]["dividendsPaid"].(*types.AttributeValueMemberN).Value, 64)
		cf.NetBorrowings, _ = strconv.ParseFloat(out.Items[i]["netBorrowings"].(*types.AttributeValueMemberN).Value, 64)
		cf.OtherFinancingCashFlows, _ = strconv.ParseFloat(out.Items[i]["otherFinancingCashFlows"].(*types.AttributeValueMemberN).Value, 64)
		cf.CashFlowFinancing, _ = strconv.ParseFloat(out.Items[i]["cashFlowFinancing"].(*types.AttributeValueMemberN).Value, 64)
		cf.ExchangeRateEffect, _ = strconv.ParseFloat(out.Items[i]["exchangeRateEffect"].(*types.AttributeValueMemberN).Value, 64)

		cashFlows = append(cashFlows, cf)
	}

	return cashFlows, nil
}

func (d *Dynamo) InsertCashFlow(ctx context.Context, symbol string, cf *model.CashFlow, filing string) error {
	item := d.MarshalListOfMapsCF(symbol, cf, filing)
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String("Financials"),
	}

	_, err := d.client.PutItem(ctx, input)
	if err != nil {
		d.log.Fatalf("Unable to insert Annual Cash Flow for symbol: %s, %v", symbol, err)
	}

	return err
}

func (d *Dynamo) InsertCashFlows(ctx context.Context, symbol string, cf []model.CashFlow, filing string) error {
	input := &dynamodb.BatchWriteItemInput{
		RequestItems: d.GetWriteRequestItemsCF(symbol, cf, filing),
	}

	_, err := d.client.BatchWriteItem(ctx, input)
	if err != nil {
		d.log.Fatalf("Unable to insert bulk Annual Cash Flow for symbol: %s, %v", symbol, err)
	}

	return nil
}

// Function to create a map[string][]types.WriteRequest of Put Requests to insert array of model.CashFlow
func (d *Dynamo) GetWriteRequestItemsCF(symbol string, cf []model.CashFlow, filing string) map[string][]types.WriteRequest {
	requestItems := map[string][]types.WriteRequest{}
	requestItems["Financials"] = []types.WriteRequest{}

	for i := 0; i < len(cf); i++ {
		requestItems["Financials"] = append(requestItems["Financials"], types.WriteRequest{
			PutRequest: &types.PutRequest{
				Item: d.MarshalListOfMapsCF(symbol, &cf[i], filing),
			},
		})
	}

	return requestItems
}

func (d *Dynamo) MarshalListOfMapsCF(symbol string, cf *model.CashFlow, filing string) map[string]types.AttributeValue {
	pk := fmt.Sprintf("SYMBOL#%s", symbol)
	sk := fmt.Sprintf("STATEMENT#CASHFLOW#FILING#%s#FISCALDATE#%s", filing, cf.FiscalDate.Format("2006-01-02"))
	d.log.Printf("Inserting cf into dynamo. pk: %s\tsk: %s\n", pk, sk)

	item := map[string]types.AttributeValue{
		"pk":                      &types.AttributeValueMemberS{Value: pk},
		"sk":                      &types.AttributeValueMemberS{Value: sk},
		"reportDate":              &types.AttributeValueMemberS{Value: cf.ReportDate.Format("2006-01-02")},
		"fiscalDate":              &types.AttributeValueMemberS{Value: cf.FiscalDate.Format("2006-01-02")},
		"currency":                &types.AttributeValueMemberS{Value: cf.Currency},
		"netIncome":               &types.AttributeValueMemberN{Value: strconv.FormatFloat(cf.NetIncome, 'f', -1, 64)},
		"depreciation":            &types.AttributeValueMemberN{Value: strconv.FormatFloat(cf.Depreciation, 'f', -1, 64)},
		"changesInReceivables":    &types.AttributeValueMemberN{Value: strconv.FormatFloat(cf.ChangesInReceivables, 'f', -1, 64)},
		"changesInInventories":    &types.AttributeValueMemberN{Value: strconv.FormatFloat(cf.ChangesInInventories, 'f', -1, 64)},
		"cashChange":              &types.AttributeValueMemberN{Value: strconv.FormatFloat(cf.CashChange, 'f', -1, 64)},
		"cashFlow":                &types.AttributeValueMemberN{Value: strconv.FormatFloat(cf.CashFlow, 'f', -1, 64)},
		"capitalExpenditures":     &types.AttributeValueMemberN{Value: strconv.FormatFloat(cf.CapitalExpenditures, 'f', -1, 64)},
		"investments":             &types.AttributeValueMemberN{Value: strconv.FormatFloat(cf.Investments, 'f', -1, 64)},
		"investingActivityOther":  &types.AttributeValueMemberN{Value: strconv.FormatFloat(cf.InvestingActivityOther, 'f', -1, 64)},
		"totalInvestingCashFlows": &types.AttributeValueMemberN{Value: strconv.FormatFloat(cf.TotalInvestingCashFlows, 'f', -1, 64)},
		"dividendsPaid":           &types.AttributeValueMemberN{Value: strconv.FormatFloat(cf.DividendsPaid, 'f', -1, 64)},
		"netBorrowings":           &types.AttributeValueMemberN{Value: strconv.FormatFloat(cf.NetBorrowings, 'f', -1, 64)},
		"otherFinancingCashFlows": &types.AttributeValueMemberN{Value: strconv.FormatFloat(cf.OtherFinancingCashFlows, 'f', -1, 64)},
		"cashFlowFinancing":       &types.AttributeValueMemberN{Value: strconv.FormatFloat(cf.CashFlowFinancing, 'f', -1, 64)},
		"exchangeRateEffect":      &types.AttributeValueMemberN{Value: strconv.FormatFloat(cf.ExchangeRateEffect, 'f', -1, 64)},
	}

	return item
}
