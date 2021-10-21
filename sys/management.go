package sys

import (
	"github.com/spartan-projects/output-reader/common"
	"log"
	"os"
	"syscall"
)

func KillProcess(pid int) {
	log.Printf("##### Killing process with id: %d #####", pid)
	process, err := os.FindProcess(pid)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = process.Signal(syscall.Signal(common.SysKillSignal))
}
