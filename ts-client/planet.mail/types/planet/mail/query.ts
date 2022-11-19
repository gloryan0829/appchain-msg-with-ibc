/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../cosmos/base/query/v1beta1/pagination";
import { Message } from "./message";
import { Params } from "./params";
import { SentMessage } from "./sent_message";
import { TimedoutMessage } from "./timedout_message";

export const protobufPackage = "planet.mail";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetMessageRequest {
  id: number;
}

export interface QueryGetMessageResponse {
  Message: Message | undefined;
}

export interface QueryAllMessageRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllMessageResponse {
  Message: Message[];
  pagination: PageResponse | undefined;
}

export interface QueryGetSentMessageRequest {
  id: number;
}

export interface QueryGetSentMessageResponse {
  SentMessage: SentMessage | undefined;
}

export interface QueryAllSentMessageRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllSentMessageResponse {
  SentMessage: SentMessage[];
  pagination: PageResponse | undefined;
}

export interface QueryGetTimedoutMessageRequest {
  id: number;
}

export interface QueryGetTimedoutMessageResponse {
  TimedoutMessage: TimedoutMessage | undefined;
}

export interface QueryAllTimedoutMessageRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllTimedoutMessageResponse {
  TimedoutMessage: TimedoutMessage[];
  pagination: PageResponse | undefined;
}

function createBaseQueryParamsRequest(): QueryParamsRequest {
  return {};
}

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    return {};
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsRequest>, I>>(_: I): QueryParamsRequest {
    const message = createBaseQueryParamsRequest();
    return message;
  },
};

function createBaseQueryParamsResponse(): QueryParamsResponse {
  return { params: undefined };
}

