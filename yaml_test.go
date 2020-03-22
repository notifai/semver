package semver

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

const (
	testYaml        = "version: 3.1.4-alpha.1.5.9+build.2.6.5\n"
	testYamlInvalid = "version: 3.1.4.1.5.9.2.6.5-other-digits-of-pi\n"
	versionString   = "3.1.4-alpha.1.5.9+build.2.6.5"
)

type testYamlObject struct {
	Version Version `yaml:"version"`
}

func TestYAMLMarshalValid(t *testing.T) {
	var v testYamlObject
	var err error

	v.Version, err = Parse(versionString)
	require.NoError(t, err)

	var versionYAML []byte

	versionYAML, err = yaml.Marshal(&v)
	require.NoError(t, err)

	require.Equal(t, testYaml,  string(versionYAML), "YAML marshaled semantic version not equal to origin")
}

func TestYAMLMarshalInValid(t *testing.T) {
	var v testYamlObject
	v.Version.SetBuild([]string{"?"})

	_, err := yaml.Marshal(&v)
	require.Error(t, err)
}

func TestYAMLUnmarshalValid(t *testing.T) {
	var v testYamlObject

	err := yaml.Unmarshal([]byte(testYaml), &v)
	require.NoError(t, err)
	require.Equal(t, versionString, v.Version.String(), "YAML unmarshaled semantic version not equal to origin")
}

func TestYAMLUnmarshalInValid(t *testing.T) {
	var v testYamlObject
	err := yaml.Unmarshal([]byte(testYamlInvalid), &v)
	require.Error(t, err)
}
