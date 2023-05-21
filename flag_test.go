package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRecoveryLevel_Set(t *testing.T) {
	assertLevel := func(t *testing.T, want recoveryLevel, s string) {
		t.Helper()
		var got recoveryLevel
		checkError(t, got.Set(s))
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("recoveryLevel.Set(%s): got %v, want %s\n%s", s, got, want, diff)
		}
	}

	assertLevel(t, recoveryLevelLow, "low")
	assertLevel(t, recoveryLevelMedium, "medium")
	assertLevel(t, recoveryLevelHigh, "high")
	assertLevel(t, recoveryLevelHighest, "highest")
}

func TestRecoveryLevel_String(t *testing.T) {
	assertString := func(t *testing.T, want string, r recoveryLevel) {
		t.Helper()
		got := r.String()
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("recoveryLevel.String(): got %q, want %q\n%s", got, want, diff)
		}
	}

	assertString(t, "low", recoveryLevelLow)
	assertString(t, "medium", recoveryLevelMedium)
	assertString(t, "high", recoveryLevelHigh)
	assertString(t, "highest", recoveryLevelHighest)
}

func checkError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Error(err, "unexpected error")
	}
}
