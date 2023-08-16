// @generated by protoc-gen-es v1.3.0
// @generated from file hello/v1alpha1/hello.proto (package hello.v1alpha1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * Request message for the SayHello method.
 *
 * @generated from message hello.v1alpha1.SayHelloRequest
 */
export const SayHelloRequest = proto3.makeMessageType(
  "hello.v1alpha1.SayHelloRequest",
  () => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * Response message for the SayHello method.
 *
 * @generated from message hello.v1alpha1.SayHelloResponse
 */
export const SayHelloResponse = proto3.makeMessageType(
  "hello.v1alpha1.SayHelloResponse",
  () => [
    { no: 1, name: "greeting", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * Request message for the SayGoodbye method.
 *
 * @generated from message hello.v1alpha1.SayGoodbyeRequest
 */
export const SayGoodbyeRequest = proto3.makeMessageType(
  "hello.v1alpha1.SayGoodbyeRequest",
  () => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * Response message for the SayGoodbye method.
 *
 * @generated from message hello.v1alpha1.SayGoodbyeResponse
 */
export const SayGoodbyeResponse = proto3.makeMessageType(
  "hello.v1alpha1.SayGoodbyeResponse",
  () => [
    { no: 1, name: "goodbye", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

