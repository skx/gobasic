

default:
	@echo "make clean    - Clean binaries"
	@echo "make coverage - Run the coverage report"
	@echo "make test     - Run the test cases"

clean:
	@rm -f gobasic gobasic-* goserver/goserver goserver/goserver-*
	@find . -name 'c.out' -delete

coverage:
	@go test -coverprofile=c.out ./... | grep coverage:
	@go tool cover -html=c.out

test:
	@go test -race  ./...
