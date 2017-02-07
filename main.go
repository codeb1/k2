package main

import "log"
import "os"
import "io"
import "sync"

var wg = new(sync.WaitGroup)

func worker(s int, c chan int) {
	log.Printf("in worker %d\n", s)
	c <- s
	wg.Done()
}

func reader(c chan int) {
	log.Printf("In reader\n")
	for true {
		s := <-c
		log.Printf("Read %d\n", s)
	}
}
func setLog(logpath string) {
	file, err := os.Create(logpath)
	if err != nil {
		panic(err)
	}
	log.SetOutput(io.Writer(file))
}
func main() {
	setLog("output.log")
	c := make(chan int)
	go reader(c)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(i, c)
	}
	for i := 0; i < 10; i++ {
	}
	wg.Wait()
}
