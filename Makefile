test:
	go test ./...

cover:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

cover-total:
	go test -coverprofile=coverage.tmp ./...
	@grep -vE 'mock|stub|dummy|fakes|mocks|testdata|report' coverage.tmp > coverage.out
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out
	rm -f coverage.tmp