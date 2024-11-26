package main

import "sync"

func RepeatFunc[T any, K any](done <-chan K, fn func() T) <-chan T {
	// This is a Generator
	// This function accepts 2 parameters of type any with
	// Generics
	stream := make(chan T)
	go func() {
		defer close(stream)
		// a while loop in golang is "for" without any argument
		for {
			select {
			case <-done:
				return
			case stream <- fn():
			}
		}

	}()
	return stream

}

func Take[T any, K any](done <-chan K, stream <-chan T, n int) <-chan T {
	//  Function accept 2  send channel type and an int return a recieve channel
	taken := make(chan T)
	go func() {
		defer close(taken)
		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			case taken <- <-stream: //writing from one channel to another
			}
		}
	}()
	return taken

}

func PrimeFinder(done <-chan int, randomIntStream <-chan int) <-chan int {
	isPrime := func(randomInt int) bool {
		//checking for prime number
		for i := randomInt - 1; i > 1; i-- {
			if randomInt%i == 0 {
				return false
			}
		}
		return true
	}
	primes := make(chan int)
	go func() {
		defer close(primes)
		for { // infinite for loop
			select {
			case <-done:
				return
			case randomInt := <-randomIntStream:
				if isPrime(randomInt) {
					primes <- randomInt
				}

			}
		}

	}()
	return primes

}

func FanIn[T any](done <-chan int, channels ...<-chan T) <-chan T {
	var wg sync.WaitGroup
	fannedInStream := make(chan T)

	transfer := func(c <-chan T) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case fannedInStream <- i:
			}
		}
	}

	for _, c := range channels {
		wg.Add(1)
		go transfer(c)
	}

	go func() {
		wg.Wait()
		close(fannedInStream)

	}()
	return fannedInStream

}
