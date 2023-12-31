package rndstr

import (
	"testing"
)

func TestMinPasswordLength(t *testing.T) {

	got := RandString(MinPasswordLength)

	if len(got) != MinPasswordLength {
		t.Errorf("got %d, wanted %d	", len(got), MinPasswordLength)
	}
}

func TestIsComplex(t *testing.T) {
	tests := []struct {
		password string
		result   bool
	}{
		{`abc`, false},
		{`d#)P:(R6fL,*`, true},
		{`abcdefghijK1!`, true},
		{`0123456789`, false},
		{`ABCDEFGHIJK`, false},
	}

	for _, tc := range tests {
		got := IsComplex(tc.password)
		if got != tc.result {
			t.Errorf("for `%s`, got %t, wanted %t", tc.password, got, tc.result)
		}
	}

}

func TestRandomString(t *testing.T) {
	tests := []struct {
		length         int
		expectedLength int
	}{
		{0, 0},
		{1, 1},
		{12, 12},
		{64, 64},
		{100, 12}, // Defaults to minPasswordLength if password length is greater than maxPasswordLength
	}

	for _, tc := range tests {
		got := RandString(tc.length)
		if len(got) != tc.expectedLength {
			t.Errorf("got %d, expected %d", tc.length, tc.expectedLength)
		}
	}
}

func BenchmarkRandString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandString(MinPasswordLength)
	}
}
