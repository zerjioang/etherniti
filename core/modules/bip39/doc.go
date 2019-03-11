// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package bip39

/*

# BIP39: Mnemonic code for generating deterministic keys

BIPs : This BIP describes the implementation of a mnemonic code or mnemonic sentence - a group of easy to remember words - for the generation of deterministic wallets.
The English-language wordlist for the BIP39 standard has 2048 words, so if the phrase contained only 12 random words, the number of possible combinations would be 204⁸¹² = ²¹³² and the phrase would have 132 bits of security. Actual security of a 12-word BIP39 seed phrase is only 128 bits.
Generally a seed phrase only works with the same wallet software that created it. If storing for a long period of time it's a good idea to write the name of the wallet too.

BIP39: Mnemonic code for generating deterministic keys , BIP39 - used to manage your recovery seed and recovery words. Abstract. This BIP describes the implementation of a mnemonic code or mnemonic sentence - a group of easy to remember words - for the generation of deterministic wallets. It consists of two parts: generating the mnemonic, and converting it into a binary seed. Example: let bip39 = require("bip39");


# initial package performance

BenchmarkBip39/bip39-generate-4         	  200000	     10377 ns/op	   0.10 MB/s	    2512 B/op	      62 allocs/op

after some minor changes

BenchmarkBip39/bip39-generate-4         	  200000	      7832 ns/op	   0.13 MB/s	    2544 B/op	      63 allocs/op
BenchmarkBip39/is-valid-4   		      	  500000	      3797 ns/op	   0.26 MB/s	     384 B/op	       1 allocs/op


*/
