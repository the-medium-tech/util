package types

import (
	"github.com/stretchr/testify/require"

	"fmt"
	"strconv"
	"testing"
)

type Case struct {
	val      uint64
	expected int
}

var cases = []Case{
	{0x0, 0},
	{0x1, 1},
	{1 << 30, 1},
	{0b11110111, 7},
	{uint64(0b10000000000000000000000000000000), 1},
	{uint64(0b1000000000000000000000000000000000000000000000000000000000000000), 1},
}

func TestOnCount(t *testing.T) {
	t.Log("IntSize", strconv.IntSize)
	for _, c := range cases {
		bitTag := &BitTag{Tag: c.val}
		require.Equal(t, c.expected, bitTag.Count(true), fmt.Sprintf("%b", bitTag.Tag))
	}
}
