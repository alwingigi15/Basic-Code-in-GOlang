package main
import "fmt"

func main(){
  var size int
  fmt.Println("Enter your limit:")
  fmt.Scanln(&size)

  arr:=make([]int,size)

  fmt.Println("Enter th array:")
  for i:=0;i<size;i++{
    fmt.Scanln(&arr[i])
  }
  Rev(arr,size)
}
func Rev(arr[] int,size int){
  rarr:=make([]int,size)
  j:=0
  for i:= size - 1; i >= 0; i-- {
       rarr[j] = arr[i]
       j++
   }

   fmt.Println("\nThe Result of a Reverse Array = ", rarr)
}
