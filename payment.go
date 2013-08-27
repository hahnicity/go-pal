package gopal

import (
    "encoding/json"
    "github.com/hahnicity/go-pal/config"
)

// Create a payment for immediate sale
func (p *PaymentRequest) CreateImmediatePayment(endpoint, token string) (*PaymentResponse, error) {
    p.Intent = "sale"
    endpoint = endpoint + config.PaymentEndpoint
    return processNewPayment(endpoint, token, p)
}

// Create a payment for later sale/authorization
func (p *PaymentRequest) CreateDelayedPayment(endpoint, token string) (*PaymentResponse, error) {
    p.Intent = "authorize"
    endpoint = endpoint + config.PaymentEndpoint
    return processNewPayment(endpoint, token, p)
}

// Execute a pending payment
func (e *ExecutionRequest) ExecutePayment(endpoint, token, id string) (response *ExecutionResponse, err error) {
    endpoint = endpoint + config.PaymentEndpoint + "/" + id + "/execute" 
    rawResponse, err := makeRequestWithToken(endpoint, token, e)
    if err != nil {
        return nil, err    
    }
    err = json.Unmarshal(rawResponse, &response)
    raiseIfError(err)
    return response, nil
}

// Process the creation of a new payment
func processNewPayment(endpoint, token string, p *PaymentRequest) (response *PaymentResponse, err error) {
    rawResponse, err := makeRequestWithToken(endpoint, token, p)
    if err != nil {
        return nil, err    
    }
    err = json.Unmarshal(rawResponse, &response)
    raiseIfError(err)
    return response, nil
}
