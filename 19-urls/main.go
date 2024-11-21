package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://www.pillartechnologybackend.com.ng/docs#/"

func main() {
	fmt.Println("Welcome to handling URLs in golang ")
	fmt.Println(myurl)

	//parsing
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query() // returns a key value pair for parameters
	fmt.Printf("The type of query parameters are: %T\n", qparams)

	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Param is:", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Path:    "lco.dev",
		Host:    "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)

}
