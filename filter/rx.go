package filter

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
)

var regexPattern = `(?s)(vxTest\sOptions\sSummary).*?(Test\sexecution\sfinished)`

func FileOutputFilter(jobIdFileName string) {
	r := regexp.MustCompile(regexPattern)

	file, err := ioutil.ReadFile(jobIdFileName)
	if err != nil {
		log.Panic(err.Error())
	}

	matches := r.FindAllString(string(file), -1)

	for _, match := range matches {
		fmt.Printf("##### MATCH FOUND #####\n %s", match)
		ioutil.WriteFile(jobIdFileName, []byte(match) , 0666)
	}
}
