package main
import "fmt"

func main(){
var row,col int
var a[10][10] int
var b[10][10] int


fmt.Println("Enter the limit of the Row")
fmt.Scanln(&row)
fmt.Println("enter The Limit of the column")
fmt.Scanln(&col)

fmt.Println("Enter the First Array:")
    for i:=0;i<row;i++{
      for j:=0;j<col;j++{
        fmt.Scanln(&a[i][j])
      }
    }
    fmt.Println("Enter the Second Array:")
        for i:=0;i<row;i++{
          for j:=0;j<col;j++{
            fmt.Scanln(&b[i][j])
          }
        }
        fmt.Println("Added Array:")
            for i:=0;i<row;i++{
              for j:=0;j<col;j++{

                fmt.Println(a[i][j]+b[i][j])

              }

}
}
