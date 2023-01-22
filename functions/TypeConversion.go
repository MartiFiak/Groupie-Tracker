package groupietrackers

import (
	"fmt"
	"strconv"
)

func AtoiWithoutErr(str string) int {
	// ! Uses the Atoi function by handling the errors it may generate
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
	}
	return i
}
