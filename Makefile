.PHONY: proto
proto:
	@protoc -I=./docs --go_out=. ./docs/*.proto

.PHONY: build
build:
	docker build -t emoine .


.PHONY: up
up:
	@docker-compose up -d --build

.PHONY: down
down:
	@docker-compose down -v
