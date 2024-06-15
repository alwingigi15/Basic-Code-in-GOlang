package customer

import (
	"fmt"
	
)

func GetCustomer(Id int ) (bool, error) {

	fmt.Println("Get Method Invoked!")

	return true, nil
}
