package sys

import (
	"log"
	"os"
	"syscall"
)

func KillProcess(pid int) {
	process, err := os.FindProcess(pid)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = process.Signal(syscall.Signal(15))
}
