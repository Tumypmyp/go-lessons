package lessons

import "fmt"

const (
	error = iota
	cat
	dog
	dragon
)

func Lesson3() {
	const myConst = 3
	fmt.Printf("%v %T\n", myConst, myConst)
	var num int32 = 10
	fmt.Printf("%v %T\n", myConst+num, myConst+num)

	var animalType int = cat
	fmt.Printf("%v - animal\n", animalType == cat)
}
