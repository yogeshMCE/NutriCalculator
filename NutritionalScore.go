package main

type ScoreType int

const (
	Food ScoreType = iota
	Beverage
	Water
	Cheese
)

type NutritionalScore struct {
	Value     int
	Positive  int
	Negative  int
	ScoreType ScoreType
}

type EnergyKj float64

type SugarGram float64

type SaturatedfattyAcids float32

type SodiumMiliGram float64

type FruitsPercents float64

type FiberGram float64

type ProtienGram float64

type NutritionalData struct {
	Energy             EnergyKj
	Sugar              SugarGram
	SaturatedfattyAcid SaturatedfattyAcids
	Sodium             SodiumMiliGram
	Fruits             FruitsPercents
	Fiber              FiberGram
	Protien            ProtienGram
	IsWater            bool
}

var ScoreToLetter = []string{"A", "B", "C", "D", "E"}

var Energylevels = []float64{3350, 3015, 2680, 2345, 2010, 1675, 1340, 1005, 670, 335}
var Sugarlevels = []float64{45, 60, 36, 31, 27, 22.5, 18, 13.5, 9, 4.5}
var saturatedfattyAcidLevels = []float64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
var SodiumLevels = []float64{900, 810, 720, 630, 540, 450, 360, 270, 180, 90}
var FiberLevels = []float64{4.7, 3.7, 2.8, 1.9, 0.9}
var ProtienLevels = []float64{8, 6.4, 4.8, 3.2, 1.6}

// for beverages
var EnergyLevelsBeverage = []float64{270, 240, 210, 180, 150, 120, 90, 60}
var SugarLevelsBeverage = []float64{13.5, 12, 10.5, 9, 7.5, 6, 4.5, 3, 1.5, 0}

func (e *EnergyKj) GetPoints(st ScoreType) int {
	if st == Beverage {
		return GetPointsFromRange(float64(*e), EnergyLevelsBeverage)
	}
	return GetPointsFromRange(float64(*e), Energylevels)
}
func (s *SugarGram) GetPoints(st ScoreType) int {
	if st == Beverage {
		return GetPointsFromRange(float64(*s), SugarLevelsBeverage)
	}
	return GetPointsFromRange(float64(*s), Sugarlevels)
}
func (fp *FruitsPercents) GetPoints(st ScoreType) int {
	if st == Beverage {
		if *fp > 80 {
			return 10
		} else if *fp > 60 {
			return 4
		} else if *fp > 40 {
			return 2
		}
		return 0
	}
	if *fp > 80 {
		return 5
	} else if *fp > 60 {
		return 2
	} else if *fp > 40 {
		return 1
	}
	return 0

}
func (sfa *SaturatedfattyAcids) GetPoints(st ScoreType) int {
	return GetPointsFromRange(float64(*sfa), saturatedfattyAcidLevels)
}
func (sm *SodiumMiliGram) GetPoints(st ScoreType) int {
	return GetPointsFromRange(float64(*sm), SodiumLevels)
}
func (fg *FiberGram) GetPoints(st ScoreType) int {
	return GetPointsFromRange(float64(*fg), FiberLevels)
}
func (pg *ProtienGram) GetPoints(st ScoreType) int {
	return GetPointsFromRange(float64(*pg), ProtienLevels)
}
func EnergyFromKcal(Kcal float64) EnergyKj {

	return EnergyKj(Kcal * 4.184)
}
func SodiumFromSalt(saltmg float64) SodiumMiliGram {
	return SodiumMiliGram(saltmg / 2.5)
}
func GetPointsFromRange(value float64, steps []float64) int {
	stepslen := len(steps)
	for i, l := range steps {
		if value > l {
			return stepslen - i
		}
	}
	return 0
}

func (ns NutritionalScore) GetNutriScore() string {
	if ns.ScoreType == Food {
		return ScoreToLetter[GetPointsFromRange(float64(ns.Value), []float64{18, 10, 2, -1})]
	}
	if ns.ScoreType == Water {
		return ScoreToLetter[0]
	}
	return ScoreToLetter[GetPointsFromRange(float64(ns.Value), []float64{9, 5, 1, -2})]

}
func GetNutritionalScore(n NutritionalData, st ScoreType) NutritionalScore {
	value := 0
	positive := 0
	negative := 0
	if st != Water {
		FruitsPoint := n.Fruits.GetPoints(st)
		FiberPoints := n.Fiber.GetPoints(st)
		negative = n.Energy.GetPoints(st) + n.SaturatedfattyAcid.GetPoints(st) + n.Sodium.GetPoints(st) + n.Sugar.GetPoints(st)
		positive = FruitsPoint + FiberPoints + n.Protien.GetPoints(st)

		if st == Cheese {
			value = negative - positive
		} else {

			if negative >= 11 && FruitsPoint < 5 {
				value = positive - negative - FruitsPoint
			} else {
				value = negative - positive
			}

		}

	}

	return NutritionalScore{
		Value:     value,
		Positive:  positive,
		Negative:  negative,
		ScoreType: st,
	}

}
