FROM node:lts-alpine AS client-build
RUN apk add --update --no-cache openjdk8-jre-base
WORKDIR /github.com/FujishigeTemma/Emoine/client
COPY ./client/package*.json ./
COPY ./client/patches ./patches
COPY ./client/scripts ./scripts
COPY ./docs ../docs
RUN npm ci --unsafe-perm
COPY ./client .
RUN npm run build

FROM golang:1.14-alpine AS server-build
RUN apk add --update --no-cache git curl make

RUN curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.12.3/protoc-3.12.3-linux-x86_64.zip
RUN mkdir -p protoc3 && \
  unzip protoc-3.12.3-linux-x86_64.zip -d protoc3
RUN ls -la /go/protoc3/bin
RUN ls -la /go/protoc3/bin
ENV PATH $PATH:/go/protoc3/bin
# protoの実行できない

WORKDIR /go/src/github.com/FujishigeTemma/Emoine
COPY ./go.* ./
RUN go mod download
RUN go install github.com/golang/protobuf/protoc-gen-go

COPY ./Makefile ./
RUN make proto

COPY . .
RUN go build

FROM alpine:3.12.0
WORKDIR /app
RUN apk --update add tzdata ca-certificates && \
  cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
  rm -rf /var/cache/apk/*
COPY --from=server-build /go/src/github.com/FujishigeTemma/Emoine ./
COPY --from=client-build /github.com/FujishigeTemma/Emoine/client/dist ./client/dist
ENTRYPOINT ./Emoine
