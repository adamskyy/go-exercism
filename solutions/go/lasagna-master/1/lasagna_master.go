package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avg int) int {
    if avg == 0 {
        avg = 2
    }
    return len(layers) * avg
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
    sauce := 0.0
    grams := 0
    for i := 0; i < len(layers); i++ {
        if layers[i] == "noodles" {
            grams += 50
        } else if layers[i] == "sauce" {
            sauce += 0.2
        }
    }
    return grams, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friends []string, myList []string) []string {
    myList[len(myList) - 1] = friends[len(friends) - 1]
    return myList
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
    newSlice := make([]float64, len(quantities))
    for i := 0; i < len(newSlice); i++ {
        newSlice[i] = quantities[i] / 2.0 * float64(portions)
    }
    return newSlice
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
