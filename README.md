# Lamport

## A Go package implementing Lamport signatures

[![Build Status](https://travis-ci.org/ureeves/lamport.svg?branch=master)](https://travis-ci.org/ureeves/lamport)
[![codecov](https://codecov.io/gh/ureeves/lamport/branch/master/graph/badge.svg)](https://codecov.io/gh/ureeves/lamport)
[![GoDoc](https://godoc.org/github.com/ureeves/lamport?status.svg)](https://godoc.org/github.com/ureeves/lamport)

Lamport signatures are a one-time signature scheme. Each key is intended to be used only once.

### Usage

Generate key

``` Go
k, err := lamport.GenerateKey(rand.Reader, crypto.SHA256)
```

Sign message

``` Go
message := []byte("Hello, Lamport!")
messageHash := sha256.Sum256(message)

sig, _ := lamport.Sign(k, messageHash[:])
```

Verify signature

``` Go
if lamport.Verify(&k.PublicKey, messageHash[:], sig); err != nil {
    // signature invalid here
}
```

### Testing & Benchmarking

To test on your machine use:

``` Bash
go test
```

If you want to run the benchmarks:

``` Bash
go test -bench=.
```

### Documentation

Documentation can be found [here](https://godoc.org/github.com/ureeves/lamport).
