package lasagnamaster

func PreparationTime(layers []string, averageTime int) int {
	if averageTime == 0 {
		averageTime = 2
	}

	return len(layers) * averageTime
}

func Quantities(layers []string) (noodleQty int, sauceQty float64) {
	for _, layer := range layers {
		switch layer {
		case "sauce":
			sauceQty += 0.2
		case "noodles":
			noodleQty += 50
		default:
			continue
		}
	}

	return
}

func AddSecretIngredient(friendsIngredients []string, myIngredients []string) {
	friendIngSize := len(friendsIngredients)
	myIngSize := len(myIngredients)

	myIngredients[myIngSize-1] = friendsIngredients[friendIngSize-1]
}

func ScaleRecipe(twoPortionAmount []float64, numberOfPortions int) (quantities []float64) {
	quantities = make([]float64, len(twoPortionAmount))

	for key, amount := range twoPortionAmount {
		quantities[key] = (amount / 2.0) * float64(numberOfPortions)
	}

	return quantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
