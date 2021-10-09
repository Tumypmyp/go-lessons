package lessons

import "fmt"

func Lesson2() {
	var b bool = true
	fmt.Printf("%v %T\n", b, b)

	c := 10
	d := 3
	fmt.Println(c - d)
	fmt.Println(c * d)
	fmt.Println(c % d)

	fmt.Println(c == d)
	fmt.Println(c >> 2)

	var c1 complex128 = 2 + 3i
	fmt.Printf("%v %T\n", c1, c1)
	fmt.Printf("%v %T\n", real(c1), real(c1))
	fmt.Printf("%v %T\n", imag(c1), imag(c1))

	c2 := complex(2, 7)
	fmt.Printf("%v %T\n", c2, c2)

	var s string = "abcde"
	v := []byte(s)
	fmt.Printf("%v %T\n", s, s)
	fmt.Printf("%v %T\n", v, v)
}
