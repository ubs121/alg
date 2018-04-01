package math

// greatest common divisor
func gcd(u, v uint) uint {
    // simple cases (termination)
    if u == v {
		return u
	}

    if u == 0 {
		return v
	}

    if v == 0 {
		return u
	}

    // look for factors of 2
    if u % 2 == 0 { // u is even
        if v % 2 == 1 { // v is odd
			return gcd(u >> 1, v)
		} 
		
		// both u and v are even
		return gcd(u >> 1, v >> 1) << 1
    }

    if v % 2 == 0 { // u is odd, v is even
		return gcd(u, v >> 1)
	}

    // reduce larger argument
    if u > v {
		return gcd((u - v) >> 1, v)
	}

    return gcd((v - u) >> 1, u)
}