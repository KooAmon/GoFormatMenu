package menu

import (
	"fmt"
	"os/exec"
)

func Execute(command *string) (int, error) {
	cmd := exec.Command("cmd.exe", "/C", "start", *command)
	if err := cmd.Run(); err != nil {
		fmt.Println("Error:", err)
		return -1, err
	}
	return cmd.Process.Pid, nil
}
