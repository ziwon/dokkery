package exec

import (
	"fmt"
	"os/exec"
	"time"
)

func Execute(cmd string) bool {
	_, err := exec.Command("sh", "-c", cmd).Output() //nolint:gosec
	if err != nil {
		fmt.Println(err)
		return false
	}

	time.Sleep(5 * time.Second)
	return true
}
