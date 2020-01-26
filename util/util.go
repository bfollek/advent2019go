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

// MustLoadStringSlice splits the result of MustLoadString() on `sep`.
// "If sep is empty, Split splits after each UTF-8 sequence." -
// https://golang.org/pkg/strings/#Split
func MustLoadStringSlice(fileName, sep string) []string {
	s := MustLoadString(fileName)
	return strings.Split(s, sep)
}

// MustLoadIntSlice splits the result of MustLoadString() on `splitString`,
// and converts each string digit to an int.
func MustLoadIntSlice(fileName, sep string) []int {
	is := []int{}
	ss := MustLoadStringSlice(fileName, sep)
	for _, s := range ss {
		is = append(is, MustAtoi(s))
	}
	return is
}

// MustAtoi converts a string to an integer and stops program execution
// if there's an error.
func MustAtoi(s string) int {
	i, err := strconv.Atoi(strings.TrimSpace(s))
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