export const QueryParamsResponse = {
  encode(message: QueryParamsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    return { params: isSet(object.params) ? Params.fromJSON(object.params) : undefined };
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsResponse>, I>>(object: I): QueryParamsResponse {
    const message = createBaseQueryParamsResponse();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    return message;
  },
};

function createBaseQueryGetMessageRequest(): QueryGetMessageRequest {
  return { id: 0 };
}

export const QueryGetMessageRequest = {
  encode(message: QueryGetMessageRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetMessageRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetMessageRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetMessageRequest {
    return { id: isSet(object.id) ? Number(object.id) : 0 };
  },

  toJSON(message: QueryGetMessageRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetMessageRequest>, I>>(object: I): QueryGetMessageRequest {
    const message = createBaseQueryGetMessageRequest();
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseQueryGetMessageResponse(): QueryGetMessageResponse {
  return { Message: undefined };
}

export const QueryGetMessageResponse = {
  encode(message: QueryGetMessageResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.Message !== undefined) {
      Message.encode(message.Message, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetMessageResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetMessageResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Message = Message.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetMessageResponse {
    return { Message: isSet(object.Message) ? Message.fromJSON(object.Message) : undefined };
  },

  toJSON(message: QueryGetMessageResponse): unknown {
    const obj: any = {};
    message.Message !== undefined && (obj.Message = message.Message ? Message.toJSON(message.Message) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetMessageResponse>, I>>(object: I): QueryGetMessageResponse {
    const message = createBaseQueryGetMessageResponse();
    message.Message = (object.Message !== undefined && object.Message !== null)
      ? Message.fromPartial(object.Message)
      : undefined;
    return message;
  },
};

function createBaseQueryAllMessageRequest(): QueryAllMessageRequest {
  return { pagination: undefined };
}

export const QueryAllMessageRequest = {
  encode(message: QueryAllMessageRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllMessageRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllMessageRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllMessageRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllMessageRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllMessageRequest>, I>>(object: I): QueryAllMessageRequest {
    const message = createBaseQueryAllMessageRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllMessageResponse(): QueryAllMessageResponse {
  return { Message: [], pagination: undefined };
}

export const QueryAllMessageResponse = {
  encode(message: QueryAllMessageResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.Message) {
      Message.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllMessageResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllMessageResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Message.push(Message.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllMessageResponse {
    return {
      Message: Array.isArray(object?.Message) ? object.Message.map((e: any) => Message.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllMessageResponse): unknown {
    const obj: any = {};
    if (message.Message) {
      obj.Message = message.Message.map((e) => e ? Message.toJSON(e) : undefined);
    } else {
      obj.Message = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllMessageResponse>, I>>(object: I): QueryAllMessageResponse {
    const message = createBaseQueryAllMessageResponse();
    message.Message = object.Message?.map((e) => Message.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetSentMessageRequest(): QueryGetSentMessageRequest {
  return { id: 0 };
}

export const QueryGetSentMessageRequest = {
  encode(message: QueryGetSentMessageRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetSentMessageRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetSentMessageRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetSentMessageRequest {
    return { id: isSet(object.id) ? Number(object.id) : 0 };
  },

  toJSON(message: QueryGetSentMessageRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetSentMessageRequest>, I>>(object: I): QueryGetSentMessageRequest {
    const message = createBaseQueryGetSentMessageRequest();
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseQueryGetSentMessageResponse(): QueryGetSentMessageResponse {
  return { SentMessage: undefined };
}

export const QueryGetSentMessageResponse = {
  encode(message: QueryGetSentMessageResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.SentMessage !== undefined) {
      SentMessage.encode(message.SentMessage, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetSentMessageResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetSentMessageResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.SentMessage = SentMessage.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetSentMessageResponse {
    return { SentMessage: isSet(object.SentMessage) ? SentMessage.fromJSON(object.SentMessage) : undefined };
  },

  toJSON(message: QueryGetSentMessageResponse): unknown {
    const obj: any = {};
    message.SentMessage !== undefined
      && (obj.SentMessage = message.SentMessage ? SentMessage.toJSON(message.SentMessage) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetSentMessageResponse>, I>>(object: I): QueryGetSentMessageResponse {
    const message = createBaseQueryGetSentMessageResponse();
    message.SentMessage = (object.SentMessage !== undefined && object.SentMessage !== null)
      ? SentMessage.fromPartial(object.SentMessage)
      : undefined;
    return message;
  },
};

function createBaseQueryAllSentMessageRequest(): QueryAllSentMessageRequest {
  return { pagination: undefined };
}

export const QueryAllSentMessageRequest = {
  encode(message: QueryAllSentMessageRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllSentMessageRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllSentMessageRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllSentMessageRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllSentMessageRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllSentMessageRequest>, I>>(object: I): QueryAllSentMessageRequest {
    const message = createBaseQueryAllSentMessageRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllSentMessageResponse(): QueryAllSentMessageResponse {
  return { SentMessage: [], pagination: undefined };
}

export const QueryAllSentMessageResponse = {
  encode(message: QueryAllSentMessageResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.SentMessage) {
      SentMessage.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllSentMessageResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllSentMessageResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.SentMessage.push(SentMessage.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllSentMessageResponse {
    return {
      SentMessage: Array.isArray(object?.SentMessage)
        ? object.SentMessage.map((e: any) => SentMessage.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllSentMessageResponse): unknown {
    const obj: any = {};
    if (message.SentMessage) {
      obj.SentMessage = message.SentMessage.map((e) => e ? SentMessage.toJSON(e) : undefined);
    } else {
      obj.SentMessage = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllSentMessageResponse>, I>>(object: I): QueryAllSentMessageResponse {
    const message = createBaseQueryAllSentMessageResponse();
    message.SentMessage = object.SentMessage?.map((e) => SentMessage.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetTimedoutMessageRequest(): QueryGetTimedoutMessageRequest {
  return { id: 0 };
}

export const QueryGetTimedoutMessageRequest = {
  encode(message: QueryGetTimedoutMessageRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetTimedoutMessageRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetTimedoutMessageRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetTimedoutMessageRequest {
    return { id: isSet(object.id) ? Number(object.id) : 0 };
  },

  toJSON(message: QueryGetTimedoutMessageRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetTimedoutMessageRequest>, I>>(
    object: I,
  ): QueryGetTimedoutMessageRequest {
    const message = createBaseQueryGetTimedoutMessageRequest();
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseQueryGetTimedoutMessageResponse(): QueryGetTimedoutMessageResponse {
  return { TimedoutMessage: undefined };
}

export const QueryGetTimedoutMessageResponse = {
  encode(message: QueryGetTimedoutMessageResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.TimedoutMessage !== undefined) {
      TimedoutMessage.encode(message.TimedoutMessage, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetTimedoutMessageResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetTimedoutMessageResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.TimedoutMessage = TimedoutMessage.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetTimedoutMessageResponse {
    return {
      TimedoutMessage: isSet(object.TimedoutMessage) ? TimedoutMessage.fromJSON(object.TimedoutMessage) : undefined,
    };
  },

  toJSON(message: QueryGetTimedoutMessageResponse): unknown {
    const obj: any = {};
    message.TimedoutMessage !== undefined
      && (obj.TimedoutMessage = message.TimedoutMessage ? TimedoutMessage.toJSON(message.TimedoutMessage) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetTimedoutMessageResponse>, I>>(
    object: I,
  ): QueryGetTimedoutMessageResponse {
    const message = createBaseQueryGetTimedoutMessageResponse();
    message.TimedoutMessage = (object.TimedoutMessage !== undefined && object.TimedoutMessage !== null)
      ? TimedoutMessage.fromPartial(object.TimedoutMessage)
      : undefined;
    return message;
  },
};

function createBaseQueryAllTimedoutMessageRequest(): QueryAllTimedoutMessageRequest {
  return { pagination: undefined };
}

export const QueryAllTimedoutMessageRequest = {
  encode(message: QueryAllTimedoutMessageRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllTimedoutMessageRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllTimedoutMessageRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllTimedoutMessageRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllTimedoutMessageRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllTimedoutMessageRequest>, I>>(
    object: I,
  ): QueryAllTimedoutMessageRequest {
    const message = createBaseQueryAllTimedoutMessageRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllTimedoutMessageResponse(): QueryAllTimedoutMessageResponse {
  return { TimedoutMessage: [], pagination: undefined };
}

export const QueryAllTimedoutMessageResponse = {
  encode(message: QueryAllTimedoutMessageResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.TimedoutMessage) {
      TimedoutMessage.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllTimedoutMessageResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllTimedoutMessageResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.TimedoutMessage.push(TimedoutMessage.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllTimedoutMessageResponse {
    return {
      TimedoutMessage: Array.isArray(object?.TimedoutMessage)
        ? object.TimedoutMessage.map((e: any) => TimedoutMessage.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllTimedoutMessageResponse): unknown {
    const obj: any = {};
    if (message.TimedoutMessage) {
      obj.TimedoutMessage = message.TimedoutMessage.map((e) => e ? TimedoutMessage.toJSON(e) : undefined);
    } else {
      obj.TimedoutMessage = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllTimedoutMessageResponse>, I>>(
    object: I,
  ): QueryAllTimedoutMessageResponse {
    const message = createBaseQueryAllTimedoutMessageResponse();
    message.TimedoutMessage = object.TimedoutMessage?.map((e) => TimedoutMessage.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a Message by id. */
  Message(request: QueryGetMessageRequest): Promise<QueryGetMessageResponse>;
  /** Queries a list of Message items. */
  MessageAll(request: QueryAllMessageRequest): Promise<QueryAllMessageResponse>;
  /** Queries a SentMessage by id. */
  SentMessage(request: QueryGetSentMessageRequest): Promise<QueryGetSentMessageResponse>;
  /** Queries a list of SentMessage items. */
  SentMessageAll(request: QueryAllSentMessageRequest): Promise<QueryAllSentMessageResponse>;
  /** Queries a TimedoutMessage by id. */
  TimedoutMessage(request: QueryGetTimedoutMessageRequest): Promise<QueryGetTimedoutMessageResponse>;
  /** Queries a list of TimedoutMessage items. */
  TimedoutMessageAll(request: QueryAllTimedoutMessageRequest): Promise<QueryAllTimedoutMessageResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.Message = this.Message.bind(this);
    this.MessageAll = this.MessageAll.bind(this);
    this.SentMessage = this.SentMessage.bind(this);
    this.SentMessageAll = this.SentMessageAll.bind(this);
    this.TimedoutMessage = this.TimedoutMessage.bind(this);
    this.TimedoutMessageAll = this.TimedoutMessageAll.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("planet.mail.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  Message(request: QueryGetMessageRequest): Promise<QueryGetMessageResponse> {
    const data = QueryGetMessageRequest.encode(request).finish();
    const promise = this.rpc.request("planet.mail.Query", "Message", data);
    return promise.then((data) => QueryGetMessageResponse.decode(new _m0.Reader(data)));
  }

  MessageAll(request: QueryAllMessageRequest): Promise<QueryAllMessageResponse> {
    const data = QueryAllMessageRequest.encode(request).finish();
    const promise = this.rpc.request("planet.mail.Query", "MessageAll", data);
    return promise.then((data) => QueryAllMessageResponse.decode(new _m0.Reader(data)));
  }

  SentMessage(request: QueryGetSentMessageRequest): Promise<QueryGetSentMessageResponse> {
    const data = QueryGetSentMessageRequest.encode(request).finish();
    const promise = this.rpc.request("planet.mail.Query", "SentMessage", data);
    return promise.then((data) => QueryGetSentMessageResponse.decode(new _m0.Reader(data)));
  }

  SentMessageAll(request: QueryAllSentMessageRequest): Promise<QueryAllSentMessageResponse> {
    const data = QueryAllSentMessageRequest.encode(request).finish();
    const promise = this.rpc.request("planet.mail.Query", "SentMessageAll", data);
    return promise.then((data) => QueryAllSentMessageResponse.decode(new _m0.Reader(data)));
  }

  TimedoutMessage(request: QueryGetTimedoutMessageRequest): Promise<QueryGetTimedoutMessageResponse> {
    const data = QueryGetTimedoutMessageRequest.encode(request).finish();
    const promise = this.rpc.request("planet.mail.Query", "TimedoutMessage", data);
    return promise.then((data) => QueryGetTimedoutMessageResponse.decode(new _m0.Reader(data)));
  }

  TimedoutMessageAll(request: QueryAllTimedoutMessageRequest): Promise<QueryAllTimedoutMessageResponse> {
    const data = QueryAllTimedoutMessageRequest.encode(request).finish();
    const promise = this.rpc.request("planet.mail.Query", "TimedoutMessageAll", data);
    return promise.then((data) => QueryAllTimedoutMessageResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
