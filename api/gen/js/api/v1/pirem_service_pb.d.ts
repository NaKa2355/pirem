import * as jspb from 'google-protobuf'

import * as api_v1_device_pb from '../../api/v1/device_pb';
import * as api_v1_irdata_pb from '../../api/v1/irdata_pb';
import * as api_v1_remote_pb from '../../api/v1/remote_pb';
import * as api_v1_button_pb from '../../api/v1/button_pb';
import * as api_v1_empty_pb from '../../api/v1/empty_pb';


export class SendIrRequest extends jspb.Message {
  getDeviceId(): string;
  setDeviceId(value: string): SendIrRequest;

  getIrData(): api_v1_irdata_pb.IrData | undefined;
  setIrData(value?: api_v1_irdata_pb.IrData): SendIrRequest;
  hasIrData(): boolean;
  clearIrData(): SendIrRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SendIrRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SendIrRequest): SendIrRequest.AsObject;
  static serializeBinaryToWriter(message: SendIrRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SendIrRequest;
  static deserializeBinaryFromReader(message: SendIrRequest, reader: jspb.BinaryReader): SendIrRequest;
}

export namespace SendIrRequest {
  export type AsObject = {
    deviceId: string,
    irData?: api_v1_irdata_pb.IrData.AsObject,
  }
}

export class SendIrResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SendIrResponse.AsObject;
  static toObject(includeInstance: boolean, msg: SendIrResponse): SendIrResponse.AsObject;
  static serializeBinaryToWriter(message: SendIrResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SendIrResponse;
  static deserializeBinaryFromReader(message: SendIrResponse, reader: jspb.BinaryReader): SendIrResponse;
}

export namespace SendIrResponse {
  export type AsObject = {
  }
}

export class ReceiveIrRequest extends jspb.Message {
  getDeviceId(): string;
  setDeviceId(value: string): ReceiveIrRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ReceiveIrRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ReceiveIrRequest): ReceiveIrRequest.AsObject;
  static serializeBinaryToWriter(message: ReceiveIrRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ReceiveIrRequest;
  static deserializeBinaryFromReader(message: ReceiveIrRequest, reader: jspb.BinaryReader): ReceiveIrRequest;
}

export namespace ReceiveIrRequest {
  export type AsObject = {
    deviceId: string,
  }
}

export class ReceiveIrResponse extends jspb.Message {
  getIrData(): api_v1_irdata_pb.IrData | undefined;
  setIrData(value?: api_v1_irdata_pb.IrData): ReceiveIrResponse;
  hasIrData(): boolean;
  clearIrData(): ReceiveIrResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ReceiveIrResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ReceiveIrResponse): ReceiveIrResponse.AsObject;
  static serializeBinaryToWriter(message: ReceiveIrResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ReceiveIrResponse;
  static deserializeBinaryFromReader(message: ReceiveIrResponse, reader: jspb.BinaryReader): ReceiveIrResponse;
}

export namespace ReceiveIrResponse {
  export type AsObject = {
    irData?: api_v1_irdata_pb.IrData.AsObject,
  }
}

export class ListDevicesRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListDevicesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListDevicesRequest): ListDevicesRequest.AsObject;
  static serializeBinaryToWriter(message: ListDevicesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListDevicesRequest;
  static deserializeBinaryFromReader(message: ListDevicesRequest, reader: jspb.BinaryReader): ListDevicesRequest;
}

export namespace ListDevicesRequest {
  export type AsObject = {
  }
}

export class ListDevicesResponse extends jspb.Message {
  getDevicesList(): Array<api_v1_device_pb.Device>;
  setDevicesList(value: Array<api_v1_device_pb.Device>): ListDevicesResponse;
  clearDevicesList(): ListDevicesResponse;
  addDevices(value?: api_v1_device_pb.Device, index?: number): api_v1_device_pb.Device;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListDevicesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListDevicesResponse): ListDevicesResponse.AsObject;
  static serializeBinaryToWriter(message: ListDevicesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListDevicesResponse;
  static deserializeBinaryFromReader(message: ListDevicesResponse, reader: jspb.BinaryReader): ListDevicesResponse;
}

export namespace ListDevicesResponse {
  export type AsObject = {
    devicesList: Array<api_v1_device_pb.Device.AsObject>,
  }
}

export class GetDeviceRequest extends jspb.Message {
  getDeviceId(): string;
  setDeviceId(value: string): GetDeviceRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetDeviceRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetDeviceRequest): GetDeviceRequest.AsObject;
  static serializeBinaryToWriter(message: GetDeviceRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetDeviceRequest;
  static deserializeBinaryFromReader(message: GetDeviceRequest, reader: jspb.BinaryReader): GetDeviceRequest;
}

export namespace GetDeviceRequest {
  export type AsObject = {
    deviceId: string,
  }
}

