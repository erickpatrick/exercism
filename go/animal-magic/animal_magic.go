package chance

import "math/rand"

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	return rand.Intn(20) + 1
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	return rand.Float64() * float64(rand.Intn(12)+1)
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	panic("Please implement the ShuffleAnimals function")
}
