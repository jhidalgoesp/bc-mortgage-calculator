package mortgage

import (
	"github.com/jhidalgoesp/bc-mortgage-calculator/pkg/tests"
	"github.com/jhidalgoesp/bc-mortgage-calculator/pkg/validate"
	"testing"
)

func TestValidateAmortizationPeriod(t *testing.T) {
	t.Run("when the amortization period is not a 5 year increment return a not multiple of 5 error", func(t *testing.T) {
		c := Calculator{
			AmortizationPeriod: 1,
		}
		got := c.validateAmortizationPeriod()
		tests.AssertEqualErrors(t, got, ErrPeriodNotAMultipleOfFive)
	})

	t.Run("when the amortization period is between the maxium and minimum amortization period constants return true", func(t *testing.T) {
		c := Calculator{
			AmortizationPeriod: 10,
		}
		got := c.validateAmortizationPeriod()
		tests.AssertNilError(t, got)
	})

	t.Run("when the amortization period is equals to the maximumAmortizationPeriod constant return true", func(t *testing.T) {
		c := Calculator{
			AmortizationPeriod: 30,
		}
		got := c.validateAmortizationPeriod()
		tests.AssertNilError(t, got)
	})

	t.Run("when the amortization period is equals to the minimumAmortizationPeriod constant return false", func(t *testing.T) {
		c := Calculator{
			AmortizationPeriod: 5,
		}
		got := c.validateAmortizationPeriod()
		tests.AssertNilError(t, got)
	})

	t.Run("when the amortization period is lower than the minimumAmortizationPeriod constant return an out of range error", func(t *testing.T) {
		c := Calculator{
			AmortizationPeriod: 0,
		}
		got := c.validateAmortizationPeriod()
		tests.AssertEqualErrors(t, got, ErrPeriodOutOfRange)
	})

	t.Run("when the amortization period is higher than the maximumAmortizationPeriod constant return an out of range error", func(t *testing.T) {
		c := Calculator{
			AmortizationPeriod: 35,
		}
		got := c.validateAmortizationPeriod()
		tests.AssertEqualErrors(t, got, ErrPeriodOutOfRange)
	})
}

func TestCalculateCMHCRate(t *testing.T) {
	t.Run(`when the percentage of the down payment is less than 5% of the property value return an down payment
		 not large enough error`, func(t *testing.T) {
		c := Calculator{
			PropertyPrice: 100000,
			DownPayment:   0,
		}
		got, err := c.calculateCMHCRate()
		tests.AssertSameFloat(t, got, 0.0)
		tests.AssertEqualErrors(t, err, ErrDownPaymentNotLargeEnough)
	})

	t.Run(`when the percentage of the down payment is equals to 5% return a 4% CMHC rate`, func(t *testing.T) {
		c := Calculator{
			PropertyPrice: 100000,
			DownPayment:   5000,
		}
		got, err := c.calculateCMHCRate()
		AssertFloatValuesAndNilError(t, err, got, 4)
	})

	t.Run(`when the percentage of the down payment is between 5% - 9.99% return a 4% CMHC rate`, func(t *testing.T) {
		c := Calculator{
			PropertyPrice: 100000,
			DownPayment:   6000,
		}
		got, err := c.calculateCMHCRate()
		AssertFloatValuesAndNilError(t, err, got, 4)
	})

	t.Run(`when the percentage of the down payment is 9.99% return a 4% CMHC rate`, func(t *testing.T) {
		c := Calculator{
			PropertyPrice: 100000,
			DownPayment:   9999.99,
		}
		got, err := c.calculateCMHCRate()
		AssertFloatValuesAndNilError(t, err, got, 4)
	})

	t.Run(`when the percentage of the down payment is equals to 10% return a 3.10% CMHC rate`, func(t *testing.T) {
		c := Calculator{
			PropertyPrice: 100000,
			DownPayment:   10000,
		}
		got, err := c.calculateCMHCRate()
		AssertFloatValuesAndNilError(t, err, got, 3.10)
	})

	t.Run(`when the percentage of the down payment is between 10% - 14.99% return a 3.10% CMHC rate`, func(t *testing.T) {
		c := Calculator{
			PropertyPrice: 100000,
			DownPayment:   11000,
		}
		got, err := c.calculateCMHCRate()
		AssertFloatValuesAndNilError(t, err, got, 3.10)
	})

	t.Run(`when the percentage of the down payment is 14.99% return a 3.10% CMHC rate`, func(t *testing.T) {
		c := Calculator{
			PropertyPrice: 100000,
			DownPayment:   14999.99,
		}
		got, err := c.calculateCMHCRate()
		AssertFloatValuesAndNilError(t, err, got, 3.10)
	})

	t.Run(`when the percentage of the down payment is equals to 15% return a 2.80% CMHC rate`, func(t *testing.T) {
		c := Calculator{
			PropertyPrice: 100000,
			DownPayment:   15000,
		}
		got, err := c.calculateCMHCRate()
		AssertFloatValuesAndNilError(t, err, got, 2.80)
	})

	t.Run(`when the percentage of the down payment is between 15% - 19.99% return a 2.80% CMHC rate`, func(t *testing.T) {
		c := Calculator{
			PropertyPrice: 100000,
			DownPayment:   16000,
		}
		got, err := c.calculateCMHCRate()
		AssertFloatValuesAndNilError(t, err, got, 2.80)
	})

	t.Run(`when the percentage of the down payment is 19.99% return a 2.80% CMHC rate`, func(t *testing.T) {
		c := Calculator{
			PropertyPrice: 100000,
			DownPayment:   19999.99,
		}
		got, err := c.calculateCMHCRate()
		AssertFloatValuesAndNilError(t, err, got, 2.80)
	})

	t.Run(`when the percentage of the down payment is 20% return a 0% CMHC rate`, func(t *testing.T) {
		c := Calculator{
			PropertyPrice: 100000,
			DownPayment:   20000,
		}
		got, err := c.calculateCMHCRate()
		AssertFloatValuesAndNilError(t, err, got, 0.0)
	})

	t.Run(`when the percentage of the down payment is higher than 20% return a 0% CMHC rate`, func(t *testing.T) {
		c := Calculator{
			PropertyPrice: 100000,
			DownPayment:   210000,
		}
		got, err := c.calculateCMHCRate()
		AssertFloatValuesAndNilError(t, err, got, 0.0)
	})
}

