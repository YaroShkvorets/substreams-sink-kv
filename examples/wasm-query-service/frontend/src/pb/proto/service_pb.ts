// @generated by protoc-gen-es v1.2.0 with parameter "target=ts"
// @generated from file proto/service.proto (package eth.service.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64 } from "@bufbuild/protobuf";

/**
 * @generated from message eth.service.v1.GetBlockInfoRequest
 */
export class GetBlockInfoRequest extends Message<GetBlockInfoRequest> {
  /**
   * @generated from field: string start = 1;
   */
  start = "";

  /**
   * @generated from field: string end = 2;
   */
  end = "";

  constructor(data?: PartialMessage<GetBlockInfoRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "eth.service.v1.GetBlockInfoRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "start", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "end", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetBlockInfoRequest {
    return new GetBlockInfoRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetBlockInfoRequest {
    return new GetBlockInfoRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetBlockInfoRequest {
    return new GetBlockInfoRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetBlockInfoRequest | PlainMessage<GetBlockInfoRequest> | undefined, b: GetBlockInfoRequest | PlainMessage<GetBlockInfoRequest> | undefined): boolean {
    return proto3.util.equals(GetBlockInfoRequest, a, b);
  }
}

/**
 * @generated from message eth.service.v1.GetMonthRequest
 */
export class GetMonthRequest extends Message<GetMonthRequest> {
  /**
   * @generated from field: string year = 1;
   */
  year = "";

  /**
   * @generated from field: string month = 2;
   */
  month = "";

  constructor(data?: PartialMessage<GetMonthRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "eth.service.v1.GetMonthRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "year", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "month", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetMonthRequest {
    return new GetMonthRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetMonthRequest {
    return new GetMonthRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetMonthRequest {
    return new GetMonthRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetMonthRequest | PlainMessage<GetMonthRequest> | undefined, b: GetMonthRequest | PlainMessage<GetMonthRequest> | undefined): boolean {
    return proto3.util.equals(GetMonthRequest, a, b);
  }
}

/**
 * @generated from message eth.service.v1.GetYearRequest
 */
export class GetYearRequest extends Message<GetYearRequest> {
  /**
   * @generated from field: string year = 1;
   */
  year = "";

  constructor(data?: PartialMessage<GetYearRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "eth.service.v1.GetYearRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "year", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetYearRequest {
    return new GetYearRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetYearRequest {
    return new GetYearRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetYearRequest {
    return new GetYearRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetYearRequest | PlainMessage<GetYearRequest> | undefined, b: GetYearRequest | PlainMessage<GetYearRequest> | undefined): boolean {
    return proto3.util.equals(GetYearRequest, a, b);
  }
}

/**
 * @generated from message eth.service.v1.Months
 */
export class Months extends Message<Months> {
  /**
   * @generated from field: repeated eth.service.v1.Month months = 1;
   */
  months: Month[] = [];

  constructor(data?: PartialMessage<Months>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "eth.service.v1.Months";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "months", kind: "message", T: Month, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Months {
    return new Months().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Months {
    return new Months().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Months {
    return new Months().fromJsonString(jsonString, options);
  }

  static equals(a: Months | PlainMessage<Months> | undefined, b: Months | PlainMessage<Months> | undefined): boolean {
    return proto3.util.equals(Months, a, b);
  }
}

/**
 * @generated from message eth.service.v1.Month
 */
export class Month extends Message<Month> {
  /**
   * @generated from field: string year = 1;
   */
  year = "";

  /**
   * @generated from field: string month = 2;
   */
  month = "";

  /**
   * @generated from field: eth.service.v1.Block first_block = 3;
   */
  firstBlock?: Block;

  /**
   * @generated from field: eth.service.v1.Block last_block = 4;
   */
  lastBlock?: Block;

  constructor(data?: PartialMessage<Month>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "eth.service.v1.Month";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "year", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "month", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "first_block", kind: "message", T: Block },
    { no: 4, name: "last_block", kind: "message", T: Block },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Month {
    return new Month().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Month {
    return new Month().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Month {
    return new Month().fromJsonString(jsonString, options);
  }

  static equals(a: Month | PlainMessage<Month> | undefined, b: Month | PlainMessage<Month> | undefined): boolean {
    return proto3.util.equals(Month, a, b);
  }
}

/**
 * @generated from message eth.service.v1.Block
 */
export class Block extends Message<Block> {
  /**
   * @generated from field: uint64 number = 1;
   */
  number = protoInt64.zero;

  /**
   * @generated from field: bytes hash = 2;
   */
  hash = new Uint8Array(0);

  /**
   * @generated from field: bytes parent_hash = 3;
   */
  parentHash = new Uint8Array(0);

  /**
   * @generated from field: string timestamp = 4;
   */
  timestamp = "";

  constructor(data?: PartialMessage<Block>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "eth.service.v1.Block";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "number", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 2, name: "hash", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
    { no: 3, name: "parent_hash", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
    { no: 4, name: "timestamp", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Block {
    return new Block().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Block {
    return new Block().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Block {
    return new Block().fromJsonString(jsonString, options);
  }

  static equals(a: Block | PlainMessage<Block> | undefined, b: Block | PlainMessage<Block> | undefined): boolean {
    return proto3.util.equals(Block, a, b);
  }
}

/**
 * @generated from message eth.service.v1.Error
 */
export class Error extends Message<Error> {
  /**
   * @generated from field: uint64 code = 1;
   */
  code = protoInt64.zero;

  /**
   * @generated from field: string message = 2;
   */
  message = "";

  constructor(data?: PartialMessage<Error>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "eth.service.v1.Error";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "code", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 2, name: "message", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Error {
    return new Error().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Error {
    return new Error().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Error {
    return new Error().fromJsonString(jsonString, options);
  }

  static equals(a: Error | PlainMessage<Error> | undefined, b: Error | PlainMessage<Error> | undefined): boolean {
    return proto3.util.equals(Error, a, b);
  }
}