export class ListRemotesRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListRemotesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListRemotesRequest): ListRemotesRequest.AsObject;
  static serializeBinaryToWriter(message: ListRemotesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListRemotesRequest;
  static deserializeBinaryFromReader(message: ListRemotesRequest, reader: jspb.BinaryReader): ListRemotesRequest;
}

export namespace ListRemotesRequest {
  export type AsObject = {
  }
}

export class ListRemotesResponse extends jspb.Message {
  getRemotesList(): Array<api_v1_remote_pb.Remote>;
  setRemotesList(value: Array<api_v1_remote_pb.Remote>): ListRemotesResponse;
  clearRemotesList(): ListRemotesResponse;
  addRemotes(value?: api_v1_remote_pb.Remote, index?: number): api_v1_remote_pb.Remote;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListRemotesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListRemotesResponse): ListRemotesResponse.AsObject;
  static serializeBinaryToWriter(message: ListRemotesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListRemotesResponse;
  static deserializeBinaryFromReader(message: ListRemotesResponse, reader: jspb.BinaryReader): ListRemotesResponse;
}

export namespace ListRemotesResponse {
  export type AsObject = {
    remotesList: Array<api_v1_remote_pb.Remote.AsObject>,
  }
}

export class GetRemoteRequest extends jspb.Message {
  getRemoteId(): string;
  setRemoteId(value: string): GetRemoteRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetRemoteRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetRemoteRequest): GetRemoteRequest.AsObject;
  static serializeBinaryToWriter(message: GetRemoteRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetRemoteRequest;
  static deserializeBinaryFromReader(message: GetRemoteRequest, reader: jspb.BinaryReader): GetRemoteRequest;
}

export namespace GetRemoteRequest {
  export type AsObject = {
    remoteId: string,
  }
}

export class UpdateRemoteRequest extends jspb.Message {
  getRemoteId(): string;
  setRemoteId(value: string): UpdateRemoteRequest;

  getName(): string;
  setName(value: string): UpdateRemoteRequest;

  getDeviceId(): string;
  setDeviceId(value: string): UpdateRemoteRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateRemoteRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateRemoteRequest): UpdateRemoteRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateRemoteRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateRemoteRequest;
  static deserializeBinaryFromReader(message: UpdateRemoteRequest, reader: jspb.BinaryReader): UpdateRemoteRequest;
}

export namespace UpdateRemoteRequest {
  export type AsObject = {
    remoteId: string,
    name: string,
    deviceId: string,
  }
}

export class DeleteRemoteRequest extends jspb.Message {
  getRemoteId(): string;
  setRemoteId(value: string): DeleteRemoteRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteRemoteRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteRemoteRequest): DeleteRemoteRequest.AsObject;
  static serializeBinaryToWriter(message: DeleteRemoteRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteRemoteRequest;
  static deserializeBinaryFromReader(message: DeleteRemoteRequest, reader: jspb.BinaryReader): DeleteRemoteRequest;
}

export namespace DeleteRemoteRequest {
  export type AsObject = {
    remoteId: string,
  }
}

export class GetButtonRequest extends jspb.Message {
  getButtonId(): string;
  setButtonId(value: string): GetButtonRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetButtonRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetButtonRequest): GetButtonRequest.AsObject;
  static serializeBinaryToWriter(message: GetButtonRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetButtonRequest;
  static deserializeBinaryFromReader(message: GetButtonRequest, reader: jspb.BinaryReader): GetButtonRequest;
}

export namespace GetButtonRequest {
  export type AsObject = {
    buttonId: string,
  }
}

export class LearnIrDataRequest extends jspb.Message {
  getButtonId(): string;
  setButtonId(value: string): LearnIrDataRequest;

  getIrData(): api_v1_irdata_pb.IrData | undefined;
  setIrData(value?: api_v1_irdata_pb.IrData): LearnIrDataRequest;
  hasIrData(): boolean;
  clearIrData(): LearnIrDataRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LearnIrDataRequest.AsObject;
  static toObject(includeInstance: boolean, msg: LearnIrDataRequest): LearnIrDataRequest.AsObject;
  static serializeBinaryToWriter(message: LearnIrDataRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LearnIrDataRequest;
  static deserializeBinaryFromReader(message: LearnIrDataRequest, reader: jspb.BinaryReader): LearnIrDataRequest;
}

export namespace LearnIrDataRequest {
  export type AsObject = {
    buttonId: string,
    irData?: api_v1_irdata_pb.IrData.AsObject,
  }
}

export class PushButtonRequest extends jspb.Message {
  getButtonId(): string;
  setButtonId(value: string): PushButtonRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PushButtonRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PushButtonRequest): PushButtonRequest.AsObject;
  static serializeBinaryToWriter(message: PushButtonRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PushButtonRequest;
  static deserializeBinaryFromReader(message: PushButtonRequest, reader: jspb.BinaryReader): PushButtonRequest;
}

export namespace PushButtonRequest {
  export type AsObject = {
    buttonId: string,
  }
}

