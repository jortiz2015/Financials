package store

import (
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"google.golang.org/protobuf/types/known/timestamppb"

	//"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	//"encoding/json"
	"context"
	pb "finsvc/pb"
	"log"
)

type Dynamo struct {
	log log.Logger
}

func NewDynamo(log *log.Logger) *Dynamo {
	return &Dynamo{log: *log}
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

	getItemInput := &dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{
			"symbol":     &types.AttributeValueMemberS{Value: "AAPL"},
			"fiscalDate": &types.AttributeValueMemberS{Value: "2021-09-11"},
		},
		TableName: aws.String("BalanceSheet"),
	}

	resp, err := client.GetItem(context.Background(), getItemInput)
	if err != nil {
		d.log.Fatalf("unable to get item, %v", err)
	}

	bs, _ := d.setBalanceSheet(resp.Item)
	return bs, nil
}
