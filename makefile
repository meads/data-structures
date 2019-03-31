
# run all tests
test:
	go test ./...

# profile test coverage in browser
test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

list-deps:
	go list -m all
