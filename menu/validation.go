package menu

import "fmt"

func ValidateUserInputIsInteger(maxLength int) *int {
	var userInput int
	for {
		fmt.Print("Select Item Index: ")
		_, err := fmt.Scanf("%d", &userInput)

		if err != nil {
			if err.Error() == "unexpected newline" {
				continue
			} else {
				fmt.Println(err, ":Not a valid Index - try again")
				continue
			}
		}
		if userInput > maxLength-1 {
			fmt.Println("Index is above TotalMenuItems - try again")
			continue
		}
		if userInput < 0 {
			fmt.Println("Index is below 0 - try again")
			continue
		}
		break
	}
	return &userInput
}
