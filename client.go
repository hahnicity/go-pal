package gopal

import (
    "bytes"
    "encoding/json"
    "github.com/hahnicity/go-pal/config"
    "io/ioutil"
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

// Add headers for all requests necessitating a Bearer token
func addBearerHeaders(token string, req *http.Request) *http.Request{
    req.Header.Set("Authorization", "Bearer " + token)
    req.Header.Set("Content-Type", "application/json")
    return req
}

// Get the response object we will use to get an oauth token
func makeOAuthRequest(app *Application) ([]byte, error) {
    client := &http.Client{}
    req, err := http.NewRequest(
        "POST", 
        app.Endpoint + config.OauthEndpoint,
        strings.NewReader("grant_type=client_credentials"),
    )
    raiseIfError(err)
    req = addOAuthHeaders(app, req)
    return postRequest(client, req)
}

// Get the response we will use for all transactions made with a token 
func makeRequestWithToken(endpoint, token string, data interface{}) ([]byte, error) {
    client := &http.Client{}
    jsonData, err := json.Marshal(data)
    raiseIfError(err)
    req, err := http.NewRequest(
        "POST",
        endpoint,
        bytes.NewBuffer(jsonData),
    )
    raiseIfError(err)
    req = addBearerHeaders(token, req)
    return postRequest(client, req)
}

// Post the constructed request to server and return the response in byte form
func postRequest(client *http.Client, req *http.Request) ([]byte, error){
    resp, err := client.Do(req)
    if err != nil {
        return nil, err    
    }
    defer resp.Body.Close()
    checkIfTokenReceived(resp.Status)
    data, err := ioutil.ReadAll(resp.Body)
    raiseIfError(err)
    return data, nil
}
