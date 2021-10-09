package lessons

import "fmt"

func Lesson7() {
	// Looping

	for i := 0; i < 5; i += 2 {
		fmt.Printf("%v ", i)
	}
	fmt.Println()

	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Printf("%v %v, ", i, j)
	}
	fmt.Println()

	i := 0
	for ; i < 5; i++ {
		if i == 3 {
			continue
		} else if i == 5 {
			break
		}
		fmt.Printf("%v ", i)
	}
	fmt.Println()

	k := 0
	for {
		k++
		if k == 5 {
			break
		}
	}

Loop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("%v ", i*j)
			if i*j >= 3 {
				break Loop
			}
		}
	}
	fmt.Println()
	//s := "hello world"
	//s := [3]int{1, 2, 3}
	s := []int{1, 2, 3, 4}
	for k, v := range s {
		fmt.Printf("%v %v, ", k, v)
	}
	fmt.Println()

}
