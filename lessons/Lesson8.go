package lessons

import "fmt"

func Lesson8() {
	// Defer, Panic and Recover
	forward()
	fmt.Println()
	oneDefer()
	fmt.Println()
	allDefer()
	fmt.Println()
	stringDefer()
	fmt.Println()
	//panicExample()
}

func forward() {
	fmt.Println("start")
	fmt.Println("middle")
	fmt.Println("end")
}
func oneDefer() {
	fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")
}
func allDefer() {
	defer fmt.Println("start")
	defer fmt.Println("middle")
	defer fmt.Println("end")
}

// parameters are saved when deferring occurs
func stringDefer() {
	s := "start"
	defer fmt.Println(s)
	s = "end"
}
func panicExample() {
	fmt.Println("start")
	panic("somethin bad happened")
	fmt.Println("end")

}
