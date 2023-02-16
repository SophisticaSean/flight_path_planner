# flight_path_calculator

## What is this project?
  This project is able to take input such as `[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]`
  and return the flight path `SFO -> ATL -> GSO -> IND -> EWR` and initial departure and final arrival airport `["SFO", "EWR"]`.

## How to use this project
  
  ### Requirements
  - go 1.19.5
  - [golangci-lint](https://golangci-lint.run/usage/install/)

  ### Testing
  `go test ./...`

  ### Linting
  `golangci-lint -v --color=always run ./...`

  ### Running the project
  - `go run ./...`
  - `curl -X POST "localhost:8080/calculate" -d '[["IND", "EWR"], ["EWR", "JFK"]]'`
