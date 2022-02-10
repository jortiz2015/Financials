/**
 * @fileoverview gRPC-Web generated client stub for 
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as finsvc_pb from './finsvc_pb';


export class FinSvcClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'binary';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoGetAnnualBalanceSheets = new grpcWeb.MethodDescriptor(
    '/FinSvc/GetAnnualBalanceSheets',
    grpcWeb.MethodType.UNARY,
    finsvc_pb.GetRequest,
    finsvc_pb.BalanceSheets,
    (request: finsvc_pb.GetRequest) => {
      return request.serializeBinary();
    },
    finsvc_pb.BalanceSheets.deserializeBinary
  );

  getAnnualBalanceSheets(
    request: finsvc_pb.GetRequest,
    metadata: grpcWeb.Metadata | null): Promise<finsvc_pb.BalanceSheets>;

  getAnnualBalanceSheets(
    request: finsvc_pb.GetRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: finsvc_pb.BalanceSheets) => void): grpcWeb.ClientReadableStream<finsvc_pb.BalanceSheets>;

  getAnnualBalanceSheets(
    request: finsvc_pb.GetRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: finsvc_pb.BalanceSheets) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/FinSvc/GetAnnualBalanceSheets',
        request,
        metadata || {},
        this.methodInfoGetAnnualBalanceSheets,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/FinSvc/GetAnnualBalanceSheets',
    request,
    metadata || {},
    this.methodInfoGetAnnualBalanceSheets);
  }

  methodInfoGetQuarterlyBalanceSheets = new grpcWeb.MethodDescriptor(
    '/FinSvc/GetQuarterlyBalanceSheets',
    grpcWeb.MethodType.UNARY,
    finsvc_pb.GetRequest,
    finsvc_pb.BalanceSheets,
    (request: finsvc_pb.GetRequest) => {
      return request.serializeBinary();
    },
    finsvc_pb.BalanceSheets.deserializeBinary
  );

  getQuarterlyBalanceSheets(
    request: finsvc_pb.GetRequest,
    metadata: grpcWeb.Metadata | null): Promise<finsvc_pb.BalanceSheets>;

  getQuarterlyBalanceSheets(
    request: finsvc_pb.GetRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: finsvc_pb.BalanceSheets) => void): grpcWeb.ClientReadableStream<finsvc_pb.BalanceSheets>;

  getQuarterlyBalanceSheets(
    request: finsvc_pb.GetRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: finsvc_pb.BalanceSheets) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/FinSvc/GetQuarterlyBalanceSheets',
        request,
        metadata || {},
        this.methodInfoGetQuarterlyBalanceSheets,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/FinSvc/GetQuarterlyBalanceSheets',
    request,
    metadata || {},
    this.methodInfoGetQuarterlyBalanceSheets);
  }

  methodInfoGetAnnualIncomeStatements = new grpcWeb.MethodDescriptor(
    '/FinSvc/GetAnnualIncomeStatements',
    grpcWeb.MethodType.UNARY,
    finsvc_pb.GetRequest,
    finsvc_pb.IncomeStatements,
    (request: finsvc_pb.GetRequest) => {
      return request.serializeBinary();
    },
    finsvc_pb.IncomeStatements.deserializeBinary
  );

  getAnnualIncomeStatements(
    request: finsvc_pb.GetRequest,
    metadata: grpcWeb.Metadata | null): Promise<finsvc_pb.IncomeStatements>;

  getAnnualIncomeStatements(
    request: finsvc_pb.GetRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: finsvc_pb.IncomeStatements) => void): grpcWeb.ClientReadableStream<finsvc_pb.IncomeStatements>;

  getAnnualIncomeStatements(
    request: finsvc_pb.GetRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: finsvc_pb.IncomeStatements) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/FinSvc/GetAnnualIncomeStatements',
        request,
        metadata || {},
        this.methodInfoGetAnnualIncomeStatements,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/FinSvc/GetAnnualIncomeStatements',
    request,
    metadata || {},
    this.methodInfoGetAnnualIncomeStatements);
  }

  methodInfoGetQuarterlyIncomeStatements = new grpcWeb.MethodDescriptor(
    '/FinSvc/GetQuarterlyIncomeStatements',
    grpcWeb.MethodType.UNARY,
    finsvc_pb.GetRequest,
    finsvc_pb.IncomeStatements,
    (request: finsvc_pb.GetRequest) => {
      return request.serializeBinary();
    },
    finsvc_pb.IncomeStatements.deserializeBinary
  );

  getQuarterlyIncomeStatements(
    request: finsvc_pb.GetRequest,
    metadata: grpcWeb.Metadata | null): Promise<finsvc_pb.IncomeStatements>;

  getQuarterlyIncomeStatements(
    request: finsvc_pb.GetRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: finsvc_pb.IncomeStatements) => void): grpcWeb.ClientReadableStream<finsvc_pb.IncomeStatements>;

  getQuarterlyIncomeStatements(
    request: finsvc_pb.GetRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: finsvc_pb.IncomeStatements) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/FinSvc/GetQuarterlyIncomeStatements',
        request,
        metadata || {},
        this.methodInfoGetQuarterlyIncomeStatements,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/FinSvc/GetQuarterlyIncomeStatements',
    request,
    metadata || {},
    this.methodInfoGetQuarterlyIncomeStatements);
  }

  methodInfoGetAnnualCashFlows = new grpcWeb.MethodDescriptor(
    '/FinSvc/GetAnnualCashFlows',
    grpcWeb.MethodType.UNARY,
    finsvc_pb.GetRequest,
    finsvc_pb.CashFlows,
    (request: finsvc_pb.GetRequest) => {
      return request.serializeBinary();
    },
    finsvc_pb.CashFlows.deserializeBinary
  );

  getAnnualCashFlows(
    request: finsvc_pb.GetRequest,
    metadata: grpcWeb.Metadata | null): Promise<finsvc_pb.CashFlows>;

  getAnnualCashFlows(
    request: finsvc_pb.GetRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: finsvc_pb.CashFlows) => void): grpcWeb.ClientReadableStream<finsvc_pb.CashFlows>;

  getAnnualCashFlows(
    request: finsvc_pb.GetRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: finsvc_pb.CashFlows) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/FinSvc/GetAnnualCashFlows',
        request,
        metadata || {},
        this.methodInfoGetAnnualCashFlows,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/FinSvc/GetAnnualCashFlows',
    request,
    metadata || {},
    this.methodInfoGetAnnualCashFlows);
  }

  methodInfoGetQuarterlyCashFlows = new grpcWeb.MethodDescriptor(
    '/FinSvc/GetQuarterlyCashFlows',
    grpcWeb.MethodType.UNARY,
    finsvc_pb.GetRequest,
    finsvc_pb.CashFlows,
    (request: finsvc_pb.GetRequest) => {
      return request.serializeBinary();
    },
    finsvc_pb.CashFlows.deserializeBinary
  );

  getQuarterlyCashFlows(
    request: finsvc_pb.GetRequest,
    metadata: grpcWeb.Metadata | null): Promise<finsvc_pb.CashFlows>;

  getQuarterlyCashFlows(
    request: finsvc_pb.GetRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: finsvc_pb.CashFlows) => void): grpcWeb.ClientReadableStream<finsvc_pb.CashFlows>;

  getQuarterlyCashFlows(
    request: finsvc_pb.GetRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: finsvc_pb.CashFlows) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/FinSvc/GetQuarterlyCashFlows',
        request,
        metadata || {},
        this.methodInfoGetQuarterlyCashFlows,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/FinSvc/GetQuarterlyCashFlows',
    request,
    metadata || {},
    this.methodInfoGetQuarterlyCashFlows);
  }

}

