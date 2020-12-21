package main

import (
	"fmt"
	"time"
)

type Dog struct {
	height    int
	happiness int
	name      string
	collar    func() string
}

type Cat struct {
	height    int
	sassiness int
	name      string
}

type Animal interface {
	makeNoise() string
	animalFacts() string
	getHeight() int
}

func aprilCollar() string {
	return "asdbjhf"
}

func main() {
	rex := Dog{5,
		2234523452,
		"rex",
		func() string {
			if time.Now().Month() == time.October {
				return "blue"
			}
			return "red"
		}}
	tibbles := Cat{3, 86, "tibbles"}

	info(rex)
	fmt.Println()
	info(tibbles)

	fmt.Println(rex.collar())

	cherry := Cat{3, 86, ""}
	moss := Cat{3, 86, ""}

	if cherry == moss {
		fmt.Println("Cherry and Moss are equal")
	}

	leo := &Cat{3, 86, ""}
	jo := &Cat{3, 86, ""}

	if leo == jo {
		fmt.Println("Leo and Jo are equal")
	}

	if *leo == *jo {
		fmt.Println("Leo and Jo's pointers  are equal")
	}
}

func info(ani Animal) {
	fmt.Printf("this animal goes: %v\n", ani.makeNoise())
	fmt.Println(ani.animalFacts())
	if ani.getHeight() < 5 {
		fmt.Printf("Tiny size, only %dkm tall\n", ani.getHeight())
	}
}

func (d Dog) makeNoise() string {
	return "bork bork"
}

func (c Cat) makeNoise() string {
	return "*silent stare*"
}

func (d Dog) animalFacts() string {
	return fmt.Sprintf("%v is the happiest dog in the world, with a happiness level of: %v", d.name, d.happiness)
}

func (c Cat) animalFacts() string {
	return fmt.Sprintf("%v is the sassiest cat in the world, with a sassiness level of: %v", c.name, c.sassiness)
}

func (d Dog) getHeight() int {
	return d.height
}

func (c Cat) getHeight() int {
	return c.height
}
