/*
---
title: "Exercise 8"
key: "ex08"
---

# Exercise 8

In this problem we're given a big list of base64 encoded strings, and we're
given the task of finding out which one has been encoded with AES-ECB.

Wait, what's ECB mode, anyway?

## Electronic Codebook mode

ECB is a particular way of using a block cipher to encrypt a message.
Essentially, we devide the message (plaintext) into identically sizes
blocks, and then we encrypt each block by performing whatever the
encryption operation is. This looks like this (thanks wikipedia!):

![](https://upload.wikimedia.org/wikipedia/commons/thumb/e/e6/ECB_decryption.svg/1202px-ECB_decryption.svg.png)

So any block cipher can be used in ECB mode. Great! How flexible! Chopping
the message up into separate blocks, each of which is encrypted
independently, also means that the encryption operation can be
parallelized (since the encryption of any block depends only upon that
block's plaintext).

ECB has a fatal flaw though - this separate encryption means that if your
message plaintext contains two blocks that are identical, then your
ciphertext will also have identical blocks (in the corresponding
positions).

This is not good. This means you're vulnerable to some chosen-ciphertext
attacks, and in particular we're going to leak information about the
plaintext contents. An attacker could use the pattern of repeated blocks
to gain information about what type of file is being transmitted, or could
even again information about file contents. Yikes!


## Solution

This should be pretty straightforward. A major failing of ECB mode (and what
allows for chosen-plaintext type attack on it and so on) is that a given
16-byte block in the input (assuming AES-128) will always produce the same
output in the ciphertext.

So we'll just look for that pattern! We're going to basically read in the data,
and then we'll look for the line which has the highest number of repeated
blocks. The exercise strongly implies that the block size is 16.

First, a little utility function to get the data for the exercise:
*/

package ex08

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readExerciseInput() []string {
	lines, _ := ioutil.ReadFile("./ex08.txt")
	return strings.Split(string(lines), "\n")
}

/*
Now we want to write a function which takes a bite slice and returns a count
of the repeated blocks in that byte slice:
*/

func scoreECBmode(line []byte) int {
	counts := make(map[[16]byte]int)

	for i := 0; i < len(line); i += 16 {
		var chunk [16]byte
		copy(chunk[:], line[i:i+16])
		counts[chunk]++
	}

	for _, k := range counts {
		fmt.Println(k)
	}
	fmt.Printf("\n\n")
	// fmt.Println(counts)
	score := 0
	for _, v := range counts {
		if v != 1 {
			score += v
		}
	}
	return score
}

func findECBString() string {
	strings := readExerciseInput()

	for _, s := range strings {
		fmt.Println(scoreECBmode([]byte(s)))
	}
	return "asdf"
}

func solveExercise() {
	findECBString()
}