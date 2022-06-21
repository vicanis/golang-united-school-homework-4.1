package reverse_string

func ReverseString(input string) (output string) {
	for _, c := range input {
		output = string(c) + output
	}

	return output
}
