package main

import "fmt"

type Animal interface {
	Move()
	Eat()
	Speak()
}

type Cow struct {
	Locomotion string
	Spoken     string
	Food       string
}

type Bird struct {
	Locomotion string
	Spoken     string
	Food       string
}

type Snake struct {
	Locomotion string
	Spoken     string
	Food       string
}

func (an Bird) Speak() {
	fmt.Println("I say " + an.Spoken)
}

func (an Bird) Eat() {
	fmt.Println("I eat " + an.Food)
}

func (an Bird) Move() {
	fmt.Println("I move by " + an.Locomotion)
}

func (an Cow) Speak() {
	fmt.Println("I say " + an.Spoken)
}

func (an Cow) Eat() {
	fmt.Println("I eat " + an.Food)
}

func (an Cow) Move() {
	fmt.Println("I move by " + an.Locomotion)
}

func (an Snake) Speak() {
	fmt.Println("I say " + an.Spoken)
}

func (an Snake) Eat() {
	fmt.Println("I eat " + an.Food)
}

func (an Snake) Move() {
	fmt.Println("I move by " + an.Locomotion)
}

func createSnake() Snake {
	return Snake{"slither", "hsss", "mice"}
}

func createCow() Cow {
	return Cow{"walk", "peep", "grass"}
}

func createBird() Bird {
	return Bird{"fly", "peep", "worms"}
}

func main() {

	seen := make(map[string]Animal)

	for {
		fmt.Printf("> ")

		var val1, val2, val3 string

		fmt.Scanf("%s %s %s", &val1, &val2, &val3)

		if val1 == "query" {
			animal, err := seen[val2]
			if err != true {
				fmt.Println("The animal by that name wasn't found")
			} else {
				switch val3 {
				case "eat":
					animal.Eat()
				case "move":
					animal.Move()
				case "speak":
					animal.Speak()
				}
			}
		} else {
			switch val3 {
			case "snake":
				seen[val2] = createSnake()
			case "cow":
				seen[val2] = createCow()
			case "bird":
				seen[val2] = createBird()
			}
		}
	}
}
