# randstr

Go library for generating random strings.

from https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go

## Usage

```go
package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/winebarrel/randstr"
)

func main() {
	src := rand.NewSource(time.Now().UnixNano())
	str := randstr.String(src, 16)
	fmt.Println(str)
}
```
