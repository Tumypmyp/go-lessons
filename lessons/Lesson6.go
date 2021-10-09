package lessons

import "fmt"

func Lesson6() {
	// If and Switch statements
	if pop, ok := map[string]int{"Anya": 20, "Olya": 17}["Olya"]; ok {
		fmt.Println(pop)
	}
	a := 5
	b := 7
	c := 10
	fmt.Println(a <= b, a >= b, a != b)
	if a < b && a < c {
		fmt.Println("a")
	} else if a > b {
		fmt.Println("b")
	}
	switch i := 2 + 3; i {
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	case 2, 4, 6:
		fmt.Println("two, four or six")
	default:
		fmt.Println("another value")
	}
	{
		i := 10
		switch {
		case i <= 10:
			fmt.Println("less than 10")
			// explicitly run next case
			fallthrough
		case i < 20:
			fmt.Println("less than twenty")
			// out of switch
			break
		}
	}
	var i interface{} = 1
	switch i.(type) {
	case int:
		fmt.Println("i is int")
	case float64:
		fmt.Println("i is float64")
	case string:
		fmt.Println("i is string")

	}

}
