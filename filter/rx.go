package filter

import (
	"fmt"
	"github.com/spartan-projects/output-reader/common"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func FileOutputFilter(jobIdFileName string) {
	r := regexp.MustCompile(common.OutputFilterRegex)

	file, err := ioutil.ReadFile(jobIdFileName)
	if err != nil {
		log.Panic(err.Error())
	}

	fileContent := string(file)

	if r.MatchString(fileContent) {
		findString := r.FindString(fileContent)
		fmt.Printf("##### MATCH FOUND #####\n %s", findString)
		ioutil.WriteFile(jobIdFileName, []byte(findString) , common.FilePermissions)
	}
}

func EofFilter(content string) bool{
	var found = false

	if strings.Contains(content, common.EofFilter) {
		log.Println("###### EOF Filter Matches!!! ######")
		found = true
	}

	return found
}
