go-pal
======

A REST-ful Paypal client for Go!

## Install

        go get github.com/hahnicity/go-pal

## Usage
Go-pal currently supports applications interfacing with PayPal REST APIs. 

It currently only supports obtaining an OAuth token, but much more is to come!!

### Getting an OAuth Token
        
        import "gopal"

        token := gopal.GetToken(<API endpoint>, <app id>, <app secret>)

If a non-200 status code is received when trying to obtain a token then an
error will be thrown and the program will fail
