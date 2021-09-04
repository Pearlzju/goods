package goods

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStr2bytes(t *testing.T) {
	s1 := "abc"
	b1 := Str2bytes(s1)
	t.Log(b1)

	s2 := Bytes2str(b1)
	t.Log(s2)

	require.Equal(t, s1, s2)
}

func TestBytes2str(t *testing.T) {
	b1 := []byte{97, 98, 99}
	s1 := Bytes2str(b1)
	t.Log(s1)

	b2 := Str2bytes(s1)
	t.Log(b2)

	require.Equal(t, b1, b2)
}
