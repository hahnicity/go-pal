package gopal

import (
    "encoding/json"
    "github.com/hahnicity/go-pal/config"
    "io/ioutil"
    "net/http"
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

// Add necessary request headers
func addOAuthHeaders(app *Application, req *http.Request) *http.Request {
    req.SetBasicAuth(app.id, app.secret)
    req.Header.Add("Accept", "application/json")
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    return req
}

// Receive a Response object we will use to get an oauth token
func makeOAuthRequest(app *Application) *http.Response {
    client := &http.Client{}
    req, err := http.NewRequest(
        "POST", 
        app.Endpoint + config.OauthEndpoint,
        strings.NewReader("grant_type=client_credentials"),
    )
    checkForError(err)
    req = addOAuthHeaders(app, req)
    resp, err := client.Do(req)
    checkForError(err)
    return resp
}

// Returns the a JSON-like response of calling the OAuth API
func GetOAuthResponse(app *Application) OAuthResponse {
    resp := makeOAuthRequest(app)
    defer resp.Body.Close()
    checkIfTokenReceived(resp.Status)
    rawResponse, _ := ioutil.ReadAll(resp.Body)
    var o OAuthResponse
    err := json.Unmarshal(rawResponse, &o)
    checkForError(err)
    return o
}

// Returns an OAuth token from PayPal
func GetToken(endpoint, id, secret string) string {
    app := MakeApplication(endpoint, id, secret)
    return GetOAuthResponse(app).Access_token
} 

// Check to ensure a generic error was not propogated
func checkForError(err error) {
    if err != nil {
        panic(err)    
    }
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
