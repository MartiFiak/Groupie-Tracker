package groupietrackers

import (
	"strings"
)

func sortByNIndex(i int, slice [][]string)[][]string{
	for pos, elementM := range slice {
		for npos, elementC := range slice[pos:]{
			if AtoiWithoutErr(strings.Split(elementM[i],"-")[0]) < AtoiWithoutErr(strings.Split(elementC[i],"-")[0]){
				slice[pos], slice[npos] = elementC, elementM
			}
		}
	}
	return slice
}