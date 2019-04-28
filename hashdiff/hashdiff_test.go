package hashdiff

import (
	"reflect"
	"testing"
)

func TestDiffMapKeys(t *testing.T) {
	first := map[string]string{
		"my.super.cool.key": "An awesome value.",
		"another_key":       "An awesome %@ with %d that is newer.",
		"and.a.sweet.key":   "Is the greatest thing ever %@!",
	}
	second := map[string]string{
		"my.super.cool.key": "An awesome value.",
		"another_key":       "An awesome %@ with %d that is old.",
	}
	want := map[string]string{
		"another_key":     "An awesome %@ with %d that is newer.",
		"and.a.sweet.key": "Is the greatest thing ever %@!",
	}

	got := DiffMap(first, second)
	eq := reflect.DeepEqual(got, want)

	if !eq {
		t.Errorf("DiffMap() = %v, want %v", got, want)
	}
}
