package main

import (
	"fmt"
	"math"
	// "slices"
)

// struct of rectangle, height and width
type Rectangle struct {
	h float32
	w float32
}

// rotation can either be 0 or 90 ah should I make a type for it?? hmmmmm probably smart to do so
//type Rotation struct {
//	degree float32
//}
//

// number of shapes
var shapes = 4

// list of rectangles
// eyeballed the values from the image
// var list = [Rectangle(8.8, 9.3) Rectangle(9.2, 7.8) Rectangle(7.2, 8.9) Rectangle(9.1, 7.2)]
var listR = []Rectangle{{8.8, 9.3}, {9.2, 7.8}, {7.2, 8.9}, {9.1, 7.2}}

func main() {
	values := []int{0, 90}
	combinations := generateCombinations(shapes, values)

	// Print all combinations
	for _, combo := range combinations {
		fmt.Println(combo)
	}

	resultR := make([][]float32, 0, int(math.Pow(2, float64(shapes))))

	for combN, combo := range combinations {

		var width float32 = 0
		var height float32 = 0

		for numb, angle := range combo {
			// nice so have everything I need here, this code is hard to improve tough
			// wish I had more time to desin this better
			// very scuffed so far

			//fmt.Println(numb, angle, width, height)

			if angle == 0 {
				width = width + listR[numb].w
				if height < listR[numb].h {
					height = height + listR[numb].h
				}
			}

			// changed this here becuase 90 degrees so this is different
			// height/width change places,
			// neat trick to not make this code complex heh???
			if angle == 90 {
				width = width + listR[numb].h
				if height < listR[numb].w {
					height = height + listR[numb].w
				}
			}

		}

		fmt.Println("This is the combination: ", combN, "Width: ", width, "Height: ", height)
		resultR = append(resultR, []float32{width, height})

	}

	for _, ans := range resultR {
		fmt.Println(ans)
	}

	var minR float32
	minR = math.MaxFloat32
	var heightMinR float32

	for _, val := range resultR {
		if val[0] < minR {
			minR = val[0]
			heightMinR = val[1]
		}

	}

	fmt.Println("Min Width: ", minR, " Height: ", heightMinR)

	fmt.Println("Area: ", minR*heightMinR)
}

func minWidth(listV [][]float32) {

	var minR float32
	var heightMinR float32

	for _, val := range listV {
		if val[0] < minR {
			minR = val[0]
			heightMinR = val[1]
		}

	}

	fmt.Println("Min width: ", minR, " Heigh: ", heightMinR)
}

// probably dumb to do this this way, really dumb but oh well man
func generateCombinations(n int, values []int) [][]int {
	totalCombinations := 1 << n
	result := make([][]int, 0, totalCombinations)

	for i := 0; i < totalCombinations; i++ {
		combination := make([]int, n)
		for j := 0; j < n; j++ {
			if (i>>j)&1 == 1 {
				combination[j] = values[1]
			} else {
				combination[j] = values[0]
			}
		}
		result = append(result, combination)
	}

	return result
}

// func find the total length
//func findWidth(rec1 *Rectangle, rec2 *Rectangle, rot1 *Rotation, rot2 *Rotation) {
//
//	var width int
//
//	// lets keep it simple for now ugh so unintuitive to do it this scuffed
//	if rot1.degree == 0 {
//		width = width + rec1.w
//	}
//	if rot1.degree == 90 {
//		width = width + rec1.h
//	}
//	if rot2.degree == 0 {
//		width = width + rec1.w
//	}
//	if rot2.degree == 90 {
//		width = width + rec1.h
//	}
//    // so that covers all the "cases" for now in the pass rectangles scenario
//
//}
