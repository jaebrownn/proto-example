import * as grpcWeb from 'grpc-web';

import * as hello_v1alpha1_hello_pb from '../../hello/v1alpha1/hello_pb';


export class HelloServiceClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  sayHello(
    request: hello_v1alpha1_hello_pb.SayHelloRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: hello_v1alpha1_hello_pb.SayHelloResponse) => void
  ): grpcWeb.ClientReadableStream<hello_v1alpha1_hello_pb.SayHelloResponse>;

  sayGoodbye(
    request: hello_v1alpha1_hello_pb.SayGoodbyeRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: hello_v1alpha1_hello_pb.SayGoodbyeResponse) => void
  ): grpcWeb.ClientReadableStream<hello_v1alpha1_hello_pb.SayGoodbyeResponse>;

}

export class HelloServicePromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  sayHello(
    request: hello_v1alpha1_hello_pb.SayHelloRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<hello_v1alpha1_hello_pb.SayHelloResponse>;

  sayGoodbye(
    request: hello_v1alpha1_hello_pb.SayGoodbyeRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<hello_v1alpha1_hello_pb.SayGoodbyeResponse>;

}

