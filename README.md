# Newton's Wells in Go

This program generates nice images by applying Newton's polynomial
root finding technique. The roots of the polynomial is given at the
command line. Then the application tries to find a root by getting
every pixel as a guess and measuring how long it takes to reach the
root from that point. If a point cannot reach the root, it is colored
black. Otherwise, the color of the pixel is darkened as it takes more
steps to reach the root.

## Installation

Assuming that you have the latest Go toolchain installed just run:

> `go get git.fazlamesai.net/boran.puhaloglu/paralleling-pi/go/newtons-wells`

If you want to build from git then first, make sure that you have the
necessary libraries installed:

> `go get github.com/ogier/pflag`
> `go get code.google.com/p/sadbox/color`

Then in this directory, run: `go build`

## Usage

```
Usage: ./newtons-wells [options]:

  -c, --concurrency=4: Number of CPU cores to use. Defaults to all the cores available.
  -y, --height=480: Height of the resulting image
  -h, --help=false: Usage information
  -o, --output="-": Output file name. If omitted, standard output will be used.
  -r, --roots="1+1i": Complex roots, separated by commas. Ex: 3+2i,1+1i
  -x, --width=640: Width of the resulting image
```

Example command:

> `./newtons-wells -x 1920 -y 1080 -o example.png -r 960+540i,480+540i,1420+540i`

## TODO

- Increase test coverage
- Newton's Wells is a problem that is embarassingly parallel by it's
  nature. The application should take advantage of this fact and use a
  number of goroutines to parallelize this process.
- The process can also be distributed among multiple hosts using some
  kind of rpc (see [net/rpc](http://golang.org/pkg/net/rpc/) package)

## Strech goals

- Instead of creating an image, let the pixel values stream separately
  so that it can be used for creating interactive visualizations.
