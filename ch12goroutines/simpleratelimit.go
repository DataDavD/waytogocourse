package main

const (
	AvailableMemory         = 10 << 20                                  // 10 MB, for example
	AverageMemoryPerRequest = 10 << 10                                  // 10 KB
	MAXREQS                 = AvailableMemory / AverageMemoryPerRequest // here amounts to 1000
)

var sem = make(chan int, MAXREQS)

type Request struct {
	a, b   int
	replyc chan int
}

func process(r *Request) {
	// Do something
	// May take a long time and use a lot of memory or CPU
}

func handle(r *Request) {
	process(r)
	<-sem // signal done: enable next request to start by making 1 empty place in the buffer
}

func Server(queue chan *Request) {
	for {
		sem <- 1 // blocks when channel is full (1000 requests are active)

		// so wait here until there is capacity to process a request
		// (doesn't matter what we put in it)
		request := <-queue
		go handle(request)
	}
}

func main() {
	queue := make(chan *Request)
	go Server(queue)
}
