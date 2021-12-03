/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
import { Relayer, Threshold } from '../relayers/relayer';
export const protobufPackage = 'stafiprotocol.stafihub.relayers';
const baseQueryIsRelayerRequest = { denom: '', address: '' };
export const QueryIsRelayerRequest = {
    encode(message, writer = Writer.create()) {
        if (message.denom !== '') {
            writer.uint32(10).string(message.denom);
        }
        if (message.address !== '') {
            writer.uint32(18).string(message.address);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryIsRelayerRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.denom = reader.string();
                    break;
                case 2:
                    message.address = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryIsRelayerRequest };
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = String(object.denom);
        }
        else {
            message.denom = '';
        }
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.denom !== undefined && (obj.denom = message.denom);
        message.address !== undefined && (obj.address = message.address);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryIsRelayerRequest };
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = object.denom;
        }
        else {
            message.denom = '';
        }
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = '';
        }
        return message;
    }
};
const baseQueryIsRelayerResponse = { flag: false };
export const QueryIsRelayerResponse = {
    encode(message, writer = Writer.create()) {
        if (message.flag === true) {
            writer.uint32(8).bool(message.flag);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryIsRelayerResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.flag = reader.bool();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryIsRelayerResponse };
        if (object.flag !== undefined && object.flag !== null) {
            message.flag = Boolean(object.flag);
        }
        else {
            message.flag = false;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.flag !== undefined && (obj.flag = message.flag);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryIsRelayerResponse };
        if (object.flag !== undefined && object.flag !== null) {
            message.flag = object.flag;
        }
        else {
            message.flag = false;
        }
        return message;
    }
};
const baseQueryAllRelayerRequest = {};
export const QueryAllRelayerRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllRelayerRequest };
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
    fromJSON(object) {
        const message = { ...baseQueryAllRelayerRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllRelayerRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllRelayerResponse = {};
export const QueryAllRelayerResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.relayers) {
            Relayer.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllRelayerResponse };
        message.relayers = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.relayers.push(Relayer.decode(reader, reader.uint32()));
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
    fromJSON(object) {
        const message = { ...baseQueryAllRelayerResponse };
        message.relayers = [];
        if (object.relayers !== undefined && object.relayers !== null) {
            for (const e of object.relayers) {
                message.relayers.push(Relayer.fromJSON(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.relayers) {
            obj.relayers = message.relayers.map((e) => (e ? Relayer.toJSON(e) : undefined));
        }
        else {
            obj.relayers = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllRelayerResponse };
        message.relayers = [];
        if (object.relayers !== undefined && object.relayers !== null) {
            for (const e of object.relayers) {
                message.relayers.push(Relayer.fromPartial(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryRelayersByDenomRequest = { denom: '' };
export const QueryRelayersByDenomRequest = {
    encode(message, writer = Writer.create()) {
        if (message.denom !== '') {
            writer.uint32(10).string(message.denom);
        }
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryRelayersByDenomRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.denom = reader.string();
                    break;
                case 2:
                    message.pagination = PageRequest.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryRelayersByDenomRequest };
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = String(object.denom);
        }
        else {
            message.denom = '';
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.denom !== undefined && (obj.denom = message.denom);
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryRelayersByDenomRequest };
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = object.denom;
        }
        else {
            message.denom = '';
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryRelayersByDenomResponse = {};
export const QueryRelayersByDenomResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.relayers) {
            Relayer.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryRelayersByDenomResponse };
        message.relayers = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.relayers.push(Relayer.decode(reader, reader.uint32()));
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
    fromJSON(object) {
        const message = { ...baseQueryRelayersByDenomResponse };
        message.relayers = [];
        if (object.relayers !== undefined && object.relayers !== null) {
            for (const e of object.relayers) {
                message.relayers.push(Relayer.fromJSON(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.relayers) {
            obj.relayers = message.relayers.map((e) => (e ? Relayer.toJSON(e) : undefined));
        }
        else {
            obj.relayers = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryRelayersByDenomResponse };
        message.relayers = [];
        if (object.relayers !== undefined && object.relayers !== null) {
            for (const e of object.relayers) {
                message.relayers.push(Relayer.fromPartial(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryThresholdRequest = { denom: '' };
export const QueryThresholdRequest = {
    encode(message, writer = Writer.create()) {
        if (message.denom !== '') {
            writer.uint32(10).string(message.denom);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryThresholdRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.denom = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryThresholdRequest };
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = String(object.denom);
        }
        else {
            message.denom = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.denom !== undefined && (obj.denom = message.denom);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryThresholdRequest };
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = object.denom;
        }
        else {
            message.denom = '';
        }
        return message;
    }
};
const baseQueryThresholdResponse = {};
export const QueryThresholdResponse = {
    encode(message, writer = Writer.create()) {
        if (message.threshold !== undefined) {
            Threshold.encode(message.threshold, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryThresholdResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.threshold = Threshold.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryThresholdResponse };
        if (object.threshold !== undefined && object.threshold !== null) {
            message.threshold = Threshold.fromJSON(object.threshold);
        }
        else {
            message.threshold = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.threshold !== undefined && (obj.threshold = message.threshold ? Threshold.toJSON(message.threshold) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryThresholdResponse };
        if (object.threshold !== undefined && object.threshold !== null) {
            message.threshold = Threshold.fromPartial(object.threshold);
        }
        else {
            message.threshold = undefined;
        }
        return message;
    }
};
const baseQueryAllThresholdRequest = {};
export const QueryAllThresholdRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllThresholdRequest };
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
    fromJSON(object) {
        const message = { ...baseQueryAllThresholdRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllThresholdRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllThresholdResponse = {};
export const QueryAllThresholdResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.threshold) {
            Threshold.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllThresholdResponse };
        message.threshold = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.threshold.push(Threshold.decode(reader, reader.uint32()));
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
    fromJSON(object) {
        const message = { ...baseQueryAllThresholdResponse };
        message.threshold = [];
        if (object.threshold !== undefined && object.threshold !== null) {
            for (const e of object.threshold) {
                message.threshold.push(Threshold.fromJSON(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.threshold) {
            obj.threshold = message.threshold.map((e) => (e ? Threshold.toJSON(e) : undefined));
        }
        else {
            obj.threshold = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllThresholdResponse };
        message.threshold = [];
        if (object.threshold !== undefined && object.threshold !== null) {
            for (const e of object.threshold) {
                message.threshold.push(Threshold.fromPartial(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
export class QueryClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    RelayerAll(request) {
        const data = QueryAllRelayerRequest.encode(request).finish();
        const promise = this.rpc.request('stafiprotocol.stafihub.relayers.Query', 'RelayerAll', data);
        return promise.then((data) => QueryAllRelayerResponse.decode(new Reader(data)));
    }
    Threshold(request) {
        const data = QueryThresholdRequest.encode(request).finish();
        const promise = this.rpc.request('stafiprotocol.stafihub.relayers.Query', 'Threshold', data);
        return promise.then((data) => QueryThresholdResponse.decode(new Reader(data)));
    }
    ThresholdAll(request) {
        const data = QueryAllThresholdRequest.encode(request).finish();
        const promise = this.rpc.request('stafiprotocol.stafihub.relayers.Query', 'ThresholdAll', data);
        return promise.then((data) => QueryAllThresholdResponse.decode(new Reader(data)));
    }
    IsRelayer(request) {
        const data = QueryIsRelayerRequest.encode(request).finish();
        const promise = this.rpc.request('stafiprotocol.stafihub.relayers.Query', 'IsRelayer', data);
        return promise.then((data) => QueryIsRelayerResponse.decode(new Reader(data)));
    }
    RelayersByDenom(request) {
        const data = QueryRelayersByDenomRequest.encode(request).finish();
        const promise = this.rpc.request('stafiprotocol.stafihub.relayers.Query', 'RelayersByDenom', data);
        return promise.then((data) => QueryRelayersByDenomResponse.decode(new Reader(data)));
    }
}