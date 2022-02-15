import { Timestamp } from "google-protobuf/google/protobuf/timestamp_pb";
import React from "react";
import style from "./FinTable.module.css";

import { IncomeStatement } from "../pages/finsvc_pb";

const ISTab: React.FC<{f: IncomeStatement[] | null}> = (props) => {
    return (
        <>
        <table className={style.table}>
            <tbody>
                <tr className={style.labels}>
                    <td className={style.years}>Fiscal Date</td>
                    <td>Reported Date</td>
                    <td>Currency</td>
                    <td>Total Revenue</td>
                    <td>Cost of Revenue</td>
                    <td>Gross Profit</td>
                    <td>Research and Development</td>
                    <td>Selling General and Admin</td>
                    <td>Operating Expense</td>
                    <td>Operating Income</td>
                    <td>Other Income Expense Net</td>
                    <td>EBIT</td>
                    <td>Interest Income</td>
                    <td>Pre-tax Income</td>
                    <td>Income Tax</td>
                    <td>Minority Interest</td>
                    <td>Net Income</td>
                    <td>Net Income Basic</td>
                </tr>
            {props.f?.map(is => {
                let reportDateTS = is.getReportdate() as Timestamp;
                let reportDate = reportDateTS.toDate().toISOString().split("T")[0];

                const fiscalDateTS = is.getFiscaldate() as Timestamp;
                // format localDate to YYYY-MM-DD in UTC
                let fiscalDate = fiscalDateTS.toDate().toISOString().split("T")[0];

                return (
                <tr key={fiscalDate} className={style.data}>
                    <td className={style.years}>{fiscalDate}</td>
                    <td>{reportDate}</td>
                    <td>{is.getCurrency()}</td>
                    <td>{is.getTotalrevenue()}</td>
                    <td>{is.getCostofrevenue()}</td>
                    <td>{is.getGrossprofit()}</td>
                    <td>{is.getResearchanddevelopment()}</td>
                    <td>{is.getSellinggeneralandadmin()}</td>
                    <td>{is.getOperatingexpense()}</td>
                    <td>{is.getOperatingincome()}</td>
                    <td>{is.getOtherincomeexpensenet()}</td>
                    <td>{is.getEbit()}</td>
                    <td>{is.getInterestincome()}</td>
                    <td>{is.getPretaxincome()}</td>
                    <td>{is.getIncometax()}</td>
                    <td>{is.getMinorityinterest()}</td>
                    <td>{is.getNetincome()}</td>
                    <td>{is.getNetincomebasic()}</td>
                </tr>
                );
            })}
            </tbody>
        </table>


        </>
    );
};

export default ISTab;