func main() { // routine 1
	x := make(chan int) // chan 1
	y := make(chan int) // chan 2
	a := make(chan int) // chan 3

	go func() { x <- 1 }()      // routine 2
	go func() { <-x; x <- 1 }() // routine 3
	go func() { y <- 1; <-x }() // routine 4
	go func() { <-y }()         // routine 5

	select {
	case <-a:
	default:
	}
}
