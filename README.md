# Knuth-Morris-Pratt-Algorithm
Implementation of the Knuth-Morris-Pratt algorithm or KMP algorithm in GO

## Technology
* [go](https://golang.org/)
* [ginkgo](https://github.com/onsi/ginkgo)

## Usage
CLI:
```bash
    go test -v test/*.go # for test
```

CODE:
```go
import "github.com/AlgoCafe/Knuth-Morris-Pratt-Algorithm/kmp"

var index, length int
var buffer = []byte("abc abcdab abcdabcdabde")
var pattern = []byte("abcdabd") // 7
var pattern1 = []byte("abcdabdr") // 8

index = kmp.Search(buffer, pattern) // 15
index = kmp.Search(buffer, pattern1) // -1
```

## KMP Explications
 - [Explications en francais](doc/FR-EXPLICATION.md)
 - [English Explications](doc/EN-EXPLICATION.md)

## Links
https://en.wikipedia.org/wiki/Knuth%E2%80%93Morris%E2%80%93Pratt_algorithm  
https://dev.to/girish3/string-matching-kmp-algorithm-cie  
https://www.geeksforgeeks.org/kmp-algorithm-for-pattern-searching/  
