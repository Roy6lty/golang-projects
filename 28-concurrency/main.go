package main

import "fmt"

func someFunc(num string) {
	fmt.Println(num)
}
func doWork(done <-chan bool) {
	for {
		select {
		case <-done: //until the channel is closed the done will b enempty
			return
		default:
			fmt.Println("Doing Work")
		}
	}
}

func sliceToChannel(nums []int) <-chan int {
	out := make(chan int) // create a channel unbuffered channel;

	//since this is a go routine the main thread will not wait it wil go in to
	// execute the next line while go routing running
	//
	go func() {
		for _, n := range nums { // values from the slice into the channel
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	//recieves a channel as arg
	//returns a channel as an arg

	out := make(chan int)
	go func() {
		for n := range in { //taking values from the first channel feeding it into the second
			out <- n * n
		}
		close(out)
	}()
	return out
}

func repeatFunc[T any, K any](done <-chan K, fn func() T) {
	go func() {
		defer close(stream)
		// a while loop in golang is "for" without any argument

	}()

}
func main() {
	go someFunc("1")
	go someFunc("2")
	go someFunc("3")

	fmt.Println("hi")

	myChannel := make(chan string) //allocates memory for channel
	//chan is a keyword for channel
	myChannel2 := make(chan string)

	go func() {
		myChannel <- "data" // sending data to a channel
	}()

	go func() {
		myChannel2 <- "data2" // sending data to a channel
	}()

	select {
	case msgFromMyChannel := <-myChannel:
		fmt.Println(msgFromMyChannel)
	case msgFromMyChannel2 := <-myChannel2:
		fmt.Println(msgFromMyChannel2)
	}

	charChannel := make(chan string, 3)
	chars := []string{"a", "b", "c"}

	for _, s := range chars {
		select {
		case charChannel <- s:
		}
	}
	close(charChannel)

	for result := range charChannel {
		fmt.Println(result)
	}

	//For Select Loop Pattern

	done := make(chan bool)

	go doWork(done)

	close(done)

}
