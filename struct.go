package main

import "fmt"

type Character struct {
	Weapon, Role, Element string
	Level, exp            int
}

type Monster struct {
	Type      string
	Boss      bool
	Health    int
	Difficult int
}

func main() {
	var changli Character
	fmt.Println(changli)

	changli.Weapon = "Sword"
	changli.Role = "Fighter"
	changli.Element = "Fusion"
	changli.Level = 80

	// ---------------------------- Struct Literals
	shorekeeper := Character{
		Role:    "Support",
		Element: "Spectro",
		Level:   70,
	}

	encore := Character{"Catalyst", "Mage", "Fusion", 80, 4000}
	zani := Character{"Gauntlet", "Tank", "Spectro", 80, 0}

	fmt.Println(changli)
	fmt.Println(shorekeeper)
	fmt.Println(encore)
	fmt.Println(zani)
	fmt.Println(zani.Role)
}
