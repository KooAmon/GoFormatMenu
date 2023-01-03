package render

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

const (
	//	Start Formatting
	ESCAPE = "\033["
	//	End Formating
	ENDC = ESCAPE + "0m"
	//	Clear Screen - Not Used/Working
	CLEAR = ESCAPE + "2J"
)

func RenderCode(code string, args ...interface{}) string {
	var message string
	if ln := len(args); ln == 0 {
		return ""
	}

	message = fmt.Sprint(args...)
	if len(code) == 0 {
		return message
	}

	return ESCAPE + code + "m" + message + ENDC
}

func RenderString(code string, str string) string {
	if len(code) == 0 || str == "" {
		return str
	}

	return ESCAPE + code + "m" + str + ENDC
}

var clear map[string]func()

func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func ClearScreen() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}
