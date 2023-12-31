// @generated by protoc-gen-connect-es v0.12.0
// @generated from file hello/v1alpha1/hello.proto (package hello.v1alpha1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { SayGoodbyeRequest, SayGoodbyeResponse, SayHelloRequest, SayHelloResponse } from "./hello_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * HelloService is a simple service consisting of methods to return greetings
 * to the callers of the service.
 *
 * @generated from service hello.v1alpha1.HelloService
 */
export declare const HelloService: {
  readonly typeName: "hello.v1alpha1.HelloService",
  readonly methods: {
    /**
     * SayHello returns a greeting message to the caller of the method.
     *
     * @generated from rpc hello.v1alpha1.HelloService.SayHello
     */
    readonly sayHello: {
      readonly name: "SayHello",
      readonly I: typeof SayHelloRequest,
      readonly O: typeof SayHelloResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * SayGoobye returns a goodbye message to the caller of the method.
     *
     * @generated from rpc hello.v1alpha1.HelloService.SayGoodbye
     */
    readonly sayGoodbye: {
      readonly name: "SayGoodbye",
      readonly I: typeof SayGoodbyeRequest,
      readonly O: typeof SayGoodbyeResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};

