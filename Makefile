.PHONY: run build compose
run:
	go run ./cmd/imob

build:
	docker build -t imob:local .

compose-up:
	docker-compose up --build
