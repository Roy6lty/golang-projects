package main

import (
	"fmt"
	"io"
	"os"
)

func main(){
	fmt.Println("welcome t files in golang")
	content := "This needs to go in  a file - LearnCodeOnline.in" 

	// create txt file
	file, err := os.Create("./mylcogofile.text")

	checkNilErr(err)

	// write string into text file
	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("length: ", length)

	// close file
	defer file.Close()
	readFile("./mylcogofile.text")  
}

// read file
func readFile(filename string){
	databyte , err := os.ReadFile(filename)

	checkNilErr(err)
	fmt.Println("Text data inside the file is \n", string(databyte))

}

func checkNilErr(err error){
	if err != nil{
		panic(err)
	}
}