package filter

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
)

var regexPattern = `(?s)(vxTest\sOptions\sSummary).*?(Test\sexecution\sfinished)`
var eofPattern = `(?m)^Au\srevoir\!$`

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

func EofFilter(jobIdFileName string) {
	r := regexp.MustCompile(eofPattern)

	file, err := ioutil.ReadFile(jobIdFileName)
	if err != nil {
		log.Panic(err.Error())
	}

	fileContent := string(file)

	if r.MatchString(fileContent) {
		findString := r.FindString(fileContent)
		log.Printf("##### EOF MATCH FOUND #####\n %s", findString)
		log.Println("##### A TRUE BOOLEAN WILL BE RETURNED AT THIS POINT #####")
	} else {
		log.Printf("##### File Content #####\n%s", fileContent)
	}
}
