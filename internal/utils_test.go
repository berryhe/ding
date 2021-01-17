package internal

import (
	"math/rand"
	"strconv"
	"testing"
)

func TestGroupAlg(t *testing.T) {
	c := 10
	randStrArr := make([]string, c)
	for i := 0; i < c; i++ {
		randStrArr[i] = strconv.Itoa(rand.Intn(7))
	}

	t.Logf("%+v\n", randStrArr)
	res := StrArrGroupAlg(randStrArr, 50)
	t.Log(len(res))
	t.Logf("%+v\n", res)
}

func TestFlipByteSlice(t *testing.T) {
	c := []byte("12345678")
	res := FlipByteSlice(c)
	t.Log(string(res))
}
