package workers

import "strings"

func ExtractOperators(rule string) []string {

	tokens := strings.Fields(rule)

	var operators []string

	i, j := 0, len(tokens)-1

	for i < j {

		for i < len(tokens) && tokens[i] != "AND" && tokens[i] != "OR" {
			i++
		}

		for j >= 0 && tokens[j] != "AND" && tokens[j] != "OR" {
			j--
		}

		if i < j {
			if tokens[i] == "AND" || tokens[i] == "OR" {
				operators = append(operators, tokens[i])
			}

			if i != j && (tokens[j] == "AND" || tokens[j] == "OR") {
				operators = append(operators, tokens[j])
			}

			i++
			j--
		}
	}

	return operators
}

func ExtractOperatorsOutwardsInwards(rule string) []string {
	var operators []string
	openParentheses := 0

	tokens := strings.Fields(rule)

	var outerOperators, innerOperators []string

	for _, token := range tokens {
		if token == "(" {
			openParentheses++
		} else if token == ")" {
			openParentheses--
		} else if token == "AND" || token == "OR" {

			if openParentheses == 1 {
				outerOperators = append(outerOperators, token)
			} else {

				innerOperators = append(innerOperators, token)
			}
		}
	}

	operators = append(outerOperators, innerOperators...)
	return operators
}

func ExtractAndReverseOperators(rule string) []string {
	// Split the rule into tokens by spaces
	tokens := strings.Fields(rule)

	// Initialize a slice to hold operators
	var operators []string

	// Iterate through the tokens and collect all AND/OR operators
	for _, token := range tokens {
		if token == "AND" || token == "OR" {
			operators = append(operators, token)
		}
	}

	// Reverse the order of the operators (for inward fashion)
	reverse(operators)

	return operators
}

// Reverse function to reverse a slice of strings
func reverse(slice []string) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func CustomExtract(rule string) []string {
	// Split the rule into tokens by spaces
	tokens := strings.Fields(rule)

	var operators []string

	// Iterate through the tokens and collect all AND/OR operators
	for _, token := range tokens {
		if token == "AND" || token == "OR" {
			operators = append(operators, token)
		}
	}

	i, j := 0, len(operators)-1
	var result []string

	for i < j {
		// if tokens[i] == "AND" || tokens[i] == "OR" {
		// 	operators = append(operators, tokens[i])
		// }

		// if i != j && (tokens[j] == "AND" || tokens[j] == "OR") {
		// 	operators = append(operators, tokens[j])
		// }

		// i++
		// j--

		result = append(result, operators[i])

		result = append(result, operators[j])
		i++
		j--
	}
	if len(operators) != len(result) {
		n := len(operators) / 2
		result = append(result, operators[n])
	}
	return result
}
