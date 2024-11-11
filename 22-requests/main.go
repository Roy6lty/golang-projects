package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main(){
	fmt.Println("Welcome to lco tutorial")
	PerformGetRequest()
	PerformPostJsonRequest()

}

func PerformPostJsonRequest(){
	const myurl = "http://127.0.0.0:8001/auth/login"

	//login details
	requestBody := strings.NewReader( `
	{
		"email":"olowoleru06@gmail.com",
		"password":"9kP?&7Ay"
    }
`)
	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))


}

func PerformGetRequest(){
	// Performing a Get request
	const myurl = "http://127.0.0.1:8001/ping"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close() 

	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content length:", response.ContentLength)

	//declared reponseString of type strings.Builder


	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	bytesCount, _ := responseString.Write(content)

	// converting the response from bytes to string type
	fmt.Println(string(content))

	// Another way of writing the response body to string instead of j
	//just wrapping the response body with  string
	fmt.Println("bytesCount:", bytesCount)
	fmt.Println(responseString.String())



}
