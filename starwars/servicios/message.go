package servicios

import "fmt"

func GetMessage(messages ...[]string) (msg string) {
	var totalMessage []string
	for _, value := range messages {
		for _, message := range value {
			totalMessage = append(totalMessage, message)
		}
	}
	// Se eliminan duplicados
	result := removeDuplicates(totalMessage)
	var stringResult string
	for _, v := range result {
		if v != "" && v != " " {
			stringResult = stringResult + " " + v
		}
	}
	fmt.Println(stringResult)
	return stringResult
}

func removeDuplicates(elements []string) []string {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{} // change string to int here if required
	result := []string{}             // change string to int here if required

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}
