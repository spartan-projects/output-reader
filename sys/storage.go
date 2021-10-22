package sys

import (
	"github.com/spartan-projects/output-reader/common"
	"io/ioutil"
	"log"
)

func ReadFile(fileName string) string {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Panic(err.Error())
	}

	return string(file)
}

func WriteFile(fileName string, content string) {
	err := ioutil.WriteFile(fileName, []byte(content) , common.FilePermissions)

	if err != nil {
		log.Panic(err.Error())
	}
}