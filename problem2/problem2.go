package main

import (
	"fmt"
	"math"
)

func GetMinMax(numbers ...*int) (min int, max int) {
    if len(numbers) == 0 {
        return 0, 0
    }
    //min = *numbers[0]
    //max = *numbers[0]
	min = math.MaxInt64
	max = math.MinInt64

	for _,nilai := range numbers{
		
		if *nilai > max{
			max =*nilai
		}
		
		if *nilai <min{
			min =*nilai
		}
	}
	return min,max
	}

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)
	min, max = GetMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println("Nilai min ", min)
	fmt.Println("Nilai max ", max)
}
