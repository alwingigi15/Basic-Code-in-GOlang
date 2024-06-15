package main
import ("fmt")

func main(){
	var x,y,opr int
	fmt.Println("enter 2 number ")
	fmt.Scanln(&x)
	fmt.Scanln(&y)
	fmt.Println("enter the operation")
	fmt.Scanf("%c",&opr)
	
	
	

	switch opr {
	
	case '+':
		
		fmt.Printf("%d+%d=%d",x,y,x+y)
		break
	case '-':
		
		fmt.Printf("%d-%d=%d",x,y,x-y)
		break
	case '/':
		
		fmt.Printf("%d/%d=%d",x,y,x/y)
		break
	case '*':
		
		fmt.Printf("%d*%d=%d",x,y,x*y)
		break
	default:
		fmt.Println("not defined")
		break
	}
}