package main

import "math"

func main() {
	// Variable declaration and initialization
	// Camel case notation is used for naming variables in Go.
	//var is used to declare a variable.
	// Here, we are declaring three variables: investmentAmount, expectedReturnRate, and years.
	// investmentAmount is of type int (integer).
	// expectedReturnRate is of type float64 (floating-point number).
	// years is of type int (integer).
	var investmentAmount = 1000
	var expectedReturnRate = 5.5
	var years = 10

	// Future Value Calculation using the formula:
	// FV = P * (1 + r)^n
	// where:
	// FV = Future Value
	// P = Principal amount (initial investment)
	// r = annual interest rate (decimal)
	// n = number of years the money is invested
	var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))

	println("Future Value of the Investment:", futureValue)

}
