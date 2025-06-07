package main

import "fmt"

type Character struct {
	Name, Role, Element string
	Level, exp          int
}

type Monster struct {
	Type      string
	Boss      bool
	Health    int
	Difficult int
	Habitat   string
}

// ---------------------------- Struct Method
func (char Character) levelUp() {
	fmt.Println(char.Name, "level is increased ++", char.Level+10)
}

func (char Character) pressQ() {
	fmt.Println(char.Name, "is unleashing liberation!")
}

func main() {
	var party1 Character
	fmt.Println(party1)

	party1.Name = "Changli"
	party1.Role = "Fighter"
	party1.Element = "Fusion"
	party1.Level = 80

	// ---------------------------- Struct Literals
	party2 := Character{
		Name:    "Shorekeeper",
		Role:    "Support",
		Element: "Spectro",
		Level:   70,
	}

	party3 := Character{"Zani", "Tank", "Spectro", 80, 0}

	fmt.Println(party1)
	fmt.Println(party2)
	fmt.Println(party3)
	fmt.Println(party3.Name)
	fmt.Println(party3.Role)

	party1.pressQ()
	party3.pressQ()
	party2.levelUp()
}

// ---------------------------- Struct Tag

// `key:"value" key2:"value2"`

type User struct {
	ID    int    `json:"id" db:"user_id"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email,omitempty" db:"email_address"`
}
