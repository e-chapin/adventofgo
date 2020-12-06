package adventofgo

import (
	"io/ioutil"
	"os"
	"strings"
)

func ReadFile(name string) []string{
	if name == "" {
		name = "input.txt"
	}
	content, err := ioutil.ReadFile(name)
	if err != nil {
		os.Exit(0)
	}
	return strings.Split(string(content), "\n")
}


//for index, line := range lines {
//
//if line == ""{
//// check for completed passport
//if checkPassportPartOne(currentPassport){
//validCountPartOne += 1
//}
//if checkPassportPartTwo(currentPassport){
//validCountPartTwo += 1
//}
//currentPassport = make(map[string]string)
//continue
//	}
