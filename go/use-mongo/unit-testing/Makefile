.PHONY: test

test:
	go test -v ./math

test-coverage:
	go test -coverprofile=coverage.out ./math

show-coverage:
	go tool cover -html=coverage.out

benchmark:
	go test -bench=. ./math
