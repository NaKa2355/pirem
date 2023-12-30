import * as jspb from 'google-protobuf'



export class Remote extends jspb.Message {
  getId(): string;
  setId(value: string): Remote;

  getName(): string;
  setName(value: string): Remote;

  getDeviceId(): string;
  setDeviceId(value: string): Remote;

  getTag(): string;
  setTag(value: string): Remote;

  getButtonsList(): Array<Remote.Button>;
  setButtonsList(value: Array<Remote.Button>): Remote;
  clearButtonsList(): Remote;
  addButtons(value?: Remote.Button, index?: number): Remote.Button;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Remote.AsObject;
  static toObject(includeInstance: boolean, msg: Remote): Remote.AsObject;
  static serializeBinaryToWriter(message: Remote, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Remote;
  static deserializeBinaryFromReader(message: Remote, reader: jspb.BinaryReader): Remote;
}

export namespace Remote {
  export type AsObject = {
    id: string,
    name: string,
    deviceId: string,
    tag: string,
    buttonsList: Array<Remote.Button.AsObject>,
  }

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

}

