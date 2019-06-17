package base58

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testValues struct {
	dec []byte
	enc string
}

var n = 5000000
var testPairs = make([]testValues, 0, n)

func initTestPairs() {
	if len(testPairs) > 0 {
		return
	}
	// pre-make the test pairs, so it doesn't take up benchmark time...
	for i := 0; i < n; i++ {
		data := make([]byte, 32)
		rand.Read(data)
		testPairs = append(testPairs, testValues{dec: data, enc: FastBase58Encoding(data)})
	}
}

func randAlphabet() *Alphabet {
	// Permutes [0, 127] and returns the first 58 elements.
	// Like (math/rand).Perm but using crypto/rand.
	var randomness [128]byte
	rand.Read(randomness[:])

	var bts [128]byte
	for i, r := range randomness {
		j := int(r) % (i + 1)
		bts[i] = bts[j]
		bts[j] = byte(i)
	}
	return NewAlphabet(string(bts[:58]))
}

func TestFastEqTrivialEncodingAndDecoding(t *testing.T) {
	for k := 0; k < 10; k++ {
		testEncDecLoop(t, randAlphabet())
	}
	testEncDecLoop(t, BTCAlphabet)
	testEncDecLoop(t, FlickrAlphabet)
}

func testEncDecLoop(t *testing.T, alph *Alphabet) {
	for j := 1; j < 256; j++ {
		var b = make([]byte, j)
		for i := 0; i < 100; i++ {
			rand.Read(b)
			fe := FastBase58EncodingAlphabet(b, alph)
			te := TrivialBase58EncodingAlphabet(b, alph)

			if fe != te {
				t.Errorf("encoding err: %#v", hex.EncodeToString(b))
			}

			fd, ferr := FastBase58DecodingAlphabet(fe, alph)
			if ferr != nil {
				t.Errorf("fast error: %v", ferr)
			}
			td, terr := TrivialBase58DecodingAlphabet(te, alph)
			if terr != nil {
				t.Errorf("trivial error: %v", terr)
			}

			if hex.EncodeToString(b) != hex.EncodeToString(td) {
				t.Errorf("decoding err: %s != %s", hex.EncodeToString(b), hex.EncodeToString(td))
			}
			if hex.EncodeToString(b) != hex.EncodeToString(fd) {
				t.Errorf("decoding err: %s != %s", hex.EncodeToString(b), hex.EncodeToString(fd))
			}
		}
	}
}

func TestFastBase58Encoding(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		testAddr := []string{
			"1QCaxc8hutpdZ62iKZsn1TCG3nh7uPZojq",
			"1DhRmSGnhPjUaVPAj48zgPV9e2oRhAQFUb",
			"17LN2oPYRYsXS9TdYdXCCDvF2FegshLDU2",
			"14h2bDLZSuvRFhUL45VjPHJcW667mmRAAn",
		}

		for ii, vv := range testAddr {
			// num := Base58Decode([]byte(vv))
			// chk := Base58Encode(num)
			num, err := FastBase58Decoding(vv)
			if err != nil {
				t.Errorf("Test %d, expected success, got error %s\n", ii, err)
			}
			chk := FastBase58Encoding(num)
			if vv != string(chk) {
				t.Errorf("Test %d, expected=%s got=%s Address did base58 encode/decode correctly.", ii, vv, chk)
			}
		}
	})
	t.Run("encode-hello", func(t *testing.T) {
		chk := FastBase58Encoding([]byte("hello-world"))
		assert.Equal(t, chk, "StV1DL6KQw7yiqZ")
	})
	t.Run("decode-hello", func(t *testing.T) {
		decoded, err := FastBase58Decoding("StV1DL6KQw7yiqZ")
		assert.Nil(t, err)
		assert.Equal(t, string(decoded), "hello-world")
	})
}
