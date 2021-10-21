package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/spartan-projects/output-reader/export"
	"github.com/spartan-projects/output-reader/filter"
	"github.com/spartan-projects/output-reader/sys"
	"io"
	"log"
	"os"
	"strings"
)

var namedPipeFile = "/vxworks/comms/input.out"

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("###### Starting Output Reader Script ######")

	jobIdParam, pid := getCmdParams()
	jobIdFileName := fmt.Sprintf("%s.log", jobIdParam)
	bucketKey := fmt.Sprintf("test-job-logs/%s", jobIdFileName)

	namedPipeFile, err := os.OpenFile(namedPipeFile, os.O_RDONLY, os.ModeNamedPipe)
	if err != nil {
		log.Fatal(err)
	}

	defer closeFile(namedPipeFile)

	fileOutput, err := os.OpenFile(jobIdFileName, os.O_CREATE | os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}

	defer closeFile(fileOutput)

	log.Println("###### Processing Named Pipe Output ######")
	processPipe(namedPipeFile, fileOutput, pid)

	log.Println("###### Filter FileOutput Content ######")
	filter.FileOutputFilter(jobIdFileName)

	log.Println("###### Uploading File to S3 ######")
	export.UploadFile(jobIdFileName, "vandv-common-store", bucketKey)
}

func getCmdParams() (string, int){
	var jobId string
	var pid int

	flag.StringVar(&jobId, "job", "ebf0001", "Test job id")
	flag.IntVar(&pid, "process", 0, "Qemu process id")

	flag.Parse()

	return jobId, pid
}

func processPipe(namedPipe *os.File, outputFile *os.File, pid int) {
	r := bufio.NewReader(namedPipe)
	buf := make([]byte, 0, 6 * 1024)
	sb := make([]string, 0)
	log.Println("###### Processing output file ######")

	for {
		n, err := r.Read(buf[:cap(buf)])
		buf = buf[:n]
		if n == 0 {
			if err == nil {
				continue
			}
			if err == io.EOF {
				log.Println("###### EOF Reached ######")
				break
			}

			log.Fatal(err)
		}

		writeBuffer(outputFile, buf, n)
		sb = append(sb, string(buf))
		closeInputPipe(sb, pid)

		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
	}
}

func writeBuffer(outputFile *os.File, buf []byte, n int) {
	if _, err := outputFile.Write(buf[:n]); err != nil {
		log.Fatal(err)
	}
}

func closeInputPipe(strBuff []string, pid int) {
	ct := strings.Join(strBuff, "")

	if filter.EofFilter(ct) {
		sys.KillProcess(pid)
	}
}

func closeFile(fileToClose *os.File) {
	log.Printf("###### Closing file %s ######", fileToClose.Name())

	if err := fileToClose.Close(); err != nil {
		log.Fatal(err)
	}
}