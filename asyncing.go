package asyncing

import "fmt"

func hello(msngs chan string) {
	msg := <-msngs

	if msg != "" {
		fmt.Println(msg)
	} else {
		fmt.Println("Hello World")
	}
}

func worker(input <-chan string, output chan<- string) {
	for in := range input {
		output <- fmt.Sprintf("Hello %s", in)
	}
}

func helloIn(in chan string) {
	for msg := range in {
		fmt.Println(msg)
	}
}

func helloOut(in, outA, outB chan string) {
	for msg := range in {
		select {
		case outA <- msg:
		case outB <- msg:
		}
	}
}
