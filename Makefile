wire:
	cd pkg/di && wire

run:
	go run ./cmd/api/main.go

swag :
	swag init -g cmd/api/main.go -o ./cmd/api/docs