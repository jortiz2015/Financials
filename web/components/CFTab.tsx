import { Timestamp } from "google-protobuf/google/protobuf/timestamp_pb";
import React from "react";
import style from "./FinTable.module.css";

import { CashFlow } from "../pages/finsvc_pb";

const CFTab: React.FC<{f: CashFlow[] | null}> = (props) => {
    return (
        <>
        <table className={style.table}>
            <tbody>
                <tr className={style.labels}>
                    <td className={style.years}>Fiscal Date</td>
                    <td>Reported Date</td>
                    <td>Currency</td>
                    <td>Net Income</td>
                    <td>Depreciation</td>
                    <td>Changes In Receivables</td>
                    <td>Changes In Inventories</td>
                    <td>Cash Change</td>
                    <td>Cash Flow</td>
                    <td>Capital Expenditures</td>
                    <td>Inbestments</td>
                    <td>Investing Activity Other</td>
                    <td>Total Investing Cash flows</td>
                    <td>Dividends Paid</td>
                    <td>Net Borrowings</td>
                    <td>Other Financing Cash Flows</td>
                    <td>Cash Flow Financing</td>
                    <td>Exchange Rate Effect</td>   
                </tr>
            {props.f?.map(cf => {
                let reportDateTS = cf.getReportdate() as Timestamp;
                let reportDate = reportDateTS.toDate().toISOString().split("T")[0];

                const fiscalDateTS = cf.getFiscaldate() as Timestamp;
                // format localDate to YYYY-MM-DD in UTC
                let fiscalDate = fiscalDateTS.toDate().toISOString().split("T")[0];

                return (
                <tr key={fiscalDate} className={style.data}>
                    <td className={style.years}>{fiscalDate}</td>
                    <td>{reportDate}</td>
                    <td>{cf.getCurrency()}</td>
                    <td>{cf.getNetincome()}</td>
                    <td>{cf.getDepreciation()}</td>
                    <td>{cf.getChangesinreceivables()}</td>
                    <td>{cf.getChangesininventories()}</td>
                    <td>{cf.getCashchange()}</td>
                    <td>{cf.getCashflow()}</td>
                    <td>{cf.getCapitalexpenditures()}</td>
                    <td>{cf.getInvestments()}</td>
                    <td>{cf.getInvestingactivityother()}</td>
                    <td>{cf.getTotalinvestingcashflows()}</td>
                    <td>{cf.getDividendspaid()}</td>
                    <td>{cf.getNetborrowings()}</td>
                    <td>{cf.getOtherfinancingcashflows()}</td>
                    <td>{cf.getCashflowfinancing()}</td>
                    <td>{cf.getExchangerateeffect()}</td>
                </tr>
                );
            })}
            </tbody>
        </table>


        </>
    );
};

export default CFTab;