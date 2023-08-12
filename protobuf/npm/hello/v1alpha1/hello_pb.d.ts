// @generated by protoc-gen-es v1.3.0
// @generated from file hello/v1alpha1/hello.proto (package hello.v1alpha1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * Request message for the SayHello method.
 *
 * @generated from message hello.v1alpha1.SayHelloRequest
 */
export declare class SayHelloRequest extends Message<SayHelloRequest> {
  /**
   * The name of the person to greet.
   *
   * @generated from field: string name = 1;
   */
  name: string;

  constructor(data?: PartialMessage<SayHelloRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "hello.v1alpha1.SayHelloRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SayHelloRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SayHelloRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SayHelloRequest;

  static equals(a: SayHelloRequest | PlainMessage<SayHelloRequest> | undefined, b: SayHelloRequest | PlainMessage<SayHelloRequest> | undefined): boolean;
}

/**
 * Response message for the SayHello method.
 *
 * @generated from message hello.v1alpha1.SayHelloResponse
 */
export declare class SayHelloResponse extends Message<SayHelloResponse> {
  /**
   * The greeting.
   *
   * @generated from field: string greeting = 1;
   */
  greeting: string;

  constructor(data?: PartialMessage<SayHelloResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "hello.v1alpha1.SayHelloResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SayHelloResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SayHelloResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SayHelloResponse;

  static equals(a: SayHelloResponse | PlainMessage<SayHelloResponse> | undefined, b: SayHelloResponse | PlainMessage<SayHelloResponse> | undefined): boolean;
}

