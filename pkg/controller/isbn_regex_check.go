package controller

import (
	"regexp"
)
func ValidISBN(isbn string) bool {
	// Define the regular expression pattern for ISBN-13
	isbnPattern := `^(978|979)-\d{1,5}-\d{1,7}-\d{1,7}-\d$`

	// Compile the regular expression
	re := regexp.MustCompile(isbnPattern)

	// Match the ISBN against the regular expression
	return re.MatchString(isbn)
}