# Lamport

## A Go package implementing Lamport signatures

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