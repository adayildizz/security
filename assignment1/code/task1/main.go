package main

import (
	"fmt"
)

// public shared key
const g = 42

// public shared prime
const p = 29837

// systems public key 
const PK = 22690

// ciphertext pair
const c1 = 23447
const c2 = 8372

func find_private_key(g, PK, p int) int {
	for x := 1; x < p-1; x++ {
		if mod_exp(g, x, p) == PK {
			return x
		}
	}
	return -1 
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

func mod_inverse(s, p int) int {
	if s%p == 0 {
		panic("no modular inverse exists (s divisible by p)")
	}
	return mod_exp(s, p-2, p)
}



func recover_message(x int, p int, c1 int , c2 int)(int){
	s := int(mod_exp(c1, x, p))
	sinv := mod_inverse(s, p)
	message := (c2 * sinv) % p
	return  message

}

func main(){
	sk := find_private_key(g,PK, p)
	message:=recover_message(sk, p, c1, c2)
	fmt.Printf("Original message: %d",message)
}


