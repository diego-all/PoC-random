package main

import (
	"fmt"
	"os"
)

func main() {
	nombre := os.Getenv("NOMBRE")
	fmt.Println("Hola,", nombre, "!")
}
