package main

import (
	"fmt"
	"github.com/jhidalgoesp/bc-mortgage-calculator/pkg/web/handlers"
	"log"
	"net/http"
)

const port = 3000

func main() {
	fmt.Printf("Starting server at port %d\n", port)
	handler := http.HandlerFunc(handlers.PaymentScheduleHandler)
	log.Fatal(http.ListenAndServe(":3000", handler))
}
