package main
import ("fmt")

type employee struct{
	eid int
	name string
	age int
	salary int
}

func main(){

	var emp1 employee
	var emp2 employee

	emp1.eid=101
	emp1.name="alwin"
	emp1.age=25
	emp1.salary=10000

	emp2.eid=102
	emp2.name="babu"
	emp2.age=24
	emp2.salary=10000

	fmt.Println("ID:",emp1.eid)
	fmt.Println("Name:",emp1.name)
	fmt.Println("AGE:",emp1.age)
	fmt.Println("SALARY:",emp1.salary)

	fmt.Println("ID:",emp2.eid)
	fmt.Println("Name:",emp2.name)
	fmt.Println("AGE:",emp2.age)
	fmt.Println("SALARY:",emp2.salary)

printemployee(emp1)

printemployee(emp2)
}
func printemployee(emp employee){
	fmt.Println("ID:",emp.eid)
	fmt.Println("Name:",emp.name)
	fmt.Println("AGE:",emp.age)
	fmt.Println("SALARY:",emp.salary)
}