package models

import "fmt"

// FindStartAndEndFlightNaive was my first/initial solution to this problem
// It satisfied a decent chunk of test cases but I was unhappy that it
// wasn't able to show me the ending path from A -> B
// so then I wrote the linked_list implementation.
func (fi FlightsInput) FindStartAndEndFlightNaive() (fo FlightOutput) {
	startFlight := ""
	endFlight := ""

	// separate the flight list
	startList, endList, err := fi.splitFlightsInput()
	if err != nil {
		fo.ErrorInformation = err.Error()
		return fo
	}

	// find the first flight, it should not exist in the last flight list
	startFlight = findItemNotInSecondList(startList, endList)
	if startFlight == "" {
		fo.ErrorInformation = "Unable to find starting flight, loop or invalid list provided."
		return fo
	}

	// find the last flight, it should not exist in the first flight list
	endFlight = findItemNotInSecondList(endList, startList)
	if endFlight == "" {
		fo.ErrorInformation = "Unable to find ending flight, loop or invalid list provided."
		return fo
	}

	fo = FlightOutput{
		CalculateResult: []string{
			startFlight,
			endFlight,
		},
		FinalArrivalAirport:   startFlight,
		FinalDepartureAirport: endFlight,
		ErrorInformation:      "",
	}

	return fo
}

func (fi FlightsInput) splitFlightsInput() (startList, endList []string, err error) {
	for _, flightPair := range fi {
		// ensure all flightPairs are exactly 2 long
		if len(flightPair) != 2 {
			return startList, endList, fmt.Errorf("Item %v does not have exactly 2 airports.", flightPair)
		}
		for i, airport := range flightPair {
			if i == 0 {
				startList = append(startList, airport)
			} else {
				endList = append(endList, airport)
			}
		}
	}
	return startList, endList, nil
}

func findItemNotInSecondList(firstList, secondList []string) string {
	uniqueItem := ""
	for _, firstListI := range firstList {
		inSecondList := false
		for _, secondListI := range secondList {
			if secondListI == firstListI {
				inSecondList = true
				break // prevents unneeded loops
			}
		}
		if !inSecondList {
			uniqueItem = firstListI
		}
	}

	return uniqueItem
}
