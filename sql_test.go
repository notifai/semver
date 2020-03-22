package semver

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type scanTest struct {
	val      interface{}
	expected string
}

var scanTestsValid = []scanTest{
	{"1.2.3", "1.2.3"},
	{[]byte("1.2.3"), "1.2.3"},
}

var scanTestsInValid = []interface{}{
	true,
	false,
	"blabla",
}


func TestScanValid(t *testing.T) {
	for _, tc := range scanTestsValid {
		s := &Version{}
		err := s.Scan(tc.val)
		require.NoError(t, err)
		val, e := s.Value()
		require.NoError(t, e)
		require.Equal(t, tc.expected, val)
	}
}

func TestScanInValid(t *testing.T) {
	for _, tc := range scanTestsInValid {
		s := &Version{}
		err := s.Scan(tc)
		require.Error(t, err)
	}
}
