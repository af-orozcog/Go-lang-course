package main

import "fmt"

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (animal *Animal) Eat() string {
	return animal.food
}

func (animal *Animal) Move() string {
	return animal.locomotion
}

func (animal *Animal) Speak() string {
	return animal.noise
}

func main() {
	animals := map[string]Animal{
		"cow":   Animal{"grass", "walk", "moo"},
		"bird":  Animal{"worms", "fly", "peep"},
		"snake": Animal{"mice", "slither", "hsss"},
	}
	for {
		var name, method string
		fmt.Printf(">")
		fmt.Scanf("%s %s", &name, &method)
		temp := animals[name]
		if method == "eat" {
			fmt.Println(temp.Eat())
		} else if method == "speak" {
			fmt.Println(temp.Speak())
		} else {
			fmt.Println(temp.Move())
		}
	}
}
