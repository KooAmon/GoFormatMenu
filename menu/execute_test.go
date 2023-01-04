package menu

import (
	"testing"
)

func Test_Menu_Execute(t *testing.T) {
	testCommand := "cls"
	_, errorCode := Execute(&testCommand)

	if errorCode != nil {
		t.Error("Error:", errorCode)
	}
}

func Test_Menu_Execute_Bad_Input(t *testing.T) {
	testCommand := "badinput"
	_, errorCode := Execute(&testCommand)

	if errorCode == nil {
		t.Error("Error: Bad input did not trigger If Clause")
	}
}
