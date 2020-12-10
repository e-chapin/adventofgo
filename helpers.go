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

func Contains(item int, list []int) bool{
	for _, value := range list {
		if value == item {
			return true
		}
	}
	return false
}