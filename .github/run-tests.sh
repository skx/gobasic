#!/bin/bash

# Install the tools we use to test our code-quality.
#
# Here we setup the tools to install only if the "CI" environmental variable
# is not empty.  This is because locally I have them installed.
#
# NOTE: Github Actions always set CI=true
#
if [ ! -z "${CI}" ] ; then
    go install golang.org/x/lint/golint@latest
    go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
    go install honnef.co/go/tools/cmd/staticcheck@latest
fi

# Run the static-check tool - we ignore errors in goserver/static.go
t=$(mktemp)
staticcheck -checks all ./...  | grep -v "is deprecated"> $t
if [ -s $t ]; then
    echo "Found errors via 'staticcheck'"
    cat $t
    rm $t
    exit 1
fi
rm $t

# At this point failures cause aborts
set -e

# Run the linter
echo "Launching linter .."
golint -set_exit_status ./...
echo "Completed linter .."

# Run the shadow-checker
echo "Launching shadowed-variable check .."
go vet -vettool=$(which shadow) ./...
echo "Completed shadowed-variable check .."

# Run golang tests
go test ./...
