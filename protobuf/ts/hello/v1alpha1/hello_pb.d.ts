import * as jspb from 'google-protobuf'



export class SayHelloRequest extends jspb.Message {
  getName(): string;
  setName(value: string): SayHelloRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SayHelloRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SayHelloRequest): SayHelloRequest.AsObject;
  static serializeBinaryToWriter(message: SayHelloRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SayHelloRequest;
  static deserializeBinaryFromReader(message: SayHelloRequest, reader: jspb.BinaryReader): SayHelloRequest;
}

export namespace SayHelloRequest {
  export type AsObject = {
    name: string,
  }
}

export class SayHelloResponse extends jspb.Message {
  getGreeting(): string;
  setGreeting(value: string): SayHelloResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SayHelloResponse.AsObject;
  static toObject(includeInstance: boolean, msg: SayHelloResponse): SayHelloResponse.AsObject;
  static serializeBinaryToWriter(message: SayHelloResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SayHelloResponse;
  static deserializeBinaryFromReader(message: SayHelloResponse, reader: jspb.BinaryReader): SayHelloResponse;
}

export namespace SayHelloResponse {
  export type AsObject = {
    greeting: string,
  }
}

