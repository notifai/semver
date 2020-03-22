package semver

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type scanTest struct {
	val         interface{}
	shouldError bool
	expected    string
}

var scanTests = []scanTest{
	{"1.2.3", false, "1.2.3"},
	{[]byte("1.2.3"), false, "1.2.3"},
	{7, true, ""},
	{7e4, true, ""},
	{true, true, ""},
}

func TestScanString(t *testing.T) {
	for _, tc := range scanTests {
		s := &Version{}
		err := s.Scan(tc.val)
		if tc.shouldError {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
			val, e := s.Value()
			require.NoError(t, e)
			require.Equal(t, tc.expected, val)
		}
	}
}

