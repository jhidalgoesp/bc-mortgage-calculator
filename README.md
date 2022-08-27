# BC-Mortgage-Calculator

This is an example of a Rest API application to calculate British Columbia mortgage schedules:

```bash
.
├── Makefile                    <-- Make to automate tasks
├── README.md                   <-- This instructions file
├── cmd                         <-- Application entrypoints
│   ├── main.go                 <-- Server
├── pkg                         <-- Library packages usable by the application
├── tests                       <-- Shared test mocks and assertions
├── validate                    <-- Support for request validation logic
├── web                         <-- Support for http transport layer requests
│   ├── handlers                <-- Http handlers used by the server
└── dockerfile.yaml
```

## Requirements
* [Docker installed](https://www.docker.com/community-edition)
* [Golang](https://golang.org)
* Optionally install godoc and staticcheck to generate documentation and lint the codebase

## Setup process


### Installing dependencies & building the target

To install all the dependencies needed to build a binary just run the following command:

```bash
make tidy
```

To run the application run: 

```bash
make run
```

or build the docker image and run a container :

```bash
docker build -t mortgage-calculator .
docker run -p 3000:3000 mortgage-calculator
```

## Example

http://localhost:3000/paymentSchedule [POST]

## Live Demo

http://bc-mortgage-calculator.s3-website-us-east-1.amazonaws.com/
http://3.235.248.35/paymentSchedule

### Request Body Example

```json
{
    "propertyPrice":      100000,
    "downPayment":        5000,
    "AnnualInterestRate": 4.29,
    "AmortizationPeriod": 5,
    "Schedule":           "Monthly" || "Biweekly" || "AcceleratedByWeekly"
}
```

### Testing

We use the `testing` package that is built-in in Golang and you can simply run the following command to run our tests:

```shell
make test
```

If you want to generate the coverage report:
```shell
make coverage
```

# Appendix

### Golang installation

Please ensure Go 1.x (where 'x' is the latest version) is installed as per the instructions on the official golang website: https://golang.org/doc/install

A quickstart way would be to use Homebrew, chocolatey or your linux package manager.

#### Homebrew (Mac)

Issue the following command from the terminal:

```shell
brew install golang
```

If it's already installed, run the following command to ensure it's the latest version:

```shell
brew update
brew upgrade golang
```

#### Chocolatey (Windows)

Issue the following command from the powershell:

```shell
choco install golang
```

If it's already installed, run the following command to ensure it's the latest version:

```shell
choco upgrade golang
```

## Generating GoDocs

To generate the project documentation run the following command:

```
make docs
```

## Built With

* [Go](https://golang.org/) - The Go programming language
