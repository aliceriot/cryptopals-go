/*
---
layout: matasano_exercise
title: "Exercise 4"
---

# Exercise 4

In [exercise 4](http://cryptopals.com/sets/1/challenges/4) we're given
326 60-character hex encoded strings, and we're tasked with figuring out
which one has been encoded with single-byte XOR (as seen in [exercise 3](/matasano/ex03.html)).

Should be fun!
*/

package main

/* First off, we're going to be re-using a couple of functions we wrote
for exercise three:
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
algo outline:

- read in lines
- iterate through and check:
	- an XOR encrypted ciphertext will have at least one key
	- when we look at the possible keys for a ciphertext, the most common
	byte will be 'e', 'a', or 't'

*/
