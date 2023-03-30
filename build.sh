protoc --go_out=proto --go_opt=module=github.com/MathPoem/mineProto/protoSamples \
    --go-grpc_out=require_unimplemented_servers=false:proto --go-grpc_opt=module=github.com/MathPoem/mineProto/protoSamples \
    proto/hello.proto