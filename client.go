package gopal

import (
    "bytes"
    "encoding/json"
    "github.com/hahnicity/go-pal/config"
    "net/http"
    "strings"
)

// Add oauth request headers
func addOAuthHeaders(app *Application, req *http.Request) *http.Request {
    req.SetBasicAuth(app.id, app.secret)
    req.Header.Add("Accept", "application/json")
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    return req
}

func addBearerHeaders(token string, req *http.Request) *http.Request{
    req.Header.Set("Authorization", "Bearer " + token)
    req.Header.Set("Content-Type", "application/json")
    return req
}

// Get the response object we will use to get an oauth token
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

// Get the response we will use for all transactions made with a token 
func makeRequestWithToken(endpoint, token string, data interface{}) *http.Response {
    client := &http.Client{}
    jsonData, err := json.Marshal(data)
    checkForError(err)
    req, err := http.NewRequest(
        "POST",
        endpoint,
        bytes.NewBuffer(jsonData),
    )
    checkForError(err)
    req = addBearerHeaders(token, req)
    resp, err := client.Do(req)
    return resp
}
