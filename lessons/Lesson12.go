package lessons

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg1 = sync.WaitGroup{}
//var m = sync.RWMutex{}
func Lesson12(){
	// Goroutines
	go sayHello()
	time.Sleep(100 * time.Millisecond)
	var msg = "hello"
	wg.Add(1) //++
	go func(msg string){
		fmt.Println(msg)
		wg.Done() //--
	}(msg)
	msg = "goodbye"
	wg.Wait()
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))
	runtime.GOMAXPROCS(1)
	// go run -race .
}

func sayHello(){
	fmt.Println("Hello")
}


