package main

func SliceToChannel(nums []int) <-chan int {
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

func Sq(in <-chan int) <-chan int {
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
