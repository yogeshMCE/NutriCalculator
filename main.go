package main

import "fmt"

func main() {
	ns := GetNutritionalScore(NutritionalData{
		Energy:             EnergyFromKcal(100),
		Sugar:              SugarGram(10),
		SaturatedfattyAcid: SaturatedfattyAcids(2),
		Sodium:             SodiumMiliGram(500),
		Fruits:             FruitsPercents(60),
		Fiber:              FiberGram(4),
		Protien:            ProtienGram(2),
	}, Food)

	fmt.Printf("The Nutritional Score is %d\n", ns.Value)
	fmt.Printf("NutriScore: %s\n", ns.GetNutriScore())
}
