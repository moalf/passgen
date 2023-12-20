package main

import "testing"

func TestMinPasswordLength(t *testing.T) {
	got := RandString(minPasswordLength)

	if len(got) != minPasswordLength {
		t.Errorf("got %d, wanted %d	", len(got), minPasswordLength)
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
		got := isComplex(tc.password)
		if got != tc.result {
			t.Errorf("for `%s`, got %t, wanted %t", tc.password, got, tc.result)
		}
	}

}

func BenchmarkRandString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandString(minPasswordLength)
	}
}
