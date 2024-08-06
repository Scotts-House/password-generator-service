package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestGeneratePasswordOnlyChars(t *testing.T) {
	length := int64(128)
	for n := range length {
		testName := fmt.Sprintf("%d chars, no special, no numbers (#%d)", length, n)
		t.Run(testName, func(t *testing.T) {
			p := generatePassword(length, true, false, false)
			if int64(len(p)) != length {
				t.Errorf("Password length does not match. Expected %d, got %d", length, len(p))
			}

			// Make sure it contains characters
			if tf, _ := stringContainsChars(p); !tf {
				t.Errorf("Password contains no characters. Expected %t, got %t; ", true, tf)
			}

			// Make sure it doesn't contains numbers
			if tf, c := stringContainsNumbers(p); tf {
				t.Errorf("Password contains no numbers. Expected %t, got %t; Number Found: %s", false, tf, c)
			}

			// Make sure there are no symbols
			if tf, c := stringContainsSymbols(p); tf {
				t.Errorf("Password contains numbers. Expected %t, got %t; Symbol Found: %s", false, tf, c)
			}
		})
	}
}

func TestGeneratePasswordLettersNumbers(t *testing.T) {
	length := int64(128)
	for n := range length {
		testName := fmt.Sprintf("%d chars, no special (#%d)", length, n)
		t.Run(testName, func(t *testing.T) {
			p := generatePassword(length, true, false, true)
			if int64(len(p)) != length {
				t.Errorf("Password length does not match. Expected %d, got %d", length, len(p))
			}

			// Make sure it contains characters
			if tf, _ := stringContainsChars(p); !tf {
				t.Errorf("Password contains no characters. Expected %t, got %t; ", false, tf)
			}

			// Make sure it contains numbers
			if tf, _ := stringContainsNumbers(p); !tf {
				t.Errorf("Password contains no numbers. Expected %t, got %t; ", false, tf)
			}

			// Make sure there are no symbols
			if tf, c := stringContainsSymbols(p); tf {
				t.Errorf("Password contains symbols. Expected %t, got %t; Symbol Found: %s", false, tf, c)
			}
		})
	}
}

func TestGeneratePasswordLettersNumbersSymbols(t *testing.T) {
	length := int64(128)
	for n := range length {
		testName := fmt.Sprintf("%d chars (#%d)", length, n)
		t.Run(testName, func(t *testing.T) {
			p := generatePassword(length, true, true, true)
			if int64(len(p)) != length {
				t.Errorf("Password length does not match. Expected %d, got %d", length, len(p))
			}

			// Make sure it contains characters
			if tf, _ := stringContainsChars(p); !tf {
				t.Errorf("Password contains no characters. Expected %t, got %t ", true, tf)
			}

			// Make sure it contains numbers
			if tf, _ := stringContainsNumbers(p); !tf {
				t.Errorf("Password contains no numbers. Expected %t, got %t ", true, tf)
			}

			// Make sure there are no symbols
			if tf, _ := stringContainsSymbols(p); !tf {
				t.Errorf("Password contains no symbols. Expected %t, got %t", true, tf)
			}
		})
	}
}

func stringContainsNumbers(p string) (bool, string) {
	for _, n := range numBytes {
		if strings.Contains(p, string(n)) {
			return true, string(n)
		}
	}
	return false, ""
}

func stringContainsSymbols(p string) (bool, string) {
	for _, n := range specialBytes {
		if strings.Contains(p, string(n)) {
			return true, string(n)
		}
	}
	return false, ""
}

func stringContainsChars(p string) (bool, string) {
	for _, n := range letterBytes {
		if strings.Contains(p, string(n)) {
			return true, string(n)
		}
	}
	return false, ""

}
