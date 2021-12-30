package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgTimePerLayer int) int {
	if avgTimePerLayer == 0 {
		avgTimePerLayer = 2
	}
	return len(layers) * avgTimePerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	noodles = 0
	sauce = 0.0
	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodles = noodles + 50
		}
		if layers[i] == "sauce" {
			sauce = sauce + 0.2
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) []string {
	newList := make([]string, len(friendsList))
	copy(newList, myList)
	newList = append(newList, friendsList[len(friendsList)-1])
	return newList
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaledRecipe := make([]float64, len(quantities))
	for i := 0; i < len(quantities); i++ {
		scaledRecipe[i] = (quantities[i] / 2) * float64(portions)
	}
	return scaledRecipe
}
