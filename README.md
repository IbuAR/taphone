# TAphone

TAphone is a phonetic algorithm for indexing Tamil words by their pronunciation, like Metaphone for English. The algorithm generates three Romanized phonetic keys (hashes) of varying phonetic affinities for a given Tamil word. This package implements the algorithm in Go.

The algorithm takes into account the context sensitivity of sounds, syntactic and phonetic gemination, compounding, modifiers, and other known exceptions to produce Romanized phonetic hashes of increasing phonetic affinity that are very faithful to the pronunciation of the original Tamil word.

- `key0` = a broad phonetic hash comparable to a Metaphone key that doesn't account for hard sounds and phonetic modifiers
- `key1` = is a slightly more inclusive hash that accounts for hard sounds.
- `key2` = highly inclusive and narrow hash that accounts for hard sounds and phonetic modifiers.

### Examples

| Word     | Pronunciation | key0 | key1  | key2    |
| -------- | ------------- | ---- | ----- | ------- |
| வணக்கம்  | Vaṇakkam      | VNKM | VN1KM | VN1K2M  |
| நீர்     | Nīr           | NR   | NR    | N4R     |
| நிலம்    | Nilam         | NLM  | NLM   | N4LM    |
| நெருப்பு | Neruppu       | NRP  | NRP   | N6R5P25 |
| காற்று   | Kāṟṟu         | KR   | KR    | K3R25   |
| ஆகாயம்   | Ākāyam        | AKYM | AKYM  | AK3YM   |

### Go implementation

Install the package:
`go get -u github.com/IbuAR/taphone`

```go
package main

import (
	"fmt"

	"github.com/IbuAR/taphone"
)

func main() {
	ta := taphone.New()
	fmt.Println(ta.Encode("வணக்கம்"))
	fmt.Println(ta.Encode("நெருப்பு"))
}

```

License: GPLv3
