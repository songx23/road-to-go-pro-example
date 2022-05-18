package main

import (
	"fmt"

	"github.com/songx23/road-to-go-pro-example/equation"
)

func main() {
	linearAnswer := equation.SolveLinear(1, 2)
	fmt.Println(linearAnswer)

	quadraticAnswer1, quadraticAnswer2 := equation.SolveQuadratic(1, -1, -6)
	fmt.Println(quadraticAnswer1, quadraticAnswer2)
}
