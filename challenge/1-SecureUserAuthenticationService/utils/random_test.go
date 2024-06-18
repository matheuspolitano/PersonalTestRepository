package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomTest(t *testing.T) {
	testRand := RandString(12)
	require.Len(t, []byte(testRand), 12)
}
