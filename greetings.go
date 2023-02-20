// declare a package to collect related functions
package greetings

import "fmt"


// Hello returns a greeting for the named person.
//* function X(parameter parameter's type) return type {}
func Hello(name string) string {

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
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}