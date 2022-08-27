package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/jhidalgoesp/bc-mortgage-calculator/pkg/mortgage"
	"github.com/jhidalgoesp/bc-mortgage-calculator/pkg/tests"
	"github.com/jhidalgoesp/bc-mortgage-calculator/pkg/validate"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPaymentScheduleHandler(t *testing.T) {
	c := mortgage.Calculator{
		PropertyPrice:      100000,
		DownPayment:        5000,
		AnnualInterestRate: 4.29,
		AmortizationPeriod: 5,
		Schedule:           mortgage.Monthly,
	}

	t.Run("returns the correct values for a monthly schedule", func(t *testing.T) {
		c.Schedule = mortgage.Monthly
		jsonBody, _ := json.Marshal(&c)
		request, _ := http.NewRequest(http.MethodPost, "/paymentSchedule", bytes.NewBuffer(jsonBody))
		response := httptest.NewRecorder()
		PaymentScheduleHandler(response, request)
		paymentSchedule := paymentScheduleResponse{}
		json.NewDecoder(response.Body).Decode(&paymentSchedule)
		want := 1832.51
		tests.AssertSameFloat(t, paymentSchedule.PaymentPerSchedule, want)
	})

	t.Run("returns the correct values for a biweekly schedule", func(t *testing.T) {
		c.Schedule = mortgage.Biweekly
		jsonBody, _ := json.Marshal(&c)
		request, _ := http.NewRequest(http.MethodPost, "/paymentSchedule", bytes.NewBuffer(jsonBody))
		response := httptest.NewRecorder()
		PaymentScheduleHandler(response, request)
		paymentSchedule := paymentScheduleResponse{}
		json.NewDecoder(response.Body).Decode(&paymentSchedule)
		want := 845.05
		tests.AssertSameFloat(t, paymentSchedule.PaymentPerSchedule, want)
	})

	t.Run("returns the correct values for a accelerated biweekly schedule", func(t *testing.T) {
		c.Schedule = mortgage.AcceleratedBiweekly
		jsonBody, _ := json.Marshal(&c)
		request, _ := http.NewRequest(http.MethodPost, "/paymentSchedule", bytes.NewBuffer(jsonBody))
		response := httptest.NewRecorder()
		PaymentScheduleHandler(response, request)
		paymentSchedule := paymentScheduleResponse{}
		json.NewDecoder(response.Body).Decode(&paymentSchedule)
		want := 916.255
		tests.AssertSameFloat(t, paymentSchedule.PaymentPerSchedule, want)
	})

	t.Run("returns not found if the path is not supported", func(t *testing.T) {
		jsonBody, _ := json.Marshal(&c)
		request, _ := http.NewRequest(http.MethodPost, "/test", bytes.NewBuffer(jsonBody))
		response := httptest.NewRecorder()
		PaymentScheduleHandler(response, request)
		err := errorResponse{}
		json.NewDecoder(response.Body).Decode(&err)
		want := http.StatusText(http.StatusNotFound)
		if err.Error != want {
			t.Errorf("got %s, want %s", err.Error, want)
		}
	})

	t.Run("returns not found if the http method is not supported", func(t *testing.T) {
		jsonBody, _ := json.Marshal(&c)
		request, _ := http.NewRequest(http.MethodGet, "/paymentSchedule", bytes.NewBuffer(jsonBody))
		response := httptest.NewRecorder()
		PaymentScheduleHandler(response, request)
		err := errorResponse{}
		json.NewDecoder(response.Body).Decode(&err)
		want := http.StatusText(http.StatusNotFound)
		if err.Error != want {
			t.Errorf("got %s, want %s", err.Error, want)
		}
	})

	t.Run("returns internal server error if request body is not present", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/paymentSchedule", http.NoBody)
		response := httptest.NewRecorder()
		PaymentScheduleHandler(response, request)
		err := errorResponse{}
		json.NewDecoder(response.Body).Decode(&err)
		want := http.StatusText(http.StatusInternalServerError)
		if err.Error != want {
			t.Errorf("got %s, want %s", err.Error, want)
		}
	})

	t.Run("returns field errors if the body does not pass validations", func(t *testing.T) {
		c.DownPayment = 0
		jsonBody, _ := json.Marshal(&c)
		request, _ := http.NewRequest(http.MethodPost, "/paymentSchedule", bytes.NewBuffer(jsonBody))
		response := httptest.NewRecorder()
		PaymentScheduleHandler(response, request)
		err := errorResponse{}
		json.NewDecoder(response.Body).Decode(&err)
		errorField := err.Fields[0]
		want := validate.FieldError{
			Field: "downPayment",
			Error: "downPayment is a required field",
		}
		if want != errorField {
			t.Errorf("got %v, want %v", err, want)
		}
	})

	t.Run("when a preflight options request should respond ok status", func(t *testing.T) {
		jsonBody, _ := json.Marshal(&c)
		request, _ := http.NewRequest(http.MethodOptions, "/paymentSchedule", bytes.NewBuffer(jsonBody))
		response := httptest.NewRecorder()
		PaymentScheduleHandler(response, request)
		if response.Code != http.StatusOK {
			t.Errorf("got %v, want %v", response.Code, http.StatusOK)
		}
	})

	t.Run("handled errors", func(t *testing.T) {
		t.Run("returns a payment not large enough error response if the payment is < 5% of price", func(t *testing.T) {
			c.DownPayment = 1
			AssertHandledErrors(t, c, mortgage.ErrDownPaymentNotLargeEnough)
		})
		t.Run("returns a period out of range error response if amortization period is out of range", func(t *testing.T) {
			c.AmortizationPeriod = 90
			AssertHandledErrors(t, c, mortgage.ErrPeriodOutOfRange)
		})
		t.Run("returns a period not multiple of 5 error response if amortization", func(t *testing.T) {
			c.AmortizationPeriod = 6
			AssertHandledErrors(t, c, mortgage.ErrPeriodNotAMultipleOfFive)
		})
	})
}

func AssertHandledErrors(tb testing.TB, c mortgage.Calculator, erro error) {
	tb.Helper()
	jsonBody, _ := json.Marshal(&c)
	request, _ := http.NewRequest(http.MethodPost, "/paymentSchedule", bytes.NewBuffer(jsonBody))
	response := httptest.NewRecorder()
	PaymentScheduleHandler(response, request)
	err := errorResponse{}
	json.NewDecoder(response.Body).Decode(&err)
	want := erro.Error()
	if err.Error != want {
		tb.Errorf("got %v, want %v", err, want)
	}
}
