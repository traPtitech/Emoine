FROM node:16-alpine@sha256:a1f9d027912b58a7c75be7716c97cfbc6d3099f3a97ed84aa490be9dee20e787 AS client-build
RUN apk add --update --no-cache openjdk8-jre-base
WORKDIR /github.com/traPtitech/Emoine/client
COPY ./client/package*.json ./
COPY ./client/scripts ./scripts
COPY ./docs ../docs
RUN npm ci --unsafe-perm
COPY ./client .
ARG HOST
RUN HOST=$HOST npm run build

FROM golang:1.17-alpine@sha256:99ddec1bbfd6d6bca3f9804c02363daee8c8524dae50df7942e8e60788fd17c9 AS server-build
RUN apk add --update --no-cache git curl make protoc
WORKDIR /go/src/github.com/traPtitech/Emoine
COPY ./go.* ./
RUN go mod download
RUN go install github.com/golang/protobuf/protoc-gen-go@latest
RUN mkdir -p ./router/handler
COPY . .
RUN make proto
RUN go build

FROM alpine:3.15.0@sha256:21a3deaa0d32a8057914f36584b5288d2e5ecc984380bc0118285c70fa8c9300
WORKDIR /app
RUN apk --update --no-cache add tzdata ca-certificates openssl && \
  cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
  rm -rf /var/cache/apk/*
ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
 && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
 && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz
COPY --from=client-build /github.com/traPtitech/Emoine/client/dist ./web/dist
COPY --from=server-build /go/src/github.com/traPtitech/Emoine ./
ENTRYPOINT ./Emoine
