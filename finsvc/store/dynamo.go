package store

import (
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"google.golang.org/protobuf/types/known/timestamppb"

	//"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	api "finsvc/api"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	//"encoding/json"
	"context"
	pb "finsvc/pb"
	"log"
)

type Dynamo struct {
	log log.Logger
	iex api.IEX
}

func NewDynamo(log *log.Logger) *Dynamo {
	iex := api.NewIEX(log)
	return &Dynamo{log: *log, iex: *iex}
}

func (d *Dynamo) setBalanceSheet(item map[string]types.AttributeValue) (*pb.BalanceSheet, error) {
	bs := &pb.BalanceSheet{}
	bs.AccountsPayable, _ = strconv.ParseFloat(item["accountsPayable"].(*types.AttributeValueMemberN).Value, 64)
	bs.CapitalSurplus, _ = strconv.ParseFloat(item["capitalSurplus"].(*types.AttributeValueMemberN).Value, 64)
	bs.CommonStock, _ = strconv.ParseFloat(item["commonStock"].(*types.AttributeValueMemberN).Value, 64)
	bs.Currency = item["currency"].(*types.AttributeValueMemberS).Value
	bs.CurrentAssets, _ = strconv.ParseFloat(item["currentAssets"].(*types.AttributeValueMemberN).Value, 64)
	bs.CurrentCash, _ = strconv.ParseFloat(item["currentCash"].(*types.AttributeValueMemberN).Value, 64)
	bs.CurrentLongTermDebt, _ = strconv.ParseFloat(item["currentLongTermDebt"].(*types.AttributeValueMemberN).Value, 64)
	bs.FilingType = item["filingType"].(*types.AttributeValueMemberS).Value
	fdte, _ := time.Parse("2006-01-02", item["fiscalDate"].(*types.AttributeValueMemberS).Value)
	bs.FiscalDate = timestamppb.New(fdte)
	fq, _ := strconv.ParseInt(item["fiscalQuarter"].(*types.AttributeValueMemberN).Value, 10, 32)
	bs.FiscalQuarter = int32(fq)
	fy, _ := strconv.ParseInt(item["fiscalYear"].(*types.AttributeValueMemberN).Value, 10, 32)
	bs.FiscalYear = int32(fy)
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
	bs.ReportDate = timestamppb.New(rdte)
	bs.RetainedEarnings, _ = strconv.ParseFloat(item["retainedEarnings"].(*types.AttributeValueMemberN).Value, 64)
	bs.ShareholderEquity, _ = strconv.ParseFloat(item["shareholderEquity"].(*types.AttributeValueMemberN).Value, 64)
	bs.ShortTermInvestments, _ = strconv.ParseFloat(item["shortTermInvestments"].(*types.AttributeValueMemberN).Value, 64)
	bs.TotalAssets, _ = strconv.ParseFloat(item["totalAssets"].(*types.AttributeValueMemberN).Value, 64)
	bs.TotalCurrentLiabilities, _ = strconv.ParseFloat(item["totalCurrentLiabilities"].(*types.AttributeValueMemberN).Value, 64)
	bs.TotalLiabilities, _ = strconv.ParseFloat(item["totalLiabilities"].(*types.AttributeValueMemberN).Value, 64)

	return bs, nil

}

func (d *Dynamo) GetBalanceSheet(ctx context.Context, ticker *pb.Ticker) (*pb.BalanceSheet, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		d.log.Fatalf("unable to load SDK config, %v", err)
	}

	// Using the Config value, create the DynamoDB client
	client := dynamodb.NewFromConfig(cfg)

	/*
		getItemInput := &dynamodb.GetItemInput{
			Key: map[string]types.AttributeValue{
				"symbol":     &types.AttributeValueMemberS{Value: ticker.Symbol},
				"fiscalDate": &types.AttributeValueMemberS{Value: "2021-09-11"},
			},
			TableName: aws.String("BalanceSheet"),
		}

		resp, err := client.GetItem(context.Background(), getItemInput)

		if err != nil {
			err = d.InsertBalanceSheet(ctx, ticker)
			return nil, err
		}

		bs, _ := d.setBalanceSheet(resp.Item)
	*/
	t := time.Now()
	today := t.Format("2006-01-02")
	var limit int32 = 1

	keyCond := expression.KeyAnd(
		expression.Key("symbol").Equal(expression.Value(ticker.Symbol)),
		expression.Key("fiscalDate").LessThanEqual(expression.Value(today)),
	)

	expr, err := expression.NewBuilder().WithKeyCondition(keyCond).Build()
	if err != nil {
		d.log.Fatalf("unable to build expression, %v", err)
	}
	out, err := client.Query(ctx, &dynamodb.QueryInput{
		TableName:                 aws.String("BalanceSheet"),
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		ScanIndexForward:          aws.Bool(false), // Descending
		Limit:                     &limit,          // Limit to 1
	})

	if err != nil {
		d.log.Fatalf("unable to query, %v", err)
	}

	d.log.Println("QUERYING BALANCE SHEET ", out.Count)
	d.log.Println("QUERYING BALANCE SHEET ", out.Items)
	d.log.Println("QUERYING BALANCE SHEET ", out.Items[0]["fiscalDate"].(*types.AttributeValueMemberS).Value)

	bs := &pb.BalanceSheet{}
	return bs, nil
}

