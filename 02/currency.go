/*

In this example we will make a call to http://www.floatrates.com/daily/eur.json and display the value of Euro in INR and USD

We will use following packages...
- encoding/json :- we will use this for decoding JSON Object
- fmt :- we will use this to display data to the console
- net/http :- we will use this to  make a call to http://www.floatrates.com/daily/eur.json

*/

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// This is the json object for INR
// "inr":{"code":"INR","alphaCode":"INR","numericCode":"356","name":"Indian Rupee","rate":84.064779503976,"date":"Sun, 31 May 2020 12:00:01 GMT","inverseRate":0.011895588210669}
// Let's Create a struct called Response which has INR and USD field.

type Response struct {
	INR INR `json:"inr"`
	USD USD `json:"usd"`
}

// "name":"Indian Rupee","rate":84.064779503976
// Now lets create a struct for INR which will have rate and name field.

type INR struct {
	Rate float64 `json:"rate"`
	Name string  `json:"name"`
}

// Here we create a struct for USD which will have rate and name field.

type USD struct {
	Rate float64 `json:"rate"`
	Name string  `json:"name"`
}

// Now let's create a function

func GetCurrency() {

	// Let's create a variable named response for type Response
	var response Response

	// We initiate a GET Request to http://www.floatrates.com/daily/eur.json
	resp, err := http.Get("http://www.floatrates.com/daily/eur.json")

	if err != nil {
		panic(err)
	}
	// Let's close the response once we return from the function.

	defer resp.Body.Close()

	// Let's Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response)

	if err != nil {
		fmt.Println(err)
	}

	// Here we are using fmt funtion to display the Rate and Name
	fmt.Printf("1 Euro = %v %v\n", response.INR.Rate, response.INR.Name)
	fmt.Printf("1 Euro = %v %v\n", response.USD.Rate, response.USD.Name)

}

// Main function which calls  GetCurrency()
func main() {

	GetCurrency()
}
