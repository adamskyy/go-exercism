package dndcharacter

import "math/rand"

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

func floorDiv(a, b int) int {

	q := a / b
	r := a % b
	if r != 0 && (a < 0) != (b < 0) {
		q--
	}
	return q

}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return floorDiv(score - 10, 2)
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
    min := 6
    sum := 0
    for i := 0; i < 4; i++ {
        rollDice := rand.Intn(6) + 1
        if rollDice < min {
            min = rollDice
        }
        sum += rollDice
    }
    sum -= min
    return sum
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	exampleChar := Character{
        Strength: Ability(),
        Dexterity: Ability(),
        Constitution: Ability(),
        Intelligence: Ability(),
        Wisdom: Ability(),
        Charisma: Ability(),
        Hitpoints: 0,
    }
    exampleChar.Hitpoints = Modifier(exampleChar.Constitution) + 10
    return exampleChar
}
