package main

import "fmt"

type Character struct {
	Name, Weapon, Element string
	Level, Exp            int
}

// Pointer to char
func (char *Character) levelUp() {
	char.Exp = 0
	char.Level += 10
	fmt.Println(char.Name, "level is increased ++")
}

func (char *Character) changeWeapon(weapon string) {
	char.Weapon = weapon
	fmt.Println(char.Name, "weapon changed")
}

func main() {

	party1 := Character{"Changli", "Sword", "Fusion", 80, 83792}
	party2 := Character{"Zani", "Gauntlet", "Spectro", 80, 232222}
	party3 := Character{"Shorekeeper", "Catalyst", "Spectro", 70, 2000}

	fmt.Println(party1)
	fmt.Println(party2)
	fmt.Println(party3)

	party3.levelUp()
	party3.levelUp()
	fmt.Println(party3)

	party2.changeWeapon("Signature")
	fmt.Println(party2)
}

// * pointer is very recommended to use in method (struct function)
