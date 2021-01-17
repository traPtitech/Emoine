FROM node:14-alpine AS client-build
RUN apk add --update --no-cache openjdk8-jre-base
WORKDIR /github.com/traPtitech/Emoine/client
COPY ./client/package*.json ./
COPY ./client/scripts ./scripts
COPY ./docs ../docs
RUN npm ci --unsafe-perm
COPY ./client .
RUN npm run build

FROM golang:1.15-alpine AS server-build
RUN apk add --update --no-cache git curl make protoc
WORKDIR /go/src/github.com/traPtitech/Emoine
COPY ./go.* ./
RUN go mod download
RUN go install github.com/golang/protobuf/protoc-gen-go
RUN mkdir -p ./router/handler
COPY . .
RUN make proto
RUN go build

FROM alpine:3.13.0
WORKDIR /app
RUN apk --update --no-cache add tzdata ca-certificates openssl && \
  cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
  rm -rf /var/cache/apk/*
ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
 && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
 && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz
COPY --from=client-build /github.com/traPtitech/Emoine/client/dist ./client/dist
COPY --from=server-build /go/src/github.com/traPtitech/Emoine ./
ENTRYPOINT ./Emoine
