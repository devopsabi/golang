package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	// The os.Args variable holds the command-line arguments

	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "https://google.com")
	}

	url := os.Args[1] // I am storing the 1st argument in the variable called url

	resp, err := http.Get(url)


	if err != nil {
		panic(err)
	}

	// In Go language, defer statements delay the execution of the function or method or an anonymous method until the nearby functions returns
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	// ioutil.Readall reads data from io.Reader
	body, err := ioutil.ReadAll(resp.Body)

	// ioutil.WriteFile writes data to output.txt
	err = ioutil.WriteFile("output.txt", body, 0644)

	if err != nil {
		panic(err)
	}

}
