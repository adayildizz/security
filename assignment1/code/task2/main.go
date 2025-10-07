package main

import "fmt"

const M = 26000
const student_number = 12345

// public shared prime
const p = 29837

// ciphertext pair
const c1 = 23447
const c2 = 8372

func mod_inverse(s, p int) int {
	if s%p == 0 {
		panic("no modular inverse exists (s divisible by p)")
	}
	return mod_exp(s, p-2, p)
}

func mod_exp(base, exp, mod int) int {
	result := 1
	base = base % mod
	for exp > 0 {
		if exp%2 == 1 {
			result = (result * base) % mod
		}
		base = (base * base) % mod
		exp /= 2
	}
	return result
}

func modify_ciphertext(c1, c2, T, M, p int) (int, int) {
	modulo_inverse_M := mod_inverse(M, p)
	// the multiplier covered
	m := (modulo_inverse_M * T) % p
	c1_modified := c1
	c2_modified := (c2 * m) % p
	return c1_modified, c2_modified

}

func main() {

	c1_modified, c2_modified := modify_ciphertext(c1, c2, student_number, M, p)
	fmt.Printf("new c1: %d, new c2 %d", c1_modified, c2_modified)

}