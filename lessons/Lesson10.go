package lessons

import (
	"errors"
	"fmt"
)

func Lesson10() {
	a := "hello world"
	sendMessage(&a)
	fmt.Println(a)
	sayGreeting("Hello", "Stacey")
	sum("sum is", 1, 2, 3, 4, 5)
	fmt.Println(divide(2, 0))

	func() {
		fmt.Println("hello world")
	}()

	for i := 0; i < 5; i++ {
		func(i int) {
			fmt.Println(i)
		}(i)
	}

	var f func() = func() {
		fmt.Println("staff")
	}
	f()

	g := greeter{
		greeting: "hello",
		name:     "go",
	}
	g.greet()
}

type greeter struct {
	greeting string
	name     string
}

//func (g *greeter) greet() { or copy:
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)

}

func sendMessage(msg *string) {
	fmt.Println(*msg)
	*msg += "asdf"
}
func sayGreeting(greeting, name string) {
	fmt.Println(greeting, name)
}

func sum(msg string, values ...int) {
	fmt.Println(values)
	sum := 0
	for _, v := range values {
		sum += v
	}
	fmt.Println(msg, sum)
}
func sum2(a int, b int) *int {
	c := a + b
	return &c
}
func sum3(a int, b int) (result int) {
	result = a + b
	return
}
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0.0, errors.New("cannot divide through zero")
	}

	return a / b, nil
}

//
//func divide(a, b float64) (float64, err) {
//	if b == 0 {
//		return 0, fmt.Errorf("Cannot divide by zero")
//	}
//
//	return a / b, nil
//}
