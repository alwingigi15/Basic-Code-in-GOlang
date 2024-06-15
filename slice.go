package main
import "fmt"

func main(){
	arr1:=[5]int{9,2,8,3,4}
	arr1[2]=10
	myslice:=arr1[1:4]
	fmt.Println("array:",arr1)
	fmt.Printf("myslice:%v\n",myslice)
	fmt.Printf("legth:%d\n",len(myslice))
	fmt.Printf("capacity:%d\n",cap(myslice))

	myslice1:=make([]int ,5,10)
	fmt.Printf("myslice1:%v\n",myslice1)
	fmt.Printf("length:%d\n",len(myslice1))
	fmt.Printf("capcity:%d\n",cap(myslice1))

	myslice2:=[]int{3,4,5,6,7,8,9,2}
	fmt.Printf("myslice2:%d\n",myslice2)
	fmt.Printf("length:%d\n",len(myslice2))
	fmt.Printf("capacity:%d\n",cap(myslice2))

	cars:=[]string{"polo","bmw","benz"}
	fmt.Printf("cars:%v\n",cars)
	fmt.Printf("length:%d\n",len(cars))
	fmt.Printf("capacity:%d\n",cap(cars))
	cars[0]="audi"
	fmt.Println(cars[0])
	fmt.Println(cars[2])

	myslice2=append(myslice2, 10,15)
	fmt.Printf("myslice2:%d\n",myslice2)
	fmt.Printf("length:%d\n",len(myslice2))
	fmt.Printf("capacity:%d\n",cap(myslice2))

	myslice3:=append(myslice1,myslice2...)
	fmt.Printf("myslice3:%d\n",myslice3)
	fmt.Printf("length:%d\n",len(myslice3))
	fmt.Printf("capacity:%d\n",cap(myslice3))

	needs:=myslice2[:len(myslice2)-5]
	ncopy:=make([]int ,len(needs))
	copy(ncopy,needs)
	fmt.Printf("number copy:%d\n",ncopy)
	fmt.Printf("length:%d\n",len(ncopy))
	fmt.Printf("capacity:%d\n",cap(ncopy))

}
