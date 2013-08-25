package gopal

import (
    "bytes"
    "encoding/json"
    "github.com/hahnicity/go-pal"
    "io/ioutil"
    "net/http"
    "strings"
)

type Application struct {
    Endpoint string
    id string    
    secret string
}

type oauth struct {
    grant_type string
}

// Add necessary request headers
func addOAuthHeaders(app *Application, resp *http.Response) *http.Response {
    resp.Header.Add("Authorization", id + ":" + secret)
    resp.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    resp.Header.Add("Accept", "application/json")
    return resp
}

// Receive a Response object we will use to get an oauth token
func makeOAuthRequest(app *Application) *http.Response {
    tr := new(http.Transport)
    client := http.Client{Transport: tr}
    body, err := json.Marshal(oauth{"client_credentials"})
    checkForError(err)
    resp, err := client.Post(
        app.endpoint + config.oauthEndpoint, 
        "jsonp", 
        bytes.NewBuffer(body),
    )
    checkForError(err)
    return resp
}

// Request an OAuth token from PayPal
func GetToken(app *Application) []byte {
    resp := makeOAuthRequest(app)
    defer resp.Body.Close()
    resp = addOAuthHeaders(app, resp)
    checkIfTokenReceived(resp.Status)
    token, _ := ioutil.ReadAll(resp.Body)
    return token
}

// Check to ensure a generic error was not propogated
func checkForError(err error) {
    if err != nil {
        panic(err)    
    }
}

func checkIfTokenReceived(status string) {
    if strings.Split(status, " ")[0] != "200" {
        panic("You received a " + status + " status code")    
    }    
}

// constructor function for Application
func MakeApplication(endpoint, id, secret string) *Application {
    app := *Application{endpoint, id, secret}
    return app
}
