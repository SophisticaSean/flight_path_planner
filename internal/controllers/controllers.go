package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/SophisticaSean/flight_path_calculator/internal/models"
)

// CalculateHandler is the controller for the /calculate endpoint
// controllers should be named similarly to the routes they serve
func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	flightInput := models.FlightsInput{}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Unable to read request body")
		return
	}

	err = json.Unmarshal(body, &flightInput)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `Request body is not valid. Valid input would be: '[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]'`)
		return
	}

	flightOutput := flightInput.FindStartAndEndFlightLinkedList()
	fmt.Println(flightOutput.Path)
	fmt.Printf("First departure: %s\n", flightOutput.FinalDepartureAirport)
	fmt.Printf("Last arrival: %s\n", flightOutput.FinalArrivalAirport)

	// return just the error string on an error case
	if flightOutput.ErrorInformation != "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, flightOutput.ErrorInformation)
		return
	}

	jsonOut, err := json.Marshal(flightOutput)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `Unable to serialize flightOutput JSON, please contact support.`)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonOut)
	if err != nil {
		panic("unable to write out JSON to client")
	}
}
