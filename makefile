setup:
	go mod download

run_devtest:
	go run ./cmd/dev-testing/main.go

run_tests:
	go test ./... $(additionalParams)

run_test_coverage:
	make run_tests additionalParams=-coverprofile="test_coverage.out"
	go tool cover -html="test_coverage.out"

run_benchmarks:
	make run_tests additionalParams=-bench=.

swag_generate:
	go run github.com/swaggo/swag/cmd/swag init -o ./pkg/docs --parseDependency