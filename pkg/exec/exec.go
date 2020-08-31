package exec

import (
	"fmt"
	"os/exec"
	"time"
)

func Execute(cmd string) string {
	_, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		fmt.Println(err)
		return fmt.Sprintf("Failed to execute command: `%s`", cmd)
	}

	time.Sleep(5 * time.Second)
	return fmt.Sprintf("Succeed: `%s`", cmd)
}
