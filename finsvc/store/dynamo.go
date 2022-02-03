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

func (d *Dynamo) XetBalanceSheet(item map[string]types.AttributeValue) (model.BalanceSheet, error) {
	bs := model.BalanceSheet{}
	bs.AccountsPayable, _ = strconv.ParseFloat(item["accountsPayable"].(*types.AttributeValueMemberN).Value, 64)
	bs.CapitalSurplus, _ = strconv.ParseFloat(item["capitalSurplus"].(*types.AttributeValueMemberN).Value, 64)
	bs.CommonStock, _ = strconv.ParseFloat(item["commonStock"].(*types.AttributeValueMemberN).Value, 64)
	bs.Currency = item["currency"].(*types.AttributeValueMemberS).Value
	bs.CurrentAssets, _ = strconv.ParseFloat(item["currentAssets"].(*types.AttributeValueMemberN).Value, 64)
	bs.CurrentCash, _ = strconv.ParseFloat(item["currentCash"].(*types.AttributeValueMemberN).Value, 64)
	bs.CurrentLongTermDebt, _ = strconv.ParseFloat(item["currentLongTermDebt"].(*types.AttributeValueMemberN).Value, 64)
	bs.FilingType = item["filingType"].(*types.AttributeValueMemberS).Value
	fdte, _ := time.Parse("2006-01-02", item["fiscalDate"].(*types.AttributeValueMemberS).Value)
	bs.FiscalDate = fdte
	fq, _ := strconv.ParseInt(item["fiscalQuarter"].(*types.AttributeValueMemberN).Value, 10, 32)
	bs.FiscalQuarter = int(fq)
	fy, _ := strconv.ParseInt(item["fiscalYear"].(*types.AttributeValueMemberN).Value, 10, 32)
	bs.FiscalYear = int(fy)
	bs.Goodwill, _ = strconv.ParseFloat(item["goodwill"].(*types.AttributeValueMemberN).Value, 64)
	bs.IntangibleAssets, _ = strconv.ParseFloat(item["intangibleAssets"].(*types.AttributeValueMemberN).Value, 64)
	bs.Inventory, _ = strconv.ParseFloat(item["inventory"].(*types.AttributeValueMemberN).Value, 64)
	bs.LongTermDebt, _ = strconv.ParseFloat(item["longTermDebt"].(*types.AttributeValueMemberN).Value, 64)
	bs.LongTermInvestments, _ = strconv.ParseFloat(item["longTermInvestments"].(*types.AttributeValueMemberN).Value, 64)
	bs.MinorityInterest, _ = strconv.ParseFloat(item["minorityInterest"].(*types.AttributeValueMemberN).Value, 64)
	bs.NetTangibleAssets, _ = strconv.ParseFloat(item["netTangibleAssets"].(*types.AttributeValueMemberN).Value, 64)
	bs.OtherAssets, _ = strconv.ParseFloat(item["otherAssets"].(*types.AttributeValueMemberN).Value, 64)
	bs.OtherCurrentAssets, _ = strconv.ParseFloat(item["otherCurrentAssets"].(*types.AttributeValueMemberN).Value, 64)
	bs.OtherCurrentLiabilities, _ = strconv.ParseFloat(item["otherCurrentLiabilities"].(*types.AttributeValueMemberN).Value, 64)
	bs.OtherLiabilities, _ = strconv.ParseFloat(item["otherLiabilities"].(*types.AttributeValueMemberN).Value, 64)
	bs.PropertyPlantEquipment, _ = strconv.ParseFloat(item["propertyPlantEquipment"].(*types.AttributeValueMemberN).Value, 64)
	bs.Receivables, _ = strconv.ParseFloat(item["receivables"].(*types.AttributeValueMemberN).Value, 64)
	rdte, _ := time.Parse("2006-01-02", item["reportDate"].(*types.AttributeValueMemberS).Value)
	bs.ReportDate = rdte
	bs.RetainedEarnings, _ = strconv.ParseFloat(item["retainedEarnings"].(*types.AttributeValueMemberN).Value, 64)
	bs.ShareholderEquity, _ = strconv.ParseFloat(item["shareholderEquity"].(*types.AttributeValueMemberN).Value, 64)
	bs.ShortTermInvestments, _ = strconv.ParseFloat(item["shortTermInvestments"].(*types.AttributeValueMemberN).Value, 64)
	bs.TotalAssets, _ = strconv.ParseFloat(item["totalAssets"].(*types.AttributeValueMemberN).Value, 64)
	bs.TotalCurrentLiabilities, _ = strconv.ParseFloat(item["totalCurrentLiabilities"].(*types.AttributeValueMemberN).Value, 64)
	bs.TotalLiabilities, _ = strconv.ParseFloat(item["totalLiabilities"].(*types.AttributeValueMemberN).Value, 64)

	return bs, nil

}

func (d *Dynamo) GetAnnualBalanceSheet(ctx context.Context, symbol string, limit int) ([]model.BalanceSheet, error) {
	_limit := int32(limit)
	pkFilter := fmt.Sprintf("SYMBOL#%s", symbol)
	skFilter := "STATEMENT#BALANCESHEET#FILING#10-K"

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
		bs.FiscalDate, _ = time.Parse("2006-01-02", out.Items[i]["fiscalDate"].(*types.AttributeValueMemberS).Value)
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

func (d *Dynamo) InsertAnnualBalanceSheet(ctx context.Context, symbol string, bs *model.BalanceSheet) error {
	pk := fmt.Sprintf("SYMBOL#%s", symbol)
	sk := fmt.Sprintf("STATEMENT#BALANCESHEET#FILING#%s#FISCALDATE#%s", bs.FilingType, bs.FiscalDate.String())
	d.log.Printf("Inserting bs into dynamo. pk: %s\tsk: %s\n", pk, sk)

	input := &dynamodb.PutItemInput{
		Item: map[string]types.AttributeValue{
			"pk":                      &types.AttributeValueMemberS{Value: pk},
			"sk":                      &types.AttributeValueMemberS{Value: sk},
			"reportDate":              &types.AttributeValueMemberS{Value: bs.ReportDate.String()},
			"filingType":              &types.AttributeValueMemberS{Value: bs.FilingType},
			"fiscalDate":              &types.AttributeValueMemberS{Value: bs.FiscalDate.String()},
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
		},
		TableName: aws.String("Financials"),
	}

	_, err := d.client.PutItem(ctx, input)
	if err != nil {
		d.log.Fatalf("Unable to insert Annual Balance Sheet for symbol: %s, %v", symbol, err)
	}

	return err
}
