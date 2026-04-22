dev:
	air

build-app:
	go build -o ./build/app cmd/main.go

run-containers:
	docker compose up --build

stop-containers:
	docker compose down -v