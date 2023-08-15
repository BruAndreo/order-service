run.api:
	go run cmd/main.go

test:
	go test ./... -v

test.cover:
	go test ./... --cover