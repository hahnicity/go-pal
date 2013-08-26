package gopal

import (
    "net/http"
)

// Add headers to the payment object
func addPaymentHeaders(token string, req *http.Request) *http.Request{
    req.Header.Set("Authorization", "Bearer " + token)
    req.Header.Set("Content-Type", "application/json")
    return req
}

func CreatePayment(intent, p Payer, )
