package hashmap

// Problem: https://leetcode.com/problems/ransom-note/description

func RansomNote(ransomNote string, magazine string) bool {
	if len(magazine) < len(ransomNote) {
		return false
	}

	ransomNoteChars, magazineChars := make(map[rune]int), make(map[rune]int)

	for _, b := range ransomNote {
		if n, ok := ransomNoteChars[b]; ok {
			ransomNoteChars[b] = n + 1
		} else {
			ransomNoteChars[b] = 1
		}
	}
	for _, b := range magazine {
		if n, ok := magazineChars[b]; ok {
			magazineChars[b] = n + 1
		} else {
			magazineChars[b] = 1
		}
	}

	for k, v := range ransomNoteChars {
		if v > magazineChars[k] {
			return false
		}
	}
	return true
}
