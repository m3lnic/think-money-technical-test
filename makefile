run_tests:
	go test ./...

run_test_coverage:
	go test ./... -coverprofile="test_coverage.out"
	go tool cover -html="test_coverage.out"