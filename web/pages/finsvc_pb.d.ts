import * as jspb from 'google-protobuf'

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';


export class GetRequest extends jspb.Message {
  getSymbol(): string;
  setSymbol(value: string): GetRequest;

  getLimit(): number;
  setLimit(value: number): GetRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetRequest): GetRequest.AsObject;
  static serializeBinaryToWriter(message: GetRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetRequest;
  static deserializeBinaryFromReader(message: GetRequest, reader: jspb.BinaryReader): GetRequest;
}

export namespace GetRequest {
  export type AsObject = {
    symbol: string,
    limit: number,
  }
}

export class BalanceSheets extends jspb.Message {
  getBalancesheetsList(): Array<BalanceSheet>;
  setBalancesheetsList(value: Array<BalanceSheet>): BalanceSheets;
  clearBalancesheetsList(): BalanceSheets;
  addBalancesheets(value?: BalanceSheet, index?: number): BalanceSheet;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BalanceSheets.AsObject;
  static toObject(includeInstance: boolean, msg: BalanceSheets): BalanceSheets.AsObject;
  static serializeBinaryToWriter(message: BalanceSheets, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BalanceSheets;
  static deserializeBinaryFromReader(message: BalanceSheets, reader: jspb.BinaryReader): BalanceSheets;
}

export namespace BalanceSheets {
  export type AsObject = {
    balancesheetsList: Array<BalanceSheet.AsObject>,
  }
}

export class BalanceSheet extends jspb.Message {
  getReportdate(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setReportdate(value?: google_protobuf_timestamp_pb.Timestamp): BalanceSheet;
  hasReportdate(): boolean;
  clearReportdate(): BalanceSheet;

  getFilingtype(): string;
  setFilingtype(value: string): BalanceSheet;

  getFiscaldate(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setFiscaldate(value?: google_protobuf_timestamp_pb.Timestamp): BalanceSheet;
  hasFiscaldate(): boolean;
  clearFiscaldate(): BalanceSheet;

  getFiscalquarter(): number;
  setFiscalquarter(value: number): BalanceSheet;

  getFiscalyear(): number;
  setFiscalyear(value: number): BalanceSheet;

  getCurrency(): string;
  setCurrency(value: string): BalanceSheet;

  getCurrentcash(): number;
  setCurrentcash(value: number): BalanceSheet;

  getShortterminvestments(): number;
  setShortterminvestments(value: number): BalanceSheet;

  getReceivables(): number;
  setReceivables(value: number): BalanceSheet;

  getInventory(): number;
  setInventory(value: number): BalanceSheet;

  getOthercurrentassets(): number;
  setOthercurrentassets(value: number): BalanceSheet;

  getCurrentassets(): number;
  setCurrentassets(value: number): BalanceSheet;

  getLongterminvestments(): number;
  setLongterminvestments(value: number): BalanceSheet;

  getPropertyplantequipment(): number;
  setPropertyplantequipment(value: number): BalanceSheet;

  getGoodwill(): number;
  setGoodwill(value: number): BalanceSheet;

  getIntangibleassets(): number;
  setIntangibleassets(value: number): BalanceSheet;

  getOtherassets(): number;
  setOtherassets(value: number): BalanceSheet;

  getTotalassets(): number;
  setTotalassets(value: number): BalanceSheet;

  getAccountspayable(): number;
  setAccountspayable(value: number): BalanceSheet;

  getCurrentlongtermdebt(): number;
  setCurrentlongtermdebt(value: number): BalanceSheet;

  getOthercurrentliabilities(): number;
  setOthercurrentliabilities(value: number): BalanceSheet;

  getTotalcurrentliabilities(): number;
  setTotalcurrentliabilities(value: number): BalanceSheet;

  getLongtermdebt(): number;
  setLongtermdebt(value: number): BalanceSheet;

  getOtherliabilities(): number;
  setOtherliabilities(value: number): BalanceSheet;

  getMinorityinterest(): number;
  setMinorityinterest(value: number): BalanceSheet;

  getTotalliabilities(): number;
  setTotalliabilities(value: number): BalanceSheet;

  getCommonstock(): number;
  setCommonstock(value: number): BalanceSheet;

  getRetainedearnings(): number;
  setRetainedearnings(value: number): BalanceSheet;

  getTreasurystock(): number;
  setTreasurystock(value: number): BalanceSheet;

  getCapitalsurplus(): number;
  setCapitalsurplus(value: number): BalanceSheet;

  getShareholderequity(): number;
  setShareholderequity(value: number): BalanceSheet;

  getNettangibleassets(): number;
  setNettangibleassets(value: number): BalanceSheet;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BalanceSheet.AsObject;
  static toObject(includeInstance: boolean, msg: BalanceSheet): BalanceSheet.AsObject;
  static serializeBinaryToWriter(message: BalanceSheet, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BalanceSheet;
  static deserializeBinaryFromReader(message: BalanceSheet, reader: jspb.BinaryReader): BalanceSheet;
}

export namespace BalanceSheet {
  export type AsObject = {
    reportdate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    filingtype: string,
    fiscaldate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    fiscalquarter: number,
    fiscalyear: number,
    currency: string,
    currentcash: number,
    shortterminvestments: number,
    receivables: number,
    inventory: number,
    othercurrentassets: number,
    currentassets: number,
    longterminvestments: number,
    propertyplantequipment: number,
    goodwill: number,
    intangibleassets: number,
    otherassets: number,
    totalassets: number,
    accountspayable: number,
    currentlongtermdebt: number,
    othercurrentliabilities: number,
    totalcurrentliabilities: number,
    longtermdebt: number,
    otherliabilities: number,
    minorityinterest: number,
    totalliabilities: number,
    commonstock: number,
    retainedearnings: number,
    treasurystock: number,
    capitalsurplus: number,
    shareholderequity: number,
    nettangibleassets: number,
  }
}

export class IncomeStatements extends jspb.Message {
  getIncomestatementsList(): Array<IncomeStatement>;
  setIncomestatementsList(value: Array<IncomeStatement>): IncomeStatements;
  clearIncomestatementsList(): IncomeStatements;
  addIncomestatements(value?: IncomeStatement, index?: number): IncomeStatement;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): IncomeStatements.AsObject;
  static toObject(includeInstance: boolean, msg: IncomeStatements): IncomeStatements.AsObject;
  static serializeBinaryToWriter(message: IncomeStatements, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): IncomeStatements;
  static deserializeBinaryFromReader(message: IncomeStatements, reader: jspb.BinaryReader): IncomeStatements;
}

export namespace IncomeStatements {
  export type AsObject = {
    incomestatementsList: Array<IncomeStatement.AsObject>,
  }
}

export class IncomeStatement extends jspb.Message {
  getReportdate(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setReportdate(value?: google_protobuf_timestamp_pb.Timestamp): IncomeStatement;
  hasReportdate(): boolean;
  clearReportdate(): IncomeStatement;

  getFiscaldate(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setFiscaldate(value?: google_protobuf_timestamp_pb.Timestamp): IncomeStatement;
  hasFiscaldate(): boolean;
  clearFiscaldate(): IncomeStatement;

  getCurrency(): string;
  setCurrency(value: string): IncomeStatement;

  getTotalrevenue(): number;
  setTotalrevenue(value: number): IncomeStatement;

  getCostofrevenue(): number;
  setCostofrevenue(value: number): IncomeStatement;

  getGrossprofit(): number;
  setGrossprofit(value: number): IncomeStatement;

  getResearchanddevelopment(): number;
  setResearchanddevelopment(value: number): IncomeStatement;

  getSellinggeneralandadmin(): number;
  setSellinggeneralandadmin(value: number): IncomeStatement;

  getOperatingexpense(): number;
  setOperatingexpense(value: number): IncomeStatement;

  getOperatingincome(): number;
  setOperatingincome(value: number): IncomeStatement;

  getOtherincomeexpensenet(): number;
  setOtherincomeexpensenet(value: number): IncomeStatement;

  getEbit(): number;
  setEbit(value: number): IncomeStatement;

  getInterestincome(): number;
  setInterestincome(value: number): IncomeStatement;

  getPretaxincome(): number;
  setPretaxincome(value: number): IncomeStatement;

  getIncometax(): number;
  setIncometax(value: number): IncomeStatement;

  getMinorityinterest(): number;
  setMinorityinterest(value: number): IncomeStatement;

  getNetincome(): number;
  setNetincome(value: number): IncomeStatement;

  getNetincomebasic(): number;
  setNetincomebasic(value: number): IncomeStatement;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): IncomeStatement.AsObject;
  static toObject(includeInstance: boolean, msg: IncomeStatement): IncomeStatement.AsObject;
  static serializeBinaryToWriter(message: IncomeStatement, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): IncomeStatement;
  static deserializeBinaryFromReader(message: IncomeStatement, reader: jspb.BinaryReader): IncomeStatement;
}

export namespace IncomeStatement {
  export type AsObject = {
    reportdate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    fiscaldate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    currency: string,
    totalrevenue: number,
    costofrevenue: number,
    grossprofit: number,
    researchanddevelopment: number,
    sellinggeneralandadmin: number,
    operatingexpense: number,
    operatingincome: number,
    otherincomeexpensenet: number,
    ebit: number,
    interestincome: number,
    pretaxincome: number,
    incometax: number,
    minorityinterest: number,
    netincome: number,
    netincomebasic: number,
  }
}

export class CashFlows extends jspb.Message {
  getCashflowsList(): Array<CashFlow>;
  setCashflowsList(value: Array<CashFlow>): CashFlows;
  clearCashflowsList(): CashFlows;
  addCashflows(value?: CashFlow, index?: number): CashFlow;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CashFlows.AsObject;
  static toObject(includeInstance: boolean, msg: CashFlows): CashFlows.AsObject;
  static serializeBinaryToWriter(message: CashFlows, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CashFlows;
  static deserializeBinaryFromReader(message: CashFlows, reader: jspb.BinaryReader): CashFlows;
}

export namespace CashFlows {
  export type AsObject = {
    cashflowsList: Array<CashFlow.AsObject>,
  }
}

export class CashFlow extends jspb.Message {
  getReportdate(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setReportdate(value?: google_protobuf_timestamp_pb.Timestamp): CashFlow;
  hasReportdate(): boolean;
  clearReportdate(): CashFlow;

  getFiscaldate(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setFiscaldate(value?: google_protobuf_timestamp_pb.Timestamp): CashFlow;
  hasFiscaldate(): boolean;
  clearFiscaldate(): CashFlow;

  getCurrency(): string;
  setCurrency(value: string): CashFlow;

  getNetincome(): number;
  setNetincome(value: number): CashFlow;

  getDepreciation(): number;
  setDepreciation(value: number): CashFlow;

  getChangesinreceivables(): number;
  setChangesinreceivables(value: number): CashFlow;

  getChangesininventories(): number;
  setChangesininventories(value: number): CashFlow;

  getCashchange(): number;
  setCashchange(value: number): CashFlow;

  getCashflow(): number;
  setCashflow(value: number): CashFlow;

  getCapitalexpenditures(): number;
  setCapitalexpenditures(value: number): CashFlow;

  getInvestments(): number;
  setInvestments(value: number): CashFlow;

  getInvestingactivityother(): number;
  setInvestingactivityother(value: number): CashFlow;

  getTotalinvestingcashflows(): number;
  setTotalinvestingcashflows(value: number): CashFlow;

  getDividendspaid(): number;
  setDividendspaid(value: number): CashFlow;

  getNetborrowings(): number;
  setNetborrowings(value: number): CashFlow;

  getOtherfinancingcashflows(): number;
  setOtherfinancingcashflows(value: number): CashFlow;

  getCashflowfinancing(): number;
  setCashflowfinancing(value: number): CashFlow;

  getExchangerateeffect(): number;
  setExchangerateeffect(value: number): CashFlow;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CashFlow.AsObject;
  static toObject(includeInstance: boolean, msg: CashFlow): CashFlow.AsObject;
  static serializeBinaryToWriter(message: CashFlow, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CashFlow;
  static deserializeBinaryFromReader(message: CashFlow, reader: jspb.BinaryReader): CashFlow;
}

export namespace CashFlow {
  export type AsObject = {
    reportdate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    fiscaldate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    currency: string,
    netincome: number,
    depreciation: number,
    changesinreceivables: number,
    changesininventories: number,
    cashchange: number,
    cashflow: number,
    capitalexpenditures: number,
    investments: number,
    investingactivityother: number,
    totalinvestingcashflows: number,
    dividendspaid: number,
    netborrowings: number,
    otherfinancingcashflows: number,
    cashflowfinancing: number,
    exchangerateeffect: number,
  }
}

