FROM golang:bookworm

RUN apt update && apt install -y unzip dos2unix \
 && wget https://github.com/protocolbuffers/protobuf/releases/download/v35.0/protoc-35.0-linux-x86_64.zip -O protoc.zip \
 && unzip protoc.zip -d /usr/local
