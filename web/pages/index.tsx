import type { NextPage } from 'next'
import Head from 'next/head'
import styles from '../styles/Home.module.css'
import {useEffect} from "react";
import * as grpcWeb from "grpc-web";
import {GetRequest, BalanceSheet, Financials, IncomeStatement, CashFlow, IncomeStatements} from "./finsvc_pb";
import {FinSvcClient} from "./FinsvcServiceClientPb";
import React from 'react';
import BSTab from '../components/BSTab';
import ISTab from '../components/ISTab';
import CFTab from '../components/CFTab';
import Tabs from '../components/Tabs';
import Tab from '../components/Tab';

interface Props {
  query: {
    tab?: string
  };
}
const Home: NextPage<Props> = ( {query} ) => {
  const [financials, setFinancials] = React.useState<Financials | null>(null);
  const [incomeStatements, setIncomeStatements] = React.useState<IncomeStatement[] | null>(null);
  const [balanceSheets, setBalanceSheets] = React.useState<BalanceSheet[] | null>(null);
  const [cashFlows, setCashFlows] = React.useState<CashFlow[] | null>(null);
  const [symbol, setSymbol] = React.useState<string>("TSLA");

  const symbolHandler = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.key != "Enter")
      return;

    event.preventDefault();
    console.log(event.currentTarget.value)
    setSymbol(event.currentTarget.value)
  }

  useEffect(() => {
    const finService = new FinSvcClient("http://localhost:8080", null, null);
    const req = new GetRequest();
    req.setSymbol(symbol);
    req.setLimit(4);
    
    const call2 = finService.getAnnualFinancials(req, {},
      (err: grpcWeb.RpcError, response: Financials) => {
        if (err) {
          if (err.code !== grpcWeb.StatusCode.OK) {
              console.log("We got an error: ", err.message);
          }
        } else {
          let fin = response
          setFinancials(fin);
          let bs = financials?.getBalancesheets()?.getBalancesheetsList() as BalanceSheet[];
          let is = financials?.getIncomestatements()?.getIncomestatementsList() as IncomeStatement[];
          let cf = financials?.getCashflows()?.getCashflowsList() as CashFlow[];

          setIncomeStatements(is);
          setBalanceSheets(bs);
          setCashFlows(cf);

          console.log(fin);
        }
      }
    );
    

  }, [symbol]);

  return (
    <div className={styles.container}>
      <Head>
        <title>Financials</title>
        <meta name="description" content="Financial Service" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className={styles.main}>

        <div>
          <input className={styles.input} onKeyDown={symbolHandler} type="text" placeholder="Enter a stock symbol" />
        </div>
        <Tabs initialTab={query}>
          <Tab tab="I/S">
            <ISTab f={incomeStatements}/> 
          </Tab>
          <Tab tab="B/S">
            <BSTab  f={balanceSheets} /> 
          </Tab>
          <Tab tab="C/F">
            <CFTab f={cashFlows}/> 
          </Tab>
        </Tabs>
        {/* <div>
          <p style={{color:"white"}}>{symbol}</p>
          <p style={{color:"white"}}>{incomeStatements?.at(0)?.getNetincome()}</p>
        </div> */}
      </main>


      <footer className={styles.footer}>
        <a
          href="https://vercel.com?utm_source=create-next-app&utm_medium=default-template&utm_campaign=create-next-app"
          target="_blank"
          rel="noopener noreferrer"
        >
          Powered by DDI
        </a>
      </footer>
    </div>
  )
}

Home.getInitialProps = ({query}) => {
  return {query};
}

export default Home;
