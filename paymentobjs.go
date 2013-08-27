/*
All objects contained within this file are golang representations of objects defined 
as "Common Payments Objects" in the PayPal REST API documentation:

https://developer.paypal.com/webapps/developer/docs/api/#payment-object
*/

package gopal

type Address struct {
    Line1 string
    Line2 string
    City string
    Country_code string
    Postal_code string
    State string
    Phone string
}

type Amount struct {
    Currency string
    Total string
    Details Details    
}

type Capture struct {
    // XXX TODO    
}

type CreditCard struct {
    Id string
    Payer_id string
    Number string
    Type string
    Expire_month int
    Expire_year int
    Cvv2 string
    First_name string
    Last_name string
    Billing_address Address
    State string
    Valid_until string
}

type CreditCardToken struct {
    Credit_card_id string
    Payer_id string
    Last4 string
    Type string
    Expire_year int
    Expire_month int
}

type Details struct {
    Shipping string    
    Subtotal string
    Tax string
    Fee string
}

type FundingInstrument struct {
    Credit_card CreditCard
    Credit_card_token CreditCardToken
}

type Item struct {
    Quantity string
    Name string
    Price string
    Currency string
    Sku string    
}

type ItemList struct {
    Items []Item
    Shipping_address ShippingAddress    
}

type Payer struct {
    Payment_method string
    Funding_instruments []FundingInstrument
    Payer_info PayerInfo
}

type PayerInfo struct {
    Email string
    First_name string
    Last_name string
    Payer_id string
    Phone string
    Shipping_address ShippingAddress    
}

type PaymentRequest struct {
    Intent string    
    Payer Payer
    Transactions []Transaction
    Redirect_urls RedirectURLs
}

type PaymentResponse struct {
    Intent string
    Payer Payer
    Transactions []Transaction
    Redirect_urls RedirectURLs
    Id string
    Create_time string // XXX date_time
    State string
    Update_time string // XXX date_time
}

type PaymentExecution struct {
    Payer_id string
    Transactions []Transaction    
}

type RedirectURLs struct {
    Return_url string
    Cancel_url string    
}

type Refund struct {
    // XXX TODO    
}

type Sale struct {
    Id string    
    Amount Amount
    Description string
    Create_time string // XXX date_time
    State string
    Sale_id string
    Parent_payment string
    Update_time string // XXX date_time
}

type ShippingAddress struct {
    Recipient_name string
    Type string
    Line1 string
    Line2 string
    City string
    Country_code string
    Postal_code string
    State string
    Phone string

}

type Transaction struct {
    Amount Amount    
    Description string
    Item_list ItemList
    Related_resources interface{} // Sale, Authorization, Capture, Refund
}
