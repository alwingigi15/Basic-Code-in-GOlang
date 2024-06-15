package main
import ("fmt")

func main(){
	a:=map[string]string{"ID":"ab101","NAME":"ALWIN","GENDER":"MALE","AGE":"18"}
	fmt.Printf("a\t%v\n",a)
	a["NAME"]="ABU"
	a["EDUATION"]="BCA"
	delete(a,"AGE")
	fmt.Println(a)


	b:=make(map[string]string)
	b["NAME"]="ALWIN"
	b["ID"]="ab102"
	b["GENDER"]="MALE"
	b["EDUCATION"]="BCA"
	fmt.Println(b["EDUCATION"])
	fmt.Println(b)
}