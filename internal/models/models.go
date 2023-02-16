package models

type FlightsInput [][]string
type FlightOutput struct {
	CalculateResult       []string
	FinalDepartureAirport string
	FinalArrivalAirport   string
	Path                  string
	ErrorInformation      string
}
