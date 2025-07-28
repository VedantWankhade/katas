package katas

import (
	"fmt"
	"unicode/utf8"
)

func strings_runes() {
	s1 := "hello"
	s2 := "helló"

	// counts bytes
	fmt.Println(len(s1))
	fmt.Println(len(s2))

	// counts runes
	fmt.Println(utf8.RuneCountInString(s1))
	fmt.Println(utf8.RuneCountInString(s2))

	for _, c := range s1 {
		fmt.Print(string(c))
	}

	fmt.Println()

	for _, c := range s2 {
		fmt.Print(string(c))
	}

	fmt.Println()
	for i := range s1 {
		fmt.Print(string(s1[i]))
	}
	fmt.Println()
	for i := range s2 {
		fmt.Print(string(s2[i]))
	}

	fmt.Println()

	for _, r := range []rune(s1) {
		fmt.Print(string(r))
	}

	fmt.Println()

	for _, r := range []rune(s2) {
		fmt.Print(string(r))
	}
	fmt.Println()

	for _, r := range []byte(s1) {
		fmt.Print(string(r))
	}

	fmt.Println()

	for _, r := range []byte(s2) {
		fmt.Print(string(r))
	}
	fmt.Println()
}
