package lessons

import "fmt"

func Lesson9() {
	// Pointers
	var a int = 42
	var b *int = &a
	*b = 6
	fmt.Println(a, *b)

	var ms *myStruct
	ms = &myStruct{foo: 42}
	var ms2 *myStruct = new(myStruct)
	(*ms2).foo = 3
	fmt.Println(ms, ms2)
}

type myStruct struct {
	foo int
}
