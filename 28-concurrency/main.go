package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

// func RepeatFunc[T any, K any](done <-chan K, fn func() T) <-chan T {
// 	// This is a Generator
// 	// This function accepts 2 parameters of type any with
// 	// Generics
// 	stream := make(chan T)
// 	go func() {
// 		defer close(stream)
// 		// a while loop in golang is "for" without any argument
// 		for {
// 			select {
// 			case <-done:
// 				return
// 			case stream <- fn():
// 			}
// 		}

// 	}()
// 	return stream

// }

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

func main() {
	done1 := make(chan int)
	defer close(done1)

	randNumFetcher := func() int { return rand.Intn(5000000000) }
	randIntStream := RepeatFunc(done1, randNumFetcher)
	//primeStream := PrimeFinder(done1, randIntStream)
	// for rando := range Take(done1, primeStream, 10) {
	// 	fmt.Println(rando)
	// }

	//fan Out
	CPUCoreCount := runtime.NumCPU() // returnnumber of cpu core
	PrimeFinderChannels := make([]<-chan int, CPUCoreCount)
	for i := 0; i < CPUCoreCount; i++ {
		PrimeFinderChannels[i] = PrimeFinder(done1, randIntStream)
	}

	//fan in
	fannedInStream := FanIn(done1, PrimeFinderChannels...)
	for rando := range Take(done1, fannedInStream, 10) {
		fmt.Println(rando)
	}
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
