package models_test

import (
	"encoding/json"
	"testing"

	"github.com/SophisticaSean/flight_path_calculator/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestDeserializeSimpleInput(t *testing.T) {
	// this tells go that this test can run in Parallel
	// with other t.parallel enabled unit tests
	t.Parallel()

	sample := `[["SFO", "EWR"]]`
	fi := models.FlightsInput{}

	err := json.Unmarshal([]byte(sample), &fi)
	// ensure no Unmarshal error
	assert.Nil(t, err)
	// check first items
	assert.Equal(t, fi[0][0], "SFO")
	assert.Equal(t, fi[0][1], "EWR")
}

func TestDeserializeComplexInput(t *testing.T) {
	// this tells go that this test can run in Parallel
	// with other t.parallel enabled unit tests
	t.Parallel()

	sample := `[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]`
	fi := models.FlightsInput{}

	err := json.Unmarshal([]byte(sample), &fi)
	// ensure no Unmarshal error
	assert.Nil(t, err)
	// check first items
	assert.Equal(t, fi[0][0], "IND")
	assert.Equal(t, fi[0][1], "EWR")

	// check last items
	assert.Equal(t, fi[3][0], "ATL")
	assert.Equal(t, fi[3][1], "GSO")
}

func TestNaiveFindStartAndEndFlightSimple(t *testing.T) {
	// this tells go that this test can run in Parallel
	// with other t.parallel enabled unit tests
	t.Parallel()

	sample := `[["SFO","SLC"],["ATL","JFK"],["PHX","ATL"],["SLC","PHX"]]`
	fi := models.FlightsInput{}

	err := json.Unmarshal([]byte(sample), &fi)
	// ensure no Unmarshal error
	assert.Nil(t, err)

	flightOutput := fi.FindStartAndEndFlightNaive()
	// should be no errors
	assert.Equal(t, flightOutput.ErrorInformation, "")

	assert.Equal(t, flightOutput.CalculateResult, []string{"SFO", "JFK"})
}

func TestNaiveFindStartAndEndFlightComplex(t *testing.T) {
	// this tells go that this test can run in Parallel
	// with other t.parallel enabled unit tests
	t.Parallel()

	sample := `[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]`
	fi := models.FlightsInput{}

	err := json.Unmarshal([]byte(sample), &fi)
	// ensure no Unmarshal error
	assert.Nil(t, err)

	flightOutput := fi.FindStartAndEndFlightNaive()
	// should be no errors
	assert.Equal(t, flightOutput.ErrorInformation, "")

	assert.Equal(t, flightOutput.CalculateResult, []string{"SFO", "EWR"})
}

func TestNaiveFindStartAndEndFlightFlightLoop(t *testing.T) {
	// this tells go that this test can run in Parallel
	// with other t.parallel enabled unit tests
	t.Parallel()

	// loop in starting flight IND -> ATL
	sample := `[["IND", "ATL"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]`
	fi := models.FlightsInput{}

	err := json.Unmarshal([]byte(sample), &fi)
	// ensure no Unmarshal error
	assert.Nil(t, err)

	flightOutput := fi.FindStartAndEndFlightNaive()
	// should error out
	assert.Contains(t, flightOutput.ErrorInformation, "Unable to find")
}

func TestNaiveFindStartAndEndFlightExtraArrival(t *testing.T) {
	// this tells go that this test can run in Parallel
	// with other t.parallel enabled unit tests
	t.Parallel()

	// this test inspired me to do the linked list implementation
	// as I didn't think my naive solution would be able
	// detect and diagnose this test case without a lot of extra work
	// I also wanted to be able to adequately show the flight plan
	// and this implementation also made that impossible.
	t.Skip()

	// GSO has two entries
	sample := `[["IND", "ATL"], ["SFO", "ATL"], ["GSO", "IND"], ["GSO", "JFK"], ["ATL", "GSO"]]`
	fi := models.FlightsInput{}

	err := json.Unmarshal([]byte(sample), &fi)
	// ensure no Unmarshal error
	assert.Nil(t, err)

	flightOutput := fi.FindStartAndEndFlightNaive()
	// should error out
	assert.NotEmpty(t, flightOutput.ErrorInformation)
}
