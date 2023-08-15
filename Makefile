
## SETUP
setup:
	cp .env.example .env

## TEST
test:
	go test ./... -v

test.cover:
	go test ./... --cover


## RUN APP
run.api:
	go run cmd/main.go