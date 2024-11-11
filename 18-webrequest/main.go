package main

import (
	"fmt"
	"io"
	"net/http"
)
const url = "https://lco.dev"

func main(){
	fmt.Println("LCO web request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response Type is %T\n", response)

	defer response.Body.Close() // callers's responsibilty to close connection

	databytes, err := io.ReadAll(response.Body)
	if err != nil{
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)
}