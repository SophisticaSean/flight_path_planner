# flight_path_calculator

## What is this project?
  This project is able to take HTTP POST input such as `[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]` on the endpoint `/calculate` and returns the flight path `SFO -> ATL -> GSO -> IND -> EWR` and initial departure and final arrival airport `["SFO", "EWR"]`.

## Continuous Integration Configuration
  Github Actions is configured on this repo for 3 jobs:
  - build: this will run `go build ./...` and ensure it passes.
  - test: this will run `go test ./...` and ensure it passes.
  - lint: this will run `golangci-lint -v --color=always run ./...` and ensure it passes.

  To view actions in this repo go [here](https://github.com/SophisticaSean/flight_path_planner/actions).

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

  ### Code Coverage
  To compute the above code coverage locally, use `go test --cover ./...`

  With 17 unit tests, we have the following coverage:
  ```bash
    go test --cover ./...
    ?       github.com/SophisticaSean/flight_path_calculator        [no test files]
    ok      github.com/SophisticaSean/flight_path_calculator/internal/controllers   (cached)        coverage: 72.0% of statements
    ok      github.com/SophisticaSean/flight_path_calculator/internal/models        (cached)        coverage: 94.2% of statements
  ```
  - models: 
  94.2% coverage on models is satisfactory, the uncovered code is difficult to induce a failure for.
  - controllers:
  To increase coverage on controllers, we'd have to implement an interface around models/flightoutputs to induce json serialization errors.
  The effort to get that increased coverage isn't always worth the extra work it requires, so I'm happy with the coverage on controllers as it stands.

