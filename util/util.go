package util

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
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

// MustReadLines reads the lines of a text file into a string slice.
func MustReadLines(fileName string) []string {
	absPath, err := filepath.Abs(fileName)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

// AbsInt64 returns the absolute value of an int64.
func AbsInt64(i int64) int64 {
	if i < 0 {
		return -i
	}
	return i
}

// CharToIntValue converts an integer char to its integer value,
// e.g. '3' => 3.
func CharToIntValue(char byte) int {
	return int(char - '0')
}
