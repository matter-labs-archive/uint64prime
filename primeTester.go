package uint64prime

const primeListSize = 256

var primeTable = [primeListSize]uint64{3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59,
	61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137,
	139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227,
	229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283, 293, 307, 311, 313,
	317, 331, 337, 347, 349, 353, 359, 367, 373, 379, 383, 389, 397, 401, 409, 419,
	421, 431, 433, 439, 443, 449, 457, 461, 463, 467, 479, 487, 491, 499, 503, 509,
	521, 523, 541, 547, 557, 563, 569, 571, 577, 587, 593, 599, 601, 607, 613, 617,
	619, 631, 641, 643, 647, 653, 659, 661, 673, 677, 683, 691, 701, 709, 719, 727,
	733, 739, 743, 751, 757, 761, 769, 773, 787, 797, 809, 811, 821, 823, 827, 829,
	839, 853, 857, 859, 863, 877, 881, 883, 887, 907, 911, 919, 929, 937, 941, 947,
	953, 967, 971, 977, 983, 991, 997, 1009, 1013, 1019, 1021, 1031, 1033, 1039, 1049, 1051,
	1061, 1063, 1069, 1087, 1091, 1093, 1097, 1103, 1109, 1117, 1123, 1129, 1151, 1153, 1163, 1171,
	1181, 1187, 1193, 1201, 1213, 1217, 1223, 1229, 1231, 1237, 1249, 1259, 1277, 1279, 1283, 1289,
	1291, 1297, 1301, 1303, 1307, 1319, 1321, 1327, 1361, 1367, 1373, 1381, 1399, 1409, 1423, 1427,
	1429, 1433, 1439, 1447, 1451, 1453, 1459, 1471, 1481, 1483, 1487, 1489, 1493, 1499, 1511, 1523,
	1531, 1543, 1549, 1553, 1559, 1567, 1571, 1579, 1583, 1597, 1601, 1607, 1609, 1613, 1619, 1621}

// var lookup = uint64(2)

// func init() {
// 	for i := 0; i < primeListSize; i++ {
// 		lookup = lookup * primeTable[i]
// 	}
// 	lookup--
// 	fmt.Println(lookup)
// }

func mulMod(a, b, m uint64) uint64 {
	/* Calculate ab (mod m)
	**
	** Decompose into sum of powers of 2 (mod m)
	 */
	r := uint64(0)
	aCopy := a
	bCopy := b

	for bCopy != 0 {
		if bCopy&1 != 0 {
			r2 := r + aCopy
			if r2 < r {
				r2 -= m
			}
			r = r2 % m
		}
		bCopy >>= 1
		if bCopy != 0 {
			a2 := a + a
			if a2 < a {
				a2 -= m
			}
			aCopy = a2 % m
		}
	}
	return r
}

func powMod(a, b, m uint64) uint64 {
	/* Calculate a^b (mod m)
	 **
	 ** Decomposes into product of squares (mod m)
	 */
	r := uint64(1)
	aCopy := a
	bCopy := b
	for bCopy != 0 {
		if bCopy&1 != 0 {
			r = mulMod(r, aCopy, m)
		}
		bCopy >>= 1
		if bCopy != 0 {
			aCopy = mulMod(aCopy, aCopy, m)
		}
	}
	return r
}

func isStrongProbablePrime(n, a uint64) bool {
	/* Calculate d/s representation of n */
	d := n - 1
	s := uint8(0)
	for d&0xff != 0 {
		d >>= 8
		s += 8
	}
	if d&0xf != 0 {
		d >>= 4
		s += 4
	}
	if d&0x3 != 0 {
		d >>= 2
		s += 2
	}
	if d&0x1 != 0 {
		d >>= 1
		s++
	}
	// Calculate a^d(mod n)
	b := powMod(a, d, n)
	if (b == 1) || b == (n-1) {
		return true
	}
	for r := uint8(1); r < s; r++ {
		b = mulMod(b, b, n)
		if b <= 1 {
			return false
		}
		if b == (n - 1) {
			return true
		}
	}
	return false
}

// IsPrime tests if uint64 is prime or not
func IsPrime(n uint64) bool {
	// Catch easy answers
	if n < 2 {
		return false
	} // 0 and 1 are not prime
	if n < 4 {
		return true
	} // 2 and 3 are prime
	if n&1 == 0 {
		return false
	} // Even numbers are not

	// Perform trial division
	tmp := uint64(0)
	for i := 0; i < primeListSize; i++ {
		tmp = primeTable[i]
		if n == tmp {
			return true
		}
		if n%tmp == 0 {
			return false
		}
	}
	// if n < 2627641 {
	// 	return true
	// }
	// Next step, SPRP tests
	//
	// Thresholds from Sloan sequence A014233
	if !isStrongProbablePrime(n, 2) {
		return false
	}
	if n < 2047 {
		return true
	}
	if !isStrongProbablePrime(n, 3) {
		return false
	}
	if n < 1373653 {
		return true
	}
	if !isStrongProbablePrime(n, 5) {
		return false
	}
	if n < 25326001 {
		return true
	}
	if !isStrongProbablePrime(n, 7) {
		return false
	}
	// if n<3215031751 {return true}
	if n == 3215031751 {
		return false
	}
	if n < 118670087467 {
		return true
	}
	if !isStrongProbablePrime(n, 11) {
		return false
	}
	if n < 2152302898747 {
		return true
	}
	if !isStrongProbablePrime(n, 13) {
		return false
	}
	if n < 3474749660383 {
		return true
	}
	if !isStrongProbablePrime(n, 17) {
		return false
	}
	if n < 341550071728321 {
		return true
	}
	if !isStrongProbablePrime(n, 19) {
		return false
	}
	if !isStrongProbablePrime(n, 23) {
		return false
	}
	if n < 3825123056546413051 {
		return true
	}
	if !isStrongProbablePrime(n, 29) {
		return false
	}
	if !isStrongProbablePrime(n, 31) {
		return false
	}
	if !isStrongProbablePrime(n, 37) {
		return false
	}
	// This test passes for n<2^64
	return true
}
