import type { NextPage } from 'next'
import Head from 'next/head'
import Image from 'next/image'
import styles from '../styles/Home.module.css'
import {useEffect} from "react";
import * as grpcWeb from "grpc-web";
import {GetRequest, BalanceSheets, BalanceSheet} from "./finsvc_pb";
import {FinSvcClient} from "./FinsvcServiceClientPb";
import React from 'react';
import FinTable from '../components/FinTable';

const Home: NextPage = () => {
  const [balanceSheets, setBalanceSheets] = React.useState<BalanceSheet[] | null>(null);
  const [symbol, setSymbol] = React.useState<string>("FB");

  const symbolHandler = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.key != "Enter")
      return;

    event.preventDefault();
    console.log(event.currentTarget.value)
    //console.log(event.target.value);
    setSymbol(event.currentTarget.value)
  }

  useEffect(() => {
      const finService = new FinSvcClient("http://localhost:8080", null, null);
      const req = new GetRequest();
      req.setSymbol(symbol);
      req.setLimit(4);
      const call = finService.getAnnualBalanceSheets(req, {},
        (err: grpcWeb.RpcError, response: BalanceSheets) => {
          if (err) {
            if (err.code !== grpcWeb.StatusCode.OK) {
                console.log("We got an error: ", err.message);
            }
          } else {
            let bsl = response.getBalancesheetsList();
            setBalanceSheets(bsl);
            console.log(bsl);
          }
        }
      );
  }, [symbol]);

  const currency = balanceSheets?.map(bs =>
    <div key={bs.getFiscalyear()}>{bs.getFiscalyear() }</div>
  );

/*
        <div>
          {balanceSheets?.map((bs: BalanceSheet) =>  
            <div key={bs.getFiscalyear()}>{bs.getFiscaldate()}</div>
          )}
        </div>
*/

  return (
    <div className={styles.container}>
      <Head>
        <title>Financials</title>
        <meta name="description" content="Financial Service" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className={styles.main}>
        <h1 className={styles.title}>
          Financial Service
        </h1>

        <div className={styles.description}>
          Get financial data for a stock symbol.
          <code className={styles.code}>Ex: "TSLA"</code>
          <div>
            <input className={styles.input} onKeyDown={symbolHandler} type="text" placeholder="Enter a stock symbol" />
          </div>
        </div>


        <div className={styles.card}>
          <FinTable f={balanceSheets} ></FinTable>
        </div>

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

export default Home
