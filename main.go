package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/spartan-projects/output-reader/export"
	"github.com/spartan-projects/output-reader/filter"
	"io"
	"log"
	"os"
)

var namedPipeFile = "/vxworks/comms/input.out"

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("###### Starting Output Reader Script ######")
	nBytes, nChunks := int64(0), int64(0)

	jobIdParam := getCmdParams()
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
	processPipe(namedPipeFile, fileOutput, nBytes, nChunks)

	log.Println("###### Filter FileOutput Content ######")
	filter.FileOutputFilter(jobIdFileName)

	log.Println("###### Uploading File to S3 ######")
	export.UploadFile(jobIdFileName, "vandv-common-store", bucketKey)
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
	log.Println("###### Processing output file ######")

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
			log.Fatal(err)
		}

		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
	}
	log.Println("Bytes:", nBytes, "Chunks:", nChunks)
}

func closeFile(fileToClose *os.File) {
	log.Printf("###### Closing file %s ######", fileToClose.Name())

	if err := fileToClose.Close(); err != nil {
		log.Fatal(err)
	}
}