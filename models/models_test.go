package models_test

import (
	"encoding/json"
	"testing"

	"github.com/SophisticaSean/flight_path_calculator/models"
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

func TestFindStartAndEndFlightSimple(t *testing.T) {
	// this tells go that this test can run in Parallel
	// with other t.parallel enabled unit tests
	t.Parallel()

	sample := `[["SFO","SLC"],["ATL","JFK"],["PHX","ATL"],["SLC","PHX"]]`
	fi := models.FlightsInput{}

	err := json.Unmarshal([]byte(sample), &fi)
	// ensure no Unmarshal error
	assert.Nil(t, err)

	flightOutput := fi.FindStartAndEndFlight()
	// should be no errors
	assert.Equal(t, flightOutput.ErrorInformation, "")

	assert.Equal(t, flightOutput.RawOutput, []string{"SFO", "JFK"})
}

func TestFindStartAndEndFlightComplex(t *testing.T) {
	// this tells go that this test can run in Parallel
	// with other t.parallel enabled unit tests
	t.Parallel()

	sample := `[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]`
	fi := models.FlightsInput{}

	err := json.Unmarshal([]byte(sample), &fi)
	// ensure no Unmarshal error
	assert.Nil(t, err)

	flightOutput := fi.FindStartAndEndFlight()
	// should be no errors
	assert.Equal(t, flightOutput.ErrorInformation, "")

	assert.Equal(t, flightOutput.RawOutput, []string{"SFO", "EWR"})
}

func TestFindStartAndEndFlightFlightLoop(t *testing.T) {
	// this tells go that this test can run in Parallel
	// with other t.parallel enabled unit tests
	t.Parallel()

	// loop in starting flight IND -> ATL
	sample := `[["IND", "ATL"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]`
	fi := models.FlightsInput{}

	err := json.Unmarshal([]byte(sample), &fi)
	// ensure no Unmarshal error
	assert.Nil(t, err)

	flightOutput := fi.FindStartAndEndFlight()
	// should error out
	assert.Contains(t, flightOutput.ErrorInformation, "Unable to find")
}
