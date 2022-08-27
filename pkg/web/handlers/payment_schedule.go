package handlers

import (
	"errors"
	"github.com/jhidalgoesp/bc-mortgage-calculator/pkg/mortgage"
	"github.com/jhidalgoesp/bc-mortgage-calculator/pkg/validate"
	"github.com/jhidalgoesp/bc-mortgage-calculator/pkg/web"
	"log"
	"net/http"
)

type paymentScheduleResponse struct {
	PaymentPerSchedule float64 `json:"paymentPerSchedule"`
}

type errorResponse struct {
	Error  string               `json:"error"`
	Fields validate.FieldErrors `json:"fields,omitempty"`
}

func notFoundResponse(w http.ResponseWriter) {
	resp := errorResponse{Error: http.StatusText(http.StatusNotFound)}
	web.Respond(w, resp, http.StatusNotFound)
}

func internalServerErrorResponse(w http.ResponseWriter) {
	resp := errorResponse{Error: http.StatusText(http.StatusInternalServerError)}
	web.Respond(w, resp, http.StatusInternalServerError)
}

func PaymentScheduleHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/paymentSchedule" {
		notFoundResponse(w)
		return
	}

	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3001")
		w.Header().Set("Access-Control-Allow-Headers", "content-type")
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		notFoundResponse(w)
		return
	}

	var calc mortgage.Calculator
	if err := web.Decode(r, &calc); err != nil {
		log.Println("unable to decode payload: %w", err)
		internalServerErrorResponse(w)
		return
	}

	paymentSchedule, err := calc.PaymentSchedule()
	if err != nil {
		switch err.(type) {
		case validate.FieldErrors:
			log.Println("data validation error: ", err)
			resp := errorResponse{Error: "data validation error", Fields: validate.GetFieldErrors(err)}
			web.Respond(w, resp, http.StatusBadRequest)
			return
		default:
			if errors.Is(err, mortgage.ErrDownPaymentNotLargeEnough) || errors.Is(err, mortgage.ErrPeriodOutOfRange) ||
				errors.Is(err, mortgage.ErrPeriodNotAMultipleOfFive) {
				log.Println("error calculating mortgage: ", err)
				resp := errorResponse{Error: err.Error()}
				web.Respond(w, resp, http.StatusBadRequest)
				return
			}
			log.Println("error: ", err.Error())
			internalServerErrorResponse(w)
			return
		}
	}
	resp := paymentScheduleResponse{PaymentPerSchedule: paymentSchedule}

	web.Respond(w, resp, http.StatusOK)
}
