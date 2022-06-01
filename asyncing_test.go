package asyncing

import (
	"fmt"
	"testing"
	"time"
)

func Test_Hello(t *testing.T) {
	t.Run("test hello", func(t *testing.T) {
		msgns := make(chan string, 3)
		go hello(msgns)
		go hello(msgns)
		go hello(msgns)

		go func() {
			msgns <- "Hello Elixir"
		}()

		msgns <- "Hello Go"
		msgns <- "Hello Ruby"

		time.Sleep(1 * time.Second)
		close(msgns)
	})

	t.Run("test input and output", func(t *testing.T) {
		langs := make(chan string, 3)
		result := make(chan string, 6)

		go worker(langs, result)
		go worker(langs, result)

		langList := []string{"Elixir", "Go", "Swift", "Kotlin", "Ruby", "Javascript"}
		for _, lang := range langList {
			langs <- lang
		}

		for msg := range result {
			fmt.Println(msg)
		}

		time.Sleep(1 * time.Second)
		close(langs)
		close(result)
	})

	t.Run("test fan in", func(t *testing.T) {
		in := make(chan string)
		go helloIn(in)

		go func() {
			in <- "Hello Elixir"
			in <- "Hello Go"
		}()

		go func() {
			in <- "Hello Ruby"
		}()

		time.Sleep(1 * time.Second)
		close(in)
	})

	t.Run("test fan out", func(t *testing.T) {
		in := make(chan string)
		outA := make(chan string)
		outB := make(chan string)

		go func() {
			in <- "Hello Elixir"
			in <- "Hello Go"
		}()

		go func() {
			in <- "Hello Ruby"
			in <- "Goodbye PHP"
		}()

		go helloOut(in, outA, outB)
		go helloIn(outA)
		go helloIn(outB)

		time.Sleep(1 * time.Second)
		close(in)
		close(outA)
		close(outB)
	})
}
