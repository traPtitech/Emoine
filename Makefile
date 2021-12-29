PROTO_TARGETS = comment reaction state viewer message

.PHONY: proto
proto:
	@rm -rf router/pb
	@mkdir -p router/pb
	@protoc -I ./docs/ --go_out=router/pb --go_opt=paths=source_relative ./docs/*.proto

.PHONY: build
build:
	docker build -t emoine .


.PHONY: up
up:
	@docker-compose up -d --build

.PHONY: down
down:
	@docker-compose down -v
