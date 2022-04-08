# Fuzz-Testing

The 1.18 release of the golang compiler/toolset has integrated support for
fuzz-testing.

Fuzz-testing is basically magical and involves generating new inputs "randomly"
and running test-cases with those inputs.

## Running

Assuming you have go 1.18 or higher you can run the fuzz-testing of the
`eval` package like so:

    $ cd eval/
    $ go test -fuzztime=60s -parallel=1 -fuzz=FuzzEval -v

Output will look something like this:

    === FUZZ  FuzzEval
    fuzz: elapsed: 0s, gathering baseline coverage: 0/1084 completed
    fuzz: elapsed: 3s, gathering baseline coverage: 108/1084 completed
    fuzz: elapsed: 6s, gathering baseline coverage: 338/1084 completed
    fuzz: elapsed: 9s, gathering baseline coverage: 637/1084 completed
    fuzz: elapsed: 12s, gathering baseline coverage: 916/1084 completed
    fuzz: elapsed: 15s, gathering baseline coverage: 1067/1084 completed
    fuzz: elapsed: 15s, gathering baseline coverage: 1084/1084 completed, now fuzzing with 1 workers
    fuzz: elapsed: 18s, execs: 9791 (2908/sec), new interesting: 0 (total: 1084)
    fuzz: elapsed: 21s, execs: 31008 (7072/sec), new interesting: 1 (total: 1085)
    fuzz: elapsed: 24s, execs: 59590 (9529/sec), new interesting: 1 (total: 1085)

If the fuzzer terminates with a `panic` then you've found a new failure, and you should examine the contents of the file it generates and displays.  There _are_ some error-cases which are expected:

* If the fuzzer generates bogus BASIC
* If the fuzzer generates an infinite loop which is terminated via a timeout
* etc.

If you find a panic which looks like it is caused by bogus BASIC then update `fuzz_test.go` to add that to the "known failure" list.   Otherwise leave it running (perhaps overnight, removing the `-fuzztime=60s` parameter), and ideally it will keep going and not crash.
