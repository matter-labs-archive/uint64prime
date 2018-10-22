package uint64prime

import (
	"testing"
)

func TestPrimeSearch(t *testing.T) {
	max := uint64(18446744073709551557) // 2^64 - 59 - largest 64bit prime
	isPrime := IsPrime(max)
	if !isPrime {
		t.Fatal("Should return prime!")
	}
}

func BenchmarkPrimeSearch(b *testing.B) {
	max := uint64(18446744073709551557) // 2^64 - 59 - largest 64bit prime
	for i := 0; i < b.N; i++ {
		_ = IsPrime(max)
	}
}
