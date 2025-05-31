package main

import (
	"fmt"
)

type lasagnaPreparedInfos struct {
	remainingTime int
	layersTime    int
}

func main() {
	var info lasagnaPreparedInfos = ElapsedTime(10, 30)

	fmt.Println("Remaining time is:", info.remainingTime)
	fmt.Println("Layers time is:", info.layersTime)
	
	var spent int = info.remainingTime + info.layersTime 
	fmt.Println("The time spent is:", spent)
}

func RemainingOvenTime(actual int) int {
	var ovenTime int = 40
	return ovenTime - actual 
}

func PreparationTime(numberOfLayers int) int {
	var eachLayerTime int = 2
	return numberOfLayers * eachLayerTime
}

func ElapsedTime(numberOfLayers, actualMinutesInOven int) lasagnaPreparedInfos {
	var remainingOven int = RemainingOvenTime(actualMinutesInOven)
	var preparationTime int = PreparationTime(numberOfLayers)

	var result lasagnaPreparedInfos = lasagnaPreparedInfos{
		remainingTime: remainingOven,
		layersTime:    preparationTime,
	}

	return result
}