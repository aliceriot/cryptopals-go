/*
---
layout: matasano_exercise
title: "Exercise 3"
---

# Exercise 3

[Matasano exercise 3](http://cryptopals.com/sets/1/challenges/3/). This asks
us to break a single byte XOR. Basically, this means that we have a
ciphertext which has been XORed against a single byte. This means we
could exhaustively try every byte until we get it, but we're going to at least
try to do something a little more elegant.
*/

package main

import (
	"encoding/hex"
	"fmt"
)

/*
First off, something we're almost certainly going to need is a function
that takes a byte slice and a byte, and XORs the contents of that byte slice
with that byte.

We'll call it `arrayXOR`:
*/

func arrayXOR(in []byte, n byte) []byte {
	out := make([]byte, len(in))
	for i, v := range in {
		out[i] = v ^ n
	}
	return out
}

/*
Then if we want to XOR a particular byte, say, 42, with a byte array we can
just do `xorResult := arrayXOR(in, byte(42))`. Nice!
*/

/*
We'll also need two little helper functions. They each take an XORed byte
array and check that the result has some attribute. The first checks that
the result contains only printable ASCII characters, and the second checks
that the result contains at least one space character (an assumption we make
about the content of the plaintext):
*/

func asciiCheck(bytes []byte) bool {
	for _, c := range bytes {
		if c > 127 || c < 32 {
			return false
		}
	}
	return true
}

func spaceCheck(bytes []byte) bool {
	found := false
	for _, c := range string(bytes) {
		if c == ' ' {
			found = true
		}
	}
	return found
}

/*
## Doing the work

Now that we've written all of the functions we'll need, we can try to
solve the problem!

Basically, we're dealing with a keyspace of all the values a single byte
can take on (0 - 255). This isn't so many, so while we don't want to just
manually check each one, it won't be too much work to check the result of
'decrypting' with each value for certain attributes.

So, we're going to use our two helper functions to now it down!
*/

func main() {
	const cipherText = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	cipherBytes, _ := hex.DecodeString(cipherText)

	possibleKeys := make([]byte, 0)
	for i := byte(0); i < 255; i++ {
		plain := arrayXOR(cipherBytes, i)
		if asciiCheck(plain) && spaceCheck(plain) {
			possibleKeys = append(possibleKeys, byte(i))
		}
	}

	for _, k := range possibleKeys {
		plain := arrayXOR(cipherBytes, k)
		fmt.Println(string(plain))
	}
}

/*
Great! After all that we get two possible keys out the end, and one of
those produces an obvious plaintext.
*/
