package gopal

// Raise if a critical error occurs
func raiseIfError(err error) {
    if err != nil {
        panic(err)    
    }
}
