package models

type FlightsInput [][]string
type FlightOutput struct {
	RawOutput             []string
	FinalDepartureAirport string
	FinalArrivalAirport   string
	Path                  string
	ErrorInformation      string
}
