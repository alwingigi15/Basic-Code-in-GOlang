package customer

import "fmt"

func Details(){
	mydata:=map[string]string {"CUSTOMER_NAME":"ALWIN","CUSTOMER_AGE":"22","EMAIL":"alw@gmail.com"}
	for index,dat:=range mydata{
		fmt.Println(index,":",dat)
	}
	
	
}

