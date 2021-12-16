package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
	
)

func main() {
	//rand.Seed(time.Now().UnixNano()) change the seed to change the random
	fmt.Println("Hello, world!")
	fmt.Println("The time is", time.Now())
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("I'm %g years old.\n", math.Sqrt(7))
}

