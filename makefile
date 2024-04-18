setup:
	go mod download

run:
	go run $(path)

run_rest_api:
	make run path=./main.go

run_devtest:
	make run path=./cmd/dev-testing/main.go

run_tests:
	go test ./... $(additionalParams)

run_test_coverage:
	make run_tests additionalParams=-coverprofile="test_coverage.out"
	go tool cover -html="test_coverage.out"

run_benchmarks:
	make run_tests additionalParams=-bench=.

swag_generate:
	go run github.com/swaggo/swag/cmd/swag init -o ./pkg/docs --parseDependency