Protocol Buffers for the masses!
================================

Protocol Buffers is a Google made encoding mechanism that offers increased speed over the wire and easy to define contracts. These allow the automation of several processes through code generation.

### For Go services
The included Makefile should have you covered:

make import_path=path/to/protos/folder service_proto=service-name

Then just ./service-name to run the rpc server


### Resources

- [Protocol Buffers Language Guide](https://developers.google.com/protocol-buffers/docs/proto3)
- [Language API Overview](https://developers.google.com/protocol-buffers/docs/reference/overview)
- [Protocol Buffers Compiler](https://github.com/google/protobuf)
- [GRPC Home](http://grpc.io)
- [All about Swagger](https://www.openapis.org/)
- [GRPC Gateway](https://github.com/grpc-ecosystem/grpc-gateway/)
