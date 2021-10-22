package filter

import (
	"fmt"
	"github.com/spartan-projects/output-reader/common"
	"log"
	"regexp"
	"strings"
)

func FileOutputFilter(fileContent string) (bool, string){
	var match = false
	var findString string
	r := regexp.MustCompile(common.OutputFilterRegex)

	if r.MatchString(fileContent) {
		findString = r.FindString(fileContent)
		fmt.Printf("##### MATCH FOUND #####\n %s", findString)
		match = true
	}

	return match, findString
}

func EofFilter(content string) bool{
	var found = false

	if strings.Contains(content, common.EofFilter) {
		log.Println("###### EOF Filter Matches!!! ######")
		found = true
	}

	return found
}
