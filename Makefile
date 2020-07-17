PROTO_TARGETS = comment reaction state message

.PHONY: proto
proto:
	for t in $(PROTO_TARGETS); do \
		echo "protobuf build ... $$t"; \
		protoc -I=./docs --go_out=./router/handler ./docs/$$t.proto; \
	done

.PHONY: build
build:
	docker build -t emoine .