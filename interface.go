package main
import "fmt"

type employee interface{
	printname(name string)
	printsalary(b int ,t int)int
}
 type emp int 

 func(e emp) printname(name string){
	 fmt.Println("ID:",e)
	 fmt.Println("Employee Name:",name)
 }
  
 func(e emp)printsalary(b int,t int)int{
	   s:=(b*t)/100
	   return b-s 
 }
 func main(){
	 var emp1 employee
	 emp1=emp(1)
	 emp1.printname("ALWIN GIGI")
	fmt.Println("Employee Salary:", emp1.printsalary(2500,5))

 }