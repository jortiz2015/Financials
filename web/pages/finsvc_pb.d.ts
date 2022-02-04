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
