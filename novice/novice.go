package novice

import "fmt"

type INovice interface {
	DisplayInfo()
	DeleteHealth(damage int)
}

type Novice struct {
	name      string
	classname string
	health    int
}

// Constructor
func NewNovice(name, classname string, health int) INovice {
	return &Novice{
		name:      name,
		classname: classname,
		health:    health,
	}
}

func (n *Novice) DisplayInfo() {
	fmt.Println("Player Name:", n.name)
	fmt.Println("Player Classname:", n.classname)
	fmt.Println("Player Health:", n.health)
}

func (n *Novice) DeleteHealth(damage int) {
	if damage >= n.health {
		n.health = 0
	} else {
		n.health -= damage
	}
}
