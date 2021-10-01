package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

var namedPipeFile = "/home/valkyrie/vxworks-images/workspace_io/vip_intel_test/comms/input.out"

func main() {
	nBytes, nChunks := int64(0), int64(0)
	file, err := os.OpenFile(namedPipeFile, os.O_RDONLY, os.ModeNamedPipe)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	// TODO change output file name with job id - example ebf0001
	fo, err := os.OpenFile("ebf0001.log", os.O_CREATE | os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	r := bufio.NewReader(file)
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
		// process buf
		if _, err := fo.Write(buf[:n]); err != nil {
			panic(err)
		}

		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
	}
	log.Println("Bytes:", nBytes, "Chunks:", nChunks)
}