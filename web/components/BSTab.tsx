import { Timestamp } from "google-protobuf/google/protobuf/timestamp_pb";
import React from "react";
import style from "./FinTable.module.css";

import { BalanceSheet } from "../pages/finsvc_pb";

const BSTab: React.FC<{f: BalanceSheet[] | null}> = (props) => {
    return (
        <>
        <table className={style.table}>
            <tbody>
                <tr className={style.labels}>
                    <td className={style.years}>Fiscal Year</td>
                    <td>Reported Date</td>
                    <td>Filing Type</td>
                    <td>Fiscal Date</td>
                    <td>Fiscal Quarter</td>
                    <td>Currency</td>
                    <td>Current Cash</td>
                    <td>Short Term Investments</td>
                    <td>Receivables</td>
                    <td>Inventory</td>
                    <td>Other Current Assets</td>
                    <td>Total Current Assets</td>
                    <td>Long Term Investments</td>
                    <td>Property Plant and Equipment</td>
                    <td>Goodwill</td>
                    <td>Intangible Assets</td>
                    <td>Other Assets</td> 
                    <td>Total Assets</td>
                    <td>Accounts Payable</td>
                    <td>Current Long Term Debt</td>
                    <td>Other Current Liabilities</td>
                    <td>Total Current Liabilities</td>
                    <td>Long Term Debt</td>
                    <td>Other Liabilities</td>
                    <td>Minority Interest</td>
                    <td>Total Liabilities</td>
                    <td>Common Stock</td>
                    <td>Retained Earnings</td>
                    <td>Treasury Stock</td>
                    <td>Capital Surplus</td>
                    <td>Shareholder Equity</td>
                    <td>Net Tangible Assets</td>
                </tr>
            {props.f?.map(bs => {
                let reportDateTS = bs.getReportdate() as Timestamp;
                let reportDate = reportDateTS.toDate().toISOString().split("T")[0];

                const fiscalDateTS = bs.getFiscaldate() as Timestamp;
                // format localDate to YYYY-MM-DD in UTC
                let fiscalDate = fiscalDateTS.toDate().toISOString().split("T")[0];

                return (
                <tr key={bs.getFiscalyear()} className={style.data}>
                    <td className={style.years}>{bs.getFiscalyear()}</td>
                    <td>{reportDate}</td>
                    <td>{bs.getFilingtype()}</td>
                    <td>{fiscalDate}</td>
                    <td>{bs.getFiscalquarter()}</td>
                    <td>{bs.getCurrency()}</td>
                    <td>{bs.getCurrentcash()}</td>
                    <td>{bs.getShortterminvestments()}</td>
                    <td>{bs.getReceivables()}</td>
                    <td>{bs.getInventory()}</td>
                    <td>{bs.getOthercurrentassets()}</td>
                    <td>{bs.getCurrentassets()}</td>
                    <td>{bs.getLongterminvestments()}</td>
                    <td>{bs.getPropertyplantequipment()}</td>
                    <td>{bs.getGoodwill()}</td>
                    <td>{bs.getIntangibleassets()}</td>
                    <td>{bs.getOtherassets()}</td>
                    <td>{bs.getTotalassets()}</td>
                    <td>{bs.getAccountspayable()}</td>
                    <td>{bs.getCurrentlongtermdebt()}</td>
                    <td>{bs.getOthercurrentliabilities()}</td>
                    <td>{bs.getTotalcurrentliabilities()}</td>
                    <td>{bs.getLongtermdebt()}</td>
                    <td>{bs.getOtherliabilities()}</td>
                    <td>{bs.getMinorityinterest()}</td>
                    <td>{bs.getTotalliabilities()}</td>
                    <td>{bs.getCommonstock()}</td>
                    <td>{bs.getRetainedearnings()}</td>
                    <td>{bs.getTreasurystock()}</td>
                    <td>{bs.getCapitalsurplus()}</td>
                    <td>{bs.getShareholderequity()}</td>
                    <td>{bs.getNettangibleassets()}</td>
                </tr>
                );
            })}
            </tbody>
        </table>


        </>
    );
};

export default BSTab;