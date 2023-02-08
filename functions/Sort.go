package groupietrackers

import (
	"strings"
)

func sortByNIndex(i int, slice [][]string) [][]string { // ! Sorts a slice of slice of type string with respect to elements at index i
	for pos := range slice {
		for npos := range slice[pos:] {
			if AtoiWithoutErr(strings.Split(slice[npos][i], "-")[0]) < AtoiWithoutErr(strings.Split(slice[pos][i], "-")[0]) { // ? Here we take into account that we will potentially sort formatted elements in the form num-num
				tmp := slice[pos]
				slice[pos] = slice[npos]
				slice[npos] = tmp
			}
		}
	}
	return slice
}
