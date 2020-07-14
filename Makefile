PROTO_TARGETS = comment reaction state

.PHONY: proto
proto:
	for t in $(PROTO_TARGETS); do \
		echo "protobuf build ... $$t"; \
		protoc -I=./docs --go_out=./handler ./docs/$$t.proto; \
	done