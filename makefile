setup:
	go mod download

run_tests:
	go test ./... $(additionalParams)

run_test_coverage:
	go test ./... -coverprofile="test_coverage.out"
	go tool cover -html="test_coverage.out"

run_benchmarks:
	make run_tests additionalParams=-bench=.