package main
import "fmt"

func main(){
	var i int
	
	var marr1[5] int
	var marr2[5] int
	var  marr3[5] int

	fmt.Print("Enter Your First Array:")
	for i=0; i<len(marr1); i++{
		fmt.Scan(&marr1[i])
	}
	fmt.Print("Enter Your Second Array:")
	for i=0 ; i<len(marr2);i++{
		fmt.Scan(&marr2[i])
	} 
		fmt.Println("Here Your Multiplied Arrays:")
		for i=0;i<len(marr3);i++{
			marr3[1]=marr1[i] * marr2[i]
			fmt.Print(marr3[i]," ")
		}
}