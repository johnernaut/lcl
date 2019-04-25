package hashdiff

import (
	"reflect"
	"testing"
)

func TestDiffMapKeys(t *testing.T) {
	original := map[string]string{
		"my.super.cool.key": "An awesome value.",
		"another_key":       "An awesome %@ with %d.",
		"and.a.sweet.key":   "Is the greatest thing ever %@!",
	}
	other := map[string]string{
		"my.super.cool.key": "An awesome value.",
		"another_key":       "An awesome %@ with %d.",
	}
	want := map[string]string{
		"and.a.sweet.key": "Is the greatest thing ever %@!",
	}

	got := DiffMapKeys(original, other)
	eq := reflect.DeepEqual(got, want)

	if !eq {
		t.Errorf("DiffMapKeys() = %v, want %v", got, want)
	}
}

func TestDiffMapValues(t *testing.T) {
	original := map[string]string{
		"my.super.cool.key": "An awesome value.",
		"another_key":       "An awesome %@ with %d.",
		"and.a.sweet.key":   "Is the greatest thing ever %@!",
	}
	other := map[string]string{
		"my.super.cool.key": "An awesome value.",
		"another_key":       "An awesome %@ with %d.",
		"and.a.sweet.key":   "New value is different!",
	}
	want := map[string]string{
		"and.a.sweet.key": "Is the greatest thing ever %@!",
	}

	got := DiffMapValues(original, other)
	eq := reflect.DeepEqual(got, want)

	if !eq {
		t.Errorf("DiffMapValues() = %v, want %v", got, want)
	}
}
