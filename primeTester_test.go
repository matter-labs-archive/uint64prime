package uint64prime

import (
	"testing"
)

func TestPrimeSearch(t *testing.T) {
	max := uint64(3825123056546413051)
	_ = IsPrime(max)
}

func BenchmarkPrimeSearch(b *testing.B) {
	max := uint64(3825123056546413051)
	for i := 0; i < b.N; i++ {
		_ = IsPrime(max)
	}
}
