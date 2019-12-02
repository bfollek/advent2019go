package util

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
	"strings"
)

// LoadString loads a text file into a string.
func LoadString(fileName string) (string, error) {
	absPath, err := filepath.Abs(fileName)
	if err != nil {
		return "", err
	}
	bytes, err := ioutil.ReadFile(absPath)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// MustLoadString stops program execution if LoadString() returns an error.
func MustLoadString(fileName string) string {
	s, err := LoadString(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return s
}

// MustLoadStringSlice splits the result of MustLoadString() on `splitString`.
func MustLoadStringSlice(fileName, splitString string) []string {
	s := MustLoadString(fileName)
	return strings.Split(s, splitString)
}

// MustAtoi converts a string to an integer and stops program execution
// if there's an error.
func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
