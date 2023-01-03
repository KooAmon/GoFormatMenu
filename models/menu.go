package models

import (
	"encoding/json"
	"fmt"
	"os"
)

type Menu struct {
	Header    string
	LineWidth int
	Colors    Menucolors
	MenuItems []MenuItem
}

func GetMenu(path *string) Menu {
	fileBytes, err := os.ReadFile(*path)
	if err != nil {
		panic(err)
	}
	fileString := string(fileBytes)

	var currentMenu Menu
	err = json.Unmarshal([]byte(fileString), &currentMenu)
	if err != nil {
		fmt.Println("error:", err)
	}

	if currentMenu.LineWidth == 0 {
		currentMenu.LineWidth = 80
	}

	return currentMenu
}
