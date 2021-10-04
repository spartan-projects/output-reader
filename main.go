package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var namedPipeFile = "/home/valkyrie/vxworks-images/workspace_io/vip_intel_test/comms/input.out"

func main() {
	nBytes, nChunks := int64(0), int64(0)

	jobIdParam := getCmdParams()
	jobIdFileName := fmt.Sprintf("%s.log", jobIdParam)

	namedPipeFile, err := os.OpenFile(namedPipeFile, os.O_RDONLY, os.ModeNamedPipe)
	if err != nil {
		log.Fatal(err)
	}

	defer closeFile(namedPipeFile)

	// TODO change output namedPipeFile name with job id - example ebf0001
	fileOutput, err := os.OpenFile(jobIdFileName, os.O_CREATE | os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	defer closeFile(fileOutput)

	processPipe(namedPipeFile, fileOutput, nBytes, nChunks)
}

func getCmdParams() string{
	var jobId string

	if len(jobId) > 0 {
		flag.StringVar(&jobId, "job", "", "Test job id")
	} else {
		flag.StringVar(&jobId, "job", "ebf0001", "Test job id")
	}

	flag.Parse()

	return jobId
}

func processPipe(namedPipe *os.File, outputFile *os.File, nBytes int64, nChunks int64) {
	r := bufio.NewReader(namedPipe)
	buf := make([]byte, 0, 6 * 1024)
	for {
		n, err := r.Read(buf[:cap(buf)])
		buf = buf[:n]
		if n == 0 {
			if err == nil {
				continue
			}
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		nChunks++
		nBytes += int64(len(buf))

		if _, err := outputFile.Write(buf[:n]); err != nil {
			panic(err)
		}

		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
	}
	log.Println("Bytes:", nBytes, "Chunks:", nChunks)
}

func closeFile(fileToClose *os.File) {
	if err := fileToClose.Close(); err != nil {
		panic(err)
	}
}