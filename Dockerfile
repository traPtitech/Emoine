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
RUN apk add --update --no-cache git curl make protoc

WORKDIR /go/src/github.com/FujishigeTemma/Emoine
COPY . .
RUN go mod download

RUN go install github.com/golang/protobuf/protoc-gen-go

RUN mkdir -p ./router/handler
RUN make proto
RUN go build

FROM alpine:3.12.0
WORKDIR /app
RUN apk --update add tzdata ca-certificates && \
  cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
  rm -rf /var/cache/apk/*
COPY --from=server-build /go/src/github.com/FujishigeTemma/Emoine ./
COPY --from=client-build /github.com/FujishigeTemma/Emoine/client/dist ./client/dist
ENTRYPOINT ./Emoine
