package domain

import "fmt"

type Person struct {
	Name        string
	Age         int
	Weight      float64
	alive       bool
	ParentsList *Parents
}

/*
	comentarios
	comentarios
*/

// comentario
type Parents []string

func (p *Person) GetName(name string) string {
	return p.Name + name
}

func (p *Person) IsAlive() bool {
	return p.alive
}

func IMC(heigh float64, weight float64) float64 {
	if heigh == weight {
		return 0
	}

	switch heigh {
	case 0.1:
		fmt.Print(heigh)
	default:
		fmt.Print("")
	}

	name := "name"
	age := 11

	fmt.Print(name, age)

	return 0.3
}
