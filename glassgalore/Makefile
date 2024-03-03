wire:
	cd pkg/di && wire

run:
	go run ./cmd/api/main.go

swag :
	swag init -g cmd/api/main.go -o ./cmd/api/docs

mock: ##make mockfile using mockgen
	mockgen -source pkg/repository/interfaces/user.go -destination pkg/repository/mock/user_mock.go - package mock
    mockgen -source pkg/usecase/interfaces/user.go -destination pkg/usecase/mock/user_mock.go - package mock
	mockgen -source pkg/helper/interfaces/helper.go - destination pkg/helper/mock/helper_mock.go - package mock
	