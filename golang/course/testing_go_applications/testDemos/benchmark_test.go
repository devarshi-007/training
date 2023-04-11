package testdemos

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"testing"
)

/*

go test
go test <pkg> ...
go test -bench [. | regex] [default is 1S]
go test -bench . -benchmen
go test -bench -benthtime 10s
go test -bench Alloc -memprofile profile.out
go test -bench=. -cpuprofile simple.out
go tool pprof simple.out

*/

func BenchmarkSHA1(b *testing.B) {
	data := []byte("this is something")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sha1.Sum(data)
	}
}

func BenchmarkSHA256(b *testing.B) {
	data := []byte("this is something")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sha256.Sum256(data)
	}
}

func BenchmarkSHA512(b *testing.B) {
	data := []byte("this is something")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sha512.Sum512(data)
	}
}

func BenchmarkSHA512B(b *testing.B) {
	data := []byte("this is something")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		h := sha512.New()
		h.Sum(data)
	}
}
