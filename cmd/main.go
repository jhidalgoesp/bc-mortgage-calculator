package main

import (
	"fmt"
	"log"
	"net/http"
	"quoter/pkg/web/handlers"
)

const port = 3000

func main() {
	fmt.Printf("Starting server at port %d\n", port)
	handler := http.HandlerFunc(handlers.PaymentScheduleHandler)
	log.Fatal(http.ListenAndServe(":3000", handler))
}
