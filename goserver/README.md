# goserver

A simple service which executes BASIC, in your browser.


## Implementation

The HTTP server binds to localhost:8080, and handles two events:

* `GET /`
  * Serves a single [index.html](data/index.html) file, containing javascript magic.
* `POST /`
  * Reads the contents of the HTTP POST and executes the BASIC code stored in the `code` parameter.

All other requests will result in a 404 error-code.



## Updating `data/index.html`

Because we want to ship a single binary we embed the contents of `data/index.html` inside our binary - meaning that if you wish to make changes to the content you need to do a little extra work.

The contents of `data/index.html` are stored in a compiled form inside the file `static.go`.  If you make changes to the source file you need to rebuild it.

Install `implant`, if you don't already have it:

    go get -u github.com/skx/implant

Then run this to update the `static.go` file:

    implant -output static.go data/

Finally you can rebuild the binary:

    go build .
