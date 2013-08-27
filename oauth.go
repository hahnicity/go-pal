package gopal

import (
    "encoding/json"
    "strings"
)

type OAuthResponse struct {
    Scope string
    Access_token string
    Token_type string
    App_id string
    Expires_in int    
}

type Application struct {
    Endpoint string
    id string    
    secret string
}

// Returns an OAuth token from PayPal
func GetToken(endpoint, id, secret string) (*string, error) {
    app := MakeApplication(endpoint, id, secret)
    rawResponse, err := makeOAuthRequest(app)
    if err != nil {
        return nil, err    
    } 
    var o OAuthResponse
    err = json.Unmarshal(rawResponse, &o)
    raiseIfError(err)
    return &o.Access_token, nil
} 

// Check to see if an authentication token was received
func checkIfTokenReceived(status string) {
    if strings.Split(status, " ")[0] != "200" {
        panic("You received a <" + status + "> status code")
        // XXX Add debugging for why this happened
    }    
}

// constructor function for Application
func MakeApplication(endpoint, id, secret string) *Application {
    app := new(Application)
    app.Endpoint = endpoint 
    app.id = id
    app.secret = secret
    return app
}
