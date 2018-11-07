# Fuzz-Testing

If you don't have the appropriate tools installed you can fetch them via:

    $ go get github.com/dvyukov/go-fuzz/go-fuzz
    $ go get github.com/dvyukov/go-fuzz/go-fuzz-build

Now you can build the `eval` package with fuzzing enabled:

    $ go-fuzz-build github.com/skx/gobasic/eval

Create a location to hold the work, and give it copies of some sample-programs:

    $ mkdir -p workdir/corpus
    $ cp examples/*.bas workdir/corpus

Now you can actually launch the fuzzer - here I use `-procs 1` so that
my desktop system isn't complete overloaded:

    $ go-fuzz -procs 1 -bin=eval-fuzz.zip -workdir workdir/

Now take a look at `workdir/crashers` to see the findings.
