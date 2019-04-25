package parser

import (
	"reflect"
	"testing"
)

func TestParseLines(t *testing.T) {
	lines := []string{
		`"my.super.cool.key" = "An awesome value."`,
		`"another_key" = "An awesome %@ with %d."`,
		`//"commented" = "Commented key"`,
	}
	want := map[string]string{
		"my.super.cool.key": "An awesome value.",
		"another_key":       "An awesome %@ with %d.",
	}

	got := parseLines(lines)
	eq := reflect.DeepEqual(got, want)

	if !eq {
		t.Errorf("parseLines() = %v, want %v", got, want)
	}
}

func TestFormatKey(t *testing.T) {
	key1 := "\n\"key1\""
	want := "key1"
	got, skip := formatKey(key1)
	if got != want || skip != false {
		t.Errorf("formatKey() = %q, want %q", got, want)
	}

	key2 := "\n//\"key1\""
	got2, skip := formatKey(key2)
	if !skip {
		t.Errorf("formatKey() = %q, not skipped", got2)
	}
}

func TestFormatValue(t *testing.T) {
	value := `"A really super cool value with %@ %d"`
	want := `A really super cool value with %@ %d`
	got := formatValue(value)
	if got != want {
		t.Errorf("formatValue() = %q, want %q", got, want)
	}
}
