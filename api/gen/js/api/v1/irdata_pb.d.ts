import * as jspb from 'google-protobuf'



export class RawIrData extends jspb.Message {
  getCarrierFreqKhz(): number;
  setCarrierFreqKhz(value: number): RawIrData;

  getOnOffPluseNsList(): Array<number>;
  setOnOffPluseNsList(value: Array<number>): RawIrData;
  clearOnOffPluseNsList(): RawIrData;
  addOnOffPluseNs(value: number, index?: number): RawIrData;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RawIrData.AsObject;
  static toObject(includeInstance: boolean, msg: RawIrData): RawIrData.AsObject;
  static serializeBinaryToWriter(message: RawIrData, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RawIrData;
  static deserializeBinaryFromReader(message: RawIrData, reader: jspb.BinaryReader): RawIrData;
}

export namespace RawIrData {
  export type AsObject = {
    carrierFreqKhz: number,
    onOffPluseNsList: Array<number>,
  }
}

export class None extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): None.AsObject;
  static toObject(includeInstance: boolean, msg: None): None.AsObject;
  static serializeBinaryToWriter(message: None, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): None;
  static deserializeBinaryFromReader(message: None, reader: jspb.BinaryReader): None;
}

export namespace None {
  export type AsObject = {
  }
}

export class IrData extends jspb.Message {
  getRaw(): RawIrData | undefined;
  setRaw(value?: RawIrData): IrData;
  hasRaw(): boolean;
  clearRaw(): IrData;

  getNone(): None | undefined;
  setNone(value?: None): IrData;
  hasNone(): boolean;
  clearNone(): IrData;

  getDataCase(): IrData.DataCase;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): IrData.AsObject;
  static toObject(includeInstance: boolean, msg: IrData): IrData.AsObject;
  static serializeBinaryToWriter(message: IrData, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): IrData;
  static deserializeBinaryFromReader(message: IrData, reader: jspb.BinaryReader): IrData;
}

export namespace IrData {
  export type AsObject = {
    raw?: RawIrData.AsObject,
    none?: None.AsObject,
  }

  export enum DataCase { 
    DATA_NOT_SET = 0,
    RAW = 1,
    NONE = 2,
  }
}

