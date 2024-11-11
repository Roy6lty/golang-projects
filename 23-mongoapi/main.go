package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/olowoleru/router"
)

func main(){
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("Sever is getting started....")
	log.Fatal(http.ListenAndServe(":4001", r))
	fmt.Println("Listening at port 4001....")
}