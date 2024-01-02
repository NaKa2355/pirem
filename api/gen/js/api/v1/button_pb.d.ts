import * as jspb from 'google-protobuf'

import * as api_v1_irdata_pb from '../../api/v1/irdata_pb';


export class Button extends jspb.Message {
  getId(): string;
  setId(value: string): Button;

  getName(): string;
  setName(value: string): Button;

  getTag(): string;
  setTag(value: string): Button;

  getHasIrData(): boolean;
  setHasIrData(value: boolean): Button;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Button.AsObject;
  static toObject(includeInstance: boolean, msg: Button): Button.AsObject;
  static serializeBinaryToWriter(message: Button, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Button;
  static deserializeBinaryFromReader(message: Button, reader: jspb.BinaryReader): Button;
}

export namespace Button {
  export type AsObject = {
    id: string,
    name: string,
    tag: string,
    hasIrData: boolean,
  }
}

