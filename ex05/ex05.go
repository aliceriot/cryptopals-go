/*
---
title: "Exercise 4"
key: "ex05"
---

# Exercise 5

This [exercise](http://cryptopals.com/sets/1/challenges/5) asks us to implement
a simple cryptosystem: repeating-key XOR, which is similar to the Vigenere cipher.

For repeating-key XOR we are going to end up sequentially XORing each byte of
key with each byte of the plaintext, so that the first byte of the plaintext will
be XORed against the first of the key, the 2nd with the 2nd, and so on. Our key is only
3 bytes long however, so in general we'll have:

\\[ C[i] = k[i%len(k)]\ XOR\ P[i] \\]

where \\(P\\) is our plaintext, \\(k\\) is our key, and \\(C\\) is our ciphertext.
The key we're going to be using is `ICE`, and the plaintext is:

```
Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal
```
*/


