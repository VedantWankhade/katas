package hashmap

import (
	"sort"
	"strings"
)

func GroupAnagrams(strs []string) [][]string {
	hash := make(map[string][]string)
	for _, s := range strs {
		canonicalStr := canonical(s)
		if arr, ok := hash[canonicalStr]; !ok {
			hash[canonicalStr] = []string{s}
		} else {
			hash[canonicalStr] = append(arr, s)
		}
	}
	out := [][]string{}
	for _, v := range hash {
		out = append(out, v)
	}

	return out
}

func canonical(str string) string {
	parts := strings.Split(str, "")
	sort.Strings(parts)
	return strings.Join(parts, "")
}

func GroupAnagramsFrequencyArray(strs []string) [][]string {
	hash := make(map[[26]int][]string) // a static array is comparable in go, so we can use if as a key in maps
	for _, s := range strs {
		key := frequencyArr(s)
		if arr, ok := hash[key]; !ok {
			hash[key] = []string{s}
		} else {
			hash[key] = append(arr, s)
		}
	}

	out := [][]string{}
	for _, v := range hash {
		out = append(out, v)
	}

	return out
}

func frequencyArr(str string) [26]int {
	arr := [26]int{}
	offset := 97
	for _, c := range str {
		arr[int(c)-offset]++
	}
	return arr
}
