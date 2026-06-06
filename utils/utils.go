package utils

import (
	"reflect"
	"sort"
)

func SlicesString2DFuzzyEqual(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	normalize := func(s [][]string) [][]string {
		res := make([][]string, len(s))

		for i, inner := range s {
			cp := append([]string(nil), inner...)
			sort.Strings(cp)
			res[i] = cp
		}

		sort.Slice(res, func(i, j int) bool {
			for k := 0; k < len(res[i]) && k < len(res[j]); k++ {
				if res[i][k] != res[j][k] {
					return res[i][k] < res[j][k]
				}
			}
			return len(res[i]) < len(res[j])
		})

		return res
	}

	return reflect.DeepEqual(normalize(a), normalize(b))
}

func SlicesInt2DFuzzyEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	normalize := func(s [][]int) [][]int {
		res := make([][]int, len(s))

		for i, inner := range s {
			cp := append([]int(nil), inner...)
			sort.Ints(cp)
			res[i] = cp
		}

		sort.Slice(res, func(i, j int) bool {
			for k := 0; k < len(res[i]) && k < len(res[j]); k++ {
				if res[i][k] != res[j][k] {
					return res[i][k] < res[j][k]
				}
			}
			return len(res[i]) < len(res[j])
		})

		return res
	}

	return reflect.DeepEqual(normalize(a), normalize(b))
}
