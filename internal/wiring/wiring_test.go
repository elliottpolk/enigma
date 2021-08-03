package wiring

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecode(t *testing.T) {
	got := Decode([]rune("EKMFLGDQVZNTOWYHXUSPAIBRCJ"))
	require.NotEmpty(t, got)

	want := Wiring([]int{4, 10, 12, 5, 11, 6, 3, 16, 21, 25, 13, 19, 14, 22, 24, 7, 23, 20, 18, 15, 0, 8, 1, 17, 2, 9})
	require.EqualValues(t, want, got)
}

func TestInverse(t *testing.T) {
	got := Inverse(Decode([]rune("EKMFLGDQVZNTOWYHXUSPAIBRCJ")))
	require.NotEmpty(t, got)

	want := Wiring([]int{20, 22, 24, 6, 0, 3, 5, 15, 21, 25, 1, 4, 2, 10, 12, 19, 7, 23, 18, 11, 17, 8, 13, 16, 14, 9})
	require.EqualValues(t, want, got)
}
