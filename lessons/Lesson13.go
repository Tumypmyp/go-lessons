package lessons

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func Lesson13() {
	fmt.Println("example1:")
	example1()
	fmt.Println("example2:")
	example2()
	fmt.Println("example3:")
	example3()

}

func example1() {
	ch := make(chan int)
	wg.Add(2)
	go func() {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()
	go func() {
		i := 42
		ch <- i
		i = 27
		wg.Done()
	}()
	wg.Wait()
}
func example2() {
	ch := make(chan int)
	for j := 0; j < 5; j++ {
		wg.Add(2)
		// receive-only channel
		go func(ch <-chan int) {
			i := <-ch
			fmt.Println(i)
			wg.Done()
		}(ch)

		// send-only channel
		// only one message in the channel
		go func(ch chan<- int) {
			ch <- 42
			wg.Done()
		}(ch)
	}
	wg.Wait()
}

func example3() {
	// buffer is 50 integers
	ch := make(chan int, 50)

	wg.Add(2)

	go func(ch <-chan int) {
		//if i, ok := <-ch; ok {

		for i := range ch {
			//i := <-ch
			fmt.Println(i)
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()
}

var doneCh = make(chan struct{})

func example4() {

	doneCh <- struct{}{}
}
//
//func logger() {
//	for {
//		select {
//		case entry := <-logCh:
//			fmt.Println("asdf")
//		case <-doneCh:
//			break
//		}
//	}
//
//}
