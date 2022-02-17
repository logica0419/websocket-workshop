.PHONY: init
init:
	go mod download && cd frontend && npm i

.PHONY: build
build:
	@cd frontend && npm run build

.PHONY: dev
dev:
	@cd frontend && npm run dev

.PHONY: run
run:
	@go run ./*.go
