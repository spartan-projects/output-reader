package sys

import (
	"flag"
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

func GetCmdParams() (string, int) {
	var jobId string
	var pid int

	flag.StringVar(&jobId, "job", common.JobParamDefaultValue, "Test job id")
	flag.IntVar(&pid, "process", common.PidParamDefaultValue, "Qemu process id")

	flag.Parse()

	return jobId, pid
}