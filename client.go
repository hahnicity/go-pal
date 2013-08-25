package main

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "net/http"
    "net/http/httputil"
    "strings"
)

type Application struct {
    Endpoint string
    id string    
    secret string
}

type PostData struct {
    Grant_Type string    
}

// XXX
const oauthEndpoint string = "/v1/oauth2/token"

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
        app.Endpoint + oauthEndpoint,
        strings.NewReader(`{"grant_type": "client_credentials"}`),
    )
    checkForError(err)
    req = addOAuthHeaders(app, req)
    
    fmt.Println(req.Body)

    resp, err := client.Do(req)
    dumpedReq, err := httputil.DumpRequest(req, true)
    fmt.Println(string(dumpedReq))
    dumpedRes, err := httputil.DumpResponse(resp, true)
    fmt.Println(string(dumpedRes))
    checkForError(err)
    return resp
}

// Request an OAuth token from PayPal
func GetToken(app *Application) []byte {
    resp := makeOAuthRequest(app)
    defer resp.Body.Close()
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

// Check to see if an authentication token was received
func checkIfTokenReceived(status string) {
    if strings.Split(status, " ")[0] != "200" {
        panic("You received a <" + status + "> status code")    
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

func main() {
    app := MakeApplication("https://api.sandbox.paypal.com", "AU3saBBLmtcRB-gglYmD1EDlrB53feI0NxE2JGWdY0_ppX-22dulztl63PYK", "EMgdYRAeCG5EN5T-5_eKcVcDd2I_lI8TvEFAR0zLe3bPtDCMyLdaXLr9xnvr")
    GetToken(app)    
}
