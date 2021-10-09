package lessons

import (
	"fmt"
	"reflect"
	_ "reflect"
)

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

type Animal struct {
	Name   string `required max:"100"`
	Origin string
}
type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

func Lesson5() {
	// Maps and Structs
	countByName := map[string]int{
		"Lola":   12,
		"Stitch": 2,
		"Vova":   15,
	}
	// reference
	cbn := countByName

	countByName["Tanya"] = 14
	delete(countByName, "Vova")
	fmt.Println(countByName["Lola"], " - Lola")
	fmt.Println(len(countByName))
	fmt.Println(countByName)
	count, ok := countByName["Masha"]
	fmt.Println(count, ok)
	fmt.Println(cbn)
	//error n := map[[]int]string{}
	m := map[[3]int]string{}
	m = make(map[[3]int]string)
	fmt.Println(m)

	aDoctor := Doctor{
		number:     3,
		actorName:  "Jon Pertwee",
		companions: []string{"Liz Shaw", "Jo Grant"},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.companions[1])

	// anonymous struct
	aUser := struct{ name string }{name: "John Pertwee"}
	fmt.Println(aUser)
	// a copy, a value type
	// a reference := &aUser
	aUser2 := aUser
	aUser2.name = "Tom Baker"
	fmt.Println(aUser2)

	// embedding, composition over inheritance
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b)
	fmt.Println(b.Name)
	c := Bird{
		Animal:   Animal{Name: "Took", Origin: "Alyaska"},
		SpeedKPH: 20,
		CanFly:   true,
	}
	fmt.Println(c)

	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
