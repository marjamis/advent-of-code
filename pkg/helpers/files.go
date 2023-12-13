package helpers

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// ReadIntArray reads from file and returns an []int
func ReadIntArray(file string) (ints []int) {
	fdata, _ := os.Open(file)
	defer fdata.Close()

	scanner := bufio.NewScanner(fdata)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		ints = append(ints, i)
	}

	return
}

// ReadStringArray reads from file and returns a []string
func ReadStringArray(file string) (strings []string) {
	fdata, _ := os.Open(file)
	defer fdata.Close()

	scanner := bufio.NewScanner(fdata)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		strings = append(strings, scanner.Text())
	}

	return
}

// ReadString reads from file and returns a string
func ReadString(file string) string {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("File reading error", err)
		return ""
	}

	return strings.TrimSuffix(string(data), "\n")
}

// ReadStringArray2d reads from file and returns a [][]string
func ReadStringArray2d(file string) (strings [][]string) {
	fdata, _ := os.Open(file)
	defer fdata.Close()

	scanner := bufio.NewScanner(fdata)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		substrings := []string{}
		for _, c := range scanner.Text() {
			substrings = append(substrings, string(c))
		}
		strings = append(strings, substrings)
	}

	return
}

// ReadRuneArray2d  reads from file and returns a [][]rune
func ReadRuneArray2d(file string) (strings [][]rune) {
	fdata, _ := os.Open(file)
	defer fdata.Close()

	scanner := bufio.NewScanner(fdata)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		substrings := []rune{}
		for _, c := range scanner.Text() {
			substrings = append(substrings, c)
		}
		strings = append(strings, substrings)
	}

	return
}

// ReadIntCSV reads a csv file and return a 2d int array
func ReadIntCSV(file string) (ints [][]int) {
	fdata, _ := os.Open(file)
	defer fdata.Close()

	r := csv.NewReader(fdata)

	var recordCount int
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		var newLine []int
		for _, value := range record {
			d, _ := strconv.Atoi(value)
			newLine = append(newLine, d)
		}
		ints = append(ints, newLine)

		recordCount++
	}

	return
}

// ReadStringCSV reads a csv file and return a 2d int array
func ReadStringCSV(file string) (strings [][]string) {
	fdata, _ := os.Open(file)
	defer fdata.Close()

	r := csv.NewReader(fdata)

	var recordCount int
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		var newLine []string
		for _, value := range record {
			newLine = append(newLine, value)
		}
		strings = append(strings, newLine)

		recordCount++
	}

	return
}

// ReadIntArray2d reads from file and returns a [][]int
func ReadIntArray2d(file string) (ints [][]int) {
	fdata, _ := os.Open(file)
	defer fdata.Close()

	scanner := bufio.NewScanner(fdata)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		subInts := []int{}
		for _, c := range scanner.Text() {
			tmpI, err := strconv.Atoi(string(c))
			if err != nil {
				log.Fatal("Couldn't convert file")
			}
			subInts = append(subInts, tmpI)
		}
		ints = append(ints, subInts)
	}

	return
}
