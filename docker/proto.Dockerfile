FROM --platform=linux/amd64 golang:1.22.3

RUN apt-get update && \
    apt-get install unzip

RUN curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.14.0/protoc-3.14.0-linux-x86_64.zip && \
    unzip -o protoc-3.14.0-linux-x86_64.zip -d /usr/local bin/protoc && \
    unzip -o protoc-3.14.0-linux-x86_64.zip -d /usr/local include/* && \
    rm -rf protoc-3.13.0-linux-x86_64.zip

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1 && \
    go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@v1.5.1 && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

#CMD sleep 100000
CMD ["protoc", "-I=/go/src/", "--proto_path=/in/", "--go_out=/out/", "/in/event.proto"]
