package models

import (
	"container/list"
	"fmt"
)

func (fi FlightsInput) FindStartAndEndFlightLinkedList() (fo FlightOutput) {
	// ensure every FlightsInput has two items
	// also ensure arrivals and departures are unique
	arrivals := make(map[string]string)
	departures := make(map[string]string)
	for _, flightPair := range fi {
		// ensure all flightPairs are exactly 2 long
		if len(flightPair) != 2 {
			fo.ErrorInformation = fmt.Sprintf("Item %v does not have exactly 2 airports.", flightPair)
			return fo
		}

		// ensure departure is unique
		departure := flightPair[0]
		_, ok := departures[departure]
		if ok {
			fo.ErrorInformation = fmt.Sprintf("Departure airport %v appears more than once in the given flight plan.", departure)
			return fo
		}
		departures[departure] = ""

		// ensure arrival is unique
		arrival := flightPair[1]
		_, ok = arrivals[arrival]
		if ok {
			fo.ErrorInformation = fmt.Sprintf("Arrival airport %v appears more than once in the given flight plan.", arrival)
			return fo
		}
		arrivals[arrival] = ""
	}

	// using a map to keep track of which airport is in what location in the list
	itemMap := make(map[string]*list.Element)
	linkedList := list.New()

	newLL, newTrackingMap, notFound := buildFlightPath(linkedList, itemMap, fi)
	// recurse up to len(inputFlights)+1 times to try to populate the linked list
	solutionFound := false
	for i := 0; i < len(fi)*len(fi)+1; i++ {
		newLL, newTrackingMap, notFound = buildFlightPath(newLL, newTrackingMap, notFound)
		// break if notFound is empty
		if len(notFound) == 0 {
			solutionFound = true
			break
		}
	}

	if !solutionFound {
		fo.ErrorInformation = "Unable to find a connecting path for given flights."
		return fo
	}

	// check that the flight path is valid
	valid := validFlightPath(newLL)
	if !valid {
		fo.ErrorInformation = "Duplicates found in flight path. There's a complete or partial loop in given flight plan, or duplicate arrival/departure airports."
		return fo
	}

	path := concatLinkedList(newLL)
	fmt.Println(path)

	startFlight := newLL.Front().Value.(string)
	fmt.Printf("First departure: %s\n", startFlight)
	endFlight := newLL.Back().Value.(string)
	fmt.Printf("Last arrival: %s\n", endFlight)

	fo.FinalArrivalAirport = endFlight
	fo.FinalDepartureAirport = startFlight
	fo.RawOutput = []string{
		startFlight,
		endFlight,
	}
	fo.Path = path

	return fo
}

// pretty print a linked list
func printLinkedList(ll *list.List) {
	for e := ll.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

// turn a linked list from "['EWR', 'SFO', 'ATL']" -> "EWR -> SFO -> ATL"
func concatLinkedList(ll *list.List) (out string) {
	for e := ll.Front(); e != nil; e = e.Next() {
		if out == "" {
			out = fmt.Sprintf("%s", e.Value)
		} else {
			out = fmt.Sprintf("%s -> %s", out, e.Value)
		}
	}
	return out
}

// ensure an airport is only in the linked list once
func validFlightPath(ll *list.List) bool {
	foundItems := make(map[string]string)
	for e := ll.Front(); e != nil; e = e.Next() {
		item := e.Value.(string)
		_, ok := foundItems[item]
		if ok {
			// if there's duplicates, immediately return false
			return false
		} else {
			foundItems[item] = ""
		}
	}
	return true
}

func buildFlightPath(ll *list.List, inputMap map[string]*list.Element, fi FlightsInput) (newLL *list.List, newTrackingMap map[string]*list.Element, notFound FlightsInput) {
	for _, flightPair := range fi {
		departure := flightPair[0]
		arrival := flightPair[1]

		// if linkedList is empty, then populate it
		if ll.Len() == 0 {
			// push the first airport into the linked list
			mark := ll.PushFront(departure)
			inputMap[departure] = mark

			// push the last airport into the linked list
			mark2 := ll.PushBack(arrival)
			inputMap[arrival] = mark2
		} else {

			// handle departure find
			found := false
			mark, ok := inputMap[departure]
			if ok {
				// if found we know where the arrival needs to be inserted
				// first check if it's a duplicate
				if !isDuplicateOfParentOrChild(arrival, mark) {
					newMark := ll.InsertAfter(arrival, mark)
					inputMap[arrival] = newMark
					found = true
				}
			}

			// handle arrival find
			mark, ok = inputMap[arrival]
			if ok {
				// if found we know where the departure needs to be inserted
				// first check if it's a duplicate
				if !isDuplicateOfParentOrChild(departure, mark) {
					newMark := ll.InsertBefore(departure, mark)
					inputMap[departure] = newMark
					found = true
				}
			}

			// handle not found
			if !found {
				// if the item isn't in the map then put it into notFound
				notFound = append(notFound, flightPair)
			}
		}
	}

	return ll, inputMap, notFound
}

func isDuplicateOfParentOrChild(potentialAddition string, element *list.Element) bool {
	parent := ""
	child := ""
	ok := false

	// coerce element's parent to a string
	prev := element.Prev()
	if prev != nil {
		parent, ok = prev.Value.(string)
		if !ok {
			parent = ""
		}
	}

	// coerce element's child to a string
	next := element.Next()
	if next != nil {
		child, ok = next.Value.(string)
		if !ok {
			child = ""
		}
	}

	// check to see if either parent or child equal potentialAddition
	if parent == potentialAddition {
		return true
	}
	if child == potentialAddition {
		return true
	}
	return false
}
