// declare a package to collect related functions
package greetings

import (
	//! import Go standard error package
	"errors"

	"fmt"

	//? math/rand generates a random number
	"math/rand"

	"time"

)


// Hello returns a greeting for the named person.
//* function X(parameter parameter's type) return type {}
//! any Go function can return mulitple values
//* Allow the caller to check the second vlue to see if an error occured
func Hello(name string) (string, error) {

	//! Error handling control flow
	// if no name is given, return an error with a message
	if name == "" {
		return "", errors.New("empty name")
	}
	/* 
		* the `:=` operator is a shorcut for delcaring and initializing a variable in one line. 
		* Go uses the value on the right to determine the variable's type

	    ? long way -
		* decalre variables with var, name the variable, type of the variable
		var message string
		message = fmt.Sprintf("Hi, %v. Welcome!", name)
	*/

    // Return a greeting that embeds the name in a message.
	// * 1st argument is a fromated string, Sprintf will sub in the name param's value for the %v format verb. 
    // message := fmt.Sprintf("Hi, %v. Welcome!", name)

	// invoke the randomFormat func
    message := fmt.Sprintf(randomFormat(), name)

	//! add nil (meaning no error) as a second value, this allows the caller to know if the function was succesful. 
    return message, nil
}

// init sets inital values for variables used in the function
// * Go executes init functions automatically at program startup, after gloabl variables have been initialized.
func init() {
	// seed random package with current time
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messagees
// the returned message is selected at random
func randomFormat() string {
	// A slice of messages formats
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v",
		"%v, you're a meat popsickle",
	}

	// Return a randomly selected message format by specifying a random index for the slice of formats
	return formats[rand.Intn(len(formats))]
}