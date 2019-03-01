#!/bin/bash

# Install the lint-tool
go get -u golang.org/x/lint/golint

# Init the modules
go mod init

# At this point failures cause aborts
set -e

# Run the linter
echo "Running linter .."
golint -set_exit_status ./...
echo "Linter complete .."

# Run golang tests
go test ./...
