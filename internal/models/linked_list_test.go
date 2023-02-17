package models_test

import (
	"encoding/json"
	"testing"

	"github.com/SophisticaSean/flight_path_calculator/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestLinkedListComplex(t *testing.T) {
	// this tells go that this test can run in Parallel
	// with other t.parallel enabled unit tests
	t.Parallel()

	fi := models.FlightsInput{}

	inputData := `[
  ["IND", "EWR"], 
  ["SFO", "ATL"], 
  ["GSO", "IND"], 
  ["ATL", "GSO"]
  ]`
	err := json.Unmarshal([]byte(inputData), &fi)
	// ensure no Unmarshal error
	assert.Nil(t, err)

	flightOutput := fi.FindStartAndEndFlightLinkedList()
	// should be no errors
	assert.Equal(t, "", flightOutput.ErrorInformation)
	assert.Equal(t, []string{"SFO", "EWR"}, flightOutput.CalculateResult)
	assert.Equal(t, "SFO - ATL - GSO - IND - EWR", flightOutput.Path)
}

func TestLinkedListHuge(t *testing.T) {
	// this tells go that this test can run in Parallel
	// with other t.parallel enabled unit tests
	t.Parallel()

	inputData := `[
  ["IND", "EWR"], 
  ["SFO", "ATL"], 
  ["ABS", "IND"], 
  ["ATL", "GSO"], 
  ["GSO", "ABS"], 
  ["EWR", "ABC"]
  ]`
	fi := models.FlightsInput{}

	err := json.Unmarshal([]byte(inputData), &fi)
	// ensure no Unmarshal error
	assert.Nil(t, err)

	flightOutput := fi.FindStartAndEndFlightLinkedList()
	// should be no errors
	assert.Equal(t, "", flightOutput.ErrorInformation)
	// assert.Equal(t, flightOutput.RawOutput, []string{"SFO", "EWR"})
	assert.Equal(t, "SFO - ATL - GSO - ABS - IND - EWR - ABC", flightOutput.Path)
}

func TestLinkedListLoop(t *testing.T) {
	// this tells go that this test can run in Parallel
	// with other t.parallel enabled unit tests
	t.Parallel()

	// ["EWR", "SFO"] creates a complete loop
	inputData := `[
  ["IND", "EWR"], 
  ["SFO", "ATL"], 
  ["GSO", "IND"], 
  ["ATL", "GSO"], 
  ["EWR", "SFO"]
  ]`
	fi := models.FlightsInput{}

	err := json.Unmarshal([]byte(inputData), &fi)
	// ensure no Unmarshal error
	assert.Nil(t, err)

	flightOutput := fi.FindStartAndEndFlightLinkedList()
	// should be no errors
	assert.Contains(t, flightOutput.ErrorInformation, "Duplicates found in")
}

func TestLinkedListPartialLoop(t *testing.T) {
	// this tells go that this test can run in Parallel
	// with other t.parallel enabled unit tests
	t.Parallel()

	// ["EWR", "ATL"] creates a partial loop
	inputData := `[
  ["IND", "EWR"], 
  ["SFO", "ATL"], 
  ["GSO", "IND"], 
  ["ATL", "GSO"], 
  ["EWR", "ATL"]
  ]`
	fi := models.FlightsInput{}

	err := json.Unmarshal([]byte(inputData), &fi)
	// ensure no Unmarshal error
	assert.Nil(t, err)

	flightOutput := fi.FindStartAndEndFlightLinkedList()
	// should be no errors
	assert.Contains(t, flightOutput.ErrorInformation, "Arrival airport ATL")
}

func TestLinkedListUnconnectedPaths(t *testing.T) {
	// this tells go that this test can run in Parallel
	// with other t.parallel enabled unit tests
	t.Parallel()

	// ["SLC", "JFK"] is completely isolated
	inputData := `[
  ["IND", "EWR"], 
  ["SFO", "ATL"], 
  ["GSO", "IND"], 
  ["ATL", "GSO"], 
  ["SLC", "JFK"]
  ]`
	fi := models.FlightsInput{}

	err := json.Unmarshal([]byte(inputData), &fi)
	// ensure no Unmarshal error
	assert.Nil(t, err)

	flightOutput := fi.FindStartAndEndFlightLinkedList()
	// should be no errors
	assert.Contains(t, flightOutput.ErrorInformation, "Unable to find a connecting path")
}

func TestLinkedListDoubleArrivals(t *testing.T) {
	// this tells go that this test can run in Parallel
	// with other t.parallel enabled unit tests
	t.Parallel()

	// ATL has two different arrival entries
	inputData := `[
  ["IND", "EWR"], 
  ["SFO", "ATL"], 
  ["ATL", "SLC"], 
  ["GSO", "IND"], 
  ["ATL", "GSO"]
  ]`
	fi := models.FlightsInput{}

	err := json.Unmarshal([]byte(inputData), &fi)
	// ensure no Unmarshal error
	assert.Nil(t, err)

	flightOutput := fi.FindStartAndEndFlightLinkedList()
	// should have an error
	assert.Contains(t, flightOutput.ErrorInformation, "Departure airport ATL")
}

func TestLinkedListDoubleDepartures(t *testing.T) {
	// this tells go that this test can run in Parallel
	// with other t.parallel enabled unit tests
	t.Parallel()

	// IND has two different arrival entries
	inputData := `[
  ["IND", "EWR"], 
  ["SFO", "ATL"], 
  ["ATL", "SLC"], 
  ["GSO", "IND"], 
  ["JFK", "IND"]
  ]`
	fi := models.FlightsInput{}

	err := json.Unmarshal([]byte(inputData), &fi)
	// ensure no Unmarshal error
	assert.Nil(t, err)

	flightOutput := fi.FindStartAndEndFlightLinkedList()
	// should have an error
	assert.Contains(t, flightOutput.ErrorInformation, "Arrival airport IND")
}

func TestLinkedListInvalidList(t *testing.T) {
	// this tells go that this test can run in Parallel
	// with other t.parallel enabled unit tests
	t.Parallel()

	// IND is missing a 2nd item
	inputData := `[
  ["IND"], 
  ["SFO", "ATL"], 
  ["ATL", "SLC"], 
  ["GSO", "IND"]
  ]`
	fi := models.FlightsInput{}

	err := json.Unmarshal([]byte(inputData), &fi)
	// ensure no Unmarshal error
	assert.Nil(t, err)

	flightOutput := fi.FindStartAndEndFlightLinkedList()
	// should have an error
	assert.Contains(t, flightOutput.ErrorInformation, "Item [IND] does not have exactly two airports.")
}
