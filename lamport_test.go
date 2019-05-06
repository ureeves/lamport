package lamport_test

import (
	"crypto"
	"crypto/rand"
	"crypto/sha256"
	_ "crypto/sha256"
	"testing"

	"github.com/ureeves/lamport"
)

var messageHash = sha256.Sum256([]byte("message to sign"))

func TestLamport(t *testing.T) {
	k, err := lamport.GenerateKey(rand.Reader, crypto.SHA256)

	if err != nil {
		t.Fatalf("failed generating key: %s", err)
	}

	sig, err := lamport.Sign(k, messageHash[:])
	if err != nil {
		t.Fatalf("failed signing with key")
	}

	if err := lamport.Verify(&k.PublicKey, messageHash[:], sig); err != nil {
		t.Fatalf("signature generated was invalid: %s", err)
	}
}

func BenchmarkGenerate(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lamport.GenerateKey(rand.Reader, crypto.SHA256) // nolint, only for benchmark purposes
	}
}

func BenchmarkSign(b *testing.B) {
	k, err := lamport.GenerateKey(rand.Reader, crypto.SHA256)

	if err != nil {
		b.Fatalf("failed generating key: %s", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lamport.Sign(k, messageHash[:]) // nolint, only for benchmark purposes
	}
}

func BenchmarkVerify(b *testing.B) {
	k, err := lamport.GenerateKey(rand.Reader, crypto.SHA256)

	if err != nil {
		b.Fatalf("failed generating key: %s", err)
	}

	sig, err := lamport.Sign(k, messageHash[:])
	if err != nil {
		b.Fatalf("failed signing with key")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lamport.Verify(&k.PublicKey, messageHash[:], sig) // nolint, only for benchmark purposes
	}
}
