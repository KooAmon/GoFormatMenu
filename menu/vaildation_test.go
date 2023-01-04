package menu

import (
	"errors"
	"testing"
)

func Test_Menu_ValidateUserInputIsInteger(t *testing.T) {
	userInput := 3
	if !ValidateUserInputIsInteger(&userInput, 10) {
		t.Error("Error: Validation of Integer Failed")
	}
}

func Test_Menu_ValidateUserInputIsInteger_Too_High(t *testing.T) {
	userInput := 3
	if ValidateUserInputIsInteger(&userInput, 2) {
		t.Error("Error: Validation of Integer Too High Failed")
	}
}

func Test_Menu_ValidateUserInputIsInteger_Too_Low(t *testing.T) {
	userInput := -1
	if ValidateUserInputIsInteger(&userInput, 10) {
		t.Error("Error: Validation of IntegerToo Low Failed")
	}
}

func Test_Menu_ValidateUserInputIsNotAnError_Has_No_Error(t *testing.T) {
	if !ValidateUserInputIsNotAnError(nil) {
		t.Error("Error: Validation of No Error threw an Error")
	}
}

func Test_Menu_ValidateUserInputIsNotAnError_Has_An_Error(t *testing.T) {
	err := errors.New("test error")
	if ValidateUserInputIsNotAnError(err) {
		t.Error("Error: Validation of an Error Failed")
	}
}
