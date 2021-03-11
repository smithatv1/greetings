package greetings

import (
	"fmt"
"rsc.io/quote"
)
func Greetings() string {
	fmt.Println("Greetings from greetings module" )
	return quote.Hello()
} 
func Hello() string {
	fmt.Println("Greetings from greetings module" )
	return quote.Hello()
}
