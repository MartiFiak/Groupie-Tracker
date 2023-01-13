package groupietrackers

import (
	"fmt"
	"strconv"
)

func AtoiWithoutErr(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
	}
	return i
}
