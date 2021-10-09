package lessons

import "fmt"

func Lesson4() {
	// Arrays and Slices

	grades := [3]int{97, 46, 43}
	fmt.Printf("Grades: %v\n", grades)

	grades2 := [...]int{97, 46, 43, 5}
	grades2[0] = 9
	fmt.Printf("Grades2: %v\n", grades2)
	fmt.Printf("Grades2 #1: %v\n", grades2[1])
	fmt.Printf("Size of Grades2: %v\n", len(grades2))
	var matrix [3][3]int
	matrix[0] = [3]int{1, 0, 0}
	matrix[1] = [3]int{0, 1, 0}

	fmt.Println(matrix)

	// array
	a := [...]int{1, 2, 3}
	b := &a
	b[0] = 10
	fmt.Println(a)
	fmt.Println(b)
	// slice
	a1 := []int{1, 2, 3, 4, 5, 6}
	b1 := a1
	b1[0] = 10
	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

	// works for arrays too
	b1 = a1[:]
	c1 := a1[3:]
	d1 := a1[3:5]
	d1[1] = 123
	fmt.Println(b1)
	fmt.Println(c1)
	fmt.Println(d1)

	{
		a := make([]int, 3, 100)
		a = append(a, 2)
		a = append(a, 3, 4, 5)
		a = append(a, []int{6, 7, 8, 9}...)
		fmt.Println(a)
		fmt.Printf("Length: %v\n", len(a))
		fmt.Printf("Capacity: %v\n", cap(a))
		b := a[:len(a)-1]
		fmt.Println(b)
		c := append(a[:4], a[6:]...)
		fmt.Println("still linked")
		fmt.Println(a, " - a")
		fmt.Println(c, " - c")
	}
}
