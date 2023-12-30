import * as jspb from 'google-protobuf'



export class Device extends jspb.Message {
  getId(): string;
  setId(value: string): Device;

  getName(): string;
  setName(value: string): Device;

  getBufferSize(): number;
  setBufferSize(value: number): Device;

  getCanSend(): boolean;
  setCanSend(value: boolean): Device;

  getCanReceive(): boolean;
  setCanReceive(value: boolean): Device;

  getDriverVersion(): string;
  setDriverVersion(value: string): Device;

  getFirmwareVersion(): string;
  setFirmwareVersion(value: string): Device;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Device.AsObject;
  static toObject(includeInstance: boolean, msg: Device): Device.AsObject;
  static serializeBinaryToWriter(message: Device, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Device;
  static deserializeBinaryFromReader(message: Device, reader: jspb.BinaryReader): Device;
}

export namespace Device {
  export type AsObject = {
    id: string,
    name: string,
    bufferSize: number,
    canSend: boolean,
    canReceive: boolean,
    driverVersion: string,
    firmwareVersion: string,
  }
}

