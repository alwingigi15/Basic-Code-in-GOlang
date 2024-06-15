package main
import "fmt"

func main(){
  var row,col int

  var m[10][10] int
  var tran[10][10] int

  fmt.Print("enter the limit of Row:")
  fmt.Scan(&row)
  fmt.Print("enter the limit of column:")
  fmt.Scan(&col)

  fmt.Println("Enter your Array:")
  for i:=0;i<row;i++{
    for j:=0;j<col;j++{
      fmt.Scan(&m[i][j])
    }
  }
  for i:=0;i<row;i++{
    for j:=0;j<col;j++{
      tran[j][i]=m[i][j]
      }
  }
  fmt.Println("The Transpose Array:")
  for i:=0;i<col;i++{
    for j:=0;j<row;j++{
      fmt.Print(tran[i][j]," ")
}
fmt.Println()
}
}
