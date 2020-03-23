package semver

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSort(t *testing.T) {
	v100, _ := Parse("1.0.0")
	v010, _ := Parse("0.1.0")
	v001, _ := Parse("0.0.1")
	versions := []Version{v010, v100, v001}
	Sort(versions)

	correct := []Version{v001, v010, v100}
	require.True(t, reflect.DeepEqual(versions, correct), "Sort returned wrong order: %s", versions)
}

func BenchmarkSort(b *testing.B) {
	v100, _ := Parse("1.0.0")
	v010, _ := Parse("0.1.0")
	v001, _ := Parse("0.0.1")
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		Sort([]Version{v010, v100, v001})
	}
}
