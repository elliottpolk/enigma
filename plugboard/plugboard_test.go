package plugboard

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPlugboard(t *testing.T) {
	pb, err := NewPlugboard([]*Connection{
		{A: 'A', B: 'D'},
		{A: 'F', B: 'T'},
		{A: 'W', B: 'H'},
		{A: 'J', B: 'O'},
		{A: 'P', B: 'N'},
	})
	require.NoError(t, err)

	got := make([]int, 26)
	for i := 0; i < 26; i++ {
		got[i] = pb.Forward(i)
	}

	want := []int{3, 1, 2, 0, 4, 19, 6, 22, 8, 14, 10, 11, 12, 15, 9, 13, 16, 17, 18, 5, 20, 21, 7, 23, 24, 25}
	require.EqualValues(t, want, got)
}
