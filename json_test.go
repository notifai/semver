package semver

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJSONMarshalValid(t *testing.T) {
	versionString := "3.1.4-alpha.1.5.9+build.2.6.5"
	v, err := Parse(versionString)
	require.NoError(t, err)

	var versionJSON []byte
	versionJSON, err = json.Marshal(v)
	require.NoError(t, err)

	quotedVersionString := strconv.Quote(versionString)

	require.Equal(t,  string(versionJSON), quotedVersionString)
}

func TestJSONMarshalInValid(t *testing.T) {
	var v Version

	v.SetBuild([]string{"?"})

	_, err := json.Marshal(v)
	require.Error(t, err)
}

func TestJSONUnmarshalValid(t *testing.T) {
	versionString := "3.1.4-alpha.1.5.9+build.2.6.5"
	quotedVersionString := strconv.Quote(versionString)

	var v Version
	err := json.Unmarshal([]byte(quotedVersionString), &v)
	require.NoError(t, err)
	require.Equal(t, versionString, v.String(), "JSON unmarshaled semantic version not equal to the origin")
}

func TestJSONUnmarshalInValid(t *testing.T) {
	var v Version
	badVersionString := strconv.Quote("3.1.4.1.5.9.2.6.5-other-digits-of-pi")
	err := json.Unmarshal([]byte(badVersionString), &v)
	require.Error(t, err)
	err = json.Unmarshal([]byte("3.1"), &v)
	require.Error(t, err)
}