func TestCalculateCMHC(t *testing.T) {
	t.Run("when CMHC is needed return the CMHC value, nil error", func(t *testing.T) {
		c := Calculator{
			PropertyPrice: 123456,
			DownPayment:   12223,
		}

		got, err := c.calculateCMHC()

		AssertFloatValuesAndNilError(t, err, got, 4449.32)
	})

	t.Run("when CMHC is not needed return 0, nil error", func(t *testing.T) {
		c := Calculator{
			PropertyPrice: 100000,
			DownPayment:   21000,
		}

		got, err := c.calculateCMHC()

		AssertFloatValuesAndNilError(t, err, got, 0)
	})

	t.Run("when down payment is not large enough return a not large enough down payment error", func(t *testing.T) {
		c := Calculator{
			PropertyPrice: 100000,
			DownPayment:   1000,
		}

		got, err := c.calculateCMHC()

		tests.AssertSameFloat(t, got, 0.0)
		tests.AssertEqualErrors(t, err, ErrDownPaymentNotLargeEnough)
	})
}

func TestCalculateTotalMortgage(t *testing.T) {
	t.Run("when a CMHC is needed the mortgage should equal the property price - payment + CMHC, nil error", func(t *testing.T) {
		c := Calculator{
			PropertyPrice: 123456,
			DownPayment:   12223,
		}
		got, err := c.calculateTotalMortgage()
		AssertFloatValuesAndNilError(t, err, got, 115682.32)
	})

	t.Run("when a CMHC is not needed the mortgage should equal the property price - payment, nil error", func(t *testing.T) {
		c := Calculator{
			PropertyPrice: 100000,
			DownPayment:   21000,
		}
		got, err := c.calculateTotalMortgage()
		AssertFloatValuesAndNilError(t, err, got, c.PropertyPrice-c.DownPayment)
	})

	t.Run("when down payment is not large enough return a not large enough down payment error", func(t *testing.T) {
		c := Calculator{
			PropertyPrice: 100000,
			DownPayment:   1000,
		}
		got, err := c.calculateTotalMortgage()
		tests.AssertSameFloat(t, got, 0.0)
		tests.AssertEqualErrors(t, err, ErrDownPaymentNotLargeEnough)
	})
}

func TestPaymentsPerYear(t *testing.T) {
	t.Run("when schedule is AcceleratedBiweekly return 26", func(t *testing.T) {
		c := Calculator{
			Schedule: AcceleratedBiweekly,
		}
		got, err := c.paymentsPerYear()
		AssertIntValuesAndNilError(t, err, got, 12)
	})

	t.Run("when schedule is Biweekly return 12", func(t *testing.T) {
		c := Calculator{
			Schedule: Biweekly,
		}
		got, err := c.paymentsPerYear()
		AssertIntValuesAndNilError(t, err, got, 26)
	})

	t.Run("when schedule is Monthly return 12", func(t *testing.T) {
		c := Calculator{
			Schedule: Monthly,
		}
		got, err := c.paymentsPerYear()
		AssertIntValuesAndNilError(t, err, got, 12)
	})

	t.Run("when schedule is not supported return an error", func(t *testing.T) {
		c := Calculator{
			Schedule: "Yearly",
		}
		got, err := c.totalNumberOfPayments()
		tests.AssertEqualErrors(t, err, errInvalidSchedule)
		tests.AssertSameInt(t, got, 0)
	})
}

