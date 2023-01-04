package menu

import "fmt"

func ValidateUserInputIsInteger(userInput *int, maxLength int) bool {
	if *userInput > maxLength-1 {
		fmt.Println("Index is above TotalMenuItems - try again")
		return false
	}

	if *userInput < 0 {
		fmt.Println("Index is below 0 - try again")
		return false
	}

	return true
}

func ValidateUserInputIsNotAnError(err error) bool {
	if err != nil && err.Error() != "unexpected newline" {
		fmt.Println(err, ":Not a valid Index - try again")
		return false
	}
	return true
}
