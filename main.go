package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"mime/multipart"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jfyne/csvd"
)

// Digit represents a single digit (e.g. "1"), the number of times it's detected within the input CSV, and the calculated percentage.
type Digit struct {
	Value   int     `json:"digit"`
	Count   int     `json:"count"`
	Percent float64 `json:"percent"`
}

// Payload constitutes the JSON-encoded data that is output.
type Payload struct {
	Values            []Digit `json:"first_digit_distribution"`
	BenfordValidation bool    `json:"benford_match"`
	Filename          string  `json:"file"`
}

// AddItem appends a single Digit struct into the greater dataset, or Payload.
func (payload *Payload) AddItem(Value Digit) []Digit {
	payload.Values = append(payload.Values, Value)
	return payload.Values
}

// This returns a sorted list of indexes from the provided map, as maps in Golang are unordered.
// This is in case one of the leading digits isn't represented in the dataset.
// Also removes any zero key.
func sortMap(distributionMap map[int]int) []int {
	var keys = make([]int, 10)
	for k := range distributionMap {
		keys[k] = k
	}
	sort.Ints(keys)
	keys = removeIndex(keys, 0)
	return keys
}

// Returns an integer (effectively) representing the percentage of instances of a digit within
// the total number of numerical rows in the CSV file.
func calculatePercent(count int, totalRows int) (percent float64) {
	return math.Round(float64(count) / float64(totalRows) * 100)
}

// Converts the submitted column into zero-based numbering.
// TODO: Handle invalid inputs ("A") better.
func sanitizeColumnValue(column string) int {
	newValue, err := strconv.Atoi(column)
	if err != nil {
		newValue = 0
	}
	if newValue > 0 {
		newValue--
	} else if newValue < 0 {
		newValue = 0
	}
	return newValue
}

// We remove any leading characters that'd allow us to have a negative or numerical non-integer first digit.
func retrieveFirstDigit(record string) int {
	badCharacters := ".-0"
	value := strings.TrimLeft(record, badCharacters)
	firstDigit, _ := strconv.Atoi(string(value[0]))
	return firstDigit
}

// Returns a new slice object with the specified index removed from "slice".
func removeIndex(slice []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, slice[:index]...)
	return append(ret, slice[index+1:]...)
}

// Primary check on the application.
func benfordValidator(percent float64) bool {
	if percent >= 30 {
		return true
	}
	return false
}

// ReceiveFiles is the default HTTP handler for POST requests in this application. This
// handler expects a CSV file as its input and outputs a JSON-encoded description of the
// digit distribution per Benford's law.
func receiveFiles(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	r.ParseForm()
	var column int = sanitizeColumnValue(r.Form.Get("column"))
	var filename string = handler.Filename

	defer file.Close()

	data, err := parseCSV(file, filename, column)

	if err != nil {
		io.WriteString(w, "There was an error parsing the sent CSV file. Please try again.\n")
	} else {
		io.WriteString(w, data+"\n")
	}
}

// ParseCSV contains most of the logic: receiving a CSV file, recording the first digit for each
// row in a specific column, storing this data in an associative array, and then building the
// eventual JSON output.
func parseCSV(csvFile multipart.File, filename string, column int) (data string, err error) {
	var totalRows int = 0
	var distributionMap = make(map[int]int, 9)

	r := csvd.NewReader(csvFile)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		// Out of range error, i.e. that column is not in use.
		if len(record) <= column {
			break
		}
		// Skip over any empty records.
		if record[column] == "" {
			continue
		}

		firstDigit := retrieveFirstDigit(record[column])
		if firstDigit != 0 {
			distributionMap[firstDigit]++
			totalRows++
		}
	}

	sortedKeys := sortMap(distributionMap)

	payload := &Payload{}
	payload.Filename = filename

	for _, digit := range sortedKeys {
		var count int = distributionMap[digit]
		var percent float64 = calculatePercent(count, totalRows)

		if digit != 0 {
			values := Digit{Value: digit, Count: count, Percent: percent}
			payload.AddItem(values)
		}
		if digit == 1 {
			payload.BenfordValidation = benfordValidator(percent)
		}
	}

	output, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	data = string(output)
	return
}

func main() {
	pages := newPage()
	router := mux.NewRouter()
	router.
		Path("/").
		Methods("POST").
		HandlerFunc(receiveFiles)
	router.
		Handle("/", pages.Home).
		Methods("GET")
	fmt.Println("Starting server.")
	http.ListenAndServe(":9021", router)
}
