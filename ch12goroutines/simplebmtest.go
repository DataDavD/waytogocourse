package main

import (
	"fmt"
	"testing"
)

func main() {
	fmt.Println(" sync", testing.Benchmark(BenchmarkChannelSync).String())
	fmt.Println("buffered", testing.Benchmark(BenchmarkChannelBuffered).String())
}

func BenchmarkChannelSync(b *testing.B) { // makes buffered channel
	ch := make(chan int)
	go func() {
		for i := 0; i < b.N; i++ {
			ch <- i
		}
		close(ch)
	}()
	for _ = range ch { // iterating over channel without doing anything
	}
}

func BenchmarkChannelBuffered(b *testing.B) { // makes buffered channel with capacity of 128
	ch := make(chan int, 128)
	go func() {
		for i := 0; i < b.N; i++ {
			ch <- i
		}
		close(ch)

	}()
	for _ = range ch {
	}
}