func TestTotalNumberOfPayments(t *testing.T) {
	t.Run("when schedule is AcceleratedBiweekly return 12 * amortization period", func(t *testing.T) {
		c := Calculator{
			Schedule:           AcceleratedBiweekly,
			AmortizationPeriod: 5,
		}
		got, err := c.totalNumberOfPayments()
		AssertIntValuesAndNilError(t, err, got, 12*c.AmortizationPeriod)
	})

	t.Run("when schedule is Biweekly return 26 * amortization period", func(t *testing.T) {
		c := Calculator{
			Schedule:           Biweekly,
			AmortizationPeriod: 5,
		}
		got, err := c.totalNumberOfPayments()
		AssertIntValuesAndNilError(t, err, got, 26*c.AmortizationPeriod)
	})

	t.Run("when schedule is Monthly return 12 * amortization period", func(t *testing.T) {
		c := Calculator{
			Schedule:           AcceleratedBiweekly,
			AmortizationPeriod: 5,
		}
		got, err := c.totalNumberOfPayments()
		AssertIntValuesAndNilError(t, err, got, 12*c.AmortizationPeriod)
	})

	t.Run("when schedule is not supported return an error", func(t *testing.T) {
		c := Calculator{
			Schedule:           "Yearly",
			AmortizationPeriod: 5,
		}
		got, err := c.totalNumberOfPayments()
		tests.AssertEqualErrors(t, err, errInvalidSchedule)
		tests.AssertSameInt(t, got, 0)
	})
}

func TestPaymentSchedule(t *testing.T) {
	t.Run("should calculate the Monthly payment schedule ", func(t *testing.T) {
		c := Calculator{
			PropertyPrice:      100000,
			DownPayment:        5000,
			AnnualInterestRate: 4.29,
			AmortizationPeriod: 5,
			Schedule:           Monthly,
		}

		got, err := c.PaymentSchedule()
		AssertFloatValuesAndNilError(t, err, got, 1832.51)
	})

	t.Run("should calculate the Biweekly payment schedule ", func(t *testing.T) {
		c := Calculator{
			PropertyPrice:      100000,
			DownPayment:        5000,
			AnnualInterestRate: 4.29,
			AmortizationPeriod: 5,
			Schedule:           Biweekly,
		}

		got, err := c.PaymentSchedule()
		AssertFloatValuesAndNilError(t, err, got, 845.05)
	})

	t.Run("should calculate the accelerated Biweekly payment schedule ", func(t *testing.T) {
		c := Calculator{
			PropertyPrice:      100000,
			DownPayment:        5000,
			AnnualInterestRate: 4.29,
			AmortizationPeriod: 5,
			Schedule:           AcceleratedBiweekly,
		}

		got, err := c.PaymentSchedule()
		AssertFloatValuesAndNilError(t, err, got, 916.255)
	})

	t.Run("should return a error if Calculator does not pass validation check", func(t *testing.T) {
		c := Calculator{
			PropertyPrice:      100000,
			DownPayment:        5000,
			AnnualInterestRate: 4.29,
			AmortizationPeriod: 5,
		}

		_, err := c.PaymentSchedule()
		fieldError := validate.GetFieldErrors(err)[0]

		wantedField := "schedule"
		wantedMessage := "schedule is a required field"
		if fieldError.Field != wantedField {
			t.Errorf("expected %s got %s", wantedField, fieldError.Field)
		}

		if fieldError.Error != wantedMessage {
			t.Errorf("expected %s got %s", wantedMessage, fieldError.Error)
		}
	})

	t.Run("should return a error if validateAmortizationPeriod fails", func(t *testing.T) {
		c := Calculator{
			PropertyPrice:      100000,
			DownPayment:        5000,
			AnnualInterestRate: 4.29,
			AmortizationPeriod: 80,
			Schedule:           Biweekly,
		}
		AssertHandledScheduleErrors(t, &c, ErrPeriodOutOfRange)
	})

	t.Run("should return a error if calculateTotalMortgage fails", func(t *testing.T) {
		c := Calculator{
			PropertyPrice:      100000,
			DownPayment:        10,
			AnnualInterestRate: 4.29,
			AmortizationPeriod: 20,
			Schedule:           Biweekly,
		}
		AssertHandledScheduleErrors(t, &c, ErrDownPaymentNotLargeEnough)
	})

	t.Run("should return a error if scheduleInterestRate fails", func(t *testing.T) {
		c := Calculator{
			PropertyPrice:      100000,
			DownPayment:        5000,
			AnnualInterestRate: 4.29,
			AmortizationPeriod: 5,
			Schedule:           "yearly",
		}
		AssertHandledScheduleErrors(t, &c, errInvalidSchedule)
	})
}

func AssertHandledScheduleErrors(t testing.TB, c *Calculator, wantedErr error) {
	_, err := c.PaymentSchedule()
	tests.AssertEqualErrors(t, err, wantedErr)
}

func AssertIntValuesAndNilError(t testing.TB, err error, got, want int) {
	tests.AssertSameInt(t, got, want)
	tests.AssertEqualErrors(t, err, nil)
}

func AssertFloatValuesAndNilError(t testing.TB, err error, got, want float64) {
	tests.AssertSameFloat(t, got, want)
	tests.AssertEqualErrors(t, err, nil)
}
