package dndcharacter

import (
	"math"
	"math/rand/v2"
	"slices"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return int(math.Floor(float64(score-10) / 2.0))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	rolls := []int{rand.IntN(5) + 1, rand.IntN(5) + 1, rand.IntN(5) + 1, rand.IntN(5) + 1}
	slices.Sort(rolls)

	total := 0
	for key, value := range rolls {
		if key == 0 {
			continue
		}
		total += value
	}
	return total
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	panic("Please implement the GenerateCharacter() function")
}
