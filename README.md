# knuth-morris-pratt
[EN] Implementation of the Knuth-Morris-Pratt algorithm or KMP algorithm in GO

[FR] Implémentation de l'algorithme de Knuth-Morris-Pratt ou de l'algorithme KMP en GO

## Technology
* [go](https://golang.org/)
* [ginkgo](https://github.com/onsi/ginkgo)

## Usage
CLI:
```bash
go test -v test/
```

CODE:
```go
import "github.com/coursecomputer/knuth-morris-pratt/kmp"

var index, length int
var buffer = []byte("abc abcdab abcdabcdabde")
var pattern = []byte("abcdabd") // 7
var pattern1 = []byte("abcdabdr") // 8

index = kmp.Search(buffer, pattern) // 15
index = kmp.Search(buffer, pattern1) // -1
```

## Explications
 - [English](documentation/EN-EXPLICATION.md)
 - [Francais](documentation/FR-EXPLICATION.md)

## Links
https://en.wikipedia.org/wiki/Knuth%E2%80%93Morris%E2%80%93Pratt_algorithm  
https://dev.to/girish3/string-matching-kmp-algorithm-cie  
https://www.geeksforgeeks.org/kmp-algorithm-for-pattern-searching/  
