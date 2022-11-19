/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Message } from "./message";
import { Params } from "./params";
import { SentMessage } from "./sent_message";
import { TimedoutMessage } from "./timedout_message";

export const protobufPackage = "planet.mail";

/** GenesisState defines the mail module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  portId: string;
  messageList: Message[];
  messageCount: number;
  sentMessageList: SentMessage[];
  sentMessageCount: number;
  timedoutMessageList: TimedoutMessage[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  timedoutMessageCount: number;
}

function createBaseGenesisState(): GenesisState {
  return {
    params: undefined,
    portId: "",
    messageList: [],
    messageCount: 0,
    sentMessageList: [],
    sentMessageCount: 0,
    timedoutMessageList: [],
    timedoutMessageCount: 0,
  };
}

export const GenesisState = {
  encode(message: GenesisState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    if (message.portId !== "") {
      writer.uint32(18).string(message.portId);
    }
    for (const v of message.messageList) {
      Message.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    if (message.messageCount !== 0) {
      writer.uint32(32).uint64(message.messageCount);
    }
    for (const v of message.sentMessageList) {
      SentMessage.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    if (message.sentMessageCount !== 0) {
      writer.uint32(48).uint64(message.sentMessageCount);
    }
    for (const v of message.timedoutMessageList) {
      TimedoutMessage.encode(v!, writer.uint32(58).fork()).ldelim();
    }
    if (message.timedoutMessageCount !== 0) {
      writer.uint32(64).uint64(message.timedoutMessageCount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenesisState();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.portId = reader.string();
          break;
        case 3:
          message.messageList.push(Message.decode(reader, reader.uint32()));
          break;
        case 4:
          message.messageCount = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.sentMessageList.push(SentMessage.decode(reader, reader.uint32()));
          break;
        case 6:
          message.sentMessageCount = longToNumber(reader.uint64() as Long);
          break;
        case 7:
          message.timedoutMessageList.push(TimedoutMessage.decode(reader, reader.uint32()));
          break;
        case 8:
          message.timedoutMessageCount = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    return {
      params: isSet(object.params) ? Params.fromJSON(object.params) : undefined,
      portId: isSet(object.portId) ? String(object.portId) : "",
      messageList: Array.isArray(object?.messageList) ? object.messageList.map((e: any) => Message.fromJSON(e)) : [],
      messageCount: isSet(object.messageCount) ? Number(object.messageCount) : 0,
      sentMessageList: Array.isArray(object?.sentMessageList)
        ? object.sentMessageList.map((e: any) => SentMessage.fromJSON(e))
        : [],
      sentMessageCount: isSet(object.sentMessageCount) ? Number(object.sentMessageCount) : 0,
      timedoutMessageList: Array.isArray(object?.timedoutMessageList)
        ? object.timedoutMessageList.map((e: any) => TimedoutMessage.fromJSON(e))
        : [],
      timedoutMessageCount: isSet(object.timedoutMessageCount) ? Number(object.timedoutMessageCount) : 0,
    };
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    message.portId !== undefined && (obj.portId = message.portId);
    if (message.messageList) {
      obj.messageList = message.messageList.map((e) => e ? Message.toJSON(e) : undefined);
    } else {
      obj.messageList = [];
    }
    message.messageCount !== undefined && (obj.messageCount = Math.round(message.messageCount));
    if (message.sentMessageList) {
      obj.sentMessageList = message.sentMessageList.map((e) => e ? SentMessage.toJSON(e) : undefined);
    } else {
      obj.sentMessageList = [];
    }
    message.sentMessageCount !== undefined && (obj.sentMessageCount = Math.round(message.sentMessageCount));
    if (message.timedoutMessageList) {
      obj.timedoutMessageList = message.timedoutMessageList.map((e) => e ? TimedoutMessage.toJSON(e) : undefined);
    } else {
      obj.timedoutMessageList = [];
    }
    message.timedoutMessageCount !== undefined && (obj.timedoutMessageCount = Math.round(message.timedoutMessageCount));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GenesisState>, I>>(object: I): GenesisState {
    const message = createBaseGenesisState();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    message.portId = object.portId ?? "";
    message.messageList = object.messageList?.map((e) => Message.fromPartial(e)) || [];
    message.messageCount = object.messageCount ?? 0;
    message.sentMessageList = object.sentMessageList?.map((e) => SentMessage.fromPartial(e)) || [];
    message.sentMessageCount = object.sentMessageCount ?? 0;
    message.timedoutMessageList = object.timedoutMessageList?.map((e) => TimedoutMessage.fromPartial(e)) || [];
    message.timedoutMessageCount = object.timedoutMessageCount ?? 0;
    return message;
  },
};

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
