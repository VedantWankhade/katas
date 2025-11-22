package caesarcipher

import (
	"slices"
	"strings"
)

/*
Encrypts plaintext into cipher text with key.

Key should be >= 1

Will convert all chars to uppercase
*/
func Encrypt(plaintext string, key int) string {
	letters := []rune{' ', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	plainTextRunes := []rune(strings.ToUpper(plaintext))
	key = key % len(letters)

	cipherRunes := []rune{}

	for _, r := range plainTextRunes {
		runeIndex := slices.Index(letters, r)
		cipheredIndex := (runeIndex + key) % len(letters)
		cipherRunes = append(cipherRunes, letters[cipheredIndex])
	}

	return string(cipherRunes)
}

/*
Decrypts ciphertext with key.

Key should be >= 1
*/
func Decrypt(cipherText string, key int) string {
	letters := []rune{' ', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	cipherTextRunes := []rune(strings.ToUpper(cipherText))

	plainRunes := []rune{}

	for _, r := range cipherTextRunes {
		runeIndex := slices.Index(letters, r)
		cipheredIndex := (runeIndex - key) % len(letters)
		if cipheredIndex < 0 {
			cipheredIndex = len(letters) + cipheredIndex
		}
		plainRunes = append(plainRunes, letters[cipheredIndex])
	}

	return string(plainRunes)
}
