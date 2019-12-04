package main

import "testing"

func TestValidPassword(t *testing.T) {
	t.Run("Must be six digits", func(t *testing.T) {
		if validPassword(11111) {
			t.Errorf("Expected 11111 to be invalid, but was valid, must be length 6")
		}
		if validPassword(1111111) {
			t.Errorf("Expected 1111111 to be invalid, but was valid, must be length 6")
		}
		if !validPassword(111111) {
			t.Errorf("Expected 111111 to be valid, but was invalid")
		}
	})

	t.Run("Must have two adjacent numbers duplicated", func(t *testing.T) {
		if validPassword(123456) {
			t.Errorf("Expected 123456 to be invalid, but was valid, must have two adjacent numbers duplicated")
		}
		if !validPassword(112345) {
			t.Errorf("Expected 112345 to be valid, but was invalid")
		}
	})

	t.Run("The digits must always increase", func(t *testing.T) {
		if validPassword(113450) {
			t.Errorf("Expected 123450 to be invalid, but was valid, the digits must always increase")
		}
		if validPassword(010111) {
			t.Errorf("Expected 010111 to be invalid, but was valid, the digits must always increase")
		}
		if !validPassword(111123) {
			t.Errorf("Expected 111123 to be valid, but was invalid")
		}
	})
}

func TestValidPasswordStrict(t *testing.T) {
	t.Run("Must have two adjacent numbers duplicated and not part of a larger group", func(t *testing.T) {
		if !validPasswordStrict(112233) {
			t.Errorf("Expected 112233 to be valid, but was invalid")
		}
		if validPasswordStrict(123444) {
			t.Errorf("Expected 123444 to be invalid, but was valid, must not be part of larger group")
		}
		if !validPasswordStrict(111122) {
			t.Errorf("Expected 111122 to be valid, but was invalid")
		}
	})
}
