# Fuzz-Testing

The 1.18 release of the golang compiler/toolset has integrated support for
fuzz-testing.

Fuzz-testing is basically magical and involves generating new inputs "randomly"
and running test-cases with those inputs.

## Running

Assuming you have go 1.18 or higher you can run the fuzz-testing of the
`eval` package like so:

    $ go test -fuzztime=60s -parallel=1 -fuzz=FuzzEval -v
