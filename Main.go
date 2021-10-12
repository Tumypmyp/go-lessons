package main

import (
	//"./lessons"
	"github.com/Tumypmyp/go-lessons/lessons"
	"github.com/fatih/color"
)

func main() {
	c := color.New(color.Bold).Add(color.FgHiYellow)

	_, _ = c.Println("\nlesson1:")
	lessons.Lesson1()

	_, _ = c.Println("\nlesson2:")
	lessons.Lesson2()

	_, _ = c.Println("\nlesson3:")
	lessons.Lesson3()

	_, _ = c.Println("\nlesson4:")
	lessons.Lesson4()

	_, _ = c.Println("\nlesson5:")
	lessons.Lesson5()

	_, _ = c.Println("\nlesson6:")
	lessons.Lesson6()

	_, _ = c.Println("\nlesson7:")
	lessons.Lesson7()

	_, _ = c.Println("\nlesson8:")
	lessons.Lesson8()

	_, _ = c.Println("\nlesson9:")
	lessons.Lesson9()

	_, _ = c.Println("\nlesson10:")
	lessons.Lesson10()

	_, _ = c.Println("\nlesson11:")
	lessons.Lesson11()

	_, _ = c.Println("\nerror example:")
	lessons.ErrorExample()

	_, _ = c.Println("\nlesson12:")
	lessons.Lesson12()

	_, _ = c.Println("\nlesson13:")
	lessons.Lesson13()

}
