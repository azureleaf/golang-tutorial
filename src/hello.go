package main

// Doesn't need "," symbol in the package list
import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	// Without using the rand.Seed(), this returns the identical number
	fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Println("The time is", time.Now())

	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

}
