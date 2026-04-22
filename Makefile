dev:
	air

build:
	go build -o ./build/main cmd/main.go

run-containers:
	docker compose up --build