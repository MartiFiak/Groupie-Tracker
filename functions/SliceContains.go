package groupietrackers

func SContains(slice []string, str string) bool { // ! Check if a slice of string contains a string
	for _, element := range slice {
		if element == str {
			return true
		}
	}
	return false
}