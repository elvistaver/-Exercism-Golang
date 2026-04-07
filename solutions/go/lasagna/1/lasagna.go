package lasagna

// TODO: define the 'OvenTime' constant
const OvenTime = 40

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
    RemainingOvenTime:= OvenTime - actualMinutesInOven
    return RemainingOvenTime

	//panic("RemainingOvenTime not implemented")
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
    Time:= 2
    preparationtime:= numberOfLayers * Time
    return preparationtime
	//panic("PreparationTime not implemented")
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
    time:= 2
    elapsedtime:= (time * numberOfLayers) + actualMinutesInOven
    return elapsedtime
	//panic("ElapsedTime not implemented")
}
