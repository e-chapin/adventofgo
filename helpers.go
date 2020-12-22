package adventofgo

import (
	"io/ioutil"
	"os"
	"strconv"
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

func GetFileString(name string) string {
	if name == "" {
		name = "input.txt"
	}
	content, err := ioutil.ReadFile(name)
	if err != nil {
		os.Exit(0)
	}
	return string(content)
}

func Contains(item int, list []int) bool{
	for _, value := range list {
		if value == item {
			return true
		}
	}
	return false
}


func IndexOf(item string, list []string)  int {
	for index, value := range list {
		if value == item {
			return index
		}
	}
	return -1
}

func IndexOfInt(item int, list []int)  int {
	for index, value := range list {
		if value == item {
			return index
		}
	}
	return -1
}
func IndexOfInt64(item int64, list []int64)  int {
	for index, value := range list {
		if value == item {
			return index
		}
	}
	return -1
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ReverseString(s string) (result string) {
	for _,v := range s {
		result = string(v) + result
	}
	return
}

func CheckError (err error) {
	if err != nil {
		panic(err)
	}

}

func BinaryToInt(s string) int64 {
	i, err := strconv.ParseInt(s, 2, 64)
	CheckError(err)
	return i

}

func AsInt(s string) int {
	i, err := strconv.Atoi(s)
	CheckError(err)
	return i
}

func RemoveEmpty(s []string) []string {
	var r []string
	for _, v := range s {
		if v != "" {
			r = append(r, v)
		}
	}
	return r
}

func Unique(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}