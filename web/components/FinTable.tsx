import { Timestamp } from "google-protobuf/google/protobuf/timestamp_pb";
import React from "react";
import { BalanceSheet, BalanceSheets } from "../pages/finsvc_pb";

const FinTable: React.FC<{f: BalanceSheet[] | null}> = (props) => {
    const years = props.f?.map(bs =>
        <th key={bs.getFiscalyear()}>{bs.getFiscalyear()}</th>
    );

    const reportDate = props.f?.map(bs => {
        let ts = bs.getReportdate() as Timestamp;
        let formattedDate = ts.toDate().toISOString().split("T")[0];
        //console.log(raw);
        return <td key={bs.getFiscalyear()}>{formattedDate}</td>
    });

    const fiscalDate = props.f?.map(bs => {
        const ts = bs.getFiscaldate() as Timestamp;
        // format localDate to YYYY-MM-DD in UTC
        let formattedDate = ts.toDate().toISOString().split("T")[0];
        return <td key={bs.getFiscalyear()}>{formattedDate}</td>
    });

    const fiscalQuarter = props.f?.map(bs =>
        <td key={bs.getFiscalyear()}>{bs.getFiscalquarter()}</td>
    );

    const fiscalYear = props.f?.map(bs =>
        <td key={bs.getFiscalyear()}>{bs.getFiscalyear()}</td>
    );

    const currency = props.f?.map(bs =>
        <td key={bs.getFiscalyear()}>{bs.getCurrency()}</td>
    );

    const currentCash = props.f?.map(bs =>
        <td key={bs.getFiscalyear()}>{bs.getCurrentcash()}</td>
    );

    const shortTermInvestments = props.f?.map(bs =>
        <td key={bs.getFiscalyear()}>{bs.getShortterminvestments()}</td>
    );

    const receivables = props.f?.map(bs =>
        <td key={bs.getFiscalyear()}>{bs.getReceivables()}</td>
    );

    const inventory = props.f?.map(bs =>
        <td key={bs.getFiscalyear()}>{bs.getInventory()}</td>
    );

    const otherCurrentAssets = props.f?.map(bs =>
        <td key={bs.getFiscalyear()}>{bs.getOthercurrentassets()}</td>
    );

    const currentAssets = props.f?.map(bs =>
        <td key={bs.getFiscalyear()}>{bs.getCurrentassets()}</td>
    );

    const longTermInvestments = props.f?.map(bs =>
        <td key={bs.getFiscalyear()}>{bs.getLongterminvestments()}</td>
    );

    const propertyPlantEquipment = props.f?.map(bs =>
        <td key={bs.getFiscalyear()}>{bs.getPropertyplantequipment()}</td>
    );

    const goodwill = props.f?.map(bs =>
        <td key={bs.getFiscalyear()}>{bs.getGoodwill()}</td>
    );

    const intangibleAssets = props.f?.map(bs =>
        <td key={bs.getFiscalyear()}>{bs.getIntangibleassets()}</td>
    );

    const otherAssets = props.f?.map(bs =>
        <td key={bs.getFiscalyear()}>{bs.getOtherassets()}</td>
    );

    const totalAssets = props.f?.map(bs =>
        <td key={bs.getFiscalyear()}>{bs.getTotalassets()}</td>
    );

    const accountsPayable = props.f?.map(bs =>
        <td key={bs.getFiscalyear()}>{bs.getAccountspayable()}</td>
    );

    const currentLongTermDebt = props.f?.map(bs =>
        <td key={bs.getFiscalyear()}>{bs.getCurrentlongtermdebt()}</td>
    );

    const otherCurrentLiabilities = props.f?.map(bs =>
        <td key={bs.getFiscalyear()}>{bs.getOthercurrentliabilities()}</td>
    );

    const totalCurrentLiabilities = props.f?.map(bs =>
        <td key={bs.getFiscalyear()}>{bs.getTotalcurrentliabilities()}</td>
    );

    const longTermDebt = props.f?.map(bs =>
        <td key={bs.getFiscalyear()}>{bs.getLongtermdebt()}</td>
    );

    const otherLiabilities = props.f?.map(bs =>
        <td key={bs.getFiscalyear()}>{bs.getOtherliabilities()}</td>
    );

    const minorityInterest = props.f?.map(bs =>
        <td key={bs.getFiscalyear()}>{bs.getMinorityinterest()}</td>
    );

    const totalLiabilities = props.f?.map(bs =>
        <td key={bs.getFiscalyear()}>{bs.getTotalliabilities()}</td>
    );

    const commonStock = props.f?.map(bs =>
        <td key={bs.getFiscalyear()}>{bs.getCommonstock()}</td>
    );

    const retainedEarnings = props.f?.map(bs =>
        <td key={bs.getFiscalyear()}>{bs.getRetainedearnings()}</td>
    );

    const treasuryStock = props.f?.map(bs =>
        <td key={bs.getFiscalyear()}>{bs.getTreasurystock()}</td>
    );

    const capitalSurplus = props.f?.map(bs =>
        <td key={bs.getFiscalyear()}>{bs.getCapitalsurplus()}</td>
    );

    const shareholderEquity = props.f?.map(bs =>
        <td key={bs.getFiscalyear()}>{bs.getShareholderequity()}</td>
    );

    const netTangibleAssets = props.f?.map(bs =>
        <td key={bs.getFiscalyear()}>{bs.getNettangibleassets()}</td>
    );

    return (
        <table>
            <thead>
                <tr>
                    <th>Items</th>
                    {years}
                </tr>
            </thead>
            <tbody>
                <tr>
                    <td>Report Date</td>
                    {reportDate}
                </tr>
                <tr>
                    <td>Filing Type</td>
                    {props.f?.map(bs =>
                        <td key={bs.getFiscalyear()}>{bs.getFilingtype()}</td>
                    )}
                </tr>
                <tr>
                    <td>Fiscal Date</td>
                    {fiscalDate}
                </tr>
                <tr>
                    <td>Fiscal Quarter</td>
                    {fiscalQuarter}
                </tr>
                <tr>
                    <td>Fiscal Year</td>
                    {fiscalYear}
                </tr>
                <tr>
                    <td>Currency</td>
                    {currency}
                </tr>
                <tr>
                    <td>Current Cash</td>
                    {currentCash}
                </tr>
                <tr>
                    <td>Short Term Investments</td>
                    {shortTermInvestments}
                </tr>
                <tr>
                    <td>Receivables</td>
                    {receivables}
                </tr>
                <tr>
                    <td>Inventory</td>
                    {inventory}
                </tr>
                <tr>
                    <td>Other Current Assets</td>
                    {otherCurrentAssets}
                </tr>
                <tr>
                    <td>Current Assets</td>
                    {currentAssets}
                </tr>
                <tr>
                    <td>Long Term Investments</td>
                    {longTermInvestments}
                </tr>
                <tr>
                    <td>Property, Plant & Equipment</td>
                    {propertyPlantEquipment}
                </tr>
                <tr>
                    <td>Goodwill</td>
                    {goodwill}
                </tr>
                <tr>
                    <td>Intangible Assets</td>
                    {intangibleAssets}
                </tr>
                <tr>
                    <td>Other Assets</td>
                    {otherAssets}
                </tr>
                <tr>
                    <td>Total Assets</td>
                    {totalAssets}
                </tr>
                <tr>
                    <td>Accounts Payable</td>
                    {accountsPayable}
                </tr>
                <tr>
                    <td>Current Long Term Debt</td>
                    {currentLongTermDebt}
                </tr>
                <tr>
                    <td>Other Current Liabilities</td>
                    {otherCurrentLiabilities}
                </tr>
                <tr>
                    <td>Total Current Liabilities</td>
                    {totalCurrentLiabilities}
                </tr>
                <tr>
                    <td>Long Term Debt</td>
                    {longTermDebt}
                </tr>
                <tr>
                    <td>Other Liabilities</td>
                    {otherLiabilities}
                </tr>
                <tr>
                    <td>Minority Interest</td>
                    {minorityInterest}
                </tr>
                <tr>
                    <td>Total Liabilities</td>
                    {totalLiabilities}
                </tr>
                <tr>
                    <td>Common Stock</td>
                    {commonStock}
                </tr>
                <tr>
                    <td>Retained Earnings</td>
                    {retainedEarnings}
                </tr>
                <tr>
                    <td>Treasury Stock</td>
                    {treasuryStock}
                </tr>
                <tr>
                    <td>Capital Surplus</td>
                    {capitalSurplus}
                </tr>
                <tr>
                    <td>Shareholder Equity</td>
                    {shareholderEquity}
                </tr>
                <tr>
                    <td>Net Tangible Assets</td>
                    {netTangibleAssets}
                </tr>
            </tbody>
        </table>
    );
};

export default FinTable;