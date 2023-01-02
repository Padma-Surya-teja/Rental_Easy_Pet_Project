package utils

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
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

// use godot package to load/read the .env file and
// return the value of the key
func GoDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(filepath.Join(`/home/suryatejap/Documents/Go Lang Practice/RENTAL EASY V-1 (copy)/`, ".env"))
	CheckErr(err)

	return os.Getenv(key)
}
