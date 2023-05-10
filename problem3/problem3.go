package main

import (
	"fmt"
	"math"
)

type Student struct {
	name  []string
	score []int
}

func (s Student) Average() float64 {
	// your code here
	var sum int
	for _,value := range s.score{
		sum +=value
	} 
	avr := float64(sum)/float64(len(s.score))
	return avr
}

func (s Student) Min() (min int, name string) {
	// your code here
	if len(s.score) == 0 {
		return 0,""
	}
    min = math.MaxInt64

	for i, num := range s.score {
		if num < min {
			min = num
			name = s.name[i]
		}
	}
	return min,name
}

func (s Student) Max() (max int, name string) {
	// your code here
	if len(s.score) == 0 {
		return 0,""
	}
	max = math.MinInt64

	for i, num := range s.score {
		if num > max {
			max = num
			name = s.name[i]
		}
	}
	return max,name
}

func main() {
	var a = Student{}

	for i := 0; i < 6; i++ {
		var name string
		fmt.Print("Input " + fmt.Sprint(i) + " Student's Name : ")
		fmt.Scan(&name)
		a.name = append(a.name, name)
		var score int
		fmt.Print("Input " + name + " Score : ")
		fmt.Scan(&score)
		a.score = append(a.score, score)
	}

	fmt.Println("\n\nAverage Score Students is", a.Average())
	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score Students is : "+nameMax+" (", scoreMax, ")")
	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score Students is : "+nameMin+" (", scoreMin, ")")
}
