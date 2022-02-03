package main

import (
	"context"
	"fin/datasource"
	"fin/store"
	"fin/svc"
	"log"
)

func main() {
	log := log.Default()
	ds := datasource.NewIEX(log)
	store, _ := store.NewDynamo(log)
	svc := svc.NewSvc(log, store, ds)
	svc.GetAnualBalanceSheet(context.Background(), "AAPL", 1)

	// Using the SDK's default configuration, loading additional config
	// and credentials values from the environment variables, shared
	// credentials, and shared configuration files
	//cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
	/*
		cfg, err := config.LoadDefaultConfig(context.TODO())
		if err != nil {
			log.Fatalf("unable to load SDK config, %v", err)
		}

		// Using the Config value, create the DynamoDB client
		dsvc := dynamodb.NewFromConfig(cfg)

		getItemInput := &dynamodb.GetItemInput{
			Key: map[string]types.AttributeValue{
				"symbol":     &types.AttributeValueMemberS{Value: "AAPL"},
				"fiscalDate": &types.AttributeValueMemberS{Value: "2021-09-11"},
			},
			TableName: aws.String("BalanceSheet"),
		}

		resp, err := dsvc.GetItem(context.Background(), getItemInput)
		if err != nil {
			log.Fatalf("unable to get item, %v", err)
		}

		item := struct {
			AccountsPayable         float64 `dynamodbav:"accountsPayable" json:"accountsPayable"`
			CapitalSurplus          float64 `dynamodbav:"capitalSurplus" json:"capitalSurplus"`
			CommonStock             float64 `dynamodbav:"commonStock" json:"commonStock"`
			Currency                string  `dynamodbav:"currency" json:"currency"`
			CurrentAssets           float64 `dynamodbav:"currentAssets" json:"currentAssets"`
			CurrentCash             float64 `dynamodbav:"currentCash" json:"currentCash"`
			CurrentLongTermDebt     float64 `dynamodbav:"currentLongTermDebt" json:"currentLongTermDebt"`
			FilingType              string  `dynamodbav:"filingType" json:"filingType"`
			FiscalDate              string  `dynamodbav:"fiscalDate" json:"fiscalDate"`
			FiscalQuarter           float64 `dynamodbav:"fiscalQuarter" json:"fiscalQuarter"`
			FiscalYear              float64 `dynamodbav:"fiscalYear" json:"fiscalYear"`
			Goodwill                float64 `dynamodbav:"goodwill" json:"goodwill"`
			IntangibleAssets        float64 `dynamodbav:"intangibleAssets" json:"intangibleAssets"`
			Inventory               float64 `dynamodbav:"inventory" json:"inventory"`
			LongTermDebt            float64 `dynamodbav:"longTermDebt" json:"longTermDebt"`
			LongTermInvestments     float64 `dynamodbav:"longTermInvestments" json:"longTermInvestments"`
			MinorityInterest        float64 `dynamodbav:"minorityInterest" json:"minorityInterest"`
			NetTangibleAssets       float64 `dynamodbav:"netTangibleAssets" json:"netTangibleAssets"`
			OtherAssets             float64 `dynamodbav:"otherAssets" json:"otherAssets"`
			OtherCurrentAssets      float64 `dynamodbav:"otherCurrentAssets" json:"otherCurrentAssets"`
			OtherCurrentLiabilities float64 `dynamodbav:"otherCurrentLiabilities" json:"otherCurrentLiabilities"`
			OtherLiabilities        float64 `dynamodbav:"otherLiabilities" json:"otherLiabilities"`
			PropertyPlantEquipment  float64 `dynamodbav:"propertyPlantEquipment" json:"propertyPlantEquipment"`
			Receivables             float64 `dynamodbav:"receivables" json:"receivables"`
			ReportDate              string  `dynamodbav:"reportDate" json:"reportDate"`
			RetainedEarnings        float64 `dynamodbav:"retainedEarnings" json:"retainedEarnings"`
			ShareholderEquity       float64 `dynamodbav:"shareholderEquity" json:"shareholderEquity"`
			ShortTermInvestments    float64 `dynamodbav:"shortTermInvestments" json:"shortTermInvestments"`
			Symbol                  string  `dynamodbav:"symbol" json:"symbol"`
			TotalAssets             float64 `dynamodbav:"totalAssets" json:"totalAssets"`
			TotalCurrentLiabilities float64 `dynamodbav:"totalCurrentLiabilities" json:"totalCurrentLiabilities"`
			TotalLiabilities        float64 `dynamodbav:"totalLiabilities" json:"totalLiabilities"`
		}{}

		err = attributevalue.UnmarshalMap(resp.Item, &item)
		if err != nil {
			log.Fatalf("unable to unmarshal, %v", err)
		} else {
			ijson, _ := json.MarshalIndent(item, "", "  ")
			log.Printf("\n\n\n%s\n\n\n", string(ijson))
		}

		lg := log.Default()
		dyn := dynamo.NewDynamo(lg)
		//dyn.GetBalanceSheet(context.Background(), &pb.Ticker{Symbol: "AAPL"})
		log.Println("GETTING STUFF FROM IEX")
	*/

	/*
		log := log.Default()

		svr := grpc.NewServer()
		pb.RegisterFinSvcServer(svr, dyn)

		reflection.Register(svr)

		l, err := net.Listen("tcp", ":9090")
		if err != nil {
			log.Fatalf("Error listening: %s", err)
			os.Exit(1)
		}
		svr.Serve(l)
	*/
}
