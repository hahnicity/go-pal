go-pal
======

A REST-ful Paypal client for Go!

## Status
Go-pal had the intention of supporting applications interfacing with PayPal REST APIs. 
However due to messing around with other projects my attention waned, and this project
was quietly dropped

It currently only supports obtaining an OAuth token as of the master branch. If the 
develop branch is used there is much cleaner code to use than what is available in 
master plus some work made in the attempt to get payments to work. Although that
work was never completed

## Install

        go get github.com/hahnicity/go-pal

## Usage
Get OAuth tokens for your application!

### Getting an OAuth Token
        
        import "gopal"

        token := gopal.GetToken(<API endpoint>, <app id>, <app secret>)

If a non-200 status code is received when trying to obtain a token then an
error will be thrown and the program will fail
