package main

import (
	"animal_magic/animal_magic"
	"fmt"
)

func main() {
	die := chance.RollADie()
	animals_shuffled := chance.ShuffleAnimals()
	fmt.Println(die)
	fmt.Println(animals_shuffled)
}
