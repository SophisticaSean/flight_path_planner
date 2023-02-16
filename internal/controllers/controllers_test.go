package controllers_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/SophisticaSean/flight_path_calculator/internal/controllers"
	"github.com/SophisticaSean/flight_path_calculator/internal/models"
	"github.com/tj/assert"
)

func TestCalculateInvalidJSON(t *testing.T) {
	// this tells go that this test can run in Parallel
	// with other t.parallel enabled unit tests
	t.Parallel()

	body := strings.NewReader("hello")
	req := httptest.NewRequest(http.MethodPost, "/calculate", body)
	w := httptest.NewRecorder()

	// handle the request
	controllers.CalculateHandler(w, req)

	response := w.Result()
	defer response.Body.Close()

	readBody, err := io.ReadAll(response.Body)
	if err != nil {
		t.Errorf("unable to read response body")
	}

	strBody := string(readBody)
	assert.Contains(t, strBody, "Request body is not valid")
}

func TestCalculateInvalidFlightInput(t *testing.T) {
	// this tells go that this test can run in Parallel
	// with other t.parallel enabled unit tests
	t.Parallel()

	body := strings.NewReader(`[["SLC", "JFK"], ["SLC", "SFO"]]`)
	req := httptest.NewRequest(http.MethodPost, "/calculate", body)
	w := httptest.NewRecorder()

	// handle the request
	controllers.CalculateHandler(w, req)

	response := w.Result()
	defer response.Body.Close()

	readBody, err := io.ReadAll(response.Body)
	if err != nil {
		t.Errorf("unable to read response body")
	}

	strBody := string(readBody)
	assert.Contains(t, strBody, "Departure airport SLC appears more than once")
}

func TestCalculate(t *testing.T) {
	// this tells go that this test can run in Parallel
	// with other t.parallel enabled unit tests
	t.Parallel()

	body := strings.NewReader(`[["SLC", "JFK"], ["JFK", "SFO"], ["SFO", "ABS"]]`)
	req := httptest.NewRequest(http.MethodPost, "/calculate", body)
	w := httptest.NewRecorder()

	// handle the request
	controllers.CalculateHandler(w, req)

	response := w.Result()
	defer response.Body.Close()

	readBody, err := io.ReadAll(response.Body)
	if err != nil {
		t.Errorf("unable to read response body")
	}

	// deserialize the body back into a models.FlightOutput
	flightOutput := models.FlightOutput{}

	err = json.Unmarshal(readBody, &flightOutput)
	if err != nil {
		t.Errorf("unable to unmarshal response body")
	}

	arrivalAirport := "ABS"
	departureAirport := "SLC"

	// should not have an error
	assert.Empty(t, flightOutput.ErrorInformation)

	// ensure departureAirport and arrivalAirport is correct
	assert.Equal(t, []string{departureAirport, arrivalAirport}, flightOutput.CalculateResult)
	assert.Equal(t, departureAirport, flightOutput.FinalDepartureAirport)
	assert.Equal(t, arrivalAirport, flightOutput.FinalArrivalAirport)

	// ensure our path is correct
	assert.Equal(t, "SLC - JFK - SFO - ABS", flightOutput.Path)
}
