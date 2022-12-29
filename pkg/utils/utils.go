package utils

import (
	"errors"
	"log"
	"strconv"
	"strings"
)

func CheckErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

var ErrAlreadyExists = errors.New("record already exists")

func ConvertToInt(array []string) []int {
	var result = []int{}
	for _, str := range array {
		num, err := strconv.Atoi(str)
		CheckErr(err)

		result = append(result, num)
	}

	return result
}

func DateFromTimeStamp(timestamp string) string {
	return strings.Split(timestamp, " ")[0]
}
