package mortgage

import (
	"errors"
	"github.com/jhidalgoesp/bc-mortgage-calculator/pkg/validate"
	"math"
	"strings"
)

const (
	AcceleratedBiweekly       = "ACCELERATEDBIWEEKLY"
	Biweekly                  = "BIWEEKLY"
	Monthly                   = "MONTHLY"
	minimumAmortizationPeriod = 5
	maximumAmortizationPeriod = 30
)

// ErrPeriodOutOfRange variables for calculation operations
var ErrPeriodOutOfRange = errors.New("amortization period out of range")
var ErrPeriodNotAMultipleOfFive = errors.New("amortization period must be a 5 years multiple")
var ErrDownPaymentNotLargeEnough = errors.New("down payment is lower than the minimum 5% of the property price")
var errInvalidSchedule = errors.New("amortization schedule not supported")

// Calculator holds the properties and exposes methods needed to perform mortgage calculations.
type Calculator struct {
	PropertyPrice      float64 `json:"propertyPrice" validate:"required,gtfield=DownPayment"`
	DownPayment        float64 `json:"downPayment" validate:"required,ltfield=PropertyPrice"`
	AnnualInterestRate float64 `json:"annualInterestRate" validate:"required"`
	AmortizationPeriod int     `json:"amortizationPeriod" validate:"required"`
	Schedule           string  `json:"schedule" validate:"required"`
}

func (c Calculator) PaymentSchedule() (float64, error) {
	err := validate.Check(c)
	if err != nil {
		return 0, err
	}

	err = c.validateAmortizationPeriod()
	if err != nil {
		return 0, err
	}

	principal, err := c.calculateTotalMortgage()
	if err != nil {
		return 0, err
	}

	scheduleRate, err := c.scheduleInterestRate()
	if err != nil {
		return 0, err
	}

	numberOfPayments, err := c.totalNumberOfPayments()
	if err != nil {
		return 0, err
	}

	paymentPerSchedule := principal * scheduleRate * (math.Pow(1+scheduleRate, float64(numberOfPayments))) /
		(math.Pow(1+scheduleRate, float64(numberOfPayments)) - 1)

	if strings.ToUpper(c.Schedule) == AcceleratedBiweekly {
		return math.Round(paymentPerSchedule*100) / 2 / 100, nil
	}

	return math.Round(paymentPerSchedule*100) / 100, nil
}

// scheduleInterestRate returns the interest rate depending on the selected schedule.
func (c *Calculator) scheduleInterestRate() (float64, error) {
	np, err := c.paymentsPerYear()
	if err != nil {
		return 0, err
	}
	return c.AnnualInterestRate / 100 / float64(np), nil
}

// validateAmortizationPeriod returns true if Amortization period is between the amortization period allowed and a 5-year increment.
func (c *Calculator) validateAmortizationPeriod() error {
	if c.AmortizationPeriod%5 != 0 {
		return ErrPeriodNotAMultipleOfFive
	}
	if c.AmortizationPeriod > maximumAmortizationPeriod || c.AmortizationPeriod < minimumAmortizationPeriod {
		return ErrPeriodOutOfRange
	}
	return nil
}

// calculateTotalMortgage performs the calculation of the CMHC mortgage value.
func (c *Calculator) calculateTotalMortgage() (float64, error) {
	CMHC, err := c.calculateCMHC()
	if err != nil {
		return 0, err
	}
	return (c.PropertyPrice - c.DownPayment) + CMHC, nil
}

// calculateCMHC performs the calculation of the CMHC mortgage value.
func (c *Calculator) calculateCMHC() (float64, error) {
	CMHCRate, err := c.calculateCMHCRate()
	if err != nil {
		return 0, err
	}
	return (c.PropertyPrice - c.DownPayment) * CMHCRate / 100, nil
}

// calculateCMHCRate performs the calculation of the percentage of the mortgage amount needed as insurance.
func (c *Calculator) calculateCMHCRate() (float64, error) {
	percentageHomePrice := c.DownPayment * 100 / c.PropertyPrice
	switch {
	case percentageHomePrice < 5:
		return 0, ErrDownPaymentNotLargeEnough
	case percentageHomePrice >= 5 && percentageHomePrice < 10:
		return 4.0, nil
	case percentageHomePrice >= 10 && percentageHomePrice < 15:
		return 3.1, nil
	case percentageHomePrice >= 15 && percentageHomePrice < 20:
		return 2.8, nil
	default:
		return 0, nil
	}
}

// paymentsPerYear return the total amount of payments done on a year.
func (c *Calculator) paymentsPerYear() (int, error) {
	switch strings.ToUpper(c.Schedule) {
	case AcceleratedBiweekly:
		return 12, nil
	case Biweekly:
		return 26, nil
	case Monthly:
		return 12, nil
	}
	return 0, errInvalidSchedule
}

// totalNumberOfPayments return the total amount of payments done on the amortization period	.
func (c *Calculator) totalNumberOfPayments() (int, error) {
	switch strings.ToUpper(c.Schedule) {
	case AcceleratedBiweekly:
		return 12 * c.AmortizationPeriod, nil
	case Biweekly:
		return 26 * c.AmortizationPeriod, nil
	case Monthly:
		return 12 * c.AmortizationPeriod, nil
	}
	return 0, errInvalidSchedule
}
