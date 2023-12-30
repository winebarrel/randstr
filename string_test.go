package randstr_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/randstr"
)

func TestString(t *testing.T) {
	assert := assert.New(t)
	src := rand.NewSource(time.Now().UnixNano())
	str := randstr.String(src, 16)
	assert.Len(str, 16)
}
