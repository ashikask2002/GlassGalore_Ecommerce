wire:
	cd pkg/di && wire

run:
	go run ./cmd/api/main.go

swagger:
	swag init -g ./cmd/api/main.go