func (d *Dynamo) InsertBalanceSheet(ctx context.Context, ticker *pb.Ticker) error {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		d.log.Fatalf("unable to load SDK config, %v", err)
	}

	// Using the Config value, create the DynamoDB client
	client := dynamodb.NewFromConfig(cfg)

	symbol := ticker.Symbol
	bs, _ := d.iex.GetBalanceSheet(ctx, symbol)
	d.log.Println(bs)
	input := &dynamodb.PutItemInput{
		Item: map[string]types.AttributeValue{
			"symbol":                  &types.AttributeValueMemberS{Value: symbol},
			"fiscalDate":              &types.AttributeValueMemberS{Value: bs.FiscalDate.String()},
			"currency":                &types.AttributeValueMemberS{Value: bs.Currency},
			"goodwill":                &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.Goodwill, 'f', -1, 64)},
			"commonStock":             &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.CommonStock, 'f', -1, 64)},
			"propertyPlantEquipment":  &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.PropertyPlantEquipment, 'f', -1, 64)},
			"retainedEarnings":        &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.RetainedEarnings, 'f', -1, 64)},
			"totalLiabilities":        &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.TotalLiabilities, 'f', -1, 64)},
			"totalAssets":             &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.TotalAssets, 'f', -1, 64)},
			"otherAssets":             &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.OtherAssets, 'f', -1, 64)},
			"currentCash":             &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.CurrentCash, 'f', -1, 64)},
			"currentLongTermDebt":     &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.CurrentLongTermDebt, 'f', -1, 64)},
			"longTermInvestments":     &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.LongTermInvestments, 'f', -1, 64)},
			"fiscalYear":              &types.AttributeValueMemberN{Value: strconv.FormatInt(int64(bs.FiscalYear), 10)},
			"otherCurrentLiabilities": &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.OtherCurrentLiabilities, 'f', -1, 64)},
			"shareholderEquity":       &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.ShareholderEquity, 'f', -1, 64)},
			"netTangibleAssets":       &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.NetTangibleAssets, 'f', -1, 64)},
			"intangibleAssets":        &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.IntangibleAssets, 'f', -1, 64)},
			"inventory":               &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.Inventory, 'f', -1, 64)},
			"accountsPayable":         &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.AccountsPayable, 'f', -1, 64)},
			"capitalSurplus":          &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.CapitalSurplus, 'f', -1, 64)},
			"otherLiabilities":        &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.OtherLiablities, 'f', -1, 64)},
			"otherCurrentAssets":      &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.OtherCurrentAssets, 'f', -1, 64)},
			"fiscalQuarter":           &types.AttributeValueMemberN{Value: strconv.FormatInt(int64(bs.FiscalQuarter), 10)},
			"minorityInterest":        &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.MinorityInterest, 'f', -1, 64)},
			"receivables":             &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.Receivables, 'f', -1, 64)},
			"shortTermInvestments":    &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.ShortTermInvestments, 'f', -1, 64)},
			"longTermDebt":            &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.LongTermDebt, 'f', -1, 64)},
			"totalCurrentLiabilities": &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.TotalCurrentLiabilities, 'f', -1, 64)},
			"currentAssets":           &types.AttributeValueMemberN{Value: strconv.FormatFloat(bs.CurrentAssets, 'f', -1, 64)},
			"reportDate":              &types.AttributeValueMemberS{Value: bs.ReportDate.String()},
			"filingType":              &types.AttributeValueMemberS{Value: bs.FilingType},
		},
		TableName: aws.String("BalanceSheet"),
	}

	_, err = client.PutItem(ctx, input)
	if err != nil {
		d.log.Printf("unable to insert item, %v", err)
	}

	return err
}
