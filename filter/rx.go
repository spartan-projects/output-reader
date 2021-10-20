package filter

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

var regexPattern = `(?s)(vxTest\sOptions\sSummary).*?(Test\sexecution\sfinished)`
var eofPattern = "revoir"

func FileOutputFilter(jobIdFileName string) {
	r := regexp.MustCompile(regexPattern)

	file, err := ioutil.ReadFile(jobIdFileName)
	if err != nil {
		log.Panic(err.Error())
	}

	fileContent := string(file)

	if r.MatchString(fileContent) {
		findString := r.FindString(fileContent)
		fmt.Printf("##### MATCH FOUND #####\n %s", findString)
		ioutil.WriteFile(jobIdFileName, []byte(findString) , 0666)
	}
}

func EofFilter(content string) bool{
	var found = false

	if strings.Contains(content, eofPattern) {
		log.Println("###### EOF Filter Matches!!! ######")
		found = true
	}

	return found
}
