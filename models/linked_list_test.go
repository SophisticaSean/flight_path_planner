package models_test

import (
	"encoding/json"
	"testing"

	"github.com/SophisticaSean/flight_path_calculator/models"
	"github.com/stretchr/testify/assert"
)

func TestLinkedListComplex(t *testing.T) {
	// this tells go that this test can run in Parallel
	// with other t.parallel enabled unit tests
	t.Parallel()

	sample := `[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]`
	fi := models.FlightsInput{}

	err := json.Unmarshal([]byte(sample), &fi)
	// ensure no Unmarshal error
	assert.Nil(t, err)

	flightOutput := fi.FindStartAndEndFlightLinkedList()
	// should be no errors
	assert.Equal(t, flightOutput.ErrorInformation, "")
	assert.Equal(t, flightOutput.RawOutput, []string{"SFO", "EWR"})
	assert.Equal(t, flightOutput.Path, "SFO -> ATL -> GSO -> IND -> EWR")
}

func TestLinkedListLoop(t *testing.T) {
	// this tells go that this test can run in Parallel
	// with other t.parallel enabled unit tests
	t.Parallel()

	// ["EWR", "SFO"] creates a complete loop
	sample := `[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"], ["EWR", "SFO"]]`
	fi := models.FlightsInput{}

	err := json.Unmarshal([]byte(sample), &fi)
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
	sample := `[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"], ["EWR", "ATL"]]`
	fi := models.FlightsInput{}

	err := json.Unmarshal([]byte(sample), &fi)
	// ensure no Unmarshal error
	assert.Nil(t, err)

	flightOutput := fi.FindStartAndEndFlightLinkedList()
	// should be no errors
	assert.Contains(t, flightOutput.ErrorInformation, "Duplicates found in")
}
