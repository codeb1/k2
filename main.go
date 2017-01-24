package main

import "log"
import "os"
import "io"
import "sync"

var wg = new(sync.WaitGroup)
func worker(s int){
//	setLog("output%s.log",s)
     log.Printf("in worker %s\n",s)
	 log.Printf("something")
    wg.Done()
}
func setLog(logpath string){
	file, err := os.Create(logpath)
	if err	!= nil {
		panic(err)
	}
	log.SetOutput(io.Writer(file))
}
func main(){
	setLog("output.log")
	log.Printf("hello jacob \n")
	for  i:=0; i < 10;  i ++ {
		wg.Add(1)
		go worker(i)
	}
	wg.Wait()
}
