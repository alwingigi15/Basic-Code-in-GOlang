package customer

import (
	"fmt"
	"sort"
)

func Sorting() {
	val := []int{10, 98, 63, 27, 2, 17, 99}
	x := 27
	fmt.Println(sort.IntsAreSorted(val)) 
	fmt.Println("VALUE BEFORE SORTING:", val)

	sort.Ints(val)
	fmt.Println("SORTED VALUES ARE:", val)
	fmt.Println(sort.IntsAreSorted(val)) 

	Pos := sort.SearchInts(val, x)
	fmt.Println("FOUNDED IN:", x, Pos, val)

	name := []string{"varun", "Dias", "Pushpa", "Rocky", "Alwin"}
	name[1] = "Britto"
	fmt.Println("ENTERED NAMES:", name)
	sort.Strings(name)
	fmt.Println("SORTED NAMES:", name)

	result := []int{33, 67, 14, 6, 47}
	y := 14
	P := sort.SearchInts(result, y)
	fmt.Println("FOUNDED IN:", x, P, result)

}