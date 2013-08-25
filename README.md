go-pal
======

A REST-ful Paypal client for Go!

Is currently under development and should not be expected to work in current form

## Install

        go get github.com/hahnicity/go-pal

## Usage
Go-pal currently supports applications interfacing with PayPal REST APIs. 

### Getting an OAuth Token
        
        import "gopal"

        app := gopal.MakeApplication(<API endpoint>, <app id>, <app secret>)
        token := gopal.GetToken(app)

If a non-200 status code is received when trying to obtain a token then an
error will be thrown and the program will fail
