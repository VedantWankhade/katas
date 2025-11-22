package caesarcipher

import "testing"

func TestCaesarCipherEncrypt(t *testing.T) {
	tests := []struct {
		plaintext          string
		key                int
		expectedCipherText string
	}{
		// existing tests
		{"ABC", 1, "BCD"},
		{"ABC", 2, "CDE"},
		{"XYZ", 1, "YZ "},
		{"XYZG", 2, "Z AI"},

		// new edge cases
		{"ABC", 0, "ABC"},                 // zero shift
		{"ABC", 27, "BCD"},                // key > alphabet length (27 ≡ 1)
		{"ABC", 54, "ABC"},                // 54 ≡ 0
		{" ", 5, "F"},                     // space shifts to letter
		{"HELLO WORLD", 1, "IFMMPAXPSME"}, // mixed spaces/letters
		{"hello world", 1, "IFMMPAXPSME"}, // lowercase → uppercase
		{"", 5, ""},                       // empty input
		{"ZZZ", 1, "   "},                 // full wrap into spaces
		{"   ", 1, "AAA"},                 // spaces wrap to A
	}

	for _, test := range tests {
		actualCipherText := Encrypt(test.plaintext, test.key)
		if actualCipherText != test.expectedCipherText {
			t.Error("\nExpected:", test.expectedCipherText, "\nACTUAL:", actualCipherText)
		}
	}
}

func TestCaesarCipherDecrypt(t *testing.T) {
	tests := []struct {
		cipherText        string
		key               int
		expectedPlainText string
	}{
		// existing tests
		{"BCD", 1, "ABC"},
		{"CDE", 2, "ABC"},
		{"YZ ", 1, "XYZ"},
		{"Z AI", 2, "XYZG"},

		// new edge cases
		{"ABC", 0, "ABC"},                 // zero shift
		{"BCD", 27, "ABC"},                // key > alphabet length
		{"ABC", 54, "ABC"},                // 54 ≡ 0
		{"F", 5, " "},                     // letter back to space
		{"IFMMPAXPSME", 1, "HELLO WORLD"}, // multi-word
		{"IFMMPAXPSME", 1, "HELLO WORLD"}, // lowercase equivalent
		{"", 3, ""},                       // empty input
		{"   ", 1, "ZZZ"},                 // spaces wrap to Z
		{"AAA", 1, "   "},                 // wrap back to spaces
	}

	for _, test := range tests {
		actualPlainText := Decrypt(test.cipherText, test.key)
		if actualPlainText != test.expectedPlainText {
			t.Error("\nExpected:", test.expectedPlainText, "\nACTUAL:", actualPlainText)
		}
	}
}
