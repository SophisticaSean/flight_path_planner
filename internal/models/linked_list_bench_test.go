package models_test

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"

	"github.com/SophisticaSean/flight_path_calculator/internal/models"
	"github.com/stretchr/testify/assert"
)

func BenchmarkLinkedList10Flights(b *testing.B) {
	fi := models.FlightsInput{}

	inputData := `[
  ["IND", "EWR"], 
  ["EWR", "AAA"],
  ["AAA", "AAB"],
  ["AAB", "AAC"],
  ["AAC", "AAD"],
  ["AAD", "AAE"],
  ["AAE", "AAF"],
  ["AAF", "AAG"],
  ["AAG", "AAH"],
  ["AAH", "AAI"]
  ]`

	err := json.Unmarshal([]byte(inputData), &fi)
	if err != nil {
		b.Fail()
	}

	flightOutput := fi.FindStartAndEndFlightLinkedList()

	for n := 0; n < b.N; n++ {
		flightOutput = fi.FindStartAndEndFlightLinkedList()
	}

	// should be no errors
	assert.Equal(b, flightOutput.ErrorInformation, "")
	assert.Equal(b, flightOutput.CalculateResult, []string{"IND", "AAI"})
	assert.Equal(b, "IND - EWR - AAA - AAB - AAC - AAD - AAE - AAF - AAG - AAH - AAI", flightOutput.Path)
}

func BenchmarkLinkedList100Flights(b *testing.B) {
	fi, solution, path := generateRandomFlightPath(1000)

	flightOutput := fi.FindStartAndEndFlightLinkedList()

	for n := 0; n < b.N; n++ {
		flightOutput = fi.FindStartAndEndFlightLinkedList()
	}

	// should be no errors
	assert.Equal(b, flightOutput.ErrorInformation, "")
	assert.Equal(b, flightOutput.CalculateResult, solution)
	assert.Equal(b, flightOutput.Path, path)
}

func BenchmarkLinkedList1000Flights(b *testing.B) {
	fi, solution, path := generateRandomFlightPath(1000)

	flightOutput := fi.FindStartAndEndFlightLinkedList()

	for n := 0; n < b.N; n++ {
		flightOutput = fi.FindStartAndEndFlightLinkedList()
	}

	// should be no errors
	assert.Equal(b, flightOutput.ErrorInformation, "")
	assert.Equal(b, flightOutput.CalculateResult, solution)
	assert.Equal(b, flightOutput.Path, path)
}

func BenchmarkNaiveImpl10Flights(b *testing.B) {
	fi := models.FlightsInput{}

	inputData := `[
  ["IND", "EWR"], 
  ["EWR", "AAA"],
  ["AAA", "AAB"],
  ["AAB", "AAC"],
  ["AAC", "AAD"],
  ["AAD", "AAE"],
  ["AAE", "AAF"],
  ["AAF", "AAG"],
  ["AAG", "AAH"],
  ["AAH", "AAI"]
  ]`

	err := json.Unmarshal([]byte(inputData), &fi)
	if err != nil {
		b.Fail()
	}

	flightOutput := fi.FindStartAndEndFlightNaive()

	for n := 0; n < b.N; n++ {
		flightOutput = fi.FindStartAndEndFlightNaive()
	}

	// should be no errors
	assert.Equal(b, "", flightOutput.ErrorInformation)
	assert.Equal(b, flightOutput.CalculateResult, []string{"IND", "AAI"})
	// path unimplemented for naive impl
	// assert.Equal(b, "SFO - ATL - GSO - IND - EWR - AAA - AAB - AAC - AAD - AAE - AAF - AAG - AAH - AAI - AAJ - AAK - AAL - AAM", flightOutput.Path)
}

func BenchmarkNaive100Flights(b *testing.B) {
	fi, solution, _ := generateRandomFlightPath(100)

	flightOutput := fi.FindStartAndEndFlightNaive()

	for n := 0; n < b.N; n++ {
		flightOutput = fi.FindStartAndEndFlightNaive()
	}

	// should be no errors
	assert.Equal(b, "", flightOutput.ErrorInformation)
	assert.Equal(b, flightOutput.CalculateResult, solution)
}
func BenchmarkNaive1000Flights(b *testing.B) {
	fi, solution, _ := generateRandomFlightPath(1000)

	flightOutput := fi.FindStartAndEndFlightNaive()

	for n := 0; n < b.N; n++ {
		flightOutput = fi.FindStartAndEndFlightNaive()
	}

	// should be no errors
	assert.Equal(b, "", flightOutput.ErrorInformation)
	assert.Equal(b, flightOutput.CalculateResult, solution)
	// path unimplemented for naive impl
}

func generateRandomFlightPath(length int) (flights models.FlightsInput, solution []string, flightPath string) {
	currItem := ""
	// always start on AAA
	lastItem := "AAA"
	firstItem := ""
	orderedFlightSlice := []string{lastItem}
	// keep track of values we've already seen
	seen := make(map[string]string)
	for i := 0; i < length; i++ {
		// attempt 100 times to generate a non-colliding airport like SFO"
		for j := 0; j < 100; j++ {
			currItem = randSeq(3)
			_, ok := seen[currItem]
			if !ok {
				seen[currItem] = ""
				break
			}
		}
		// set first path if we're just starting the loop
		if i == 0 {
			firstItem = lastItem
		}

		orderedFlightSlice = append(orderedFlightSlice, currItem)
		calcResult := []string{lastItem, currItem}
		flights = append(flights, calcResult)
		lastItem = currItem
	}

	// set final item after loop exits
	finalItem := currItem

	solution = []string{firstItem, finalItem}

	flightPath = genTestFlightPath(orderedFlightSlice)
	return flights, solution, flightPath
}

func genTestFlightPath(flights []string) (output string) {
	for i, flight := range flights {
		if i == 0 {
			output = flight
		} else {
			output = fmt.Sprintf("%s - %s", output, flight)
		}
	}

	return output
}

func randSeq(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
