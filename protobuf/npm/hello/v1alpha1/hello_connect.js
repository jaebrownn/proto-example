// @generated by protoc-gen-connect-es v0.12.0
// @generated from file hello/v1alpha1/hello.proto (package hello.v1alpha1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { SayHelloRequest, SayHelloResponse } from "./hello_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * HelloService is a simple service consisting of methods to return greetings
 * to the callers of the service.
 *
 * @generated from service hello.v1alpha1.HelloService
 */
export const HelloService = {
  typeName: "hello.v1alpha1.HelloService",
  methods: {
    /**
     * SayHello returns a greeting message to the caller of the method.
     *
     * @generated from rpc hello.v1alpha1.HelloService.SayHello
     */
    sayHello: {
      name: "SayHello",
      I: SayHelloRequest,
      O: SayHelloResponse,
      kind: MethodKind.Unary,
    },
  }
};
