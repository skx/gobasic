

clean:
	go clean .
	find . -name 'c.out' -delete

test:
	go test -race  ./...
